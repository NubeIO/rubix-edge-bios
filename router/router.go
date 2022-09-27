package router

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-edge-bios/config"
	"github.com/NubeIO/rubix-edge-bios/constants"
	"github.com/NubeIO/rubix-edge-bios/controller"
	"github.com/NubeIO/rubix-edge-bios/logger"
	"github.com/NubeIO/rubix-edge-bios/model"
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"time"
)

func NotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := fmt.Sprintf("%s %s [%d]: %s", ctx.Request.Method, ctx.Request.URL, 404, "rubix-edge-bios: api not found")
		ctx.JSON(http.StatusNotFound, model.Message{Message: message})
	}
}

func Setup() *gin.Engine {
	engine := gin.New()
	// Set gin access logs
	if viper.GetBool("gin.log.store") {
		fileLocation := fmt.Sprintf("%s/rubix-edge-bios.access.log", config.Config.GetAbsDataDir())
		f, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, constants.Permission)
		if err != nil {
			logger.Logger.Errorf("Failed to create access log file: %v", err)
		} else {
			gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		}
	}
	gin.SetMode(viper.GetString("gin.log.level"))
	engine.NoRoute(NotFound())
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders: []string{
			"X-FLOW-Key", "Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host",
		},
		ExposeHeaders:          []string{"Content-Length"},
		AllowCredentials:       true,
		AllowAllOrigins:        true,
		AllowBrowserExtensions: true,
		MaxAge:                 12 * time.Hour,
	}))

	systemCtl := systemctl.New(false, 30)
	api := controller.Controller{SystemCtl: systemCtl, RubixRegistry: rubixregistry.New(), FileMode: 0755}
	engine.POST("/api/users/login", api.Login)
	engine.GET("/api/users", api.GetUser)
	engine.GET("/api/tokens", api.GetTokens)
	systemApi := engine.Group("/api/system")
	{
		systemApi.GET("/ping", api.Ping)
		systemApi.GET("/device", api.GetDeviceInfo)
	}

	handleAuth := func(c *gin.Context) { c.Next() }
	if config.Config.Auth() {
		// handleAuth = api.HandleAuth() // TODO add back in auth
	}

	apiRoutes := engine.Group("/api", handleAuth)
	apiRoutes.PATCH("/system/device", api.UpdateDeviceInfo)

	appControl := apiRoutes.Group("/systemctl")
	{
		appControl.POST("/enable", api.SystemCtlEnable)
		appControl.POST("/disable", api.SystemCtlDisable)
		appControl.GET("/show", api.SystemCtlShow)
		appControl.POST("/start", api.SystemCtlStart)
		appControl.GET("/status", api.SystemCtlStatus)
		appControl.POST("/stop", api.SystemCtlStop)
		appControl.POST("/reset-failed", api.SystemCtlResetFailed)
		appControl.POST("/daemon-reload", api.SystemCtlDaemonReload)
		appControl.POST("/restart", api.SystemCtlRestart)
		appControl.POST("/mask", api.SystemCtlMask)
		appControl.POST("/unmask", api.SystemCtlUnmask)
		appControl.GET("/state", api.SystemCtlState)
		appControl.GET("/is-enabled", api.SystemCtlIsEnabled)
		appControl.GET("/is-active", api.SystemCtlIsActive)
		appControl.GET("/is-running", api.SystemCtlIsRunning)
		appControl.GET("/is-failed", api.SystemCtlIsFailed)
		appControl.GET("/is-installed", api.SystemCtlIsInstalled)
	}

	files := apiRoutes.Group("/files")
	{
		files.GET("/exists", api.FileExists)            // needs to be a file
		files.GET("/walk", api.WalkFile)                // similar as find in linux command
		files.GET("/list", api.ListFiles)               // list all files and folders
		files.POST("/create", api.CreateFile)           // create file only
		files.POST("/copy", api.CopyFile)               // copy either file or folder
		files.POST("/rename", api.RenameFile)           // rename either file or folder
		files.POST("/move", api.MoveFile)               // move file only
		files.POST("/upload", api.UploadFile)           // upload single file
		files.POST("/download", api.DownloadFile)       // download single file
		files.GET("/read", api.ReadFile)                // read single file
		files.PUT("/write", api.WriteFile)              // write single file
		files.DELETE("/delete", api.DeleteFile)         // delete single file
		files.DELETE("/delete-all", api.DeleteAllFiles) // deletes file or folder
	}

	dirs := apiRoutes.Group("/dirs")
	{
		dirs.GET("/exists", api.DirExists)  // needs to be a folder
		dirs.POST("/create", api.CreateDir) // create folder
	}

	zip := apiRoutes.Group("/zip")
	{
		zip.POST("/unzip", api.Unzip)
		zip.POST("/zip", api.ZipDir)
	}

	user := apiRoutes.Group("/users")
	{
		user.PUT("", api.UpdateUser)
	}

	token := apiRoutes.Group("/tokens")
	{
		token.POST("/generate", api.GenerateToken)
		token.PUT("/:uuid/block", api.BlockToken)
		token.PUT("/:uuid/regenerate", api.RegenerateToken)
		token.DELETE("/:uuid", api.DeleteToken)
	}

	return engine
}

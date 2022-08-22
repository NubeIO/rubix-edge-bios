package router

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-edge/controller"
	dbase "github.com/NubeIO/rubix-edge/database"
	"github.com/NubeIO/rubix-edge/pkg/config"
	"github.com/NubeIO/rubix-edge/pkg/logger"
	"github.com/NubeIO/rubix-edge/service/apps"
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"time"
)

func NotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := fmt.Sprintf("%s %s [%d]: %s", ctx.Request.Method, ctx.Request.URL, 404, "rubix-edge: api not found")
		ctx.JSON(http.StatusNotFound, controller.Message{Message: message})
	}
}

func Setup(db *gorm.DB) *gin.Engine {
	engine := gin.New()
	// Set gin access logs
	if viper.GetBool("gin.log.store") {
		fileLocation := fmt.Sprintf("%s/edge.access.log", config.Config.GetAbsDataDir())
		f, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
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

	appDB := &dbase.DB{
		DB: db,
	}
	edgeSystem := system.New(&system.System{})
	rubixApps, _ := apps.New(&apps.EdgeApps{App: &installer.App{}})
	api := controller.Controller{DB: appDB, Rubix: rubixApps, System: edgeSystem}
	engine.POST("/api/users/login", api.Login)

	handleAuth := func(c *gin.Context) { c.Next() }
	apiPublicRoutes := engine.Group("/api", handleAuth)

	public := apiPublicRoutes.Group("/public") // THESE ARE PUBLIC APIs
	{
		public.GET("/device", api.GetDeviceProduct)
	}

	if config.Config.Auth() {
		//handleAuth = api.HandleAuth() // TODO add back in auth
	}

	apiRoutes := engine.Group("/api", handleAuth)
	apiProxyRoutes := engine.Group("/ff", handleAuth)
	apiProxyRoutes.Any("/*proxyPath", api.FFProxy) // FLOW-FRAMEWORK PROXY

	deviceInfo := apiRoutes.Group("/device")
	{
		deviceInfo.GET("/", api.GetDeviceInfo)
		deviceInfo.PATCH("/", api.UpdateDeviceInfo)
		deviceInfo.DELETE("/", api.DropDeviceInfo)
	}

	edgeApps := apiRoutes.Group("/apps")
	{
		edgeApps.GET("/", api.ListApps)
		edgeApps.GET("/services", api.ListAppsAndService)
		edgeApps.GET("/services/nube", api.ListNubeServices)
		edgeApps.POST("/add", api.AddUploadApp)
		edgeApps.POST("/service/upload", api.UploadService)
		edgeApps.POST("/service/install", api.InstallService)
		edgeApps.DELETE("/", api.UninstallApp) // uninstall an app
	}

	appControl := apiRoutes.Group("/apps/control")
	{
		appControl.POST("/action", api.CtlAction)              // start, stop
		appControl.POST("/action/mass", api.ServiceMassAction) // mass operation start, stop
		appControl.POST("/status", api.CtlStatus)              // isRunning, isInstalled and so on
		appControl.POST("/status/mass", api.ServiceMassStatus) // mass isRunning, isInstalled and so on
	}

	appBackups := apiRoutes.Group("/backup")
	{
		appBackups.POST("/restore/full", api.RestoreBackup)
		appBackups.POST("/restore/app", api.RestoreAppBackup)
		appBackups.POST("/run/full", api.FullBackUp)
		appBackups.POST("/run/app", api.BackupApp)
		appBackups.GET("/list/full", api.ListFullBackups)
		appBackups.GET("/list/apps", api.ListAppBackupsDirs)
		appBackups.GET("/list/app", api.ListBackupsByApp)
	}

	systemTime := apiRoutes.Group("/time")
	{
		systemTime.GET("/", api.SystemTime)
		systemTime.POST("/", api.SetSystemTime)
	}

	systemTimeZone := apiRoutes.Group("/timezone")
	{
		systemTimeZone.GET("/", api.GetHardwareTZ)
		systemTimeZone.POST("/", api.UpdateTimezone)
		systemTimeZone.GET("/list", api.GetTimeZoneList)
		systemTimeZone.POST("/config", api.GenerateTimeSyncConfig)
	}

	systemApi := apiRoutes.Group("/system")
	{
		systemApi.GET("/ping", api.Ping)
		systemApi.GET("/product", api.GetProduct)
		systemApi.POST("/scanner", api.RunScanner)
	}

	networking := apiRoutes.Group("/networking")
	{
		networking.GET("/", api.Networking)
		networking.GET("/interfaces", api.GetInterfacesNames)
		networking.GET("/internet", api.InternetIP)

	}

	networks := apiRoutes.Group("/networking/networks")
	{
		networks.POST("/restart", api.RestartNetworking)
		networks.POST("/up", api.InterfaceUp)
		networks.POST("/down", api.InterfaceDown)
	}

	networkAddress := apiRoutes.Group("/networking/interfaces")
	{
		networkAddress.POST("/exists", api.DHCPPortExists)
		networkAddress.POST("/auto", api.DHCPSetAsAuto)
		networkAddress.POST("/static", api.DHCPSetStaticIP)
	}

	networkFirewall := apiRoutes.Group("/networking/firewall")
	{
		networkFirewall.GET("/", api.UWFStatusList)
		networkFirewall.POST("/status", api.UWFStatus)
		networkFirewall.POST("/active", api.UWFActive)
		networkFirewall.POST("/enable", api.UWFEnable)
		networkFirewall.POST("/disable", api.UWFDisable)
		networkFirewall.POST("/port/open", api.UWFOpenPort)
		networkFirewall.POST("/port/close", api.UWFClosePort)
	}

	files := apiRoutes.Group("/files")
	{
		files.GET("/walk", api.WalkFile)
		files.GET("/list", api.ListFiles) // /api/files/list?file=/data
		files.POST("/rename", api.RenameFile)
		files.POST("/copy", api.CopyFile)
		files.POST("/move", api.MoveFile)
		files.POST("/upload", api.UploadFile)
		files.POST("/download", api.DownloadFile)
		files.DELETE("/delete", api.DeleteFile)
		files.DELETE("/delete/all", api.DeleteAllFiles)
	}

	dirs := apiRoutes.Group("/dirs")
	{
		dirs.POST("/create", api.CreateDir)
		dirs.POST("/copy", api.CopyDir)
		dirs.DELETE("/delete", api.DeleteDir)
	}

	zip := apiRoutes.Group("/zip")
	{
		zip.POST("/unzip", api.Unzip)
		zip.POST("/zip", api.ZipDir)
	}

	user := apiRoutes.Group("/users")
	{
		user.PUT("", api.UpdateUser)
		user.GET("", api.GetUser)
	}

	token := apiRoutes.Group("/tokens")
	{
		token.GET("", api.GetTokens)
		token.POST("/generate", api.GenerateToken)
		token.PUT("/:uuid/block", api.BlockToken)
		token.PUT("/:uuid/regenerate", api.RegenerateToken)
		token.DELETE("/:uuid", api.DeleteToken)
	}

	return engine
}

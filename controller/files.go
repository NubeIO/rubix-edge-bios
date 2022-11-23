package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/rubix-edge-bios/model"
	"github.com/gin-gonic/gin"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

type FileExistence struct {
	File   string `json:"file"`
	Exists bool   `json:"exists"`
}

type UploadResponse struct {
	Destination string `json:"destination"`
	File        string `json:"file"`
	Size        string `json:"size"`
	UploadTime  string `json:"upload_time"`
}

func (inst *Controller) FileExists(c *gin.Context) {
	file := c.Query("file")
	exists := fileutils.FileExists(file)
	fileExistence := FileExistence{File: file, Exists: exists}
	responseHandler(fileExistence, nil, c)
}

func (inst *Controller) WalkFile(c *gin.Context) {
	path_ := c.Query("path")
	files := make([]string, 0)
	err := filepath.WalkDir(path_, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		files = append(files, p)
		return nil
	})
	responseHandler(files, err, c)
}

func (inst *Controller) ListFiles(c *gin.Context) {
	files, err := inst.listFiles(c.Query("path"))
	responseHandler(files, err, c)
}

func (inst *Controller) listFiles(_path string) ([]fileutils.FileDetails, error) {
	fileInfo, err := os.Stat(_path)
	dirContent := make([]fileutils.FileDetails, 0)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(_path)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			dirContent = append(dirContent, fileutils.FileDetails{Name: file.Name(), IsDir: file.IsDir()})
		}
	} else {
		return nil, errors.New("it needs to be a directory, found a file")
	}
	return dirContent, nil
}

func (inst *Controller) CreateFile(c *gin.Context) {
	file := c.Query("file")
	if file == "" {
		responseHandler(nil, errors.New("file can not be empty"), c)
		return
	}
	_, err := fileutils.CreateFile(file, os.FileMode(inst.FileMode))
	responseHandler(model.Message{Message: fmt.Sprintf("created file: %s", file)}, err, c)
}

func (inst *Controller) CopyFile(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	if from == "" || to == "" {
		responseHandler(nil, errors.New("from and to names can not be empty"), c)
		return
	}
	err := fileutils.Copy(from, to)
	responseHandler(model.Message{Message: "copied successfully"}, err, c)
}

func (inst *Controller) RenameFile(c *gin.Context) {
	oldPath := c.Query("old_path")
	newPath := c.Query("new_path")
	if oldPath == "" || newPath == "" {
		responseHandler(nil, errors.New("old_path & new_path names can not be empty"), c)
		return
	}
	err := os.Rename(oldPath, newPath)
	responseHandler(model.Message{Message: "renamed successfully"}, err, c)
}

func (inst *Controller) MoveFile(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	if from == "" || to == "" {
		responseHandler(nil, errors.New("from and to names can not be empty"), c)
		return
	}
	if from == to {
		responseHandler(nil, errors.New("from and to names are same"), c)
		return
	}
	err := os.Rename(from, to)
	responseHandler(model.Message{Message: "moved successfully"}, err, c)
}

func (inst *Controller) DownloadFile(c *gin.Context) {
	path_ := c.Query("path")
	fileName := c.Query("file")
	c.FileAttachment(fmt.Sprintf("%s/%s", path_, fileName), fileName)
}

// UploadFile
// curl -X POST http://localhost:1661/api/files/upload?destination=/data/ -F "file=@/home/user/Downloads/bios-master.zip" -H "Content-Type: multipart/form-data"
func (inst *Controller) UploadFile(c *gin.Context) {
	now := time.Now()
	destination := c.Query("destination")
	file, err := c.FormFile("file")
	resp := &UploadResponse{}
	if err != nil || file == nil {
		responseHandler(resp, err, c)
		return
	}
	if found := fileutils.DirExists(destination); !found {
		responseHandler(nil, errors.New(fmt.Sprintf("destination not found %s", destination)), c)
		return
	}
	toFileLocation := path.Join(destination, filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, toFileLocation); err != nil {
		responseHandler(resp, err, c)
		return
	}
	if err := os.Chmod(toFileLocation, os.FileMode(inst.FileMode)); err != nil {
		responseHandler(resp, err, c)
		return
	}
	size, err := fileutils.GetFileSize(toFileLocation)
	if err != nil {
		responseHandler(resp, err, c)
		return
	}
	resp = &UploadResponse{
		Destination: toFileLocation,
		File:        file.Filename,
		Size:        size.String(),
		UploadTime:  TimeTrack(now),
	}
	responseHandler(resp, nil, c)
}

func (inst *Controller) ReadFile(c *gin.Context) {
	file := c.Query("file")
	if file == "" {
		responseHandler(nil, errors.New("file can not be empty"), c)
		return
	}
	found := fileutils.FileExists(file)
	if !found {
		responseHandler(nil, errors.New(fmt.Sprintf("file not found: %s", file)), c)
		return
	}
	c.File(file)
}

type WriteFile struct {
	Data string `json:"data"`
}

type WriteFormatFile struct {
	FilePath     string      `json:"path"`
	Body         interface{} `json:"body"`
	BodyAsString string      `json:"body_as_string"`
}

func (inst *Controller) WriteFile(c *gin.Context) {
	file := c.Query("file")
	if file == "" {
		responseHandler(nil, errors.New("file can not be empty"), c)
		return
	}
	var m *WriteFile
	err := c.ShouldBindJSON(&m)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	err = fileutils.WriteFile(file, m.Data, fs.FileMode(inst.FileMode))
	responseHandler(model.Message{Message: fmt.Sprintf("wrote the file: %s", file)}, err, c)
}

func (inst *Controller) DeleteFile(c *gin.Context) {
	file := c.Query("file")
	if !fileutils.FileExists(file) {
		responseHandler(nil, errors.New(fmt.Sprintf("file doesn't exist: %s", file)), c)
		return
	}
	err := fileutils.Rm(file)
	responseHandler(model.Message{Message: fmt.Sprintf("deleted file: %s", file)}, err, c)
}

func (inst *Controller) DeleteAllFiles(c *gin.Context) {
	filePath := c.Query("path")
	if !fileutils.FileOrDirExists(filePath) {
		responseHandler(nil, errors.New(fmt.Sprintf("doesn't exist: %s", filePath)), c)
		return
	}
	err := fileutils.RemoveAllFiles(filePath)
	responseHandler(model.Message{Message: fmt.Sprintf("deleted path: %s", filePath)}, err, c)
}

func TimeTrack(start time.Time) (out string) {
	elapsed := time.Since(start)
	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	out = fmt.Sprintf("%s took %s", name, elapsed)
	return out
}

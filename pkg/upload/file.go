package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/Apriil15/blog-server/global"
	"github.com/Apriil15/blog-server/pkg/util"
)

type FileType int

const (
	TypeImage FileType = iota + 1
)

// Get encoded file name
// because i don't want the origin file name showed directly
func GetFileName(name string) string {
	extension := GetFileExtension(name)
	fileName := strings.TrimSuffix(name, extension) // trim extension: test.jpg -> test
	fileName = util.EncodeMD5(fileName)

	return fileName + extension
}

// Get file's extension
func GetFileExtension(name string) string {
	return path.Ext(name)
}

// Get upload save path from config
func GetUploadSavePath() string {
	return global.AppSetting.UploadSavePath
}

// Check whether save path is exist
func CheckSavePath(destination string) bool {
	_, err := os.Stat(destination) // get FileInfo
	return os.IsNotExist(err)
}

// Check extension is supported.
// return true means extension is valid, false means invalid
func CheckContainExtension(fileType FileType, name string) bool {
	extension := GetFileExtension(name)
	extension = strings.ToUpper(extension) // maybe extension with upper, lower, or mix, so i use ToUpper to uniform

	switch fileType {
	case TypeImage:
		for _, allowExtension := range global.AppSetting.UploadImageAllowExtensions {
			if extension == strings.ToUpper(allowExtension) {
				return true
			}
		}
	}
	return false
}

// Check image size.
// return true means over maximum size limit
func CheckMaxSize(fileType FileType, file multipart.File) bool {
	content, _ := ioutil.ReadAll(file)
	size := len(content)

	switch fileType {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// Check permission
func CheckPermission(destination string) bool {
	_, err := os.Stat(destination)
	return os.IsPermission(err)
}

// Create save path
func CreateSavePath(destination string, permission os.FileMode) error {
	err := os.MkdirAll(destination, permission)
	if err != nil {
		return err
	}
	return nil
}

// Save file to destinationPath
func SaveFile(file *multipart.FileHeader, destination string) error {
	source, err := file.Open()
	if err != nil {
		return err
	}
	defer source.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, source) // copy from source to destinationFile
	return err
}

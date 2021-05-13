package upload

import (
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

// Get upload save path
func GetUploadSavePath() string {
	return global.AppSetting.UploadSavePath
}

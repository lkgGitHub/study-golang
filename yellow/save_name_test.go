package yellow

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSaveName(t *testing.T) {
	//db, err := gorm.Open(sqlite.Open(filepath.Join("yellow.db")), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(filepath.Join(reNameDir, "yellow.db")), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	//
	db.AutoMigrate(&Video{})

	fileMap := make(map[string]struct{}, 1000)
	files, _ := GetAllFile(reNameDir)
	fmt.Println("目录下的所有文件如下")
	for i := 0; i < len(files); i++ {
		fileName := filepath.Base(files[i])
		if strings.HasPrefix(fileName, ".") ||
			strings.HasSuffix(fileName, ".jpg") || strings.HasSuffix(fileName, ".png") {
			continue
		}

		fileMap[fileName] = struct{}{}
		fmt.Println(fileName)
	}
	for k, _ := range fileMap {
		video := &Video{Name: k}
		if err = db.Save(video).Error; err != nil {
			t.Error(err)
			continue
		}
	}

}

// GetAllFile 递归获取指定目录下的所有文件名
func GetAllFile(pathname string) ([]string, error) {
	var result []string

	fis, err := os.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullName := filepath.Join(pathname, fi.Name())

		// $RECYCLE.BIN 是系统重要的隐藏文件，一般存在于磁盘根目录下。是系统“回收站”在每一个磁盘上的链接文件夹，用于保存磁盘上删除的文件或者文件夹信息，我们恢复误删除到回收站中的文件或者文件夹时大有用处。
		if strings.HasPrefix(fi.Name(), ".") || strings.Contains(fi.Name(), "$RECYCLE.BIN") {
			continue
		}

		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := GetAllFile(fullName)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullName, err)
				return result, err
			}
			result = append(result, temp...)
		} else {
			result = append(result, fullName)
		}
	}

	return result, nil
}

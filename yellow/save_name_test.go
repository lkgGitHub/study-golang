package yellow

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestSaveName(t *testing.T) {
	//db, err := gorm.Open(sqlite.Open(filepath.Join("yellow.db")), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(filepath.Join(reNameDir, "yellow.db")), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	db.AutoMigrate(&Video{})

	fileMap := make(map[string]struct{}, 1000)
	files, _ := getAllFile(reNameDir)
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

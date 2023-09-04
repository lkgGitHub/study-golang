package yellow

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"testing"
)

const reNameDir = ""

func TestRename(t *testing.T) {
	listFiles(reNameDir)
}

func listFiles(dirname string) {
	fileInfos, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		if !strings.HasSuffix(fi.Name(), ".mp4") {
			continue
		}
		newName, ok := getNewName(fi.Name())
		if ok {
			err = os.Rename(path.Join(dirname, fi.Name()), path.Join(dirname, newName+".mp4"))
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		}
	}
}

func getNewName(name string) (string, bool) {
	if len(name) < 6 {
		return name, false
	}

	prefix := name[0:6]
	nameTmp := strings.TrimPrefix(name, prefix)
	index := strings.Index(nameTmp, prefix)
	if index > 0 {
		return name[0 : index+5], true
	}

	return name, false
}

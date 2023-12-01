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
	dirname := reNameDir

	fileInfos, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		if !strings.HasSuffix(strings.ToLower(fi.Name()), ".mp4") {
			continue
		}
		if newName, ok := deleteRepetitionCharacter(fi.Name()); ok {
			fmt.Println(fi.Name(), "->", newName)
			err = os.Rename(path.Join(dirname, fi.Name()), path.Join(dirname, newName+".mp4"))
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		}
		if newName, ok := deleteInvalid(fi.Name()); ok {
			fmt.Println(fi.Name(), "->", newName)
			err = os.Rename(path.Join(dirname, fi.Name()), path.Join(dirname, newName))
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		}
	}
}

func deleteRepetitionCharacter(name string) (string, bool) {
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

func deleteInvalid(name string) (string, bool) {
	rename := false

	if strings.Contains(name, ".mp4.mp4") {
		name = strings.ReplaceAll(name, ".mp4.mp4", ".mp4")
		rename = true
	}

	if strings.Contains(name, "-.mp4") {
		name = strings.ReplaceAll(name, "-.mp4", ".mp4")
		rename = true
	}

	for _, str := range deleteStrSlice {
		if strings.Contains(name, str) {
			name = strings.ReplaceAll(name, str, "")
			rename = true
		}
	}

	for _, substr := range deleteSuffixSlice {
		if index := strings.Index(name, substr); index > 0 {
			nameSlice := strings.Split(name, ".")
			suffix := nameSlice[len(nameSlice)-1]
			name = name[:index]
			name = strings.TrimSpace(name)
			name = strings.TrimSuffix(name, "-")
			name = name + "." + suffix
			rename = true
		}
	}

	return name, rename
}

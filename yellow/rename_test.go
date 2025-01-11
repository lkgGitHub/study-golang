package yellow

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestRename(t *testing.T) {
	dirname := reNameDir

	for i := 0; i < 2; i++ {
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
}

func TestCheckName(t *testing.T) {
	files, err := getAllFile(reNameDir)
	if err != nil {
		t.Error(err.Error())
		return
	}

	// 实际所有文件的 map
	fileMap := make(map[string]string, len(files))
	for _, f := range files {
		n := filepath.Base(f)
		if v, ok := fileMap[n]; !ok {
			fileMap[n] = f
		} else {
			// 重复文件名打印出来
			fmt.Printf("duplicate name: %s \n %s\n %s \n", n, v, f)
		}
	}

	body, err := os.ReadFile("yellow.txt")
	if err != nil {
		panic(err)
	}

	// 已保存所有文件的 map
	bodyMap := make(map[string]struct{}, 3000)
	for _, line := range strings.Split(string(body), "\n") {
		bodyMap[line] = struct{}{}
	}

	var content strings.Builder
	for k := range fileMap {
		if _, ok := bodyMap[k]; !ok {
			content.WriteString(k)
			content.WriteString("\n")
		}
	}

	file, err := os.OpenFile("yellow.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	write := bufio.NewWriter(file)
	if _, err = write.WriteString(content.String()); err != nil {
		panic(err)
	}
	//Flush将缓存的文件真正写入到文件中
	if err = write.Flush(); err != nil {
		panic(err)
	}

}

// 递归获取指定目录下的所有文件名
func getAllFile(pathname string) ([]string, error) {
	result := make([]string, 0, 1000)

	fis, err := os.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		if strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		if strings.Contains(fi.Name(), "22,0622f7b76eba00") {
			fmt.Println("aaaa")
		}
		if strings.HasSuffix(fi.Name(), ".jpg") || strings.HasSuffix(fi.Name(), ".jpeg") {
			continue
		}

		fullName := pathname + "/" + fi.Name()
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

func deleteRepetitionCharacter(name string) (string, bool) {
	if len(name) < 9 {
		return name, false
	}

	p := name[0:9]
	nameTmp := strings.TrimPrefix(name, p)
	index := strings.Index(nameTmp, p)
	if index > 0 {
		return name[0 : index+9], true
	}

	return name, false
}

func deleteInvalid(name string) (string, bool) {
	rename := false

	if strings.Contains(name, ".mp4.mp4") {
		name = strings.ReplaceAll(name, ".mp4.mp4", ".mp4")
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

	if strings.Contains(name, "-.mp4") {
		name = strings.ReplaceAll(name, "-.mp4", ".mp4")
		rename = true
	}

	return name, rename
}

func TestPrefix(t *testing.T) {
	dirname := reNameDir
	fileInfos, err := os.ReadDir(reNameDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		if !strings.HasSuffix(strings.ToLower(fi.Name()), ".mp4") {
			continue
		}
		if strings.Contains(fi.Name(), prefix) && !strings.HasPrefix(fi.Name(), prefix) {
			newName := strings.Replace(fi.Name(), prefix, "", -1)
			newName, _ = deleteInvalid(newName)
			err = os.Rename(path.Join(dirname, fi.Name()), path.Join(dirname, prefix+"-"+newName))
			if err != nil {
				fmt.Println(err.Error())
				continue
			} else {
				fmt.Println(fi.Name(), "->", prefix+"-"+newName)
			}
		}
	}
}

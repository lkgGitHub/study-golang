package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const dirPath = "./yulan/testData"
const hf = 3 // 段的个数

//var df = ""     // 文件名
var dataFileCount int    // 文件个数，不固定
var serverCount int = 10 // 服务器个数，不固定

var words = []string{"kafka", "Kubernetes", "PostgreSQL"}
var wordToFile = make(map[string][]string)

// 标识符三元组，{文件名, 文件中第几段, 服务器id}
var identifiersMap = make(map[string]*Identifier)

type Identifier struct {
	fileID    string
	paragraph int
	serverID  int
}

func main() {
	//currentPath, _ := os.Getwd()
	//filepath.Join(currentPath, dirPath)
	dirs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("read dir failed. error:", err.Error())
		os.Exit(1)
	}
	dataFileCount = len(dirs)
	for i, df := range dirs {
		contentByte, err := os.ReadFile(filepath.Join(dirPath, df.Name()))
		if err != nil {
			fmt.Println("read file error:", err.Error())
			continue
		}
		content := string(contentByte)
		contentArr := strings.Split(content, "\n")
		segmentCap := len(contentArr) / hf
		for hji := 0; hji < hf; hji++ {
			var serverID int
			if hji != hf-1 {
				serverID, err = sendToServer(df.Name()+strconv.Itoa(i),
					strings.Join(contentArr[hji*segmentCap:(hji+1)*segmentCap], " "))
			} else {
				serverID, err = sendToServer(df.Name()+strconv.Itoa(i),
					strings.Join(contentArr[hji*segmentCap:], " "))
			}
			if err != nil {
				fmt.Println("send to server error:", err.Error())
				continue
			}
			identifier := &Identifier{
				fileID:    df.Name(),
				paragraph: i,
				serverID:  serverID,
			}
			//identifiers = append(identifiers, identifier)
			identifiersMap[df.Name()+strconv.Itoa(i)] = identifier
		}
	}

	for _, df := range dirs {
		contentByte, err := os.ReadFile(filepath.Join(dirPath, df.Name()))
		if err != nil {
			fmt.Println("read file error:", err.Error())
			continue
		}
		for _, word := range words {
			if strings.Contains(string(contentByte), word) {
				wordToFile[word] = append(wordToFile[word], df.Name())
			}
		}
	}

	println("dataFileCount:", dataFileCount)
	fmt.Printf("wordToFile: %v \n", wordToFile)
	fmt.Printf("identifiersMap: %v \n", identifiersMap)

}

// 每段随机分配到服务器上，获取服务器标识
func getServerID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(serverCount)
}

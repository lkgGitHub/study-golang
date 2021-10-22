package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
	"io/ioutil"
	"pare/start"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

//源的个数，份数
const Group_num int = 3

//单词数量//单词数量
const WordsCount int = 2000

var Words [WordsCount]string

const DirPath = "./data_1000" //从data文件加中读取文件D
//文件数量
const DataFileCount int = 1000

//存储文件关键字和文件标识符对

//索引
var ListL = make(map[string]string)
var ListL_add = make(map[string]string)
var ListL_del []string

//关键字和文件标识符对
var wordToFile = make(map[string][]string)

//包含某关键字文件个数计数
var Tw = make(map[string]int)
var T_add = make(map[string]int)

func IndexGen() {
	var size_index int = 0
	for _, word := range Words {
		Tw[word] = 0
		T_add[word] = 0
		//提取关键字处理
		for j := 0; j < DataFileCount; j++ {
			contentByte, err := ioutil.ReadFile(filepath.Join(DirPath, strconv.Itoa(j)+".txt"))
			if err != nil {
				fmt.Println("read file error:", err.Error())
				continue
			}
			if strings.Contains(string(contentByte), word) {
				wordToFile[word] = append(wordToFile[word], strconv.Itoa(j)+".txt")
			}
		}
	}

	wg := &sync.WaitGroup{}
	for _, word := range Words {
		go func(word string, wg *sync.WaitGroup) {
			wg.Add(1)
			defer wg.Done()
			K_1 := PRF(K_a, "1"+word)
			K_2 := PRF(K_a, "2"+word)
			c := Tw[word]
			temp_len := len(wordToFile[word])
			temp_con := wordToFile[word]
			Tw[word] = c + temp_len
			for j := 0; j < temp_len; j++ {
				l := PRF(K_1, strconv.Itoa(c))
				d, err := start.AesEncrypt([]byte(temp_con[j]), K_2)
				if err != nil {
					fmt.Println("索引加密失败", err)
					return
				}
				c++
				ListL[string(l)] = string(d)
				size_index += (len(l) + len(d))
			}
		}(word, wg)
	}
	wg.Wait()
	fmt.Println("size_index", size_index)
}

//PRF随机函数--------F
func PRF(key []byte, plainText string) []byte {
	//fmt.Println("key:", key)
	//plaintext := []byte(plainText)
	var myhash hash.Hash
	var hashText_temp [][]byte //切片
	////var hashText []byte        //切片
	len_key := len(key)
	//fmt.Println("len_key:",len_key)
	len_hash := 0
	for i := 0; len_hash <= len_key; i++ {
		// 1.创建哈希接口, 需要指定使用的哈希算法, 和秘钥
		myhash = hmac.New(sha256.New, key)
		// 2. 给哈希对象添加数据
		myhash.Write([]byte(plainText + strconv.Itoa(i)))
		// 3. 计算散列值
		////hashText_temp[i] = myhash.Sum(nil)
		hashText_temp = append(hashText_temp, myhash.Sum(nil))
		//fmt.Println("myhash.Sum(nil):", myhash.Sum(nil))
		////len_hash += len(hashText_temp[i])
		len_hash += len(myhash.Sum(nil))
	}

	sep := []byte("")
	hashText := bytes.Join(hashText_temp, sep) //多个[]byte结合

	//fmt.Println("PRF-hashText:", hashText)
	//如果hashText的长度超过key或者不够key
	if len_hash > len_key {
		hashText = hashText[:len_key]
	}
	//fmt.Println("PRF-hashText:", hashText)
	//fmt.Println("len(hashText):", len(hashText))
	return hashText
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	ErrorReply  = '-'
	StatusReply = '+'
	IntReply    = ':'
	StringReply = '$'
	ArrayReply  = '*'
)
type RedisError string
const Nil = RedisError("redis: nil")
func (e RedisError) Error() string { return string(e) }


type Reader struct {
	rd   *bufio.Reader
	_buf []byte
}

// 处理redis返回信息
func ParseReply(reply []byte) (interface{}, error) {
	responseType := reply[0]

	switch responseType {
	case ErrorReply:
		return string(reply[1:len(reply)-2]), RedisError(string(reply[1:len(reply)-2]))
	case StatusReply:
		return parseStatusReply(reply)
	case IntReply:
		return parseInt(reply)
	case StringReply:
		return parseStringReply(reply)
	case ArrayReply:
		return parseArrayLen(reply)
	default:
		return nil, RedisError("proto wrong!")
	}
}

//func parseArrayReply(reply []byte) (string, error) {
//	statusByte := reply[1:]
//
//	// 获取响应文本第一行标示的响应字符串长度
//	pos := 0
//
//	for _, v := range statusByte {
//		if v == '\r' {
//			break
//		}
//		pos++
//	}
//
//	arrayByte :=  statusByte[pos+2:]
//
//}
func parseArrayLen(line []byte) (int64, error) {
	if isNilReply(line) {
		return 0, Nil
	}
	//return strconv.Atoi(string(line[1:]))
	return strconv.ParseInt(string(line[1:]), 10, 64)
}

func isNilReply(b []byte) bool {
	return len(b) == 3 &&
		(b[0] == StringReply || b[0] == ArrayReply) &&
		b[1] == '-' && b[2] == '1'
}

func parseInt(reply []byte) (interface{}, error) {
	statusByte := reply[1:]
	pos := 0
	for _, v := range statusByte {
		if v == '\r' {
			break
		}
		pos++
	}
	status := statusByte[:pos]

	return strconv.ParseInt(string(status), 10, 64)
}

// 处理状态响应
func parseStatusReply(reply []byte) (interface{}, error) {
	statusByte := reply[1:]

	pos := 0
	for _, v := range statusByte {
		if v == '\r' {
			break
		}
		pos++
	}
	status := statusByte[:pos]

	return string(status), nil
}

// 处理主体响应
func parseStringReply(reply []byte) (interface{}, error) {
	statusByte := reply[1:]

	// 获取响应文本第一行标示的响应字符串长度
	pos := 0

	for _, v := range statusByte {
		if v == '\r' {
			break
		}
		pos++
	}

	strlen, err := strconv.Atoi(string(statusByte[:pos]))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if strlen == -1 {
		return "nil", nil
	}
	nextLinePost := 1
	for _, v := range statusByte {
		if v == '\n' {
			break
		}
		nextLinePost++
	}

	result := string(statusByte[nextLinePost:nextLinePost+strlen])
	return result, nil
}

func main() {
	statusReply, _ := ParseReply([]byte("+OK\r\n"))
	intReply, _ := ParseReply([]byte(":1000\r\n"))
	errorReply, _ := ParseReply([]byte("-Error message\r\n"))
	stringReply, _ := ParseReply([]byte("$5\r\nhello\r\n"))
	arrayReply, _ := ParseReply([]byte("*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"))
	fmt.Println("statusReply: ", statusReply)
	fmt.Println("intReply: ", intReply)
	fmt.Println("errorReply: ", errorReply)
	fmt.Println("stringReply: ", stringReply)
	fmt.Println("arrayReply: ", arrayReply)
}
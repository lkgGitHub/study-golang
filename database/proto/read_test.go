package proto_test

import (
	"bytes"
	"fmt"
	"proto"
	"strings"
	"testing"
)

func TestParseReply(t *testing.T)  {
	replyMap := make(map[string]interface{})
	replyMap["+OK\r\n"] = "OK"
	replyMap["-Error message\r\n"] = nil
	replyMap[":100\r\n"] = int64(100)
	replyMap["$5\r\nhello\r\n"] = "hello"
	replyMap["*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"] = [2]interface{}{"hello",  "world"}

	for key, value := range replyMap{
		var reply interface{}
		if strings.HasPrefix(key, "*"){
			reply = ParseReply(t, key, multiBulkParse)
		}else{
			reply = ParseReply(t, key, nil)
		}
		if value == reply{
			fmt.Println("success: ", value)
		}else {
			fmt.Printf("error: value: %s, reply: %s \r\n", value, reply)
		}
	}

	//ParseReply(t, "+OK\r\n", nil)
	//ParseReply(t, ":100\r\n", nil)
	//ParseReply(t, "-Error message\r\n", nil)
	//ParseReply(t, "$5\r\nhello\r\n", nil)
	//ParseReply(t, "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n", multiBulkParse)
}


func BenchmarkReader_ParseReply_Status(b *testing.B) {
	benchmarkParseReply(b, "+OK\r\n", nil, false)
}

func BenchmarkReader_ParseReply_Int(b *testing.B) {
	benchmarkParseReply(b, ":100\r\n", nil, false)
}

func BenchmarkReader_ParseReply_Error(b *testing.B) {
	benchmarkParseReply(b, "-Error message\r\n", nil, true)
}

func BenchmarkReader_ParseReply_String(b *testing.B) {
	benchmarkParseReply(b, "$5\r\nhello\r\n", nil, false)
}

func BenchmarkReader_ParseReply_Slice(b *testing.B) {
	benchmarkParseReply(b, "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n", multiBulkParse, false)
}

func ParseReply(t *testing.T, reply string, m proto.MultiBulkParse) interface{} {
	buf := new(bytes.Buffer)
	buf.WriteString(reply)
	p := proto.NewReader(buf)
	replyStr, err := p.ReadReply(m)
	//fmt.Println(replyStr)
	if err != nil {
		//t.Errorf(err.Error())
	}
	return replyStr
}

func benchmarkParseReply(b *testing.B, reply string, m proto.MultiBulkParse, wanterr bool) {
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		buf.WriteString(reply)
	}
	p := proto.NewReader(buf)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		replyStr, err := p.ReadReply(m)
		fmt.Println(replyStr)
		if !wanterr && err != nil {
			b.Fatal(err)
		}
	}
}

func multiBulkParse(p *proto.Reader, n int64) (interface{}, error) {
	vv := make([]interface{}, 0, n)
	for i := int64(0); i < n; i++ {
		v, err := p.ReadReply(multiBulkParse)
		if err != nil {
			return nil, err
		}
		vv = append(vv, v)
	}
	return vv, nil
}

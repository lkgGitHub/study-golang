package proto_test

import (
	"bytes"
	"fmt"
	"proto"
	"strings"
	"testing"
)
/*
For Simple Strings the first byte of the reply is "+"
For Errors the first byte of the reply is "-"
For Integers the first byte of the reply is ":"
For Bulk Strings the first byte of the reply is "$"
For Arrays the first byte of the reply is "*"

"+OK\r\n" ==> "OK";
"-Error message\r\n" ==> nil
":1000\r\n" ==> 1000。
"$6\r\nfoobar\r\n" ==> foobar   , "$-1\r\n" ==> key 不存在
"*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n" ==> ["foo","bar"]
"*3\r\n:1\r\n:2\r\n:3\r\n" ==> [1, 2, 3]

*/

func TestParseReply(t *testing.T)  {
	replyMap := make(map[string]interface{})
	replyMap["+OK\r\n"] = "OK"
	replyMap["-Error message\r\n"] = nil
	replyMap[":100\r\n"] = int64(100)
	replyMap["$0\r\n\r\n"] = "" // Null Bulk String.
	replyMap["$-1\r\n"] = nil   // Null Bulk String.
	replyMap["$5\r\nhello\r\n"] = "hello"

	for key, value := range replyMap{
		var parseReply interface{}
		parseReply = ParseReply(t, key, true)

		if value == parseReply{
			fmt.Println("success: ", parseReply)
		}else {
			fmt.Printf("error: key: %s, reply: %s \r\n", key, parseReply)
		}
	}

	//reply := "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
	//reply := "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
	//reply := "*3\r\n:1\r\n:2\r\n:3\r\n"
	//reply := "*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$6\r\nfoobar\r\n"
	//reply := "*0\r\n"  // 返回[]
	reply := "*-1\r\n"  //返回 nil
	parseReply := ParseReply(t, reply, true)
	fmt.Println("parseReply: ", parseReply)
}


 // reply: 需要解析的字符串; wanter: 当有错误时，是否通过测试。 true 返回值（空或者nil),false 不返回,报错；
func ParseReply(t *testing.T, reply string, wanter bool) interface{} {
	buf := new(bytes.Buffer)
	buf.WriteString(reply)
	p := proto.NewReader(buf)
	var multi proto.MultiBulkParse
	if strings.HasPrefix(reply, "*"){
		multi = multiBulkParse
	}
	replyStr, err := p.ReadReply(multi)

	//fmt.Println(replyStr)
	if !wanter && err != nil {
		fmt.Println("replyStr: ", replyStr)
		t.Errorf(err.Error())
	}
	return replyStr
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

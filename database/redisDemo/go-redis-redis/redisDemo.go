package main
/*
For Simple Strings the first byte of the reply is "+"
For Errors the first byte of the reply is "-"
For Integers the first byte of the reply is ":"
For Bulk Strings the first byte of the reply is "$"
For Arrays the first byte of the reply is "*"

"+OK\r\n" ==> "OK";
"-Error message\r\n" ==> Error message
":1000\r\n" ==> 1000。
"$6\r\nfoobar\r\n" ==> foobar   , "$-1\r\n" ==> key 不存在
"*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n" ==> ["foo","bar"]
"*3\r\n:1\r\n:2\r\n:3\r\n" ==> [1, 2, 3]

*/
import (
	"fmt"
	"net"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:6379")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err, conn)
		return
	}
	req := "*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"

	conn.Write([]byte(req))

	req = "*2\r\n$3\r\nget\r\n$3\r\nfoo\r\n"
	buffer := make([]byte, 2048)
	conn.Write([]byte(req))
	time.Sleep(10 * time.Millisecond)
	conn.Read(buffer)
	fmt.Println(buffer)
}

package main

import (
	"fmt"
	"testing"
)


func TestParseReply(t *testing.T) {
	statusReply, _ := ParseReply([]byte("+OK\r\n"))
	fmt.Println("statusReply: ", statusReply)
}
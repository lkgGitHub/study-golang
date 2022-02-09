/*
	Twitter 的 snowflake 算法

	+--------------------------------------------------------------------------+
	| 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
	+--------------------------------------------------------------------------+
*/
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/snowflake"
)

func main2() {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	for i := 0; i < 100; i++ {
		// Generate a snowflake ID.
		id := node.Generate()
		// Print out the ID in a few different ways.
		fmt.Printf("Int64  ID: %d\n", id)
		fmt.Printf("String ID: %s\n", id)
		fmt.Printf("Base2  ID: %s\n", id.Base2())
		fmt.Printf("Base64 ID: %s\n", id.Base64())

		// Print out the ID's timestamp
		fmt.Printf("ID Time  : %d\n", id.Time())

		// Print out the ID's node number
		fmt.Printf("ID Node  : %d\n", id.Node())

		// Print out the ID's sequence number
		fmt.Printf("ID Step  : %d\n", id.Step())

		fmt.Println(strings.Repeat("=", 80))
	}

}

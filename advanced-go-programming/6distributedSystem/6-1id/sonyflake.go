package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sony/sonyflake"
)

func getMachineID() (uint16, error) {
	return 1, nil
}

func checkMachineID(machineID uint16) bool {
	return true
}

func main() {
	t, _ := time.Parse("2006-01-02", "2022-01-01")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,
	}

	sf := sonyflake.NewSonyflake(settings)
	if sf == nil {
		fmt.Println("sonyFlake is nil")
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		id, err := sf.NextID()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(id)
	}
}

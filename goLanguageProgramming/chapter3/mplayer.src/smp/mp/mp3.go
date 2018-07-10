package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat int
	progress int
}

func (p *MP3Player) Play(source string) {
	fmt.Println()

	p.progress = 0
	for p.progress < 100 {
		time.Sleep(1000 * time.Microsecond)
		fmt.Print(".")
		p.progress += 10
	}

	fmt.Println("\nFinished playing", source)
}

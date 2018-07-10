package mp

import (
	"time"
	"fmt"
)

type WAVPlayer struct {
	stat int
	propress int
}

func (p *WAVPlayer) Play(source string) {
	for p.propress < 100 {
		time.Sleep(1000 * time.Microsecond)
		fmt.Print(".")
		p.propress += 10

	}
	fmt.Println("\nFinished playing", source)
}

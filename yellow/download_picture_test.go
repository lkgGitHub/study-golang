package yellow

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"sync"
	"testing"
	"time"
)

const baseDir = "/Volumes/SamSung/picture"

func TestDownloadPicture(t *testing.T) {
	start := time.Now()
	db := InitPG()
	videos := make([]*Video, 0, 30000)
	db.Find(&videos)

	wg := &sync.WaitGroup{}
	threadNum := make(chan struct{}, 100)
	for n, v := range videos {
		if n%100 == 0 {
			log.Println(n)
		}

		wg.Add(1)
		threadNum <- struct{}{}
		go func(name, url string) {
			defer func() {
				wg.Done()
				<-threadNum
			}()
			for i := 0; i < 5; i++ {
				resp, err := http.Get(url)
				if err != nil {
					if i == 4 {
						fmt.Println("http error:", err.Error())
						time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
					}
					continue
				}
				b, _ := io.ReadAll(resp.Body)
				_ = resp.Body.Close()
				err = os.WriteFile(path.Join(baseDir, name+".jpg"), b, 0666)
				if err != nil {
					fmt.Println("write file error:", err.Error())
					continue
				}
				break
			}

		}(v.Name, v.Picture)
	}

	wg.Wait()
	fmt.Printf("spend time: %v", time.Since(start))
}

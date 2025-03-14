package yellow

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGoQuery(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("video.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	db.AutoMigrate(&Video{})

	for count := 1; count <= 24; count++ {
	loop:
		for page := 1; ; page++ {
			log.Println("count:", count, "page:", page)
			res, err := http.Get(fmt.Sprintf(baseURL+"/vodtype/%d-%d.html", count, page))
			if err != nil {
				log.Fatal(err)
			}

			if res.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}

			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			// Find the review items
			videos := make([]*Video, 0, 40)
			doc.Find(".stui-vodlist__box").Each(func(i int, s *goquery.Selection) {
				if i <= 2 {
					return
				}
				a := s.Find("a[class='stui-vodlist__thumb lazyload']")
				title, _ := a.Attr("title")
				v := &Video{}
				if err = db.Where("name=?", title).First(v).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
					href, _ := a.Attr("href")
					dataOriginal, _ := a.Attr("data-original")
					duration := s.Find("span[class='pic-text text-right']").Text()
					hearts := s.Find("span[class='number pull-right']").Text()
					views := s.Find("span[class='pull-right']").Text()

					var year, month, day string
					for j, data := range strings.Split(dataOriginal, "/") {
						switch j {
						case 4:
							year = data
						case 5:
							month = data
						case 6:
							day = data
						}
					}
					publishDate := fmt.Sprintf("%s-%s-%s", year, month, day)

					v = &Video{
						Name:     title,
						URL:      href,
						Picture:  dataOriginal,
						Duration: duration,
					}
					v.Heart, _ = strconv.Atoi(strings.TrimSpace(hearts))
					v.Views, _ = strconv.Atoi(strings.TrimSpace(views))
					v.Publish, _ = time.Parse(time.DateOnly, publishDate)
					videos = append(videos, v)
				}
			})
			_ = res.Body.Close()

			switch len(videos) {
			case 0:
				break loop
			default:
				if err = db.Create(videos).Error; err != nil {
					for _, v := range videos {
						db.Create(v)
					}
				}
			}
		}
	}
}

func TestUpdateIsDownload(t *testing.T) {
	videoDB, err := gorm.Open(sqlite.Open("video.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	videoDB.AutoMigrate(&Video{})

	yellowDB, err := gorm.Open(sqlite.Open("yellow.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}

	yVideos := make([]*Video, 0, 2000)
	yellowDB.Model(&Video{}).Find(&yVideos)
	for i, v := range yVideos {
		name := strings.TrimSuffix(v.Name, ".mp4")
		err = videoDB.Model(&Video{}).Where("name=?", name).Update("is_download", true).Error
		if err != nil {
			t.Error("update error:", err)
			continue
		}
		if i%10 == 0 {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}

}

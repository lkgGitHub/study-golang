package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	name     = "博天堂资源站"
	uri      = "https://bttzyw.com"
	httpApi  = "http://bttcj.com/inc/sapi.php"
	httpsApi = "http://bttcj.com/inc/sapi.php"
	//httpApi = "http://sscj8.com/inc/api.php"
	//http://bttcj.com/inc/sapi.php?ac=videolist&ids=
)

func main() {
	url := fmt.Sprintf("%s?ac=list&t=2&pg=1", httpApi) // t=2

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	rss := &Rss{}
	err := xml.Unmarshal(body, rss)
	if err != nil {
		println(err)
	}
	ids := make([]string, 0)
	for _, v := range rss.List.Video {
		ids = append(ids, v.ID)
	}

	fmt.Printf("ids len: %d \n", len(ids))
	idsStr := strings.Join(ids, ",")
	videoListURL := "http://bttcj.com/inc/sapi.php?ac=videolist&ids=" + idsStr

	videoResp, _ := http.Get(videoListURL)
	defer videoResp.Body.Close()
	videoRespByte, _ := ioutil.ReadAll(videoResp.Body)
	err = xml.Unmarshal(videoRespByte, rss)
	if err != nil {
		println(err)
	}
	for _, v := range rss.List.Video {
		if v.Dl.Dd.Text != "" {
			t := v.Dl.Dd.Text
			fmt.Println(v.Name, t[8:len(t)-9])
			imagResp := &http.Response{}
			for i := 0; i < 3; i++ {
				imagResp, err = http.Get(v.Pic)
				if err != nil {
					println(err.Error())
					continue
				}
				break
			}
			if imagResp != nil {
				imagBody, _ := ioutil.ReadAll(imagResp.Body)
				out, err := os.Create(v.Name + ".jpg")
				if err != nil {
					println(err.Error())
					continue
				}
				_, err = io.Copy(out, bytes.NewReader(imagBody))
				if err != nil {
					println(err.Error())
					continue
				}
			}
		}
	}

}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Class   struct {
		Text string `xml:",chardata"`
		Ty   []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"ty"`
	} `xml:"class"`
	List struct {
		Text        string `xml:",chardata"`
		Page        string `xml:"page,attr"`
		Pagecount   string `xml:"pagecount,attr"`
		Pagesize    string `xml:"pagesize,attr"`
		Recordcount string `xml:"recordcount,attr"`
		Video       []struct {
			Text     string `xml:",chardata"`
			Last     string `xml:"last"`
			ID       string `xml:"id"`
			Tid      string `xml:"tid"`
			Name     string `xml:"name"`
			Type     string `xml:"type"`
			Pic      string `xml:"pic"`
			Lang     string `xml:"lang"`
			Area     string `xml:"area"`
			Year     string `xml:"year"`
			State    string `xml:"state"`
			Note     string `xml:"note"`
			Actor    string `xml:"actor"`
			Director string `xml:"director"`
			Dl       struct {
				Text string `xml:",chardata"`
				Dd   struct {
					Text string `xml:",chardata"`
					Flag string `xml:"flag,attr"`
				} `xml:"dd"`
			} `xml:"dl"`
			Des string `xml:"des"`
		} `xml:"video"`
	} `xml:"list"`
}

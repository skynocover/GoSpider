package spider

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//使用網址找出html
func Gethtml(website string) string {
	res, err := http.Get(website)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	sitemap, err := ioutil.ReadAll(res.Body) //使用ioutil讀取body
	if err != nil {
		log.Fatal(err)
	}
	return string(sitemap)
}

//下載圖片
func GetImg(fname string, url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		//name = fname + "/" + path[len(path)-1]
		name = fname + "/" + path[len(path)-1]
	}
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	if resp == nil {
		defer resp.Body.Close()
	}
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return
}

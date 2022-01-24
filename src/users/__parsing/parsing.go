package parser

import (
	"fmt"
	"github.com/opesun/goquery"
	"io"
	"net/http"
	"os"
	"go_project/src/config"
	"path/filepath"
	"time"

	//Пакеты, которые пригодятся для работы с файлами и сигналами:
	//"io"
	//"os"
	//"os/signal"
	////А вот эти - для высчитывания хешей:
	//"crypto/md5"
	//"encoding/hex"
)

var workers int = config.InitWorkers.Workers

func Grab() <-chan []string { //функция вернет канал, из которого мы будем читать данные типа string
	c := make(chan []string)
	urls := []string{}

	for i := 0; i < workers; i++ { //в цикле создадим нужное нам количество гоурутин - worker'oв
		go func() {
			defer close(c)
			for { //в вечном цикле собираем данные
				x, err := goquery.ParseUrl("https://www.atbmarket.com/trademark/goods")
				x.Find(".promo_image_link").Each(func(index int, element *goquery.Node) {
					cb := element.Child[1].Attr[0].Val
					fileUrl := "https://www.atbmarket.com/" + cb
					file := filepath.Base(cb)
					filePath, _ := filepath.Abs("../../src/parser/public/images/" + file)
					err := DownloadFile(filePath, fileUrl)
					if err != nil {
						println(err)
					}
					urls = append(urls, cb)
				})
				if err == nil {
					c <- urls
					//if s := x.Find(".promo_image_link").Attr("src"); s != "" {
					//	c <- s //и отправляем их в канал
					//}
				}
				time.Sleep(10000 * time.Millisecond)
			}
		}()
	}
	fmt.Println("Запущено потоков: ", workers)
	return c
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

//func check(e error) {
//	if e != nil {
//		panic(e)
//	}
//}

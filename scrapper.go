package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/html"
	"log"
	"mvdan.cc/xurls/v2"
	"net/http"
	"regexp"
)

// #js-dom-data-prefetched-data
// ss.childNodes[0].data

// получаем ссылку и айди объекта
// фетчим для ссылки данные и пишем в доп поле по айди объекта
func Scrape(url string, uid string) {

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var comment *html.Node
	// Find the review items
	doc.Find("#js-dom-data-prefetched-data").Each(func(i int, s *goquery.Selection) {
		for _, node := range s.Nodes {
			if comment = findComment(node); comment != nil {
				break
			}
		}

		// TODO: нужно сохранить текст описания сцены или в БД или в файл
		fmt.Println(comment.Data + uid)

		_ = SaveSceneData(comment.Data, uid)

		rxStrict := xurls.Strict()
		rxStrict.FindAllString(comment.Data, -1)
		// вытаскиваем все ссылки
		dirturls := rxStrict.FindAllString(comment.Data, -1)

		// проходим по всем ссылкам и берем только содержащие https://media.sketchfab.com

		var cleanurls []string // выбранные ссылки

		for index := range dirturls {
			result, err := regexp.MatchString("(https://media.sketchfab.com)", dirturls[index])
			// если найдено вхождение и нет ошибки
			if result && err == nil {
				// объединяем нужные ссылки в новый массив
				cleanurls = append(cleanurls, dirturls[index])
			}
		}

		fmt.Println(cleanurls)

		// TODO: нужно отдать ссылки блоку скачивания ссылок
		// func (cleanurls)

		err = downloadservice(cleanurls)

		// TODO: если происходи тошибка при скачивании метка все равно ставится что скачало
		if err == nil {	_ = MarkAsDownloaded(uid) }
	})
}

func findComment(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	if n.Type == html.CommentNode {
		return n
	}
	if res := findComment(n.FirstChild); res != nil {
		return res
	}
	if res := findComment(n.NextSibling); res != nil {
		return res
	}
	return nil
}

func SaveSceneData(data string, uid string) error {
	var collection = MongoClient.Database(DATABASE).Collection(COLLECTION)
	var filter = bson.D{{"uid", uid}}
	update := bson.D{
		{"$set", bson.D{{"checked", true}, {"downloaded", false}, {"SceneData", data}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func MarkAsDownloaded(uid string) error {
	var collection = MongoClient.Database(DATABASE).Collection(COLLECTION)
	var filter = bson.D{{"uid", uid}}
	update := bson.D{
		{"$set", bson.D{{"downloaded", true}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

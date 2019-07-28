package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
	"github.com/riy0/menu_recommend/db"
	"log"
	"net/url"
)

var (
	urls      []string
	menus     []string
	imageUrls []string
)

func ebaraFoodCollector() error {
	db := db.GetDB()
	ebaraUrl := "https://www.ebarafoods.com/recipe/cla_menu/49/?&limit=100"

	doc, err := goquery.NewDocument(ebaraUrl)
	if err != nil {
		return err
	}

	doc.Find("ul.list-results li a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		burl, _ := url.Parse(ebaraUrl)
		var fullUrl = toAbsUrl(burl, href)
		urls = append(urls, fullUrl)
	})

	doc.Find("li dl").Each(func(i int, s *goquery.Selection) {
		menu := s.Find("dt").Text()
		menus = append(menus, menu)
	})

	doc.Find("ul.list-results > li > a > figure > img").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("src")
		burl, _ := url.Parse(ebaraUrl)
		var fullUrl = toAbsUrl(burl, href)
		imageUrls = append(imageUrls, fullUrl)
	})

	fmt.Printf("ngo")
	sql := "INSERT INTO menus(title, cooking_time, image_url, url, calorie, category) VALUES($1, $2, $3, $4, $5, $6)"
	stmt, err := db.Prepare(Sql)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ngo")
	for j := range menus {
		_, err = stmt.Exec(menus[j], 0, imageUrls[j], urls[j], 0, 0)
	}

	return nil
}

func toAbsUrl(baseUrl *url.URL, webUrl string) string {
	relUrl, err := url.Parse(WebUrl)
	if err != nil {
		return ""
	}
	absUrl := baseUrl.ResolveReference(relUrl)
	return absUrl.String()
}

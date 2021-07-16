package parser

import (
	goquery"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	
	"fmt"

)

func GetFilms(r []string) []string{
	res, err := http.Get("https://www.kinopoisk.ru/afisha/new/city/1/")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("statuse 200")
	}

	doc , err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	LinkAll := doc.Find(".best_res").Find("table").Find("tbody").Find("tr").Find("b")
	// fmt.Println(link)
	// fmt.Println(*&k.Parent.Data)
	// text,_:=LinkAll.Find(".cinema-today-film-link-4299813").Find(".title").Html()
	// var b bool = true
	LinkAll.EachWithBreak(
		
			func(i int, s *goquery.Selection) (b bool) {  // For each item found, get the title
			
		
		title,_ := s.Find("a").Attr("href")  
		fmt.Printf("Review %d: %s\n",i, title)  
		r = append(r, title)
		// if i > 5 {
		// 	return false
		// }
		return true
		}) 
		
		return r
}
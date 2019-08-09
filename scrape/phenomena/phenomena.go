package phenomena

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type scheduling struct {
	Time  string `json:"time"`
	Title string `json:"title"`
}

type day []scheduling

type Month map[string]day

func trim(s string) string {
	return strings.TrimSpace(s)
}

// FetchMonth returns the calendar for a month
func FetchMonth(year, month int) Month {
	url := fmt.Sprintf("http://www.phenomena-experience.com/programacion-calendario/%d-%d.html", month, year)
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

	m := Month{}
	doc.Find(".cal-columnadia").Each(func(i int, s *goquery.Selection) {
		dayText := s.Find(".cal-titulodia").Text()
		dayText = strings.Split(dayText, " ")[1]
		dayNumber, _ := strconv.Atoi((trim(dayText)))
		day := day{}

		s.Find(".cal-film").Each(func(i int, s *goquery.Selection) {
			title := s.Find(".cal-film-texto").Text()
			time := s.Find(".cal-film-hora").Text()

			day = append(day, scheduling{Time: trim(time), Title: trim(title)})

			fullDate := fmt.Sprintf("%d-%d-%d", year, month, dayNumber)
			m[fullDate] = day
		})
	})

	return m
}

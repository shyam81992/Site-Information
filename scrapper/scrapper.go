package scrapper

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

//go:generate mockgen -destination=./mock/scrapper.go -package=mock github.com/shyam81992/Site-Information/scrapper IScrapper

var loginregex = regexp.MustCompile("(?i)(username|password|email|otp|forgot)")

type IScrapper interface {
	GetPageInfo(string) gin.H
}

type Scrapper struct {
}

// urlstatus is safe to use concurrently.
type urlStatus struct {
	mu sync.Mutex
	v  map[string]bool
}

// Sets the current valuefor the given key.
func (us *urlStatus) Set(key string, val bool) {
	us.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer us.mu.Unlock()
	us.v[key] = val
}

// Value returns the current value of the counter for the given key.
func (us *urlStatus) Value(key string) bool {
	us.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer us.mu.Unlock()
	return us.v[key]
}

// headingTags is safe to use concurrently.
type headingTags struct {
	mu sync.Mutex
	v  map[string]int
}

// Sets the current valuefor the given key.
func (ht *headingTags) Inc(key string) {
	ht.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer ht.mu.Unlock()
	ht.v[key]++
}

func (s *Scrapper) GetPageInfo(URL string) gin.H {
	col := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
	)

	us := urlStatus{v: make(map[string]bool)}
	loginform := false
	pagetitle := ""
	ht := headingTags{v: make(map[string]int)}
	internalLinks := 0
	externalLinks := 0
	Status := 200
	// Find and visit all links

	var linkregex = regexp.MustCompile("(?i)^(" + URL + "|\\./|\\.\\./|#|/)")

	col.OnHTML("html", func(e *colly.HTMLElement) {

		if e.Request.URL.String() == URL {
			e.ForEach("a[href]", func(i int, e *colly.HTMLElement) {

				// mutex is not requiered because code execution here is synchronous
				if linkregex.MatchString(e.Attr("href")) {
					fmt.Println("matched", e.Attr("href"))
					internalLinks++
				} else {
					fmt.Println("not matched", e.Attr("href"))
					externalLinks++
				}

				e.Request.Visit(e.Attr("href"))
			})

			for _, val := range [6]string{"h1", "h2", "h3", "h4", "h5", "h6"} {
				e.ForEach(val, func(i int, h *colly.HTMLElement) {
					ht.Inc(val)
				})
			}
			e.ForEach("form", func(i int, e *colly.HTMLElement) {

				text := e.Text
				stringsmatched := loginregex.FindAllString(strings.ToLower(text), -1)

				username := false
				password := false
				email := false
				otp := false
				forgot := false

				for i := 0; i < len(stringsmatched); i++ {
					if stringsmatched[i] == "username" {
						username = true
					}
					if stringsmatched[i] == "password" {
						password = true
					}
					if stringsmatched[i] == "email" {
						email = true
					}
					if stringsmatched[i] == "otp" {
						otp = true
					}
					if stringsmatched[i] == "forgot" {
						forgot = true
					}
				}

				if (username || email) && (password || otp || forgot) {
					loginform = true
				}

			})

			pagetitle = e.ChildText("title")
		}

	})

	col.OnRequest(func(r *colly.Request) {
		us.Set(r.URL.String(), false)
		//fmt.Println("Visiting", r.URL)
	})

	// attach callbacks after login
	col.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 200 {
			us.Set(r.Request.URL.String(), true)
		}
		//log.Println("response received", r.StatusCode, r.Request.URL)
	})
	col.OnError(func(r *colly.Response, e error) {
		if r.Request.URL.String() == URL {
			Status = r.StatusCode
		}
		fmt.Println(e)
	})

	col.Visit(URL)

	col.Wait()

	invalidurls := 0

	for _, value := range us.v {
		if !value {
			invalidurls++
		}
	}
	return gin.H{
		"Status":             Status,
		"Page Title":         pagetitle,
		"Headings":           ht.v,
		"Internal Links":     internalLinks,
		"External Links":     externalLinks,
		"Inaccessible Links": invalidurls,
		"loginform":          loginform,
	}

}

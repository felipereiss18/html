package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Titulos é responsável em buscar os títulos das páginas html
func Titulos(urls ...string) (ch chan string) {
	ch = make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return
}

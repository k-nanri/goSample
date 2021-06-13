package threadsample

import (
	f "fmt"
	"log"
	"net/http"
)

func Execute() {

	f.Println("=== Thread Sample Execute =====")

	urls := []string {
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	for _, url := range urls {
		// ゴルーチンで実行
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			f.Println(url, res.Status)
		}(url)
	}

}
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

	statusChan := make(chan string)
	for _, url := range urls {
		
		// ゴルーチンで実行
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			statusChan <- res.Status

		}(url)
	}
	// 待ち合わせ
	for i := 0; i< len(urls); i++ {
		f.Println(<-statusChan)
	}

	f.Println("Chanel Finish!!")

}
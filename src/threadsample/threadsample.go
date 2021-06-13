package threadsample

import (
	f "fmt"
	"log"
	"net/http"
	"sync"
)

func Execute() {

	f.Println("=== Thread Sample Execute =====")

	wait := new (sync.WaitGroup)
	urls := []string {
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	for _, url := range urls {

		// waitGroupに追加
		wait.Add(1)
		// ゴルーチンで実行
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			f.Println(url, res.Status)
			// waitGroupから削除
			wait.Done()
		}(url)
	}
	// 待ち合わせ
	wait.Wait()
	f.Println("Wait Finish!!")

}
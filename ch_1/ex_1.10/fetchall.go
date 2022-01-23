// Fetchall fetches urls in parallel
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	file, err := os.Create("./results.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // starts a goroutine
	}
	for range os.Args[1:] {
		file.WriteString(<-ch) //receive from channel ch
	}

	line := fmt.Sprintf("%.2fs elasped\n", time.Since(start).Seconds())
	file.WriteString(line)
	// fmt.Printf("%.2fs elasped\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while read %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}

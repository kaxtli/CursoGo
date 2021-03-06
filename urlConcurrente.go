package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "time"
  "strings"
)

func main() {
  start := time.Now()
  ch := make(chan string)

  for _, url := range os.Args[1:] {
    go fetch(url, ch)
  }

  for range os.Args[1:] {
    fmt.Printf( <- ch)
  }

  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<-string)  {
  start := time.Now()
  resp, err := http.Get(checkPrefix(url))
  if err != nil{
    ch <- fmt.Sprint(err)
    return
  }
  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close()

  if err != nil {
    ch <- fmt.Sprintf("while reading %s: %v", url, err)
  }

  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7d %s ", secs, nbytes, url)
}

func checkPrefix(url string) string{
	newURL := url
	if ! (strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		newURL = "http://" + url
	}
	return newURL
}

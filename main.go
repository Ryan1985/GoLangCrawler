package main

import (
	"fmt"
	"strconv"
	"time"

	"GOLANGCRAWLER/parseUrls"
)

func main() {
	start := time.Now()
	ch := make(chan bool)
	for j := 0; j <= 15; j++ {
		for i := 0; i < 10; i++ {
			if i == 0 {
				go parseUrls.ParseUrls("https://www.dy2018.com/"+strconv.Itoa(j)+"/index.html", "dy2018movieList.log", ch)
			} else {
				go parseUrls.ParseUrls("https://www.dy2018.com/"+strconv.Itoa(j)+"/index_"+strconv.Itoa(i)+".html", "dy2018movieList.log", ch)
			}
		}
	}

	for i := 0; i < 15*10; i++ {
		<-ch
	}

	// n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
}

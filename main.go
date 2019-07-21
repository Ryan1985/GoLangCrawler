package main

import (
	"fmt"
	"time"
	"strconv"

	"SPIDER/parseUrls"
)


func main(){
	start := time.Now()

for	j:=0;j<=15;j++{
	for i:=0;i<10;i++{
		if i == 0{
			parseUrls.ParseUrls("https://www.dy2018.com/"+strconv.Itoa(j)+"/index.html","dy2018movieList.log")
		}else{
			parseUrls.ParseUrls("https://www.dy2018.com/"+strconv.Itoa(j)+"/index_"+ strconv.Itoa(i)+".html","dy2018movieList.log")
		}
	}
}
	// n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	elapsed:= time.Since(start)
	fmt.Printf("Took %s",elapsed)
}
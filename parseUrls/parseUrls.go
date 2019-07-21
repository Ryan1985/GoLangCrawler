package parseUrls

import (
//	"fmt"
	"regexp"
	"strings"
	"os"
	
	"GOLANGCRAWLER/fetch"
)


func ParseUrls(url string,fileName string)  {
	var dstFile (*os.File)
	_, err := os.Stat(fileName)  
    if err != nil {  
		dstFile,_ = os.Create(fileName)
    }  else{
		dstFile,_ = os.OpenFile(fileName, os.O_APPEND, 0666) 
	}
///html/tv/hytv/index.html
	body := fetch.Fetch(url)
	body = strings.Replace(body,"\n","",-1)
	rp := regexp.MustCompile(`<a href="(.*?)" class="ulink" title="(.*?)">(.*?)</a>`)
	idRe:=regexp.MustCompile(`(/i/(\d+)|/html/(.*?)).html`)
	items:=rp.FindAllStringSubmatch(body,-1)
	for _,item:= range items{
		resultString:="https://www.dy2018.com"+idRe.FindStringSubmatch(item[1])[1]+".html"+"\t"+item[3]
		//fmt.Println(resultString)
		dstFile.WriteString(resultString + "\n")
	}
	defer dstFile.Close()
}
package fetch

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func Fetch (url string) string{
	fmt.Println("Fetch Url",url)
	client:= &http.Client{}
	req,_:=http.NewRequest("GET",url,nil)
	req.Header.Set("User-Agent","Mozilla/5.0  (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	req.Header.Set("content-type","text/html")
	req.Header.Set("charset","gb2312")
	resp,err :=client.Do(req)
	
	if err != nil{
		fmt.Println("Http get err:",err)
		return ""
	}

	if resp.StatusCode !=200{
		fmt.Println("Http status code:",resp.StatusCode)
		return ""
	}

	defer resp.Body.Close()
	body,err:= ioutil.ReadAll(resp.Body)
	if  err !=nil{
		fmt.Println("Read error", err)
		return ""
	}

	return string(body)
}
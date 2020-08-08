package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/////调用第三方API/////
func wangYiXinWenDataGet() []StdNew {
	url := "https://3g.163.com/touch/reconstruct/article/list/BBM54PGAwangning/0-10.html"

	resp, err := http.Get(url)
	if err !=nil {
		fmt.Println("Couldn't fetch news", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err !=nil {
		fmt.Println("Couldn't read news", err)
	}

	//fmt.Println(string(body))

	data := WangYiXinWenRaw{}
	Json2Go([]byte(body[9: len(body)-1]), &data)

	//fmt.Println(data)

	stdNews := WangYiXinWenGO2StdGo(data)

	//for i:= range stdNews{
	//	stdJson := Go2Json(stdNews[i])
	//}

	return stdNews
}

func tianXingTouTiaoDataGet() []StdNew{
	url := "http://api.tianapi.com/topnews/index?key=4703f453980fd2a17b6413701b591c4b"

	resp, err := http.Get(url)
	if err !=nil {
		fmt.Println("Couldn't fetch news", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err !=nil {
		fmt.Println("Couldn't read news", err)
	}

	data := TianXingTouTiaoRaw{}
	Json2Go([]byte(body[9: len(body)-1]), &data)

	stdNews := TianXingTouTiaoGO2StdGo(data)
	return stdNews
}

/////提取正文/////
//func getArticleContent(value GeneralDataList) {
//	panic("not implemented")
//}
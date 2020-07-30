package main

import (
	"encoding/json"
	"fmt"
)

//////////数据归一化///////

// 网易新闻struct
type WangYiXinWenRaw struct {
	BBM54PGAwangning []struct {
		LiveInfo interface{} `json:"liveInfo"`
		Docid string `json:"docid"`
		Source string `json:"source"`
		Title string `json:"title"`
		Priority int `json:"priority"`
		HasImg int `json:"hasImg"`
		URL string `json:"url"`
		CommentCount int `json:"commentCount"`
		Imgsrc3Gtype string `json:"imgsrc3gtype"`
		Stitle string `json:"stitle"`
		Digest string `json:"digest"`
		Imgsrc string `json:"imgsrc"`
		Ptime string `json:"ptime"`
	} `json:"BBM54PGAwangning"`
}

// 天行头条struct
type TianXingTouTiaoRaw struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Newslist []struct {
		Ctime string `json:"ctime"`
		Title string `json:"title"`
		Description string `json:"description"`
		PicURL string `json:"picUrl"`
		URL string `json:"url"`
		Source string `json:"source"`
	} `json:"newslist"`
}

func Json2Go(body []byte, rawNews interface{}) {
	err := json.Unmarshal([]byte(body), &rawNews)
	if err != nil {
		fmt.Println("Couldn't convert json to go", err)
	}
	//fmt.Println("AAA: ",rawNews)
}
func Go2Json(data interface{}) []byte{
	stdJson, err := json.Marshal(data)
	if err !=nil {
		fmt.Println("Couldn't convert go objects back to json", err)
	}
	//fmt.Println(fmt.Sprintf(string(stdJson)))
	return stdJson
}

// 标准struct
type StdNew struct{
	Timestamp string `json:"timestamp"`
	Source string `json:"source"`
	Title string `json:"title"`
	Body string `json:"body"`
	URL string `json:"URL"`
	//Types []string `json:"types"`
}

func TianXingTouTiaoGO2Std(txtt TianXingTouTiaoRaw) []StdNew {
	ret := []StdNew{}
	for i:= range txtt.Newslist{
		var item StdNew
		item.Source = txtt.Newslist[i].Source
		item.Timestamp = txtt.Newslist[i].Ctime
		item.Title = txtt.Newslist[i].Title
		item.Body = txtt.Newslist[i].Description
		item.URL = txtt.Newslist[i].URL
		ret = append(ret, item)
	}
	return ret
}
func WangYiXinWenGO2StdGo(wyxw WangYiXinWenRaw) []StdNew {
	ret := []StdNew{}
	for i:= range wyxw.BBM54PGAwangning{
		var item StdNew
		item.Source = wyxw.BBM54PGAwangning[i].Source
		item.Timestamp = wyxw.BBM54PGAwangning[i].Ptime
		item.Title = wyxw.BBM54PGAwangning[i].Title
		item.Body = wyxw.BBM54PGAwangning[i].Digest
		item.URL = wyxw.BBM54PGAwangning[i].URL
		ret = append(ret, item)
	}
	return ret
}
//func ZHGO2Std(zh ZongHeNewsRaw) []StdNew {
//	ret := []StdNew{}
//	for i:= range zh.Newslist{
//		var item StdNew
//		item.Source = "zonghe"
//		item.Timestamp = time.Now().Format("2006-01-02 15:04:05")
//		item.Title = zh.Newslist[i].Title
//		item.Body = zh.Newslist[i].Description
//		item.URL = zh.Newslist[i].URL
//		ret = append(ret, item)
//	}
//	return ret
//}
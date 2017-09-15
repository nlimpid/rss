package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"encoding/json"

	"io/ioutil"

	"path"

	"github.com/gorilla/feeds"
	"github.com/nlimpid/rss/models"
	"github.com/pressly/chi"
)

var baseURL = "https://zhuanlan.zhihu.com/"
var postURL = "https://zhuanlan.zhihu.com/api/columns/"

// GetArticle 获取文章
func GetArticle(w http.ResponseWriter, r *http.Request) {
	// Assume if we've reach this far, we can access the article
	// context because this handler is a child of the ArticleCtx
	// middleware. The worst case, the recoverer middleware will save us.

	articleName := chi.URLParam(r, "articleName")

	url := fmt.Sprintf("%v%v", baseURL, articleName)
	log.Println(url)
	// w.Write([]byte(ToRss()))
	w.Write([]byte(ToZhihuRss(articleName)))

}

func ToZhihuRss(name string) string {
	now := time.Now()
	zhihuFeed, err := getPost(name)
	if err != nil {
		return ""
	}
	feed := &feeds.Feed{
		Title:       zhihuFeed.Name,
		Link:        &feeds.Link{Href: fmt.Sprintf("%v%v", baseURL, name)},
		Description: zhihuFeed.Description,
		Updated:     now,
	}
	zhihuItems := getItem(name)
	// feed.Items = *feeds.Item{}
	for _, v := range zhihuItems {
		publishTime, _ := time.Parse(time.RFC3339, v["publishedTime"].(string))
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       v["title"].(string),
			Link:        &feeds.Link{Href: path.Join(baseURL, v["url"].(string))},
			Description: v["content"].(string),
			Author:      &feeds.Author{Name: v["author"].(map[string]interface{})["name"].(string)},
			Created:     publishTime,
		})
	}
	log.Println(feed.Items)
	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(rss)
	return rss
}

// ZhihuPost 总的post
func getPost(name string) (models.ZhihuPost, error) {
	c := http.Client{
		Timeout: 1 * time.Minute,
	}
	url := fmt.Sprintf("%v%v", postURL, name)
	req, err := http.NewRequest("GET", url, nil)
	log.Printf("url=%v\n", url)
	if err != nil {
		log.Printf("getPost NewRequest %v\n", err)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Printf("getPost Do %v\n", err)
		return models.ZhihuPost{}, err
	}
	zhiPost := models.ZhihuPost{}
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &zhiPost)
	if err != nil {
		log.Println(err)
	}
	return zhiPost, nil
}

func getItem(name string) []map[string]interface{} {
	c := http.Client{
		Timeout: 1 * time.Minute,
	}
	url := fmt.Sprintf("%v%v%v", postURL, name, "/posts")
	req, err := http.NewRequest("GET", url, nil)
	log.Printf("url=%v\n", url)
	if err != nil {
		log.Printf("getPost NewRequest %v\n", err)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Printf("getPost Do %v\n", err)
	}
	// log.Println(resp.Body)
	v := []map[string]interface{}{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		log.Println(err)
	}
	regex, err := regexp.Compile(`<img src="(v2-[0-9a-zA-Z]{32}.(png|jpg|gif))"`)
	for i, j := range v {
		res := regex.FindAllString(j["content"].(string), 3)
		if res != nil {
			fmt.Printf("img res  =%v\n", res[0])
		}
		refControl := "https://rss.nlimpid.com/zhihu_image?image="
		replacedstr := regex.ReplaceAllString(j["content"].(string), fmt.Sprintf("<img src=\"%vhttps://pic4.zhimg.com/$1\"", refControl))
		titleImage := j["titleImage"]
		log.Printf("origin title Image =%s\n", titleImage.(string))
		if titleImage.(string) != "" {
			replaceImage := fmt.Sprintf("%s%s", "https://rss.nlimpid.com/zhihu_image?image=", j["titleImage"])
			titleImageStr := fmt.Sprintf("<img src=\"%v\">", replaceImage)
			log.Printf("titleImageStr=%s\n", titleImageStr)
			v[i]["content"] = fmt.Sprintf("%s%s", titleImageStr, replacedstr)
		} else {
			v[i]["content"] = replacedstr
		}
	}
	return v
}

func getItems(name, limit, offset string) (items []models.ZhihuItem) {
	c := http.Client{
		Timeout: 1 * time.Minute,
	}
	url := fmt.Sprintf("%v%v%v", postURL, name, "/posts")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("getPost NewRequest %v\n", err)
	}
	values := req.URL.Query()
	values.Add("limit", limit)
	values.Add("offset", offset)
	req.URL.RawQuery = values.Encode()
	resp, err := c.Do(req)
	if err != nil {
		log.Printf("getPost Do %v\n", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(body, &items)
	if err != nil {
		log.Println(err)
	}
	return
}

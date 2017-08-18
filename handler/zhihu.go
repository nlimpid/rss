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

func ToRsfffs() string {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "jmoiron.net blog",
		Link:        &feeds.Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
		Created:     now,
	}
	feed.Items = []*feeds.Item{
		&feeds.Item{
			Title:       "Logic-less Template Redux",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		&feeds.Item{
			Title:       "Limiting Concurrency in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
			Created:     now,
		},
	}
	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(rss)
	return rss
}

// ZhihuPost 总的post
type ZhihuPost struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ZhihuItem 每一条文章
type ZhihuItem struct {
	Title       string `json:"title"`
	Link        string
	Description string `json:"content"`
	Created     time.Time
}

func getPost(name string) (ZhihuPost, error) {
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
		return ZhihuPost{}, err
	}

	zhiPost := ZhihuPost{}
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
	for k, _ := range v[0] {
		log.Printf("key=%v\n", k)
	}
	fmt.Printf("titleImage =%v\n", v[0]["titleImage"])
	fmt.Printf("snapshorURL=%v\n", v[0]["snapshortUrl"])
	regex, err := regexp.Compile(`<img src="(v2-[0-9a-zA-Z]{32}.jpg)"`)
	res := regex.FindAllString(v[0]["content"].(string), 3)
	if res != nil {
		fmt.Printf("img res  =%v\n", res[0])
	}
	// refControl := "http://read.html5.qq.com/image?src=forum&q=5&r=0&imgflag=7&imageUrl="
	refControl := "http://172.104.105.215:6334/zhihu_image?image="
	replacedstr := regex.ReplaceAllString(v[0]["content"].(string), fmt.Sprintf("<img src=\"%vhttps://pic4.zhimg.com/$1\"", refControl))
	v[0]["content"] = replacedstr

	return v
}

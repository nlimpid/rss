package models

import (
	"fmt"
	"regexp"

	"github.com/gorilla/feeds"
)

var baseURL = "https://zhuanlan.zhihu.com"

// ZhihuPostAvatar represnet zhihu zhuanlan post avatar
type ZhihuPostAvatar struct {
	ID       string `json:"id"`
	Template string `json:"template"`
}

// ZhihuPost represent zhihu zhuanlan struct
type ZhihuPost struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Link        string          `json:"url"`
	Avatar      ZhihuPostAvatar `json:"avatar"`
}

// FullLink get the complete url of link
func (z ZhihuPost) FullLink() string {
	return fmt.Sprintf("%s%s", baseURL, z.Link)
}

// GetRssImage generate the image struct
func (z ZhihuPost) GetRssImage() *feeds.RssImage {
	fi := feeds.RssImage{
		Title: "avatar",
		Url:   z.Avatar.FullAvatar(),
	}
	return &fi

}

// FullAvatar get the biggest img from the template
func (a ZhihuPostAvatar) FullAvatar() string {
	re, _ := regexp.Compile("^(.*){id}_{size}(.jpg)$")
	replaced := fmt.Sprintf("${1}%v${2}", a.ID)
	return re.ReplaceAllString(a.Template, replaced)
}

// ReplacedImg to avoid origin check
func (a ZhihuPostAvatar) ReplacedImg() string {
	originImg := a.FullAvatar()
	refBase := "https://rss.nlimpid.com/zhihu_image?image="
	image := fmt.Sprintf("%s%s", refBase, originImg)
	return image
}

// ZhihuItem represent single article
type ZhihuItem struct {
	Title       string `json:"title"`
	TitleImage  string `json:"titleImage"`
	Link        string `json:"url"`
	Description string `json:"content"`
	Created     string `json:"publishedTime"`
}

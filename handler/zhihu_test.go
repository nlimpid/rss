package handler

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/nlimpid/rss/models"
)

func Test_getPost(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want models.ZhihuPost
	}{
		// TODO: Add test cases.
		{
			"test1", args{"oh-hard"},
			models.ZhihuPost{
				Name:        "硬派健身",
				Description: "每日一篇质量长文，微信公众：硬派健身。",
				Link:        "/oh-hard",
				Avatar: models.ZhihuPostAvatar{
					ID:       "d8752eeb8ddb0ca382afc836de6224c0",
					Template: "https://pic1.zhimg.com/{id}_{size}.jpg",
				}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPost(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPost() = %v, want %v", got, tt.want)
			}
			if err != nil {
				t.Errorf("getPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getItem(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want models.ZhihuPost
	}{
		{
			"test1", args{"oh-hard"},
			models.ZhihuPost{
				Name:        "硬派健身",
				Description: "每日一篇质量长文，微信公众：硬派健身。",
				Avatar: models.ZhihuPostAvatar{
					ID:       "0a47432f4ef552ceaf1da5a9fe11a443",
					Template: "https://pic4.zhimg.com/{id}_{size}.jpg",
				}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getItem(tt.args.name)
			fmt.Printf("got content=%v\n", got[0]["content"])
			if !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("getPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	regex, _ := regexp.Compile(`<img src="(v2-[0-9a-zA-Z]{32}.(png|jpg|gif))"`)
	beforeStr := "</p><img src=\"v2-47487ea266bdf33aa0c0552331ecc3e0.png\" data-rawwidth=\"773\" data-rawheight=\"457\"><p>"
	res := regex.FindAllString(beforeStr, 3)
	fmt.Printf("res = %v\n", res)
	refControl := "https://rss.nlimpid.com/zhihu_image?image="
	replacedstr := regex.ReplaceAllString(beforeStr, fmt.Sprintf("<img src=\"%vhttps://pic4.zhimg.com/$1\"", refControl))
	fmt.Printf("after str = %v\n", replacedstr)
}

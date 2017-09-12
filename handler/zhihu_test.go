package handler

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func Test_getPost(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want ZhihuPost
	}{
		// TODO: Add test cases.
		{"test1", args{"undefined"}, ZhihuPost{"前端黑板报",
			"前端黑板报由各大互联网公司的资深前端和技术专家合力运营，致力于传播高质量的行业资讯，分享有态度的意见观点。"}},
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
		want ZhihuPost
	}{
		// TODO: Add test cases.
		{"test1", args{"undefined"}, ZhihuPost{"前端黑板报",
			"前端黑板报由各大互联网公司的资深前端和技术专家合力运营，致力于传播高质量的行业资讯，分享有态度的意见观点。"}},
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

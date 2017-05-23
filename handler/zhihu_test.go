package handler

import (
	"reflect"
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
			if got := getPost(tt.args.name); !reflect.DeepEqual(got, tt.want) {
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
			if got := getItem(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("getPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

package models

import "testing"

func TestZhihuPostAvatar_FullAvatar(t *testing.T) {
	type fields struct {
		ID       string
		Template string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"test1",
			fields{
				ID:       "d8752eeb8ddb0ca382afc836de6224c0",
				Template: "https://pic1.zhimg.com/{id}_{size}.jpg",
			},
			"https://pic1.zhimg.com/d8752eeb8ddb0ca382afc836de6224c0.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ZhihuPostAvatar{
				ID:       tt.fields.ID,
				Template: tt.fields.Template,
			}
			if got := a.FullAvatar(); got != tt.want {
				t.Errorf("ZhihuPostAvatar.FullAvatar() = %v, want %v", got, tt.want)
			}
		})
	}
}

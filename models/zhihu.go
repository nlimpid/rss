package models

import "time"

// ZhihuPost represent zhihu zhuanlan struct
type ZhihuPost struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ZhihuItem represent single article
type ZhihuItem struct {
	Title       string `json:"title"`
	Link        string
	Description string `json:"content"`
	Created     time.Time
}

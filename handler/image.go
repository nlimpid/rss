package handler

import (
	"net/http"

	"io/ioutil"

	"github.com/Sirupsen/logrus"
)

var imageBase = "https://pic4.zhimg.com"

// GetImage 获取图片（实际上是中转图片）
func GetImage(w http.ResponseWriter, r *http.Request) {
	imageURL := r.URL.Query().Get("image")

	resp, err := http.Get(imageURL)
	if err != nil {
		logrus.Error(err)
		return
	}

	w.Header().Add("Content-Type", "image/*")
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return
	}
	w.Write(result)

}

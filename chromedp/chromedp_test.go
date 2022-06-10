package chromedptool

import (
	"log"
	"os"
	"testing"
)

func TestPrintf(t *testing.T) {
	path := "./baidu.pdf"
	defer os.RemoveAll(path)
	h := ChromedpHandler{}

	url, err := h.GetConnUrl()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(url)

	err = h.GenPdf("https://baidu.com", path)
	if err != nil {
		t.Error(err)
	}
}

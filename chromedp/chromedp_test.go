package chromedptool

import (
	"os"
	"testing"
)

func TestPrintf(t *testing.T) {
	path := "./baidu.pdf"
	defer os.RemoveAll(path)

	err := ChromedpHandler{}.GenPdf("https://baidu.com", path)
	if err != nil {
		t.Error(err)
	}
}

package chromedptool

import (
	"context"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	_blank = "about:blank"
)

// "description": "",
// "devtoolsFrontendUrl": "/devtools/inspector.html?ws=localhost:9222/devtools/page/E8CB7F8AAB1FF9ECB2D76C5363C6FE7A",
// "id": "E8CB7F8AAB1FF9ECB2D76C5363C6FE7A",
// "title": "about:blank",
// "type": "page",
// "url": "about:blank",
// "webSocketDebuggerUrl": "ws://localhost:9222/devtools/page/E8CB7F8AAB1FF9ECB2D76C5363C6FE7A"
// curl http://localhost:9222/json,we can get output like this which contain the info for connecting chromedp
type ChromedpConnParam struct {
	Description          string
	DevtoolsFrontendUrl  string
	Id                   string
	Title                string
	Type                 string
	Url                  string
	WebSocketDebuggerUrl string
}

type ChromedpHandler struct {
}

/**
 * @description: get connection info from chromedp
 * @param {...string} urls   chromdp service url,default using http://localhost:9222/json
 * @return {*}
 */
func (c ChromedpHandler) GetConnUrl(urls ...string) (*ChromedpConnParam, error) {
	client := resty.New()
	url := ""
	if len(urls) == 0 || urls[0] == "" {
		url = "http://localhost:9222/json"
	} else {
		url = urls[0]
	}
	res := []ChromedpConnParam{}
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&res).Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "get conn url failed")
	}

	if len(res) <= 0 {
		return nil, errors.New("cannot find conn url")
	}
	connParam := res[len(res)-1]
	go c.Clear(res)
	return &connParam, nil
}

/**
 * @description: generate pdf file by chromedp
 * @param {*} printUrl : url to print
 * @param {string} outputPdfFilePath : output pdf file path
 * @return {*}
 */
func (c ChromedpHandler) GenPdf(printUrl, outputPdfFilePath string) error {
	connectUrl, err := c.GetConnUrl()
	if err != nil {
		return err
	}

	ctxTime, cancelTime := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelTime()

	ctx, cancel := chromedp.NewRemoteAllocator(ctxTime, connectUrl.WebSocketDebuggerUrl)
	defer cancel()
	ctx1, cancel1 := chromedp.NewContext(ctx)
	defer func() {
		cancel1()
		_ = chromedp.Cancel(ctx1) // cancel to release resource (a tab in browser)
	}()

	chromedp.Env("TZ=" + "Asia/Shanghai")
	var buf []byte
	err = chromedp.Run(ctx1, chromedp.Tasks{
		chromedp.Navigate(printUrl),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				WithScale(0.8).
				WithMarginLeft(0.6).
				WithMarginRight(0.6).
				Do(ctx)
			return err
		}),
	})
	if err != nil {
		return errors.Wrapf(err, "chromedp Run failed")
	}

	err = os.WriteFile(outputPdfFilePath, buf, 0666)
	if err != nil {
		return errors.Wrapf(err, "write to file failed")
	}
	return nil
}

/** release resource(page etc.) that not need any more,except the blank page
 * @description: 清理资源,仅保留最后一个blank的标签页,其他标签页全部关闭
 * @param {[]ChromedpConnParam} res 资源数据,通过id可以关闭
 * @return {*}
 * 参考: https://github.com/chromedp/chromedp/issues/1096
 */
func (pf ChromedpHandler) Clear(res []ChromedpConnParam) {
	if len(res) <= 1 {
		return
	}
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), res[len(res)-1].WebSocketDebuggerUrl)
	defer cancel()
	ctx1, cancel1 := chromedp.NewContext(ctx)

	ts, err := chromedp.Targets(ctx1)
	if err != nil {
		return
	}

	hasBlank := false
	for _, info := range ts {
		if info.Attached {
			continue
		}
		if info.Title == _blank && !hasBlank { // 保留一个空白标签页
			hasBlank = true
			continue
		}

		newCtx, newCancel := chromedp.NewContext(ctx, chromedp.WithTargetID(info.TargetID))
		_ = chromedp.Run(newCtx, page.Close())
		newCancel()
	}

	defer func() {
		_ = chromedp.Cancel(ctx1)
		cancel1()
	}()
}

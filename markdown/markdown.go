package markdown

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/iris-contrib/blackfriday"
	"github.com/sirupsen/logrus"

	"github.com/mlogclub/simple"
)

// option
type MdOption func(*SimpleMd)

// EnableTOC 开启TOC
func EnableTOC() MdOption {
	return func(md *SimpleMd) {
		md._enableToc = true
	}
}

// SummaryLen 生成摘要的长度
func SummaryLen(summaryLen int) MdOption {
	return func(md *SimpleMd) {
		md._summaryLen = summaryLen
	}
}

// simple md
type SimpleMd struct {
	_summaryLen int  // 摘要长度
	_enableToc  bool // 是否开启TOC
}

// NewMd new simple md
func NewMd(options ...MdOption) *SimpleMd {
	simpleMd := &SimpleMd{
		_summaryLen: 256,
		_enableToc:  false,
	}
	for _, option := range options {
		option(simpleMd)
	}
	return simpleMd
}

// Run
func (md *SimpleMd) Run(mdText string) (htmlStr, summary string) {
	mdText = strings.Replace(mdText, "\r\n", "\n", -1)

	var htmlRenderer blackfriday.Option
	if md._enableToc {
		htmlRenderer = blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
			Flags: blackfriday.TOC,
		}))
	} else {
		htmlRenderer = blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		}))
	}
	data := blackfriday.Run([]byte(mdText), htmlRenderer)
	if doc, err := md.doRender(data); err == nil {
		htmlStr, _ = doc.Find("body").Html()
		if md._summaryLen > 0 {
			summary = md.summaryText(doc)
		}
	} else {
		logrus.Error(err)
	}
	return
}

func (md SimpleMd) doRender(data []byte) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// doc.Find("*").Contents().FilterFunction(func(i int, ele *goquery.Selection) bool {
	// 	if "#text" != goquery.NodeName(ele) {
	// 		return false
	// 	}
	// 	parent := goquery.NodeName(ele.Parent())
	//
	// 	return "span" != parent && "code" != parent && "pre" != parent
	// }).Each(func(i int, ele *goquery.Selection) {
	// 	text := ele.Text()
	// 	text = pangu.SpacingText(text)
	// 	ele.ReplaceWithHtml(text)
	// })
	//
	// doc.Find("code").Each(func(i int, ele *goquery.Selection) {
	// 	code, err := ele.Html()
	// 	if nil != err {
	// 		logrus.Error("get element HTML failed", ele, err)
	// 	} else {
	// 		code = strings.Replace(code, "<", "&lt;", -1)
	// 		code = strings.Replace(code, ">", "&gt;", -1)
	// 		ele.SetHtml(code)
	// 	}
	// })

	if md._enableToc {
		navLi := doc.Find("nav > ul > li")
		// 说明外面有一层空的ul包裹，需要去掉它（这个地方不知道是不是markdown渲染器的BUG）
		if navLi.Size() > 0 && navLi.Size() == 1 && doc.Find("nav > ul > li > a").Size() == 0 {
			if tocHtml, err := navLi.Html(); err != nil {
				doc.Find("nav").First().SetHtml(tocHtml)
			} else {
				logrus.Error(err)
			}
		}
	}
	return doc, nil
}

// summaryText 摘要
func (md *SimpleMd) summaryText(doc *goquery.Document) string {
	if md._summaryLen <= 0 {
		return ""
	}
	doc.Find("nav").Remove()
	text := doc.Text()
	text = strings.TrimSpace(text)
	return simple.GetSummary(text, md._summaryLen)
}

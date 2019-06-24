package simple

import (
	"crypto/md5"
	"github.com/PuerkitoBio/goquery"
	"github.com/iris-contrib/blackfriday"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"github.com/vinta/pangu"
	"regexp"
	"strings"
)

type MarkdownResult struct {
	ContentHtml string
	SummaryText string
	thumbUrl    string
}

// Markdown process the specified markdown text to HTML.
func Markdown(mdText string) *MarkdownResult {
	mdText = strings.Replace(mdText, "\r\n", "\n", -1)

	digest := md5.New()
	digest.Write([]byte(mdText))

	unsafe := blackfriday.Run([]byte(mdText))
	contentHTML := string(unsafe)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(contentHTML))

	// 处理图片
	// doc.Find("img").Each(func(i int, ele *goquery.Selection) {
	// 	src, _ := ele.Attr("src")
	// 	ele.SetAttr("data-src", src)
	// 	ele.RemoveAttr("src")
	// })

	doc.Find("*").Contents().FilterFunction(func(i int, ele *goquery.Selection) bool {
		if "#text" != goquery.NodeName(ele) {
			return false
		}
		parent := goquery.NodeName(ele.Parent())

		return "span" != parent && "code" != parent && "pre" != parent
	}).Each(func(i int, ele *goquery.Selection) {
		text := ele.Text()
		text = pangu.SpacingText(text)
		ele.ReplaceWithHtml(text)
	})

	doc.Find("code").Each(func(i int, ele *goquery.Selection) {
		code, err := ele.Html()
		if nil != err {
			logrus.Error("get element [%+v]' HTML failed: %s", ele, err)
		} else {
			code = strings.Replace(code, "<", "&lt;", -1)
			code = strings.Replace(code, ">", "&gt;", -1)
			ele.SetHtml(code)
		}
	})

	contentHTML, _ = doc.Find("body").Html()
	contentHTML = bluemonday.UGCPolicy().AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code").
		AllowAttrs("data-src").OnElements("img").
		AllowAttrs("class", "target", "id", "style").Globally().
		AllowAttrs("src", "width", "height", "border", "marginwidth", "marginheight").OnElements("iframe").
		AllowAttrs("controls", "src").OnElements("audio").
		AllowAttrs("color").OnElements("font").
		AllowAttrs("controls", "src", "width", "height").OnElements("video").
		AllowAttrs("src", "media", "type").OnElements("source").
		AllowAttrs("width", "height", "data", "type").OnElements("object").
		AllowAttrs("name", "value").OnElements("param").
		AllowAttrs("src", "type", "width", "height", "wmode", "allowNetworking").OnElements("embed").
		Sanitize(contentHTML)

	summaryText := summaryText(doc)

	return &MarkdownResult{
		ContentHtml: contentHTML,
		SummaryText: summaryText,
		thumbUrl:    thumbnailUrl(doc),
	}
}

// 缩略图
func thumbnailUrl(doc *goquery.Document) string {
	selection := doc.Find("img").First()
	thumbnailURL, _ := selection.Attr("src")
	if "" == thumbnailURL {
		thumbnailURL, _ = selection.Attr("data-src")
	}
	return thumbnailURL
}

// 摘要
func summaryText(doc *goquery.Document) string {
	text := doc.Text()
	return GetSummary(text, 256)
}

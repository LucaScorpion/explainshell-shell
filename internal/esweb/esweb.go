package esweb

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const ES_EXPLAIN_URL = "https://explainshell.com/explain?cmd="

func GetCommandHelp(cmd string) (*CommandHelp, error) {
	source := ES_EXPLAIN_URL + url.QueryEscape(cmd)
	doc, err := loadDocument(source)
	if err != nil {
		return nil, err
	}

	help := &CommandHelp{
		Source:  source,
		Command: strings.TrimSpace(doc.Find("#command").Text()),
	}

	doc.Find("[helpref]").Each(func(_ int, sel *goquery.Selection) {
		helpText := ""
		if helpRef, ok := sel.Attr("helpref"); ok {
			helpText = doc.Find("#" + helpRef).Text()
		}

		help.Parts = append(help.Parts, &CommandPart{
			Part: sel.Text(),
			Help: helpText,
		})
	})

	return help, nil
}

func loadDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("unexpected status code " + strconv.Itoa(res.StatusCode))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

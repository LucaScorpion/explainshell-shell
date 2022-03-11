package esweb

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

const ES_EXPLAIN_URL = "https://explainshell.com/explain?cmd="

func DoThing(cmd string) {
	getCommandHelp(cmd)
}

func getCommandHelp(cmd string) (*CommandHelp, error) {
	doc, err := loadDocument(cmd)
	if err != nil {
		return nil, err
	}

	help := &CommandHelp{}
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

func loadDocument(cmd string) (*goquery.Document, error) {
	res, err := http.Get(ES_EXPLAIN_URL + cmd)
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

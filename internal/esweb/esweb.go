package esweb

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const esExplainUrl = "https://explainshell.com/explain?cmd="
const mpUrl = "https://manpages.ubuntu.com/manpages/jammy/en/man{{section}}/{{cmd}}.{{section}}.html"

func GetCommandHelp(cmd string) (*CommandHelp, error) {
	source := esExplainUrl + url.QueryEscape(cmd)
	doc, err := loadDocument(source)
	if err != nil {
		return nil, err
	}

	help := &CommandHelp{
		Source: source,
		Command: strings.Join(
			doc.Find("#command .command0").Map(func(_ int, c *goquery.Selection) string {
				return c.Text()
			}),
			" ",
		),
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

	if len(help.Parts) >= 1 {
		firstPartName := help.Parts[0].Part
		if len(firstPartName) >= 3 && firstPartName[len(firstPartName)-3] == '(' && firstPartName[len(firstPartName)-1] == ')' {
			manPageSection := string(firstPartName[len(firstPartName)-2])
			cmdName := firstPartName[:len(firstPartName)-3]
			manPage := strings.ReplaceAll(mpUrl, "{{section}}", manPageSection)
			manPage = strings.ReplaceAll(manPage, "{{cmd}}", cmdName)
			help.ManPage = manPage
		}
	}

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

package stations

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"tobloggan/code/contracts"
)

const divider = "+++"

type ArticleParser struct{}

func NewArticleParser() *ArticleParser {
	return &ArticleParser{}
}
func (this *ArticleParser) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.SourceFile:
		strInput := strings.Split(string(input), divider)
		if len(strInput) != 2 {
			output(fmt.Errorf("invalid input: %s", input))
			return
		}
		content := strInput[0]
		body := strings.TrimSpace(strInput[1])
		article := contracts.Article{}
		err := json.Unmarshal([]byte(content), &article)
		if err != nil {
			output(err)
			return
		}
		article.Body = body
		output(article)
	default:
		output(input)
	}

	//    TODO: given a contracts.SourceFile, parse the JSON metadata and save the body on a contracts.Article.
	//    input: contracts.SourceFile
	//    input: contracts.Article
}

var errMalformedContent = errors.New("malformed content")

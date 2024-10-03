package stations

import (
	"fmt"

	"tobloggan/code/contracts"
	"tobloggan/code/set"
)

type ArticleValidator struct {
	seenSlugs []string
}

func NewArticleValidator() *ArticleValidator {
	return &ArticleValidator{}
}
func (this *ArticleValidator) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		for _, r := range input.Slug {
			if !validSlugCharacters.Contains(r) {
				output(fmt.Errorf("slug contains invalid character: %s", string(r)))
				return
			}
		}
		if len(input.Title) == 0 {
			output(fmt.Errorf("title is empty"))
			return
		}

		output(input)
	default:
		output(input)
	}
	//    TODO: given a contracts.Article, validate the Slug and the Title fields and emit the contracts.Article (or an error)
	//    input: contracts.Article
	//    output: contracts.Article (or error)
}

var validSlugCharacters = set.New([]rune("abcdefghijklmnopqrstuvwxyz0123456789-/")...)

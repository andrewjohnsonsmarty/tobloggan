package stations

import (
	"testing"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleValidatorFixture(t *testing.T) {
	gunit.Run(new(ArticleValidatorFixture), t)
}

type ArticleValidatorFixture struct {
	StationFixture
}

func (this *ArticleValidatorFixture) Setup() {
	this.station = NewArticleValidator()
}

func (this *ArticleValidatorFixture) TestValidArticle() {
	this.do(articleExpected)
	this.So(len(this.outputs), should.Equal, 1)
	this.So(this.outputs[0], should.HaveSameTypeAs, contracts.Article{})
}

func (this *ArticleValidatorFixture) TestInvalidSlugs() {
	article.Slug = "/article/1!"
	this.do(article)
	this.So(len(this.outputs), should.Equal, 1)
	this.So(this.outputs[0], should.BeError)
}

func (this *ArticleValidatorFixture) TestInvalidTitles() {
	article.Title = ""
	article.Slug = "/article/1"
	this.do(article)
	this.So(len(this.outputs), should.Equal, 1)
	this.So(this.outputs[0], should.BeError)
}

//func (this *ArticleValidatorFixture) TestSlugsMustBeUnique() {}

var article = contracts.Article{
	Draft: false,
	Slug:  "/article/1",
	Title: "Article 1",
	Date:  timeExpected,
	Body:  "The contents of article 1.",
}

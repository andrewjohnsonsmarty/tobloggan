package stations

import (
	"encoding/json"
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleParserFixture(t *testing.T) {
	gunit.Run(new(ArticleParserFixture), t)
}

type ArticleParserFixture struct {
	StationFixture
}

func (this *ArticleParserFixture) Setup() {
	this.station = &ArticleParser{}
}

func (this *ArticleParserFixture) TestBasicReturn() {
	this.do(article1Content)
	this.So(len(this.outputs), should.Equal, 1)
}

func (this *ArticleParserFixture) TestArticleMetaAndContentReadFromDiskAndEmitted() {
	this.do(contracts.SourceFile(article1Content))
	this.So(this.outputs[0], should.Equal, articleExpected)
}

func (this *ArticleParserFixture) TestMissingDivider() {
	this.do(contracts.SourceFile("nothing"))
	this.So(this.outputs[0], should.BeError)
}

func (this *ArticleParserFixture) TestMalformedMetadata() {
	this.do(contracts.SourceFile("malformed+++hello"))
	this.So(this.outputs[0], should.HaveSameTypeAs, &json.SyntaxError{})
}

var timeExpected, _ = time.Parse("2006-01-02T15:04:05Z", "2024-09-04T00:00:00Z")
var articleExpected = contracts.Article{
	Draft: false,
	Slug:  "/article/1",
	Title: "Article 1",
	Date:  timeExpected,
	Body:  "The contents of article 1.",
}

const article1Content = `{
	"date": "2024-09-04T00:00:00Z",
	"slug": "/article/1",
	"title": "Article 1"
}

+++

The contents of article 1.`

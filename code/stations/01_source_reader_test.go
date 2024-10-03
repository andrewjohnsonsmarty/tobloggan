package stations

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

// `Emit unrecognizeed types

func TestSourceReaderFixture(t *testing.T) {
	gunit.Run(new(SourceReaderFixture), t)
}

type SourceReaderFixture struct {
	*gunit.Fixture
	files   fstest.MapFS
	station contracts.Station
	outputs []any
}

func (this *SourceReaderFixture) Setup() {
	this.files = fstest.MapFS{}
	this.station = NewSourceReader(this.files)
}

func (this *SourceReaderFixture) output(v any) {
	this.outputs = append(this.outputs, v)
}

func (this *SourceReaderFixture) TestIOError() {
	this.station.Do(contracts.SourceFilePath("asdfasdf"), this.output)
	this.So(len(this.outputs), should.Equal, 1)
	this.So(this.outputs[0], should.Wrap, fs.ErrNotExist)

}

func (this *SourceReaderFixture) TestSourceFileContentReadFromDiskAndEmitted() {
	this.files["src/article.md"] = &fstest.MapFile{Data: []byte("article content")}
	this.station.Do(contracts.SourceFilePath("src/article.md"), this.output)
	this.So(len(this.outputs), should.Equal, 1)
	this.So(this.outputs[0], should.Equal, contracts.SourceFile("article content"))
}

func (this *SourceReaderFixture) TestUnrecognizedType() {
	this.station.Do(420, this.output)
	this.So(len(this.outputs), should.Equal, 1)
	this.So(this.outputs[0], should.Equal, 420)
}

func (this *SourceReaderFixture) TestReceiveContractsSourceFile() {
	this.files["src/article.md"] = &fstest.MapFile{Data: []byte("article content")}
	this.station.Do(contracts.SourceFilePath("src/article.md"), this.output)
	this.So(len(this.outputs), should.Equal, 1)
	this.So(this.outputs[0], should.HaveSameTypeAs, contracts.SourceFile{})
}

package stations

import (
	"fmt"
	"io/fs"
	"strings"

	"tobloggan/code/contracts"
)

type SourceScanner struct {
	fs fs.FS
}

func NewSourceScanner(fs fs.FS) contracts.Station {
	return &SourceScanner{fs: fs}
}

func (this *SourceScanner) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.SourceDirectory:
		err := fs.WalkDir(this.fs, string(input), func(path string, d fs.DirEntry, err error) error {
			if err == nil {
				if strings.HasSuffix(path, ".md") && strings.Contains(path, "article") {
					output(contracts.SourceFilePath(path))
				}
			}
			return err
		})
		if err != nil {
			output(fmt.Errorf("error walking source directory: %w", err))
		}
	default:
		output(input)
	}
}

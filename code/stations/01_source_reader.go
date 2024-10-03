package stations

import (
	"io/fs"

	"tobloggan/code/contracts"
)

type SourceReader struct {
	files fs.FS
}

func NewSourceReader(files fs.FS) contracts.Station {
	return &SourceReader{files: files}
}

func (this *SourceReader) Do(input any, output func(v any)) {
	switch input := input.(type) {

	case contracts.SourceFilePath:
		file, err := fs.ReadFile(this.files, string(input))
		if err != nil {
			//output(fmt.Errorf("error reading file: %w", err))
			output(err)
			return
		}
		output(contracts.SourceFile(file))

	default:
		output(input)
	}

	//   TODO: given a contracts.SourceFilePath, read its contents and emit contracts.SourceFile
	//    input: contracts.SourceFilePath
	//    output: contracts.SourceFile, or error from fs.ReadFile, or input w/ unrecognized type
}

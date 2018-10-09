package converter

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Converter struct {
	InputFiles      []string
	OutputDirectory string
}

func NewConverter() *Converter {
	return &Converter{}
}

func PathExists(pathTo string) bool {
	_, err := os.Stat(pathTo)
	return !os.IsNotExist(err)
}

func (c *Converter) AddInputFile(pathTo string) (err error) {

	if !PathExists(pathTo) {
		return errors.New(fmt.Sprintf("input file %s does not exist or is not accessible", pathTo))
	}

	if !strings.HasSuffix(strings.ToLower(pathTo), ".m3u8") {
		return errors.New("input file must have suffix .m3u8")
	}

	for _, inputFile := range c.InputFiles {
		if inputFile == pathTo {
			return errors.New(fmt.Sprintf("%s is already added", pathTo))
		}
	}

	c.InputFiles = append(c.InputFiles, pathTo)

	return nil
}

func (c *Converter) SetOutputDirectory(pathTo string) (err error) {

	if !PathExists(pathTo) {
		return errors.New(fmt.Sprintf("output file %s does not exist or is not accessible", pathTo))
	}

	c.OutputDirectory = pathTo
	return nil

}

package converter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func (c *Converter) Convert() (err error) {
	if len(c.InputFiles) == 0 {
		return errors.New("please select at least one .m3u8 file")
	}

	if c.OutputDirectory == "" {
		return errors.New("please select a valid output directory")
	}

	for _, inputFile := range c.InputFiles {

		m3u8Lines, err := c.readLines(inputFile)

		var tracks []string

		if err != nil {
			return err
		}

		for _, m3u8Line := range m3u8Lines {
			if !strings.HasPrefix(m3u8Line, "#") {
				tracks = append(tracks, m3u8Line)
			}
		}

		err = c.writeLines(tracks, c.getFileOut(inputFile))

		if err != nil {
			return err
		}

	}

	return nil
}

func (c *Converter) getFileOut(filePath string) string {

	newFilename := strings.Replace(filepath.Base(filePath), "m3u8", "m3u", -1)
	return path.Join(c.OutputDirectory, newFilename)

}

func (c *Converter) readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (c *Converter) writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
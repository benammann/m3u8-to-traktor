package converter

import (
	"errors"
)

func (c *Converter) Convert() (err error) {
	if len(c.InputFiles) == 0 {
		return errors.New("please select at least one .m3u8 file")
	}

	if c.OutputDirectory == "" {
		return errors.New("please select a valid output directory")
	}

	return nil
}
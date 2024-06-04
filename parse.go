package fontinfo

import (
	"errors"
	"fmt"
	"io"

	"golang.org/x/image/font/sfnt"
)

type fontMetadata struct {
	FontFamily string
}

func readMetadata(r io.ReadSeeker) ([]fontMetadata, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read font: %w", err)
	}
	col, err := sfnt.ParseCollection(data)
	if err != nil {
		return nil, fmt.Errorf("sfnt parse collection: %w", err)
	}
	if col.NumFonts() == 0 {
		return nil, errors.New("no fonts found in file")
	}
	ret := make([]fontMetadata, col.NumFonts())
	for i := 0; i < col.NumFonts(); i++ {
		font, err := col.Font(i)
		if err != nil {
			return nil, fmt.Errorf("font %d of collection: %w", i, err)
		}
		var buf sfnt.Buffer
		family, err := font.Name(&buf, sfnt.NameIDFamily)
		if err != nil {
			return nil, fmt.Errorf("font %d read family: %w", i, err)
		}
		ret[i] = fontMetadata{
			FontFamily: family,
		}
	}
	return ret, nil
}

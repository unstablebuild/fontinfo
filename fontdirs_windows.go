//go:build windows

// Copyright 2016 Florian Pigorsch. All rights reserved.
package fontinfo

import (
	"os"
	"path/filepath"
)

func getFontDirectories() (paths []string) {
	return []string{
		filepath.Join(os.Getenv("windir"), "Fonts"),
		filepath.Join(os.Getenv("localappdata"), "Microsoft", "Windows", "Fonts"),
	}
}

//go:build darwin

// Copyright 2016 Florian Pigorsch. All rights reserved.
package fontinfo

func getFontDirectories() (paths []string) {
	return []string{
		expandUser("~/Library/Fonts/"),
		"/Library/Fonts/",
		"/System/Library/Fonts/",
	}
}

// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package xlsui

import (
	"github.com/mdhender/semver"
)

var (
	// version is the semantic version of the daleri package
	version = semver.Version{Major: 0, Minor: 0, Patch: 0}
)

func Version() semver.Version {
	return version
}

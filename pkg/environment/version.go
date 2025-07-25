/*
Copyright 2024 The Karmada Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package environment

import (
	"fmt"
)

const (
	userAgent = "karmada-dashboard"
	dev       = "0.0.0-dev"
)

var (
	// Version is the version of this binary.
	Version      = dev
	gitVersion   = "v0.0.0-master" // nolint:unused
	gitCommit    = "unknown"       // nolint:unused // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState = "unknown"       // nolint:unused // state of git tree, either "clean" or "dirty"
	buildDate    = "unknown"       // nolint:unused // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// IsDev returns true if the version is dev.
func IsDev() bool {
	return Version == dev
}

// UserAgent returns the user agent of this binary.
func UserAgent() string {
	return fmt.Sprintf("%s:%s", userAgent, Version)
}

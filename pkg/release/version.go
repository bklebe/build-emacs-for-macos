package release

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Errors
var (
	Err             = errors.New("release")
	ErrInvalidName  = fmt.Errorf("%w: invalid name", Err)
	ErrEmptyVersion = fmt.Errorf("%w: empty version", Err)
	ErrNotStableRef = fmt.Errorf(
		"%w: git ref is not stable tagged release", Err,
	)
)

var (
	stableVersion = regexp.MustCompile(`^\d+\.\d+(?:[a-z]+)?$`)
	stableGitRef  = regexp.MustCompile(`^emacs-(\d+\.\d+(?:[a-z]+)?)$`)
)

func VersionToName(version string) (string, error) {
	if version == "" {
		return "", ErrEmptyVersion
	}

	if stableVersion.MatchString(version) ||
		strings.HasSuffix(version, "-pretest") {
		return "Emacs-" + version, nil
	}

	return "Emacs." + version, nil
}

func GitRefToStableVersion(ref string) (string, error) {
	if m := stableGitRef.FindStringSubmatch(ref); len(m) > 1 {
		return m[1], nil
	}

	return "", fmt.Errorf("%w: \"%s\"", ErrNotStableRef, ref)
}

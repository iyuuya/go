package linux

import (
	"errors"
	"strings"

	"github.com/iyuuya/go/ioutil"
	"github.com/iyuuya/go/os"
)

// OSRelease see also https://www.linux.org/docs/man5/os-release.html.
type OSRelease struct {
	Name             string
	Version          string
	ID               string
	VersionID        string
	PrettyName       string
	ANSIColor        string
	CPEName          string
	HomeURL          string
	SupportURL       string
	BugReportURL     string
	PrivacyPolicyURL string
	BuildID          string
	Variant          string
	VariantID        string
}

// NewOSRelease returns OSRelease.
func NewOSRelease() (*OSRelease, error) {
	if !os.IsLinux() {
		return nil, errors.New("OS is not linux")
	}

	lines, err := ioutil.ReadLines("/etc/os-release")
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	osrelease := &OSRelease{}

	for _, line := range lines {
		var value string

		pair := strings.SplitN(line, "=", 2)

		if len(pair) == 2 {
			value = osrelease.parseOSReleaseValue(pair[1])
		}

		osrelease.assign(pair[0], value)
	}

	return osrelease, nil
}

// [todo] - escape.
// [todo] - comment.
func (*OSRelease) parseOSReleaseValue(value string) string {
	vlen := len(value)
	if vlen > 2 && value[0] == '"' && value[vlen-1] == '"' {
		value = value[1 : vlen-1]
	}

	return value
}

func (osrelease *OSRelease) assign(key, value string) {
	switch key {
	case "NAME":
		osrelease.Name = value
	case "VERSION":
		osrelease.Version = value
	case "ID":
		osrelease.ID = value
	case "VERSION_ID":
		osrelease.VersionID = value
	case "PRETTY_NAME":
		osrelease.PrettyName = value
	case "ANSI_COLOR":
		osrelease.ANSIColor = value
	case "CPE_NAME":
		osrelease.CPEName = value
	case "HOME_URL":
		osrelease.HomeURL = value
	case "SUPPORT_URL":
		osrelease.SupportURL = value
	case "BUG_REPORT_URL":
		osrelease.BugReportURL = value
	case "PRIVACY_POLICY_URL":
		osrelease.PrivacyPolicyURL = value
	case "BUILD_ID":
		osrelease.BuildID = value
	case "VARIANT":
		osrelease.Variant = value
	case "VARIANT_ID":
		osrelease.VariantID = value
	}
}

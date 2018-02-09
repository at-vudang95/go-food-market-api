package validator

import "regexp"

const (
	// imageNameRegexString validate request im_name
	imageNameRegexString = "^[a-zA-Z0-9._:-]+$"
)

var (
	imageNameRegex = regexp.MustCompile(imageNameRegexString)
)

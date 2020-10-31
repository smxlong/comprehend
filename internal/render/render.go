package render

import "errors"

// ErrorUnsupportedMimeType is returned when the requested MIME type is unsupported.
var ErrorUnsupportedMimeType = errors.New("unsupported MIME type")

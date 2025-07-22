package e

import "errors"

var PreloadErrInitDBEmptySetting = errors.New("PreloadErrInitDBEmptySetting")
var PreloadErrInitDBConnFailed = errors.New("PreloadErrInitDBConnFailed")
var PreloadErrInitDBTypeUnsupport = errors.New("PreloadErrInitDBTypeUnsupport")

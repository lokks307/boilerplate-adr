package e

import "errors"

var PreloadErrInitDBEmptySetting = errors.New("PreloadErrInitDBEmptySetting")
var PreloadErrInitDBConnFailed = errors.New("PreloadErrInitDBConnFailed")
var PreloadErrInitDBTypeUnsupport = errors.New("PreloadErrInitDBTypeUnsupport")

var ActionErrGetDjsonFromBody1 = errors.New("[action][getDjsonFromBody1] request body is nil ")
var ActionErrGetDjsonFromBody2 = errors.New("[action][getDjsonFromBody2] cannot read body ")
var ActionErrGetDjsonFromBody3 = errors.New("[action][getDjsonFromBody3] cannot parse body raw ")
var ActionErrGetDjsonFromBody4 = errors.New("[action][getDjsonFromBody4] validation failed")
var ActionErrGetDjsonFromBody5 = errors.New("[action][getDjsonFromBody5] validator not exists")

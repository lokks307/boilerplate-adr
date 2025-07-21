package e

import "errors"

var PreloadErrInitDBEmptySetting = errors.New("PreloadErrInitDBEmptySetting")
var PreloadErrInitDBConnFailed = errors.New("PreloadErrInitDBConnFailed")
var PreloadErrInitDBTypeUnsupport = errors.New("PreloadErrInitDBTypeUnsupport")

var ActionErrGetDjsonFromBody1 = errors.New("[action][getDjsonFromBody1] request body is nil ")
var ActionErrGetDjsonFromBody2 = errors.New("[action][getDjsonFromBody2] cannot read body ")
var ActionErrGetDjsonFromBody3 = errors.New("[action][getDjsonFromBody3] validator not exists")
var ActionErrGetDjsonFromBody4 = errors.New("[action][getDjsonFromBody4] validation failed")

var ActionErrInvalidPathParam = errors.New("[Error] Invalid Path Parameter")
var ActionErrInvalidQueryParam = errors.New("[Error] Invalid Query Parameter")
var ActionErrInvalidBody = errors.New("[Error] Invalid Requset Body")
var ActionErrInvalidBodyValue = errors.New("[Error] Invalid Requset Body's Data")

package e

import "errors"

var ErrInitValidator = errors.New("[action][InitValidator] compilation failed")
var ActionErrGetDjsonFromBody1 = errors.New("[action][getDjsonFromBody1] request body is nil ")
var ActionErrGetDjsonFromBody2 = errors.New("[action][getDjsonFromBody2] cannot read body ")
var ActionErrGetDjsonFromBody3 = errors.New("[action][getDjsonFromBody3] validator not exists")
var ActionErrGetDjsonFromBody4 = errors.New("[action][getDjsonFromBody4] validation failed")

var ActionErrInvalidPathParam = errors.New("[Error] Invalid Path Parameter")
var ActionErrInvalidQueryParam = errors.New("[Error] Invalid Query Parameter")
var ActionErrInvalidBody = errors.New("[Error] Invalid Requset Body")
var ActionErrInvalidBodyValue = errors.New("[Error] Invalid Requset Body's Data")

var ActionErrNotFound = errors.New("[Error] Resource Not Found")
var ActionErrExpired = errors.New("[Error] Expired Request")
var ActionErrProcessData = errors.New("[Error] Can't Process Data")
var ActionErrInternalServer = errors.New("[Error] Internal Server")
var ActionErrExternalServer = errors.New("[Error] External Server")

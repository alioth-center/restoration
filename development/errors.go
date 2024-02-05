package restoration

import "github.com/alioth-center/infrastructure/utils/values"

type invalidCallTimeError struct {
	CallTime string `json:"call_time"`
}

func (e *invalidCallTimeError) Error() string {
	return values.BuildStrings("invalid call time: ", e.CallTime)
}

type DialRestorationServiceError struct {
	Endpoint      string `json:"endpoint"`
	ErrorOccurred error  `json:"error_occurred"`
}

func (e *DialRestorationServiceError) Error() string {
	return values.BuildStrings("dial restoration service [", e.Endpoint, "] error: ", e.ErrorOccurred.Error())
}

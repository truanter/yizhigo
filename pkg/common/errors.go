package common

import (
	"fmt"
	"github.com/truanter/yizhigo/pkg/log"
)

type runtimeError struct {
	msg string
}

func (r *runtimeError) Error() string {
	return r.msg
}

func NewRuntimeError(msg string) error {
	log.Logger.Errorf("Runtime error: %s", msg)
	return &runtimeError{msg: msg}
}

func IsRuntimeError(err error) bool {
	_, ok := err.(*runtimeError)
	return ok
}

type tbkError struct {
	code      int
	msg       string
	subCode   string
	subMsg    string
	requestID string
}

func (t *tbkError) Error() string {
	return fmt.Sprintf("code: %d, sub_code: %s, msg: %s, sub_msg: %s, request_id: %s", t.code, t.subCode, t.msg, t.subMsg, t.requestID)
}

func NewTbkError(code int, subCode, msg, subMsg, requestID string) error {
	t := &tbkError{
		code:      code,
		msg:       msg,
		subCode:   subCode,
		subMsg:    subMsg,
		requestID: requestID,
	}
	log.Logger.Errorf("Tbk error: %s", t.Error())
	return t
}

func IsTbkError(err error) bool {
	_, ok := err.(*tbkError)
	return ok
}

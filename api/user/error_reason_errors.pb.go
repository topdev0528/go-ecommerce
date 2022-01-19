// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package user

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsOk(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Ok.String() && e.Code == 200
}

func ErrorOk(format string, args ...interface{}) *errors.Error {
	return errors.New(200, ErrorReason_Ok.String(), fmt.Sprintf(format, args...))
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotFound.String() && e.Code == 404
}

func ErrorNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_NotFound.String(), fmt.Sprintf(format, args...))
}

func IsInvalidParams(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_InvalidParams.String() && e.Code == 400
}

func ErrorInvalidParams(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_InvalidParams.String(), fmt.Sprintf(format, args...))
}

func IsUnauthorized(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Unauthorized.String() && e.Code == 401
}

func ErrorUnauthorized(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_Unauthorized.String(), fmt.Sprintf(format, args...))
}

func IsForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Forbidden.String() && e.Code == 403
}

func ErrorForbidden(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_Forbidden.String(), fmt.Sprintf(format, args...))
}

func IsInternalServer(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_InternalServer.String() && e.Code == 500
}

func ErrorInternalServer(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_InternalServer.String(), fmt.Sprintf(format, args...))
}
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/coupon/service/coupon.proto

package service

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on IdRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IdRequestMultiError, or nil
// if none found.
func (m *IdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := IdRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IdRequestMultiError(errors)
	}

	return nil
}

// IdRequestMultiError is an error wrapping multiple validation errors returned
// by IdRequest.ValidateAll() if the designated constraints aren't met.
type IdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdRequestMultiError) AllErrors() []error { return m }

// IdRequestValidationError is the validation error returned by
// IdRequest.Validate if the designated constraints aren't met.
type IdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdRequestValidationError) ErrorName() string { return "IdRequestValidationError" }

// Error satisfies the builtin error interface
func (e IdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdRequestValidationError{}

// Validate checks the field values on Coupons with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Coupons) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Coupons with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CouponsMultiError, or nil if none found.
func (m *Coupons) ValidateAll() error {
	return m.validate(true)
}

func (m *Coupons) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCoupon() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CouponsValidationError{
						field:  fmt.Sprintf("Coupon[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CouponsValidationError{
						field:  fmt.Sprintf("Coupon[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CouponsValidationError{
					field:  fmt.Sprintf("Coupon[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CouponsMultiError(errors)
	}

	return nil
}

// CouponsMultiError is an error wrapping multiple validation errors returned
// by Coupons.ValidateAll() if the designated constraints aren't met.
type CouponsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CouponsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CouponsMultiError) AllErrors() []error { return m }

// CouponsValidationError is the validation error returned by Coupons.Validate
// if the designated constraints aren't met.
type CouponsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CouponsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CouponsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CouponsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CouponsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CouponsValidationError) ErrorName() string { return "CouponsValidationError" }

// Error satisfies the builtin error interface
func (e CouponsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCoupons.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CouponsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CouponsValidationError{}

// Validate checks the field values on CouponBo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CouponBo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CouponBo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CouponBoMultiError, or nil
// if none found.
func (m *CouponBo) ValidateAll() error {
	return m.validate(true)
}

func (m *CouponBo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for StartTime

	// no validation rules for EndTime

	// no validation rules for Description

	// no validation rules for FullMoney

	// no validation rules for Minus

	// no validation rules for Rate

	// no validation rules for Type

	// no validation rules for Remark

	// no validation rules for WholeStore

	if len(errors) > 0 {
		return CouponBoMultiError(errors)
	}

	return nil
}

// CouponBoMultiError is an error wrapping multiple validation errors returned
// by CouponBo.ValidateAll() if the designated constraints aren't met.
type CouponBoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CouponBoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CouponBoMultiError) AllErrors() []error { return m }

// CouponBoValidationError is the validation error returned by
// CouponBo.Validate if the designated constraints aren't met.
type CouponBoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CouponBoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CouponBoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CouponBoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CouponBoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CouponBoValidationError) ErrorName() string { return "CouponBoValidationError" }

// Error satisfies the builtin error interface
func (e CouponBoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCouponBo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CouponBoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CouponBoValidationError{}
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/blog/v1/blog.proto

package v1

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

// Validate checks the field values on BlogDetail with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BlogDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BlogDetail with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BlogDetailMultiError, or
// nil if none found.
func (m *BlogDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *BlogDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Content

	// no validation rules for Like

	if len(errors) > 0 {
		return BlogDetailMultiError(errors)
	}
	return nil
}

// BlogDetailMultiError is an error wrapping multiple validation errors
// returned by BlogDetail.ValidateAll() if the designated constraints aren't met.
type BlogDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BlogDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BlogDetailMultiError) AllErrors() []error { return m }

// BlogDetailValidationError is the validation error returned by
// BlogDetail.Validate if the designated constraints aren't met.
type BlogDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BlogDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BlogDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BlogDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BlogDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BlogDetailValidationError) ErrorName() string { return "BlogDetailValidationError" }

// Error satisfies the builtin error interface
func (e BlogDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBlogDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BlogDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BlogDetailValidationError{}

// Validate checks the field values on CreateBlogRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateBlogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateBlogRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateBlogRequestMultiError, or nil if none found.
func (m *CreateBlogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateBlogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTitle()); l < 5 || l > 50 {
		err := CreateBlogRequestValidationError{
			field:  "Title",
			reason: "value length must be between 5 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Content

	if len(errors) > 0 {
		return CreateBlogRequestMultiError(errors)
	}
	return nil
}

// CreateBlogRequestMultiError is an error wrapping multiple validation errors
// returned by CreateBlogRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateBlogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateBlogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateBlogRequestMultiError) AllErrors() []error { return m }

// CreateBlogRequestValidationError is the validation error returned by
// CreateBlogRequest.Validate if the designated constraints aren't met.
type CreateBlogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateBlogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateBlogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateBlogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateBlogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateBlogRequestValidationError) ErrorName() string {
	return "CreateBlogRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateBlogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateBlogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateBlogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateBlogRequestValidationError{}

// Validate checks the field values on CreateBlogReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateBlogReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateBlogReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateBlogReplyMultiError, or nil if none found.
func (m *CreateBlogReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateBlogReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBlogDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateBlogReplyValidationError{
					field:  "BlogDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateBlogReplyValidationError{
					field:  "BlogDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBlogDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateBlogReplyValidationError{
				field:  "BlogDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateBlogReplyMultiError(errors)
	}
	return nil
}

// CreateBlogReplyMultiError is an error wrapping multiple validation errors
// returned by CreateBlogReply.ValidateAll() if the designated constraints
// aren't met.
type CreateBlogReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateBlogReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateBlogReplyMultiError) AllErrors() []error { return m }

// CreateBlogReplyValidationError is the validation error returned by
// CreateBlogReply.Validate if the designated constraints aren't met.
type CreateBlogReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateBlogReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateBlogReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateBlogReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateBlogReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateBlogReplyValidationError) ErrorName() string { return "CreateBlogReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateBlogReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateBlogReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateBlogReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateBlogReplyValidationError{}

// Validate checks the field values on DeleteBlogRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteBlogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteBlogRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteBlogRequestMultiError, or nil if none found.
func (m *DeleteBlogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteBlogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteBlogRequestMultiError(errors)
	}
	return nil
}

// DeleteBlogRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteBlogRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteBlogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteBlogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteBlogRequestMultiError) AllErrors() []error { return m }

// DeleteBlogRequestValidationError is the validation error returned by
// DeleteBlogRequest.Validate if the designated constraints aren't met.
type DeleteBlogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteBlogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteBlogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteBlogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteBlogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteBlogRequestValidationError) ErrorName() string {
	return "DeleteBlogRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteBlogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteBlogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteBlogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteBlogRequestValidationError{}

// Validate checks the field values on DeleteBlogReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteBlogReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteBlogReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteBlogReplyMultiError, or nil if none found.
func (m *DeleteBlogReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteBlogReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteBlogReplyMultiError(errors)
	}
	return nil
}

// DeleteBlogReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteBlogReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteBlogReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteBlogReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteBlogReplyMultiError) AllErrors() []error { return m }

// DeleteBlogReplyValidationError is the validation error returned by
// DeleteBlogReply.Validate if the designated constraints aren't met.
type DeleteBlogReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteBlogReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteBlogReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteBlogReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteBlogReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteBlogReplyValidationError) ErrorName() string { return "DeleteBlogReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteBlogReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteBlogReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteBlogReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteBlogReplyValidationError{}

// Validate checks the field values on UpdateBlogRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateBlogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateBlogRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateBlogRequestMultiError, or nil if none found.
func (m *UpdateBlogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateBlogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateBlogRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetTitle()); l < 5 || l > 50 {
		err := UpdateBlogRequestValidationError{
			field:  "Title",
			reason: "value length must be between 5 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Content

	if len(errors) > 0 {
		return UpdateBlogRequestMultiError(errors)
	}
	return nil
}

// UpdateBlogRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateBlogRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateBlogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateBlogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateBlogRequestMultiError) AllErrors() []error { return m }

// UpdateBlogRequestValidationError is the validation error returned by
// UpdateBlogRequest.Validate if the designated constraints aren't met.
type UpdateBlogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateBlogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateBlogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateBlogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateBlogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateBlogRequestValidationError) ErrorName() string {
	return "UpdateBlogRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateBlogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateBlogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateBlogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateBlogRequestValidationError{}

// Validate checks the field values on UpdateBlogReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateBlogReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateBlogReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateBlogReplyMultiError, or nil if none found.
func (m *UpdateBlogReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateBlogReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBlogDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateBlogReplyValidationError{
					field:  "BlogDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateBlogReplyValidationError{
					field:  "BlogDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBlogDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateBlogReplyValidationError{
				field:  "BlogDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateBlogReplyMultiError(errors)
	}
	return nil
}

// UpdateBlogReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateBlogReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateBlogReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateBlogReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateBlogReplyMultiError) AllErrors() []error { return m }

// UpdateBlogReplyValidationError is the validation error returned by
// UpdateBlogReply.Validate if the designated constraints aren't met.
type UpdateBlogReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateBlogReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateBlogReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateBlogReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateBlogReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateBlogReplyValidationError) ErrorName() string { return "UpdateBlogReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateBlogReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateBlogReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateBlogReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateBlogReplyValidationError{}

// Validate checks the field values on GetBlogRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetBlogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBlogRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetBlogRequestMultiError,
// or nil if none found.
func (m *GetBlogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBlogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetBlogRequestMultiError(errors)
	}
	return nil
}

// GetBlogRequestMultiError is an error wrapping multiple validation errors
// returned by GetBlogRequest.ValidateAll() if the designated constraints
// aren't met.
type GetBlogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBlogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBlogRequestMultiError) AllErrors() []error { return m }

// GetBlogRequestValidationError is the validation error returned by
// GetBlogRequest.Validate if the designated constraints aren't met.
type GetBlogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBlogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBlogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBlogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBlogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBlogRequestValidationError) ErrorName() string { return "GetBlogRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetBlogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBlogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBlogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBlogRequestValidationError{}

// Validate checks the field values on GetBlogReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetBlogReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBlogReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetBlogReplyMultiError, or
// nil if none found.
func (m *GetBlogReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBlogReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBlogDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetBlogReplyValidationError{
					field:  "BlogDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetBlogReplyValidationError{
					field:  "BlogDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBlogDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetBlogReplyValidationError{
				field:  "BlogDetail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetBlogReplyMultiError(errors)
	}
	return nil
}

// GetBlogReplyMultiError is an error wrapping multiple validation errors
// returned by GetBlogReply.ValidateAll() if the designated constraints aren't met.
type GetBlogReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBlogReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBlogReplyMultiError) AllErrors() []error { return m }

// GetBlogReplyValidationError is the validation error returned by
// GetBlogReply.Validate if the designated constraints aren't met.
type GetBlogReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBlogReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBlogReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBlogReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBlogReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBlogReplyValidationError) ErrorName() string { return "GetBlogReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetBlogReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBlogReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBlogReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBlogReplyValidationError{}

// Validate checks the field values on ListBlogRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListBlogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListBlogRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListBlogRequestMultiError, or nil if none found.
func (m *ListBlogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListBlogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListBlogRequestMultiError(errors)
	}
	return nil
}

// ListBlogRequestMultiError is an error wrapping multiple validation errors
// returned by ListBlogRequest.ValidateAll() if the designated constraints
// aren't met.
type ListBlogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListBlogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListBlogRequestMultiError) AllErrors() []error { return m }

// ListBlogRequestValidationError is the validation error returned by
// ListBlogRequest.Validate if the designated constraints aren't met.
type ListBlogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBlogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBlogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBlogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBlogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBlogRequestValidationError) ErrorName() string { return "ListBlogRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListBlogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBlogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBlogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBlogRequestValidationError{}

// Validate checks the field values on ListBlogReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListBlogReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListBlogReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListBlogReplyMultiError, or
// nil if none found.
func (m *ListBlogReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListBlogReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListBlogReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListBlogReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListBlogReplyValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListBlogReplyMultiError(errors)
	}
	return nil
}

// ListBlogReplyMultiError is an error wrapping multiple validation errors
// returned by ListBlogReply.ValidateAll() if the designated constraints
// aren't met.
type ListBlogReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListBlogReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListBlogReplyMultiError) AllErrors() []error { return m }

// ListBlogReplyValidationError is the validation error returned by
// ListBlogReply.Validate if the designated constraints aren't met.
type ListBlogReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBlogReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBlogReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBlogReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBlogReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBlogReplyValidationError) ErrorName() string { return "ListBlogReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListBlogReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBlogReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBlogReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBlogReplyValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: account.proto

package account

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

// Validate checks the field values on EditPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EditPasswordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EditPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EditPasswordRequestMultiError, or nil if none found.
func (m *EditPasswordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *EditPasswordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCode() != "" {

		if utf8.RuneCountInString(m.GetCode()) < 1 {
			err := EditPasswordRequestValidationError{
				field:  "Code",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPhone() != "" {

		if utf8.RuneCountInString(m.GetPhone()) != 11 {
			err := EditPasswordRequestValidationError{
				field:  "Phone",
				reason: "value length must be 11 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)

		}

		if !_EditPasswordRequest_Phone_Pattern.MatchString(m.GetPhone()) {
			err := EditPasswordRequestValidationError{
				field:  "Phone",
				reason: "value does not match regex pattern \"^[0-9]{11}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetEmial() != "" {

		if err := m._validateEmail(m.GetEmial()); err != nil {
			err = EditPasswordRequestValidationError{
				field:  "Emial",
				reason: "value must be a valid email address",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPassword() != "" {

		if utf8.RuneCountInString(m.GetPassword()) < 8 {
			err := EditPasswordRequestValidationError{
				field:  "Password",
				reason: "value length must be at least 8 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return EditPasswordRequestMultiError(errors)
	}

	return nil
}

func (m *EditPasswordRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *EditPasswordRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// EditPasswordRequestMultiError is an error wrapping multiple validation
// errors returned by EditPasswordRequest.ValidateAll() if the designated
// constraints aren't met.
type EditPasswordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EditPasswordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EditPasswordRequestMultiError) AllErrors() []error { return m }

// EditPasswordRequestValidationError is the validation error returned by
// EditPasswordRequest.Validate if the designated constraints aren't met.
type EditPasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EditPasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EditPasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EditPasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EditPasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EditPasswordRequestValidationError) ErrorName() string {
	return "EditPasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e EditPasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEditPasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EditPasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EditPasswordRequestValidationError{}

var _EditPasswordRequest_Phone_Pattern = regexp.MustCompile("^[0-9]{11}$")

// Validate checks the field values on LoginOutRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginOutRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginOutRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginOutRequestMultiError, or nil if none found.
func (m *LoginOutRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginOutRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LoginOutRequestMultiError(errors)
	}

	return nil
}

// LoginOutRequestMultiError is an error wrapping multiple validation errors
// returned by LoginOutRequest.ValidateAll() if the designated constraints
// aren't met.
type LoginOutRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginOutRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginOutRequestMultiError) AllErrors() []error { return m }

// LoginOutRequestValidationError is the validation error returned by
// LoginOutRequest.Validate if the designated constraints aren't met.
type LoginOutRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginOutRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginOutRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginOutRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginOutRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginOutRequestValidationError) ErrorName() string { return "LoginOutRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginOutRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginOutRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginOutRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginOutRequestValidationError{}

// Validate checks the field values on UserRegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserRegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRegisterRequestMultiError, or nil if none found.
func (m *UserRegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPhone() != "" {

		if utf8.RuneCountInString(m.GetPhone()) != 11 {
			err := UserRegisterRequestValidationError{
				field:  "Phone",
				reason: "value length must be 11 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)

		}

		if !_UserRegisterRequest_Phone_Pattern.MatchString(m.GetPhone()) {
			err := UserRegisterRequestValidationError{
				field:  "Phone",
				reason: "value does not match regex pattern \"^[0-9]{11}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetCode() != "" {

		if utf8.RuneCountInString(m.GetCode()) < 1 {
			err := UserRegisterRequestValidationError{
				field:  "Code",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetAccount() != "" {

		if l := utf8.RuneCountInString(m.GetAccount()); l < 3 || l > 20 {
			err := UserRegisterRequestValidationError{
				field:  "Account",
				reason: "value length must be between 3 and 20 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPassword() != "" {

		if utf8.RuneCountInString(m.GetPassword()) < 8 {
			err := UserRegisterRequestValidationError{
				field:  "Password",
				reason: "value length must be at least 8 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetEmail() != "" {

		if err := m._validateEmail(m.GetEmail()); err != nil {
			err = UserRegisterRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetRegisterType() != "" {

		if utf8.RuneCountInString(m.GetRegisterType()) < 1 {
			err := UserRegisterRequestValidationError{
				field:  "RegisterType",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return UserRegisterRequestMultiError(errors)
	}

	return nil
}

func (m *UserRegisterRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UserRegisterRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UserRegisterRequestMultiError is an error wrapping multiple validation
// errors returned by UserRegisterRequest.ValidateAll() if the designated
// constraints aren't met.
type UserRegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRegisterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRegisterRequestMultiError) AllErrors() []error { return m }

// UserRegisterRequestValidationError is the validation error returned by
// UserRegisterRequest.Validate if the designated constraints aren't met.
type UserRegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRegisterRequestValidationError) ErrorName() string {
	return "UserRegisterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UserRegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRegisterRequestValidationError{}

var _UserRegisterRequest_Phone_Pattern = regexp.MustCompile("^[0-9]{11}$")

// Validate checks the field values on UserLoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserLoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserLoginRequestMultiError, or nil if none found.
func (m *UserLoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserLoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAccount() != "" {

		if utf8.RuneCountInString(m.GetAccount()) < 1 {
			err := UserLoginRequestValidationError{
				field:  "Account",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPassword() != "" {

		if utf8.RuneCountInString(m.GetPassword()) < 0 {
			err := UserLoginRequestValidationError{
				field:  "Password",
				reason: "value length must be at least 0 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPhone() != "" {

		if utf8.RuneCountInString(m.GetPhone()) != 11 {
			err := UserLoginRequestValidationError{
				field:  "Phone",
				reason: "value length must be 11 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)

		}

		if !_UserLoginRequest_Phone_Pattern.MatchString(m.GetPhone()) {
			err := UserLoginRequestValidationError{
				field:  "Phone",
				reason: "value does not match regex pattern \"^[0-9]{11}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetLoginType() != "" {

		if utf8.RuneCountInString(m.GetLoginType()) < 1 {
			err := UserLoginRequestValidationError{
				field:  "LoginType",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetEmail() != "" {

		if err := m._validateEmail(m.GetEmail()); err != nil {
			err = UserLoginRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetCode() != "" {

		if utf8.RuneCountInString(m.GetCode()) < 1 {
			err := UserLoginRequestValidationError{
				field:  "Code",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return UserLoginRequestMultiError(errors)
	}

	return nil
}

func (m *UserLoginRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UserLoginRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UserLoginRequestMultiError is an error wrapping multiple validation errors
// returned by UserLoginRequest.ValidateAll() if the designated constraints
// aren't met.
type UserLoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserLoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserLoginRequestMultiError) AllErrors() []error { return m }

// UserLoginRequestValidationError is the validation error returned by
// UserLoginRequest.Validate if the designated constraints aren't met.
type UserLoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserLoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserLoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserLoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserLoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserLoginRequestValidationError) ErrorName() string { return "UserLoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserLoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserLoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserLoginRequestValidationError{}

var _UserLoginRequest_Phone_Pattern = regexp.MustCompile("^[0-9]{11}$")

// Validate checks the field values on UpdateWarnOpenReqeust with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateWarnOpenReqeust) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateWarnOpenReqeust with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateWarnOpenReqeustMultiError, or nil if none found.
func (m *UpdateWarnOpenReqeust) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateWarnOpenReqeust) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WarnOpen

	if len(errors) > 0 {
		return UpdateWarnOpenReqeustMultiError(errors)
	}

	return nil
}

// UpdateWarnOpenReqeustMultiError is an error wrapping multiple validation
// errors returned by UpdateWarnOpenReqeust.ValidateAll() if the designated
// constraints aren't met.
type UpdateWarnOpenReqeustMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateWarnOpenReqeustMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateWarnOpenReqeustMultiError) AllErrors() []error { return m }

// UpdateWarnOpenReqeustValidationError is the validation error returned by
// UpdateWarnOpenReqeust.Validate if the designated constraints aren't met.
type UpdateWarnOpenReqeustValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateWarnOpenReqeustValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateWarnOpenReqeustValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateWarnOpenReqeustValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateWarnOpenReqeustValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateWarnOpenReqeustValidationError) ErrorName() string {
	return "UpdateWarnOpenReqeustValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateWarnOpenReqeustValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateWarnOpenReqeust.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateWarnOpenReqeustValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateWarnOpenReqeustValidationError{}

// Validate checks the field values on WechatLoginCallbackRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WechatLoginCallbackRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WechatLoginCallbackRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WechatLoginCallbackRequestMultiError, or nil if none found.
func (m *WechatLoginCallbackRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WechatLoginCallbackRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetCode()) < 1 {
		err := WechatLoginCallbackRequestValidationError{
			field:  "Code",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetState()) < 1 {
		err := WechatLoginCallbackRequestValidationError{
			field:  "State",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return WechatLoginCallbackRequestMultiError(errors)
	}

	return nil
}

// WechatLoginCallbackRequestMultiError is an error wrapping multiple
// validation errors returned by WechatLoginCallbackRequest.ValidateAll() if
// the designated constraints aren't met.
type WechatLoginCallbackRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WechatLoginCallbackRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WechatLoginCallbackRequestMultiError) AllErrors() []error { return m }

// WechatLoginCallbackRequestValidationError is the validation error returned
// by WechatLoginCallbackRequest.Validate if the designated constraints aren't met.
type WechatLoginCallbackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WechatLoginCallbackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WechatLoginCallbackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WechatLoginCallbackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WechatLoginCallbackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WechatLoginCallbackRequestValidationError) ErrorName() string {
	return "WechatLoginCallbackRequestValidationError"
}

// Error satisfies the builtin error interface
func (e WechatLoginCallbackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWechatLoginCallbackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WechatLoginCallbackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WechatLoginCallbackRequestValidationError{}

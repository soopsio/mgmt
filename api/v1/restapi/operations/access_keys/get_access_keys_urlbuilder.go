// Code generated by go-swagger; DO NOT EDIT.

package access_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetAccessKeysURL generates an URL for the get access keys operation
type GetAccessKeysURL struct {
	AccessKeyID      string
	Signature        string
	SignatureMethod  string
	SignatureNonce   string
	SignatureVersion string
	Timestamp        string
	Version          string
	Action           *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAccessKeysURL) WithBasePath(bp string) *GetAccessKeysURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAccessKeysURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetAccessKeysURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/accesskeys"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	accessKeyID := o.AccessKeyID
	if accessKeyID != "" {
		qs.Set("AccessKeyId", accessKeyID)
	}

	signature := o.Signature
	if signature != "" {
		qs.Set("Signature", signature)
	}

	signatureMethod := o.SignatureMethod
	if signatureMethod != "" {
		qs.Set("SignatureMethod", signatureMethod)
	}

	signatureNonce := o.SignatureNonce
	if signatureNonce != "" {
		qs.Set("SignatureNonce", signatureNonce)
	}

	signatureVersion := o.SignatureVersion
	if signatureVersion != "" {
		qs.Set("SignatureVersion", signatureVersion)
	}

	timestamp := o.Timestamp
	if timestamp != "" {
		qs.Set("Timestamp", timestamp)
	}

	version := o.Version
	if version != "" {
		qs.Set("Version", version)
	}

	var action string
	if o.Action != nil {
		action = *o.Action
	}
	if action != "" {
		qs.Set("action", action)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetAccessKeysURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetAccessKeysURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetAccessKeysURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetAccessKeysURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetAccessKeysURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetAccessKeysURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetPolicyURL generates an URL for the get policy operation
type GetPolicyURL struct {
	ID string

	AccessKeyID      string
	Signature        string
	SignatureMethod  string
	SignatureNonce   string
	SignatureVersion string
	Timestamp        string
	Version          string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPolicyURL) WithBasePath(bp string) *GetPolicyURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPolicyURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetPolicyURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/policies/{id}"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("ID is required on GetPolicyURL")
	}
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

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetPolicyURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetPolicyURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetPolicyURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetPolicyURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetPolicyURL")
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
func (o *GetPolicyURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

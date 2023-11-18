// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// GetCustomerByIDParams is parameters of getCustomerByID operation.
type GetCustomerByIDParams struct {
	CustomerID int64
}

func unpackGetCustomerByIDParams(packed middleware.Parameters) (params GetCustomerByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "customerID",
			In:   "path",
		}
		params.CustomerID = packed[key].(int64)
	}
	return params
}

func decodeGetCustomerByIDParams(args [1]string, argsEscaped bool, r *http.Request) (params GetCustomerByIDParams, _ error) {
	// Decode path: customerID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "customerID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CustomerID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "customerID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PutModifyCustomerByIDParams is parameters of putModifyCustomerByID operation.
type PutModifyCustomerByIDParams struct {
	CustomerID int64
}

func unpackPutModifyCustomerByIDParams(packed middleware.Parameters) (params PutModifyCustomerByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "customerID",
			In:   "path",
		}
		params.CustomerID = packed[key].(int64)
	}
	return params
}

func decodePutModifyCustomerByIDParams(args [1]string, argsEscaped bool, r *http.Request) (params PutModifyCustomerByIDParams, _ error) {
	// Decode path: customerID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "customerID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.CustomerID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "customerID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
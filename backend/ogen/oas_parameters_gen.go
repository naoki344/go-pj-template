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

// GetNoteByIDParams is parameters of getNoteByID operation.
type GetNoteByIDParams struct {
	NoteID int64
}

func unpackGetNoteByIDParams(packed middleware.Parameters) (params GetNoteByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "noteID",
			In:   "path",
		}
		params.NoteID = packed[key].(int64)
	}
	return params
}

func decodeGetNoteByIDParams(args [1]string, argsEscaped bool, r *http.Request) (params GetNoteByIDParams, _ error) {
	// Decode path: noteID.
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
				Param:   "noteID",
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

				params.NoteID = c
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
			Name: "noteID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
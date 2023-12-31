// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func encodeGetCustomerByIDResponse(response GetCustomerByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CustomerHeaders:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetCustomerByIDBadRequest:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetCustomerByIDNotFound:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetCustomerByIDInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostCreateCustomerResponse(response PostCreateCustomerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CustomerHeaders:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PostCreateCustomerBadRequest:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PostCreateCustomerInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostSearchCustomerResponse(response PostSearchCustomerRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PostSearchCustomer200ResponseHeaders:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PostSearchCustomerBadRequest:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PostSearchCustomerNotFound:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PostSearchCustomerInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePutModifyCustomerByIDResponse(response PutModifyCustomerByIDRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CustomerHeaders:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PutModifyCustomerByIDBadRequest:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PutModifyCustomerByIDNotFound:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PutModifyCustomerByIDInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Access-Control-Allow-Headers" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Headers",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowHeaders.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Headers header")
				}
			}
			// Encode "Access-Control-Allow-Methods" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Methods",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowMethods.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Methods header")
				}
			}
			// Encode "Access-Control-Allow-Origin" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Access-Control-Allow-Origin",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AccessControlAllowOrigin.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Access-Control-Allow-Origin header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

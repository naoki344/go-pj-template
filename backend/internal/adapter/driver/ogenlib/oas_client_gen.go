// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// GetCustomerByID invokes getCustomerByID operation.
	//
	// 顧客情報参照.
	//
	// GET /customers/{customerID}
	GetCustomerByID(ctx context.Context, params GetCustomerByIDParams) (GetCustomerByIDRes, error)
	// PostCreateCustomer invokes postCreateCustomer operation.
	//
	// 顧客情報登録.
	//
	// POST /customers
	PostCreateCustomer(ctx context.Context, request *PostCreateCustomerRequest) (PostCreateCustomerRes, error)
	// PostSearchCustomer invokes postSearchCustomer operation.
	//
	// 顧客情報検索.
	//
	// POST /customers/search
	PostSearchCustomer(ctx context.Context, request *PostSearchCustomerRequest) (PostSearchCustomerRes, error)
	// PutModifyCustomerByID invokes putModifyCustomerByID operation.
	//
	// 顧客情報更新.
	//
	// PUT /customers/{customerID}
	PutModifyCustomerByID(ctx context.Context, request *PutModifyCustomerByIDRequest, params PutModifyCustomerByIDParams) (PutModifyCustomerByIDRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// GetCustomerByID invokes getCustomerByID operation.
//
// 顧客情報参照.
//
// GET /customers/{customerID}
func (c *Client) GetCustomerByID(ctx context.Context, params GetCustomerByIDParams) (GetCustomerByIDRes, error) {
	res, err := c.sendGetCustomerByID(ctx, params)
	return res, err
}

func (c *Client) sendGetCustomerByID(ctx context.Context, params GetCustomerByIDParams) (res GetCustomerByIDRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getCustomerByID"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/customers/{customerID}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "GetCustomerByID",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/customers/"
	{
		// Encode "customerID" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "customerID",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CustomerID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetCustomerByIDResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PostCreateCustomer invokes postCreateCustomer operation.
//
// 顧客情報登録.
//
// POST /customers
func (c *Client) PostCreateCustomer(ctx context.Context, request *PostCreateCustomerRequest) (PostCreateCustomerRes, error) {
	res, err := c.sendPostCreateCustomer(ctx, request)
	return res, err
}

func (c *Client) sendPostCreateCustomer(ctx context.Context, request *PostCreateCustomerRequest) (res PostCreateCustomerRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("postCreateCustomer"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/customers"),
	}
	// Validate request before sending.
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "PostCreateCustomer",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/customers"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodePostCreateCustomerRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodePostCreateCustomerResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PostSearchCustomer invokes postSearchCustomer operation.
//
// 顧客情報検索.
//
// POST /customers/search
func (c *Client) PostSearchCustomer(ctx context.Context, request *PostSearchCustomerRequest) (PostSearchCustomerRes, error) {
	res, err := c.sendPostSearchCustomer(ctx, request)
	return res, err
}

func (c *Client) sendPostSearchCustomer(ctx context.Context, request *PostSearchCustomerRequest) (res PostSearchCustomerRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("postSearchCustomer"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/customers/search"),
	}
	// Validate request before sending.
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "PostSearchCustomer",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/customers/search"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodePostSearchCustomerRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodePostSearchCustomerResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PutModifyCustomerByID invokes putModifyCustomerByID operation.
//
// 顧客情報更新.
//
// PUT /customers/{customerID}
func (c *Client) PutModifyCustomerByID(ctx context.Context, request *PutModifyCustomerByIDRequest, params PutModifyCustomerByIDParams) (PutModifyCustomerByIDRes, error) {
	res, err := c.sendPutModifyCustomerByID(ctx, request, params)
	return res, err
}

func (c *Client) sendPutModifyCustomerByID(ctx context.Context, request *PutModifyCustomerByIDRequest, params PutModifyCustomerByIDParams) (res PutModifyCustomerByIDRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putModifyCustomerByID"),
		semconv.HTTPMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/customers/{customerID}"),
	}
	// Validate request before sending.
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "PutModifyCustomerByID",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/customers/"
	{
		// Encode "customerID" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "customerID",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CustomerID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodePutModifyCustomerByIDRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodePutModifyCustomerByIDResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

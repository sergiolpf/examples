// Code generated by goa v3.0.6, DO NOT EDIT.
//
// chatter client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o
// $(GOPATH)/src/goa.design/examples/streaming

package client

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the chatter service endpoint HTTP clients.
type Client struct {
	// Login Doer is the HTTP client used to make requests to the login endpoint.
	LoginDoer goahttp.Doer

	// Echoer Doer is the HTTP client used to make requests to the echoer endpoint.
	EchoerDoer goahttp.Doer

	// Listener Doer is the HTTP client used to make requests to the listener
	// endpoint.
	ListenerDoer goahttp.Doer

	// Summary Doer is the HTTP client used to make requests to the summary
	// endpoint.
	SummaryDoer goahttp.Doer

	// Subscribe Doer is the HTTP client used to make requests to the subscribe
	// endpoint.
	SubscribeDoer goahttp.Doer

	// History Doer is the HTTP client used to make requests to the history
	// endpoint.
	HistoryDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme     string
	host       string
	encoder    func(*http.Request) goahttp.Encoder
	decoder    func(*http.Response) goahttp.Decoder
	dialer     goahttp.Dialer
	configurer *ConnConfigurer
}

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "chatter" service.
type ConnConfigurer struct {
	EchoerFn    goahttp.ConnConfigureFunc
	ListenerFn  goahttp.ConnConfigureFunc
	SummaryFn   goahttp.ConnConfigureFunc
	SubscribeFn goahttp.ConnConfigureFunc
	HistoryFn   goahttp.ConnConfigureFunc
}

// EchoerClientStream implements the chatter.EchoerClientStream interface.
type EchoerClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// ListenerClientStream implements the chatter.ListenerClientStream interface.
type ListenerClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// SummaryClientStream implements the chatter.SummaryClientStream interface.
type SummaryClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// SubscribeClientStream implements the chatter.SubscribeClientStream interface.
type SubscribeClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// HistoryClientStream implements the chatter.HistoryClientStream interface.
type HistoryClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
	// view is the view to render  result type before sending to the websocket
	// connection.
	view string
}

// NewClient instantiates HTTP clients for all the chatter service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
	dialer goahttp.Dialer,
	cfn *ConnConfigurer,
) *Client {
	if cfn == nil {
		cfn = &ConnConfigurer{}
	}
	return &Client{
		LoginDoer:           doer,
		EchoerDoer:          doer,
		ListenerDoer:        doer,
		SummaryDoer:         doer,
		SubscribeDoer:       doer,
		HistoryDoer:         doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
		dialer:              dialer,
		configurer:          cfn,
	}
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "chatter" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		EchoerFn:    fn,
		ListenerFn:  fn,
		SummaryFn:   fn,
		SubscribeFn: fn,
		HistoryFn:   fn,
	}
}

// Login returns an endpoint that makes HTTP requests to the chatter service
// login server.
func (c *Client) Login() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoginRequest(c.encoder)
		decodeResponse = DecodeLoginResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoginRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoginDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("chatter", "login", err)
		}
		return decodeResponse(resp)
	}
}

// Echoer returns an endpoint that makes HTTP requests to the chatter service
// echoer server.
func (c *Client) Echoer() goa.Endpoint {
	var (
		encodeRequest  = EncodeEchoerRequest(c.encoder)
		decodeResponse = DecodeEchoerResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildEchoerRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "echoer", err)
		}
		if c.configurer.EchoerFn != nil {
			conn = c.configurer.EchoerFn(conn, cancel)
		}
		stream := &EchoerClientStream{conn: conn}
		return stream, nil
	}
}

// Recv reads instances of "string" from the "echoer" endpoint websocket
// connection.
func (s *EchoerClientStream) Recv() (string, error) {
	var (
		rv   string
		body string
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}

// Send streams instances of "string" to the "echoer" endpoint websocket
// connection.
func (s *EchoerClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}

// Close closes the "echoer" endpoint websocket connection.
func (s *EchoerClientStream) Close() error {
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return s.conn.Close()
}

// Listener returns an endpoint that makes HTTP requests to the chatter service
// listener server.
func (c *Client) Listener() goa.Endpoint {
	var (
		encodeRequest  = EncodeListenerRequest(c.encoder)
		decodeResponse = DecodeListenerResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListenerRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "listener", err)
		}
		if c.configurer.ListenerFn != nil {
			conn = c.configurer.ListenerFn(conn, cancel)
		}
		stream := &ListenerClientStream{conn: conn}
		return stream, nil
	}
}

// Send streams instances of "string" to the "listener" endpoint websocket
// connection.
func (s *ListenerClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}

// Close closes the "listener" endpoint websocket connection.
func (s *ListenerClientStream) Close() error {
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return s.conn.Close()
}

// Summary returns an endpoint that makes HTTP requests to the chatter service
// summary server.
func (c *Client) Summary() goa.Endpoint {
	var (
		encodeRequest  = EncodeSummaryRequest(c.encoder)
		decodeResponse = DecodeSummaryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSummaryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "summary", err)
		}
		if c.configurer.SummaryFn != nil {
			conn = c.configurer.SummaryFn(conn, cancel)
		}
		stream := &SummaryClientStream{conn: conn}
		return stream, nil
	}
}

// CloseAndRecv stops sending messages to the "summary" endpoint websocket
// connection and reads instances of "chatter.ChatSummaryCollection" from the
// connection.
func (s *SummaryClientStream) CloseAndRecv() (chatter.ChatSummaryCollection, error) {
	var (
		rv   chatter.ChatSummaryCollection
		body SummaryResponseBody
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewSummaryChatSummaryCollectionOK(body)
	vres := chatterviews.ChatSummaryCollection{res, "default"}
	if err := chatterviews.ValidateChatSummaryCollection(vres); err != nil {
		return rv, goahttp.ErrValidationError("chatter", "summary", err)
	}
	return chatter.NewChatSummaryCollection(vres), nil
}

// Send streams instances of "string" to the "summary" endpoint websocket
// connection.
func (s *SummaryClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}

// Subscribe returns an endpoint that makes HTTP requests to the chatter
// service subscribe server.
func (c *Client) Subscribe() goa.Endpoint {
	var (
		encodeRequest  = EncodeSubscribeRequest(c.encoder)
		decodeResponse = DecodeSubscribeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSubscribeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "subscribe", err)
		}
		if c.configurer.SubscribeFn != nil {
			conn = c.configurer.SubscribeFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &SubscribeClientStream{conn: conn}
		return stream, nil
	}
}

// Recv reads instances of "chatter.Event" from the "subscribe" endpoint
// websocket connection.
func (s *SubscribeClientStream) Recv() (*chatter.Event, error) {
	var (
		rv   *chatter.Event
		body SubscribeResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	err = ValidateSubscribeResponseBody(&body)
	if err != nil {
		return rv, err
	}
	res := NewSubscribeEventOK(&body)
	return res, nil
}

// History returns an endpoint that makes HTTP requests to the chatter service
// history server.
func (c *Client) History() goa.Endpoint {
	var (
		encodeRequest  = EncodeHistoryRequest(c.encoder)
		decodeResponse = DecodeHistoryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildHistoryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		conn, resp, err := c.dialer.DialContext(ctx, req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "history", err)
		}
		if c.configurer.HistoryFn != nil {
			conn = c.configurer.HistoryFn(conn, cancel)
		}
		go func() {
			<-ctx.Done()
			conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client closing connection"),
				time.Now().Add(time.Second),
			)
			conn.Close()
		}()
		stream := &HistoryClientStream{conn: conn}
		view := resp.Header.Get("goa-view")
		stream.SetView(view)
		return stream, nil
	}
}

// Recv reads instances of "chatter.ChatSummary" from the "history" endpoint
// websocket connection.
func (s *HistoryClientStream) Recv() (*chatter.ChatSummary, error) {
	var (
		rv   *chatter.ChatSummary
		body HistoryResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewHistoryChatSummaryOK(&body)
	vres := &chatterviews.ChatSummary{res, s.view}
	if err := chatterviews.ValidateChatSummary(vres); err != nil {
		return rv, goahttp.ErrValidationError("chatter", "history", err)
	}
	return chatter.NewChatSummary(vres), nil
}

// SetView sets the view to render the  type before sending to the "history"
// endpoint websocket connection.
func (s *HistoryClientStream) SetView(view string) {
	s.view = view
}

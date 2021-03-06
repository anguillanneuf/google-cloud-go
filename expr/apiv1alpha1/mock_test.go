// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package expr

import (
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockConformanceServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	exprpb.ConformanceServiceServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockConformanceServer) Parse(ctx context.Context, req *exprpb.ParseRequest) (*exprpb.ParseResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*exprpb.ParseResponse), nil
}

func (s *mockConformanceServer) Check(ctx context.Context, req *exprpb.CheckRequest) (*exprpb.CheckResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*exprpb.CheckResponse), nil
}

func (s *mockConformanceServer) Eval(ctx context.Context, req *exprpb.EvalRequest) (*exprpb.EvalResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*exprpb.EvalResponse), nil
}

type mockCelServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	exprpb.CelServiceServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockCelServer) Parse(ctx context.Context, req *exprpb.ParseRequest) (*exprpb.ParseResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*exprpb.ParseResponse), nil
}

func (s *mockCelServer) Check(ctx context.Context, req *exprpb.CheckRequest) (*exprpb.CheckResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*exprpb.CheckResponse), nil
}

func (s *mockCelServer) Eval(ctx context.Context, req *exprpb.EvalRequest) (*exprpb.EvalResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*exprpb.EvalResponse), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockConformance mockConformanceServer
	mockCel         mockCelServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	exprpb.RegisterConformanceServiceServer(serv, &mockConformance)
	exprpb.RegisterCelServiceServer(serv, &mockCel)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestConformanceServiceParse(t *testing.T) {
	var expectedResponse *exprpb.ParseResponse = &exprpb.ParseResponse{}

	mockConformance.err = nil
	mockConformance.reqs = nil

	mockConformance.resps = append(mockConformance.resps[:0], expectedResponse)

	var celSource string = "celSource912645552"
	var request = &exprpb.ParseRequest{
		CelSource: celSource,
	}

	c, err := NewConformanceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Parse(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockConformance.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestConformanceServiceParseError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockConformance.err = gstatus.Error(errCode, "test error")

	var celSource string = "celSource912645552"
	var request = &exprpb.ParseRequest{
		CelSource: celSource,
	}

	c, err := NewConformanceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Parse(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestConformanceServiceCheck(t *testing.T) {
	var expectedResponse *exprpb.CheckResponse = &exprpb.CheckResponse{}

	mockConformance.err = nil
	mockConformance.reqs = nil

	mockConformance.resps = append(mockConformance.resps[:0], expectedResponse)

	var parsedExpr *exprpb.ParsedExpr = &exprpb.ParsedExpr{}
	var request = &exprpb.CheckRequest{
		ParsedExpr: parsedExpr,
	}

	c, err := NewConformanceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Check(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockConformance.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestConformanceServiceCheckError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockConformance.err = gstatus.Error(errCode, "test error")

	var parsedExpr *exprpb.ParsedExpr = &exprpb.ParsedExpr{}
	var request = &exprpb.CheckRequest{
		ParsedExpr: parsedExpr,
	}

	c, err := NewConformanceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Check(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestConformanceServiceEval(t *testing.T) {
	var expectedResponse *exprpb.EvalResponse = &exprpb.EvalResponse{}

	mockConformance.err = nil
	mockConformance.reqs = nil

	mockConformance.resps = append(mockConformance.resps[:0], expectedResponse)

	var request *exprpb.EvalRequest = &exprpb.EvalRequest{}

	c, err := NewConformanceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Eval(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockConformance.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestConformanceServiceEvalError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockConformance.err = gstatus.Error(errCode, "test error")

	var request *exprpb.EvalRequest = &exprpb.EvalRequest{}

	c, err := NewConformanceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Eval(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestCelServiceParse(t *testing.T) {
	var expectedResponse *exprpb.ParseResponse = &exprpb.ParseResponse{}

	mockCel.err = nil
	mockCel.reqs = nil

	mockCel.resps = append(mockCel.resps[:0], expectedResponse)

	var celSource string = "celSource912645552"
	var request = &exprpb.ParseRequest{
		CelSource: celSource,
	}

	c, err := NewCelClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Parse(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCel.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestCelServiceParseError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCel.err = gstatus.Error(errCode, "test error")

	var celSource string = "celSource912645552"
	var request = &exprpb.ParseRequest{
		CelSource: celSource,
	}

	c, err := NewCelClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Parse(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestCelServiceCheck(t *testing.T) {
	var expectedResponse *exprpb.CheckResponse = &exprpb.CheckResponse{}

	mockCel.err = nil
	mockCel.reqs = nil

	mockCel.resps = append(mockCel.resps[:0], expectedResponse)

	var parsedExpr *exprpb.ParsedExpr = &exprpb.ParsedExpr{}
	var request = &exprpb.CheckRequest{
		ParsedExpr: parsedExpr,
	}

	c, err := NewCelClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Check(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCel.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestCelServiceCheckError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCel.err = gstatus.Error(errCode, "test error")

	var parsedExpr *exprpb.ParsedExpr = &exprpb.ParsedExpr{}
	var request = &exprpb.CheckRequest{
		ParsedExpr: parsedExpr,
	}

	c, err := NewCelClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Check(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestCelServiceEval(t *testing.T) {
	var expectedResponse *exprpb.EvalResponse = &exprpb.EvalResponse{}

	mockCel.err = nil
	mockCel.reqs = nil

	mockCel.resps = append(mockCel.resps[:0], expectedResponse)

	var request *exprpb.EvalRequest = &exprpb.EvalRequest{}

	c, err := NewCelClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Eval(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockCel.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestCelServiceEvalError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockCel.err = gstatus.Error(errCode, "test error")

	var request *exprpb.EvalRequest = &exprpb.EvalRequest{}

	c, err := NewCelClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.Eval(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}

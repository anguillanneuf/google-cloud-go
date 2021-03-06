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

package expr_test

import (
	"cloud.google.com/go/expr/apiv1alpha1"
	"golang.org/x/net/context"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

func ExampleNewCelClient() {
	ctx := context.Background()
	c, err := expr.NewCelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleCelClient_Parse() {
	ctx := context.Background()
	c, err := expr.NewCelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &exprpb.ParseRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Parse(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCelClient_Check() {
	ctx := context.Background()
	c, err := expr.NewCelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &exprpb.CheckRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Check(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCelClient_Eval() {
	ctx := context.Background()
	c, err := expr.NewCelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &exprpb.EvalRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Eval(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

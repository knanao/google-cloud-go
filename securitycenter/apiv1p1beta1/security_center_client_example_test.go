// Copyright 2021 Google LLC
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

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package securitycenter_test

import (
	"context"

	securitycenter "cloud.google.com/go/securitycenter/apiv1p1beta1"
	"google.golang.org/api/iterator"
	securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateSource() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.CreateSourceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateSource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateFinding() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.CreateFindingRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateFinding(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateNotificationConfig() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.CreateNotificationConfigRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateNotificationConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteNotificationConfig() {
	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.DeleteNotificationConfigRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteNotificationConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetIamPolicy() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetNotificationConfig() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.GetNotificationConfigRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetNotificationConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetOrganizationSettings() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.GetOrganizationSettingsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetOrganizationSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetSource() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.GetSourceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetSource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GroupAssets() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.GroupAssetsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.GroupAssets(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GroupFindings() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.GroupFindingsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.GroupFindings(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListAssets() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.ListAssetsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListAssets(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListFindings() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.ListFindingsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListFindings(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListNotificationConfigs() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.ListNotificationConfigsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListNotificationConfigs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListSources() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.ListSourcesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListSources(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_RunAssetDiscovery() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.RunAssetDiscoveryRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.RunAssetDiscovery(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_SetFindingState() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.SetFindingStateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetFindingState(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_SetIamPolicy() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_TestIamPermissions() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateFinding() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.UpdateFindingRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateFinding(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateNotificationConfig() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.UpdateNotificationConfigRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateNotificationConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateOrganizationSettings() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.UpdateOrganizationSettingsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateOrganizationSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateSource() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.UpdateSourceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateSource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateSecurityMarks() {
	// import securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1p1beta1"

	ctx := context.Background()
	c, err := securitycenter.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &securitycenterpb.UpdateSecurityMarksRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateSecurityMarks(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

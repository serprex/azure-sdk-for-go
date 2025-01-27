//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// PermissionsClient contains the methods for the Permissions group.
// Don't use this type directly, use NewPermissionsClient() instead.
type PermissionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPermissionsClient creates a new instance of PermissionsClient with the specified values.
func NewPermissionsClient(con *arm.Connection, subscriptionID string) *PermissionsClient {
	return &PermissionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// ListForResource - Gets all permissions the caller has for a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PermissionsClient) ListForResource(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *PermissionsListForResourceOptions) *PermissionsListForResourcePager {
	return &PermissionsListForResourcePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForResourceCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, options)
		},
		advancer: func(ctx context.Context, resp PermissionsListForResourceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PermissionGetResult.NextLink)
		},
	}
}

// listForResourceCreateRequest creates the ListForResource request.
func (client *PermissionsClient) listForResourceCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *PermissionsListForResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/permissions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", resourceProviderNamespace)
	if parentResourcePath == "" {
		return nil, errors.New("parameter parentResourcePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourcePath}", parentResourcePath)
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", resourceType)
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForResourceHandleResponse handles the ListForResource response.
func (client *PermissionsClient) listForResourceHandleResponse(resp *http.Response) (PermissionsListForResourceResponse, error) {
	result := PermissionsListForResourceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PermissionGetResult); err != nil {
		return PermissionsListForResourceResponse{}, err
	}
	return result, nil
}

// listForResourceHandleError handles the ListForResource error response.
func (client *PermissionsClient) listForResourceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListForResourceGroup - Gets all permissions the caller has for a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PermissionsClient) ListForResourceGroup(resourceGroupName string, options *PermissionsListForResourceGroupOptions) *PermissionsListForResourceGroupPager {
	return &PermissionsListForResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp PermissionsListForResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PermissionGetResult.NextLink)
		},
	}
}

// listForResourceGroupCreateRequest creates the ListForResourceGroup request.
func (client *PermissionsClient) listForResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PermissionsListForResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Authorization/permissions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForResourceGroupHandleResponse handles the ListForResourceGroup response.
func (client *PermissionsClient) listForResourceGroupHandleResponse(resp *http.Response) (PermissionsListForResourceGroupResponse, error) {
	result := PermissionsListForResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PermissionGetResult); err != nil {
		return PermissionsListForResourceGroupResponse{}, err
	}
	return result, nil
}

// listForResourceGroupHandleError handles the ListForResourceGroup error response.
func (client *PermissionsClient) listForResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

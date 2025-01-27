//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// LinkedServerClient contains the methods for the LinkedServer group.
// Don't use this type directly, use NewLinkedServerClient() instead.
type LinkedServerClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLinkedServerClient creates a new instance of LinkedServerClient with the specified values.
func NewLinkedServerClient(con *arm.Connection, subscriptionID string) *LinkedServerClient {
	return &LinkedServerClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreate - Adds a linked server to the Redis cache (requires Premium SKU).
// If the operation fails it returns the *ErrorResponse error type.
func (client *LinkedServerClient) BeginCreate(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters RedisLinkedServerCreateParameters, options *LinkedServerBeginCreateOptions) (LinkedServerCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, name, linkedServerName, parameters, options)
	if err != nil {
		return LinkedServerCreatePollerResponse{}, err
	}
	result := LinkedServerCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LinkedServerClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return LinkedServerCreatePollerResponse{}, err
	}
	result.Poller = &LinkedServerCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Adds a linked server to the Redis cache (requires Premium SKU).
// If the operation fails it returns the *ErrorResponse error type.
func (client *LinkedServerClient) create(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters RedisLinkedServerCreateParameters, options *LinkedServerBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, name, linkedServerName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *LinkedServerClient) createCreateRequest(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters RedisLinkedServerCreateParameters, options *LinkedServerBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if linkedServerName == "" {
		return nil, errors.New("parameter linkedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServerName}", url.PathEscape(linkedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleError handles the Create error response.
func (client *LinkedServerClient) createHandleError(resp *http.Response) error {
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

// Delete - Deletes the linked server from a redis cache (requires Premium SKU).
// If the operation fails it returns the *ErrorResponse error type.
func (client *LinkedServerClient) Delete(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerDeleteOptions) (LinkedServerDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, linkedServerName, options)
	if err != nil {
		return LinkedServerDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedServerDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return LinkedServerDeleteResponse{}, client.deleteHandleError(resp)
	}
	return LinkedServerDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LinkedServerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if linkedServerName == "" {
		return nil, errors.New("parameter linkedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServerName}", url.PathEscape(linkedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *LinkedServerClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the detailed information about a linked server of a redis cache (requires Premium SKU).
// If the operation fails it returns the *ErrorResponse error type.
func (client *LinkedServerClient) Get(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerGetOptions) (LinkedServerGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, linkedServerName, options)
	if err != nil {
		return LinkedServerGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedServerGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedServerGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LinkedServerClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if linkedServerName == "" {
		return nil, errors.New("parameter linkedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServerName}", url.PathEscape(linkedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkedServerClient) getHandleResponse(resp *http.Response) (LinkedServerGetResponse, error) {
	result := LinkedServerGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisLinkedServerWithProperties); err != nil {
		return LinkedServerGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LinkedServerClient) getHandleError(resp *http.Response) error {
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

// List - Gets the list of linked servers associated with this redis cache (requires Premium SKU).
// If the operation fails it returns the *ErrorResponse error type.
func (client *LinkedServerClient) List(resourceGroupName string, name string, options *LinkedServerListOptions) *LinkedServerListPager {
	return &LinkedServerListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, name, options)
		},
		advancer: func(ctx context.Context, resp LinkedServerListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RedisLinkedServerWithPropertiesList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LinkedServerClient) listCreateRequest(ctx context.Context, resourceGroupName string, name string, options *LinkedServerListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LinkedServerClient) listHandleResponse(resp *http.Response) (LinkedServerListResponse, error) {
	result := LinkedServerListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisLinkedServerWithPropertiesList); err != nil {
		return LinkedServerListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *LinkedServerClient) listHandleError(resp *http.Response) error {
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

//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// CustomDomainsClient contains the methods for the CustomDomains group.
// Don't use this type directly, use NewCustomDomainsClient() instead.
type CustomDomainsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCustomDomainsClient creates a new instance of CustomDomainsClient with the specified values.
func NewCustomDomainsClient(con *arm.Connection, subscriptionID string) *CustomDomainsClient {
	return &CustomDomainsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update custom domain of one lifecycle application.
// If the operation fails it returns the *CloudError error type.
func (client *CustomDomainsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, domainResource CustomDomainResource, options *CustomDomainsBeginCreateOrUpdateOptions) (CustomDomainsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, appName, domainName, domainResource, options)
	if err != nil {
		return CustomDomainsCreateOrUpdatePollerResponse{}, err
	}
	result := CustomDomainsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CustomDomainsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return CustomDomainsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &CustomDomainsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update custom domain of one lifecycle application.
// If the operation fails it returns the *CloudError error type.
func (client *CustomDomainsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, domainResource CustomDomainResource, options *CustomDomainsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, appName, domainName, domainResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CustomDomainsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, domainResource CustomDomainResource, options *CustomDomainsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, domainResource)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *CustomDomainsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Delete the custom domain of one lifecycle application.
// If the operation fails it returns the *CloudError error type.
func (client *CustomDomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, options *CustomDomainsBeginDeleteOptions) (CustomDomainsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, appName, domainName, options)
	if err != nil {
		return CustomDomainsDeletePollerResponse{}, err
	}
	result := CustomDomainsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CustomDomainsClient.Delete", "azure-async-operation", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return CustomDomainsDeletePollerResponse{}, err
	}
	result.Poller = &CustomDomainsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete the custom domain of one lifecycle application.
// If the operation fails it returns the *CloudError error type.
func (client *CustomDomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, options *CustomDomainsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, appName, domainName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomDomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, options *CustomDomainsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CustomDomainsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the custom domain of one lifecycle application.
// If the operation fails it returns the *CloudError error type.
func (client *CustomDomainsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, options *CustomDomainsGetOptions) (CustomDomainsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, appName, domainName, options)
	if err != nil {
		return CustomDomainsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomDomainsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomDomainsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomDomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, options *CustomDomainsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomDomainsClient) getHandleResponse(resp *http.Response) (CustomDomainsGetResponse, error) {
	result := CustomDomainsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomDomainResource); err != nil {
		return CustomDomainsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CustomDomainsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List the custom domains of one lifecycle application.
// If the operation fails it returns the *CloudError error type.
func (client *CustomDomainsClient) List(resourceGroupName string, serviceName string, appName string, options *CustomDomainsListOptions) *CustomDomainsListPager {
	return &CustomDomainsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, serviceName, appName, options)
		},
		advancer: func(ctx context.Context, resp CustomDomainsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomDomainResourceCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *CustomDomainsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, options *CustomDomainsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/domains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CustomDomainsClient) listHandleResponse(resp *http.Response) (CustomDomainsListResponse, error) {
	result := CustomDomainsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomDomainResourceCollection); err != nil {
		return CustomDomainsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *CustomDomainsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Update custom domain of one lifecycle application.
// If the operation fails it returns the *CloudError error type.
func (client *CustomDomainsClient) BeginUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, domainResource CustomDomainResource, options *CustomDomainsBeginUpdateOptions) (CustomDomainsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, serviceName, appName, domainName, domainResource, options)
	if err != nil {
		return CustomDomainsUpdatePollerResponse{}, err
	}
	result := CustomDomainsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CustomDomainsClient.Update", "azure-async-operation", resp, client.pl, client.updateHandleError)
	if err != nil {
		return CustomDomainsUpdatePollerResponse{}, err
	}
	result.Poller = &CustomDomainsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update custom domain of one lifecycle application.
// If the operation fails it returns the *CloudError error type.
func (client *CustomDomainsClient) update(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, domainResource CustomDomainResource, options *CustomDomainsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, appName, domainName, domainResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *CustomDomainsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, domainResource CustomDomainResource, options *CustomDomainsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, domainResource)
}

// updateHandleError handles the Update error response.
func (client *CustomDomainsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

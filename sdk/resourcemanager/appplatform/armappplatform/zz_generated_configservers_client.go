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

// ConfigServersClient contains the methods for the ConfigServers group.
// Don't use this type directly, use NewConfigServersClient() instead.
type ConfigServersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewConfigServersClient creates a new instance of ConfigServersClient with the specified values.
func NewConfigServersClient(con *arm.Connection, subscriptionID string) *ConfigServersClient {
	return &ConfigServersClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Get the config server and its properties.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigServersClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *ConfigServersGetOptions) (ConfigServersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ConfigServersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigServersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigServersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConfigServersClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ConfigServersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/default"
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
func (client *ConfigServersClient) getHandleResponse(resp *http.Response) (ConfigServersGetResponse, error) {
	result := ConfigServersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigServerResource); err != nil {
		return ConfigServersGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ConfigServersClient) getHandleError(resp *http.Response) error {
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

// BeginUpdatePatch - Update the config server.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigServersClient) BeginUpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersBeginUpdatePatchOptions) (ConfigServersUpdatePatchPollerResponse, error) {
	resp, err := client.updatePatch(ctx, resourceGroupName, serviceName, configServerResource, options)
	if err != nil {
		return ConfigServersUpdatePatchPollerResponse{}, err
	}
	result := ConfigServersUpdatePatchPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConfigServersClient.UpdatePatch", "azure-async-operation", resp, client.pl, client.updatePatchHandleError)
	if err != nil {
		return ConfigServersUpdatePatchPollerResponse{}, err
	}
	result.Poller = &ConfigServersUpdatePatchPoller{
		pt: pt,
	}
	return result, nil
}

// UpdatePatch - Update the config server.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigServersClient) updatePatch(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersBeginUpdatePatchOptions) (*http.Response, error) {
	req, err := client.updatePatchCreateRequest(ctx, resourceGroupName, serviceName, configServerResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updatePatchHandleError(resp)
	}
	return resp, nil
}

// updatePatchCreateRequest creates the UpdatePatch request.
func (client *ConfigServersClient) updatePatchCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersBeginUpdatePatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, configServerResource)
}

// updatePatchHandleError handles the UpdatePatch error response.
func (client *ConfigServersClient) updatePatchHandleError(resp *http.Response) error {
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

// BeginUpdatePut - Update the config server.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigServersClient) BeginUpdatePut(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersBeginUpdatePutOptions) (ConfigServersUpdatePutPollerResponse, error) {
	resp, err := client.updatePut(ctx, resourceGroupName, serviceName, configServerResource, options)
	if err != nil {
		return ConfigServersUpdatePutPollerResponse{}, err
	}
	result := ConfigServersUpdatePutPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConfigServersClient.UpdatePut", "azure-async-operation", resp, client.pl, client.updatePutHandleError)
	if err != nil {
		return ConfigServersUpdatePutPollerResponse{}, err
	}
	result.Poller = &ConfigServersUpdatePutPoller{
		pt: pt,
	}
	return result, nil
}

// UpdatePut - Update the config server.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigServersClient) updatePut(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersBeginUpdatePutOptions) (*http.Response, error) {
	req, err := client.updatePutCreateRequest(ctx, resourceGroupName, serviceName, configServerResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updatePutHandleError(resp)
	}
	return resp, nil
}

// updatePutCreateRequest creates the UpdatePut request.
func (client *ConfigServersClient) updatePutCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersBeginUpdatePutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, configServerResource)
}

// updatePutHandleError handles the UpdatePut error response.
func (client *ConfigServersClient) updatePutHandleError(resp *http.Response) error {
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

// BeginValidate - Check if the config server settings are valid.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigServersClient) BeginValidate(ctx context.Context, resourceGroupName string, serviceName string, configServerSettings ConfigServerSettings, options *ConfigServersBeginValidateOptions) (ConfigServersValidatePollerResponse, error) {
	resp, err := client.validate(ctx, resourceGroupName, serviceName, configServerSettings, options)
	if err != nil {
		return ConfigServersValidatePollerResponse{}, err
	}
	result := ConfigServersValidatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConfigServersClient.Validate", "location", resp, client.pl, client.validateHandleError)
	if err != nil {
		return ConfigServersValidatePollerResponse{}, err
	}
	result.Poller = &ConfigServersValidatePoller{
		pt: pt,
	}
	return result, nil
}

// Validate - Check if the config server settings are valid.
// If the operation fails it returns the *CloudError error type.
func (client *ConfigServersClient) validate(ctx context.Context, resourceGroupName string, serviceName string, configServerSettings ConfigServerSettings, options *ConfigServersBeginValidateOptions) (*http.Response, error) {
	req, err := client.validateCreateRequest(ctx, resourceGroupName, serviceName, configServerSettings, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.validateHandleError(resp)
	}
	return resp, nil
}

// validateCreateRequest creates the Validate request.
func (client *ConfigServersClient) validateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, configServerSettings ConfigServerSettings, options *ConfigServersBeginValidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/validate"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, configServerSettings)
}

// validateHandleError handles the Validate error response.
func (client *ConfigServersClient) validateHandleError(resp *http.Response) error {
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

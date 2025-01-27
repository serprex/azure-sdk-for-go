//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// ConnectionTypeClient contains the methods for the ConnectionType group.
// Don't use this type directly, use NewConnectionTypeClient() instead.
type ConnectionTypeClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewConnectionTypeClient creates a new instance of ConnectionTypeClient with the specified values.
func NewConnectionTypeClient(con *arm.Connection, subscriptionID string) *ConnectionTypeClient {
	return &ConnectionTypeClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create a connection type.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionTypeClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, parameters ConnectionTypeCreateOrUpdateParameters, options *ConnectionTypeCreateOrUpdateOptions) (ConnectionTypeCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, connectionTypeName, parameters, options)
	if err != nil {
		return ConnectionTypeCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionTypeCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return ConnectionTypeCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectionTypeClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, parameters ConnectionTypeCreateOrUpdateParameters, options *ConnectionTypeCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes/{connectionTypeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if connectionTypeName == "" {
		return nil, errors.New("parameter connectionTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionTypeName}", url.PathEscape(connectionTypeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectionTypeClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectionTypeCreateOrUpdateResponse, error) {
	result := ConnectionTypeCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionType); err != nil {
		return ConnectionTypeCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ConnectionTypeClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Delete the connection type.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionTypeClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, options *ConnectionTypeDeleteOptions) (ConnectionTypeDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, connectionTypeName, options)
	if err != nil {
		return ConnectionTypeDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionTypeDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ConnectionTypeDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ConnectionTypeDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectionTypeClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, options *ConnectionTypeDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes/{connectionTypeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if connectionTypeName == "" {
		return nil, errors.New("parameter connectionTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionTypeName}", url.PathEscape(connectionTypeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ConnectionTypeClient) deleteHandleError(resp *http.Response) error {
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

// Get - Retrieve the connection type identified by connection type name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionTypeClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, options *ConnectionTypeGetOptions) (ConnectionTypeGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, connectionTypeName, options)
	if err != nil {
		return ConnectionTypeGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionTypeGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectionTypeGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectionTypeClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, options *ConnectionTypeGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes/{connectionTypeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if connectionTypeName == "" {
		return nil, errors.New("parameter connectionTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionTypeName}", url.PathEscape(connectionTypeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectionTypeClient) getHandleResponse(resp *http.Response) (ConnectionTypeGetResponse, error) {
	result := ConnectionTypeGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionType); err != nil {
		return ConnectionTypeGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ConnectionTypeClient) getHandleError(resp *http.Response) error {
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

// ListByAutomationAccount - Retrieve a list of connection types.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionTypeClient) ListByAutomationAccount(resourceGroupName string, automationAccountName string, options *ConnectionTypeListByAutomationAccountOptions) *ConnectionTypeListByAutomationAccountPager {
	return &ConnectionTypeListByAutomationAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
		},
		advancer: func(ctx context.Context, resp ConnectionTypeListByAutomationAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ConnectionTypeListResult.NextLink)
		},
	}
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *ConnectionTypeClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *ConnectionTypeListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *ConnectionTypeClient) listByAutomationAccountHandleResponse(resp *http.Response) (ConnectionTypeListByAutomationAccountResponse, error) {
	result := ConnectionTypeListByAutomationAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionTypeListResult); err != nil {
		return ConnectionTypeListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// listByAutomationAccountHandleError handles the ListByAutomationAccount error response.
func (client *ConnectionTypeClient) listByAutomationAccountHandleError(resp *http.Response) error {
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

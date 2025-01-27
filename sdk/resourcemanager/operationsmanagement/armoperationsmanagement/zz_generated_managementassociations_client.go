//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationsmanagement

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

// ManagementAssociationsClient contains the methods for the ManagementAssociations group.
// Don't use this type directly, use NewManagementAssociationsClient() instead.
type ManagementAssociationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
	providerName   string
	resourceType   string
	resourceName   string
}

// NewManagementAssociationsClient creates a new instance of ManagementAssociationsClient with the specified values.
func NewManagementAssociationsClient(con *arm.Connection, subscriptionID string, providerName string, resourceType string, resourceName string) *ManagementAssociationsClient {
	return &ManagementAssociationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID, providerName: providerName, resourceType: resourceType, resourceName: resourceName}
}

// CreateOrUpdate - Creates or updates the ManagementAssociation.
// If the operation fails it returns the *CodeMessageError error type.
func (client *ManagementAssociationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managementAssociationName string, parameters ManagementAssociation, options *ManagementAssociationsCreateOrUpdateOptions) (ManagementAssociationsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managementAssociationName, parameters, options)
	if err != nil {
		return ManagementAssociationsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementAssociationsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementAssociationsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagementAssociationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managementAssociationName string, parameters ManagementAssociation, options *ManagementAssociationsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.providerName == "" {
		return nil, errors.New("parameter client.providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(client.providerName))
	if client.resourceType == "" {
		return nil, errors.New("parameter client.resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(client.resourceType))
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if managementAssociationName == "" {
		return nil, errors.New("parameter managementAssociationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementAssociationName}", url.PathEscape(managementAssociationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagementAssociationsClient) createOrUpdateHandleResponse(resp *http.Response) (ManagementAssociationsCreateOrUpdateResponse, error) {
	result := ManagementAssociationsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementAssociation); err != nil {
		return ManagementAssociationsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ManagementAssociationsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CodeMessageError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes the ManagementAssociation in the subscription.
// If the operation fails it returns the *CodeMessageError error type.
func (client *ManagementAssociationsClient) Delete(ctx context.Context, resourceGroupName string, managementAssociationName string, options *ManagementAssociationsDeleteOptions) (ManagementAssociationsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managementAssociationName, options)
	if err != nil {
		return ManagementAssociationsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementAssociationsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementAssociationsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ManagementAssociationsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagementAssociationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managementAssociationName string, options *ManagementAssociationsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.providerName == "" {
		return nil, errors.New("parameter client.providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(client.providerName))
	if client.resourceType == "" {
		return nil, errors.New("parameter client.resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(client.resourceType))
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if managementAssociationName == "" {
		return nil, errors.New("parameter managementAssociationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementAssociationName}", url.PathEscape(managementAssociationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ManagementAssociationsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CodeMessageError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves the user ManagementAssociation.
// If the operation fails it returns the *CodeMessageError error type.
func (client *ManagementAssociationsClient) Get(ctx context.Context, resourceGroupName string, managementAssociationName string, options *ManagementAssociationsGetOptions) (ManagementAssociationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managementAssociationName, options)
	if err != nil {
		return ManagementAssociationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementAssociationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementAssociationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagementAssociationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managementAssociationName string, options *ManagementAssociationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.providerName == "" {
		return nil, errors.New("parameter client.providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(client.providerName))
	if client.resourceType == "" {
		return nil, errors.New("parameter client.resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(client.resourceType))
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if managementAssociationName == "" {
		return nil, errors.New("parameter managementAssociationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementAssociationName}", url.PathEscape(managementAssociationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagementAssociationsClient) getHandleResponse(resp *http.Response) (ManagementAssociationsGetResponse, error) {
	result := ManagementAssociationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementAssociation); err != nil {
		return ManagementAssociationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagementAssociationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CodeMessageError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Retrieves the ManagementAssociations list.
// If the operation fails it returns the *CodeMessageError error type.
func (client *ManagementAssociationsClient) ListBySubscription(ctx context.Context, options *ManagementAssociationsListBySubscriptionOptions) (ManagementAssociationsListBySubscriptionResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return ManagementAssociationsListBySubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementAssociationsListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementAssociationsListBySubscriptionResponse{}, client.listBySubscriptionHandleError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ManagementAssociationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ManagementAssociationsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OperationsManagement/ManagementAssociations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ManagementAssociationsClient) listBySubscriptionHandleResponse(resp *http.Response) (ManagementAssociationsListBySubscriptionResponse, error) {
	result := ManagementAssociationsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementAssociationPropertiesList); err != nil {
		return ManagementAssociationsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ManagementAssociationsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CodeMessageError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

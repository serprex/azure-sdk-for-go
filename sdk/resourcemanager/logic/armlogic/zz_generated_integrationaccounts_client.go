//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// IntegrationAccountsClient contains the methods for the IntegrationAccounts group.
// Don't use this type directly, use NewIntegrationAccountsClient() instead.
type IntegrationAccountsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIntegrationAccountsClient creates a new instance of IntegrationAccountsClient with the specified values.
func NewIntegrationAccountsClient(con *arm.Connection, subscriptionID string) *IntegrationAccountsClient {
	return &IntegrationAccountsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates an integration account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount, options *IntegrationAccountsCreateOrUpdateOptions) (IntegrationAccountsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, integrationAccount, options)
	if err != nil {
		return IntegrationAccountsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IntegrationAccountsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount, options *IntegrationAccountsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, integrationAccount)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountsClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountsCreateOrUpdateResponse, error) {
	result := IntegrationAccountsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccount); err != nil {
		return IntegrationAccountsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *IntegrationAccountsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Deletes an integration account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountsDeleteOptions) (IntegrationAccountsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
	if err != nil {
		return IntegrationAccountsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IntegrationAccountsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return IntegrationAccountsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *IntegrationAccountsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets an integration account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountsGetOptions) (IntegrationAccountsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
	if err != nil {
		return IntegrationAccountsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationAccountsClient) getHandleResponse(resp *http.Response) (IntegrationAccountsGetResponse, error) {
	result := IntegrationAccountsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccount); err != nil {
		return IntegrationAccountsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IntegrationAccountsClient) getHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Gets a list of integration accounts by resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountsClient) ListByResourceGroup(resourceGroupName string, options *IntegrationAccountsListByResourceGroupOptions) *IntegrationAccountsListByResourceGroupPager {
	return &IntegrationAccountsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp IntegrationAccountsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IntegrationAccountListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IntegrationAccountsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IntegrationAccountsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IntegrationAccountsClient) listByResourceGroupHandleResponse(resp *http.Response) (IntegrationAccountsListByResourceGroupResponse, error) {
	result := IntegrationAccountsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountListResult); err != nil {
		return IntegrationAccountsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *IntegrationAccountsClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// ListBySubscription - Gets a list of integration accounts by subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountsClient) ListBySubscription(options *IntegrationAccountsListBySubscriptionOptions) *IntegrationAccountsListBySubscriptionPager {
	return &IntegrationAccountsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp IntegrationAccountsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IntegrationAccountListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *IntegrationAccountsClient) listBySubscriptionCreateRequest(ctx context.Context, options *IntegrationAccountsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *IntegrationAccountsClient) listBySubscriptionHandleResponse(resp *http.Response) (IntegrationAccountsListBySubscriptionResponse, error) {
	result := IntegrationAccountsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountListResult); err != nil {
		return IntegrationAccountsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *IntegrationAccountsClient) listBySubscriptionHandleError(resp *http.Response) error {
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

// ListCallbackURL - Gets the integration account callback URL.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountsClient) ListCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, parameters GetCallbackURLParameters, options *IntegrationAccountsListCallbackURLOptions) (IntegrationAccountsListCallbackURLResponse, error) {
	req, err := client.listCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, parameters, options)
	if err != nil {
		return IntegrationAccountsListCallbackURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountsListCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountsListCallbackURLResponse{}, client.listCallbackURLHandleError(resp)
	}
	return client.listCallbackURLHandleResponse(resp)
}

// listCallbackURLCreateRequest creates the ListCallbackURL request.
func (client *IntegrationAccountsClient) listCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, parameters GetCallbackURLParameters, options *IntegrationAccountsListCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/listCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// listCallbackURLHandleResponse handles the ListCallbackURL response.
func (client *IntegrationAccountsClient) listCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountsListCallbackURLResponse, error) {
	result := IntegrationAccountsListCallbackURLResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CallbackURL); err != nil {
		return IntegrationAccountsListCallbackURLResponse{}, err
	}
	return result, nil
}

// listCallbackURLHandleError handles the ListCallbackURL error response.
func (client *IntegrationAccountsClient) listCallbackURLHandleError(resp *http.Response) error {
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

// ListKeyVaultKeys - Gets the integration account's Key Vault keys.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountsClient) ListKeyVaultKeys(ctx context.Context, resourceGroupName string, integrationAccountName string, listKeyVaultKeys ListKeyVaultKeysDefinition, options *IntegrationAccountsListKeyVaultKeysOptions) (IntegrationAccountsListKeyVaultKeysResponse, error) {
	req, err := client.listKeyVaultKeysCreateRequest(ctx, resourceGroupName, integrationAccountName, listKeyVaultKeys, options)
	if err != nil {
		return IntegrationAccountsListKeyVaultKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountsListKeyVaultKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountsListKeyVaultKeysResponse{}, client.listKeyVaultKeysHandleError(resp)
	}
	return client.listKeyVaultKeysHandleResponse(resp)
}

// listKeyVaultKeysCreateRequest creates the ListKeyVaultKeys request.
func (client *IntegrationAccountsClient) listKeyVaultKeysCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, listKeyVaultKeys ListKeyVaultKeysDefinition, options *IntegrationAccountsListKeyVaultKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/listKeyVaultKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, listKeyVaultKeys)
}

// listKeyVaultKeysHandleResponse handles the ListKeyVaultKeys response.
func (client *IntegrationAccountsClient) listKeyVaultKeysHandleResponse(resp *http.Response) (IntegrationAccountsListKeyVaultKeysResponse, error) {
	result := IntegrationAccountsListKeyVaultKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.KeyVaultKeyCollection); err != nil {
		return IntegrationAccountsListKeyVaultKeysResponse{}, err
	}
	return result, nil
}

// listKeyVaultKeysHandleError handles the ListKeyVaultKeys error response.
func (client *IntegrationAccountsClient) listKeyVaultKeysHandleError(resp *http.Response) error {
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

// LogTrackingEvents - Logs the integration account's tracking events.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountsClient) LogTrackingEvents(ctx context.Context, resourceGroupName string, integrationAccountName string, logTrackingEvents TrackingEventsDefinition, options *IntegrationAccountsLogTrackingEventsOptions) (IntegrationAccountsLogTrackingEventsResponse, error) {
	req, err := client.logTrackingEventsCreateRequest(ctx, resourceGroupName, integrationAccountName, logTrackingEvents, options)
	if err != nil {
		return IntegrationAccountsLogTrackingEventsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountsLogTrackingEventsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountsLogTrackingEventsResponse{}, client.logTrackingEventsHandleError(resp)
	}
	return IntegrationAccountsLogTrackingEventsResponse{RawResponse: resp}, nil
}

// logTrackingEventsCreateRequest creates the LogTrackingEvents request.
func (client *IntegrationAccountsClient) logTrackingEventsCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, logTrackingEvents TrackingEventsDefinition, options *IntegrationAccountsLogTrackingEventsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/logTrackingEvents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, logTrackingEvents)
}

// logTrackingEventsHandleError handles the LogTrackingEvents error response.
func (client *IntegrationAccountsClient) logTrackingEventsHandleError(resp *http.Response) error {
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

// RegenerateAccessKey - Regenerates the integration account access key.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountsClient) RegenerateAccessKey(ctx context.Context, resourceGroupName string, integrationAccountName string, regenerateAccessKey RegenerateActionParameter, options *IntegrationAccountsRegenerateAccessKeyOptions) (IntegrationAccountsRegenerateAccessKeyResponse, error) {
	req, err := client.regenerateAccessKeyCreateRequest(ctx, resourceGroupName, integrationAccountName, regenerateAccessKey, options)
	if err != nil {
		return IntegrationAccountsRegenerateAccessKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountsRegenerateAccessKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountsRegenerateAccessKeyResponse{}, client.regenerateAccessKeyHandleError(resp)
	}
	return client.regenerateAccessKeyHandleResponse(resp)
}

// regenerateAccessKeyCreateRequest creates the RegenerateAccessKey request.
func (client *IntegrationAccountsClient) regenerateAccessKeyCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, regenerateAccessKey RegenerateActionParameter, options *IntegrationAccountsRegenerateAccessKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/regenerateAccessKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, regenerateAccessKey)
}

// regenerateAccessKeyHandleResponse handles the RegenerateAccessKey response.
func (client *IntegrationAccountsClient) regenerateAccessKeyHandleResponse(resp *http.Response) (IntegrationAccountsRegenerateAccessKeyResponse, error) {
	result := IntegrationAccountsRegenerateAccessKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccount); err != nil {
		return IntegrationAccountsRegenerateAccessKeyResponse{}, err
	}
	return result, nil
}

// regenerateAccessKeyHandleError handles the RegenerateAccessKey error response.
func (client *IntegrationAccountsClient) regenerateAccessKeyHandleError(resp *http.Response) error {
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

// Update - Updates an integration account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountsClient) Update(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount, options *IntegrationAccountsUpdateOptions) (IntegrationAccountsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, integrationAccountName, integrationAccount, options)
	if err != nil {
		return IntegrationAccountsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *IntegrationAccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount, options *IntegrationAccountsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, integrationAccount)
}

// updateHandleResponse handles the Update response.
func (client *IntegrationAccountsClient) updateHandleResponse(resp *http.Response) (IntegrationAccountsUpdateResponse, error) {
	result := IntegrationAccountsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccount); err != nil {
		return IntegrationAccountsUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *IntegrationAccountsClient) updateHandleError(resp *http.Response) error {
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

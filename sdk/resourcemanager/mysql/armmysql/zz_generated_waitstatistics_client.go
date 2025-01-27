//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

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

// WaitStatisticsClient contains the methods for the WaitStatistics group.
// Don't use this type directly, use NewWaitStatisticsClient() instead.
type WaitStatisticsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWaitStatisticsClient creates a new instance of WaitStatisticsClient with the specified values.
func NewWaitStatisticsClient(con *arm.Connection, subscriptionID string) *WaitStatisticsClient {
	return &WaitStatisticsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Retrieve wait statistics for specified identifier.
// If the operation fails it returns the *CloudError error type.
func (client *WaitStatisticsClient) Get(ctx context.Context, resourceGroupName string, serverName string, waitStatisticsID string, options *WaitStatisticsGetOptions) (WaitStatisticsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, waitStatisticsID, options)
	if err != nil {
		return WaitStatisticsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WaitStatisticsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WaitStatisticsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WaitStatisticsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, waitStatisticsID string, options *WaitStatisticsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/waitStatistics/{waitStatisticsId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if waitStatisticsID == "" {
		return nil, errors.New("parameter waitStatisticsID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{waitStatisticsId}", url.PathEscape(waitStatisticsID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WaitStatisticsClient) getHandleResponse(resp *http.Response) (WaitStatisticsGetResponse, error) {
	result := WaitStatisticsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WaitStatistic); err != nil {
		return WaitStatisticsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WaitStatisticsClient) getHandleError(resp *http.Response) error {
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

// ListByServer - Retrieve wait statistics for specified aggregation window.
// If the operation fails it returns the *CloudError error type.
func (client *WaitStatisticsClient) ListByServer(resourceGroupName string, serverName string, parameters WaitStatisticsInput, options *WaitStatisticsListByServerOptions) *WaitStatisticsListByServerPager {
	return &WaitStatisticsListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, parameters, options)
		},
		advancer: func(ctx context.Context, resp WaitStatisticsListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WaitStatisticsResultList.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *WaitStatisticsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, parameters WaitStatisticsInput, options *WaitStatisticsListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/waitStatistics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// listByServerHandleResponse handles the ListByServer response.
func (client *WaitStatisticsClient) listByServerHandleResponse(resp *http.Response) (WaitStatisticsListByServerResponse, error) {
	result := WaitStatisticsListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WaitStatisticsResultList); err != nil {
		return WaitStatisticsListByServerResponse{}, err
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *WaitStatisticsClient) listByServerHandleError(resp *http.Response) error {
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

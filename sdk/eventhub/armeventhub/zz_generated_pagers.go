// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"reflect"
)

type ClustersListByResourceGroupPager interface {
	azcore.Pager
	// PageResponse returns the current ClustersListByResourceGroupResponse.
	PageResponse() ClustersListByResourceGroupResponse
}

type clustersListByResourceGroupPager struct {
	client    *ClustersClient
	current   ClustersListByResourceGroupResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, ClustersListByResourceGroupResponse) (*azcore.Request, error)
}

func (p *clustersListByResourceGroupPager) Err() error {
	return p.err
}

func (p *clustersListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClusterListResult.NextLink == nil || len(*p.current.ClusterListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *clustersListByResourceGroupPager) PageResponse() ClustersListByResourceGroupResponse {
	return p.current
}

type ConsumerGroupsListByEventHubPager interface {
	azcore.Pager
	// PageResponse returns the current ConsumerGroupsListByEventHubResponse.
	PageResponse() ConsumerGroupsListByEventHubResponse
}

type consumerGroupsListByEventHubPager struct {
	client    *ConsumerGroupsClient
	current   ConsumerGroupsListByEventHubResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, ConsumerGroupsListByEventHubResponse) (*azcore.Request, error)
}

func (p *consumerGroupsListByEventHubPager) Err() error {
	return p.err
}

func (p *consumerGroupsListByEventHubPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ConsumerGroupListResult.NextLink == nil || len(*p.current.ConsumerGroupListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listByEventHubHandleError(resp)
		return false
	}
	result, err := p.client.listByEventHubHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *consumerGroupsListByEventHubPager) PageResponse() ConsumerGroupsListByEventHubResponse {
	return p.current
}

type DisasterRecoveryConfigsListAuthorizationRulesPager interface {
	azcore.Pager
	// PageResponse returns the current DisasterRecoveryConfigsListAuthorizationRulesResponse.
	PageResponse() DisasterRecoveryConfigsListAuthorizationRulesResponse
}

type disasterRecoveryConfigsListAuthorizationRulesPager struct {
	client    *DisasterRecoveryConfigsClient
	current   DisasterRecoveryConfigsListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, DisasterRecoveryConfigsListAuthorizationRulesResponse) (*azcore.Request, error)
}

func (p *disasterRecoveryConfigsListAuthorizationRulesPager) Err() error {
	return p.err
}

func (p *disasterRecoveryConfigsListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listAuthorizationRulesHandleError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *disasterRecoveryConfigsListAuthorizationRulesPager) PageResponse() DisasterRecoveryConfigsListAuthorizationRulesResponse {
	return p.current
}

type DisasterRecoveryConfigsListPager interface {
	azcore.Pager
	// PageResponse returns the current DisasterRecoveryConfigsListResponse.
	PageResponse() DisasterRecoveryConfigsListResponse
}

type disasterRecoveryConfigsListPager struct {
	client    *DisasterRecoveryConfigsClient
	current   DisasterRecoveryConfigsListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, DisasterRecoveryConfigsListResponse) (*azcore.Request, error)
}

func (p *disasterRecoveryConfigsListPager) Err() error {
	return p.err
}

func (p *disasterRecoveryConfigsListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ArmDisasterRecoveryListResult.NextLink == nil || len(*p.current.ArmDisasterRecoveryListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *disasterRecoveryConfigsListPager) PageResponse() DisasterRecoveryConfigsListResponse {
	return p.current
}

type EventHubsListAuthorizationRulesPager interface {
	azcore.Pager
	// PageResponse returns the current EventHubsListAuthorizationRulesResponse.
	PageResponse() EventHubsListAuthorizationRulesResponse
}

type eventHubsListAuthorizationRulesPager struct {
	client    *EventHubsClient
	current   EventHubsListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, EventHubsListAuthorizationRulesResponse) (*azcore.Request, error)
}

func (p *eventHubsListAuthorizationRulesPager) Err() error {
	return p.err
}

func (p *eventHubsListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listAuthorizationRulesHandleError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *eventHubsListAuthorizationRulesPager) PageResponse() EventHubsListAuthorizationRulesResponse {
	return p.current
}

type EventHubsListByNamespacePager interface {
	azcore.Pager
	// PageResponse returns the current EventHubsListByNamespaceResponse.
	PageResponse() EventHubsListByNamespaceResponse
}

type eventHubsListByNamespacePager struct {
	client    *EventHubsClient
	current   EventHubsListByNamespaceResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, EventHubsListByNamespaceResponse) (*azcore.Request, error)
}

func (p *eventHubsListByNamespacePager) Err() error {
	return p.err
}

func (p *eventHubsListByNamespacePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EventHubListResult.NextLink == nil || len(*p.current.EventHubListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listByNamespaceHandleError(resp)
		return false
	}
	result, err := p.client.listByNamespaceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *eventHubsListByNamespacePager) PageResponse() EventHubsListByNamespaceResponse {
	return p.current
}

type NamespacesListAuthorizationRulesPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListAuthorizationRulesResponse.
	PageResponse() NamespacesListAuthorizationRulesResponse
}

type namespacesListAuthorizationRulesPager struct {
	client    *NamespacesClient
	current   NamespacesListAuthorizationRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListAuthorizationRulesResponse) (*azcore.Request, error)
}

func (p *namespacesListAuthorizationRulesPager) Err() error {
	return p.err
}

func (p *namespacesListAuthorizationRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listAuthorizationRulesHandleError(resp)
		return false
	}
	result, err := p.client.listAuthorizationRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *namespacesListAuthorizationRulesPager) PageResponse() NamespacesListAuthorizationRulesResponse {
	return p.current
}

type NamespacesListByResourceGroupPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListByResourceGroupResponse.
	PageResponse() NamespacesListByResourceGroupResponse
}

type namespacesListByResourceGroupPager struct {
	client    *NamespacesClient
	current   NamespacesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListByResourceGroupResponse) (*azcore.Request, error)
}

func (p *namespacesListByResourceGroupPager) Err() error {
	return p.err
}

func (p *namespacesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EHNamespaceListResult.NextLink == nil || len(*p.current.EHNamespaceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *namespacesListByResourceGroupPager) PageResponse() NamespacesListByResourceGroupResponse {
	return p.current
}

type NamespacesListIPFilterRulesPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListIPFilterRulesResponse.
	PageResponse() NamespacesListIPFilterRulesResponse
}

type namespacesListIPFilterRulesPager struct {
	client    *NamespacesClient
	current   NamespacesListIPFilterRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListIPFilterRulesResponse) (*azcore.Request, error)
}

func (p *namespacesListIPFilterRulesPager) Err() error {
	return p.err
}

func (p *namespacesListIPFilterRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IPFilterRuleListResult.NextLink == nil || len(*p.current.IPFilterRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listIPFilterRulesHandleError(resp)
		return false
	}
	result, err := p.client.listIPFilterRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *namespacesListIPFilterRulesPager) PageResponse() NamespacesListIPFilterRulesResponse {
	return p.current
}

type NamespacesListPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListResponse.
	PageResponse() NamespacesListResponse
}

type namespacesListPager struct {
	client    *NamespacesClient
	current   NamespacesListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListResponse) (*azcore.Request, error)
}

func (p *namespacesListPager) Err() error {
	return p.err
}

func (p *namespacesListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EHNamespaceListResult.NextLink == nil || len(*p.current.EHNamespaceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *namespacesListPager) PageResponse() NamespacesListResponse {
	return p.current
}

type NamespacesListVirtualNetworkRulesPager interface {
	azcore.Pager
	// PageResponse returns the current NamespacesListVirtualNetworkRulesResponse.
	PageResponse() NamespacesListVirtualNetworkRulesResponse
}

type namespacesListVirtualNetworkRulesPager struct {
	client    *NamespacesClient
	current   NamespacesListVirtualNetworkRulesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, NamespacesListVirtualNetworkRulesResponse) (*azcore.Request, error)
}

func (p *namespacesListVirtualNetworkRulesPager) Err() error {
	return p.err
}

func (p *namespacesListVirtualNetworkRulesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworkRuleListResult.NextLink == nil || len(*p.current.VirtualNetworkRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listVirtualNetworkRulesHandleError(resp)
		return false
	}
	result, err := p.client.listVirtualNetworkRulesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *namespacesListVirtualNetworkRulesPager) PageResponse() NamespacesListVirtualNetworkRulesResponse {
	return p.current
}

type OperationsListPager interface {
	azcore.Pager
	// PageResponse returns the current OperationsListResponse.
	PageResponse() OperationsListResponse
}

type operationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*azcore.Request, error)
}

func (p *operationsListPager) Err() error {
	return p.err
}

func (p *operationsListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *operationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

type PrivateEndpointConnectionsListPager interface {
	azcore.Pager
	// PageResponse returns the current PrivateEndpointConnectionsListResponse.
	PageResponse() PrivateEndpointConnectionsListResponse
}

type privateEndpointConnectionsListPager struct {
	client    *PrivateEndpointConnectionsClient
	current   PrivateEndpointConnectionsListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionsListResponse) (*azcore.Request, error)
}

func (p *privateEndpointConnectionsListPager) Err() error {
	return p.err
}

func (p *privateEndpointConnectionsListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateEndpointConnectionListResult.NextLink == nil || len(*p.current.PrivateEndpointConnectionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *privateEndpointConnectionsListPager) PageResponse() PrivateEndpointConnectionsListResponse {
	return p.current
}

type RegionsListBySKUPager interface {
	azcore.Pager
	// PageResponse returns the current RegionsListBySKUResponse.
	PageResponse() RegionsListBySKUResponse
}

type regionsListBySKUPager struct {
	client    *RegionsClient
	current   RegionsListBySKUResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, RegionsListBySKUResponse) (*azcore.Request, error)
}

func (p *regionsListBySKUPager) Err() error {
	return p.err
}

func (p *regionsListBySKUPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MessagingRegionsListResult.NextLink == nil || len(*p.current.MessagingRegionsListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listBySKUHandleError(resp)
		return false
	}
	result, err := p.client.listBySKUHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *regionsListBySKUPager) PageResponse() RegionsListBySKUResponse {
	return p.current
}
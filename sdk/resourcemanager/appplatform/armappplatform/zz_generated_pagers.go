//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AppsListPager provides operations for iterating over paged responses.
type AppsListPager struct {
	client    *AppsClient
	current   AppsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AppsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AppsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AppsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AppResourceCollection.NextLink == nil || len(*p.current.AppResourceCollection.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
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

// PageResponse returns the current AppsListResponse page.
func (p *AppsListPager) PageResponse() AppsListResponse {
	return p.current
}

// BindingsListPager provides operations for iterating over paged responses.
type BindingsListPager struct {
	client    *BindingsClient
	current   BindingsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BindingsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BindingsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BindingsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.BindingResourceCollection.NextLink == nil || len(*p.current.BindingResourceCollection.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
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

// PageResponse returns the current BindingsListResponse page.
func (p *BindingsListPager) PageResponse() BindingsListResponse {
	return p.current
}

// CertificatesListPager provides operations for iterating over paged responses.
type CertificatesListPager struct {
	client    *CertificatesClient
	current   CertificatesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CertificatesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CertificatesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CertificatesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateResourceCollection.NextLink == nil || len(*p.current.CertificateResourceCollection.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
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

// PageResponse returns the current CertificatesListResponse page.
func (p *CertificatesListPager) PageResponse() CertificatesListResponse {
	return p.current
}

// CustomDomainsListPager provides operations for iterating over paged responses.
type CustomDomainsListPager struct {
	client    *CustomDomainsClient
	current   CustomDomainsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CustomDomainsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CustomDomainsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CustomDomainsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CustomDomainResourceCollection.NextLink == nil || len(*p.current.CustomDomainResourceCollection.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
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

// PageResponse returns the current CustomDomainsListResponse page.
func (p *CustomDomainsListPager) PageResponse() CustomDomainsListResponse {
	return p.current
}

// DeploymentsListForClusterPager provides operations for iterating over paged responses.
type DeploymentsListForClusterPager struct {
	client    *DeploymentsClient
	current   DeploymentsListForClusterResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DeploymentsListForClusterResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DeploymentsListForClusterPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DeploymentsListForClusterPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeploymentResourceCollection.NextLink == nil || len(*p.current.DeploymentResourceCollection.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listForClusterHandleError(resp)
		return false
	}
	result, err := p.client.listForClusterHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DeploymentsListForClusterResponse page.
func (p *DeploymentsListForClusterPager) PageResponse() DeploymentsListForClusterResponse {
	return p.current
}

// DeploymentsListPager provides operations for iterating over paged responses.
type DeploymentsListPager struct {
	client    *DeploymentsClient
	current   DeploymentsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DeploymentsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DeploymentsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DeploymentsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeploymentResourceCollection.NextLink == nil || len(*p.current.DeploymentResourceCollection.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
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

// PageResponse returns the current DeploymentsListResponse page.
func (p *DeploymentsListPager) PageResponse() DeploymentsListResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AvailableOperations.NextLink == nil || len(*p.current.AvailableOperations.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
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

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// SKUsListPager provides operations for iterating over paged responses.
type SKUsListPager struct {
	client    *SKUsClient
	current   SKUsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SKUsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SKUsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SKUsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ResourceSKUCollection.NextLink == nil || len(*p.current.ResourceSKUCollection.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
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

// PageResponse returns the current SKUsListResponse page.
func (p *SKUsListPager) PageResponse() SKUsListResponse {
	return p.current
}

// ServicesListBySubscriptionPager provides operations for iterating over paged responses.
type ServicesListBySubscriptionPager struct {
	client    *ServicesClient
	current   ServicesListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServicesListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServicesListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServicesListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServiceResourceList.NextLink == nil || len(*p.current.ServiceResourceList.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ServicesListBySubscriptionResponse page.
func (p *ServicesListBySubscriptionPager) PageResponse() ServicesListBySubscriptionResponse {
	return p.current
}

// ServicesListPager provides operations for iterating over paged responses.
type ServicesListPager struct {
	client    *ServicesClient
	current   ServicesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServicesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServicesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServicesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServiceResourceList.NextLink == nil || len(*p.current.ServiceResourceList.NextLink) == 0 {
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
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
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

// PageResponse returns the current ServicesListResponse page.
func (p *ServicesListPager) PageResponse() ServicesListResponse {
	return p.current
}

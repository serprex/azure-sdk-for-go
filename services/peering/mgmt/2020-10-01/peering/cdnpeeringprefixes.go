package peering

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CdnPeeringPrefixesClient is the peering Client
type CdnPeeringPrefixesClient struct {
	BaseClient
}

// NewCdnPeeringPrefixesClient creates an instance of the CdnPeeringPrefixesClient client.
func NewCdnPeeringPrefixesClient(subscriptionID string) CdnPeeringPrefixesClient {
	return NewCdnPeeringPrefixesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCdnPeeringPrefixesClientWithBaseURI creates an instance of the CdnPeeringPrefixesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewCdnPeeringPrefixesClientWithBaseURI(baseURI string, subscriptionID string) CdnPeeringPrefixesClient {
	return CdnPeeringPrefixesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all of the advertised prefixes for the specified peering location
// Parameters:
// peeringLocation - the peering location.
func (client CdnPeeringPrefixesClient) List(ctx context.Context, peeringLocation string) (result CdnPeeringPrefixListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CdnPeeringPrefixesClient.List")
		defer func() {
			sc := -1
			if result.cpplr.Response.Response != nil {
				sc = result.cpplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, peeringLocation)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.CdnPeeringPrefixesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.cpplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.CdnPeeringPrefixesClient", "List", resp, "Failure sending request")
		return
	}

	result.cpplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.CdnPeeringPrefixesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.cpplr.hasNextLink() && result.cpplr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client CdnPeeringPrefixesClient) ListPreparer(ctx context.Context, peeringLocation string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version":     APIVersion,
		"peeringLocation": autorest.Encode("query", peeringLocation),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/cdnPeeringPrefixes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CdnPeeringPrefixesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CdnPeeringPrefixesClient) ListResponder(resp *http.Response) (result CdnPeeringPrefixListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CdnPeeringPrefixesClient) listNextResults(ctx context.Context, lastResults CdnPeeringPrefixListResult) (result CdnPeeringPrefixListResult, err error) {
	req, err := lastResults.cdnPeeringPrefixListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "peering.CdnPeeringPrefixesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "peering.CdnPeeringPrefixesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.CdnPeeringPrefixesClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CdnPeeringPrefixesClient) ListComplete(ctx context.Context, peeringLocation string) (result CdnPeeringPrefixListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CdnPeeringPrefixesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, peeringLocation)
	return
}
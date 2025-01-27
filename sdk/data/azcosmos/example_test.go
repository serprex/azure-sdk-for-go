// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

// This example shows you how to get started using the Azure Cosmos DB SDK for Go. NewCosmosClient creates a new instance of Cosmos client with the specified values. It uses the default pipeline configuration.
func Example() {

	endpoint, _ := os.LookupEnv("SOME_ENDPOINT")
	key, _ := os.LookupEnv("SOME_KEY")

	// Create new Cosmos DB client.
	cred, _ := azcosmos.NewKeyCredential(key)
	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	// All operations for Go operate on a context.Context, allowing you to control cancellation/timeout.
	ctx := context.Background()

	// This example showcases several common operations to help you get started, such as:

	// ===== 1. Creating a database =====

	databaseProperties := azcosmos.DatabaseProperties{ID: "databaseName"}
	databaseResponse, err := client.CreateDatabase(ctx, databaseProperties, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Database created. ActivityId %s", databaseResponse.ActivityId)

	// ===== 2. Creating a container =====

	properties := azcosmos.ContainerProperties{
		ID: "aContainer",
		PartitionKeyDefinition: azcosmos.PartitionKeyDefinition{
			Paths: []string{"/myPartitionKey"},
		},
	}

	throughput := azcosmos.NewManualThroughputProperties(400)

	database, err := client.NewDatabase("databaseName")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := database.CreateContainer(ctx, properties, &azcosmos.CreateContainerOptions{ThroughputProperties: &throughput})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Container created. ActivityId %s", resp.ActivityId)

	container, err := client.NewContainer("databaseName", "aContainer")
	if err != nil {
		log.Fatal(err)
	}

	// ===== 3. Update container properties =====

	resp, err = container.Read(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	updatedProperties := azcosmos.ContainerProperties{
		ID: "aContainer",
		PartitionKeyDefinition: azcosmos.PartitionKeyDefinition{
			Paths: []string{"/myPartitionKey"},
		},
		IndexingPolicy: &azcosmos.IndexingPolicy{
			IncludedPaths: []azcosmos.IncludedPath{},
			ExcludedPaths: []azcosmos.ExcludedPath{},
			Automatic:     false,
			IndexingMode:  azcosmos.IndexingModeNone,
		},
	}

	// Replace container properties
	resp, err = container.Replace(ctx, updatedProperties, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Read the manual throughput property
	throughputResponse, err := container.ReadThroughput(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	manualThroughput, hasManual := throughputResponse.ThroughputProperties.ManualThroughput()
	if !hasManual {
		log.Fatal("Expected to have manual throughput")
	}
	fmt.Printf("Container is provisioned with %v RU/s", manualThroughput)

	// Replace manual throughput property

	newScale := azcosmos.NewManualThroughputProperties(500)
	_, err = container.ReplaceThroughput(ctx, newScale, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate from manual throughput to autoscale
	newScale = azcosmos.NewAutoscaleThroughputProperties(10000)
	replaceThroughputResponse, err := container.ReplaceThroughput(ctx, newScale, nil)
	if err != nil {
		log.Fatal(err)
	}

	autoscaleMaxThroughputResponse, hasAutoscale := replaceThroughputResponse.ThroughputProperties.AutoscaleMaxThroughput()
	if !hasAutoscale {
		log.Fatal("Expected to have autoscale throughput")
	}

	fmt.Printf("Container is provisioned with %v RU/s", autoscaleMaxThroughputResponse)

	// ===== 4. Item CRUD =====

	// Items in an Azure Cosmos container are uniquely identified by their id and partition key value.
	pk := azcosmos.NewPartitionKeyString("newPartitionKey")

	item := map[string]string{
		"id":             "1",
		"value":          "2",
		"myPartitionKey": "newPartitionKey",
	}

	marshalled, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	// Create item.
	itemResponse, err := container.CreateItem(ctx, pk, marshalled, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Read item.
	itemResponse, err = container.ReadItem(ctx, pk, "1", nil)
	if err != nil {
		log.Fatal(err)
	}

	var itemResponseBody map[string]string
	err = json.Unmarshal(itemResponse.Value, &itemResponseBody)
	if err != nil {
		log.Fatal(err)
	}

	// Modify some property
	itemResponseBody["value"] = "newValue"
	marshalledReplace, err := json.Marshal(itemResponseBody)
	if err != nil {
		log.Fatal(err)
	}

	// Replace item
	itemResponse, err = container.ReplaceItem(ctx, pk, "1", marshalledReplace, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Delete item.
	itemResponse, err = container.DeleteItem(ctx, pk, "1", nil)
	if err != nil {
		log.Fatal(err)
	}

	// ===== 5. Session consistency =====

	itemResponse, err = container.UpsertItem(ctx, pk, marshalled, nil)
	if err != nil {
		log.Fatal(err)
	}

	itemSessionToken := itemResponse.SessionToken
	itemResponse, err = container.ReadItem(ctx, pk, "1", &azcosmos.ItemOptions{SessionToken: itemSessionToken})
	if err != nil {
		log.Fatal(err)
	}

	// ===== 6. Optimistic Concurrency Etag PreConditionFail =====

	// Azure Cosmos DB supports optimistic concurrency control to prevent lost updates or deletes and detection of conflicting operations.
	// Check the item response status code. If an error is imitted and the response code is 412 then retry operation.
	numberRetry := 3
	err = retryOptimisticConcurrency(numberRetry, 1000*time.Millisecond, func() (bool, error) {
		itemResponse, err = container.ReadItem(ctx, pk, "1", nil)
		if err != nil {
			log.Fatal(err)
		}

		var itemResponseBody map[string]string
		err = json.Unmarshal(itemResponse.Value, &itemResponseBody)
		if err != nil {
			log.Fatal(err)
		}

		// Change a value in the item response body.
		itemResponseBody["value"] = "newValue"

		marshalledReplace, err := json.Marshal(itemResponseBody)
		if err != nil {
			log.Fatal(err)
		}

		// Replace with Etag
		etag := itemResponse.ETag
		itemResponse, err = container.ReplaceItem(ctx, pk, "1", marshalledReplace, &azcosmos.ItemOptions{IfMatchEtag: &etag})
		var httpErr azcore.HTTPResponse

		return (errors.As(err, &httpErr) && itemResponse.RawResponse.StatusCode == 412), err
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func retryOptimisticConcurrency(retryAttempts int, wait time.Duration, retry func() (bool, error)) (result error) {
	for i := 0; ; i++ {
		retryResult, err := retry()
		if err != nil {
			log.Fatal(err)
			break
		}

		if !(retryResult) {
			break
		}

		if i >= (retryAttempts - 1) {
			break
		}

		log.Fatal("retrying after error:", err)

		time.Sleep(wait)
	}
	return fmt.Errorf("Cosmos DB retry attempts %d, error: %s", retryAttempts, result)
}

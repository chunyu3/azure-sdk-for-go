package armagrifood

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"golang.org/x/net/http2"
)
import "github.com/Azure/azure-sdk-for-go/sdk/to"

const (
	mockHost = "https://localhost:8443"
)

var (
	ctx            context.Context
	subscriptionId string
	cred           *azidentity.EnvironmentCredential
	err            error
	con            *armcore.Connection
)

func TestExtensions_Create(t *testing.T) {

	// From example Extensions_Create
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewExtensionsClient(con,
		"11111111-2222-3333-4444-555555555555")
	_, err := client.Create(ctx,
		"provider.extension",
		"examples-farmbeatsResourceName",
		"examples-rg",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestExtensions_Get(t *testing.T) {

	// From example Extensions_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewExtensionsClient(con,
		"11111111-2222-3333-4444-555555555555")
	_, err := client.Get(ctx,
		"provider.extension",
		"examples-farmbeatsResourceName",
		"examples-rg",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestExtensions_Update(t *testing.T) {

	// From example Extensions_Update
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewExtensionsClient(con,
		"11111111-2222-3333-4444-555555555555")
	_, err := client.Update(ctx,
		"provider.extension",
		"examples-farmbeatsResourceName",
		"examples-rg",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestExtensions_Delete(t *testing.T) {

	// From example Extensions_Delete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewExtensionsClient(con,
		"11111111-2222-3333-4444-555555555555")
	_, err := client.Delete(ctx,
		"provider.extension",
		"examples-farmbeatsResourceName",
		"examples-rg",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestExtensions_ListByFarmBeats(t *testing.T) {

	// From example Extensions_ListByFarmBeats
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewExtensionsClient(con,
		"11111111-2222-3333-4444-555555555555")
	client.ListByFarmBeats("examples-rg",
		"examples-farmbeatsResourceName",
		nil)

}

func TestFarmBeatsExtensions_List(t *testing.T) {

	// From example FarmBeatsExtensions_List
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewFarmBeatsExtensionsClient(con)
	client.List(nil)

}

func TestFarmBeatsExtensions_Get(t *testing.T) {

	// From example FarmBeatsExtensions_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewFarmBeatsExtensionsClient(con)
	_, err := client.Get(ctx,
		"DTN.ContentServices",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestFarmBeatsModels_Get(t *testing.T) {

	// From example FarmBeatsModels_Get
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewFarmBeatsModelsClient(con,
		"11111111-2222-3333-4444-555555555555")
	_, err := client.Get(ctx,
		"examples-rg",
		"examples-farmBeatsResourceName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestFarmBeatsModels_CreateOrUpdate(t *testing.T) {

	// From example FarmBeatsModels_CreateOrUpdate
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewFarmBeatsModelsClient(con,
		"11111111-2222-3333-4444-555555555555")
	_, err := client.CreateOrUpdate(ctx,
		"examples-farmbeatsResourceName",
		"examples-rg",
		FarmBeats{
			TrackedResource: TrackedResource{
				Location: to.StringPtr("eastus2"),
				Tags: map[string]*string{
					"key1": to.StringPtr("value1"),
					"key2": to.StringPtr("value2"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestFarmBeatsModels_Update(t *testing.T) {

	// From example FarmBeatsModels_Update
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewFarmBeatsModelsClient(con,
		"11111111-2222-3333-4444-555555555555")
	_, err := client.Update(ctx,
		"examples-farmBeatsResourceName",
		"examples-rg",
		FarmBeatsUpdateRequestModel{
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
				"key2": to.StringPtr("value2"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestFarmBeatsModels_Delete(t *testing.T) {

	// From example FarmBeatsModels_Delete
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewFarmBeatsModelsClient(con,
		"11111111-2222-3333-4444-555555555555")
	_, err := client.Delete(ctx,
		"examples-rg",
		"examples-farmBeatsResourceName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestFarmBeatsModels_ListBySubscription(t *testing.T) {

	// From example FarmBeatsModels_ListBySubscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewFarmBeatsModelsClient(con,
		"11111111-2222-3333-4444-555555555555")
	client.ListBySubscription(nil)

}

func TestFarmBeatsModels_ListByResourceGroup(t *testing.T) {

	// From example FarmBeatsModels_ListByResourceGroup
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewFarmBeatsModelsClient(con,
		"11111111-2222-3333-4444-555555555555")
	client.ListByResourceGroup("examples-rg",
		nil)

}

func TestLocations_CheckNameAvailability(t *testing.T) {

	// From example Locations_CheckNameAvailability_AlreadyExists
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewLocationsClient(con,
		"11111111-2222-3333-4444-555555555555")
	_, err := client.CheckNameAvailability(ctx,
		CheckNameAvailabilityRequest{
			Name: to.StringPtr("existingaccountname"),
			Type: to.StringPtr("Microsoft.AgFoodPlatform/farmBeats"),
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Locations_CheckNameAvailability_Available
	_, err = client.CheckNameAvailability(ctx,
		CheckNameAvailabilityRequest{
			Name: to.StringPtr("newaccountname"),
			Type: to.StringPtr("Microsoft.AgFoodPlatform/farmBeats"),
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestOperations_List(t *testing.T) {

	// From example Operations_List
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewOperationsClient(con)
	client.List(nil)

}

// TestMain will exec each test
func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run() // exec test and this returns an exit code to pass to os
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	ctx = context.Background()
	subscriptionId = os.Getenv("SUBSCRIPTION_ID")

	tr := &http.Transport{}
	if err := http2.ConfigureTransport(tr); err != nil {
		fmt.Printf("Failed to configure http2 transport: %v", err)
	}
	tr.TLSClientConfig.InsecureSkipVerify = true
	client := &http.Client{Transport: tr}
	cred, err = azidentity.NewEnvironmentCredential(&azidentity.EnvironmentCredentialOptions{AuthorityHost: mockHost, HTTPClient: client})
	if err != nil {
		panic(err)
	}

	con = armcore.NewConnection(mockHost, cred, &armcore.ConnectionOptions{
		Logging: azcore.LogOptions{
			IncludeBody: true,
		},
		HTTPClient: client,
	})
}

func tearDown() {

}

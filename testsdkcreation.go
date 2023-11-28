// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package testsdkcreation

import (
	"context"
	"fmt"
	"net/http"
	"testsdkcreation/pkg/models/shared"
	"testsdkcreation/pkg/utils"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Smartcar API
	"https://api.smartcar.com/v2.0",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Testsdkcreation - Smartcar API: # How do I use Postman with Smartcar?
// We've detailed how to get started with Smartcar in Postman [here](https://www.notion.so/smartcarapi/How-do-I-use-Postman-with-Smartcar-b8e8483bae8b43a986715582beb54bd4).
type Testsdkcreation struct {
	// Operations about compatibility
	Compatibility *Compatibility
	User          *User
	// Operations about vehicles
	Vehicles *Vehicles
	Tesla    *Tesla
	// Operations about electric vehicles
	Evs       *Evs
	Cadillac  *Cadillac
	Chevrolet *Chevrolet
	Webhooks  *Webhooks

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Testsdkcreation)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Testsdkcreation) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Testsdkcreation) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Testsdkcreation) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Testsdkcreation) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details

func WithSecurity(bearerAuth string) SDKOption {
	return func(sdk *Testsdkcreation) {
		security := shared.Security{BearerAuth: bearerAuth}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Testsdkcreation) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Testsdkcreation {
	sdk := &Testsdkcreation{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "0.2.3",
			GenVersion:        "2.202.2",
			UserAgent:         "speakeasy-sdk/go 0.2.3 2.202.2 1.0.0 testsdkcreation",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Compatibility = newCompatibility(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	sdk.Vehicles = newVehicles(sdk.sdkConfiguration)

	sdk.Tesla = newTesla(sdk.sdkConfiguration)

	sdk.Evs = newEvs(sdk.sdkConfiguration)

	sdk.Cadillac = newCadillac(sdk.sdkConfiguration)

	sdk.Chevrolet = newChevrolet(sdk.sdkConfiguration)

	sdk.Webhooks = newWebhooks(sdk.sdkConfiguration)

	return sdk
}

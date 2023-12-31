// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package pushcash

import (
	"fmt"
	"net/http"
	"push-cash/pkg/models/shared"
	"push-cash/pkg/utils"
	"time"
)

const (
	// ServerProd - Production API
	ServerProd string = "prod"
	// ServerSandbox - Sandbox API used for developing an integration with Push
	ServerSandbox string = "sandbox"
)

// ServerList contains the list of servers available to the SDK
var ServerList = map[string]string{
	ServerProd:    "https://api.pushcash.co",
	ServerSandbox: "https://sandbox.pushcash.co",
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
	Security          *shared.Security
	ServerURL         string
	Server            string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	if c.Server == "" {
		c.Server = "prod"
	}

	return ServerList[c.Server], nil
}

// PushCash
type PushCash struct {
	Balance  *balance
	Events   *events
	Intent   *intent
	Transfer *transfer
	User     *user

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*PushCash)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *PushCash) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *PushCash) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServer allows the overriding of the default server by name
func WithServer(server string) SDKOption {
	return func(sdk *PushCash) {
		_, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		sdk.sdkConfiguration.Server = server
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *PushCash) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *PushCash) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *PushCash {
	sdk := &PushCash{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.0.1",
			SDKVersion:        "1.1.0",
			GenVersion:        "2.73.0",
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

	sdk.Balance = newBalance(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.Intent = newIntent(sdk.sdkConfiguration)

	sdk.Transfer = newTransfer(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	return sdk
}

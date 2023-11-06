<div align="center">
   <img src="https://github.com/speakeasy-sdks/push-cash/assets/68016351/40cf454f-d017-4aa4-ae41-a546e26a7d3a">
   <h1>Push Cash SDK</h1>
   <p>Instant Transfers, Gauranteed</p>
   <a href="https://www.pushcash.co/docs"><img src="https://img.shields.io/static/v1?label=Docs&message=Docs&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/push-cash/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/push-cash/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/push-cash
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"log"
	pushcash "push-cash"
	"push-cash/pkg/models/shared"
)

func main() {
	s := pushcash.New(
		pushcash.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Balance.GetBalance(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountBalance != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Balance](docs/sdks/balance/README.md)

* [GetBalance](docs/sdks/balance/README.md#getbalance) - Get Push Account balance
* [GetTransaction](docs/sdks/balance/README.md#gettransaction) - Get a transaction
* [List](docs/sdks/balance/README.md#list) - List transactions

### [Events](docs/sdks/events/README.md)

* [GetEvent](docs/sdks/events/README.md#getevent) - Retrieve an event
* [List](docs/sdks/events/README.md#list) - List events

### [Intent](docs/sdks/intent/README.md)

* [CancelIntent](docs/sdks/intent/README.md#cancelintent) - Cancel an intent
* [CreateIntent](docs/sdks/intent/README.md#createintent) - Create intent
* [GetIntent](docs/sdks/intent/README.md#getintent) - Get an intent
* [List](docs/sdks/intent/README.md#list) - List intents

### [Transfer](docs/sdks/transfer/README.md)

* [CreateTransfer](docs/sdks/transfer/README.md#createtransfer) - Create a transfer
* [GetTransfer](docs/sdks/transfer/README.md#gettransfer) - Retrieve a transfer
* [List](docs/sdks/transfer/README.md#list) - List transfers

### [User](docs/sdks/user/README.md)

* [CreateUser](docs/sdks/user/README.md#createuser) - Create user
* [GetUser](docs/sdks/user/README.md#getuser) - Get a user
* [List](docs/sdks/user/README.md#list) - List users
* [UpdateUser](docs/sdks/user/README.md#updateuser) - Update a user's status
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->



<!-- End Dev Containers -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Name

You can override the default server globally using the `WithServer` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name | Server | Variables |
| ----- | ------ | --------- |
| `prod` | `https://api.pushcash.co` | None |
| `sandbox` | `https://sandbox.pushcash.co` | None |

For example:


```go
package main

import (
	"context"
	"log"
	pushcash "push-cash"
	"push-cash/pkg/models/shared"
)

func main() {
	s := pushcash.New(
		pushcash.WithSecurity(""),
		pushcash.WithServer("sandbox"),
	)

	ctx := context.Background()
	res, err := s.Balance.GetBalance(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountBalance != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:


```go
package main

import (
	"context"
	"log"
	pushcash "push-cash"
	"push-cash/pkg/models/shared"
)

func main() {
	s := pushcash.New(
		pushcash.WithSecurity(""),
		pushcash.WithServerURL("https://api.pushcash.co"),
	)

	ctx := context.Background()
	res, err := s.Balance.GetBalance(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountBalance != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Go Types -->
# Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

## Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

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

import(
	"context"
	"log"
	pushcash "push-cash"
	"push-cash/pkg/models/shared"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity(shared.Security{
            Bearer: "",
        }),
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



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



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

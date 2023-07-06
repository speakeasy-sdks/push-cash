<div align="center">
   <img src="https://github.com/speakeasy-sdks/push-cash/assets/68016351/40cf454f-d017-4aa4-ae41-a546e26a7d3a">
   <h1>Push Cash SDK</h1>
   <p><strong>Instant Transfers, Gauranteed></p>
   <a href="https://www.pushcash.co/docs"><img src="https://img.shields.io/static/v1?label=Docs&message=Docs&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/push-cash/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/push-cash/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation
![Uploading 64812dddd94ab3eeda765cc5_logo.svg…]()

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
	"push-cash"
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
* [ListTransactions](docs/sdks/balance/README.md#listtransactions) - List transactions

### [Events](docs/sdks/events/README.md)

* [GetEvent](docs/sdks/events/README.md#getevent) - Retrieve an event
* [ListEvents](docs/sdks/events/README.md#listevents) - List events

### [Intent](docs/sdks/intent/README.md)

* [CancelIntent](docs/sdks/intent/README.md#cancelintent) - Cancel an intent
* [CreateIntent](docs/sdks/intent/README.md#createintent) - Create intent
* [GetIntent](docs/sdks/intent/README.md#getintent) - Get an intent
* [ListIntents](docs/sdks/intent/README.md#listintents) - List intents

### [Transfer](docs/sdks/transfer/README.md)

* [CreateTransfer](docs/sdks/transfer/README.md#createtransfer) - Create a transfer
* [GetTransfer](docs/sdks/transfer/README.md#gettransfer) - Retrieve a transfer
* [ListTransfers](docs/sdks/transfer/README.md#listtransfers) - List transfers

### [User](docs/sdks/user/README.md)

* [CreateUser](docs/sdks/user/README.md#createuser) - Create user
* [GetUser](docs/sdks/user/README.md#getuser) - Get a user
* [ListUsers](docs/sdks/user/README.md#listusers) - List users
* [UpdateUser](docs/sdks/user/README.md#updateuser) - Update a user's status
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

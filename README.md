# push-cash

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

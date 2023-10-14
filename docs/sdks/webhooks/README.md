# Webhooks
(*Webhooks*)

### Available Operations

* [Subscribe](#subscribe) - Subscribe Webhook
* [Unsubscribe](#unsubscribe) - Unsubscribe Webhook

## Subscribe

__Description__

Subscribe to a webhook for a vehicle.

__Permission__

`required: webhook:read`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  status|   string|  If the request is successful, Smartcar will return “success” (HTTP 200 status).|

### Example Usage

```go
package main

import(
	"context"
	"log"
	"testsdkcreation"
	"testsdkcreation/pkg/models/shared"
)

func main() {
    s := testsdkcreation.New(
        testsdkcreation.WithSecurity(""),
    )


    var vehicleID string = "Chicken"

    var webhookID string = "Cedi"

    webhookInfo := &shared.WebhookInfo{
        Vehicleid: testsdkcreation.String("dc6ea99e-57d1-4e41-b129-27e7eb58713e"),
        Webhookid: testsdkcreation.String("9b6ae692-60cc-4b3e-89d8-71e7549cf805"),
    }

    ctx := context.Background()
    res, err := s.Webhooks.Subscribe(ctx, vehicleID, webhookID, webhookInfo)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `vehicleID`                                               | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `webhookID`                                               | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `webhookInfo`                                             | [*shared.WebhookInfo](../../models/shared/webhookinfo.md) | :heavy_minus_sign:                                        | N/A                                                       |


### Response

**[*operations.SubscribeResponse](../../models/operations/subscriberesponse.md), error**


## Unsubscribe

__Description__

Delete a webhook for a vehicle.

__Permission__

`required: webhook:read`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  status|   string|  If the request is successful, Smartcar will return “success” (HTTP 200 status).|

### Example Usage

```go
package main

import(
	"context"
	"log"
	"testsdkcreation"
	"testsdkcreation/pkg/models/shared"
)

func main() {
    s := testsdkcreation.New(
        testsdkcreation.WithSecurity(""),
    )


    var vehicleID string = "deposit"

    var webhookID string = "royal"

    ctx := context.Background()
    res, err := s.Webhooks.Unsubscribe(ctx, vehicleID, webhookID)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `vehicleID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `webhookID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.UnsubscribeResponse](../../models/operations/unsubscriberesponse.md), error**


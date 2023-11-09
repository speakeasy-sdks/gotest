# Cadillac
(*Cadillac*)

### Available Operations

* [GetChargeTime](#getchargetime) - Retrieve charging completion time for a Cadillac.
* [GetVoltage](#getvoltage) - Retrieve charging voltmeter time for a Cadillac.

## GetChargeTime

__Description__

When the vehicle is charging, this endpoint returns the date and time the vehicle expects to complete this charging session. When the vehicle is not charging, this endpoint results in a vehicle state error.

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


    var vehicleID string = "string"

    ctx := context.Background()
    res, err := s.Cadillac.GetChargeTime(ctx, vehicleID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeTime != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `vehicleID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetCadillacChargeTimeResponse](../../pkg/models/operations/getcadillacchargetimeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetVoltage

__Description__

When the vehicle is plugged in, this endpoint returns the voltage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.

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


    var vehicleID string = "string"

    ctx := context.Background()
    res, err := s.Cadillac.GetVoltage(ctx, vehicleID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeVoltage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `vehicleID`                                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetCadillacVoltageResponse](../../pkg/models/operations/getcadillacvoltageresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

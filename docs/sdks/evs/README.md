# Evs
(*Evs*)

## Overview

Operations about electric vehicles

### Available Operations

* [GetBatteryCapacity](#getbatterycapacity) - EV Battery Capacity
* [GetBatteryLevel](#getbatterylevel) - EV Battery Level
* [GetChargingLimit](#getcharginglimit) - EV Charging Limit
* [GetChargingStatus](#getchargingstatus) - EV Charging Status
* [SetChargingLimit](#setcharginglimit) - Set EV Charging Limit
* [StartStopCharge](#startstopcharge) - Start or stop charging an electric vehicle. Please contact us to request early access.

## GetBatteryCapacity

__Description__

Returns the total capacity of an electric vehicle's battery.

__Permission__

`read_battery`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  capacity|   number|  The total capacity of the vehicle's battery (in kilowatt-hours). 	|

### Example Usage

```go
package main

import(
	"testsdkcreation/pkg/models/shared"
	"testsdkcreation"
	"context"
	"log"
)

func main() {
    s := testsdkcreation.New(
        testsdkcreation.WithSecurity(""),
    )


    var vehicleID string = "string"

    ctx := context.Background()
    res, err := s.Evs.GetBatteryCapacity(ctx, vehicleID)
    if err != nil {
        log.Fatal(err)
    }

    if res.BatteryCapacity != nil {
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

**[*operations.GetBatteryCapacityResponse](../../pkg/models/operations/getbatterycapacityresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetBatteryLevel

__Description__

Retrieve EV battery level of a vehicle.

__Permission__

`read_battery`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  `percentRemaining`|   number|  An EV battery’s state of charge (in percent). 	|
|   `range`|   number	|   The estimated remaining distance the vehicle can travel (in kilometers by default or in miles using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers).	|

### Example Usage

```go
package main

import(
	"testsdkcreation/pkg/models/shared"
	"testsdkcreation"
	"context"
	"log"
)

func main() {
    s := testsdkcreation.New(
        testsdkcreation.WithSecurity(""),
    )


    var vehicleID string = "string"

    ctx := context.Background()
    res, err := s.Evs.GetBatteryLevel(ctx, vehicleID)
    if err != nil {
        log.Fatal(err)
    }

    if res.BatteryLevel != nil {
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

**[*operations.GetBatteryLevelResponse](../../pkg/models/operations/getbatterylevelresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetChargingLimit

__Description__

Returns the current charge limit of an electric vehicle.

### Example Usage

```go
package main

import(
	"testsdkcreation/pkg/models/shared"
	"testsdkcreation"
	"context"
	"log"
)

func main() {
    s := testsdkcreation.New(
        testsdkcreation.WithSecurity(""),
    )


    var vehicleID string = "string"

    ctx := context.Background()
    res, err := s.Evs.GetChargingLimit(ctx, vehicleID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeLimit != nil {
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

**[*operations.GetChargingLimitResponse](../../pkg/models/operations/getcharginglimitresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetChargingStatus

__Description__

Returns the current charge status of an electric vehicle.

__Permission__

`read_charge`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  `isPluggedIn` 	|   boolean	|  Indicates whether a charging cable is currently plugged into the vehicle’s charge port. 	|
|   `state`	|   string	|   Indicates whether the vehicle is currently charging. Options: `CHARGING` `FULLY_CHARGED` `NOT_CHARGING`	|

### Example Usage

```go
package main

import(
	"testsdkcreation/pkg/models/shared"
	"testsdkcreation"
	"context"
	"log"
)

func main() {
    s := testsdkcreation.New(
        testsdkcreation.WithSecurity(""),
    )


    var vehicleID string = "string"

    ctx := context.Background()
    res, err := s.Evs.GetChargingStatus(ctx, vehicleID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeStatus != nil {
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

**[*operations.GetChargingStatusResponse](../../pkg/models/operations/getchargingstatusresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## SetChargingLimit

__Description__

Returns the current charge limit of an electric vehicle.

### Example Usage

```go
package main

import(
	"testsdkcreation/pkg/models/shared"
	"testsdkcreation"
	"context"
	"log"
)

func main() {
    s := testsdkcreation.New(
        testsdkcreation.WithSecurity(""),
    )


    var vehicleID string = "string"

    chargeLimit := &shared.ChargeLimit{
        Limit: testsdkcreation.Float32(1),
    }

    ctx := context.Background()
    res, err := s.Evs.SetChargingLimit(ctx, vehicleID, chargeLimit)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `vehicleID`                                                      | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `chargeLimit`                                                    | [*shared.ChargeLimit](../../../pkg/models/shared/chargelimit.md) | :heavy_minus_sign:                                               | N/A                                                              |


### Response

**[*operations.SetChargingLimitResponse](../../pkg/models/operations/setcharginglimitresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## StartStopCharge

__Description__

Returns the current charge status of an electric vehicle.

__Permission__

`read_charge`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  `isPluggedIn` 	|   boolean	|  Indicates whether a charging cable is currently plugged into the vehicle’s charge port. 	|
|   `state`	|   string	|   Indicates whether the vehicle is currently charging. Options: `CHARGING` `FULLY_CHARGED` `NOT_CHARGING`	|

### Example Usage

```go
package main

import(
	"testsdkcreation/pkg/models/shared"
	"testsdkcreation"
	"context"
	"log"
)

func main() {
    s := testsdkcreation.New(
        testsdkcreation.WithSecurity(""),
    )


    var vehicleID string = "string"

    chargeAction := &shared.ChargeAction{
        Action: shared.ActionStart.ToPointer(),
    }

    ctx := context.Background()
    res, err := s.Evs.StartStopCharge(ctx, vehicleID, chargeAction)
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `vehicleID`                                                        | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `chargeAction`                                                     | [*shared.ChargeAction](../../../pkg/models/shared/chargeaction.md) | :heavy_minus_sign:                                                 | N/A                                                                |


### Response

**[*operations.StartStopChargeResponse](../../pkg/models/operations/startstopchargeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

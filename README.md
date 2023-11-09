# testsdkcreation

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/gotest.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/gotest/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README
<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/gotest
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"log"
	"testsdkcreation"
	"testsdkcreation/pkg/models/shared"
)

func main() {
	s := testsdkcreation.New(
		testsdkcreation.WithSecurity(""),
	)

	var vehicleID string = "36ab27d0-fd9d-4455-823a-ce30af709ffc"

	ctx := context.Background()
	res, err := s.Vehicles.GetLocation(ctx, vehicleID)
	if err != nil {
		log.Fatal(err)
	}

	if res.Location != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Compatibility](docs/sdks/compatibility/README.md)

* [ListCompatibility](docs/sdks/compatibility/README.md#listcompatibility) - Compatibility

### [User](docs/sdks/user/README.md)

* [GetInfo](docs/sdks/user/README.md#getinfo) - User Info

### [Vehicles](docs/sdks/vehicles/README.md)

* [Batch](docs/sdks/vehicles/README.md#batch) - Batch
* [Disconnect](docs/sdks/vehicles/README.md#disconnect) - Revoke Access
* [Get](docs/sdks/vehicles/README.md#get) - Vehicle Info
* [GetEngineOil](docs/sdks/vehicles/README.md#getengineoil) - Engine Oil Life
* [GetFuelTank](docs/sdks/vehicles/README.md#getfueltank) - Fuel Tank (US Only)
* [GetLocation](docs/sdks/vehicles/README.md#getlocation) - Location
* [GetOdometer](docs/sdks/vehicles/README.md#getodometer) - Odometer
* [GetPermissions](docs/sdks/vehicles/README.md#getpermissions) - Application Permissions
* [GetTirePressure](docs/sdks/vehicles/README.md#gettirepressure) - Tire pressure
* [GetVin](docs/sdks/vehicles/README.md#getvin) - Returns the vehicle’s manufacturer identifier.
* [ListVehicles](docs/sdks/vehicles/README.md#listvehicles) - All Vehicles
* [LockUnlock](docs/sdks/vehicles/README.md#lockunlock) - Lock/Unlock Vehicle

### [Tesla](docs/sdks/tesla/README.md)

* [GetAmmeter](docs/sdks/tesla/README.md#getammeter) - Retrieve charging ammeter time for a Tesla.
* [GetChargeTime](docs/sdks/tesla/README.md#getchargetime) - Retrieve charging completion time for a Tesla.
* [GetCompass](docs/sdks/tesla/README.md#getcompass) - Retrieve compass heading for a Tesla.
* [GetExteriorTemperature](docs/sdks/tesla/README.md#getexteriortemperature) - Retrieve exterior temperature for a Tesla.
* [GetInteriorTemperature](docs/sdks/tesla/README.md#getinteriortemperature) - Retrieve interior temperature for a Tesla.
* [GetSpeedometer](docs/sdks/tesla/README.md#getspeedometer) - Retrieve speed for a Tesla.
* [GetVoltage](docs/sdks/tesla/README.md#getvoltage) - Retrieve charging voltmeter time for a Tesla.
* [GetWattmeter](docs/sdks/tesla/README.md#getwattmeter) - Retrieve charging wattmeter time for a Tesla.
* [SetAmmeter](docs/sdks/tesla/README.md#setammeter) - Set charging ammeter time for a Tesla.

### [Evs](docs/sdks/evs/README.md)

* [GetBatteryCapacity](docs/sdks/evs/README.md#getbatterycapacity) - EV Battery Capacity
* [GetBatteryLevel](docs/sdks/evs/README.md#getbatterylevel) - EV Battery Level
* [GetChargingLimit](docs/sdks/evs/README.md#getcharginglimit) - EV Charging Limit
* [GetChargingStatus](docs/sdks/evs/README.md#getchargingstatus) - EV Charging Status
* [SetChargingLimit](docs/sdks/evs/README.md#setcharginglimit) - Set EV Charging Limit
* [StartStopCharge](docs/sdks/evs/README.md#startstopcharge) - Start or stop charging an electric vehicle. Please contact us to request early access.

### [Cadillac](docs/sdks/cadillac/README.md)

* [GetChargeTime](docs/sdks/cadillac/README.md#getchargetime) - Retrieve charging completion time for a Cadillac.
* [GetVoltage](docs/sdks/cadillac/README.md#getvoltage) - Retrieve charging voltmeter time for a Cadillac.

### [Chevrolet](docs/sdks/chevrolet/README.md)

* [GetChargeTime](docs/sdks/chevrolet/README.md#getchargetime) - Retrieve charging completion time for a Chevrolet.
* [GetVoltage](docs/sdks/chevrolet/README.md#getvoltage) - Retrieve charging voltmeter time for a Chevrolet.

### [Webhooks](docs/sdks/webhooks/README.md)

* [Subscribe](docs/sdks/webhooks/README.md#subscribe) - Subscribe Webhook
* [Unsubscribe](docs/sdks/webhooks/README.md#unsubscribe) - Unsubscribe Webhook
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->

<!-- End Dev Containers -->

<!-- Start Error Handling -->
# Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |


## Example

```go
package main

import (
	"context"
	"log"
	"testsdkcreation"
	"testsdkcreation/pkg/models/shared"
)

func main() {
	s := testsdkcreation.New(
		testsdkcreation.WithSecurity(""),
	)

	var country *string = "{country}"

	var scope *string = "{scope}"

	var vin *string = "{vin}"

	ctx := context.Background()
	res, err := s.Compatibility.ListCompatibility(ctx, country, scope, vin)
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->

<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.smartcar.com/v2.0` | None |

For example:

```go
package main

import (
	"context"
	"log"
	"testsdkcreation"
	"testsdkcreation/pkg/models/shared"
)

func main() {
	s := testsdkcreation.New(
		testsdkcreation.WithServerIndex(0),
		testsdkcreation.WithSecurity(""),
	)

	var country *string = "{country}"

	var scope *string = "{scope}"

	var vin *string = "{vin}"

	ctx := context.Background()
	res, err := s.Compatibility.ListCompatibility(ctx, country, scope, vin)
	if err != nil {
		log.Fatal(err)
	}

	if res.CompatibilityResponse != nil {
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
	"testsdkcreation"
	"testsdkcreation/pkg/models/shared"
)

func main() {
	s := testsdkcreation.New(
		testsdkcreation.WithServerURL("https://api.smartcar.com/v2.0"),
		testsdkcreation.WithSecurity(""),
	)

	var country *string = "{country}"

	var scope *string = "{scope}"

	var vin *string = "{vin}"

	ctx := context.Background()
	res, err := s.Compatibility.ListCompatibility(ctx, country, scope, vin)
	if err != nil {
		log.Fatal(err)
	}

	if res.CompatibilityResponse != nil {
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

<!-- End Go Types -->



<!-- Start Authentication -->
# Authentication

## Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `BearerAuth` | http         | HTTP Bearer  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	"log"
	"testsdkcreation"
	"testsdkcreation/pkg/models/shared"
)

func main() {
	s := testsdkcreation.New(
		testsdkcreation.WithSecurity(""),
	)

	var country *string = "{country}"

	var scope *string = "{scope}"

	var vin *string = "{vin}"

	ctx := context.Background()
	res, err := s.Compatibility.ListCompatibility(ctx, country, scope, vin)
	if err != nil {
		log.Fatal(err)
	}

	if res.CompatibilityResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

# testsdkcreation

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

import(
	"context"
	"log"
	"testsdkcreation"
	"testsdkcreation/pkg/models/shared"
)

func main() {
    s := testsdkcreation.New(
        testsdkcreation.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
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


### [Cadillac](docs/sdks/cadillac/README.md)

* [GetChargeTime](docs/sdks/cadillac/README.md#getchargetime) - Retrieve charging completion time for a Cadillac.
* [GetVoltage](docs/sdks/cadillac/README.md#getvoltage) - Retrieve charging voltmeter time for a Cadillac.

### [Chevrolet](docs/sdks/chevrolet/README.md)

* [GetChargeTime](docs/sdks/chevrolet/README.md#getchargetime) - Retrieve charging completion time for a Chevrolet.
* [GetVoltage](docs/sdks/chevrolet/README.md#getvoltage) - Retrieve charging voltmeter time for a Chevrolet.

### [Compatibility](docs/sdks/compatibility/README.md)

* [ListCompatibility](docs/sdks/compatibility/README.md#listcompatibility) - Compatibility

### [Evs](docs/sdks/evs/README.md)

* [GetBatteryCapacity](docs/sdks/evs/README.md#getbatterycapacity) - EV Battery Capacity
* [GetBatteryLevel](docs/sdks/evs/README.md#getbatterylevel) - EV Battery Level
* [GetChargingLimit](docs/sdks/evs/README.md#getcharginglimit) - EV Charging Limit
* [GetChargingStatus](docs/sdks/evs/README.md#getchargingstatus) - EV Charging Status
* [SetChargingLimit](docs/sdks/evs/README.md#setcharginglimit) - Set EV Charging Limit
* [StartStopCharge](docs/sdks/evs/README.md#startstopcharge) - Start or stop charging an electric vehicle. Please contact us to request early access.

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

### [Webhooks](docs/sdks/webhooks/README.md)

* [Subscribe](docs/sdks/webhooks/README.md#subscribe) - Subscribe Webhook
* [Unsubscribe](docs/sdks/webhooks/README.md#unsubscribe) - Unsubscribe Webhook
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

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

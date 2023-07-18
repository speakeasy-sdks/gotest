// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package testsdkcreation

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testsdkcreation/pkg/models/operations"
	"testsdkcreation/pkg/models/sdkerrors"
	"testsdkcreation/pkg/models/shared"
	"testsdkcreation/pkg/utils"
)

// vehicles - Operations about vehicles
type vehicles struct {
	sdkConfiguration sdkConfiguration
}

func newVehicles(sdkConfig sdkConfiguration) *vehicles {
	return &vehicles{
		sdkConfiguration: sdkConfig,
	}
}

// Batch - Batch
// __Description__ Returns a list of responses from multiple Smartcar endpoints, all combined into a single request. Note: Batch requests is a paid feature. Please contact us to upgrade your plan and obtain access.
func (s *vehicles) Batch(ctx context.Context, vehicleID string, requestBody []string) (*operations.BatchResponse, error) {
	request := operations.BatchRequest{
		VehicleID:   vehicleID,
		RequestBody: requestBody,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/batch", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.BatchResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.BatchResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.BatchResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// Disconnect - Revoke Access
// __Description__
//
// Revoke access for the current requesting application. This is the correct way to disconnect a vehicle.
//
// __Response body__
//
// |  Name 	|Type   	|Boolean   	|
// |---	|---	|---	|
// |  status|   string|  If the request is successful, Smartcar will return “success” (HTTP 200 status).|
func (s *vehicles) Disconnect(ctx context.Context, vehicleID string) (*operations.DisconnectResponse, error) {
	request := operations.DisconnectRequest{
		VehicleID: vehicleID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/application", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DisconnectResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.DisconnectStatus
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Status = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// Get - Vehicle Info
// __Permission__
//
// Returns a single vehicle object, containing identifying information.
//
// __Permission__
//
// `read_vehicle_info`
//
// __Response Body__
//
// |Name| Type|Description|
// |--|--|--|
// |`id`|string|A vehicle ID (UUID v4).|
// |`make`|string|The manufacturer of the vehicle.|
// |`model`|integer|The model of the vehicle.|
// |`year`|integer|The model year.|
func (s *vehicles) Get(ctx context.Context, vehicleID string) (*operations.GetVehicleResponse, error) {
	request := operations.GetVehicleRequest{
		VehicleID: vehicleID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetVehicleResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.VehicleInfo
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.VehicleInfo = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetEngineOil - Engine Oil Life
// __Description__
//
// Returns the remaining life span of a vehicle’s engine oil.
//
// __Permission__
//
// `read_engine_oil`
//
// __Response body__
//
// |  Name 	|Type   	|Boolean   	|
// |---	|---	|---	|
// |  `lifeRemaining`|   number|  The engine oil’s remaining life span (as a percentage). Oil life is based on the current quality of the oil. (in percent).|
func (s *vehicles) GetEngineOil(ctx context.Context, vehicleID string) (*operations.GetEngineOilResponse, error) {
	request := operations.GetEngineOilRequest{
		VehicleID: vehicleID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/engine/oil", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetEngineOilResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.EngineOil
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.EngineOil = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetFuelTank - Fuel Tank (US Only)
// __Description__
//
// Returns the status of the fuel remaining in the vehicle’s gas tank. Note: The fuel tank API is only available for vehicles sold in the United States.
//
// __Permission__
//
// `read_fuel`
//
// __Response Body__
//
// |Name| Type|Desciprtion|
// |--|--|--|
// |`range`|number|The estimated remaining distance the car can travel (in kilometers by default or in miles using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
// |`percentRemaining`|number|The remaining level of fuel in the tank (in percent).|
// |`amountRemaining`|number|The amount of fuel in the tank (in liters by default or in gallons (U.S.) using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
func (s *vehicles) GetFuelTank(ctx context.Context, vehicleID string) (*operations.GetFuelTankResponse, error) {
	request := operations.GetFuelTankRequest{
		VehicleID: vehicleID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/fuel", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetFuelTankResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.FuelTank
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.FuelTank = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetLocation - Location
// __Description__
//
// Retrieve latitude and longitude of a vehicle.
//
// __Permission__
//
// `read_location`
//
// __Response Body__
//
// |Name| Type|Desciprtion|
// |--|--|--|
// |`latitude`|number|The latitude (in degrees).|
// |`longitude`|number|The longitude (in degrees).|
func (s *vehicles) GetLocation(ctx context.Context, vehicleID string) (*operations.GetLocationResponse, error) {
	request := operations.GetLocationRequest{
		VehicleID: vehicleID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/location", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetLocationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Location
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Location = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetOdometer - Odometer
// __Description__
//
// Returns the vehicle’s last known odometer reading.
//
// __Permission__
//
// `read_odometer`
//
// __Response Body__
//
// |Name| Type|Desciprtion|
// |--|--|--|
// |`distance`|number|The current odometer of the vehicle (in kilometers by default or in miles using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
func (s *vehicles) GetOdometer(ctx context.Context, vehicleID string) (*operations.GetOdometerResponse, error) {
	request := operations.GetOdometerRequest{
		VehicleID: vehicleID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/odometer", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetOdometerResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Odometer
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Odometer = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetPermissions - Application Permissions
// __Description__
//
// Returns a list of the permissions that have been granted to your application in relation to this vehicle.
//
// __Query Parameters__
//
// |Parameter| Type|Required|Description|
// |--|--|--|--|
// |`limit`|integer|false|Number of permissions to return (default: 25).|
// |`offset`|integer|false|A vehicle ID (UUID v4).|Index to start permission list at|
//
// __Response Body__
//
// |Name| Type|Desciprtion|
// |--|--|--|
// |`permissions`|array|An array of permissions.|
// |`permissions.[]`|string|A permission.|
// |`paging`|object|Metadata about the current list of elements.|
// |`paging.count`|integer|The total number of elements for the entire query (not just the given page).|
// |`paging.offset`|integer|The current start index of the returned list of elements.|
func (s *vehicles) GetPermissions(ctx context.Context, vehicleID string, limit *int64, offset *int64) (*operations.GetPermissionsResponse, error) {
	request := operations.GetPermissionsRequest{
		VehicleID: vehicleID,
		Limit:     limit,
		Offset:    offset,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/permissions", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetPermissionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Permission
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Permission = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetTirePressure - Tire pressure
// __Description__
//
// Returns the air pressure of each of the vehicle’s tires.
// __Permission__
//
// `read_tires`
//
// __Example Response__
//
// |Name| Type|Description|
// |--|--|--|
// |`frontLeft`|number|The current air pressure of the front left tire (in kilopascals by default or in pounds per square inch using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
// |`frontRight`|number|The current air pressure of the front right tire (in kilopascals by default or in pounds per square inch using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
// |`backLeft`|number|The current air pressure of the back left tire (in kilopascals by default or in pounds per square inch using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
// |`backRight`|number|The current air pressure of the back right tire (in kilopascals by default or in pounds per square inch using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers)).|
func (s *vehicles) GetTirePressure(ctx context.Context, vehicleID string) (*operations.GetTirePressureResponse, error) {
	request := operations.GetTirePressureRequest{
		VehicleID: vehicleID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/tires/pressure", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTirePressureResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.TirePressure
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.TirePressure = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetVin - Returns the vehicle’s manufacturer identifier.
// __Description__
//
// Returns the vehicle’s manufacturer identifier.
func (s *vehicles) GetVin(ctx context.Context, vehicleID string) (*operations.GetVinResponse, error) {
	request := operations.GetVinRequest{
		VehicleID: vehicleID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/vin", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetVinResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.VinInfo
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.VinInfo = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// ListVehicles - All Vehicles
// __Description__
//
// Returns a paged list of all vehicles connected to the application for the current authorized user.
//
// __Query Parameters__
//
// |Parameter| Type|Required|Description|
// |--|--|--|--|
// |`limit`|integer|false|Number of vehicles to return (default: 10).|
// |`offset`|integer|false|A vehicle ID (UUID v4).|Index to start vehicle list at|
//
// __Response Body__
//
// |Name| Type|Desciprtion|
// |--|--|--|
// |`vehicles`|array|An array of vehicle IDs.|
// |`vehicles.[]`|string|A vehicle ID (UUID v4).|
// |`paging`|object|Metadata about the current list of elements.|
// |`paging.count`|integer|The total number of elements for the entire query (not just the given page).|
// |`paging.offset`|integer|The current start index of the returned list of elements.|
func (s *vehicles) ListVehicles(ctx context.Context, limit *int64, offset *int64) (*operations.ListVehiclesResponse, error) {
	request := operations.ListVehiclesRequest{
		Limit:  limit,
		Offset: offset,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/vehicles"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListVehiclesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.VehiclesResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.VehiclesResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// LockUnlock - Lock/Unlock Vehicle
// __Description__
//
// # Unlock the vehicle
//
// __Permission__
//
// `control_security`
//
// __Response body__
//
// |  Name 	|Type   	|Boolean   	|
// |---	|---	|---	|
// |  status|   string|  If the request is successful, Smartcar will return “success” (HTTP 200 status).|
func (s *vehicles) LockUnlock(ctx context.Context, vehicleID string, securityAction *shared.SecurityAction) (*operations.LockUnlockResponse, error) {
	request := operations.LockUnlockRequest{
		VehicleID:      vehicleID,
		SecurityAction: securityAction,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/security", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "SecurityAction", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.LockUnlockResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SuccessResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.SuccessResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_SubscriberDataManagement

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/antihax/optional"

	"encoding/json"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
)

// Linger please
var (
	_ context.Context
)

type SessionManagementSubscriptionDataRetrievalApiService service

/*
SessionManagementSubscriptionDataRetrievalApiService retrieve a UE's Session Management Subscription Data
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param supi Identifier of the UE
 * @param optional nil or *GetSmDataParamOpts - Optional Parameters:
 * @param "SupportedFeatures" (optional.String) -  Supported Features
 * @param "SingleNssai" (optional.Interface of models.Snssai) -
 * @param "Dnn" (optional.String) -
 * @param "PlmnId" (optional.Interface of models.PlmnId) -
 * @param "IfNoneMatch" (optional.String) -  Validator for conditional requests, as described in RFC 7232, 3.2
 * @param "IfModifiedSince" (optional.String) -  Validator for conditional requests, as described in RFC 7232, 3.3
@return []models.SessionManagementSubscriptionData
*/

type GetSmDataParamOpts struct {
	SupportedFeatures optional.String
	SingleNssai       optional.Interface
	Dnn               optional.String
	PlmnId            optional.Interface
	IfNoneMatch       optional.String
	IfModifiedSince   optional.String
}

func (a *SessionManagementSubscriptionDataRetrievalApiService) GetSmData(ctx context.Context, supi string, localVarOptionals *GetSmDataParamOpts) ([]models.SessionManagementSubscriptionData, *http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []models.SessionManagementSubscriptionData
		tmpLocalVarReturnValue  interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{supi}/sm-data"
	localVarPath = strings.Replace(localVarPath, "{"+"supi"+"}", fmt.Sprintf("%v", supi), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.SupportedFeatures.IsSet() && localVarOptionals.SupportedFeatures.Value() != "" {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(localVarOptionals.SupportedFeatures.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SingleNssai.IsSet() {
		localVarQueryParams.Add("single-nssai", openapi.ParameterToString(localVarOptionals.SingleNssai.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Dnn.IsSet() && localVarOptionals.Dnn.Value() != "" {
		localVarQueryParams.Add("dnn", openapi.ParameterToString(localVarOptionals.Dnn.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PlmnId.IsSet() {
		localVarQueryParams.Add("plmn-id", openapi.ParameterToString(localVarOptionals.PlmnId.Value(), ""))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() && localVarOptionals.IfNoneMatch.Value() != "" {
		localVarHeaderParams["If-None-Match"] = openapi.ParameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.IfModifiedSince.IsSet() && localVarOptionals.IfModifiedSince.Value() != "" {
		localVarHeaderParams["If-Modified-Since"] = openapi.ParameterToString(localVarOptionals.IfModifiedSince.Value(), "")
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		jsonout, _ := json.Marshal(tmpLocalVarReturnValue)
		fmt.Println(string(jsonout))
		v := reflect.TypeOf(tmpLocalVarReturnValue)
		fmt.Println("res date type is ", v.Kind())
		if v.Kind() == reflect.Slice{
			fmt.Println("res date type is slice")
			if err := json.Unmarshal(jsonout, &localVarReturnValue); err != nil {
				panic(err)
			}
		} else {
			var tmp models.SessionManagementSubscriptionData;
			if err := json.Unmarshal(jsonout, &tmp); err != nil {
				panic(err)
			}
			localVarReturnValue = []models.SessionManagementSubscriptionData{tmp}
		}
		return localVarReturnValue, localVarHTTPResponse, nil
	case 400:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 404:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 500:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	case 503:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	default:
		return localVarReturnValue, localVarHTTPResponse, nil
	}
}

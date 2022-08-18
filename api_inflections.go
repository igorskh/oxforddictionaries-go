/*
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.5.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxforddictionaries

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// InflectionsApiService InflectionsApi service
type InflectionsApiService service

type ApiInflectionsSourceLangWordIdGetRequest struct {
	ctx _context.Context
	ApiService *InflectionsApiService
	sourceLang string
	wordId string
	appId *string
	appKey *string
	strictMatch *bool
}

func (r ApiInflectionsSourceLangWordIdGetRequest) AppId(appId string) ApiInflectionsSourceLangWordIdGetRequest {
	r.appId = &appId
	return r
}
func (r ApiInflectionsSourceLangWordIdGetRequest) AppKey(appKey string) ApiInflectionsSourceLangWordIdGetRequest {
	r.appKey = &appKey
	return r
}
func (r ApiInflectionsSourceLangWordIdGetRequest) StrictMatch(strictMatch bool) ApiInflectionsSourceLangWordIdGetRequest {
	r.strictMatch = &strictMatch
	return r
}

func (r ApiInflectionsSourceLangWordIdGetRequest) Execute() (Inflections, *_nethttp.Response, error) {
	return r.ApiService.InflectionsSourceLangWordIdGetExecute(r)
}

/*
 * InflectionsSourceLangWordIdGet Retrieves the inflected forms of a given word.
 * Retrieve all the [inflections](documentation/glossary?term=inflection) of a given word_id. The inflections
are given for a lexicalEntry with a specific lexicalCategory.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sourceLang Language code of the source language in a monolingual dataset.
 * @param wordId The identifier for an Entry (case-sensitive).
 * @return ApiInflectionsSourceLangWordIdGetRequest
 */
func (a *InflectionsApiService) InflectionsSourceLangWordIdGet(ctx _context.Context, sourceLang string, wordId string) ApiInflectionsSourceLangWordIdGetRequest {
	return ApiInflectionsSourceLangWordIdGetRequest{
		ApiService: a,
		ctx: ctx,
		sourceLang: sourceLang,
		wordId: wordId,
	}
}

/*
 * Execute executes the request
 * @return Inflections
 */
func (a *InflectionsApiService) InflectionsSourceLangWordIdGetExecute(r ApiInflectionsSourceLangWordIdGetRequest) (Inflections, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Inflections
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InflectionsApiService.InflectionsSourceLangWordIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inflections/{source_lang}/{word_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"source_lang"+"}", _neturl.PathEscape(parameterToString(r.sourceLang, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"word_id"+"}", _neturl.PathEscape(parameterToString(r.wordId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.appId == nil {
		return localVarReturnValue, nil, reportError("appId is required and must be specified")
	}
	if r.appKey == nil {
		return localVarReturnValue, nil, reportError("appKey is required and must be specified")
	}

	if r.strictMatch != nil {
		localVarQueryParams.Add("strictMatch", parameterToString(*r.strictMatch, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["app_id"] = parameterToString(*r.appId, "")
	localVarHeaderParams["app_key"] = parameterToString(*r.appKey, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ErrorSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

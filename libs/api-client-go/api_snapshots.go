/*
Daytona

Daytona AI platform API Docs

API version: 1.0
Contact: support@daytona.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type SnapshotsAPI interface {

	/*
		ActivateSnapshot Activate a snapshot

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Snapshot ID
		@return SnapshotsAPIActivateSnapshotRequest
	*/
	ActivateSnapshot(ctx context.Context, id string) SnapshotsAPIActivateSnapshotRequest

	// ActivateSnapshotExecute executes the request
	//  @return SnapshotDto
	ActivateSnapshotExecute(r SnapshotsAPIActivateSnapshotRequest) (*SnapshotDto, *http.Response, error)

	/*
		CanCleanupImage Check if an image can be cleaned up

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return SnapshotsAPICanCleanupImageRequest
	*/
	CanCleanupImage(ctx context.Context) SnapshotsAPICanCleanupImageRequest

	// CanCleanupImageExecute executes the request
	//  @return bool
	CanCleanupImageExecute(r SnapshotsAPICanCleanupImageRequest) (bool, *http.Response, error)

	/*
		CreateSnapshot Create a new snapshot

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return SnapshotsAPICreateSnapshotRequest
	*/
	CreateSnapshot(ctx context.Context) SnapshotsAPICreateSnapshotRequest

	// CreateSnapshotExecute executes the request
	//  @return SnapshotDto
	CreateSnapshotExecute(r SnapshotsAPICreateSnapshotRequest) (*SnapshotDto, *http.Response, error)

	/*
		DeactivateSnapshot Deactivate a snapshot

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Snapshot ID
		@return SnapshotsAPIDeactivateSnapshotRequest
	*/
	DeactivateSnapshot(ctx context.Context, id string) SnapshotsAPIDeactivateSnapshotRequest

	// DeactivateSnapshotExecute executes the request
	DeactivateSnapshotExecute(r SnapshotsAPIDeactivateSnapshotRequest) (*http.Response, error)

	/*
		GetAllSnapshots List all snapshots

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return SnapshotsAPIGetAllSnapshotsRequest
	*/
	GetAllSnapshots(ctx context.Context) SnapshotsAPIGetAllSnapshotsRequest

	// GetAllSnapshotsExecute executes the request
	//  @return PaginatedSnapshotsDto
	GetAllSnapshotsExecute(r SnapshotsAPIGetAllSnapshotsRequest) (*PaginatedSnapshotsDto, *http.Response, error)

	/*
		GetSnapshot Get snapshot by ID or name

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Snapshot ID or name
		@return SnapshotsAPIGetSnapshotRequest
	*/
	GetSnapshot(ctx context.Context, id string) SnapshotsAPIGetSnapshotRequest

	// GetSnapshotExecute executes the request
	//  @return SnapshotDto
	GetSnapshotExecute(r SnapshotsAPIGetSnapshotRequest) (*SnapshotDto, *http.Response, error)

	/*
		GetSnapshotBuildLogs Get snapshot build logs

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Snapshot ID
		@return SnapshotsAPIGetSnapshotBuildLogsRequest
	*/
	GetSnapshotBuildLogs(ctx context.Context, id string) SnapshotsAPIGetSnapshotBuildLogsRequest

	// GetSnapshotBuildLogsExecute executes the request
	GetSnapshotBuildLogsExecute(r SnapshotsAPIGetSnapshotBuildLogsRequest) (*http.Response, error)

	/*
		RemoveSnapshot Delete snapshot

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Snapshot ID
		@return SnapshotsAPIRemoveSnapshotRequest
	*/
	RemoveSnapshot(ctx context.Context, id string) SnapshotsAPIRemoveSnapshotRequest

	// RemoveSnapshotExecute executes the request
	RemoveSnapshotExecute(r SnapshotsAPIRemoveSnapshotRequest) (*http.Response, error)

	/*
		SetSnapshotGeneralStatus Set snapshot general status

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id Snapshot ID
		@return SnapshotsAPISetSnapshotGeneralStatusRequest
	*/
	SetSnapshotGeneralStatus(ctx context.Context, id string) SnapshotsAPISetSnapshotGeneralStatusRequest

	// SetSnapshotGeneralStatusExecute executes the request
	//  @return SnapshotDto
	SetSnapshotGeneralStatusExecute(r SnapshotsAPISetSnapshotGeneralStatusRequest) (*SnapshotDto, *http.Response, error)
}

// SnapshotsAPIService SnapshotsAPI service
type SnapshotsAPIService service

type SnapshotsAPIActivateSnapshotRequest struct {
	ctx                    context.Context
	ApiService             SnapshotsAPI
	id                     string
	xDaytonaOrganizationID *string
}

// Use with JWT to specify the organization ID
func (r SnapshotsAPIActivateSnapshotRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) SnapshotsAPIActivateSnapshotRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r SnapshotsAPIActivateSnapshotRequest) Execute() (*SnapshotDto, *http.Response, error) {
	return r.ApiService.ActivateSnapshotExecute(r)
}

/*
ActivateSnapshot Activate a snapshot

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Snapshot ID
	@return SnapshotsAPIActivateSnapshotRequest
*/
func (a *SnapshotsAPIService) ActivateSnapshot(ctx context.Context, id string) SnapshotsAPIActivateSnapshotRequest {
	return SnapshotsAPIActivateSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return SnapshotDto
func (a *SnapshotsAPIService) ActivateSnapshotExecute(r SnapshotsAPIActivateSnapshotRequest) (*SnapshotDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SnapshotDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.ActivateSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/snapshots/{id}/activate"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SnapshotsAPICanCleanupImageRequest struct {
	ctx                    context.Context
	ApiService             SnapshotsAPI
	imageName              *string
	xDaytonaOrganizationID *string
}

// Image name with tag to check
func (r SnapshotsAPICanCleanupImageRequest) ImageName(imageName string) SnapshotsAPICanCleanupImageRequest {
	r.imageName = &imageName
	return r
}

// Use with JWT to specify the organization ID
func (r SnapshotsAPICanCleanupImageRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) SnapshotsAPICanCleanupImageRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r SnapshotsAPICanCleanupImageRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.CanCleanupImageExecute(r)
}

/*
CanCleanupImage Check if an image can be cleaned up

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SnapshotsAPICanCleanupImageRequest
*/
func (a *SnapshotsAPIService) CanCleanupImage(ctx context.Context) SnapshotsAPICanCleanupImageRequest {
	return SnapshotsAPICanCleanupImageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return bool
func (a *SnapshotsAPIService) CanCleanupImageExecute(r SnapshotsAPICanCleanupImageRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.CanCleanupImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/snapshots/can-cleanup-image"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.imageName == nil {
		return localVarReturnValue, nil, reportError("imageName is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "imageName", r.imageName, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SnapshotsAPICreateSnapshotRequest struct {
	ctx                    context.Context
	ApiService             SnapshotsAPI
	createSnapshot         *CreateSnapshot
	xDaytonaOrganizationID *string
}

func (r SnapshotsAPICreateSnapshotRequest) CreateSnapshot(createSnapshot CreateSnapshot) SnapshotsAPICreateSnapshotRequest {
	r.createSnapshot = &createSnapshot
	return r
}

// Use with JWT to specify the organization ID
func (r SnapshotsAPICreateSnapshotRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) SnapshotsAPICreateSnapshotRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r SnapshotsAPICreateSnapshotRequest) Execute() (*SnapshotDto, *http.Response, error) {
	return r.ApiService.CreateSnapshotExecute(r)
}

/*
CreateSnapshot Create a new snapshot

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SnapshotsAPICreateSnapshotRequest
*/
func (a *SnapshotsAPIService) CreateSnapshot(ctx context.Context) SnapshotsAPICreateSnapshotRequest {
	return SnapshotsAPICreateSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SnapshotDto
func (a *SnapshotsAPIService) CreateSnapshotExecute(r SnapshotsAPICreateSnapshotRequest) (*SnapshotDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SnapshotDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.CreateSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/snapshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createSnapshot == nil {
		return localVarReturnValue, nil, reportError("createSnapshot is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	// body params
	localVarPostBody = r.createSnapshot
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SnapshotsAPIDeactivateSnapshotRequest struct {
	ctx                    context.Context
	ApiService             SnapshotsAPI
	id                     string
	xDaytonaOrganizationID *string
}

// Use with JWT to specify the organization ID
func (r SnapshotsAPIDeactivateSnapshotRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) SnapshotsAPIDeactivateSnapshotRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r SnapshotsAPIDeactivateSnapshotRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeactivateSnapshotExecute(r)
}

/*
DeactivateSnapshot Deactivate a snapshot

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Snapshot ID
	@return SnapshotsAPIDeactivateSnapshotRequest
*/
func (a *SnapshotsAPIService) DeactivateSnapshot(ctx context.Context, id string) SnapshotsAPIDeactivateSnapshotRequest {
	return SnapshotsAPIDeactivateSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *SnapshotsAPIService) DeactivateSnapshotExecute(r SnapshotsAPIDeactivateSnapshotRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.DeactivateSnapshot")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/snapshots/{id}/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SnapshotsAPIGetAllSnapshotsRequest struct {
	ctx                    context.Context
	ApiService             SnapshotsAPI
	xDaytonaOrganizationID *string
	limit                  *float32
	page                   *float32
}

// Use with JWT to specify the organization ID
func (r SnapshotsAPIGetAllSnapshotsRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) SnapshotsAPIGetAllSnapshotsRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

// Number of items per page
func (r SnapshotsAPIGetAllSnapshotsRequest) Limit(limit float32) SnapshotsAPIGetAllSnapshotsRequest {
	r.limit = &limit
	return r
}

// Page number
func (r SnapshotsAPIGetAllSnapshotsRequest) Page(page float32) SnapshotsAPIGetAllSnapshotsRequest {
	r.page = &page
	return r
}

func (r SnapshotsAPIGetAllSnapshotsRequest) Execute() (*PaginatedSnapshotsDto, *http.Response, error) {
	return r.ApiService.GetAllSnapshotsExecute(r)
}

/*
GetAllSnapshots List all snapshots

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SnapshotsAPIGetAllSnapshotsRequest
*/
func (a *SnapshotsAPIService) GetAllSnapshots(ctx context.Context) SnapshotsAPIGetAllSnapshotsRequest {
	return SnapshotsAPIGetAllSnapshotsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PaginatedSnapshotsDto
func (a *SnapshotsAPIService) GetAllSnapshotsExecute(r SnapshotsAPIGetAllSnapshotsRequest) (*PaginatedSnapshotsDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedSnapshotsDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.GetAllSnapshots")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/snapshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SnapshotsAPIGetSnapshotRequest struct {
	ctx                    context.Context
	ApiService             SnapshotsAPI
	id                     string
	xDaytonaOrganizationID *string
}

// Use with JWT to specify the organization ID
func (r SnapshotsAPIGetSnapshotRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) SnapshotsAPIGetSnapshotRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r SnapshotsAPIGetSnapshotRequest) Execute() (*SnapshotDto, *http.Response, error) {
	return r.ApiService.GetSnapshotExecute(r)
}

/*
GetSnapshot Get snapshot by ID or name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Snapshot ID or name
	@return SnapshotsAPIGetSnapshotRequest
*/
func (a *SnapshotsAPIService) GetSnapshot(ctx context.Context, id string) SnapshotsAPIGetSnapshotRequest {
	return SnapshotsAPIGetSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return SnapshotDto
func (a *SnapshotsAPIService) GetSnapshotExecute(r SnapshotsAPIGetSnapshotRequest) (*SnapshotDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SnapshotDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.GetSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/snapshots/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SnapshotsAPIGetSnapshotBuildLogsRequest struct {
	ctx                    context.Context
	ApiService             SnapshotsAPI
	id                     string
	xDaytonaOrganizationID *string
	follow                 *bool
}

// Use with JWT to specify the organization ID
func (r SnapshotsAPIGetSnapshotBuildLogsRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) SnapshotsAPIGetSnapshotBuildLogsRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

// Whether to follow the logs stream
func (r SnapshotsAPIGetSnapshotBuildLogsRequest) Follow(follow bool) SnapshotsAPIGetSnapshotBuildLogsRequest {
	r.follow = &follow
	return r
}

func (r SnapshotsAPIGetSnapshotBuildLogsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetSnapshotBuildLogsExecute(r)
}

/*
GetSnapshotBuildLogs Get snapshot build logs

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Snapshot ID
	@return SnapshotsAPIGetSnapshotBuildLogsRequest
*/
func (a *SnapshotsAPIService) GetSnapshotBuildLogs(ctx context.Context, id string) SnapshotsAPIGetSnapshotBuildLogsRequest {
	return SnapshotsAPIGetSnapshotBuildLogsRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *SnapshotsAPIService) GetSnapshotBuildLogsExecute(r SnapshotsAPIGetSnapshotBuildLogsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.GetSnapshotBuildLogs")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/snapshots/{id}/build-logs"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.follow != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "follow", r.follow, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SnapshotsAPIRemoveSnapshotRequest struct {
	ctx                    context.Context
	ApiService             SnapshotsAPI
	id                     string
	xDaytonaOrganizationID *string
}

// Use with JWT to specify the organization ID
func (r SnapshotsAPIRemoveSnapshotRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) SnapshotsAPIRemoveSnapshotRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r SnapshotsAPIRemoveSnapshotRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveSnapshotExecute(r)
}

/*
RemoveSnapshot Delete snapshot

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Snapshot ID
	@return SnapshotsAPIRemoveSnapshotRequest
*/
func (a *SnapshotsAPIService) RemoveSnapshot(ctx context.Context, id string) SnapshotsAPIRemoveSnapshotRequest {
	return SnapshotsAPIRemoveSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *SnapshotsAPIService) RemoveSnapshotExecute(r SnapshotsAPIRemoveSnapshotRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.RemoveSnapshot")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/snapshots/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SnapshotsAPISetSnapshotGeneralStatusRequest struct {
	ctx                         context.Context
	ApiService                  SnapshotsAPI
	id                          string
	setSnapshotGeneralStatusDto *SetSnapshotGeneralStatusDto
	xDaytonaOrganizationID      *string
}

func (r SnapshotsAPISetSnapshotGeneralStatusRequest) SetSnapshotGeneralStatusDto(setSnapshotGeneralStatusDto SetSnapshotGeneralStatusDto) SnapshotsAPISetSnapshotGeneralStatusRequest {
	r.setSnapshotGeneralStatusDto = &setSnapshotGeneralStatusDto
	return r
}

// Use with JWT to specify the organization ID
func (r SnapshotsAPISetSnapshotGeneralStatusRequest) XDaytonaOrganizationID(xDaytonaOrganizationID string) SnapshotsAPISetSnapshotGeneralStatusRequest {
	r.xDaytonaOrganizationID = &xDaytonaOrganizationID
	return r
}

func (r SnapshotsAPISetSnapshotGeneralStatusRequest) Execute() (*SnapshotDto, *http.Response, error) {
	return r.ApiService.SetSnapshotGeneralStatusExecute(r)
}

/*
SetSnapshotGeneralStatus Set snapshot general status

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Snapshot ID
	@return SnapshotsAPISetSnapshotGeneralStatusRequest
*/
func (a *SnapshotsAPIService) SetSnapshotGeneralStatus(ctx context.Context, id string) SnapshotsAPISetSnapshotGeneralStatusRequest {
	return SnapshotsAPISetSnapshotGeneralStatusRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return SnapshotDto
func (a *SnapshotsAPIService) SetSnapshotGeneralStatusExecute(r SnapshotsAPISetSnapshotGeneralStatusRequest) (*SnapshotDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SnapshotDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsAPIService.SetSnapshotGeneralStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/snapshots/{id}/general"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setSnapshotGeneralStatusDto == nil {
		return localVarReturnValue, nil, reportError("setSnapshotGeneralStatusDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDaytonaOrganizationID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Daytona-Organization-ID", r.xDaytonaOrganizationID, "simple", "")
	}
	// body params
	localVarPostBody = r.setSnapshotGeneralStatusDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

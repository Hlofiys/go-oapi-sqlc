// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.0 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Delete branch by branchId
	// (DELETE /api/branches)
	DeleteBranch(w http.ResponseWriter, r *http.Request)
	// Get all branches
	// (GET /api/branches)
	GetAllBranches(w http.ResponseWriter, r *http.Request)
	// Create a branch
	// (POST /api/branches)
	CreateBranch(w http.ResponseWriter, r *http.Request)
	// Check for branch max users limit
	// (GET /api/branches/limit)
	CheckBranchLimit(w http.ResponseWriter, r *http.Request, params CheckBranchLimitParams)
	// Get branch by branchId
	// (GET /api/branches/{branchId})
	GetBranchById(w http.ResponseWriter, r *http.Request, branchId int32)
	// Update branch by branchId
	// (PATCH /api/branches/{branchId})
	UpdateBranch(w http.ResponseWriter, r *http.Request, branchId int32)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Delete branch by branchId
// (DELETE /api/branches)
func (_ Unimplemented) DeleteBranch(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get all branches
// (GET /api/branches)
func (_ Unimplemented) GetAllBranches(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a branch
// (POST /api/branches)
func (_ Unimplemented) CreateBranch(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Check for branch max users limit
// (GET /api/branches/limit)
func (_ Unimplemented) CheckBranchLimit(w http.ResponseWriter, r *http.Request, params CheckBranchLimitParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get branch by branchId
// (GET /api/branches/{branchId})
func (_ Unimplemented) GetBranchById(w http.ResponseWriter, r *http.Request, branchId int32) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update branch by branchId
// (PATCH /api/branches/{branchId})
func (_ Unimplemented) UpdateBranch(w http.ResponseWriter, r *http.Request, branchId int32) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// DeleteBranch operation middleware
func (siw *ServerInterfaceWrapper) DeleteBranch(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{"role:Admin"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteBranch(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetAllBranches operation middleware
func (siw *ServerInterfaceWrapper) GetAllBranches(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAllBranches(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateBranch operation middleware
func (siw *ServerInterfaceWrapper) CreateBranch(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{"role:Admin"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateBranch(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CheckBranchLimit operation middleware
func (siw *ServerInterfaceWrapper) CheckBranchLimit(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params CheckBranchLimitParams

	// ------------- Required query parameter "usersAmount" -------------

	if paramValue := r.URL.Query().Get("usersAmount"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "usersAmount"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "usersAmount", r.URL.Query(), &params.UsersAmount)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "usersAmount", Err: err})
		return
	}

	// ------------- Optional query parameter "branchId" -------------

	err = runtime.BindQueryParameter("form", true, false, "branchId", r.URL.Query(), &params.BranchId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "branchId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CheckBranchLimit(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetBranchById operation middleware
func (siw *ServerInterfaceWrapper) GetBranchById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "branchId" -------------
	var branchId int32

	err = runtime.BindStyledParameterWithOptions("simple", "branchId", chi.URLParam(r, "branchId"), &branchId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "branchId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetBranchById(w, r, branchId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UpdateBranch operation middleware
func (siw *ServerInterfaceWrapper) UpdateBranch(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "branchId" -------------
	var branchId int32

	err = runtime.BindStyledParameterWithOptions("simple", "branchId", chi.URLParam(r, "branchId"), &branchId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "branchId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{"role:Admin"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateBranch(w, r, branchId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/branches", wrapper.DeleteBranch)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/branches", wrapper.GetAllBranches)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/branches", wrapper.CreateBranch)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/branches/limit", wrapper.CheckBranchLimit)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/branches/{branchId}", wrapper.GetBranchById)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/api/branches/{branchId}", wrapper.UpdateBranch)
	})

	return r
}

type DeleteBranchRequestObject struct {
	Body *DeleteBranchJSONRequestBody
}

type DeleteBranchResponseObject interface {
	VisitDeleteBranchResponse(w http.ResponseWriter) error
}

type DeleteBranch204Response struct {
}

func (response DeleteBranch204Response) VisitDeleteBranchResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type GetAllBranchesRequestObject struct {
}

type GetAllBranchesResponseObject interface {
	VisitGetAllBranchesResponse(w http.ResponseWriter) error
}

type GetAllBranches200JSONResponse []Branch

func (response GetAllBranches200JSONResponse) VisitGetAllBranchesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateBranchRequestObject struct {
	Body *CreateBranchJSONRequestBody
}

type CreateBranchResponseObject interface {
	VisitCreateBranchResponse(w http.ResponseWriter) error
}

type CreateBranch201JSONResponse Branch

func (response CreateBranch201JSONResponse) VisitCreateBranchResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CheckBranchLimitRequestObject struct {
	Params CheckBranchLimitParams
}

type CheckBranchLimitResponseObject interface {
	VisitCheckBranchLimitResponse(w http.ResponseWriter) error
}

type CheckBranchLimit200JSONResponse bool

func (response CheckBranchLimit200JSONResponse) VisitCheckBranchLimitResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetBranchByIdRequestObject struct {
	BranchId int32 `json:"branchId"`
}

type GetBranchByIdResponseObject interface {
	VisitGetBranchByIdResponse(w http.ResponseWriter) error
}

type GetBranchById200JSONResponse Branch

func (response GetBranchById200JSONResponse) VisitGetBranchByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateBranchRequestObject struct {
	BranchId int32 `json:"branchId"`
	Body     *UpdateBranchJSONRequestBody
}

type UpdateBranchResponseObject interface {
	VisitUpdateBranchResponse(w http.ResponseWriter) error
}

type UpdateBranch200JSONResponse Branch

func (response UpdateBranch200JSONResponse) VisitUpdateBranchResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Delete branch by branchId
	// (DELETE /api/branches)
	DeleteBranch(ctx context.Context, request DeleteBranchRequestObject) (DeleteBranchResponseObject, error)
	// Get all branches
	// (GET /api/branches)
	GetAllBranches(ctx context.Context, request GetAllBranchesRequestObject) (GetAllBranchesResponseObject, error)
	// Create a branch
	// (POST /api/branches)
	CreateBranch(ctx context.Context, request CreateBranchRequestObject) (CreateBranchResponseObject, error)
	// Check for branch max users limit
	// (GET /api/branches/limit)
	CheckBranchLimit(ctx context.Context, request CheckBranchLimitRequestObject) (CheckBranchLimitResponseObject, error)
	// Get branch by branchId
	// (GET /api/branches/{branchId})
	GetBranchById(ctx context.Context, request GetBranchByIdRequestObject) (GetBranchByIdResponseObject, error)
	// Update branch by branchId
	// (PATCH /api/branches/{branchId})
	UpdateBranch(ctx context.Context, request UpdateBranchRequestObject) (UpdateBranchResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// DeleteBranch operation middleware
func (sh *strictHandler) DeleteBranch(w http.ResponseWriter, r *http.Request) {
	var request DeleteBranchRequestObject

	var body DeleteBranchJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteBranch(ctx, request.(DeleteBranchRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteBranch")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteBranchResponseObject); ok {
		if err := validResponse.VisitDeleteBranchResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetAllBranches operation middleware
func (sh *strictHandler) GetAllBranches(w http.ResponseWriter, r *http.Request) {
	var request GetAllBranchesRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetAllBranches(ctx, request.(GetAllBranchesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAllBranches")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetAllBranchesResponseObject); ok {
		if err := validResponse.VisitGetAllBranchesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateBranch operation middleware
func (sh *strictHandler) CreateBranch(w http.ResponseWriter, r *http.Request) {
	var request CreateBranchRequestObject

	var body CreateBranchJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CreateBranch(ctx, request.(CreateBranchRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateBranch")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CreateBranchResponseObject); ok {
		if err := validResponse.VisitCreateBranchResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// CheckBranchLimit operation middleware
func (sh *strictHandler) CheckBranchLimit(w http.ResponseWriter, r *http.Request, params CheckBranchLimitParams) {
	var request CheckBranchLimitRequestObject

	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CheckBranchLimit(ctx, request.(CheckBranchLimitRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CheckBranchLimit")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CheckBranchLimitResponseObject); ok {
		if err := validResponse.VisitCheckBranchLimitResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetBranchById operation middleware
func (sh *strictHandler) GetBranchById(w http.ResponseWriter, r *http.Request, branchId int32) {
	var request GetBranchByIdRequestObject

	request.BranchId = branchId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetBranchById(ctx, request.(GetBranchByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetBranchById")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetBranchByIdResponseObject); ok {
		if err := validResponse.VisitGetBranchByIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdateBranch operation middleware
func (sh *strictHandler) UpdateBranch(w http.ResponseWriter, r *http.Request, branchId int32) {
	var request UpdateBranchRequestObject

	request.BranchId = branchId

	var body UpdateBranchJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateBranch(ctx, request.(UpdateBranchRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateBranch")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(UpdateBranchResponseObject); ok {
		if err := validResponse.VisitUpdateBranchResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

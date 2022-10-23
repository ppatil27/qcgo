// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"qcgo/models"
)

// UpdateResourcesByIDOKCode is the HTTP code returned for type UpdateResourcesByIDOK
const UpdateResourcesByIDOKCode int = 200

/*UpdateResourcesByIDOK OK

swagger:response updateResourcesByIdOK
*/
type UpdateResourcesByIDOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Resource `json:"body,omitempty"`
}

// NewUpdateResourcesByIDOK creates UpdateResourcesByIDOK with default headers values
func NewUpdateResourcesByIDOK() *UpdateResourcesByIDOK {

	return &UpdateResourcesByIDOK{}
}

// WithPayload adds the payload to the update resources by Id o k response
func (o *UpdateResourcesByIDOK) WithPayload(payload []*models.Resource) *UpdateResourcesByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update resources by Id o k response
func (o *UpdateResourcesByIDOK) SetPayload(payload []*models.Resource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateResourcesByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Resource, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateResourcesByIDBadRequestCode is the HTTP code returned for type UpdateResourcesByIDBadRequest
const UpdateResourcesByIDBadRequestCode int = 400

/*UpdateResourcesByIDBadRequest Bad Request

swagger:response updateResourcesByIdBadRequest
*/
type UpdateResourcesByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewUpdateResourcesByIDBadRequest creates UpdateResourcesByIDBadRequest with default headers values
func NewUpdateResourcesByIDBadRequest() *UpdateResourcesByIDBadRequest {

	return &UpdateResourcesByIDBadRequest{}
}

// WithPayload adds the payload to the update resources by Id bad request response
func (o *UpdateResourcesByIDBadRequest) WithPayload(payload *models.SimpleResponse) *UpdateResourcesByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update resources by Id bad request response
func (o *UpdateResourcesByIDBadRequest) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateResourcesByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateResourcesByIDUnauthorizedCode is the HTTP code returned for type UpdateResourcesByIDUnauthorized
const UpdateResourcesByIDUnauthorizedCode int = 401

/*UpdateResourcesByIDUnauthorized Unauthorized

swagger:response updateResourcesByIdUnauthorized
*/
type UpdateResourcesByIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewUpdateResourcesByIDUnauthorized creates UpdateResourcesByIDUnauthorized with default headers values
func NewUpdateResourcesByIDUnauthorized() *UpdateResourcesByIDUnauthorized {

	return &UpdateResourcesByIDUnauthorized{}
}

// WithPayload adds the payload to the update resources by Id unauthorized response
func (o *UpdateResourcesByIDUnauthorized) WithPayload(payload *models.SimpleResponse) *UpdateResourcesByIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update resources by Id unauthorized response
func (o *UpdateResourcesByIDUnauthorized) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateResourcesByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateResourcesByIDNotFoundCode is the HTTP code returned for type UpdateResourcesByIDNotFound
const UpdateResourcesByIDNotFoundCode int = 404

/*UpdateResourcesByIDNotFound Not Found

swagger:response updateResourcesByIdNotFound
*/
type UpdateResourcesByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewUpdateResourcesByIDNotFound creates UpdateResourcesByIDNotFound with default headers values
func NewUpdateResourcesByIDNotFound() *UpdateResourcesByIDNotFound {

	return &UpdateResourcesByIDNotFound{}
}

// WithPayload adds the payload to the update resources by Id not found response
func (o *UpdateResourcesByIDNotFound) WithPayload(payload *models.SimpleResponse) *UpdateResourcesByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update resources by Id not found response
func (o *UpdateResourcesByIDNotFound) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateResourcesByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateResourcesByIDUnprocessableEntityCode is the HTTP code returned for type UpdateResourcesByIDUnprocessableEntity
const UpdateResourcesByIDUnprocessableEntityCode int = 422

/*UpdateResourcesByIDUnprocessableEntity Unprocessable

swagger:response updateResourcesByIdUnprocessableEntity
*/
type UpdateResourcesByIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewUpdateResourcesByIDUnprocessableEntity creates UpdateResourcesByIDUnprocessableEntity with default headers values
func NewUpdateResourcesByIDUnprocessableEntity() *UpdateResourcesByIDUnprocessableEntity {

	return &UpdateResourcesByIDUnprocessableEntity{}
}

// WithPayload adds the payload to the update resources by Id unprocessable entity response
func (o *UpdateResourcesByIDUnprocessableEntity) WithPayload(payload *models.SimpleResponse) *UpdateResourcesByIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update resources by Id unprocessable entity response
func (o *UpdateResourcesByIDUnprocessableEntity) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateResourcesByIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateResourcesByIDInternalServerErrorCode is the HTTP code returned for type UpdateResourcesByIDInternalServerError
const UpdateResourcesByIDInternalServerErrorCode int = 500

/*UpdateResourcesByIDInternalServerError Internal Server Error

swagger:response updateResourcesByIdInternalServerError
*/
type UpdateResourcesByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewUpdateResourcesByIDInternalServerError creates UpdateResourcesByIDInternalServerError with default headers values
func NewUpdateResourcesByIDInternalServerError() *UpdateResourcesByIDInternalServerError {

	return &UpdateResourcesByIDInternalServerError{}
}

// WithPayload adds the payload to the update resources by Id internal server error response
func (o *UpdateResourcesByIDInternalServerError) WithPayload(payload *models.SimpleResponse) *UpdateResourcesByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update resources by Id internal server error response
func (o *UpdateResourcesByIDInternalServerError) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateResourcesByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateResourcesByIDServiceUnavailableCode is the HTTP code returned for type UpdateResourcesByIDServiceUnavailable
const UpdateResourcesByIDServiceUnavailableCode int = 503

/*UpdateResourcesByIDServiceUnavailable Service Unavailable

swagger:response updateResourcesByIdServiceUnavailable
*/
type UpdateResourcesByIDServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewUpdateResourcesByIDServiceUnavailable creates UpdateResourcesByIDServiceUnavailable with default headers values
func NewUpdateResourcesByIDServiceUnavailable() *UpdateResourcesByIDServiceUnavailable {

	return &UpdateResourcesByIDServiceUnavailable{}
}

// WithPayload adds the payload to the update resources by Id service unavailable response
func (o *UpdateResourcesByIDServiceUnavailable) WithPayload(payload *models.SimpleResponse) *UpdateResourcesByIDServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update resources by Id service unavailable response
func (o *UpdateResourcesByIDServiceUnavailable) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateResourcesByIDServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
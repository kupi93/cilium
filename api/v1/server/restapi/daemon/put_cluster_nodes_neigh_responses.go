// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// PutClusterNodesNeighCreatedCode is the HTTP code returned for type PutClusterNodesNeighCreated
const PutClusterNodesNeighCreatedCode int = 201

/*PutClusterNodesNeighCreated Created

swagger:response putClusterNodesNeighCreated
*/
type PutClusterNodesNeighCreated struct {
}

// NewPutClusterNodesNeighCreated creates PutClusterNodesNeighCreated with default headers values
func NewPutClusterNodesNeighCreated() *PutClusterNodesNeighCreated {

	return &PutClusterNodesNeighCreated{}
}

// WriteResponse to the client
func (o *PutClusterNodesNeighCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PutClusterNodesNeighInvalidCode is the HTTP code returned for type PutClusterNodesNeighInvalid
const PutClusterNodesNeighInvalidCode int = 400

/*PutClusterNodesNeighInvalid Invalid request; IPv6 is not supported

swagger:response putClusterNodesNeighInvalid
*/
type PutClusterNodesNeighInvalid struct {
}

// NewPutClusterNodesNeighInvalid creates PutClusterNodesNeighInvalid with default headers values
func NewPutClusterNodesNeighInvalid() *PutClusterNodesNeighInvalid {

	return &PutClusterNodesNeighInvalid{}
}

// WriteResponse to the client
func (o *PutClusterNodesNeighInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PutClusterNodesNeighFailureCode is the HTTP code returned for type PutClusterNodesNeighFailure
const PutClusterNodesNeighFailureCode int = 500

/*PutClusterNodesNeighFailure Error while inserting node neighbor

swagger:response putClusterNodesNeighFailure
*/
type PutClusterNodesNeighFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutClusterNodesNeighFailure creates PutClusterNodesNeighFailure with default headers values
func NewPutClusterNodesNeighFailure() *PutClusterNodesNeighFailure {

	return &PutClusterNodesNeighFailure{}
}

// WithPayload adds the payload to the put cluster nodes neigh failure response
func (o *PutClusterNodesNeighFailure) WithPayload(payload models.Error) *PutClusterNodesNeighFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put cluster nodes neigh failure response
func (o *PutClusterNodesNeighFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutClusterNodesNeighFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

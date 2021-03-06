// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewMoveEsClusterInstancesParams creates a new MoveEsClusterInstancesParams object
// with the default values initialized.
func NewMoveEsClusterInstancesParams() *MoveEsClusterInstancesParams {
	var (
		forceUpdateDefault   = bool(false)
		ignoreMissingDefault = bool(false)
		instancesDownDefault = bool(false)
		moveOnlyDefault      = bool(false)
		validateOnlyDefault  = bool(false)
	)
	return &MoveEsClusterInstancesParams{
		ForceUpdate:   &forceUpdateDefault,
		IgnoreMissing: &ignoreMissingDefault,
		InstancesDown: &instancesDownDefault,
		MoveOnly:      &moveOnlyDefault,
		ValidateOnly:  &validateOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMoveEsClusterInstancesParamsWithTimeout creates a new MoveEsClusterInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMoveEsClusterInstancesParamsWithTimeout(timeout time.Duration) *MoveEsClusterInstancesParams {
	var (
		forceUpdateDefault   = bool(false)
		ignoreMissingDefault = bool(false)
		instancesDownDefault = bool(false)
		moveOnlyDefault      = bool(false)
		validateOnlyDefault  = bool(false)
	)
	return &MoveEsClusterInstancesParams{
		ForceUpdate:   &forceUpdateDefault,
		IgnoreMissing: &ignoreMissingDefault,
		InstancesDown: &instancesDownDefault,
		MoveOnly:      &moveOnlyDefault,
		ValidateOnly:  &validateOnlyDefault,

		timeout: timeout,
	}
}

// NewMoveEsClusterInstancesParamsWithContext creates a new MoveEsClusterInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewMoveEsClusterInstancesParamsWithContext(ctx context.Context) *MoveEsClusterInstancesParams {
	var (
		forceUpdateDefault   = bool(false)
		ignoreMissingDefault = bool(false)
		instancesDownDefault = bool(false)
		moveOnlyDefault      = bool(false)
		validateOnlyDefault  = bool(false)
	)
	return &MoveEsClusterInstancesParams{
		ForceUpdate:   &forceUpdateDefault,
		IgnoreMissing: &ignoreMissingDefault,
		InstancesDown: &instancesDownDefault,
		MoveOnly:      &moveOnlyDefault,
		ValidateOnly:  &validateOnlyDefault,

		Context: ctx,
	}
}

// NewMoveEsClusterInstancesParamsWithHTTPClient creates a new MoveEsClusterInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMoveEsClusterInstancesParamsWithHTTPClient(client *http.Client) *MoveEsClusterInstancesParams {
	var (
		forceUpdateDefault   = bool(false)
		ignoreMissingDefault = bool(false)
		instancesDownDefault = bool(false)
		moveOnlyDefault      = bool(false)
		validateOnlyDefault  = bool(false)
	)
	return &MoveEsClusterInstancesParams{
		ForceUpdate:   &forceUpdateDefault,
		IgnoreMissing: &ignoreMissingDefault,
		InstancesDown: &instancesDownDefault,
		MoveOnly:      &moveOnlyDefault,
		ValidateOnly:  &validateOnlyDefault,
		HTTPClient:    client,
	}
}

/*MoveEsClusterInstancesParams contains all the parameters to send to the API endpoint
for the move es cluster instances operation typically these are written to a http.Request
*/
type MoveEsClusterInstancesParams struct {

	/*Body
	  Overrides defaults for the move, including setting the configuration of instances specified in the path

	*/
	Body *models.TransientElasticsearchPlanConfiguration
	/*ClusterID
	  The Elasticsearch cluster identifier.

	*/
	ClusterID string
	/*ForceUpdate
	  When `true`, cancels and overwrites the pending plans, or treats the instance as an error.

	*/
	ForceUpdate *bool
	/*IgnoreMissing
	  When `true` and the instance does not exist, proceeds to the next instance, or treats the instance as an error.

	*/
	IgnoreMissing *bool
	/*InstanceIds
	  A comma-separated list of instance identifiers.

	*/
	InstanceIds []string
	/*InstancesDown
	  When `true`, the instances specified by `instance_ids` permanently shut down for data migration logic.

	*/
	InstancesDown *bool
	/*MoveOnly
	  When `true`, moves the specified instances and ignores the changes for the cluster state.

	*/
	MoveOnly *bool
	/*ValidateOnly
	  When `true`, validates the move request, then returns the calculated plan without applying the plan.

	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithTimeout(timeout time.Duration) *MoveEsClusterInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithContext(ctx context.Context) *MoveEsClusterInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithHTTPClient(client *http.Client) *MoveEsClusterInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithBody(body *models.TransientElasticsearchPlanConfiguration) *MoveEsClusterInstancesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetBody(body *models.TransientElasticsearchPlanConfiguration) {
	o.Body = body
}

// WithClusterID adds the clusterID to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithClusterID(clusterID string) *MoveEsClusterInstancesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithForceUpdate adds the forceUpdate to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithForceUpdate(forceUpdate *bool) *MoveEsClusterInstancesParams {
	o.SetForceUpdate(forceUpdate)
	return o
}

// SetForceUpdate adds the forceUpdate to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetForceUpdate(forceUpdate *bool) {
	o.ForceUpdate = forceUpdate
}

// WithIgnoreMissing adds the ignoreMissing to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithIgnoreMissing(ignoreMissing *bool) *MoveEsClusterInstancesParams {
	o.SetIgnoreMissing(ignoreMissing)
	return o
}

// SetIgnoreMissing adds the ignoreMissing to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetIgnoreMissing(ignoreMissing *bool) {
	o.IgnoreMissing = ignoreMissing
}

// WithInstanceIds adds the instanceIds to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithInstanceIds(instanceIds []string) *MoveEsClusterInstancesParams {
	o.SetInstanceIds(instanceIds)
	return o
}

// SetInstanceIds adds the instanceIds to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetInstanceIds(instanceIds []string) {
	o.InstanceIds = instanceIds
}

// WithInstancesDown adds the instancesDown to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithInstancesDown(instancesDown *bool) *MoveEsClusterInstancesParams {
	o.SetInstancesDown(instancesDown)
	return o
}

// SetInstancesDown adds the instancesDown to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetInstancesDown(instancesDown *bool) {
	o.InstancesDown = instancesDown
}

// WithMoveOnly adds the moveOnly to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithMoveOnly(moveOnly *bool) *MoveEsClusterInstancesParams {
	o.SetMoveOnly(moveOnly)
	return o
}

// SetMoveOnly adds the moveOnly to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetMoveOnly(moveOnly *bool) {
	o.MoveOnly = moveOnly
}

// WithValidateOnly adds the validateOnly to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) WithValidateOnly(validateOnly *bool) *MoveEsClusterInstancesParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the move es cluster instances params
func (o *MoveEsClusterInstancesParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *MoveEsClusterInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.ForceUpdate != nil {

		// query param force_update
		var qrForceUpdate bool
		if o.ForceUpdate != nil {
			qrForceUpdate = *o.ForceUpdate
		}
		qForceUpdate := swag.FormatBool(qrForceUpdate)
		if qForceUpdate != "" {
			if err := r.SetQueryParam("force_update", qForceUpdate); err != nil {
				return err
			}
		}

	}

	if o.IgnoreMissing != nil {

		// query param ignore_missing
		var qrIgnoreMissing bool
		if o.IgnoreMissing != nil {
			qrIgnoreMissing = *o.IgnoreMissing
		}
		qIgnoreMissing := swag.FormatBool(qrIgnoreMissing)
		if qIgnoreMissing != "" {
			if err := r.SetQueryParam("ignore_missing", qIgnoreMissing); err != nil {
				return err
			}
		}

	}

	valuesInstanceIds := o.InstanceIds

	joinedInstanceIds := swag.JoinByFormat(valuesInstanceIds, "csv")
	// path array param instance_ids
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedInstanceIds) > 0 {
		if err := r.SetPathParam("instance_ids", joinedInstanceIds[0]); err != nil {
			return err
		}
	}

	if o.InstancesDown != nil {

		// query param instances_down
		var qrInstancesDown bool
		if o.InstancesDown != nil {
			qrInstancesDown = *o.InstancesDown
		}
		qInstancesDown := swag.FormatBool(qrInstancesDown)
		if qInstancesDown != "" {
			if err := r.SetQueryParam("instances_down", qInstancesDown); err != nil {
				return err
			}
		}

	}

	if o.MoveOnly != nil {

		// query param move_only
		var qrMoveOnly bool
		if o.MoveOnly != nil {
			qrMoveOnly = *o.MoveOnly
		}
		qMoveOnly := swag.FormatBool(qrMoveOnly)
		if qMoveOnly != "" {
			if err := r.SetQueryParam("move_only", qMoveOnly); err != nil {
				return err
			}
		}

	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool
		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {
			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

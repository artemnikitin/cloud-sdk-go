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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActiveDirectorySecurityRealmRoleMappingRule Elasticsearch Security Active Directory role mapping rule
// swagger:model ActiveDirectorySecurityRealmRoleMappingRule
type ActiveDirectorySecurityRealmRoleMappingRule struct {

	// The roles that are applied when the mapping rule is successfully evaluated
	// Required: true
	Roles []string `json:"roles"`

	// The type of role mapping rule
	// Required: true
	// Enum: [user_dn group_dn]
	Type *string `json:"type"`

	// The value to match when evaluating this rule
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this active directory security realm role mapping rule
func (m *ActiveDirectorySecurityRealmRoleMappingRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectorySecurityRealmRoleMappingRule) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

var activeDirectorySecurityRealmRoleMappingRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user_dn","group_dn"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activeDirectorySecurityRealmRoleMappingRuleTypeTypePropEnum = append(activeDirectorySecurityRealmRoleMappingRuleTypeTypePropEnum, v)
	}
}

const (

	// ActiveDirectorySecurityRealmRoleMappingRuleTypeUserDn captures enum value "user_dn"
	ActiveDirectorySecurityRealmRoleMappingRuleTypeUserDn string = "user_dn"

	// ActiveDirectorySecurityRealmRoleMappingRuleTypeGroupDn captures enum value "group_dn"
	ActiveDirectorySecurityRealmRoleMappingRuleTypeGroupDn string = "group_dn"
)

// prop value enum
func (m *ActiveDirectorySecurityRealmRoleMappingRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, activeDirectorySecurityRealmRoleMappingRuleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ActiveDirectorySecurityRealmRoleMappingRule) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectorySecurityRealmRoleMappingRule) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActiveDirectorySecurityRealmRoleMappingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectorySecurityRealmRoleMappingRule) UnmarshalBinary(b []byte) error {
	var res ActiveDirectorySecurityRealmRoleMappingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

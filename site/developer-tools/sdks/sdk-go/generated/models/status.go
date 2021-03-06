// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Status status
// swagger:model Status
type Status struct {

	// difficulty
	// Required: true
	Difficulty *uint64 `json:"difficulty"`

	// genesis key block hash
	// Required: true
	GenesisKeyBlockHash EncodedHash `json:"genesis_key_block_hash"`

	// listening
	// Required: true
	Listening *bool `json:"listening"`

	// network id
	// Required: true
	NetworkID *string `json:"network_id"`

	// node revision
	// Required: true
	NodeRevision *string `json:"node_revision"`

	// node version
	// Required: true
	NodeVersion *string `json:"node_version"`

	// peer count
	// Required: true
	// Minimum: 0
	PeerCount *uint64 `json:"peer_count"`

	// pending transactions count
	// Required: true
	// Minimum: 0
	PendingTransactionsCount *uint64 `json:"pending_transactions_count"`

	// protocols
	// Required: true
	Protocols []*Protocol `json:"protocols"`

	// solutions
	// Required: true
	// Minimum: 0
	Solutions *uint64 `json:"solutions"`

	// syncing
	// Required: true
	Syncing *bool `json:"syncing"`
}

// Validate validates this status
func (m *Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDifficulty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenesisKeyBlockHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListening(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePendingTransactionsCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocols(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSolutions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) validateDifficulty(formats strfmt.Registry) error {

	if err := validate.Required("difficulty", "body", m.Difficulty); err != nil {
		return err
	}

	return nil
}

func (m *Status) validateGenesisKeyBlockHash(formats strfmt.Registry) error {

	if err := m.GenesisKeyBlockHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("genesis_key_block_hash")
		}
		return err
	}

	return nil
}

func (m *Status) validateListening(formats strfmt.Registry) error {

	if err := validate.Required("listening", "body", m.Listening); err != nil {
		return err
	}

	return nil
}

func (m *Status) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("network_id", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

func (m *Status) validateNodeRevision(formats strfmt.Registry) error {

	if err := validate.Required("node_revision", "body", m.NodeRevision); err != nil {
		return err
	}

	return nil
}

func (m *Status) validateNodeVersion(formats strfmt.Registry) error {

	if err := validate.Required("node_version", "body", m.NodeVersion); err != nil {
		return err
	}

	return nil
}

func (m *Status) validatePeerCount(formats strfmt.Registry) error {

	if err := validate.Required("peer_count", "body", m.PeerCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("peer_count", "body", int64(*m.PeerCount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Status) validatePendingTransactionsCount(formats strfmt.Registry) error {

	if err := validate.Required("pending_transactions_count", "body", m.PendingTransactionsCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("pending_transactions_count", "body", int64(*m.PendingTransactionsCount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Status) validateProtocols(formats strfmt.Registry) error {

	if err := validate.Required("protocols", "body", m.Protocols); err != nil {
		return err
	}

	for i := 0; i < len(m.Protocols); i++ {
		if swag.IsZero(m.Protocols[i]) { // not required
			continue
		}

		if m.Protocols[i] != nil {
			if err := m.Protocols[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protocols" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Status) validateSolutions(formats strfmt.Registry) error {

	if err := validate.Required("solutions", "body", m.Solutions); err != nil {
		return err
	}

	if err := validate.MinimumInt("solutions", "body", int64(*m.Solutions), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Status) validateSyncing(formats strfmt.Registry) error {

	if err := validate.Required("syncing", "body", m.Syncing); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Status) UnmarshalBinary(b []byte) error {
	var res Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

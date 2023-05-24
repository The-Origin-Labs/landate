package models

import "gorm.io/gorm"

// Property represents a real estate
// property with its ownership information,
// property details, and other relevant attributes.
type Property struct {
	gorm.Model
	// OnwershipInfo: contains information
	// about the owner of the property, including
	// the owner's name, public address, ownership
	// type, and percentage ownership.
	PropertyOwner       string  `json:"property_owner"`
	PublicAddress       string  `json:"public_address"`
	OwernershipType     string  `json:"owernership_type"`
	PercentageOwnership float64 `json:"percentage_ownership"`

	// PropertyDetails: contains detailed
	// information about the property, such as its location,
	// area, landmark, and zone restriction.
	AreaPincode     int     `json:"area_pincode"`
	District        string  `json:"district"`
	State           string  `json:"state"`
	Country         string  `json:"country"`
	Area            float64 `json:"area"`
	Landmark        string  `json:"landmark"`
	ZoneRestriction string  `json:"zone_restriction"`

	// Transparency: describes the
	// transparency level of the property.
	Transparency string `json:"transparency"`

	// TransferMechanism: refers to
	// the method of transferring property ownership.
	TransferMechanism string `json:"transfer_mechanism"`

	// PaymentProcessing: describes the
	// payment processing method for property transactions.
	PaymentProcessing string `json:"payment_processing"`

	// RegulatoryCompliance represents
	// the property's compliance with relevant regulations.
	RegulatoryCompliance string `json:"regulatory_compliance"`

	// AccessControl: contains information
	// about the access control of the property,
	// including the number of owners.
	NumberOfOnwners int `json:"number_of_onwners"`

	// DisputeResolution: describes the
	// dispute resolution process for the property.
	DisputeResolution string `json:"dispute_resolution"`

	// CurrentPropertyPrice: is the
	// current market price of the property.
	CurrentPropertyPrice int `json:"current_property_price"`

	// ActualPropertyPrice: is the actual price of the property.
	ActualPropertyPrice int `json:"actual_property_price"`
}

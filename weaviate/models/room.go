package models


// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Room room
// swagger:model Room
type Room struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// place Id
	PlaceID string `json:"placeId,omitempty"`

	// room type
	RoomType string `json:"roomType,omitempty"`
}

// Validate validates this room
func (m *Room) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoomType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var roomTypeRoomTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["attic","backyard","basement","bathroom","bedroom","default","den","diningRoom","entryway","familyRoom","frontyard","garage","hallway","kitchen","livingRoom","masterBedroom","office","shed","unknownRoomType"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roomTypeRoomTypePropEnum = append(roomTypeRoomTypePropEnum, v)
	}
}

const (
	roomRoomTypeAttic           string = "attic"
	roomRoomTypeBackyard        string = "backyard"
	roomRoomTypeBasement        string = "basement"
	roomRoomTypeBathroom        string = "bathroom"
	roomRoomTypeBedroom         string = "bedroom"
	roomRoomTypeDefault         string = "default"
	roomRoomTypeDen             string = "den"
	roomRoomTypeDiningRoom      string = "diningRoom"
	roomRoomTypeEntryway        string = "entryway"
	roomRoomTypeFamilyRoom      string = "familyRoom"
	roomRoomTypeFrontyard       string = "frontyard"
	roomRoomTypeGarage          string = "garage"
	roomRoomTypeHallway         string = "hallway"
	roomRoomTypeKitchen         string = "kitchen"
	roomRoomTypeLivingRoom      string = "livingRoom"
	roomRoomTypeMasterBedroom   string = "masterBedroom"
	roomRoomTypeOffice          string = "office"
	roomRoomTypeShed            string = "shed"
	roomRoomTypeUnknownRoomType string = "unknownRoomType"
)

// prop value enum
func (m *Room) validateRoomTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, roomTypeRoomTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Room) validateRoomType(formats strfmt.Registry) error {

	if swag.IsZero(m.RoomType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoomTypeEnum("roomType", "body", m.RoomType); err != nil {
		return err
	}

	return nil
}

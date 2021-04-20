package property

import "github.com/JoseFMP/resharmonics/contracts"

type PropertyData struct {
	ID                  PropertyID `json:"id"`
	Name                string     `json:"name"`
	BuildingID          BuildingID `json:"buildingId"`
	BuildingName        string     `json:"buildingName"`
	BuildingDescription string     `json:"buildingDescription"`

	UnitTypeId int `json:"unitTypeId"`

	Unavailable                 bool   `json:"unavailable"`
	PropertyNameName            string `json:"propertyName"`
	PropertyTypeDescription     string `json:"propertyTpeDescription"`
	PropertyTypeLongDescription string `json:"propertyTypeLongDescription"`
	ShortDescription            string `json:"shortDescription"`
	MaxOccupancy                int    `json:"maxOccupancy"`
	FloorSpace                  string `json:"floorSpace"`
	Location                    string `json:"location"`

	Contracts    []contracts.Contract `json:"unitContracts"`
	CustomFields []CustomField        `json:"customFields"`
}

type PropertyID string
type BuildingID int

type CustomField struct {
	Name  string          `json:"name"`
	Type  CustomFieldType `json:"type"`
	Value string          `json:"value"`
}

type CustomFieldType string

package property

type PropertyData struct {
	ID         PropertyID `json:"id"`
	Name       string     `json:"name"`
	BuildingID BuildingID `json:"buildingId"`
}

type PropertyID string
type BuildingID int

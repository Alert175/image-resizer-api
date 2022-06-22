package dto

// Структура запроса загородных домов
type ReqFilterHousesDto struct {
	DealStatus        string         `json:"dealStatus"`
	Location          ReqLocationDto `json:"location"`
	LotType           string         `json:"lotType"`
	Renovation        string         `json:"renovation"`
	Rooms             []int          `json:"rooms"`
	Floors            []int          `json:"floors"`
	BuildingType      string         `json:"buildingType"`
	Price             []int          `json:"price"`
	AllArea           []int          `json:"allArea"`
	LivingArea        []int          `json:"livingArea"`
	KitchenArea       []int          `json:"kitchen"`
	LotArea           []int          `json:"lotArea"`
	Lift              bool           `json:"lift,omitempty"`
	ElectricitySupply bool           `json:"electricitySupply"`
	WaterSupply       bool           `json:"waterSupply"`
	GasSupply         bool           `json:"gasSupply"`
	SewerageSupply    bool           `json:"sewerageSupply"`
	HeatingSupply     bool           `json:"heatingSupply"`
	Apartments        bool           `json:"apartments"`
	Mortgage          bool           `json:"mortgage"`
	Order             ReqOrderDto    `json:"order"`
	Config            ConfigDto      `json:"config"`
}

type ReqLocationDto struct {
	Country   string    `json:"country"`
	Region    string    `json:"region"`
	Name      string    `json:"name"`
	District  string    `json:"district"`
	Longitude []float32 `json:"longitude"`
	Latitude  []float32 `json:"latitude"`
}

type ReqOrderDto struct {
	Field string `json:"field" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type ConfigDto struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

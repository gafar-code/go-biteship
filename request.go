package biteship

type Metadata interface{}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ProductItem struct {
	Id          interface{} `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Value       uint        `json:"value"`
	Quantity    uint        `json:"quantity"`
	Height      uint        `json:"height"`
	Width       uint        `json:"width"`
	Length      uint        `json:"length"`
	Weight      uint        `json:"weight"`
}

type CreateOrderRequestParam struct {
	//	SENDER SHIPPER DATA
	ShipperContactName  string `json:"shipper_contact_name"`
	ShipperContactPhone string `json:"shipper_contact_phone"`
	ShipperContactEmail string `json:"shipper_contact_email"`
	ShipperOrganization string `json:"shipper_organization"`

	//	SENDER ORIGIN DATA
	OriginContactName  string     `json:"origin_contact_name"`
	OriginContactPhone string     `json:"origin_contact_phone"`
	OriginAddress      string     `json:"origin_address"`
	OriginNote         string     `json:"origin_note"`
	OriginPostalCode   uint       `json:"origin_postal_code"`
	OriginCoordinate   Coordinate `json:"origin_coordinate"`

	//	DESTINATION CONTACT NAME
	DestinationContactName        string     `json:"destination_contact_name"`
	DestinationContactPhone       string     `json:"destination_contact_phone"`
	DestinationContactEmail       string     `json:"destination_contact_email"`
	DestinationAddress            string     `json:"destination_address"`
	DestinationPostalCode         uint       `json:"destination_postal_code"`
	DestinationNote               string     `json:"destination_note"`
	DestinationCoordinate         Coordinate `json:"destination_coordinate"`
	DestinationCashOnDelivery     *uint      `json:"destination_cash_on_delivery"` // Optional
	DestinationCashOnDeliveryType *string    `json:"destination_cash_on_delivery_type"`

	//	COURIER DATA
	CourierCompany   string `json:"courier_company"`
	CourierType      string `json:"courier_type"`
	CourierInsurance uint   `json:"courier_insurance"`

	//	DELIVERY TIME DATA
	DeliveryType string `json:"delivery_type"` // "later" or "now"
	DeliveryDate string `json:"delivery_date"` // yyyy-mm-dd
	DeliveryTime string `json:"delivery_time"` // hh:mm

	PaymentType string   `json:"payment_type"` // Set to be 'online'
	OrderNote   string   `json:"order_note"`
	Metadata    Metadata `json:"metadata"` // Optional

	Items []ProductItem `json:"items"`
}

type ItemCourierRate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       uint   `json:"value"`
	Length      uint   `json:"length"`
	Width       uint   `json:"width"`
	Height      uint   `json:"height"`
	Weight      uint   `json:"weight"`
	Quantity    uint   `json:"quantity"`
}
type RequestCourierRates struct {
	OriginLatitude        float64           `json:"origin_latitude" binding:"required"`
	OriginLongitude       float64           `json:"origin_longitude" binding:"required"`
	DestinationLatitude   float64           `json:"destination_latitude" binding:"required"`
	DestinationLongitude  float64           `json:"destination_longitude" binding:"required"`
	OriginPostalCode      uint              `json:"origin_postal_code"`
	DestinationPostalCode uint              `json:"destination_postal_code"`
	Couriers              string            `json:"couriers" binding:"required"` // "grab, gojek, jne, tiki, jet, jnt, sicepat, wahana, pos, lion, ninja, anteraja, rpx, paxel, mrspeedy, lalamove, deliveree, sap"
	Items                 []ItemCourierRate `json:"items"`
}

type ReasonRequest struct {
	CancellationReason string `json:"cancellation_reason"`
}

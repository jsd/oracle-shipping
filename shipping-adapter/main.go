package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/linkpoolio/bridges"
)

type EAShipping struct {
}

type ShippingTracker struct {
	Data []struct {
		Archived        bool      `json:"archived"`
		CarrierCode     string    `json:"carrier_code"`
		Comment         string    `json:"comment"`
		CreatedAt       time.Time `json:"created_at"`
		CustomerEmail   []string  `json:"customer_email"`
		CustomerName    string    `json:"customer_name"`
		DestinationInfo struct {
			ArrivalfromAbroad  interface{} `json:"ArrivalfromAbroad"`
			CustomsClearance   interface{} `json:"CustomsClearance"`
			DepartfromAirport  interface{} `json:"DepartfromAirport"`
			DestinationArrived interface{} `json:"DestinationArrived"`
			ItemDispatched     interface{} `json:"ItemDispatched"`
			ItemReceived       interface{} `json:"ItemReceived"`
			CarrierCode        interface{} `json:"carrier_code"`
			Phone              interface{} `json:"phone"`
			Trackinfo          interface{} `json:"trackinfo"`
			Weblink            interface{} `json:"weblink"`
		} `json:"destination_info"`
		ID              string `json:"id"`
		ItemTimeLength  int    `json:"itemTimeLength"`
		LastEvent       string `json:"lastEvent"`
		LastUpdateTime  string `json:"lastUpdateTime"`
		OrderCreateTime string `json:"order_create_time"`
		OrderID         string `json:"order_id"`
		OriginInfo      struct {
			ArrivalfromAbroad  interface{} `json:"ArrivalfromAbroad"`
			CustomsClearance   interface{} `json:"CustomsClearance"`
			DepartfromAirport  string      `json:"DepartfromAirport"`
			DestinationArrived interface{} `json:"DestinationArrived"`
			ItemDispatched     interface{} `json:"ItemDispatched"`
			ItemReceived       string      `json:"ItemReceived"`
			ReferenceNumber    interface{} `json:"ReferenceNumber"`
			CarrierCode        string      `json:"carrier_code"`
			Phone              interface{} `json:"phone"`
			Trackinfo          []struct {
				Date              string `json:"Date"`
				Details           string `json:"Details"`
				StatusDescription string `json:"StatusDescription"`
				CheckpointStatus  string `json:"checkpoint_status"`
				Substatus         string `json:"substatus"`
				ItemNode          string `json:"ItemNode,omitempty"`
			} `json:"trackinfo"`
			Weblink string `json:"weblink"`
		} `json:"origin_info"`
		OriginalCountry string      `json:"original_country"`
		ServiceCode     interface{} `json:"service_code"`
		Status          string      `json:"status"`
		StatusInfo      interface{} `json:"status_info"`
		StayTimeLength  int         `json:"stayTimeLength"`
		Title           string      `json:"title"`
		TrackingNumber  string      `json:"tracking_number"`
		UpdatedAt       time.Time   `json:"updated_at"`
	} `json:"data"`
}

// Run implements Bridge Run for querying the ETH Gas Station
func (eas *EAShipping) Run(h *bridges.Helper) (interface{}, error) {
	object := ShippingTracker{}
	param := h.GetParam("id")
	url := fmt.Sprintf("https://api.tracktry.com/v1/trackings/4px/%v", param)
	err := h.HTTPCallWithOpts(
		http.MethodGet,
		url,
		&object,
		bridges.CallOpts{
			Auth: bridges.NewAuth(bridges.AuthHeader, "Tracktry-Api-Key", "APIKEY"),
		},
	)
	return object.Data[0], err
}

// Opts is the bridge.Bridge implementation
func (eas *EAShipping) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "EAShipping",
		Lambda: true,
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	bridges.NewServer(&EAShipping{}).Handler(w, r)
}

func main() {
	bridges.NewServer(&EAShipping{}).Start(8080)
}

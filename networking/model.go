package networking

type Api struct {
	calculateBody CalculateBody // manage the calculate API endpoint body
	response      Response      // manage the universal response for use in all API response
}

type CalculateBody struct {
	A *uint `json:"a,omitempty"`
	B *uint `json:"b,omitempty"`
}

type Response struct {
	Error   string `json:"error,omitempty"`
	Product uint   `json:"product,omitempty"`
}

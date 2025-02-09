package controllers

// Car struct for request body
type carVariantRequest struct {
	CarId        uint
	VariantCode  string
	VariantName  string
	Transmission string
	Color        string
	Engine       string
}

// Defining struct for response
// type carVariantResponse struct {
// 	carVariantRequest
// 	ID uint
// }

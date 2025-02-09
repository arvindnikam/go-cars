package controllers

// Car struct for request body
type carRequest struct {
	Make     string
	CarModel string
	Year     int
	BodyType string
}

// Defining struct for response
// type carResponse struct {
// 	carRequest
// 	ID uint
// }

type SearchRequest struct {
	Conditions map[string]map[string]interface{}
	Page       int
	Limit      int
	Offset     int
	SortColumn string
	SortOrder  string
}

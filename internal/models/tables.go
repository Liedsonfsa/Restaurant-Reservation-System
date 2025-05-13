package models

type Tables struct {
	ID 			uint64  `json:"id"`
	Name		string	`json:"name"`
	Capacity	uint64 	`json:"capacity"`
	Status 		string	`json:"status"`
}
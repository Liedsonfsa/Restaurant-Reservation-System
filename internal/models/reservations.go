package models

import "time"

type Reservation struct {
	ID              	uint64 		`json:"id"`
	UserID          	uint64 		`json:"user_id"`
	TableID         	uint64 		`json:"table_id"`
	DateReservation 	time.Time	`json:"date_reservation"`
	Status				string		`json:"status"`
}
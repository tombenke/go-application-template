package domain

import "github.com/google/uuid"

type Contact struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

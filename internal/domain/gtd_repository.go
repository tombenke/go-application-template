package domain

import (
	"context"

	"github.com/google/uuid"
)

// GTDRepository defines the interface for managing contacts in the application.
type GTDRepository interface {
	CreateContact(ctx context.Context, contact *Contact) error
	GetContactByID(ctx context.Context, id uuid.UUID) (*Contact, error)
	UpdateContact(ctx context.Context, contact *Contact) error
	DeleteContact(ctx context.Context, id uuid.UUID) error
	ListContacts(ctx context.Context) ([]*Contact, error)
}

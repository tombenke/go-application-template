package inmemory

import (
	"context"

	"github.com/google/uuid"
	"github.com/tombenke/go-application-template/internal/domain"
)

// No real mapper is needed for this simple in-memory implementation,
// but we define one for the sake of demonstration if you want to separate the domain models from the database models.
type ContactDbEntry domain.Contact

// GTDRepository defines the interface for managing contacts in the application.
type GTDRepository struct {
	contacts map[uuid.UUID]ContactDbEntry
}

func NewGTDRepository() *GTDRepository {
	return &GTDRepository{
		contacts: make(map[uuid.UUID]ContactDbEntry),
	}
}

func (r *GTDRepository) CreateContact(_ context.Context, contact *domain.Contact) error {
	newContactEntry := r.mapToDbEntry(contact)
	r.contacts[newContactEntry.ID] = newContactEntry
	return nil
}

func (r *GTDRepository) GetContactByID(_ context.Context, id uuid.UUID) (*domain.Contact, error) {
	if contactDbEntry, exists := r.contacts[id]; exists {
		return r.mapToDomain(contactDbEntry), nil
	}
	return &domain.Contact{}, nil
}

func (r *GTDRepository) UpdateContact(_ context.Context, contact *domain.Contact) error {
	if _, exists := r.contacts[contact.ID]; exists {
		r.contacts[contact.ID] = r.mapToDbEntry(contact)
		return nil
	}
	return nil
}

func (r *GTDRepository) DeleteContact(_ context.Context, id uuid.UUID) error {
	if _, exists := r.contacts[id]; exists {
		delete(r.contacts, id)
		return nil
	}
	return nil
}

func (r *GTDRepository) ListContacts(_ context.Context) ([]*domain.Contact, error) {
	contacts := make([]*domain.Contact, 0, len(r.contacts))
	for _, contactDbEntry := range r.contacts {
		contacts = append(contacts, r.mapToDomain(contactDbEntry))
	}
	return contacts, nil
}

func (*GTDRepository) mapToDbEntry(contact *domain.Contact) ContactDbEntry {
	return ContactDbEntry(*contact)
}

func (*GTDRepository) mapToDomain(contactDbEntry ContactDbEntry) *domain.Contact {
	contact := domain.Contact(contactDbEntry)
	return &contact
}

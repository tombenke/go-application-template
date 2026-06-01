package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/tombenke/go-application-template/internal/domain"
)

type GTDManager interface {
	GetContacts(ctx context.Context) (*ContactsDTO, error)
	AddContact(ctx context.Context, contact ContactDTO) error
	UpdateContact(ctx context.Context, contact ContactDTO) error
	DeleteContact(ctx context.Context, id uuid.UUID) error
}

type gtdManagerImpl struct {
	// add dependencies here, e.g. a repository for managing contacts
	gtdRepository domain.GTDRepository
}

func NewGTDManager(gtdRepository domain.GTDRepository) GTDManager {
	return &gtdManagerImpl{
		gtdRepository: gtdRepository,
	}
}

func (c *gtdManagerImpl) GetContacts(ctx context.Context) (*ContactsDTO, error) {
	// Implement the logic to retrieve contacts, e.g. from a database or in-memory store
	contacts, err := c.gtdRepository.ListContacts(ctx)
	if err != nil {
		return nil, err
	}
	contactsDTO := ToContactsDTO(contacts)
	return &contactsDTO, nil
}

func (c *gtdManagerImpl) AddContact(ctx context.Context, contact ContactDTO) error {
	newContact := domain.Contact{
		ID:        contact.ID,
		Email:     contact.Email,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
	}
	return c.gtdRepository.CreateContact(ctx, &newContact)
}

func (c *gtdManagerImpl) UpdateContact(ctx context.Context, contact ContactDTO) error {
	updatedContact := domain.Contact{
		ID:        contact.ID,
		Email:     contact.Email,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
	}
	return c.gtdRepository.UpdateContact(ctx, &updatedContact)
}

func (c *gtdManagerImpl) DeleteContact(ctx context.Context, contactID uuid.UUID) error {

	return c.gtdRepository.DeleteContact(ctx, contactID)
}

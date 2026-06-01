package application

import (
	"github.com/google/uuid"
	"github.com/tombenke/go-application-template/internal/domain"
)

type ContactDTO struct {
	ID        uuid.UUID
	Email     string
	FirstName string
	LastName  string
	Phone     string
}

func ToContactDTO(contact *domain.Contact) ContactDTO {
	return ContactDTO{
		ID:        contact.ID,
		Email:     contact.Email,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
	}
}

type ContactsDTO []ContactDTO

func ToContactsDTO(contacts []*domain.Contact) ContactsDTO {
	contactsDTO := make(ContactsDTO, len(contacts))
	for i, contact := range contacts {
		contactsDTO[i] = ToContactDTO(contact)
	}
	return contactsDTO
}

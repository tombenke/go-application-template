package gtd

import (
	"github.com/tombenke/go-application-template/internal/application"
	"github.com/tombenke/go-application-template/internal/infrastructure/presentation/web/components/common"
)

type ContactDTOWithErrors struct {
	application.ContactDTO
	Errors map[string]string
}

type contactsPageData struct {
	common.PageData
	Contacts *application.ContactsDTO
}

type newContactPageData struct {
	common.PageData
	Contact *application.ContactDTO
}

type indexPageData struct {
	common.PageData
}

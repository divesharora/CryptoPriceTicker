package schemas

import (
	"github.com/google/uuid"
)

type Solr struct {
	Id                          uuid.UUID `json:"id"`
	FullTitle                   string    `json:"full_title"`
	Claimant                    string    `json:"claimant"`
	DocumentNumber              string    `json:"document_number"`
	RegistrationNumber          string    `json:"registration_number"`
	CopyrightSerialNumber       string    `json:"copyright_serial_number"`
	TypeOfWork                  string    `json:"type_of_work"`
	RecordType                  string    `json:"record_type"`
	CopyrightPublicationStatus  string    `json:"copyright_publication_status"`
	CopyrightRecordControlNumer string    `json:"copyright_record_control_number"`

	// Dates
	LocalCreationDate int64 `json:"local_creation_date,omitempty"`
	RegistrationDate  int64 `json:"registration_date,omitempty"`

	RecordationDate int64 `json:"recordation_date,omitempty"`

	PublicationDate          int64 `json:"publication_date,omitempty"`
	CopyrightPublicationDate int64 `json:"copyright_publication_date,omitempty"`

	CopyrightDateOfCertification int64 `json:"copyright_date_of_certification,omitempty"`
}
package schemas

import (
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

// CREATE TABLE COPYRIGHT_V2(
//     COPYRIGHT_ID SERIAL PRIMARY KEY,
//     ANALYTICAL_TITLE TEXT,
//     COPYRIGHT_PUBLICATION_STATUS TEXT   ,
//     COPYRIGHT_REGISTRATION_CLASS TEXT,
//     FULL_TITLE TEXT,
//     IMPRINT_INFORMATION TEXT,
//     NATION_OF_FIRST_PUBLICATION TEXT,
//     PLACE_OF_PUBLICATION TEXT,
//     RECORD_STATUS TEXT,
//     RECORD_TYPE TEXT,
//     SOURCE TEXT,
//     TYPE_OF_WORK TEXT,
//     COPYRIGHT_RECORD_CONTROL_NUMBER TEXT UNIQUE,

//     FOREIGN KEY(CONTENT_DESCRIPTION_DATA_ID) REFERENCES CONTENT_DESCRIPTION_DATA_V2(CONTENT_DESCRIPTION_DATA_ID) ON DELETE SET NULL,
//     FOREIGN KEY(FORM_DATA_ID) REFERENCES FORM_DATA_V2(FORM_DATA_ID) ON DELETE SET NULL,
//     FOREIGN KEY(NUMBER_DATA_ID) REFERENCES NUMBER_DATA_V2(NUMBER_DATA_ID) ON DELETE SET NULL,
//     FOREIGN KEY(DATE_DATA_ID) REFERENCES DATE_DATA_V2(DATE_DATA_ID) ON DELETE SET NULL,
//     FOREIGN KEY(LOCAL_DATA_ID) REFERENCES LOCAL_DATA_V2(LOCAL_DATA_ID) ON DELETE SET NULL,
//     FOREIGN KEY(OWNER_DATA_ID) REFERENCES OWNER_DATA_V2(OWNER_DATA_ID) ON DELETE SET NULL
// );

type Copyright struct {
	CopyrightID                  int    `json:"copyright_id"`
	AnalyticalTitle              string `json:"analytical_title"`
	CopyrightPublicationStatus   string `json:"copyright_publication_status"`
	CopyrightRegistrationClass   string `json:"copyright_registration_class"`
	FullTitle                    string `json:"full_title"`
	ImprintInformation           string `json:"imprint_information"`
	NationOfFirstPublication     string `json:"nation_of_first_publication"`
	PlaceOfPublication           string `json:"place_of_publication"`
	RecordStatus                 string `json:"record_status"`
	RecordType                   string `json:"record_type"`
	Source                       string `json:"source"`
	TypeOfWork                   string `json:"type_of_work"`
	CopyrightRecordControlNumber string `json:"copyright_record_control_number"`
	ContentDescriptionDataID     int    `json:"content_description_data_id"`
	FormDataID                   int    `json:"form_data_id"`
	NumberDataID                 int    `json:"number_data_id"`
	DateDataID                   int    `json:"date_data_id"`
	LocalDataID                  int    `json:"local_data_id"`
	OwnerDataID                  int    `json:"owner_data_id"`
}

// CREATE TABLE CONTENT_DESCRIPTION_DATA_V2(
//     CONTENT_DESCRIPTION_DATA_ID SERIAL PRIMARY KEY,
//     CONTENTS_NOTE TEXT,
//     COPYRIGHT_UNKNOWN_FIELD TEXT,
//     CURRENT_PUBLICATION_FREQUENCY TEXT,
//     DESCRIPTIVE_CATALOGING_FORM TEXT,
//     EDITION_STATEMENT TEXT,
//     ELECTRONIC_LOCATION_AND_ACCESS TEXT,
//     RETRIEVAL_CODE TEXT,
//     FORMER_PUBLICATION_FREQUENCY TEXT,
//     FREQUENCY TEXT,
//     GENERAL_BIBLIOGRAPHIC_NOTE TEXT,
//     LANGUAGE_NOTE TEXT,
//     TYPE_OF_WORK_PREREGISTERED TEXT,
//     IS_RECORDATION  TEXT,
//     LOGICAL_RECORD_LENGTH TEXT,
//     PHYSICAL_DESCRIPTION TEXT,
//     COPYRIGHT_RECORD_CONTROL_NUMBER TEXT UNIQUE
// );

type ContentDescriptionData struct {
	ContentDescriptionDataID     int    `json:"content_description_data_id"`
	ContentsNote                 string `json:"contents_note"`
	CopyrightUnknownField        string `json:"copyright_unknown_field"`
	CurrentPublicationFrequency  string `json:"current_publication_frequency"`
	DescriptiveCatalogingForm    string `json:"descriptive_cataloging_form"`
	EditionStatement             string `json:"edition_statement"`
	ElectronicLocationAndAccess  string `json:"electronic_location_and_access"`
	RetrievalCode                string `json:"retrieval_code"`
	FormerPublicationFrequency   string `json:"former_publication_frequency"`
	Frequency                    string `json:"frequency"`
	GeneralBibliographicNote     string `json:"general_bibliographic_note"`
	LanguageNote                 string `json:"language_note"`
	TypeOfWorkPreregistered      string `json:"type_of_work_preregistered"`
	IsRecordation                string `json:"is_recordation"`
	LogicalRecordLength          string `json:"logical_record_length"`
	PhysicalDescription          string `json:"physical_description"`
	CopyrightRecordControlNumber string `json:"copyright_record_control_number"`
}

// CREATE TABLE FORM_DATA_V2(
//     FORM_DATA_ID SERIAL PRIMARY KEY,
//     UNKNOWN_FIELD_917K TEXT,
//     UNKNOWN_FIELD_917L TEXT,
//     CATALOGING_SOURCE TEXT,
//     CONSTITUENT_UNIT_ENTRY TEXT,
//     CREATION_CREDITS_NOTE TEXT,
//     DOCUMENT_TYPE TEXT,
//     FORM_TITLE TEXT,
//     GEOGRAPHIC_NAME TEXT,
//     HOST_ITEM_ENTRY TEXT,
//     KEY_TITLE TEXT,
//     PARENT_RECORD_ENTRY TEXT,
//     PRECEDING_ENTRY TEXT,
//     SCALE_NOTE_FOR_GRAPHIC_MATERIALS TEXT,
//     SERIES_STATEMENT TEXT,
//     SUCCEEDING_ENTRY TEXT,
//     TITLE_LOCAL TEXT,
//     TO_PAGE_VALUE TEXT,
//     UNCONTROLLED_MISC TEXT,
//     UNCONTROLLED_NAME TEXT,
//     UNIFORM_TITLE TEXT,
//     VARYING_FORM_OF_TITLE TEXT,
//     COPYRIGHT_RECORD_CONTROL_NUMBER TEXT UNIQUE
// );

type FormData struct {
	FormDataID                   int    `json:"form_data_id"`
	UnknownField917K             string `json:"unknown_field_917k"`
	UnknownField917L             string `json:"unknown_field_917l"`
	CatalogingSource             string `json:"cataloging_source"`
	ConstituentUnitEntry         string `json:"constituent_unit_entry"`
	CreationCreditsNote          string `json:"creation_credits_note"`
	DocumentType                 string `json:"document_type"`
	FormTitle                    string `json:"form_title"`
	GeographicName               string `json:"geographic_name"`
	HostItemEntry                string `json:"host_item_entry"`
	KeyTitle                     string `json:"key_title"`
	ParentRecordEntry            string `json:"parent_record_entry"`
	PrecedingEntry               string `json:"preceding_entry"`
	ScaleNoteForGraphicMaterials string `json:"scale_note_for_graphic_materials"`
	SeriesStatement              string `json:"series_statement"`
	SucceedingEntry              string `json:"succeeding_entry"`
	TitleLocal                   string `json:"title_local"`
	ToPageValue                  string `json:"to_page_value"`
	UncontrolledMisc             string `json:"uncontrolled_misc"`
	UncontrolledName             string `json:"uncontrolled_name"`
	UniformTitle                 string `json:"uniform_title"`
	VaryingFormOfTitle           string `json:"varying_form_of_title"`
	CopyrightRecordControlNumber string `json:"copyright_record_control_number"`
}

// CREATE TABLE NUMBER_DATA_V2(
//     NUMBER_DATA_ID SERIAL PRIMARY KEY,
//     REGISTRATION_NUMBER TEXT,
//     COPYRIGHT_RECORD_CONTROL_NUMBER TEXT UNIQUE,
//     PUBLISHER_NUMBER TEXT,
//     OTHER_STANDARD_NUMBERS TEXT,
//     IN_PROCESS_NUMBER TEXT,
//     INTERNATIONAL_STANDARD_BOOK_NUMBER TEXT,
//     INTERNATIONAL_STANDARD_SERIAL_NUMBER TEXT,
//     COPYRIGHT_SERIAL_NUMBER TEXT,
//     DOCUMENT_NUMBER TEXT,
//     DOCUMENT_NUMBER_RANGE TEXT
// );

type NumberData struct {
	NumberDataID                      int    `json:"number_data_id"`
	RegistrationNumber                string `json:"registration_number"`
	CopyrightRecordControlNumber      string `json:"copyright_record_control_number"`
	PublisherNumber                   string `json:"publisher_number"`
	OtherStandardNumbers              string `json:"other_standard_numbers"`
	InProcessNumber                   string `json:"in_process_number"`
	InternationalStandardBookNumber   string `json:"international_standard_book_number"`
	InternationalStandardSerialNumber string `json:"international_standard_serial_number"`
	CopyrightSerialNumber             string `json:"copyright_serial_number"`
	DocumentNumber                    string `json:"document_number"`
	DocumentNumberRange               string `json:"document_number_range"`
}

// CREATE TABLE DATE_DATA_V2(
//     DATE_DATA_ID SERIAL PRIMARY KEY,
//     COPYRIGHT_BEGIN_DATE TEXT,
//     COPYRIGHT_END_DATE TEXT,
//     COPYRIGHT_MASK_WORK_COMMERCIAL_EXPLOITATION_DATE TEXT,
//     COPYRIGHT_REPRESENTATIVE_DATE TEXT,
//     COPYRIGHT_SIGNED_DATE_UNFORMATTED TEXT,
//     DATE_ENTERED_ON_FILE TEXT,
//     DATE_OF_LATEST_TRANSACTION TEXT,
//     TIME_OF_LATEST_TRANSACTION TEXT,
//     DATE_OF_RECORDATION TEXT,
//     DATES_OF_PUBLICATION_AND_ENUMERATION TEXT,
//     REGISTRATION_DATE TEXT,
//     PROJECTED_PUBLICATION_DATE TEXT,
//     PUBLICATION_DATE TEXT,
//     COPYRIGHT_DATE_OF_CERTIFICATION TEXT,
//     COPYRIGHT_DATE_OF_PUBLICATION TEXT,
//     LOCAL_DATE_OF_CREATION TEXT,
//     COPYRIGHT_WORK_CREATION_YEAR TEXT,
//     COPYRIGHT_YEAR_OF_PUBLICATION_STATED_IN_COPYRIGHT_NOTICE TEXT,
//     COPYRIGHT_RECORD_CONTROL_NUMBER TEXT UNIQUE
// );

type DateData struct {
	DateDataID                                        int    `json:"date_data_id"`
	CopyrightBeginDate                                string `json:"copyright_begin_date"`
	CopyrightEndDate                                  string `json:"copyright_end_date"`
	CopyrightMaskWorkCommercialExploitationDate       string `json:"copyright_mask_work_commercial_exploitation_date"`
	CopyrightRepresentativeDate                       string `json:"copyright_representative_date"`
	CopyrightSignedDateUnformatted                    string `json:"copyright_signed_date_unformatted"`
	DateEnteredOnFile                                 string `json:"date_entered_on_file"`
	DateOfLatestTransaction                           string `json:"date_of_latest_transaction"`
	TimeOfLatestTransaction                           string `json:"time_of_latest_transaction"`
	DateOfRecordation                                 string `json:"date_of_recordation"`
	DatesOfPublicationAndEnumeration                  string `json:"dates_of_publication_and_enumeration"`
	RegistrationDate                                  string `json:"registration_date"`
	ProjectedPublicationDate                          string `json:"projected_publication_date"`
	PublicationDate                                   string `json:"publication_date"`
	CopyrightDateOfCertification                      string `json:"copyright_date_of_certification"`
	CopyrightDateOfPublication                        string `json:"copyright_date_of_publication"`
	LocalDateOfCreation                               string `json:"local_date_of_creation"`
	CopyrightWorkCreationYear                         string `json:"copyright_work_creation_year"`
	CopyrightYearOfPublicationStatedInCopyrightNotice string `json:"copyright_year_of_publication_stated_in_copyright_notice"`
	CopyrightRecordControlNumber                      string `json:"copyright_record_control_number"`
}

// CREATE TABLE LOCAL_DATA_V2(
//     LOCAL_DATA_ID SERIAL PRIMARY KEY,
//     LOCAL_APPLICATION_AUTHOR_STATEMENT TEXT,
//     LOCAL_APPLICATION_TITLE_STATEMENT TEXT,
//     LOCAL_CANCELLATION_NOTE TEXT,
//     LOCAL_COPYRIGHT_NOTE TEXT,
//     LOCAL_DESC_OF_PREREGISTERED_NOTE TEXT,
//     LOCAL_FORM_EDITION_STATEMENT TEXT,
//     LOCAL_HOLDINGS TEXT,
//     LOCAL_MATERIAL_EXCLUDED TEXT,
//     LOCAL_MATERIAL_NOTE TEXT,
//     LOCAL_PARTY_ONE_STATEMENT TEXT,
//     LOCAL_PARTY_TWO_STATEMENT TEXT,
//     LOCAL_PREREGISTRATION_NOTE TEXT,
//     LOCAL_PREVIOUS_REGISTRATION_NOTE TEXT,
//     LOCAL_REFERENCE_ENTRY TEXT,
//     LOCAL_RIGHTS_AND_PERMISSION_NOTE TEXT,
//     LOCAL_SPLIT_CLAIM_NOTE TEXT,
//     LOCAL_SUBSEQUENT_REGISTRATION_NOTE TEXT,
//     LOCAL_SUPPLEMENTARY_REGISTRATION TEXT,
//     LOCAL_TITLE_QUALIFYING_INFORMATION TEXT,
//     LOCAL_VARYING_FORM_OF_TITLE_FROM_APPLICATION TEXT,
//     COPYRIGHT_RECORD_CONTROL_NUMBER TEXT UNIQUE
// );

type LocalData struct {
	LocalDataID                            int    `json:"local_data_id"`
	LocalApplicationAuthorStatement        string `json:"local_application_author_statement"`
	LocalApplicationTitleStatement         string `json:"local_application_title_statement"`
	LocalCancellationNote                  string `json:"local_cancellation_note"`
	LocalCopyrightNote                     string `json:"local_copyright_note"`
	LocalDescOfPreregisteredNote           string `json:"local_desc_of_preregistered_note"`
	LocalFormEditionStatement              string `json:"local_form_edition_statement"`
	LocalHoldings                          string `json:"local_holdings"`
	LocalMaterialExcluded                  string `json:"local_material_excluded"`
	LocalMaterialNote                      string `json:"local_material_note"`
	LocalPartyOneStatement                 string `json:"local_party_one_statement"`
	LocalPartyTwoStatement                 string `json:"local_party_two_statement"`
	LocalPreregistrationNote               string `json:"local_preregistration_note"`
	LocalPreviousRegistrationNote          string `json:"local_previous_registration_note"`
	LocalReferenceEntry                    string `json:"local_reference_entry"`
	LocalRightsAndPermissionNote           string `json:"local_rights_and_permission_note"`
	LocalSplitClaimNote                    string `json:"local_split_claim_note"`
	LocalSubsequentRegistrationNote        string `json:"local_subsequent_registration_note"`
	LocalSupplementaryRegistration         string `json:"local_supplementary_registration"`
	LocalTitleQualifyingInformation        string `json:"local_title_qualifying_information"`
	LocalVaryingFormOfTitleFromApplication string `json:"local_varying_form_of_title_from_application"`
	CopyrightRecordControlNumber           string `json:"copyright_record_control_number"`
}

// CREATE TABLE OWNER_DATA_V2(
//     OWNER_DATA_ID SERIAL PRIMARY KEY,
//     CLAIMANT TEXT,
//     PERFORMER_NOTE TEXT,
//     PERSONAL_NAME_ADDED TEXT,
//     PERSONAL_NAME_LOCAL TEXT,
//     CORPORATE_NAME_ADDED TEXT,
//     CORPORATE_NAME_LOCAL TEXT,
//     COPYRIGHT_RECORD_CONTROL_NUMBER TEXT UNIQUE
// );

type OwnerData struct {
	OwnerDataID                  int    `json:"owner_data_id"`
	Claimant                     string `json:"claimant"`
	PerformerNote                string `json:"performer_note"`
	PersonalNameAdded            string `json:"personal_name_added"`
	PersonalNameLocal            string `json:"personal_name_local"`
	CorporateNameAdded           string `json:"corporate_name_added"`
	CorporateNameLocal           string `json:"corporate_name_local"`
	CopyrightRecordControlNumber string `json:"copyright_record_control_number"`
}

type CopyrightableSearchSchema struct {
	InputQuery                         string   `json:"input_query"`
	TypeOfWorkSelected                 []string `json:"type_of_work_selected"`
	RecordTypeSelected                 []string `json:"record_type_selected"`
	CopyrightPublicationStatusSelected []string `json:"copyright_publication_status_selected"`
	FromLocalCreationDate              []int64  `json:"from_local_creation_date"`
	ToLocalCreationDate                []int64  `json:"to_local_creation_date"`
	FromRegistrationDate               []int64  `json:"from_registration_date"`
	ToRegistrationDate                 []int64  `json:"to_registration_date"`
	FromRecordationDate                []int64  `json:"from_recordation_date"`
	ToRecordationDate                  []int64  `json:"to_recordation_date"`
	FromPublicationDate                []int64  `json:"from_publication_date"`
	ToPublicationDate                  []int64  `json:"to_publication_date"`
	FromCopyrightPublicationDate       []int64  `json:"from_copyright_publication_date"`
	ToCopyrightPublicationDate         []int64  `json:"to_copyright_publication_date"`
	FromCopyrightDateOfCertification   []int64  `json:"from_copyright_date_of_certification"`
	ToCopyrightDateOfCertification     []int64  `json:"to_copyright_date_of_certification"`
	PageNumber                         int      `json:"page_number"`
}

func (c *CopyrightableSearchSchema) Validate() []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	if c.PageNumber == 0 {
		c.PageNumber = 1
	}
	return errors
}

type FacetFields struct {
	TypeOfWork                 []interface{} `json:"type_of_work"`
	RecordType                 []interface{} `json:"record_type"`
	CopyrightPublicationStatus []interface{} `json:"copyright_publication_status"`
}

type FacetStruct struct {
	FacetFields FacetFields `json:"facet_fields"`
}

type EditRecord struct {
	FullTitle                    *string        `json:"full_title"`
	Description                  *string        `json:"description"`
	Image_url                    pq.StringArray `json:"image_url"`
	CopyrightRecordControlNumber string         `json:"copyright_record_control_number"`
	VersionID                    string         `json:"version_id"`
	UserID                       int            `json:"user_id"`
	EditType                     string         `json:"edit_type"`
	EditDate                     int            `json:"edit_date"`
	ApprovalStatus               bool           `json:"approval_status"`
	ProfileImageUrl              *string        `json:"profile_image_url"`
	UserName                     *string        `json:"user_name"`
}

type CommentSchema struct {
	CommentId                    int     `json:"comment_id"`
	UserId                       int     `json:"user_id"`
	ParentId                     *int64  `json:"parent_id"`
	CommentText                  *string `json:"comment_text"`
	CopyrightRecordControlNumber *string `json:"copyright_record_control_number"`
	CreatedAt                    int     `json:"created_at"`
}

type PageDataSchema struct {
	PageDataId                   *string        `json:"page_data_id"`
	CopyrightRecordControlNumber *string        `json:"copyright_record_control_number"`
	ImageUrl                     pq.StringArray `json:"image_url"`
	Rating                       float64        `json:"rating"`
	OwnerUserId                  *int           `json:"owner_user_id"`
	ReviewCount                  int            `json:"review_count"`
	Description                  *string        `json:"description"`
}

type InsertCommentSchema struct {
	UserId                       int     `json:"user_id"`
	ParentId                     *int64  `json:"parent_id"`
	CommentText                  *string `json:"comment_text"`
	CopyrightRecordControlNumber *string `json:"copyright_record_control_number"`
}

type InfoResponse struct {
	CopyrightRecordControlNumber *string        `json:"copyright_record_control_number"`
	FullTitle                    *string        `json:"full_title"`
	Claimant                     *string        `json:"claimant"`
	PersonalNameAdded            *string        `json:"personal_name_added"`
	LocalDateOfCreation          *string        `json:"local_date_of_creation"`
	DateOfRecordation            *string        `json:"date_of_recordation"`
	OwnerUserId                  *int           `json:"owner_user_id"`
	Rating                       float64        `json:"rating"`
	ImageUrlsArray               pq.StringArray `json:"image_url"`
	Description                  *string        `json:"description"`
}

func (c *InsertCommentSchema) Validate() []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

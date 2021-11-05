package models

type TestRequest struct {
	ID             string `json:"id"`
	NamePartner    string `json:"name" validate:"required"`
	ToDateTime     string `json:"username" validate:"required"`
	EmailPartner   string `json:"email" validate:"required,email"`
	AddressPartner AddressStruct
	PhonePartner   string `json:"phone" validate:"required"`
	WebsitePartner string `json:"website" validate:"required"`
	CompanyPartner CompanyStruct
}

type SalaryDataPrtnerStruct struct {
	IDPartnerSalary string `json:"id" `
	SalaryPartner   string `json:"salaryInIDR" validate:"required"`
}

type AddressStruct struct {
	StreetPartner  string `json:"street" validate:"required"`
	SuitePartner   string `json:"suite" validate:"required"`
	CityPartner    string `json:"city" validate:"required"`
	ZIPCodePartner string `json:"zipcode" validate:"required"`
	GeoPartner     GeoStruct
}

type GeoStruct struct {
	LatGeoPartner string `json:"lat" validate:"required"`
	LngGeoPartner string `json:"lng" validate:"required"`
}

type CompanyStruct struct {
	NameCompanyPartner        string `json:"name" validate:"required"`
	CatchPhraseCompanyPartner string `json:"catchPhrase" validate:"required"`
	BsCompanyPartner          string `json:"bs" validate:"required"`
}

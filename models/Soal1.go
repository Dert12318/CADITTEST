package models

type DataResponse struct {
	ID              int64  `json:"id" validate:"required"`
	NamePartner     string `json:"name" validate:"required"`
	UsernamePartner string `json:"username" validate:"required"`
	EmailPartner    string `json:"email" validate:"required,email"`
	AddressPartner  AddressStruct
	PhonePartner    string  `json:"phone" validate:"required"`
	SalaryIDR       float64 `json:"salaryIDR" validate:"required"`
	SalaryUSD       float64 `json:"salaryUSD" validate:"required"`
}
type TestRequest struct {
	ID              int64  `json:"id" validate:"required"`
	NamePartner     string `json:"name" validate:"required"`
	UsernamePartner string `json:"username" validate:"required"`
	EmailPartner    string `json:"email" validate:"required,email"`
	AddressPartner  AddressStruct
	PhonePartner    string `json:"phone" validate:"required"`
	WebsitePartner  string `json:"website" validate:"required"`
	CompanyPartner  CompanyStruct
}

type SalaryDataPrtnerStruct struct {
	IDPartnerSalary int     `json:"id" validate:"required"`
	SalaryPartner   float64 `json:"salaryInIDR" validate:"required"`
}
type SalaryDataPrtner struct {
	Array []SalaryDataPrtnerStruct `json:"array" validate:"required"`
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

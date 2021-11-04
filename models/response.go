package models

type TestResponse struct {
	ResponseCode       string //`json:"responseCode" validate:"required`  //disesuaikan dengan ESB
	ResponseMessage    string //`json:"mailto" validate:"required,email"` //disesuaikan dengan ESB
	ReferenceNo        string //`json:"cc"`								//disesuaikan dengan ESB
	PartnerReferenceNo string //`json:"subject" validate:"required"`		//disesuaikan dengan ESB
	Balance            BalanceStruct
	TotalCreditEntries TotalCreditEntriesStruct
	TotalDebitEntries  TotalDebitEntriesStruct
	DetailData         DetailDataStruct
	AdditionalInfo     AdditionalInfoStruct
}

type AdditionalInfoStruct struct {
	DeviceId string
	Channel  string
}

type DetailDataStruct struct {
	DetailBalance           []AmountStruct
	Amount                  AmountStruct
	OriginAmount            AmountStruct
	TransactionDate         string
	Remark                  string
	TransactionId           string
	Type                    string
	TransactionDetailStatus string
}

type BalanceStruct struct {
	Amount   AmountStruct
	DateTime string
}

type AmountStruct struct {
	Value    string
	Currency string
}

type TotalCreditEntriesStruct struct {
	NumberOfEntries string
	Amount          AmountStruct
}

type TotalDebitEntriesStruct struct {
	NumberOfEntries string
	Amount          AmountStruct
}

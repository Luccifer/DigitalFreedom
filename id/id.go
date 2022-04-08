package id

type ID struct {
	Name       Name
	Surname    Surname
	Patron     Patron
	Sex        Sex
	Birth      Birth
	Language   Language
	Education  []Edu
	Work       []Work
	Religion   Religion
	Duty       Duty
	DrivingLic []Driving
	Wallet     Wallet
	Issuer     Issuer
	Relatives  []ID
	Blood      ABO
	Residence  []Residence
	Passports  []Passport
	Signature  []byte
	Criminal   CriminalRecord
	Refugee    Refugee
	Photo      []byte //base64
}

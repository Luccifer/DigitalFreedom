package id

type ID struct {
	Name       string
	Surname    string
	Patronomic string
	Sex        Sex
	Birth      Birth
	Religion   Religion
	DrivingLic []Driving
	Wallet     string
	Issuer     string
	Relatives  []ID
	Blood      ABO
}

type Birth struct {
	Date    int64
	Country string
	City    string
}

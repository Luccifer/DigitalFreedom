package id

type Passport struct {
	country     string
	cityOfIssue string
	issuer      string
	dateOfIssue string
	number      string
	photo       []byte //base64
}

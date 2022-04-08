package id

type Passport struct {
	country         string
	cityOfIssue     string
	issuer          string
	dateOfIssue     int64
	dateOfDismissal int64
	ordinatory      bool
	number          string
	photo           []byte //base64
}

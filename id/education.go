package id

type Education int

type Edu struct {
	level  Education
	issuer string
	year   string
	study  string
}

const (

	// Early childhood education 0
	/* Education designed to support early development in preparation for participation in school and society. Programmes designed for children below the age of 3. */
	ECE1 Education = iota + 0

	// Early childhood education 1
	/* Education designed to support early development in preparation for participation in school and society. Programmes designed for children from age 3 to the start of primary education.*/
	ECE2 Education = iota + 1

	// Primary education 2
	/* Programmes typically designed to provide students with fundamental skills in reading, writing and mathematics and to establish a solid foundation for learning.*/
	PE Education = iota + 2

	// Lower secondary education
	/* First stage of secondary education building on primary education, typically with a more subject-oriented curriculum. */
	LSE Education = iota + 3

	// Upper secondary education
	/* Second/final stage of secondary education preparing for tertiary education and/or providing skills relevant to employment. Usually with an increased range of subject options and streams. */
	USE Education = iota + 4

	// Post-secondary non-tertiary education
	/* Programmes providing learning experiences that build on secondary education and prepare for labour market entry and/or tertiary education. The content is broader than secondary but not as complex as tertiary education.*/
	PSE Education = iota + 5

	// Short-cycle tertiary education
	/* Short first tertiary programmes that are typically practically-based, occupationally-specific and prepare for labour market entry. These programmes may also provide a pathway to other tertiary programmes. */
	STE Education = iota + 6

	// Bachelor's or equivalent
	/* Programmes designed to provide intermediate academic and/or professional knowledge, skills and competencies leading to a first tertiary degree or equivalent qualification.*/
	Bs Education = iota + 7

	// Master's or equivalent
	/* Programmes designed to provide advanced academic and/or professional knowledge, skills and competencies leading to a second tertiary degree or equivalent qualification.*/
	Ms Education = iota + 8

	// Master's or equivalent+
	/* Programmes designed to provide advanced academic and/or professional knowledge, skills and competencies leading to a second tertiary degree or additional extra equivalent qualification. */
	MsPlus Education = iota + 9

	// Doctorate or equivalent
	/* Programmes designed primarily to lead to an advanced research qualification, usually concluding with the submission and defense of a substantive dissertation of publishable quality based on original research.*/
	DOC Education = iota + 10

	// Doctorate or equivalent+
	/* Programmes designed primarily to lead to an advanced research qualification, usually concluding with the submission and defense of a substantive dissertation of publishable quality based on original extra plus research +.*/
	DOCPlus Education = iota + 11
)

func (e Education) String() string {
	return [...]string{"ECE1", "ECE2", "PE", "LSE", "LSE", "USE", "PSE", "STE", "Bs", "Ms", "MsPlus", "DOC", "DOCPlus"}[e-1]
}

func (e Education) EnumIndex() int {
	return int(e)
}

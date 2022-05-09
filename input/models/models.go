package model

type Search struct {
	Id          string `bson:"_id"`
	Url         string
	Container 	string
	Patterns    []string
	Title       []string
	Desc        []string
	Benefits    []string
	Eligibility []string

	Completed   bool
}

type InsuranceData struct {
	Id          string `bson:"_id"`
	Url         string
	Title       string
	Desc        string
	Benefits    []string
	Eligibility []string
}

package sheets

type Entry struct {
	Timestamp    string
	EmailAddress string
	Result       Result
}

type Result struct {
	Winner      string
	RunnerUp    string
	ThirdPlace  string
	FourthPlace string
}

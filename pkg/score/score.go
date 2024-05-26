package score

type Score struct {
	Name       string `json:"name"`
	Teams      []Team `json:"teams"`
	TotalScore int    `json:"totalScore"`
}

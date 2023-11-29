package domain

type University struct {
	WebPages []string `json:"web_pages"`
	AlphaTwoCode string `json:"alpha_two_code"`
	Domains []string `json:"domains"`
	Name string `json:"name"`
	StateProvince *string `json:"state-province,omitempty"`
	Country string	`json:"country"`
}
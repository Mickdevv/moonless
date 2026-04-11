package email

type ContactFormEmailPayload struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Reason      string `json:"reason"`
	Message     string `json:"message"`
}

type SendEmail struct {
	ToList  []string `json:"to_list"`
	Msg     string   `json:"message"`
	Subject string   `json:"subject"`
}

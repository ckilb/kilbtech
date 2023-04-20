package route

import (
	"ckilb/kilbtech/mail"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ContactRequest struct {
	Name    string // anti spam
	Email   string
	Message string
}

type Contact struct {
	sender mail.Sender
}

func (r *Contact) Path() string {
	return "/contact"
}

func (r *Contact) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			w.WriteHeader(405)

			return
		}

		var body ContactRequest
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			log.Println(fmt.Errorf("decoding request body: %w", err))

			w.WriteHeader(400)

			return
		}

		if body.Name != "" {
			w.WriteHeader(204)

			return
		}

		content := fmt.Sprintf(
			"E-Mail: %s\n\nMessage:\n=====\n%s",
			body.Email,
			body.Message,
		)

		if err := r.sender.Send(content); err != nil {
			log.Println(fmt.Errorf("sending contact form: %w", err))

			w.WriteHeader(500)

			return
		}

		w.WriteHeader(204)
	})
}

func NewContact(sender mail.Sender) Route {
	return &Contact{
		sender: sender,
	}
}

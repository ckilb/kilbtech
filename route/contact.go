package route

import (
	"ckilb/kilbtech/mail"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (r *Contact) Method() string {
	return http.MethodPost
}

func (r *Contact) Page() string {
	return ""
}

func (r *Contact) Templates() []string {
	return []string{}
}

func (r *Contact) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body ContactRequest

		if err := c.BindJSON(&body); err != nil {
			log.Println(fmt.Errorf("decoding request body: %w", err))

			c.AbortWithStatus(400)

			return
		}

		if body.Name != "" {
			c.AbortWithStatus(204)

			return
		}

		content := fmt.Sprintf(
			"E-Mail: %s\n\nMessage:\n=====\n%s",
			body.Email,
			body.Message,
		)

		if err := r.sender.Send(content); err != nil {
			log.Println(fmt.Errorf("sending contact form: %w", err))

			c.AbortWithStatus(500)

			return
		}

		c.AbortWithStatus(204)
	}
}

func NewContact(sender mail.Sender) Route {
	return &Contact{
		sender: sender,
	}
}

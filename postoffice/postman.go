package postoffice

import (
	"github.com/go-mail/mail"
)

// Postman wraps mail dialer.
type Postman struct {
	dialer *mail.Dialer
}

// NewPostman creates a new instance of PostOffice
// Deprecate.
func NewPostman(host string, port int, user, pass string) Postman {
	return Postman{
		dialer: mail.NewDialer(host, port, user, pass),
	}
}

// NewPostman creates a new instance of PostOffice
func New(host string, port int, user, pass string) Postman {
	return Postman{
		dialer: mail.NewDialer(host, port, user, pass),
	}
}

// Deliver asks the postman to deliver a parcel.
func (pm Postman) Deliver(p Parcel) error {
	m := mail.NewMessage()

	m.SetAddressHeader("From", p.FromAddress, p.FromName)
	m.SetAddressHeader("To", p.ToAddress, p.ToName)
	m.SetHeader("Subject", p.Subject)
	m.SetBody("text/plain", p.Body)

	if err := pm.dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

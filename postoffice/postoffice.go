package postoffice

import (
	"github.com/FTChinese/go-rest/connect"
	"github.com/go-mail/mail"
)

// PostOffice wraps mail dialer.
type PostOffice struct {
	dialer *mail.Dialer
}

// New creates a new instance of PostOffice
func New(c connect.Connect) PostOffice {
	return PostOffice{
		dialer: mail.NewDialer(c.Host, c.Port, c.User, c.Pass),
	}
}

// Deliver asks the postman to deliver a parcel.
func (pm PostOffice) Deliver(p Parcel) error {
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

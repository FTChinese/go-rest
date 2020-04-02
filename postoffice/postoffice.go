package postoffice

import "github.com/go-mail/mail"

// Postman wraps mail dialer.
type PostOffice struct {
	dialer *mail.Dialer
}

// NewPostman creates a new instance of PostOffice
func New(host string, port int, user, pass string) PostOffice {
	return PostOffice{
		dialer: mail.NewDialer(host, port, user, pass),
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

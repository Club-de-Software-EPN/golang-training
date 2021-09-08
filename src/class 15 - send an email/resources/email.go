package resources

import (
  "crypto/tls"
  "fmt"
  gomail "gopkg.in/mail.v2"
)


func SendEmail(
		emailSender string,
		passwordSender string,
		emailReceiver string,
		emailSubject string,
		pathsImages []string,
	){
	mail := gomail.NewMessage()

	// Headers
	mail.SetHeader("From", emailSender)
	mail.SetHeader("To", emailReceiver)
	mail.SetHeader("Subject", emailSubject)

	for _, path := range pathsImages {
		mail.Embed(path)
		mail.SetBody("text/html", fmt.Sprintf(`<img src="cid:%s"/>`,path))
	}

	// Settings for SMTP server
	emailDialer := gomail.NewDialer("smtp.gmail.com", 587, emailSender, passwordSender)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	emailDialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := emailDialer.DialAndSend(mail); err != nil {
		fmt.Println(err)
		panic(err)
	}
}


package notifier

import (
	"bytes"
	"fmt"
	"github.com/mrahbar/my-bloody-hetzner-sb-notifier/hetzner"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

const (
	gmailSmtpServer = "smtp.gmail.com"
	contentTypeHtml = "text/html"
	contentTypePlain = "text/plain"
)

type Notifier struct {
	to string
	contentType string
	user     string
	password string
	server   string
	port     int
	history map[string]float64
}

func NewNotifier(Recipient, Username, Password string) Notifier {
	notifier := Notifier{
		Recipient,
		contentTypePlain,
		Username,
		Password,
		gmailSmtpServer,
		587,
		make(map[string]float64),
	}
	return notifier
}

func (n Notifier) Act(servers []hetzner.Server) {
	//absentFromHistory := false
	//
	//for _, s := range servers {
	//	// TODO check if server in history, if a single server is missing send email and store state
	//}
}

func (n Notifier) sendMail(Dest []string, Subject, bodyMessage string) {

	msg := "From: " + n.user + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(fmt.Sprintf("%s:%d", n.server, n.port),
		smtp.PlainAuth("", n.user, n.password, n.server),
		n.user, Dest, []byte(msg))

	if err != nil {

		fmt.Printf("smtp error: %s", err)
		return
	}

	fmt.Println("Mail sent successfully!")
}

func (n Notifier) writeEmail(subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = n.user

	header["To"] = n.to
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", n.contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}
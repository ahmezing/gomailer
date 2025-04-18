package email

import (
	"bytes"
	"fmt"
)

type Message struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func NewMessage(from string, to []string, subject, body string) *Message {
	return &Message{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func (m *Message) Build() []byte {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString("MIME-version: 1.0;\n")
	buf.WriteString("Content-Type: text/html; charset=\"UTF-8\";\n\n")
	
	buf.WriteString(m.Body)

	return buf.Bytes()
}
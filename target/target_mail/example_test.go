package target_mail

import "github.com/admpub/log"

func ExampleNewMailTarget() {
	logger := log.NewLogger()

	// creates a MailTarget which sends emails to admin@example.com
	target := NewMailTarget()
	target.Host = "smtp.example.com"
	target.Username = "foo"
	target.Password = "bar"
	target.Subject = "log messages for foobar"
	target.Sender = "admin@example.com"
	target.Recipients = []string{"admin@example.com"}

	logger.Targets = append(logger.Targets, target)

	logger.Open()

	// ... logger is ready to use ...
}

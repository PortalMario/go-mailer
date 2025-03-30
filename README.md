# go-mailer
A library to send mails using an secure SMTP-Relay. Credentials are only sent to the relay host, if TLS is being used. (as stated within the [PlainAuth function of "net/smtp"](https://pkg.go.dev/net/smtp#PlainAuth))

## Usage
Take a look at the provided [example](./examples/main.go) to integrate the mailer. 
package main

import (
	"bytes"
	"flag"
	"html/template"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

type Recipient struct {
	Name     string
	Email    string
	Subject  string
	Body     string
	FromName string
}

type SMTPConfig struct {
	Host     string
	Port     string
	Email    string
	Pass     string
	FromName string
}

func main() {

	// CLI arguments (user input)
	subject := flag.String("subject", "Hello!", "Email subject")
	body := flag.String("body", "This is a test message!", "Email body")
	flag.Parse()

	// Load SMTP config from .env
	cfg := SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Email:    os.Getenv("SMTP_EMAIL"),
		Pass:     os.Getenv("SMTP_PASS"),
		FromName: os.Getenv("SMTP_FROM_NAME"),
	}

	recipientChannel := make(chan Recipient)

	// producer: load CSV into channel
	go func() {
		loadRecipient("./emails.csv", recipientChannel, *subject, *body, cfg.FromName)
	}()


	var wg sync.WaitGroup
	workerCount := 5

	// start workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg, cfg)
	}

	wg.Wait()
}

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}

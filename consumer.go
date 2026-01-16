package main

import (
	"fmt"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup, cfg SMTPConfig) {
	defer wg.Done()

	// Setup Gmail auth
	auth := smtp.PlainAuth("", cfg.Email, cfg.Pass, cfg.Host)
	smtpAddr := cfg.Host + ":" + cfg.Port

	for recipient := range ch {

		msg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker %d: template error for %s\n", id, recipient.Email)
			continue
		}

		fmt.Printf("Worker %d: sending to %s\n", id, recipient.Email)

		err = smtp.SendMail(
			smtpAddr,
			auth,
			cfg.Email,
			[]string{recipient.Email},
			[]byte(msg),
		)

		if err != nil {
			fmt.Printf("Worker %d: failed to send to %s, error: %v\n", id, recipient.Email, err)
			continue
		}

		fmt.Printf("Worker %d: success for %s\n", id, recipient.Email)

		// Gmail anti-spam throttle
		time.Sleep(40 * time.Millisecond)
	}
}

package main

import (
	"fmt"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		emailMsg, err := excuteTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker %d: Failed to execute template for %s: %v\n", id, recipient.Email, err)
			// todo: future add to deadlatter queue here
			continue
		}
		fmt.Printf("Worker %d: Sending email to %s\n", id, recipient.Email)

		err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "108tiwari.g@gmail.com", []string{recipient.Email}, emailMsg)

		if err != nil {
			fmt.Printf("Worker %d: Failed to send email to %s: %v\n", id, recipient.Email, err)
			continue
		}

		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: Email sent to %s successfully\n", id, recipient.Email)

	}
}

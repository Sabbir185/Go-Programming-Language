package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
)

func sendEmail(workerid int, user User, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Sending email to %s at %s\n", user.Name, user.Email)

	msg, temp_err := executeTemplate(user)
	if temp_err != nil {
		log.Fatal(temp_err)
		return
	}

	smtpHost := "localhost"
	smtpPort := "1025"

	fmt.Printf("Worker %d: Sending email to %s \n", workerid, user.Email)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		nil, "sabbir.py@gmail.com",
		[]string{user.Email},
		[]byte(msg),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Worker %d: Sent email to %s \n", workerid, user.Email)
}

package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(workerid int, ch chan User, wg *sync.WaitGroup) {
	defer wg.Done()

	for user := range ch {
		log.Printf("Sending email to %s at %s\n", user.Name, user.Email)

		msg, temp_err := executeTemplate(user)
		if temp_err != nil {
			log.Fatal(temp_err)
			continue
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

		time.Sleep(time.Millisecond * 50)

		fmt.Printf("Worker %d: Sent email to %s \n", workerid, user.Email)
	}
}

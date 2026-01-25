package main

import (
	"bytes"
	"html/template"
	"log"
	"sync"
)

type User struct {
	Name  string
	Email string
}

func main() {
	ch := make(chan User)
	go sendReceipients("./users.csv", ch)

	var wg sync.WaitGroup
	workerCount := 5

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, ch, &wg)
	}

	wg.Wait()
}

func executeTemplate(u User) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, u)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return tpl.String(), nil
}

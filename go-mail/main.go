package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"gopkg.in/gomail.v2"
)

func sendGomail(templatePath string) {

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ name string }{name: "Shaheer"})

	if err != nil {
		fmt.Println(err)
		return
	}
	//send with gomail

	m := gomail.NewMessage()
	m.SetHeader("From", "shaheer252001@gmail.com")
	m.SetHeader("To", "mailshaheer316@gmai;.com" /*,another@email.com"*/)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	m.Attach("./cat.avif")

	d := gomail.NewDialer("smtp.example.com", 587, "shaheer252001@gmail.com", "qvgggmwlxldjzedw")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func sendMailSimple(subject string, body string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"shaheer252001@gmail.com",
		"qvgggmwlxldjzedw",
		"smtp.gmail.com",
	)

	msg := "Subject:" + subject + "\n" + body
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"shaheer252001@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func sendMailSimpleHtml(subject string, templatePath string, to []string) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ name string }{name: "Shaheer"})
	if err != nil {
		fmt.Println(err)
		return
	}
	auth := smtp.PlainAuth(
		"",
		"shaheer252001@gmail.com",
		"Shaheer_1549",
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject:" + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"shaheer252001@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// sendMailSimple(
	//  "Resume",
	// "Another Body",
	// []string{"shaheer252001@gmail.com"})

	// 	sendMailSimpleHtml(
	// 	"Resume",
	// 	"./test.html ",
	// 	[]string{"shaheer252001@gmail.com"})

	sendGomail("./test.html")

}

//we can also send emails with third party website called as sendgrid
//to Do that we have to need sendgrid Api key from Twiloo

package main

import(
	"crypto/tls"
	"gopkg.in/gomail.v2"
)



func main(){
	
    m := gomail.NewMessage()
	m.SetHeader("From", "mailing@disnovo.com")
	m.SetHeader("To", "vflores@disnovo.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com",587,"mailing@disnovo.com","uhdfjzrhnidkrxig")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
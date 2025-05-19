package internal

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
)

// 使用 STARTTLS 协议连接
func SendMailUsingSTARTTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("TCP Dial error:", err)
		return err
	}

	host := strings.Split(addr, ":")[0]
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Println("Create SMTP client error:", err)
		return err
	}
	defer client.Quit()

	tlsConfig := &tls.Config{ServerName: host}
	if ok, _ := client.Extension("STARTTLS"); ok {
		if err = client.StartTLS(tlsConfig); err != nil {
			log.Println("StartTLS error:", err)
			return err
		}
	} else {
		return fmt.Errorf("server does not support STARTTLS")
	}

	if auth != nil {
		if ok, _ := client.Extension("AUTH"); ok {
			if err = client.Auth(auth); err != nil {
				log.Println("Auth error:", err)
				return err
			}
		}
	}

	if err = client.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	return w.Close()
}

package main

import "fmt"

// Паттерн Command относится к поведенческим паттернам
// Обесппечивает обработку команды ввиде объекта, что позволяет сохранять её, передавать и возвращать в качесве параметра
// Шаблон применяют в тех случаях, когда воспроизведение команды нужно отделить от ее определения (логирование, выполнение на другом устройстве и т.п.)

type Command interface {
	Execute()
	Cancel()
}

type EmailBox struct {
	mails map[string]interface{}
}

func (b *EmailBox) SendLetter(author, mail string) Command {
	return &SendEmail{
		box:    b,
		author: author,
		mail:   mail,
	}
}
func (b *EmailBox) Clear() Command {
	return &CleanMails{
		box:     b,
		dropped: make(map[string]interface{}),
	}
}

func NewEmailBox() *EmailBox {
	return &EmailBox{
		mails: make(map[string]interface{}),
	}
}

type SendEmail struct {
	box          *EmailBox
	author, mail string
}

func (s *SendEmail) Execute() {
	s.box.mails[s.author] = s.mail
}
func (s *SendEmail) Cancel() {
	delete(s.box.mails, s.author)
}

type CleanMails struct {
	box     *EmailBox
	dropped map[string]interface{}
}

func (c *CleanMails) Execute() {
	c.dropped = c.box.mails
	c.box.mails = make(map[string]interface{})
}
func (c *CleanMails) Cancel() {
	c.box.mails = c.dropped
	c.dropped = make(map[string]interface{})
}

func main() {
	var myMail = NewEmailBox()
	var commands = []Command{
		myMail.SendLetter("Jack", "hi"),
		myMail.SendLetter("Nick", "hi"),
		myMail.SendLetter("John", "hi"),
		myMail.Clear(),
	}
	var even = false
	for _, c := range commands {
		c.Execute()
		if even {
			c.Cancel()
		}
		even = !even
	}
	fmt.Println(myMail.mails)
}

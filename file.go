package main

import (
	"io"
	"io/ioutil"
	"os"
)

type Message interface {
	Show() ([]byte, error)
	Add(bytes []byte) error
}

type message struct {
}

func New() Message {
	return &message{}
}

func (m *message) Show() ([]byte, error) {
	return ioutil.ReadFile("./message.log")
}

func (m *message) Add(data []byte) (err error) {
	var file *os.File

	if _, err = os.Stat("./message.log"); err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		file, err = os.Create("./message.log")
		if err != nil {
			return err
		}
	} else {
		file, err = os.OpenFile("./message.log", os.O_RDONLY|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	if _, err = io.WriteString(file, string(data)+"\n"); err != nil {
		return err
	}

	return nil
}

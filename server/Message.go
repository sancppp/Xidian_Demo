package server

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"tianzhenxiongProject/dal"
	"tianzhenxiongProject/model"
)

type Message struct {
	Conn      net.Conn
	Operation int
	Name      string
	Body      string
}

func (m Message) Quit() {
	msg := "quit"
	SendAll(msg, m.Name)
}

func (m Message) CreateStudent() {
	stu := &model.Student{}
	json.Unmarshal([]byte(m.Body), stu)
	err := dal.StudentDal{}.CreateStudent(stu)
	if err != nil {
		m.Conn.Write([]byte("Create Student error!" + fmt.Sprint(err)))
	} else {
		bytes, _ := json.Marshal(stu)
		m.Conn.Write([]byte("Create Student success!\n"))
		m.Conn.Write(bytes)
	}
}

func (m Message) ModifyStudent() {
	stu := &model.Student{}
	json.Unmarshal([]byte(m.Body), stu)
	err := dal.StudentDal{}.ModifyStudent(stu)
	if err != nil {
		m.Conn.Write([]byte("Modify Student error!" + fmt.Sprint(err)))
	} else {
		bytes, _ := json.Marshal(stu)
		m.Conn.Write([]byte("Modify Student success!\n"))
		m.Conn.Write(bytes)
	}
}

func (m Message) GetStudent() {
	var id int
	var err error
	if id, err = strconv.Atoi(m.Body); err != nil {
		m.Conn.Write([]byte("Get Student error!" + fmt.Sprint(err)))
		return
	}

	student, err := dal.StudentDal{}.GetStudentByID(id)
	if err != nil {
		m.Conn.Write([]byte("Get Student error!" + fmt.Sprint(err)))
		return
	}

	res, _ := json.Marshal(student)
	m.Conn.Write([]byte("Get Student success!\n"))
	m.Conn.Write(res)
}
func (m Message) DeleteStudent() {
	var id int
	var err error
	if id, err = strconv.Atoi(m.Body); err != nil {
		m.Conn.Write([]byte("Delete Student error!" + fmt.Sprint(err)))
	}

	err = dal.StudentDal{}.DeleteStudentByID(id)
	if err != nil {
		m.Conn.Write([]byte("Delete Student error!" + fmt.Sprint(err)))
	} else {
		m.Conn.Write([]byte("Delete Success!\n"))
	}
}
func (m Message) SendMessage() {
	SendAll(m.Body, m.Name)
}

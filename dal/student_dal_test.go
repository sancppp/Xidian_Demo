package dal

import (
	"fmt"
	"testing"
	"tianzhenxiongProject/model"
	"tianzhenxiongProject/mysql"
)

func TestStudentDal_CreateStudent(t *testing.T) {
	mysql.Default()
	stu := &model.Student{
		Name:       "tianzhenxiong",
		Department: "123",
		Major:      "1234",
	}
	err := StudentDal{}.CreateStudent(stu)
	fmt.Println(stu)
	if err != nil {
		t.Failed()
	}
}

func TestStudentDal_GetStudentByID(t *testing.T) {
	mysql.Default()
	student, err := StudentDal{}.GetStudentByID(1)
	if err != nil {
		t.Failed()
	}
	fmt.Println(student)

}

func TestStudentDal_ModifyStudent(t *testing.T) {
	mysql.Default()
	err := StudentDal{}.ModifyStudent(&model.Student{
		ID:         2,
		Name:       "tzx",
		Department: "321",
		Major:      "321",
	})
	if err != nil {
		t.Failed()
	}
}

func TestStudentDal_DeleteStudentByID(t *testing.T) {
	mysql.Default()
	err := StudentDal{}.DeleteStudentByID(1)
	if err != nil {
		t.Failed()
	}
}

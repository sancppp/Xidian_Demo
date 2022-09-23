package dal

import (
	"tianzhenxiongProject/model"
	"tianzhenxiongProject/mysql"
)

type StudentDal struct {
}

func (d StudentDal) CreateStudent(student *model.Student) error {
	return mysql.MysqlDB.GetConn().Create(&student).Error
}
func (d StudentDal) DeleteStudentByID(ID int) error {
	return mysql.MysqlDB.GetConn().Delete(&model.Student{}, ID).Error
}
func (d StudentDal) GetStudentByID(ID int) (*model.Student, error) {
	user := &model.Student{}
	if err := mysql.MysqlDB.GetConn().Where("Id = ?", ID).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (d StudentDal) ModifyStudent(student *model.Student) error {
	return mysql.MysqlDB.GetConn().Save(&student).Error
}

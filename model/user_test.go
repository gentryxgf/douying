package model

import "testing"

func TestQueryUserByNameAndPassword(t *testing.T) {
	InitMysql()
	username := "xgf"
	password := "25f9e794323b453885f5181f1b624d0b"

	user, err := NewUserDaoInstance().QueryUserByNameAndPassword(username, password)

	if err != nil {
		t.Log("err", err)
	}
	t.Log("userid", (*user).ID)
}

func TestAddUser(t *testing.T) {
	InitMysql()
	username := "123124"
	password := "21412124"
	user, err := NewUserDaoInstance().AddUser(username, password)

	if err != nil {
		t.Log("err", err)
	}
	t.Log("userid", (*user).ID)
}

func TestQueryUserByName(t *testing.T) {
	InitMysql()
	username := "eeee"
	_, err := NewUserDaoInstance().QueryUserByName(username)

	t.Log(err)
}

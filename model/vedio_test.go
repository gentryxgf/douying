package model

import (
	"testing"
	"time"
)

func TestAddVedio(t *testing.T) {
	InitMysql()
	vedio := Vedio{
		UserId: 2,
		Title:  "3333333333",
	}
	v, err := NewVedioDaoInstance().AddVedio(&vedio)
	if err != nil {
		t.Log(err)
	}
	t.Log(v.CreatedAt)
}

func TestQueryVedioListByTime(t *testing.T) {
	InitMysql()
	vedios, err := NewVedioDaoInstance().QueryVedioListByTime(2, time.Now())
	if err != nil {
		t.Log(err)
	}
	for _, v := range vedios {
		t.Log(v.ID)
	}
}

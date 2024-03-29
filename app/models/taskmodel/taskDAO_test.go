package taskmodel

import (
	"TaskRESTfulExercise/services/mongodb"
	"fmt"
	"testing"
	"time"
)

var (
	IP       = "127.0.0.1"
	PORT     = "27017"
	USERNAME = "user"
	PASSWORD = "user_password"
	POOLSIZE = "100"
	DATABASE = "task-restful-exercise"
	COLNAME  = "tasks"
)

func TestCreate(t *testing.T) {
	var tests = []struct {
		arg1 TaskDTO
	}{
		{
			TaskDTO{
				Name:      "task-0001",
				Status:    0,
				DueDate:   time.Now().Add(time.Duration(time.Hour * 24)),
				Timestamp: time.Now(),
			},
		},
	}
	mgClient, err := mongodb.ConnectMongoDB(IP, PORT, USERNAME, PASSWORD, POOLSIZE, DATABASE)
	if err != nil {
		t.Errorf("ConnectMongoDB :%v", err)
	}

	taskDAO := NewTaskDAOwithName(mgClient, DATABASE, COLNAME)

	for _, tt := range tests {
		if _, err := taskDAO.Create(&tt.arg1); err != nil {
			t.Errorf("%v", err)
		}
	}
}

func TestGetByID(t *testing.T) {
	var tests = []struct {
		arg1 string
	}{
		{
			arg1: "6606a575afc8ae7b93c642b4",
		},
	}
	mgClient, err := mongodb.ConnectMongoDB(IP, PORT, USERNAME, PASSWORD, POOLSIZE, DATABASE)
	if err != nil {
		t.Errorf("ConnectMongoDB :%v", err)
	}

	taskDAO := NewTaskDAOwithName(mgClient, DATABASE, COLNAME)

	for _, tt := range tests {
		if task, err := taskDAO.GetByID(tt.arg1); err != nil {
			t.Errorf("%v", err)
		} else {
			fmt.Printf("task : %+v\n", task)
		}
	}
}

func TestUpdate(t *testing.T) {
	var tests = []struct {
		arg1 string
		arg2 map[string]interface{}
	}{
		{
			arg1: "6606a575afc8ae7b93c642b4",
			arg2: map[string]interface{}{
				"name":   "task-0002",
				"status": 1,
			},
		},
	}
	mgClient, err := mongodb.ConnectMongoDB(IP, PORT, USERNAME, PASSWORD, POOLSIZE, DATABASE)
	if err != nil {
		t.Errorf("ConnectMongoDB :%v", err)
	}

	taskDAO := NewTaskDAOwithName(mgClient, DATABASE, COLNAME)

	for _, tt := range tests {
		if task, err := taskDAO.Update(tt.arg1, tt.arg2); err != nil {
			t.Errorf("%v", err)
		} else {
			fmt.Printf("task : %+v\n", task)
		}
	}
}

func TestDelete(t *testing.T) {
	var tests = []struct {
		arg1 string
	}{
		{
			arg1: "6606a575afc8ae7b93c642b4",
		},
		{
			arg1: "6606a6185d2090adab3ca1a5",
		},
	}
	mgClient, err := mongodb.ConnectMongoDB(IP, PORT, USERNAME, PASSWORD, POOLSIZE, DATABASE)
	if err != nil {
		t.Errorf("ConnectMongoDB :%v", err)
	}

	taskDAO := NewTaskDAOwithName(mgClient, DATABASE, COLNAME)

	for _, tt := range tests {
		if task, err := taskDAO.Delete(tt.arg1); err != nil {
			t.Errorf("%v", err)
		} else {
			fmt.Printf("task : %+v\n", task)
		}
	}
}

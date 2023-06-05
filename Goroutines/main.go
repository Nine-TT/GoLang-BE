package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

const (
	DB_UserName = "root"
	DB_PassWord = "2002"
	DB_Name     = "users_demo"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Gender    string
	IpAddress string
}

func ConnectDB() (*gorm.DB, error) {
	dsn := DB_UserName + ":" + DB_PassWord + "@tcp(localhost:3306)/" + DB_Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getUsers(db *gorm.DB, start, limit int, resultCh chan<- []User, wg *sync.WaitGroup) {
	var users []User

	db.Limit(limit).Offset(start).Find(&users)

	resultCh <- users

	wg.Done()
}

func collectResults(resultCh <-chan []User, userList *[]User, wg *sync.WaitGroup) {
	for users := range resultCh {
		*userList = append(*userList, users...)
	}

	wg.Done()
}

func main() {
	db, err := ConnectDB()

	if err != nil {
		fmt.Println(err)
	}

	var ListUser []User
	resultUser := make(chan []User)
	//defer close(resultUser)

	totalUsers := 1000
	batchSize := 200

	var wg sync.WaitGroup
	wg.Add(6)

	for i := 0; i <= totalUsers; i += batchSize {
		go getUsers(db, i, batchSize, resultUser, &wg)

	}

	//go getUsers(db, 0, 200, resultUser, &wg)
	//go getUsers(db, 201, 400, resultUser, &wg)
	//go getUsers(db, 401, 600, resultUser, &wg)
	//go getUsers(db, 601, 800, resultUser, &wg)
	//go getUsers(db, 801, totalUsers, resultUser, &wg)

	go collectResults(resultUser, &ListUser, &wg)

	wg.Wait()

	for _, user := range ListUser {

		fmt.Println(user.Id)
	}

	fmt.Println("len listuser :", len(ListUser))

}

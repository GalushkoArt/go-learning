package concurrency

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var actions = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	output := fmt.Sprintf("UID: %d; Email: %s;\nActivity Log:\n", u.id, u.email)
	for index, item := range u.logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.action, item.timestamp.Format(time.RFC3339))
	}

	return output
}

func Refactoring() {
	err := os.MkdirAll(filepath.Join(".", "users"), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().Unix())
	usersNumber, workerNumber := 100, 50
	usersChan := make(chan User, workerNumber)
	done := make(chan int, workerNumber)
	for i := 0; i < workerNumber; i++ {
		go usersWorker(i, usersChan, done)
	}

	startTime := time.Now()

	go generateUsers(usersNumber, usersChan)
	for i := 0; i < usersNumber; i++ {
		<-done
	}
	close(done)

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
	err = os.RemoveAll("users")
	if err != nil {
		log.Fatal(err)
	}
}

func saveUserInfo(user User) {
	fmt.Printf("WRITING FILE FOR UID %d\n", user.id)

	filename := fmt.Sprintf("users/uid%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
}

func generateUsers(count int, usersChan chan<- User) {
	waitGroup := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		i := i
		go func() {
			user := User{
				id:    i + 1,
				email: fmt.Sprintf("user%d@company.com", i+1),
				logs:  generateLogs(rand.Intn(1000)),
			}
			fmt.Printf("generated user %d\n", i+1)
			time.Sleep(time.Millisecond * 100)
			usersChan <- user
			waitGroup.Done()
		}()
	}
	go func() {
		waitGroup.Wait()
		close(usersChan)
	}()
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}

func usersWorker(id int, usersChan <-chan User, done chan<- int) {
	for u := range usersChan {
		saveUserInfo(u)
		fmt.Printf("Worker %d saved user %d\n", id, u.id)
		done <- 1
	}
}

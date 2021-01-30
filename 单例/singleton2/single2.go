package singleton2

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var once sync.Once
type single struct {}

var singleInstance *single

func GetInstance(wg *sync.WaitGroup) *single {
	defer wg.Done()
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}
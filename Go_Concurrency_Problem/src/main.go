package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ContinuumLLC/GoTest/src/utility"
)

func main() {

	//read the json File
	file := utility.ReadJsonFile()

	//sort the file
	sortedFile := utility.SortJsonValuesByTimeStamp(file)
	//	printFirstNameAndEmail(sortedFile)

	//start Index in usersSlice
	UserAtIndex := 0

	//number of times I need to spawn 10 routines
	SpawnCounter := len(sortedFile) / 10

	var wg sync.WaitGroup

	executeRoutinesAfterDelay(&UserAtIndex, sortedFile, SpawnCounter, &wg)

}

func executeRoutinesAfterDelay(userAtIndex *int, sortedFile []utility.User, spawnConter int, wg *sync.WaitGroup) {

	for i := 0; i < spawnConter; i++ {
		spawnTenGoRoutines(userAtIndex, sortedFile, wg)
		time.Sleep(10 * time.Second)
	}

}

func spawnTenGoRoutines(userAtIndex *int, sortedFile []utility.User, wg *sync.WaitGroup) {

	for i := 0; i < 10; i++ {
		if *userAtIndex >= len(sortedFile) {
			break
		}
		wg.Add(1)
		go printUser(*userAtIndex, sortedFile, wg)
		wg.Wait()
		*userAtIndex++
	}

}

func printUser(userAtIndex int, sortedFile []utility.User, wg *sync.WaitGroup) {

	fmt.Println("First Name : " + sortedFile[userAtIndex].FirstName)
	fmt.Println(" Email : " + sortedFile[userAtIndex].Email)
	fmt.Println("")
	fmt.Println("")
	wg.Done()

}

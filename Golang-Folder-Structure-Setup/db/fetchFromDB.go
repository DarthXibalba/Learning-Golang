package db

import "fmt"

// This can be called inside services
func FetchFromDB() {
	fmt.Println("Query DB to fetch data")
}

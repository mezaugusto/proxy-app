package middleware

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kataras/iris"
)

// Queue s
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Repository s
type Repository interface {
	Read() []*Queue
}

func (q *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			fmt.Println("Out")
			continue
		}
		fmt.Println("In:", scanner.Text())
	}
	return MockQueue()
}

// MockQueue is
func MockQueue() []*Queue {
	return []*Queue{
		{
			Domain:   "alpha",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "beta",
			Weight:   1,
			Priority: 2,
		},
		{
			Domain:   "omega",
			Weight:   3,
			Priority: 4,
		},
		{
			Domain:   "alpha",
			Weight:   4,
			Priority: 1,
		},
		{
			Domain:   "beta",
			Weight:   5,
			Priority: 1,
		},
		{
			Domain:   "alpha",
			Weight:   1,
			Priority: 6,
		},
	}
}

// Q empty Queue
var Q []string

// Deprecated InitQueue s
// func InitQueue() {
// 	Q = append(Q, &Queue{})
// }

// ProxyMiddleware extracts the domain from the header
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var repo Repository
	repo = &Queue{}
	fmt.Println("Header Domain:", domain)
	for _, row := range repo.Read() {
		fmt.Println("Source Domain:", row.Domain)
		// TODO: Priorization algorithm
	}
	Q = append(Q, domain)
	c.Next()
}

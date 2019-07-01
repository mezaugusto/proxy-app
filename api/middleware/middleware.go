package middleware

import (
	"fmt"

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
var Q []*Queue

// InitQueue s
func InitQueue() {
	Q = append(Q, &Queue{})
}

// ProxyMiddleware extracts the domain from the header
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	var repo Repository
	repo = &Queue{}
	fmt.Println("Header Domain:", domain)
	for _, row := range repo.Read() {
		fmt.Println("Source Domain:", row.Domain)
		// TODO: Priorization algorithm
	}
	c.Next()
}

package middleware

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	var requests []*Queue
	request := &Queue{}
	count := 0

	for scanner.Scan() {
		count++
		if scanner.Text() == "" {
			count = 0
			continue
		}
		switch count {
		case 1:
			request.Domain = scanner.Text()
		case 2:
			val := strings.Split(scanner.Text(), ":")[1]
			weight, _ := strconv.Atoi(val)
			request.Weight = weight
		case 3:
			val := strings.Split(scanner.Text(), ":")[1]
			priority, _ := strconv.Atoi(val)
			request.Priority = priority

			requests = append(requests, request)
			request = &Queue{}
		}
	}

	return requests
}

// Q empty Queue
var Q []string

var computedPriorities map[string]int

func getPriority(domain string) int {
	var repo Repository
	repo = &Queue{}
	priority := 0

	for _, row := range repo.Read() {
		if domain == row.Domain {
			if row.Weight >= 5 && row.Priority >= 5 {
				priority = 1
			} else if row.Weight < 5 && row.Priority < 5 {
				priority = 3
			} else {
				priority = 2
			}
			break
		}
	}
	computedPriorities[domain] = priority
	return priority
}

// ProxyMiddleware extracts the domain from the header
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	fmt.Println("Domain:", domain)

	if computedPriorities == nil {
		computedPriorities = make(map[string]int)
	}

	var priority int
	if val, ok := computedPriorities[domain]; ok {
		priority = val
	}
	priority = getPriority(domain)

	if priority == 0 {
		c.JSON(iris.Map{"status": 400, "result": "domain not valid"})
		return
	}

	for i, iDomain := range Q {
		iPriority := computedPriorities[iDomain]
		if priority <= iPriority {
			Q = append(Q, "")
			copy(Q[i+1:], Q[i:])
			Q[i] = domain
			c.Next()
			return
		}
	}

	Q = append(Q, domain)
	c.Next()
}

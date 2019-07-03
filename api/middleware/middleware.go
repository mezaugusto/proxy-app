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

func getPriority(w int, p int) string {
	if w >= 5 && p >= 5 {
		return "High"
	} else if w < 5 && p < 5 {
		return "Low"
	}
	return "Medium"
}

// ProxyMiddleware extracts the domain from the header
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	fmt.Println("Header Domain:", domain)

	var repo Repository
	repo = &Queue{}
	for _, row := range repo.Read() {
		fmt.Println("Source Domain:", row.Domain)
		if domain == row.Domain {
			fmt.Println(getPriority(row.Weight, row.Priority))
		}
	}
	Q = append(Q, domain)
	c.Next()
}

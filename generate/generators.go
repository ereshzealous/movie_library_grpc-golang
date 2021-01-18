package generate

import (
	"math/rand"

	"strconv"
	"strings"

	"github.com/google/uuid"
)

func randomID() string {
	return uuid.New().String()
}

func randomString() string {
	len := randomInt(2, 15)
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomYear() int {
	return 1930 + rand.Intn(2020-1930)
}

func randomBudget() string {
	budget := 1 + rand.Intn(500-1)
	return strings.Join([]string{strconv.Itoa(budget), "Millions"}, " ")
}

func randomStringArray() []string {
	index := randomInt(1, 20)
	value := make([]string, index)
	for i := 0; i < index; i++ {
		value = append(value, randomString())
	}
	return value
}

package util

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func FindMatch(key string, msgStr string) []string {
	reg := regexp.MustCompile(key)
	return reg.FindAllString(msgStr, -1)
}

func CheckRegexpMatch(key string, msgStr string) bool {
	match, _ := regexp.MatchString(key, msgStr)
	return match
}

func CheckWordExist(key string, msgStr string) bool {
	return strings.Contains(msgStr, key)
}

func IsInStringList(key string, List []string) bool {
	for _, k := range List {
		if key == k {
			return true
		}
	}
	return false
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// GetRandom get an int in the `min` to `max` range
func GetRandom(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return randomInt(min, max)
}

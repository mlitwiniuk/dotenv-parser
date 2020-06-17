package main

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("You must give full path to file as a param (can be multiple)")
	}
	for i := 1; i < len(os.Args); i++ {
		parseFile(os.Args[i])
	}
}

func parseFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = godotenv.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	f.Seek(0, 0)

	pattern := regexp.MustCompile(`^[#A-Z][A-Z0-9_]+\s+=\s['"]?.*['"]?\s*[a-zA-Z0-9_]+\s+=`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		loc := pattern.FindIndex([]byte(text))
		if len(loc) > 0 {
			log.Fatal(text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

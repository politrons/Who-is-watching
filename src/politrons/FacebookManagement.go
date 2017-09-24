package main

import (
	"fmt"
	"os/exec"
	"io/ioutil"
	"encoding/json"
	"strings"
)

const run_script = "sh facebook.sh"

func getVisitors() map[string]int {
	command(run_script)
	visitorsJson, _ := ioutil.ReadFile("visitors.json")
	return transformVisitorJsonToMap(visitorsJson)
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
	command(get_visitor_query)
}

func command(cmd string) []byte {
	data, error := exec.Command("/bin/bash", "-c", cmd).Output()
	checkError(error)
	return data
	//checkError(exec.Command("/bin/bash", "-c", cmd).Run())
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error occured")
		fmt.Printf("%s", err)
	}
}



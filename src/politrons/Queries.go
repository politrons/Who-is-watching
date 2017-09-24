package main

const cookie = "curl -X  GET 'https://www.facebook.com/home.php' --verbose --user-agent 'Firefox' --cookie 'cookies.txt' --cookie-jar 'cookies.txt' --location"
const login = "curl -X  POST 'https://login.facebook.com/login.php' --verbose --user-agent 'Firefox' --data-urlencode 'email=${EMAIL}' --data-urlencode 'pass=${PASS}' --cookie 'cookies.txt' --cookie-jar 'cookies.txt'"
const home = "curl -X  GET 'https://www.facebook.com/home.php' --verbose --user-agent 'Firefox' --cookie 'cookies.txt' --cookie-jar 'cookies.txt' > home.html"
const visitors = "perl -lne 'print $& if /(?<=lastActiveTimes:).*?(?=,chatNotif)/g' home.html > visitors.json"




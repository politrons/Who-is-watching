
#!/bin/bash

# If it redirects to http://www.facebook.com/login.php at the end, wait a few minutes and try again



COOKIES='cookies.txt'
USER_AGENT='Firefox'

curl -X  GET 'https://www.facebook.com/home.php' --verbose --user-agent $USER_AGENT --cookie $COOKIES --cookie-jar $COOKIES --location
curl -X  POST 'https://login.facebook.com/login.php' --verbose --user-agent $USER_AGENT --data-urlencode "email=${EMAIL}" --data-urlencode "pass=${PASS}" --cookie $COOKIES --cookie-jar $COOKIES
curl -X  GET 'https://www.facebook.com/home.php' --verbose --user-agent $USER_AGENT --cookie $COOKIES --cookie-jar $COOKIES > home.html
perl -lne 'print $& if /(?<=lastActiveTimes:).*?(?=,chatNotif)/g' home.html > visitors.json



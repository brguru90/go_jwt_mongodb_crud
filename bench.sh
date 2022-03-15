SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiOTBZTzNAWDFJOVkuY29tIiwiX2lkIjoiNjIzMGE5OGM5ZGY2YTBkZTk3YTk0NzFhIn0sInVuYW1lIjoiOTBZTzNAWDFJOVkuY29tIiwidG9rZW5faWQiOiI2MjMwYTk4YzlkZjZhMGRlOTdhOTQ3MWFfQWF1YUVzSzNPSEMzakZ6OE1taDlPdkw1cjBpd2JoaFh4WlM2WjZMWkJtY0VBeEpVRjFoNnRia0E4QnVBQ1NRVzBZYjVvWkViYTd1ekFXUEZJOE1acFROV0w2TDJQWkRhclhBQWZQN2NCOHlnNU1LL0cwcnA0QWpmYndXYjVPWE5PUGhoZHc9PV8xNjQ3MzYwMDg4NjE4IiwiZXhwIjoxNjQ3MzYzNjg4NjE4LCJpc3N1ZWRfYXQiOjE2NDczNjAwODg2MTgsImNzcmZfdG9rZW4iOiIifX0.bly13GdbW8HWg9uu_Pp_-StqyLf76_GmrJMZ-YAE3uQ"

csrf_token="JTZBtb/L+xH03dL1Oq4RiP+9JHQ7VNcVJ2+hFA+o0nDN90Cxzgf6B+fSyCnTSd9QUK6E4oMdFAxpPX+0xGNRjSI7I5j6YVpq+BFgKnjerfIyRCBcsibmAh3opNLzXzxdkeObjA=="

echo -e "\n====> /api/login_status/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/login_status/" 

echo -e "\n\n====>/api/user/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/user/" 

echo -e "\n\n\n\n\n\n\n\n"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

sleep 5

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

sleep 5

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

sleep 5

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"




SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiOTBZTzNAWDFJOVkuY29tIiwiX2lkIjoiNjIzMGE5OGM5ZGY2YTBkZTk3YTk0NzFhIn0sInVuYW1lIjoiOTBZTzNAWDFJOVkuY29tIiwidG9rZW5faWQiOiI2MjMwYTk4YzlkZjZhMGRlOTdhOTQ3MWFfTWJaVTMrdUhWQW03NHV6dy95dG1oa1lPK0VxeG95cUFmODdSZFNCT29iN0IzOHFDSjRJakZPSUU4OFQvMnU1SHllaVhkMVU5dm1QRmlGRFU5YVIvZWV4V3h1akhFVkpwWW1Xc1IveDlCMVlYeDI0V3NLeDVNRXN0VlVqbXZNMHpLY2QrL1E9PV8xNjQ3ODM4MDkxNTUxIiwiZXhwIjoxNjQ3ODQxNjkxNTUxLCJpc3N1ZWRfYXQiOjE2NDc4MzgwOTE1NTEsImNzcmZfdG9rZW4iOiIifX0.d04L_TlaQVEy_rtqGIFrUaF7IXQDwuhTcIoKe9RUJkw"

csrf_token="3qKbYxVJM9CJQ/eGmBD8xPUw0BdfwaRwYdES53NL39TbbpWpCwfCSxfwM2ZJpBH3tgsPBFB+yRNePC0b6p6bAgJkvmlKWm3RXVXDxVi0P6buh4l5oaDhUz5RywEyKt2pHvlz/Q=="

echo -e "\n====> /api/login_status/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/login_status/" 

echo -e "\n\n====>/api/user/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/user/" 

echo -e "\n\n\n\n\n\n\n\n"

ulimit -n 1000000

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 


ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"


ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"


ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"




SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiOTBZTzNAWDFJOVkuY29tIiwiX2lkIjoiNjIzMGE5OGM5ZGY2YTBkZTk3YTk0NzFhIn0sInVuYW1lIjoiOTBZTzNAWDFJOVkuY29tIiwidG9rZW5faWQiOiI2MjMwYTk4YzlkZjZhMGRlOTdhOTQ3MWFfVFVrNlZ1dzBzMzgvcTBFNU9FRzUwRkQxMURGV2ZWN0VWemJPeE1mMTJzSVdQbFNRbnA3UzVuREtRRktRRmtWVk5qeXFBQ0N1VjVXWkZlckx3UTdBa0VlZXFtUnBjYzhka2JIQ1JYY1ozaE9NL2tnaklxc0NzVnNHcUExaXhTckg2b24zSnc9PV8xNjQ3MzY3MDQ1MjM2IiwiZXhwIjoxNjQ3MzcwNjQ1MjM2LCJpc3N1ZWRfYXQiOjE2NDczNjcwNDUyMzYsImNzcmZfdG9rZW4iOiIifX0.J3RiOHTLVQo7DFtKgxOXImRgDkIy__JO6JV-oyeMBGs"

csrf_token="Kc8Q5HXak4NXRnpoHDwniHW9Bj6y291UX1IMgnkFo11wDOhiGHCOjgPhCJnZ5tg0qMXixI+bT4+Fkhg/E0Buv4UnRERAL6X8z/1EkSOaid5VHw1loAbcb97dmuAi1izDT0Nkgg=="

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




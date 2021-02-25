# User Rest API

## How to run the application

1. Create a mysql database

2. Use the mySQL dump `protectednet_user.sql` to create the tables and insert some dump data.

3. There are number of dependencies which need to be imported before running the application. Please get the dependenices through the following commands -

        go get "github.com/gorilla/mux"
        go get "github.com/stretchr/testify/assert"
        go get "github.com/go-sql-driver/mysql"
        go get "github.com/joho/godotenv"

4. Modify the `.env` file with the database name created and db URL + port

6. To run the application, please use the following commands -

        go build
		go_backend
		

## Endpoints Description

### Create User

```
    URL - *http://localhost:8080/createUser*
    Method - POST
	Body - (APIKey = eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV)
    {
    	"firstName":"John",
    	"lastName":"Smith",
    	"username":"jsmith",
    	"darkMode":true
    }
```
curl -H "APIKey: eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV" -X POST http://localhost:8080/createUser -d "{\"firstName\":\"John\", \"lastName\": \"Smith\" ,  \"username\": \"jsmith\" ,   \"darkMode\": true }"

### Update Name

```JSON
    URL - *http://localhost:8080/updateName*
    Method - PUT
	Body - (APIKey = eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV)
    {
    	"firstName":"John",
    	"lastName":"Smith",
    	"username":"jsmith"
    }
```

curl -H "APIKey: eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV" -X PUT http://localhost:8080/updateName -d "{\"firstName\":\"John\", \"lastName\": \"Snow\" ,  \"username\": \"jsmith\" }"

### Toggle DarkMode

```JSON
    URL - *http://localhost:8080/toggleDarkMode*
    Method - PUT
	Body - (APIKey = eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV)
    {
    	"username":"jsmith"
    }
```

curl -H "APIKey: eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV" -X PUT http://localhost:8080/toggleDarkMode -d "{\"username\": \"jsmith\" }"

### Delete User

```JSON
    URL - *http://localhost:8080/deleteUser"
    Method - DELETE
    Body - (APIKey = eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV)
    {
    	"username":"jsmith"
    }
```

curl -H "APIKey: eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV" -X DELETE http://localhost:8080/deleteUser -d "{\"username\": \"jsmith\" }"

### List Users

```JSON
    URL - *http://localhost:8080/listUsers*
    Method - GET
	Body - (APIKey = eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV)
```

curl -H "APIKey: eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV" -X GET http://localhost:8080/listUsers

### Search

```JSON
    URL - *http://localhost:8080/search*
    Method - GET
    Body - (APIKey = eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV)
    {
    	"searchString":"jsmith"
    }
```

curl -H "APIKey: eyJpdiI6IktMMjRUZ0hBWVMwN1dLWVEzV" -X GET http://localhost:8080/search -d "{\"searchString\": \"jsmith\" }"

## Test Driven Development Description

To run all the unit test cases, please do the following -

1. `go test -v`


## Hope everything works. Thank you.
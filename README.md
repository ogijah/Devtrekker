# Devtrekker
Further explanation how to start and test tasks:

# Task 4:

1) *docker-compose up*
2) testing the endpoints (by using postman for example):
http://localhost:9090/telephone (GET)- get all telephone numbers
http://localhost:9090/telephone/{id} (GET)- get telephone number by an Id (has to be an int number)
http://localhost:9090/telephone/{id} (DELETE)- delete telephone number by an Id
http://localhost:9090/telephone (POST) - upload a new number (it requires JSON object in the body)
3) starting the tests: position yourself in the data folder and then run: *go test -v* 


# Task 5:

1) In case bash.sh is not executable on your machine: *chmod +x bash.sh*
2) run *./bash.sh* - it will automatically run tests, build the main.go and run the main program with 
   a parameter from the command line

# glofox_task
A SaaS RESTful API platform for a studio manager to create classes and gets bookings from clients

## Prerequisites
You will need Go version 1.14+ installed on your development machine.
visit my article 
https://hashnode.com/post/the-proper-way-to-install-golang-and-set-up-your-workspace-ck8qhtjbu009oehs1nri77i5h to correctly download golang and setup your GOPATH~

## How to run the application

1. Clone the application with `https://github.com/AlexSwiss/glofox_task.git`

2. Download the postman configuration file (optional).

3. There are number of dependencies which need to be imported before running the application. Please get the dependenices through the following commands -

    ```shell
        go get "github.com/julienschmidt/httprouter"
    ```

6. To run the application, please use the following command -

    ```shell
        go build
        ./glofox_task
    ```
> Note: By default the port number is running on **8080**.

## Endpoints Description

### Create New Class

```JSON
    URL - *http://localhost:8080/api/classes*
    Method - POST
    Body - (content-type = application/json)
    {
    	"name":"Workout 101",
    	"startdate":"01 July 2020",
    	"enddate":"01 August 2020",
    	"capacity":"10"
    }
```

### Get All Classes

```
    URL - *http://localhost:8080/api/classes_list*
    Method - GET
```

### Add New Booking

```JSON
    URL - *http://localhost:8080/api/bookings*
    Method - POST
    Body - (content-type = application/json)
    {
    	"name":"Alexander Swiss",
    	"date":"01 July 2020"
    }
```
### Get All Bookings

```
    URL - *http://localhost:8080/api/bookings_list*
    Method - GET
```

## Test Driven Development Description

To run all the unit test cases, please do the following -

1. `go build`
2. `./glofox_task`
2. `go test -v`

## Hope everything works. Thank you.
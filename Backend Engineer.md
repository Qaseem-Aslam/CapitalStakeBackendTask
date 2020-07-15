# Backend Engineer Test (Golang)

This tasks is designed to help you work with concepts frequently used in backend systems development. We don't expect you to be fluent in language and concepts so take your time to familiarize yourself while attempting this task.

You will have 2 weeks to complete this task.

## Programming Language
[Go v1.14](https://golang.org/)


## Task

Write a **TCP Server** that loads dataset from a **CSV file** and provide interface to query the dataset. TCP server will expose port **4040**.

You are required to create a public repository on **GitHub.com**, place your source code and relevant dataset into that repository and share the public repository URL with us.

### Dataset

You will be required to download **Corona Virus Pakistan Dataset 2020** from the **OpenData.com.pk**, convert the dataset into a CSV file and use this file to load data into the server.

- Download `COVID_FINAL_DATA.xlsx` from [OpenData.com.pk](https://opendata.com.pk/dataset/corona-virus-pakistan-dataset-2020)
- Convert first sheet "TimeSeries_KeyIndicators" into a CSV file
- Place the CSV file into the repository

Following columns are required in CSV:
- Cumulative Test Positive
- Cumulative Tests Performed
- Date
- Discharged
- Expired
- Admitted
- Region

### Server Communication

User should be able to connect to the server using **NetCat** `nc localhost 4040` command on Linux/Unix based systems.

Once connected to TCP, user should be able communicate with the application by sending queries in JSON format.

```
{
	"query": {
		"region": "Sindh"
	}
}
```

```
{
	"query": {
		"date": "2020-03-20"
	}
}
```

User can query data based on two fields: Region and Date.

In response, server will return a **list** of records that will match query. An individual record's JSON format will look like:

```
{
   "date": "2020-03-11",
   "positive": 2,
   "tests": 171,
   "expired": 0,
   "admitted": 13,
   "discharged": 0,
   "region": "Sindh"
}
```


## Query Examples

```
> nc localhost 4040
> {"query": {"region": "Sindh"}}
> {"response": [
   {
	   "date": "2020-03-11",
	   "positive": 2,
	   "tests": 171,
	   "expired": 0,
	   "admitted": 13,
	   "discharged": 0,
	   "region": "Sindh"
   },
   {
	   "date": "2020-03-12",
	   "positive": 14,
	   "tests": 324,
	   "expired": 0,
	   "admitted": 13,
	   "discharged": 0,
	   "region": "Sindh"
   },
   {
	   "date": "2020-03-13",
	   "positive": 43,
	   "tests": 324,
	   "expired": 0,
	   "admitted": 12,
	   "discharged": 0,
	   "region": "Sindh"
   }
  ]}
```

```
> nc localhost 4040
> {"query": {"date": "2020-03-20"}}
> {"response": [
   {
	   "date": "2020-03-20",
	   "positive": 2,
	   "tests": 171,
	   "expired": 0,
	   "admitted": 13,
	   "discharged": 0,
	   "region": "Sindh"
   },
   {
	   "date": "2020-03-20",
	   "positive": 14,
	   "tests": 324,
	   "expired": 0,
	   "admitted": 13,
	   "discharged": 0,
	   "region": "Punjab"
   },
   {
	   "date": "2020-03-20",
	   "positive": 43,
	   "tests": 324,
	   "expired": 0,
	   "admitted": 12,
	   "discharged": 0,
	   "region": "KP"
   }
  ]}
```


## Submission

You are required to create a public repository on **GitHub.com**, place your source code and relevant dataset into that repository and share the public repository URL with us.

## Evaluation
We will evaluate your submission on following basis:
- Task Completion
- Code Organization
- Data Structures

## Learning Resources

[Effective Go](https://golang.org/doc/effective_go.html)

[Reading a simple CSV in Go](https://medium.com/@ankurraina/reading-a-simple-csv-in-go-36d7a269cecd)

[Read a CSV File into a Struct](https://golangcode.com/how-to-read-a-csv-file-into-a-struct/)

[Go by Example: JSON](https://gobyexample.com/json)

[Network Programming with Go: A TCP Server with a Custom Protocol](https://www.youtube.com/watch?v=yW1ltZidh7g)

[How to Use Netcat Commands: Examples and Cheat Sheets](https://www.varonis.com/blog/netcat-commands/)

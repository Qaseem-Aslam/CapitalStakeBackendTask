package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	coronaData "./lib"
)

var (
	data_list = coronaData.GetData("corona_data.csv")
)


func main() {
	// setup flags
	var addr string
	var network string
	flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
	flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
	flag.Parse()

	// validate supported network protocols
	switch network {
	case "tcp", "tcp4", "tcp6", "unix":
	default:
		fmt.Println("unsupported network protocol")
		os.Exit(1)
	}

	// create a listener for provided network and host address
	ln, err := net.Listen(network, addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer ln.Close()
	log.Println("--------Corona Data TCP Server--------")
	log.Printf("Server Running: (%s) %s\n", network, addr)

	// connection loop3
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}
		log.Println("Connected to ", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

// handle client connection
func handleConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error closing connection:", err)
		}
	}()

	fmt.Fprintf(conn, "Connected to server\n")
	fmt.Fprintf(conn, "Example Query for fetching data by Date: {\"query\":{\"date\":\"MM/DD/YYYY\"}}\n")
	fmt.Fprintf(conn, "By region: {\"query\":{\"region\":\"region_name\"}\n")

	reader := bufio.NewReaderSize(conn, 256)

	// command-loop
	for {
		// reader will read bytes until enter is pressed
		buf, err := reader.ReadSlice('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("connection read error:", err)
				return
			}
		}
		reader.Reset(conn)

		// unmarshal request into value of type coronaData.DataRequest
		var req coronaData.DataRequest
		if err := json.Unmarshal(buf, &req); err != nil {
			log.Println("failed to unmarshal request:", err)
			// Sending error msg as JSON
			cerr, jerr := json.Marshal(&coronaData.DataError{Error: err.Error()})
			if jerr != nil {
				log.Println("failed to marshal DataError:", jerr)
				continue
			}

			if _, werr := conn.Write(cerr); werr != nil {
				log.Println("failed to write to DataError:", werr)
				return
			}
			continue
		}

		// search data_list, result is []coronaData.DataObject
		result := coronaData.Find(data_list, req.Query)
		if result == nil {
			fmt.Fprintf(conn, "Invalid query! Cannot query for both region and date.\n")
			continue
		}
		// marshal result to JSON array
		resp, err := json.MarshalIndent(&result,"","    ")
		if err != nil {
			log.Println("failed to marshal data:", err)
			// Note fmt.Fprintf prints a raw JSON object directly to the client without encoding.
			if _, err := fmt.Fprintf(conn, `{"data_error":"internal error"}`); err != nil {
				log.Printf("failed to write to client: %v", err)
				return
			}
			continue
		}

		// write response to client
		if _, err := conn.Write(resp); err != nil {
			log.Println("failed to write response:", err)
			return
		}
	}
}

package coronaDatalib

import (
	"encoding/csv"
	"io"
	"os"
	//"fmt"
	"strings"
	"strconv"
)

type DataObject struct {
	CumulativeTestsPositive     int `json:"cumulative_test_positive"`
	CumulativeTestsPerformed    int `json:"cumulative_tests_performed"`
	Date 						string `json:"date"`
	Discharged					int `json:"discharged"`
	Expired						int `json:"expired"`
	Region						string `json:"region"`
	StillAdmitted				int `json:"still_admitted"`
}

type Filter struct{
	Date string `json:"date,omitempty"`
	Region string `json:"region,omitempty"`
}

type DataRequest struct {
	Query Filter `json:"query"`
}

type DataError struct {
	Error string `json:"data_error"`
}

func GetData(path string) []DataObject {
	data_list := make([]DataObject, 0)
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err.Error())
		}
		//ignoring error as data is correctly formated
		testPositive, _  	:= strconv.Atoi(row[0])
		testPerformed, _ 	:= strconv.Atoi(row[1])
		discharged, _ 		:= strconv.Atoi(row[3])
		expired, _ 			:= strconv.Atoi(row[4])
		stillAdmitted, _ 	:= strconv.Atoi(row[6])

		c := DataObject{
			CumulativeTestsPositive:	testPositive,
			CumulativeTestsPerformed:	testPerformed,
			Date:						row[2],
			Discharged:					discharged,
			Expired: 					expired,
			Region:						row[5],
			StillAdmitted:				stillAdmitted,
		}

		data_list = append(data_list, c)
	}
	return data_list
}

func Find(data_list []DataObject, query Filter) []DataObject {
	result := make([]DataObject, 0)
	if query.Date != "" && query.Region != "" {
		return nil
	}else if query.Date != "" {
		for _, data := range data_list {
			if data.Date == query.Date {
				result = append(result, data)
			}
		}
	}else {
		for _, data := range data_list {
			if strings.ToUpper(data.Region) == strings.ToUpper(query.Region) {
				result = append(result, data)
			}
		}
	}
	return result
}

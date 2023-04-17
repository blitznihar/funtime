package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_readCSVfromS3(t *testing.T) {
	arn := "arn:aws:s3:::source042023/segments.csv"
	objectURL := "https://source042023.s3.us-east-2.amazonaws.com/segments.csv"
	fmt.Println(arn)
	fmt.Println(objectURL)
	type args struct {
		s3SourceLocation string
		s3SourceBucket   string
		csvFileName      string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readCSVfromS3(tt.args.s3SourceLocation, tt.args.s3SourceBucket, tt.args.csvFileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCSVfromS3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeCsvToS3(t *testing.T) {
	type args struct {
		s3DestinationLocation string
		s3DestinationBucket   string
		csvFileName           string
		data                  []TestResults
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := writeCsvToS3(tt.args.s3DestinationLocation, tt.args.s3DestinationBucket, tt.args.csvFileName, tt.args.data); got != tt.want {
				t.Errorf("writeCsvToS3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_callTestDataAPI(t *testing.T) {

	apiUrl := "https://2207fef8-d751-4ebc-bc33-8e23a591452f.mock.pstmn.io/getTestResults"
	token := ""
	segmentList := []string{"string", "stirng2"}
	want := "[{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1},{\"jittermin\":0,\"jittermax\":1}]"

	got := callTestDataAPI(apiUrl, token, segmentList)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("callTestDataAPI() = %v, want %v", got, want)
	}
}

func Test_convertJsonToList(t *testing.T) {
	type args struct {
		json string
	}
	jsonString := "[{\"jittermin\": 100, \"jittermax\": 102}, {\"jittermin\": 200, \"jittermax\": 202}]"
	want := []TestResults{{Jittermin: 100, Jittermax: 102}, {Jittermin: 200, Jittermax: 202}}
	got := convertJsonToList(jsonString)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("convertJsonToList() = %v, want %v", got, want)
	}
	fmt.Println(got)
	fmt.Println(want)

}

func Test_convertStructToFlatCsvData(t *testing.T) {
	type args struct {
		testResults []TestResults
	}
	input := []TestResults{{Jittermin: 100, Jittermax: 102}, {Jittermin: 200, Jittermax: 202}}
	want := [][]string{{"jittermin", "jittermax"}, {"100", "102"}, {"200", "202"}}
	got := convertStructToFlatCsvData(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("convertStructToFlatCsvData() = %v, want %v", got, want)
	}
}

func Test_writeCSV(t *testing.T) {
	filename := "Result.csv"
	data := [][]string{{"jittermin", "jittermax"}, {"100", "102"}, {"200", "202"}}
	got := writeCSV(data, filename)
	want := true
	if got != want {
		t.Errorf("writeCSV() = %v, want %v", got, want)
	}
}

func TestDownloadFile(t *testing.T) {
	bucketName := "source042023"
	objectKey := ""
	fileName := "segment.csv"

	err := DownloadFile(bucketName, objectKey, fileName)
	if err != nil {
		t.Errorf("DownloadFile() error = %v, wantErr %v", err, "")
	}

}

func TestUploadFile(t *testing.T) {

	bucketName := "destination042023"
	fileName := "testrecords"
	body := [][]string{{"jittermin", "jittermax"}, {"100", "102"}, {"200", "202"}}
	err := UploadFile(bucketName, fileName, body)
	if err != nil {
		t.Errorf("UploadFile() error = %v", err)
	}
}

package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	fmt.Println("Hello World!!!")

}

type TestResults struct {
	Jittermin int
	Jittermax int
}

var sess *session.Session

// TODO: Read CSV file from S3
// Method
// Input: S3 location
// Output: if file exists read the connect and result as list
// bucketName := "destination042023"
// fileName := "testrecords.csv"
// objectkey := ""

// bucketName := "source042023"
// fileName := "segment.csv"
// objectkey := ""
func DownloadFile(bucketName string, objectKey string, fileName string) error {

	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("AKIASS3CL3LCQWIAU55D", "1vai8n4KLxVd8zPDwsAaILEDKy3Kppfy+yTH2AZo", ""),
	})

	s3Client := s3.New(awsSession)

	getInput := &s3.GetObjectInput{
		Bucket: aws.String("source042023"),
		Key:    aws.String("segments.csv"),
	}

	resp, err := s3Client.GetObjectWithContext(context.TODO(), getInput)

	// if err != nil {
	// 	return nil, fmt.Errorf("error downloading file: %v", err)
	// }
	defer resp.Body.Close()
	rd, err := ioutil.ReadAll(resp.Body)
	rdString := string(rd)
	fmt.Println(rdString)
	return err
}

func UploadFile(bucketName string, fileName string, body [][]string) error {

	var strs []string
	for _, v1 := range body {
		s := strings.Join(v1, ", ")
		s = s + "\n"
		strs = append(strs, s)
	}
	bodyToUpload := strings.Join(strs, ", ")
	myReader := strings.NewReader(bodyToUpload)
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("AKIASS3CL3LCQWIAU55D", "1vai8n4KLxVd8zPDwsAaILEDKy3Kppfy+yTH2AZo", ""),
	})

	uploader := s3manager.NewUploader(awsSession)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("destination042023"),
		Key:    aws.String("testrecords" + strings.Replace(strings.Replace(strings.Replace(strings.Replace(time.Now().UTC().Format(time.RFC3339), ":", "", -1), "-", "", -1), "Z", "", -1), "T", "", -1)),
		Body:   myReader,
	})
	if err != nil {
		// Print the error and exit.
		log.Fatal(err)
	}

	return err
}

func readCSVfromS3(s3SourceLocation string, s3SourceBucket string, csvFileName string) []string {
	var segments []string

	return segments
}

// TODO: write CSV file to S3
// Method
// Input: S3 Location, filename, List write to CSV
// Output: success
func writeCsvToS3(s3DestinationLocation string, s3DestinationBucket string, csvFileName string, data []TestResults) bool {

	return false
}

// TODO: call API to fetch JSON records
// Method
// Input: API URL, Input Parameters, 0Auth Token
// Output: List of Records
func callTestDataAPI(apiUrl string, token string, segmentList []string) string {
	response, err := http.Get(apiUrl)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	strResponseData := string(responseData)
	fmt.Println(strResponseData)
	return strResponseData
}

//TODO get 0Auth Token
// Method
// Input: Credentials
// Output: OAuth Tokens

// TODO: Convert StructToFlatenedCSV
// Method
// Input: testResults
// Output: 2 dimensional string array
func convertStructToFlatCsvData(testResults []TestResults) [][]string {
	var data [][]string
	headerRow := []string{"jittermin", "jittermax"}
	data = append(data, headerRow)
	for _, record := range testResults {
		row := []string{strconv.Itoa(record.Jittermin), strconv.Itoa(record.Jittermax)}
		data = append(data, row)
	}
	return data
}

// TODO: Convert JSON to CSV
// Method
// Input: JSON
// Output: Flattened
func convertJsonToList(jsonString string) []TestResults {
	var testResults []TestResults
	json.Unmarshal([]byte(jsonString), &testResults)

	return testResults
}

// TODO: For Testing CSV Files
func writeCSV(data [][]string, filename string) bool {
	flag := true
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	err = w.WriteAll(data) // calls Flush internally

	if err != nil {
		log.Fatal(err)
		flag = false
	}
	return flag
}

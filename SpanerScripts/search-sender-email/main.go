package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	apiURL = ""
	token  = os.Getenv("SPANNER_TOKEN")
)

func main() {
	sql := `select description,reference_key,price from sku where cost_type = 'sms_credits' and deactivated_at is null limit 999;`
	columnData, err := getData(sql)
	header := [][]string{{"description", "reference_key", "price"}}
	rows := lo.Map(columnData, func(item ColumnData, _ int) []string {
		return []string{item["description"], item["reference_key"], item["price"]}
	})
	if err != nil {
		panic(err)
	}
	WriteDataToCSV("./sku_data.csv", header, rows)
}

func getData(sql string) (data []ColumnData, err error) {
	jsonStr, err := json.Marshal(map[string]string{
		"db_type":     "spanner",
		"gcp_project": "aftership-pro",
		"query":       sql,
		"query_mode":  "profile",
	})
	if err != nil {
		panic(err)
	}
	resp, err := sendPostRequest(apiURL, jsonStr, map[string]string{
		"Connection":     "keep-alive",
		"Content-type":   "application/json",
		"Origin":         "",
		"Referer":        "",
		"Sec-Fetch-Dest": "empty",
		"Sec-Fetch-Mode": "cors",
		"Sec-Fetch-Site": "same-site",
		"Authorization":  token,
	})
	if err != nil {
		panic(err)
	}
	// read response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[read response body error]:" + string(bodyBytes))
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("[response code]:" + resp.Status + " [response body]:" + string(bodyBytes))
		return nil, fmt.Errorf("response code is not 200")
	}

	// resolve response data to [][]string
	var res Response
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		panic(err)
	}

	return res.Data.Data, nil
}

func ReadDataFromCSV(path string) []string {
	// 打开CSV文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		panic(err)
	}
	defer file.Close()

	// 创建CSV阅读器
	reader := csv.NewReader(file)

	var orgIDs []string
	// 读取并输出CSV内容
	for {
		record, err := reader.Read()
		if err != nil {
			// 到达文件末尾时退出循环
			break
		}

		// 处理CSV行数据
		for _, value := range record {
			orgIDs = append(orgIDs, value)
		}
	}
	return orgIDs
}

func WriteDataToCSV(path string, head [][]string, data [][]string) {
	// 创建或打开CSV文件
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("无法创建文件:", err)
		panic(err)
	}
	defer file.Close()

	// 创建CSV写入器
	writer := csv.NewWriter(file)

	// 写入CSV内容
	var allData [][]string
	allData = append(allData, head...)
	allData = append(allData, data...)

	for _, row := range allData {
		err := writer.Write(row)
		if err != nil {
			fmt.Println("无法写入数据:", err)
			panic(err)
		}
	}

	// 刷新缓冲区，确保所有数据被写入文件
	writer.Flush()

	if err := writer.Error(); err != nil {
		fmt.Println("写入错误:", err)
		panic(err)
	}

	fmt.Println("数据已成功写入CSV文件")
}

func sendPostRequest(url string, jsonStr []byte, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return client.Do(req)
}

type ColumnData map[string]string

type Response struct {
	Data struct {
		Data []ColumnData `json:"data"`
	} `json:"data"`
}

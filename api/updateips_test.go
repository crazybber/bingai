package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

// 定义一个结构体，表示从API返回的数据的格式
type IPData struct {
	Data [][]string `json:"data"`
}

// 定义一个结构体，表示IP范围
type IPRange struct {
	Start string // 起始IP
	End   string // 结束IP
	Count int    // IP数量
}

// 定义一个函数，获取IP范围
func getIPRange() ([]IPRange, error) {
	// 构造API的URL，使用当前时间戳作为参数
	url := fmt.Sprintf("https://cdn-lite.ip2location.com/datasets/US.json?_=%d", time.Now().Unix())
	// 发起GET请求，获取响应
	resp, err := http.Get(url)
	if err != nil {
		return nil, err // 如果有错误发生，返回错误
	}
	defer resp.Body.Close() // 延迟关闭响应体
	// 读取响应体的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err // 如果有错误发生，返回错误
	}
	// 定义一个IPData变量，用于存储解析后的数据
	var data IPData
	// 使用json包的Unmarshal函数，将响应体的内容解析为IPData结构体
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err // 如果有错误发生，返回错误
	}
	// 定义一个切片，用于存储过滤后的IP范围
	var ranges []IPRange
	// 定义一个常量，表示过滤的阈值
	const limitCount = 10000
	// 遍历data.Data中的每个元素，它是一个字符串切片，表示一行数据
	for _, row := range data.Data {
		// 将第三个元素（IP数量）转换为整数，并去掉逗号分隔符
		count, err := strconv.Atoi(strings.ReplaceAll(row[2], ",", ""))
		if err != nil {
			return nil, err // 如果有错误发生，返回错误
		}
		// 如果IP数量大于等于阈值
		if count >= limitCount {
			// 创建一个IPRange结构体，将第一个元素（起始IP）和第二个元素（结束IP）赋值给它，并将count赋值给它
			r := IPRange{
				Start: row[0],
				End:   row[1],
				Count: count,
			}
			// 将r追加到ranges切片中
			ranges = append(ranges, r)
		}
	}
	return ranges, nil // 返回ranges切片和nil表示没有错误发生
}

func TestRefreshIP(t *testing.T) {
	ranges, err := getIPRange() // 调用getIPRange函数，获取结果和错误
	if err != nil {             // 如果有错误发生
		fmt.Println(err) // 打印错误并退出程序
		return
	}
	// 定义一个期望的结果，是一个IPRange切片
	expected := []IPRange{
		{"1.0.0.0", "1.0.0.255", 10000},  //100000
		{"1.0.1.0", "1.0.3.255", 200000}, //200000
		{"1.0.4.0", "1.0.7.255", 300000}, //300000
	}
	// 使用len函数，比较ranges和expected的长度是否相等
	if len(ranges) != len(expected) {
		t.Logf("got %d ranges, want %d", len(ranges), len(expected)) // 如果不相等，使用t.Logf说明
	}
	for _, r := range ranges { // 遍历ranges切片中的每个元素，它是一个IPRange结构体
		t.Logf("{\"%s\", \"%s\"}, //%d\n", r.Start, r.End, r.Count) // 按照指定的格式打印它的字段值
	}
}

// 定义一个单元测试函数，以Test开头，接受一个*testing.T类型的参数
func TestGetIPRange(t *testing.T) {
	ranges, err := getIPRange() // 调用getIPRange函数，获取结果和错误
	if err != nil {             // 如果有错误发生
		t.Fatal(err) // 使用t.Fatal方法，打印错误并终止测试
	}
	// 定义一个期望的结果，是一个IPRange切片
	expected := []IPRange{
		{"1.0.0.0", "1.0.0.255", 10000},  //100000
		{"1.0.1.0", "1.0.3.255", 200000}, //200000
		{"1.0.4.0", "1.0.7.255", 300000}, //300000
	}
	// 使用t.Log方法，打印出控制台数据
	t.Log("ranges:", ranges)
	t.Log("expected:", expected)
	// 使用len函数，比较ranges和expected的长度是否相等
	if len(ranges) != len(expected) {
		t.Logf("got %d ranges, want %d", len(ranges), len(expected)) // 如果不相等，使用t.Logf说明
	}
	// 使用for循环，遍历ranges和expected的每个元素，并比较它们的字段值是否相等
	for i := range ranges {
		if ranges[i].Start != expected[i].Start {
			t.Errorf("got start %s, want %s", ranges[i].Start, expected[i].Start) // 如果不相等，使用t.Errorf方法，打印出错误信息
		}
		if ranges[i].End != expected[i].End {
			t.Errorf("got end %s, want %s", ranges[i].End, expected[i].End) // 如果不相等，使用t.Errorf方法，打印出错误信息
		}
		if ranges[i].Count != expected[i].Count {
			t.Errorf("got count %d, want %d", ranges[i].Count, expected[i].Count) // 如果不相等，使用t.Errorf方法，打印出错误信息
		}
	}
}

package action

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"orm_test_platform/utils"
	"strconv"
	"time"

	"github.com/toodm/boomer"
)

//增加记录
func Action_10001() {
	var (
		bodyData  map[string][]interface{}
		body      []byte
		respData  map[string]interface{}
		respBody  []byte
		err       error
		mothod    string = "POST"
		url       string = "http://127.0.0.1:7999/add/test123456/test/TbTestModel/"
		startTime int64  = boomer.Now()
		elapsed   int64
		funName   string = "Action_10001"
	)
	defer func() {
		elapsed = boomer.Now() - startTime
		if err2 := recover(); err2 != nil {
			fmt.Println(err2.(error).Error())
			boomer.Events.Publish("request_failure", "http", funName, elapsed, err2.(error).Error())
		} else {
			boomer.Events.Publish("request_success", "http", funName, elapsed, int64(10))
		}
	}()
	//捏请求数据
	rand.Seed(time.Now().UnixNano())
	rand_v1 := rand.Intn(10000)
	rand_v2 := rand.Intn(10000)
	bodyData = map[string][]interface{}{
		"values": []interface{}{map[string]string{"RoleGuid": strconv.Itoa(rand_v1), "TwoKey": strconv.Itoa(rand_v2)}},
	}
	body, err = json.Marshal(bodyData)
	if err != nil {
		panic(err)
	}
	respBody, err = utils.HttpRequest(mothod, url, nil, body)
	if err != nil {
		panic(err)
	}
	//根据返回定义数据结构反序列化
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok.....!!!....", respData)
	return
}

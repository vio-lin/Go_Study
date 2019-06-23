package main

import "fmt"

// 构造一个包含人信息的实体
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func nap_test() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	//存储数据
	personDB["1"] = PersonInfo{"12345", "tom", "street1"}
	personDB["2"] = PersonInfo{"2345", "tomc", "street2"}
	personDB["3"] = PersonInfo{"345", "tomcat", "street3"}
	personDB["4"] = PersonInfo{"45", "tomcat1", "street4"}

	person, OK := personDB["4"]

	if OK {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
}

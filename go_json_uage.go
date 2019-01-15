package main

import (
	"fmt"
	"encoding/json"
	"strconv"
)

/*
type EachPartition struct {
        Topic string
        Partition string
        Replicas  []string
}
*/

func main(){
	var arr = []string{"ch_mapp","2","216","659","660"}
	type EachPartition struct {
		Topic string `json:"topic"`
		Partition int `json:"partition"`
		Replicas  []int `json:"replicas"`
	}
	partitionNum,_ := strconv.Atoi(arr[1])
	replica_1,_ := strconv.Atoi(arr[2])
	replica_2,_ := strconv.Atoi(arr[3])
	replica_3,_ := strconv.Atoi(arr[4])
	eachpartition := EachPartition{
		Topic: arr[0],
		Partition: partitionNum,
		Replicas: []int{replica_1,replica_2,replica_3},
	}
	rs, err := json.Marshal(eachpartition)
	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(rs))
}

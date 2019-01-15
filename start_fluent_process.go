/*
Auther :chenglinguang
date: 2019-01-11
*/

package main

import (
    "os/exec"
    //"log"
    "io/ioutil"
    "fmt"
    "strings"
)

func main(){


	myFolder := "/etc/fluent"
	var fileConf map[int]string

	fileConf = listFile(myFolder)

	//start the fluentd process in this folder
        for _,file := range(fileConf){
		conf:= file[7:9]
		startFluent(conf)		
	}
	
}

//get the fluent config files and return a map with fluentd config file as value "fluent-01.conf"
func listFile(myFolder string) map[int]string{
	var i int = 0
	var fileConf map[int]string = map[int]string{}
	//var fileConf map[int]string = make(map[int]string)
	files,_:=ioutil.ReadDir(myFolder)
	for _,file := range files{
		if strings.Contains(file.Name(),"fluent") &&strings.HasPrefix(file.Name(),"fluent") &&strings.HasSuffix(file.Name(),"conf"){
			//fmt.Println(file.Name())
			fileConf[i] = string(file.Name())
			i = i+1
		}
	}
	return fileConf
}
//start fluent service with the fluentd config number
func startFluent(conf string){
	command := `./init_env.sh conf`
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("the process %s has been there!\n",conf)
	}else{
		fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	}
}

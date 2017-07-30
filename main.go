package main

import (
	"example/cloud/raport/cases"
)

func main() {
	config := cases.Config{}
	config.SetConfig("/home/acel/Desktop/deneme.xlsx")

	noneDevice := cases.NoneDeviceUser{}
	noneDevice.StartCase()
	//go noneDevice.NoneDeviceID(1)

	opertaionSys := cases.OperationSystem{}
	opertaionSys.StartCase()
	go opertaionSys.OperationSystem(1, 12)

	//a := ""
	//fmt.Scan(&a)
}

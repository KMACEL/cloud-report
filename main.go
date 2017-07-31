package main

import (
	"github.com/KMACEL/cloud-report/cases"
)

func main() {
	config := cases.Config{}
	config.SetConfig("/home/acel/Desktop/deneme.xlsx")
	config.StartCase()

	//noneDevice := cases.NoneDeviceUser{}
	//noneDevice.NoneDeviceID(1)

	opertaionSys := cases.OperationSystem{}
	opertaionSys.OperationSystem(1, 12)

	//a := ""
	//fmt.Scan(&a)
}

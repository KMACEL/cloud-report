package cases

import (
	"fmt"
	"github.com/KMACEL/cloud-report/lib"
)

//OperationSystem is
type OperationSystem struct{}

//StartCase is
func (n OperationSystem) StartCase() {
	if Path != "" {
		readLib.SetReadFilePath(Path)
	} else {
		fmt.Println("Please Enter Path")
	}
}

//OperationSystem is
func (n OperationSystem) OperationSystem(getDeviceID int, getOsVersion int) {
	var device string
	var deviceUser string
	writeLib := lib.CreateExel{}
	writeLib.Open()
	writeLib.Create("OperationSystemVersion", "OS")

	writeLib.CreateNewRow()
	writeLib.WriteByOne("DeviceID : ")
	writeLib.WriteByOne("None Operation System's Device ID")

	readLib.ReadColumnByTwo(getDeviceID, getOsVersion,
		func(arg string) {
			device = arg
		},
		func(arg1 string) {
			deviceUser = arg1
		}, func(arg2 string) {
			if deviceUser != "" && arg2 == "" {
				writeLib.CreateNewRow()
				writeLib.WriteByOne(device)
			}
		})

	fmt.Println("Finish OperationSys")
}

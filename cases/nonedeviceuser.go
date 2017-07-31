package cases

import (
	"fmt"
	"github.com/KMACEL/cloud-report/lib"
)

//NoneDeviceUser Nil Control
type NoneDeviceUser struct {
}

//StartCase StartNoneDevice
func (n NoneDeviceUser) StartCase() {
	readLib.SetReadFilePath(Path)
}

//NoneDeviceID Control
func (n NoneDeviceUser) NoneDeviceID(getDevice int) {
	var device string

	writeLib := lib.CreateExel{}
	writeLib.Open()
	writeLib.Create("NoneDeviceUser", "User")

	writeLib.CreateNewRow()
	writeLib.WriteByOne("DeviceID : ")
	writeLib.WriteByOne("NoneDeviceUser's Device ID")

	readLib.ReadColumnByOne(getDevice, func(arg2 string) {
		device = arg2
	},
		func(arg3 string) {
			if arg3 == "" {
				writeLib.CreateNewRow()
				writeLib.WriteByOne(device)
			}
		})

	fmt.Println("Finish NoneDeviceUser")
}

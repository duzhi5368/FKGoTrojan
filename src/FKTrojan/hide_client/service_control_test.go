package hide_client

import "testing"

var ()

func TestInstallService(t *testing.T) {
	testServiceInfo := &ServiceInfos[1]
	err := testServiceInfo.installService()
	t.Log(err)
}

func TestStartService(t *testing.T) {
	testServiceInfo := &ServiceInfos[1]
	t.Log(testServiceInfo.startService())
}
func TestStopService(t *testing.T) {
	testServiceInfo := &ServiceInfos[1]
	t.Log(testServiceInfo.stopService())
}
func TestRemoveService(t *testing.T) {
	testServiceInfo := &ServiceInfos[1]
	err := testServiceInfo.removeService()
	t.Log(err)
}

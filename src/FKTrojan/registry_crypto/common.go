package registry_crypto

var (
	BASEPATH1 = "SOFTWARE\\Trion Softworks\\"
	BASEPATH2 = "SOFTWARE\\Microsoft\\MediaPlayer\\PluginsUpgrade"
)

type RegistryKeyType string

var (
	UIDKEY      RegistryKeyType = "ID"
	CONFIGKEY   RegistryKeyType = "config"
	SERVERKEY   RegistryKeyType = "server"
	FROMEXEPATH RegistryKeyType = "path"
)

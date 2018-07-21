package antivirus_blocker

import "os"

func saveIni() error {
	err := saveWinAnvirIni(anvirIni)
	if err != nil {
		return err
	}
	return nil
}

func removeIni() error {
	os.Remove(anvirIni)
	return nil
}

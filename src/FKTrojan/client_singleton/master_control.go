package client_singleton

import "FKTrojan/hide_client"

func startMaster() error {
	ok, err := hide_client.StartingOrStopping()
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	hide_client.Start()
	return nil
}

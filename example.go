package main

import "github.com/gurupras/go-wireless/iwlib"
import "fmt"

func main() {
	networks, err := iwlib.GetWirelessNetworks("wlan0")
	if err != nil {
		fmt.Printf("Error retrieve wireless networks:", err)
		return
	}

	for _, network := range networks {
		fmt.Printf("SSID: %v\n", network.SSID)
	}
}

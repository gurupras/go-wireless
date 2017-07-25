package iwlib

// #cgo LDFLAGS: -liw
// #include <iwlib.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	//	"fmt"
	"unsafe"
)

type WirelessScanResult struct {
	SSID string `json:"SSID"`
	Mode int    `json:"mode"`
	Freq int    `json:"freq"`
	Key  string `json:"key"`
}

func GetWirelessNetworks(iface string) ([]*WirelessScanResult, error) {
	sock := C.iw_sockets_open()

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var iwrange C.struct_iw_range
	ok := (C.iw_get_range_info(sock, c_iface, &iwrange) >= 0)
	if !ok {
		return nil, errors.New("Error in iw_get_range_info")
	}

	var head C.struct_wireless_scan_head
	ok = (C.iw_scan(sock, c_iface, C.int(iwrange.we_version_compiled), &head) >= 0)
	if !ok {
		return nil, errors.New("Error in iw_scan")
	}

	results := make([]*WirelessScanResult, 0)

	result := head.result
	for result != nil {
		wsresult := &WirelessScanResult{}
		wsresult.SSID = C.GoString(&result.b.essid[0])
		if len(wsresult.SSID) > 0 {
			results = append(results, wsresult)
			//fmt.Printf("SSID: %v\n", wsresult)
		}
		wsresult.Freq = int(result.b.freq)
		wsresult.Mode = int(result.b.mode)
		keyChar := (*C.char)(unsafe.Pointer(&result.b.key[0]))
		wsresult.Key = C.GoString(keyChar)

		result = result.next
	}

	return results, nil
}

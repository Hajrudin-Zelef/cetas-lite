package metrics

import (
	"encoding/binary"
	"os"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modKernel32              = windows.NewLazySystemDLL("kernel32.dll")
	modIphlpapi              = windows.NewLazySystemDLL("iphlpapi.dll")
	procGlobalMemoryStatusEx = modKernel32.NewProc("GlobalMemoryStatusEx")
	procGetDiskFreeSpaceExW  = modKernel32.NewProc("GetDiskFreeSpaceExW")
	procGetSystemTimes       = modKernel32.NewProc("GetSystemTimes")
	procGetIfTable           = modIphlpapi.NewProc("GetIfTable")
)

// --- CPU via GetSystemTimes ---

type filetime struct {
	low  uint32
	high uint32
}

func (ft filetime) u64() uint64 { return uint64(ft.low) | uint64(ft.high)<<32 }

func cpuTimes() (idle, total uint64, ok bool) {
	var idleFT, kernelFT, userFT filetime
	ret, _, _ := procGetSystemTimes.Call(
		uintptr(unsafe.Pointer(&idleFT)),
		uintptr(unsafe.Pointer(&kernelFT)),
		uintptr(unsafe.Pointer(&userFT)),
	)
	if ret == 0 {
		return 0, 0, false
	}
	return idleFT.u64(), kernelFT.u64() + userFT.u64(), true
}

func cpuPercent() *float64 {
	i1, t1, ok := cpuTimes()
	if !ok {
		return nil
	}
	time.Sleep(120 * time.Millisecond)
	i2, t2, ok := cpuTimes()
	if !ok || t2 <= t1 {
		return nil
	}
	pct := 100 * (1 - float64(i2-i1)/float64(t2-t1))
	if pct < 0 {
		pct = 0
	}
	if pct > 100 {
		pct = 100
	}
	return fptr(pct)
}

// --- RAM via GlobalMemoryStatusEx ---

type memoryStatusEx struct {
	dwLength                uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func ramStat() *MemStat {
	var ms memoryStatusEx
	ms.dwLength = uint32(unsafe.Sizeof(ms))
	ret, _, _ := procGlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&ms)))
	if ret == 0 || ms.ullTotalPhys == 0 {
		return nil
	}
	return &MemStat{Used: ms.ullTotalPhys - ms.ullAvailPhys, Total: ms.ullTotalPhys}
}

// --- Disque via GetDiskFreeSpaceExW (volume du dossier home) ---

func diskStat() *MemStat {
	dir, err := os.UserHomeDir()
	if err != nil || dir == "" {
		dir = `C:\`
	}
	path, err := windows.UTF16PtrFromString(dir)
	if err != nil {
		return nil
	}
	var freeAvail, total, freeTotal uint64
	ret, _, _ := procGetDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(path)),
		uintptr(unsafe.Pointer(&freeAvail)),
		uintptr(unsafe.Pointer(&total)),
		uintptr(unsafe.Pointer(&freeTotal)),
	)
	if ret == 0 || total == 0 || freeTotal > total {
		return nil
	}
	return &MemStat{Used: total - freeTotal, Total: total}
}

// --- Reseau via GetIfTable (iphlpapi) ---
//
// Layout de MIB_IFROW (SDK Windows) :
// wszName[256] WCHAR = 512 octets, puis dwIndex(512), dwType(516),
// dwMtu(520), dwSpeed(524), dwPhysAddrLen(528), bPhysAddr[8](532),
// dwAdminStatus(540), dwOperStatus(544), dwLastChange(548),
// dwInOctets(552), ..., dwOutOctets(576), ..., bDescr[256](604).
// Taille d'une ligne : 860 octets.

const (
	ifRowSize    = 860
	offIfType    = 516
	offOperState = 544
	offInOctets  = 552
	offOutOctets = 576

	ifOperOperational = 1
	ifTypeLoopback    = 24
)

func netStat() *NetStat {
	var size uint32
	// Premier appel : recupere la taille requise.
	_, _, _ = procGetIfTable.Call(0, uintptr(unsafe.Pointer(&size)), 0)
	if size == 0 {
		return nil
	}
	buf := make([]byte, size)
	ret, _, _ := procGetIfTable.Call(
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&size)),
		0,
	)
	if ret != 0 {
		return nil
	}
	n := int(binary.LittleEndian.Uint32(buf[:4]))
	var rx, tx uint64
	for i := 0; i < n; i++ {
		base := 4 + i*ifRowSize
		if base+ifRowSize > len(buf) {
			break
		}
		ifType := binary.LittleEndian.Uint32(buf[base+offIfType : base+offIfType+4])
		oper := binary.LittleEndian.Uint32(buf[base+offOperState : base+offOperState+4])
		if ifType == ifTypeLoopback || oper != ifOperOperational {
			continue
		}
		rx += uint64(binary.LittleEndian.Uint32(buf[base+offInOctets : base+offInOctets+4]))
		tx += uint64(binary.LittleEndian.Uint32(buf[base+offOutOctets : base+offOutOctets+4]))
	}
	return &NetStat{RxBytes: rx, TxBytes: tx}
}

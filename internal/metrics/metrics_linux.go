package metrics

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// cpuTimes lit les compteurs CPU depuis /proc/stat.
func cpuTimes() (idle, total uint64, ok bool) {
	f, err := os.Open("/proc/stat")
	if err != nil {
		return 0, 0, false
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if !strings.HasPrefix(line, "cpu ") {
			continue
		}
		fields := strings.Fields(line)[1:]
		var vals []uint64
		for _, fd := range fields {
			v, err := strconv.ParseUint(fd, 10, 64)
			if err != nil {
				return 0, 0, false
			}
			vals = append(vals, v)
		}
		if len(vals) < 5 {
			return 0, 0, false
		}
		var tot uint64
		for _, v := range vals {
			tot += v
		}
		idle = vals[3]
		if len(vals) > 4 {
			idle += vals[4] // iowait compte comme inactif
		}
		return idle, tot, true
	}
	return 0, 0, false
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
	idleDelta := float64(i2 - i1)
	totalDelta := float64(t2 - t1)
	pct := 100 * (1 - idleDelta/totalDelta)
	if pct < 0 {
		pct = 0
	}
	if pct > 100 {
		pct = 100
	}
	return fptr(pct)
}

func ramStat() *MemStat {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil
	}
	defer f.Close()
	var total, avail uint64
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) < 2 {
			continue
		}
		v, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}
		v *= 1024 // kB -> octets
		switch fields[0] {
		case "MemTotal:":
			total = v
		case "MemAvailable:":
			avail = v
		}
	}
	if total == 0 {
		return nil
	}
	used := total - avail
	return &MemStat{Used: used, Total: total}
}

func diskStat() *MemStat {
	var st syscall.Statfs_t
	if err := syscall.Statfs("/", &st); err != nil {
		return nil
	}
	bsize := uint64(st.Bsize)
	total := st.Blocks * bsize
	free := st.Bavail * bsize
	if total == 0 || free > total {
		return nil
	}
	return &MemStat{Used: total - free, Total: total}
}

func netStat() *NetStat {
	f, err := os.Open("/proc/net/dev")
	if err != nil {
		return nil
	}
	defer f.Close()
	var rx, tx uint64
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		if strings.TrimSpace(parts[0]) == "lo" {
			continue
		}
		fields := strings.Fields(parts[1])
		if len(fields) < 9 {
			continue
		}
		r, err1 := strconv.ParseUint(fields[0], 10, 64)
		t, err2 := strconv.ParseUint(fields[8], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}
		rx += r
		tx += t
	}
	return &NetStat{RxBytes: rx, TxBytes: tx}
}

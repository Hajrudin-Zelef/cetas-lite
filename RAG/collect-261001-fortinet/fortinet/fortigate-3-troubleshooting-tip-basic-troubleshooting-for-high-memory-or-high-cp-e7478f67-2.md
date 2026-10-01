---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-basic-troubleshooting-for-high-memory-or-high-cp-e7478f67-2
title: "fortigate-3-troubleshooting-tip-basic-troubleshooting-for-high-memory-or-high-cp-e7478f67"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2019-11-19"]
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-basic-troubleshooting-for-high-memory-or-high-cp-e7478f67.md
source_anchor: ""
source_lines: [114, 203]
sha256: 5ca669051e13784a7df210292a361d1c22234fce22647db41bac1a7beb8ce7bf
---

# fortigate-3-troubleshooting-tip-basic-troubleshooting-for-high-memory-or-high-cp-e7478f67

323: 2019-11-19 18:47:02 scanunit=manager pid=207 str="AV database changed (1); restarting workers"
324: 2019-11-19 18:47:02 <00207> scanunit=manager str="Success loading anti-virus database."
325: 2019-11-19 18:51:36 the killed daemon is /bin/pyfcgid: status=0x0
Crash log interval is 3600 seconds 
Enable the Interface Bandwidth monitoring on the FortiGate Dashboard:
- Select in the left column Dashboard -> Status. Then select 'Add widget'.
- Select Interface Bandwidth.
- Select the interface that is used on the FortiGate.
- Go to the Dashboard to see the interfaces with the bandwidth usage widget (in this scenario: the WAN interface).
The purpose of Interface Bandwidth usage is to see whether there is high bandwidth on the FortiGate that is exceeding the supported traffic.
This information may be useful in figuring out the cause of High CPU or High Memory consumption.
Example:
Command 5:
diagnose hardware sysinfo memory
By using 'diagnose hardware system memory', all of the memory counters involved in the conserve mode and kernel conserve mode calculation can be seen.
Consider the following example:
diagnose hardware sysinfo memory
total: used: free: shared: buffers: cached: shm:
Mem: 260435968 146337792 114098176 0 221184 65974272 59985920
Swap: 0 0 0
MemTotal: 254332 kB
MemFree: 111424 kB
MemShared: 0 kB
Buffers: 216 kB
Cached: 64428 kB
SwapCached: 0 kB
Active: 26844 kB
Inactive: 37856 kB
HighTotal: 0 kB
HighFree: 0 kB
LowTotal: 254332 kB (2)
LowFree: 111424 kB (1)
SwapTotal: 0 kB
SwapFree: 0 kB
Explaining the value 'Cached, Active, Inactive' that may take significant memory.
Cached = Active + Inactive.
This is information cached by the FortiGate for its system (basically, I/O buffering). The inactive part is claimed back from the system when it requires more memory.
Command 6:
diagnose sys session stat
By using 'diagnose sys session stat', it is possible to view all detailed statistics about the session table on a FortiGate.
An increase in the number of sessions will automatically lead to higher slab memory usage, as the slab allocator is responsible for storing session-related information in the kernel.
Sample result:
FG101F-2 # diagnose sys session stat
misc info: session_count=24 setup_rate=0 exp_count=0 reflect_count=0 clash=0
memory_tension_drop=0 ephemeral=0/239104 removeable=0 extreme_low_mem=0
npu_session_count=0
nturbo_session_count=0
delete=0, flush=1, dev_down=52/767 ses_walkers=0
TCP sessions:
10 in ESTABLISHED state
firewall error stat:
error1=00000000
error2=00000000
error3=00000000
error4=00000000
tt=00000000
cont=00000000
ips_recv=00000000
policy_deny=000116b8
av_recv=00000000
fqdn_count=00000009
fqdn6_count=00000000
global: ses_limit=0 ses6_limit=0 rt_limit=0 rt6_limit=0
Command 7:
diagnose sys top-mem 10
By using 'diagnose sys top-mem <integer>', it is possible to view high-consuming processes with their corresponding pids. The < integer> can be set from 1 to 99. If set to 99, this command will show the sum of all 99 processes at the bottom of the list. This value should be close to what is visible under get hardware memory under Active. 
If it is not, there may be an issue where FortiOS spawns lots of duplicate processes, which slowly consume all available memory.
Sample result:
FG101F-2 # diagnose sys top-mem 99
node (2111): 81197kB
cid (2168): 48226kB
wad (2277): 37462kB
ipshelper (2159): 33024kB
httpsd (6298): 28507kB
wad (2275): 19904kB
wad (2269): 19840kB
cmdbsvr (1992): 17288kB
forticldd (2101): 16374kB
lnkmtd (2139): 15431kB
If the issue persists and assistance from the support team is required, open a ticket through the Support portal.
Execute the commands listed below, collect the outputs, and attach them to the ticket along with the configuration file.
get system status
get hardware status
diagnose hardware sysinfo memory
diagnose debug crashlog read
diagnose sys top-mem 99
diagnose system top 5 40 <----- To sort by high CPU, use the 'p' key. To sort by high memory, use the 'm' key.
get system performance status <----- Run it 3 times.
Run the capture for 2 minutes and then stop it with Ctrl + C.
Related article:

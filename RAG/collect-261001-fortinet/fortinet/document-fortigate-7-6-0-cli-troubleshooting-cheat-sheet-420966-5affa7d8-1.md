---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8-1
title: "document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "license", "memory"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8.md
source_anchor: ""
source_lines: [1, 104]
sha256: e463bdfb27e5814c4135adac6ebecc68ed3d4da7a21aa217882e389d7190c6cc
---

# document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8

Managed FortiAPs
CLI troubleshooting cheat sheet
This reference lists some important command line interface (CLI) commands that can be used for log gathering, analysis, and troubleshooting.
It provides a basic understanding of CLI usage for users with different skill levels. Exploring additional commands beyond the ones listed here to gain a comprehensive understanding of the CLI is recommended. 
|  | Real-time debugging is a CPU intensive task. Use it with caution. | 
Enable/Disable debugging
| Command | Description | 
|---|---|
| diagnose debug reset | Stop all the prior debugs that were enabled and running in the foreground or background. | 
| diagnose debug enable | Start printing debugs in the console. | 
| diagnose debug disable | Stop printing  debugs in the console. The debugs are still running in the background; use diagnose debug reset to completely stop them. | 
| diagnose debug duration 0 | Start debugging for infinite duration. By default, debug is set for 30 minutes. | 
System
System
| Command | Description | 
|---|---|
| get system status | Show system information. | 
| execute time | Show current system time. | 
| get system performance status | Show CPU and memory utilization. | 
| execute tac report | Execute TAC report used to open a support ticket with Fortinet Support. | 
| diagnose sys top {s} {n} {i} | Show a list of the first n processes every s seconds for i iterations.  | 
| diagnose debug crashlog read | Show system and application crashes. | 
| diagnose sys process pidof <daemon> | Show PID of the daemon that is running. The names of currently running daemons can be found using diagnose sys top . For example: diagnose sys process pidof httpsd | 
| diagnose sys kill 11 <pid> | Kill the PID with signal 11. | 
| diagnose sys session stat | Show session statistics. | 
| diagnose sys session exp-stat | Show expectation session statistics. | 
| diagnose sys vd list | Show virtual domain information and system statistics. | 
| diagnose sys cmdb info | Show information about the latest configuration change performed by the daemon. | 
| execute factoryreset [keepvmlicense] | Immediately reset to factory defaults and reboot. If keepvmlicense is specified (VM models only), the VM license is retained after reset. | 
| execute factoryreset-shutdown [keepvmlicense] | Immediately reset to factory defaults and shutdown. If keepvmlicense is specified (VM models only), the VM license is retained after reset. | 
| execute factoryreset2 [keepvmlicense] | Reset to factory default, except system settings, system interfaces, VDOMs, static routes, and virtual switches. If keepvmlicense is specified (VM models only), the VM license is retained after reset. | 
| diagnose debug config-error-log read | Show errors in the configuration file. | 
| diagnose snmp ip frags | Show fragmentation and reassembly information. | 
| diagnose sys process dump <PID> diagnose sys process pstack <PID> diagnose sys process trace <PID> | Show essential process related information for a particular process PID. | 
| diagnose sys mpstat {n} | Show CPU usage every n seconds. | 
| diagnose hardware sysinfo memory | Show system memory information. | 
| diagnose firewall packet distribution | Show packet distribution statistics. | 
| execute reboot | Reboot the device. | 
Hardware
| Command | Description | 
|---|---|
| diagnose hardware sysinfo interrupts | Show hardware interrupts statistics. | 
| diagnose hardware test suite all | Execute a hardware diagnostic test, also known as an HQIP test. | 
| diagnose hardware deviceinfo disk | Show disk information. | 
| diagnose sys flash list | Show flash partitions. | 
| execute disk list | Show available mounted disks. | 
| execute disk format <partition ref> | Format the referenced partition. | 
| diagnose disktest device <device> diagnose disktest block <block> diagnose disktest size <mb> diagnose disk test run | Execute a disk check to check if disk is faulty.  | 
| execute formatlogdisk | Format the log disk. | 
| diagnose hardware sysinfo cpu | Show CPU information. | 
| diagnose sys modem detect diagnose debug application modemd -1 diagnose debug enable | Detect the modem and start real-time debugging of the modem daemon. | 
FortiGuard
| Command | Description | 
|---|---|
| diagnose webfilter fortiguard statistics | Show rating cache and daemon statistics. | 
| diagnose debug rating | Show web filter rating server information. | 
| diagnose debug application update -1 diagnose debug enable | Start debugging for updated daemon to troubleshoot FortiGuard update issues. | 
| execute update-now | Execute the FortiGuard update manually. | 
| diagnose autoupdate status diagnose autoupdate versions | Show license information. | 
Session table
| Command | Description | 
|---|---|
| diagnose sys session filter <filter> | Set session table filters. | 
| diagnose sys session filter | Show session filters, if set. | 
| diagnose sys session list | Show session table after filtering. | 
| diagnose sys session clear | Clear the session table for the specified filter. | 
| diagnose firewall iprope list | Show FortiGate's internal firewall table. | 
Network diagnostics
| Command | Description | 
|---|---|
| execute ping-options {options} execute ping <x.x.x.x> | Ping IP address <x.x.x.x> using the specified options. | 
| execute ssh-options {options} execute ssh <x.x.x.x> | SSH to IP address <x.x.x.x> using the specified options. | 
| execute traceroute-options {options} execute traceroute <x.x.x.x> | Traceroute IP address <x.x.x.x> using the specified options. | 
| get system arp diagnose ip arp list | Show ARP entries. | 
| diagnose netlink brctl list | Show the names of all of the switches on the FortiGate. | 
| diagnose netlink brctl name host <switch-name> | Show the switching table of the specified switch. | 
| get system interface get sys interface physical | Show a summary of interface details, including IP address information. | 
| diagnose ip address list | Show IP address information. | 
| diagnose hardware deviceinfo nic <interface> get hardware nic <interface> | Show detailed interface information. | 
| get sys interface transceiver | Show connected transceivers. | 
Packet sniffer
| Command | Description | 
|---|---|
| diagnose sniffer packet <interface> <'filter'> <verbose> <count> <a\|l> | Execute the inbuilt packet sniffer, filtered on a particular interface with the specified filter. For more information, see Performing a sniffer trace or packet capture. | 
Debug flow
| Command | Description | 
|---|---|
| diagnose debug reset | Stop all the prior debugs that were enabled and running in the foreground or background. | 
| diagnose debug flow filter clear | Clear any IPv4 debug flow filters. | 
| diagnose debug flow filter6 clear | Clear any IPv6 debug flow filters. | 
| diagnose debug flow filter <filter> | Set a filter for running IPv4 traffic debug flows. | 
| diagnose debug flow filter6 <filter> | Set a filter for running IPv6 traffic debug flows. | 
| diagnose debug flow show function-name enable | Show the function name of the code that the traffic accesses. | 
| diagnose debug flow show iprope enable | Show which internal firewall policy that the traffic is going through. | 
| diagnose debug console timestamp enable | Start printing timestamps on debugs. | 
| diagnose debug flow trace start <n> | Show n lines of IPv4 debugs. | 
| diagnose debug flow trace start6 <n> | Show n lines of IPv6 debugs. | 
| diagnose debug enable | Start printing debugs in the console. | 
|  | For more detailed debug flow filter information, see Technical Tip: Using filters to review traffic traversing the FortiGate. | 
UTM
| Command | Description | 
|---|---|
| diagnose debug urlfilter <filter> diagnose debug application urlfilter -1 diagnose debug enable | Start real-time debugging for web filter traffic. | 
| diagnose debug enable diagnose test application urlfilter | List the web filter debug outputs. | 

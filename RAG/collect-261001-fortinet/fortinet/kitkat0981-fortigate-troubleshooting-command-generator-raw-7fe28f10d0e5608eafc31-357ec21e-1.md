---
id: collect-261001-fortinet/fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e-1
title: "kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["distribution", "license", "memory"]
source: docs/RAG/collect-261001-fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e.md
source_anchor: ""
source_lines: [1, 133]
sha256: 6bffa870dd71124dcf9ed9bb594ba6a8c3c76e7790c2f9663b9b279372866ad1
---

# kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e

FortiOS 7.6 Troubleshooting Cheat Sheet Fortinet Inc. 01-760-1051988-20251120
CLI troubleshooting cheat sheet
This reference lists some important command line interface (CLI)
commands that can be used for log gathering, analysis, and
troubleshooting.
It provides a basic understanding of CLI usage for users with different skill
levels. Exploring additional commands beyond the ones listed here to gain
a comprehensive understanding of the CLI is recommended.
Enable/Disable debugging
Command Description
diagnose debug reset Stop all the prior debugs that were
enabled and running in the
foreground or background.
diagnose debug enable Start printing debugs in the console.
diagnose debug disable Stop printing debugs in the console.
The debugs are still running in the
background; use diagnose debug
reset to completely stop them.
diagnose debug duration 0 Start debugging for infinite duration.
By default, debug is set for 30
minutes.
System
Command Description
get system status Show system information.
execute time Show current system time.
get system performance status Show CPU and memory utilization.
execute tac report Execute TAC report used to open a
support ticket with Fortinet Support.
diagnose sys top {s} {n} {i} Show a list of the first n processes
every s seconds for i iterations.
l Shift +C: Sort by highest CPU
l Shift + M: Sort by highest
memory
diagnose debug crashlog read Show system and application
crashes.
diagnose sys process pidof <daemon> Show PID of the daemon that is
running. The names of currently
running daemons can be found
using diagnose sys top.
For example: diagnose sys
process pidof httpsd
diagnose sys kill 11 <pid> Kill the PID with signal 11.
diagnose sys session stat Show session statistics.
diagnose sys session exp-stat Show expectation session statistics.
diagnose sys vd list Show virtual domain information and
system statistics.
diagnose sys cmdb info Show information about the latest
configuration change performed by
the daemon.
execute factoryreset [keepvmlicense] Immediately reset to factory
defaults and reboot.
If keepvmlicense is specified (VM
models only), the VM license is
retained after reset.
execute factoryreset-shutdown
[keepvmlicense]
Immediately reset to factory
defaults and shutdown.
If keepvmlicense is specified (VM
models only), the VM license is
retained after reset.
Command Description
execute factoryreset2 [keepvmlicense] Reset to factory default, except
system settings, system interfaces,
VDOMs, static routes, and virtual
switches.
If keepvmlicense is specified (VM
models only), the VM license is
retained after reset.
diagnose debug config-error-log read Show errors in the configuration file.
diagnose snmp ip frags Show fragmentation and
reassembly information.
diagnose sys process dump <PID>
diagnose sys process pstack <PID>
diagnose sys process trace <PID>
Show essential process related
information for a particular process
PID.
diagnose sys mpstat {n} Show CPU usage every n seconds.
diagnose hardware sysinfo memory Show system memory information.
diagnose firewall packet distribution Show packet distribution statistics.
execute reboot Reboot the device.
Hardware
Command Description
diagnose hardware sysinfo interrupts Show hardware interrupts statistics.
diagnose hardware test suite all Execute a hardware diagnostic test,
also known as an HQIP test.
diagnose hardware deviceinfo disk Show disk information.
diagnose sys flash list Show flash partitions.
execute disk list Show available mounted disks.
execute disk format <partition ref> Format the referenced partition.
diagnose disktest device <device>
diagnose disktest block <block>
diagnose disktest size <mb>
diagnose disk test run
Execute a disk check to check if disk
is faulty.
l <device>: Device to test
l <block>: Block size of each
read/write operation.
l <mb>: Test size limit for each
cycle
execute formatlogdisk Format the log disk.
diagnose hardware sysinfo cpu Show CPU information.
diagnose sys modem detect
diagnose debug application modemd -1
diagnose debug enable
Detect the modem and start real-
time debugging of the modem
daemon.
FortiGuard
Command Description
diagnose webfilter fortiguard
statistics
Show rating cache and daemon
statistics.
diagnose debug rating Show web filter rating server
information.
diagnose debug application update -1
diagnose debug enable
Start debugging for updated
daemon to troubleshoot FortiGuard
update issues.
execute update-now Execute the FortiGuard update
manually.
diagnose autoupdate status
diagnose autoupdate versions
Show license information.
Session table
Command Description
diagnose sys session filter <filter> Set session table filters.
diagnose sys session filter Show session filters, if set.


---
id: collect-261001-fortinet/fortinet/fortigate-ha-high-availability-68bd0e9c
title: "Fortinet FortiGate HA (High Availability): Detailed Guide"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2022-04-16"]
keywords: ["ethernet", "memory", "parameters"]
source: docs/RAG/collect-261001-fortinet/fortigate-ha-high-availability-68bd0e9c.md
source_anchor: ""
source_lines: [1, 76]
sha256: db14a26be69ce13d2312ede700aee55b33ffbdfb14c6772ea7a0c1008e4bde8b
---

# Fortinet FortiGate HA (High Availability): Detailed Guide

*Source : https://networkinterview.com/fortigate-ha-high-availability/*

High Availability (HA) is a feature of Firewalls in which two or more devices are grouped together to provide redundancy in the network. HA links and synchronises two or more devices. In FortiGate HA one device will act as a primary device (also called Active FortiGate). Active device synchronises its configuration with another device in the group. Other FortiGate devices are called Secondary or Standby devices.
There are two Fortigate HA modes available:
HA Protocol used by FortiGate Cluster to communicate. FGCP travels between FortiGate cluster devices over the heartbeat links and uses TCP port 703 with Ethernet type values:
TCP port 23 is used by FGCP for configuration synchronisation. Firewall cluster uses FGCP to elect the primary, synchronize configuration, discover another firewall that belongs to the same HA and detect failover when any of the HA device fails.
In Active/Passive, Primary Firewall performs below tasks:
Secondary Firewall performs below tasks:
Virtual IP addresses are assigned to heartbeat Interfaces based on the serial number of FortiGate Firewall
Cluster uses these virtual IP addresses to differentiate cluster members and update configuration changes in clustered devices.
Fortigate HA Configuration
Configuring Primary FortiGate for HA
1. Go to System ->Select HA
2. Select mode Active-Passive Mode
3. Once Active-Passive mode selected multiple parameters are required
4. Mode- Active/ Passive
5. Set Device Priority -200. More numerical value higher the priority. Here Priority is set 200, secondary devices must have lower numerical value than Primary Firewall.
6. Device Group– Group name must be the same for both primary and secondary devices. Here we have given the name HA-GROUP. Device Group is used in HA to assign two or more devices to be part of the same HA Group.
7. Password – same password must be provided to both primary and secondary Firewall.
8. Heartbeat Interface—Add Port 3/HA1 and Port 4/ HA2 port in heartbeat interfaces through which both primary and secondary devices can interchange hello messages to check liveliness of the peer device.
9. Select OK
Repeat Step 1 to Step 9 in Secondary Firewall.
—————————————————————————————————————————————–
Check HA status in Secondary devices. Refresh the entries and check sync status in Primary and Secondary HA monitoring Dashboard.
Dashboard widget shows below status if HA status is in sync.
Use Config Global Mode
get system ha status –>        shows HA and Cluster failover Information
FortiGate (global) # get sys ha status
HA Health Status: OK
Model: FortiGate-VM64-KVM
Mode: HA Active Passive
Group: HA-Group
Debug: 0
Cluster Uptime: 211 days 5:9:44
Cluster state change time: 2022-04-16 14:21:15
Master selected using: <2022/04/13 14:21:15> FGVMXXXXXXXXXX14 is selected as the master because it has the largest value of uptime.
<2022/04/13 14:15:46> FGVMXXXXXXXXXX16 is selected as the master because it has the largest value of uptime.
<2022/04/12 11:17:04> FGVMXXXXXXXXXX44 is selected as the master because it has the largest value of override priority.
ses_pickup: enable, ses_pickup_delay=disable
override: disable
Configuration Status: FGVMXXXXXXXXXX14(updated 2 seconds ago): in-sync FGVMXXXXXXXXXX16(updated 3 seconds ago): in-sync
System Usage stats: FGVMXXXXXXXXXX14(updated 2 seconds ago): FGVMXXXXXXXXXX16(updated 3 seconds ago):
sessions=12, average-cpu-user/nice/system/idle=0%/0%/0%/100%, memory=44%
sessions=2, average-cpu-user/nice/system/idle=0%/0%/0%/100%, memory=14%
HBDEV stats: FGVMXXXXXXXXXX14(updated 2 seconds ago): FGVMXXXXXXXXXX16(updated 3 seconds ago):
port3: physical/10000full, up, rx-bytes/packets/dropped/errors=2232258636/6463321/0/0, tx=3266257061/8035173/0/0
port3: physical/10000full, up, rx-bytes/packets/dropped/errors=3366612632/70886621/0/0, tx=1232321221/4564123/0/0
MONDEV stats: FGVMXXXXXXXXXX14(updated 1 seconds ago): FGVMXXXXXXXXXX16(updated 3 seconds ago): Master: Active-FW         , FGVMXXXXXXXXXX14, cluster index = 1
port4: physical/10000full, up, rx-bytes/packets/dropped/errors=5543991879/3242247/0/0, tx=554325343/4321945/0/0
port1: physical/10000full, up, rx-bytes/packets/dropped/errors=22183223/2218321/0/0, tx=216832/1211/0/0
Slave : Secondary-Fw         , FGVMXXXXXXXXXX16, cluster index = 0
number of vcluster: 1
vcluster 1: work 169.254.0.2
Master: FGVMXXXXXXXXXX14, operating cluster index = 0
Slave : FGVMXXXXXXXXXX16, operating cluster index = 1
Check the checksum mismatch and compare for the cluster checksum. Run command to go in rough for discrepancy VDOM’s by using command: diag sys ha checksum show <vdom>
diag sys ha checksum show <global>
Use grep to filter the configuration
diagnose sys ha checksum show root | grep system
diagnose sys ha checksum show global | grep log
Repeat above commands on secondary device to compare the mismatch output
Initiate and re-calculate checksum if no mismatch found.
Command to re-calculate the checksum diagnose sys ha checksum recalculate [<vdom-name> | global]
Above command re-calculates the checksum for all the devices.
Debug HA logs
diag debug app hasync 255 diagnose debug application hatalk -1
diag debug enable
execute ha synchronize start
communication between HA devices
Mismatch in HA can be calculated by using below command 1.diag debug config-error-log read
2. diag hardware device disk
3. show sys storage
4. show wanopt storage
IPSec VPN Configuration: Fortigate Firewall

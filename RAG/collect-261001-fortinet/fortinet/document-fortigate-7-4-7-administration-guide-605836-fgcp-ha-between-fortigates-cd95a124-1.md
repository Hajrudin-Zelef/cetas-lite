---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates-cd95a124-1
title: "document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates--cd95a124"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2023-05-29"]
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates--cd95a124.md
source_anchor: ""
source_lines: [1, 139]
sha256: ea98dbf4efece96abc3a8678a2c8063d77b7c6744745c107e2900e3e61862b57
---

# document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates--cd95a124

FGCP HA between FortiGates of the same model with different AC and DC PSUs
To improve power redundancy, FGCP HA clusters can support forming HA between units of the same model but with different AC PSU and DC PSU power supplies. This enables redundancy in a situation where power is completely lost on the AC grid, but traffic can fail over to a cluster member running on an independent DC grid.
The cluster members must be the same model with the same firmware installed, and must have the same hardware configuration other than the PSU.
In the following examples, there is an FGCP cluster with AC and DC PSU members: a FortiGate 1800F-DC (primary) and FortiGate 1800F (secondary).
Basic configuration
To configure the FGCP cluster in the GUI:
- 
                                                    On the primary FortiGate (FG-1800F-DC), go to System > HA.
- 
                                                    Configure the following settings: Mode Active-Passive Device priority 128 Group ID 0 Group name Example_cluster Password Enter a password. Session pickup Enable this setting. Monitor interfaces Click the + to add port5 and port6. Heartbeat interfaces Click the + to add ha1 and ha2.
- 
                                                    Click OK.
- 
                                                    On the secondary FortiGate (FG-1800F), go to System > HA.
- 
                                                    Configure the following settings: Mode Active-Passive Device priority 127 Group ID 0 Group name Example_cluster Password Enter a password. Session pickup Enable this setting. Monitor interfaces Click the + to add port5 and port6. Heartbeat interfaces Click the + to add ha1 and ha2.
- 
                                                    Click OK.
- 
                                                    Verify that the cluster status is Synchronized.
To configure the FGCP cluster in the CLI:
- 
                                                    Configure the primary FortiGate (FG-1800F-DC): config system ha
    set group-name "Example_cluster"
    set mode a-p
    set password **********
    set hbdev "ha2" 0 "ha1" 0 
    set session-pickup enable
    set override disable
    set monitor "port5" "port6" 
end
- 
                                                    Configure the secondary FortiGate (FG-1800F): config system ha
    set group-name "Example_cluster"
    set mode a-p
    set password **********
    set hbdev "ha2" 0 "ha1" 0 
    set session-pickup enable
    set override disable
    set priority 127
    set monitor "port5" "port6" 
end
- 
                                                    Verify the cluster status on the primary FortiGate: # get system ha status  
HA Health Status: OK
Model: FortiGate-1800F
Mode: HA A-P
Group Name: Example_cluster
Group ID: 0
Debug: 0
Cluster Uptime: 0 days 0:56:11
Cluster state change time: 2023-05-29 19:11:14
Primary selected using:
    <2023/05/29 19:11:14> vcluster-1: FG180FTK*******1 is selected as the primary because its uptime is larger than peer member FG180FTK*******2.
    <2023/05/29 18:59:45> vcluster-1: FG180FTK*******2 is selected as the primary because its uptime is larger than peer member FG180FTK*******1.
    <2023/05/29 18:59:45> vcluster-1: FG180FTK*******1 is selected as the primary because its override priority is larger than peer member FG180FTK*******2.
ses_pickup: enable, ses_pickup_delay=disable
override: disable
Configuration Status:
    FG180FTK*******1(updated 4 seconds ago): in-sync
    FG180FTK*******1 chksum dump: 95 4e 92 c3 39 75 8e 0e db 83 8d b7 b2 b1 9f 04 
    FG180FTK*******2(updated 5 seconds ago): in-sync
    FG180FTK*******2 chksum dump: 95 4e 92 c3 39 75 8e 0e db 83 8d b7 b2 b1 9f 04 
System Usage stats:
    FG180FTK*******1(updated 4 seconds ago):
        sessions=4, npu-sessions=0, average-cpu-user/nice/system/idle=0%/0%/0%/99%, memory=22%
    FG180FTK*******2(updated 5 seconds ago):
        sessions=0, npu-sessions=0, average-cpu-user/nice/system/idle=0%/0%/0%/99%, memory=22%
HBDEV stats:
    FG180FTK*******1(updated 4 seconds ago):
        ha2: physical/10000full, up, rx-bytes/packets/dropped/errors=18367581/33512/0/0, tx=9563450/16609/0/0
        ha1: physical/10000full, up, rx-bytes/packets/dropped/errors=11543018/22166/0/0, tx=12359673/22151/0/0
    FG180FTK*******2(updated 5 seconds ago):
        ha2: physical/10000full, up, rx-bytes/packets/dropped/errors=19133123/35087/0/0, tx=10685583/18475/0/0
        ha1: physical/10000full, up, rx-bytes/packets/dropped/errors=17011332/25876/0/0, tx=11919050/24991/0/0
MONDEV stats:
    FG180FTK*******1(updated 4 seconds ago):
        port5: physical/1000full, up, rx-bytes/packets/dropped/errors=988220/13742/0/0, tx=106998000/73260/0/0
        port6: physical/1000full, up, rx-bytes/packets/dropped/errors=107084264/73624/0/0, tx=953158/13611/0/0
    FG180FTK*******2(updated 5 seconds ago):
        port5: physical/1000full, up, rx-bytes/packets/dropped/errors=38194/128/0/0, tx=0/0/0/0
        port6: physical/1000full, up, rx-bytes/packets/dropped/errors=99019/448/0/0, tx=0/0/0/0
Primary     : FortiGate-1800F , FG180FTK*******1, HA cluster index = 1
Secondary   : FortiGate-1800F , FG180FTK*******2, HA cluster index = 0
number of vcluster: 1
vcluster 1: work 169.254.0.2
Primary: FG180FTK*******1, HA operating index = 0
Secondary: FG180FTK*******2, HA operating index = 1
- 
                                                    Verify the cluster status on the secondary FortiGate: # get system ha status  
HA Health Status: OK
Model: FortiGate-1800F
Mode: HA A-P
Group Name: Example_cluster
Group ID: 0
Debug: 0
Cluster Uptime: 0 days 0:56:53
Cluster state change time: 2023-05-29 19:11:14
Primary selected using:
    <2023/05/29 19:11:14> vcluster-1: FG180FTK*******1 is selected as the primary because its uptime is larger than peer member FG180FTK*******2.
    <2023/05/29 18:59:45> vcluster-1: FG180FTK*******2 is selected as the primary because its uptime is larger than peer member FG180FTK*******1.
    <2023/05/29 18:55:03> vcluster-1: FG180FTK*******2 is selected as the primary because it's the only member in the cluster.
    <2023/05/29 18:54:57> vcluster-1: FG180FTK*******2 is selected as the primary because SET_AS_SECONDARY flag is set on peer member FG180FTK*******1.
ses_pickup: enable, ses_pickup_delay=disable
override: disable
...
Secondary   : FortiGate-1800F , FG180FTK*******2, HA cluster index = 0
Primary     : FortiGate-1800F , FG180FTK*******1, HA cluster index = 1
number of vcluster: 1
vcluster 1: standby 169.254.0.2
Secondary: FG180FTK*******2, HA operating index = 1
Primary: FG180FTK*******1, HA operating index = 0
Testing synchronization in the cluster
Based on the preceding example, the interface and firewall policy configurations are changed on the primary FortiGate. These configuration changes and sessions are synchronized to the secondary FortiGate. If the switch interface connected to the primary's port5 is down (port2), this triggers the monitor interface to be down, and the PC1 traffic will fail over to the secondary FortiGate.
To test configuration synchronization in the FGCP cluster:
- 
                                                    Modify configurations on the primary FortiGate (FG-1800F-DC). 
  - 
                                                            Edit the interface settings: config system interface
    edit "port5"
        set ip 10.1.100.1 255.255.255.0
        set allowaccess ping https ssh http telnet
        set alias "To_Client_PC"
        config ipv6
            set ip6-address 2000:10:1:100::1/64
            set ip6-allowaccess ping https ssh http
        end
    next
    edit "port6"
        set ip 172.16.200.1 255.255.255.0
        set allowaccess ping https ssh http fgfm
        set alias "To_Server"
        config ipv6
            set ip6-address 2000:172:16:200::1/64
            set ip6-allowaccess ping https ssh http
        end
    next
end
  - 

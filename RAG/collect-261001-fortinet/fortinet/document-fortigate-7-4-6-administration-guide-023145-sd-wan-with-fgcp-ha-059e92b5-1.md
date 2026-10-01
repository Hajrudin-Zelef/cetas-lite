---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-administration-guide-023145-sd-wan-with-fgcp-ha-059e92b5-1
title: "document-fortigate-7-4-6-administration-guide-023145-sd-wan-with-fgcp-ha-059e92b5"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["licenses"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-administration-guide-023145-sd-wan-with-fgcp-ha-059e92b5.md
source_anchor: ""
source_lines: [1, 77]
sha256: f098148e550d74c5673151f637e36205a302b5ef791e86f3d295e6ab66877c4f
---

# document-fortigate-7-4-6-administration-guide-023145-sd-wan-with-fgcp-ha-059e92b5

SD-WAN with FGCP HA
SD-WAN with FGCP HA
This example shows how to convert a standalone FortiGate SD-WAN solution to a FGCP HA cluster with full-mesh WAN set up. This configuration allows you to load balance your internet traffic between multiple ISP links. It also provides redundancy for your internet connection if your primary ISP in unavailable, or if one of the FortiGates in the HA cluster fails.
This example assumes that a standalone FortiGate has already been configured for SD-WAN by following the SD-WAN quick start.
Standalone FortiGate:
FGCP HA cluster:
The following devices are required to convert the topology to HA:
- A second FortiGate that is the same model running the same firmware version.
- Two switches for connecting each FortiGate's WAN interface to the corresponding ISP modem.
Before you begin:
- Ensure that the licenses and subscriptions on both HA members match.
- Ensure that there are one or more ports reserved for HA heartbeat.
- Ensure you have physical access to both HA members.
|  | Enabling HA and re-cabling the WAN interfaces will cause network interruptions. This procedure should be performed during a maintenance window. | 
Configuring the standalone FortiGate for HA
After running the following commands, the FortiGate negotiates to establish an HA cluster. You might temporarily lose connectivity with the FortiGate as FGCP negotiations take place and the MAC addresses of the FortiGate interfaces are changed to HA virtual MAC addresses.
This configurations sets the HA mode to active-passive.
The ha1 and ha2 interfaces are configured as the heartbeat interfaces, with priorities set to 200 and 100 respectively. Setting different priorities for the heartbeat interfaces is a best practice, but is not required.
If you have more than one cluster on the same network, each cluster should have a different group ID. Changing the group ID changes the cluster interface's virtual MAC addresses. If the group IP causes a MAC address conflict on your network, select a different group ID.
Enabling override and increasing the device priority means that this FortiGate always becomes the primary unit.
To configure the standalone FortiGate for HA in the GUI:
- Go to System > Settings and change the Host name so that the FortiGate can be easily identified as the primary unit.
- Go to System > HA and configure the following options:Mode Active-Passive Device priority 250 Group name My-cluster Password <password> Heartbeat interfaces ha1 and ha2 Heartbeat Interface Priority port2 (ha1): 200 port3 (ha2): 100 Override and the group ID can only be configured from the CLI.
- Click OK.Connectivity with the FortiGate will temporarily be lost.
To configure the standalone FortiGate for HA in the CLI:
- Change the host name so that the FortiGate can be easily identified:config system global
    set hostname primary_FG
end
- Configure HA:config system ha
    set mode a-p
    set group-id 100
    set group-name My-cluster
    set password <password>
    set priority 250
    set override enable
    set hbdev ha1 200 ha2 100
end
|  | If HA mode does not start after running the above steps, ensure that none of the FortiGate's interfaces use DHCP or PPPoE addressing. | 
Configuring the secondary FortiGate for HA
The secondary FortiGate must be the same model and running the same firmware version as the primary FortiGate. The HA settings are the same as the for the primary unit, except the secondary device has a lower priority and override is not enabled.
|  | It is best practice to reset the FortiGate to factory default settings prior to configuring HA. This reduces the chance of synchronization problems. # execute factoryreset  This operation will reset the system to factory default! Do you want to continue? (y/n) y This is unnecessary if the device is new from the factory. | 
To configure the secondary FortiGate for HA in the GUI:
- Go to System > Settings and change the Host name so that the FortiGate can be easily identified as the backup unit.
- Go to System > HA and configure the options the same as for the primary FortiGate, except with a lower priority:Mode Active-Passive Device priority 128 Group name My-cluster Password <password> Heartbeat interfaces ha1 and ha2 Heartbeat Interface Priority port2 (ha1): 200 port3 (ha2): 100
- Click OK.
To configure the secondary FortiGate for HA in the CLI:
- Change the host name so that the secondary FortiGate can be easily identified:config system global
    set hostname secondary_FG
end
- Configure HA:config system ha
    set mode a-p
    set group-id 100
    set group-name My-cluster
    set password <password>
    set priority 128
    set hbdev ha1 200 ha2 100
end
Connecting the heartbeat interfaces between the FortiGates
To connect and check the heartbeat interfaces:
- Connect the heartbeat interfaces ha1 and ha2 between the primary and secondary FortiGate.
  - An HA primary device is selected. Because the primary FortiGate has a higher priority and override enabled, it assumes the role of HA primary.
  - The secondary FortiGate synchronizes its configuration from the primary device.
- Verify that the checksums match between the primary and secondary FortiGates:# diagnose sys ha checksum cluster ================== FG5H0XXXXXXXXXX0 ================== is_manage_primary()=1, is_root_primary()=1 debugzone global: 2b e9 81 38 c2 9d 4f db b7 0e 1f 49 42 c6 1e fb root: af a6 48 c5 c2 9a 8b 81 a5 53 fb 27 e9 ae 01 6a all: 89 1f 63 77 48 8a 30 ee 57 06 ca eb 71 e6 8e ad checksum global: 2b e9 81 38 c2 9d 4f db b7 0e 1f 49 42 c6 1e fb root: af a6 48 c5 c2 9a 8b 81 a5 53 fb 27 e9 ae 01 6a all: 89 1f 63 77 48 8a 30 ee 57 06 ca eb 71 e6 8e ad ================== FG5H0XXXXXXXXXX1 ================== is_manage_primary()=0, is_root_primary()=0 debugzone global: 2b e9 81 38 c2 9d 4f db b7 0e 1f 49 42 c6 1e fb root: af a6 48 c5 c2 9a 8b 81 a5 53 fb 27 e9 ae 01 6a all: 89 1f 63 77 48 8a 30 ee 57 06 ca eb 71 e6 8e ad checksum global: 2b e9 81 38 c2 9d 4f db b7 0e 1f 49 42 c6 1e fb root: af a6 48 c5 c2 9a 8b 81 a5 53 fb 27 e9 ae 01 6a all: 89 1f 63 77 48 8a 30 ee 57 06 ca eb 71 e6 8e ad If all of the cluster members have identical checksums, then their configurations are synchronized. If the checksums are not the same, wait for a few minutes, then repeat the command. Some parts of the configuration might take a significant amount of time to synchronize (tens of minutes).
Connecting other traffic interfaces
After the device configurations are synchronized, you can connect the rest of the traffic interfaces. Making these connections will disrupt traffic as cables are disconnected and reconnected.
Switches must be used between the cluster and the ISPs, and between the cluster and the internal network, as shown in the topology diagram.
Checking cluster operations
The HA Status dashboard widget shows the synchronization status. Hover over the host names of each FortiGate in the widget to verify that they are synchronized and have the same checksum.
To view more information about the cluster status, including the number of sessions passing through the cluster members, go to System > HA.
See Check HA synchronization status for more information.
Results
- Browse the internet on a computer in the internal network.
- Go to Network > SD-WAN and select the SD-WAN Zones tab to see the bandwidth, volume, and sessions for traffic on the SD-WAN interfaces. See Results for details.
- Go to Dashboard > Network, and expand the SD-WAN widget to see information about each interface, such as the number of sessions and the bit rate.
Testing HA failover
All traffic should currently be flowing through the primary FortiGate. If it becomes unavailable, traffic fails over to the secondary FortiGate. When the primary FortiGate rejoins the cluster, the secondary FortiGate continues to operate as the primary FortiGate.
To test this, ping a reliable IP address from a computer in the internal network, and then power off the primary FortiGate.

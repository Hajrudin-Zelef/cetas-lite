---
id: collect-261001-fortinet/fortinet/fortinet-fortigate-ha-high-availability-detailed-guide-network-interview
title: "What is High Availability?"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-fortinet/fortinet-fortigate-ha-high-availability-detailed-guide-network-interview.md
source_anchor: ""
source_lines: [1, 132]
sha256: 9310161a68f8ec810b286e27115982093ab14548cb7a62d476d9b3ac6a9c35fb
---

# What is High Availability?

### **Objectives** 

- High Availability
- HA Modes
- FGCP (FortiGate Clustering Protocol)
- Heartbeat Interfaces and Virtual IP Interfaces
- HA Requirement
- Configure Primary FortiGate Firewall
- Configure Secondary FortiGate Firewall
- HA-Troubleshooting

# What is High Availability?

High Availability (HA) is a feature of Firewalls in which two or more devices are grouped together to provide redundancy in the network. HA links and synchronises two or more devices. In FortiGate HA one device will act as a ***primary device*** (also called ***Active FortiGate***). Active device synchronises its configuration with another device in the group. Other FortiGate devices are called **Secondary or Standby devices.**

## **Fortigate HA Modes**

There are two Fortigate HA modes available:

- **Active / Passive-** Configuration of primary and secondary devices are in synchronisation. In Active/Passive mode the primary device is the only equipment which can actively process the traffic. Secondary FortiGate device remains in Passive mode and monitors the status of the primary device. If the problem is detected in the Primary FortiGate, the secondary device takes over the primary role.**This event is called HA failover.**


- **Active / Active** -All HA configuration must be in-synchronisation. Only difference in Active / Active mode is that in A/A mode all the FortiGate devices are processing the traffic.

## **FGCP (FortiGate Clustering Protocol)** 

HA Protocol used by FortiGate Cluster to communicate. FGCP travels between FortiGate cluster devices over the heartbeat links and uses TCP port 703 with Ethernet type values:

**TCP port 23** is used by FGCP for configuration synchronisation.  Firewall cluster uses FGCP to elect the primary, synchronize configuration, discover another firewall that belongs to the same HA and detect failover when any of the HA device fails.

**In Active/Passive, Primary Firewall performs below tasks:**

- Exchange heartbeat *Hello messages* with secondary device over control link
- Synchronizes routing table, DHCP information, running configuration
- Traffic sessions


**Secondary Firewall performs below tasks:**

- Monitor Primary device as to check if reachability is working in-between cluster or not
- If problem encountered with the Primary Firewall, secondary device take-over the traffic sessions
- Maintain Data Plane Processes like Forwarding Table, NAT Table, Authentication record


## **Heartbeat Interfaces and Corresponding IP addresses**

Virtual IP addresses are assigned to heartbeat Interfaces based on the serial number of FortiGate Firewall


Cluster uses these virtual IP addresses to differentiate cluster members and update configuration changes in clustered devices.

## **Fortigate HA Requirements**

**Fortigate HA Configuration****Configuring Primary FortiGate for HA**

**1.** Go to System **->**Select HA

**2.** Select mode Active-Passive Mode

**3.** Once ***Active-Passive mode*** selected multiple parameters are required

**4.** **Mode-** Active/ Passive

**5.** Set Device Priority -200. More numerical value higher the priority. Here Priority is set 200, secondary devices must have lower numerical value than Primary Firewall.

**6. Device Group**–  Group name must be the same for both primary and secondary devices. Here we have given the name HA-GROUP. Device Group is used in HA to assign two or more devices to be part of the same HA Group.

**7. Password** – same password must be provided to both primary and secondary Firewall.

**8. Heartbeat Interface**—Add **Port 3/HA1** and **Port 4/ HA2 port** in heartbeat interfaces through which both primary and secondary devices can interchange hello messages to check liveliness of the peer device.

**9.** Select **OK**

- The FortiGate exchanges messages to peer devices to establish an HA cluster. When Admin **select OK** connectivity can be lost with the FortiGate as the HA cluster negotiates and the FGCP initiate new MAC address of the FortiGate interfaces.
- Power off the FortiGate.
- Repeat the steps in Secondary devices and connect Port 3 and Port 4 with Secondary FortiGate Firewall.

### ***Configuring Secondary FortiGate for HA***

**Configuring Secondary FortiGate for HA**

Repeat Step 1 to Step 9 in Secondary Firewall.

—————————————————————————————————————————————–

**Check HA status in Secondary devices. Refresh the entries and check sync status in Primary and Secondary HA monitoring Dashboard.**

—————————————————————————————————————————————–

**Dashboard widget shows below status if HA status is in sync.**

## **Troubleshooting Commands: Fortigate HA**

**Use Config Global Mode**

**get system ha status –>        shows HA and Cluster failover Information**


**Master selected using:**


**Configuration Status:**


**System Usage stats:**


**HBDEV stats:**


**MONDEV stats:**


**Check the checksum mismatch and compare for the cluster checksum. Run command to go in rough for discrepancy VDOM’s by using command:**

**Use grep to filter the configuration**  

**Repeat above commands on secondary device to compare the mismatch output**

**Initiate and re-calculate checksum if no mismatch found.**

**Command to re-calculate the checksum**


*Above command re-calculates the checksum for all the devices.*

**Debug HA logs**

**communication between HA devices**


**Mismatch in HA can be calculated by using below command**

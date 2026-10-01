---
id: collect-261001-general-networking/general-networking/switching-ms-switches-operate-and-maintain-monitoring-and-reporting-ms-netflow-a-fbe0a20b
title: "switching-ms-switches-operate-and-maintain-monitoring-and-reporting-ms-netflow-a-fbe0a20b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["asic"]
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-operate-and-maintain-monitoring-and-reporting-ms-netflow-a-fbe0a20b.md
source_anchor: ""
source_lines: [1, 66]
sha256: fe9e5a973449d16111bc0f8e3a80e7e546d9360f87c2d2b52544f9fe966267e9
---

# switching-ms-switches-operate-and-maintain-monitoring-and-reporting-ms-netflow-a-fbe0a20b

How to Configure MS NetFlow and Encrypted Traffic Analytics
Overview
Starting with CS15+, all Catalyst-based (MS390 and 9300 series) switches support Network Based Application Recognition (NBAR) NetFlow v10 (IPFIX) for IPv4 and IPv6 traffic. These switches also support Encrypted Traffic Analytics (ETA) flow export for use with NetFlow analyzers like Cisco Secure Network Analytics (formerly Stealthwatch Enterprise and Cloud). NetFlow and ETA enable organizations to export incredibly granular flows of any traffic through a supported switch, whether flowing east, west, north, or southbound.
The addition of ETA allows an organization to correlate cryptographic standards to identify security issues in protocols used locally. ETA also validates security audits for compliance requirements and fingerprint traffic without performing man-in-the-middle (MiTM) decryption. ETA is incredibly useful for products like Cisco Secure Network Analytics and Cisco Extended Detection and Response (XDR) to perform fingerprinting of potential security events in the network, such as encrypted malware and botnet activity.
Starting with the Cloud Native IOS XE release (17.15+), all 9K-M support AVC NetFlow exporting. ETA is limited to the MS390 and 9300 platforms due to hardware requirements and software support.
Features
Flow recording configurations are implemented for both IPv4 and IPv6 by default. When the feature is enabled, every interface will collect flow records in both the input and output directions as per the AVC NetFlow specification in the UADP ASIC. This means that all traffic will be captured and exported to the specified NetFlow collector. There is support for up to 1x exporter configured per-network.
Encrypted Traffic Analytics (ETA) support when configured is also enabled on every interface on every switch in the network that supports the feature and is configured correctly. This configuration enables Cisco's ETA functionality. For further information, refer to the Cisco Encrypted Traffic Analytics
Prerequisites
Hardware:
- MS390 (all models for AVC + NetFlow + ETA)
- 9300/X/L-M (all models for AVC + NetFlow + ETA)
- 9350 ( all models for AVC + NetFlow )
- 9200/L -M (only supports AVC + NetFlow)
- (AVC NOT SUPPORTED) C9500/c9500x/C9550/C94xx/C96xx (Netflow yes)
Software:
- CS 15.X and later or Cloud Native IOS XE 17.15+ (earliest, minimum sw for platform appliess)
Licensing:
- Advanced Licensing
Configuration:
- L3 SVI configured on exporting switch/stack (this is default on cloud native IOS XE 17.15+)
- Collector that is reachable via L3 SVI
For devices running cloud native IOS XE 17.15+ , the network-wide > settings for traffic analytics requires being set to "detailed" for the flow telemetry to be exported.
Step-by-step instructions
NetFlow and ETA configuration are broken down into a few steps:
- Deploy NetFlow and ETA configurations under Network-wide > General where you will provide the NetFlow collector port and IP. This configuration is shared with any MXs deployed in the network. Optionally, you can enable ETA and provide the UDP port for those ETA flow records to be shipped to. By default, these records are expected to be exported to the same flow collector. 
- For NetFlow to work on Layer-2 switches in the initial release, the switch is required to have an SVI created for routing of the NetFlow traffic. This is due to ASIC forwarding functions and thus must be a separate IP address from the management IP that is used to communicate with dashboard. This does allow for NetFlow traffic to take an alternative path from management but must not be confused with Alternate Management Interface which is a separate feature for RADIUS, SNMP, and Syslog traffic that is outlined here: Alternate Management Interface on MS Devices
- This can be performed by navigating to Switch > Routing & DHCP and creating an SVI for the switch or switch stack to leverage for NetFlow and optionally ETA forwarding as well. Once the SVI is created, the traffic will egress via the data-plane forwarding table.
SNMP polling is enabled by default in the latest Cisco's Secure Network Analytics (SNA) version. SNMP polling on SVIs is not supported on Catalyst based (MS390/9300) switches running CS-based firmware.
Verification
Technical breakdown
NetFlow and ETA are enabled on every port on the switch or switch stack as long as the switches are supported and properly configured in the network as outlined above. This means that all traffic that flows through the switch will be recorded and exported as a flow record for further inspection by a central analysis tool. For customers that require NetFlow or ETA to be exported to more than one destination, a product like Cisco's Telemetry Broker or Secure Network Analytics UDP Director (Stealthwatch UDP Director) can be utilized for flow duplication and re-routing.
The NetFlow recorder configurations are very granular and cover the following fields for both IPv4 and IPv6:
Match is a required field for the flow record to be created, and collect is a best effort field that is recorded when available.
It is not possible at this time to change these flow recording configurations. 
| MATCH | COLLECT | 
| NBAR application | connection client counter bytes network long | 
| interface input | connection client counter packets long | 
| source address | connection client ip address | 
| destination address | connection client transport port | 
| protocol | connection initiator | 
| version | connection new-connections | 
| transport source port | connection server counter bytes network long | 
| transport destination port | connection server counter packets long | 
|  | connection server ipv4 address | 
|  | connection server transport port | 
|  | counter bytes long | 
|  | counter packets long | 
|  | adaptive policy source group-tag | 
|  | dot1q vlan input | 
|  | destination mac address | 
|  | source mac address | 
|  | flow direction | 
|  | flow observation point | 
|  | interface output | 
|  | timestamp absolute first | 
|  | timestamp absolute last | 
|  | transport tcp flags | 
Troubleshooting
You can troubleshoot flow export by grabbing a packet capture on the uplink port and filtering in Wireshark for the UDP ports for ETA and NetFlow. You can capture packets by navigating to Network-wide > Packet capture (intelligent capture), select the switch in question, populate the uplink port, and select download for Wireshark. A 60 second capture should be sufficient in capturing some flow traffic if there are more than a few active clients behind the switch.
Example:
The following filter/s are handy:
cflow || udp.port == {netflow port} || udp.port == {eta port}
NOTE: Wireshark already matches a number of ports with the CFLOW protocol type which can be modified to include any other ports you may be using for flow exporting via Wireshark > Preferences > Protocols > CFLOW like below: 
If the configuration is adjusted, the only required filter would be "cflow" to properly filter the capture.
Example packet capture:

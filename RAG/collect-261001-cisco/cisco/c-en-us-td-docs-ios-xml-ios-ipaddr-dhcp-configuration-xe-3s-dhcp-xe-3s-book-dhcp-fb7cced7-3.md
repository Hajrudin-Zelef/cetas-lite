---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7-3
title: "c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7.md
source_anchor: ""
source_lines: [158, 272]
sha256: 06b81e2e0b9ed84c54fc2f0e6152fb2c7787aa8e5b2cd91139023857b372d0af
---

# c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7

DETAILED STEPS
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Router> enable | Enables privileged EXEC mode.  | 
| Step 2 | configure                                   terminal Example: Router# configure terminal | Enters global configuration mode. | 
| Step 3 | snmp-server             enable             traps             dhcp                                   duplicate           ]   [interface]   [pool]   [subnet]   [time Example: Router(config)# snmp-server enable traps dhcp | Enables the sending of DHCP SNMP trap notifications.  | 
| Step 4 | end Example: Router(config)# end | Returns the router to privileged EXEC mode. | 
Troubleshooting Tips
If you are using secondary IP addresses under a single loopback interface and using secondary subnets under a DHCP pool, use one DHCP pool to configure networks for all the secondary subnets instead of using one pool per secondary subnet. The network network-number [mask | /prefix-length] [secondary] command must be configured under a single DHCP address pool rather than multiple DHCP address pools.
The following is the correct configuration:
!
ip dhcp pool dhcp_1
 network 172.16.1.0 255.255.255.0
 network 172.16.2.0 255.255.255.0 secondary
 network 172.16.3.0 255.255.255.0 secondary
 network 172.16.4.0 255.255.255.0 secondary
!
interface Loopback111
 ip address 172.16.1.1 255.255.255.255 secondary
 ip address 172.16.2.1 255.255.255.255 secondary
 ip address 172.16.3.1 255.255.255.255 secondary
 ip address 172.16.4.1 255.255.255.255 secondary
The following is the incorrect configuration:
! 
ip dhcp pool dhcp_1
 network 172.16.1.0 255.255.255.0
 lease 1 20 30
 accounting default
!
ip dhcp pool dhcp_2
 network 172.16.2.0 255.255.255.0
 lease 1 20 30
 accounting default
!
ip dhcp pool dhcp_3
 network 172.16.3.0 255.255.255.0
 lease 1 20 30
 accounting default
!
ip dhcp pool dhcp_4
 network 172.16.4.0 255.255.255.0
 lease 1 20 30
 accounting default
!
interface Loopback111
 ip address 172.16.1.1 255.255.255.255 secondary
 ip address 172.16.2.1 255.255.255.255 secondary
 ip address 172.16.3.1 255.255.255.255 secondary
 ip address 172.16.4.1 255.255.255.255 secondary
Configuration Examples for the DHCP Server MIB
- DHCP Server MIB--Secondary Subnet Trap Example
- DHCP Server MIB--Address Pool Trap Example
- DHCP Server MIB--Lease Limit Violation Trap Example
DHCP Server MIB--Secondary Subnet Trap Example
The following example configures 192.0.2.0/24 as the subnetwork number and mask of the DHCP pool named pool2 and then adds the DHCP pool secondary subnet specified by the subnet number and mask 192.0.4.0/30. The IP addresses in pool2 consist of two disjoint subnets: the addresses from 192.0.2.1 to 192.0.2.254 and the addresses from 192.0.4.1 to 192.0.4.2.
The address pool utilization mark, configured at the global level, will be overridden at the secondary subnet level. A trap is sent to the SNMP manager if the subnet size of the secondary subnet exceeds or goes below the level specified by the override utilization commands.
The utilization mark{high| low}log command enables a system message to be generated for a DHCP address pool or secondary subnet when the utilization exceeds the configured high utilization threshold or falls below the configured low utilization threshold.
!
ip dhcp pool pool2 
 utilization mark high 80 log 
 utilization mark low 70 log 
 network 192.0.2.0 255.255.255.0 
 network 192.0.4.0 255.255.255.252 secondary 
 override utilization high 40 
 override utilization low 30
!
snmp-server enable traps dhcp subnet
DHCP Server MIB--Address Pool Trap Example
In the following example, if the address utilization exceeds the high threshold or drops below the low threshold, an SNMP trap will be sent to the SNMP manager and a system message will be generated.
ip dhcp pool pool3 
 utilization mark high 80 log 
 utilization mark low 70 log 
!
snmp-server enable traps dhcp pool
DHCP Server MIB--Lease Limit Violation Trap Example
In the following example, four DHCP clients are allowed to receive IP addresses. If a fifth client tries to obtain an IP address, the DHCPDISCOVER messages will not be forwarded to the DHCP server and a trap will be sent to the SNMP manager.
ip dhcp limit lease log
interface Serial 0/0
 ip dhcp limit lease 4
 exit 
snmp-server enable traps dhcp interface
Additional References
The following sections provide references related to the DHCP Server MIB feature.
Related Documents
| Related Topic | Document Title | 
|---|---|
| SNMP configuration tasks | “Configuring SNMP Support” module | 
| DHCP commands: complete command syntax, command mode, command history, defaults, usage guidelines, and examples | Cisco IOS IP Addressing Services Command Reference | 
| DHCP server configuration tasks including subnet utilization tasks | “Configuring the Cisco IOS DHCP Server” module | 
| DHCP per interface lease limit functionality | “Configuring DHCP Services for Accounting and Security” module | 
| DHCP ODAP tasks including address pool utilization tasks | “Configuring the DHCP Server On-Demand Address Pool Manager” module | 
Standards
| Standard | Title | 
|---|---|
| No new or modified standards are supported by this feature. | -- | 
MIBs
| MIB | MIBs Link | 
|---|---|
|  | To locate and download MIBs for selected platforms, Cisco IOS releases, and feature sets, use Cisco MIB Locator found at the following URL: http://www.cisco.com/go/mibs | 
RFCs
| RFC | Title | 
|---|---|
| Draft RFC: draft-ietf-dhc-server-mib-10.txt | Dynamic Host Configuration Protocol for IPv4 (DHCPv4) Server MIB | 
Technical Assistance
| Description | Link | 
|---|---|
| The Cisco Support and Documentation website provides online resources to download documentation, software, and tools. Use these resources to install and configure the software and to troubleshoot and resolve technical issues with Cisco products and technologies. Access to most tools on the Cisco Support and Documentation website requires a Cisco.com user ID and password. | http://www.cisco.com/cisco/web/support/index.html | 
Feature Information for DHCP Server MIB
The following table provides release information about the feature or features described in this module. This table lists only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise, subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
| Table 10 Feature Information for 		  DHCP Server MIB |  |  | 
|---|---|---|
| Feature Name | Releases | Feature Information | 
|---|---|---|
| DHCP Server MIB | Cisco IOS XE Release 3.8S | The DHCP Server MIB feature provides SNMP access to and control of Cisco IOS DHCP server software on a Cisco router by an external network management device. The following commands were introduced by this feature: snmp-server enable traps dhcpand debug ip dhcp server snmp. |

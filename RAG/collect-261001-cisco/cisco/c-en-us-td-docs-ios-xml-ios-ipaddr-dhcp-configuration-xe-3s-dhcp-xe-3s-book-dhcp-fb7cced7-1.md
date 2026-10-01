---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7-1
title: "c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7.md
source_anchor: ""
source_lines: [1, 79]
sha256: f4d0ad631304e73ffde014fcf73a89f54b8e1804e7fd31755cbf961a8e91f76e
---

# c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7

DHCP Server MIB
The DHCP Server MIB feature provides Simple Network Management Protocol (SNMP) access to and control of Cisco IOS Dynamic Host Configuration Protocol (DHCP) server software on a Cisco router by an external network management device.
Finding Feature Information
Your software release may not support all the features documented in this module. For the latest caveats and feature information, see Bug Search Tool and the release notes for your platform and software release. To find information about the features documented in this module, and to see a list of the releases in which each feature is supported, see the feature information table.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
Prerequisites for the DHCP Server MIB
SNMP must be enabled on the router before DHCP server trap notifications can be configured.
Information About the DHCP Server MIB
SNMP Overview
SNMP is an application-layer protocol that provides a message format for communication between SNMP managers and agents. SNMP provides a standardized framework and a common language that is used for monitoring and managing devices in a network.
SNMP defines two main types of entities: managers and agents. The SNMP manager is a system that controls and monitors the activities of network hosts using SNMP. The agent is the software component within a remote networking device that maintains the data and reports this data, as needed, to the manager. The manager and agent share a Management Information Base (MIB) that defines the information that the agent can make available to the manager.
An important feature of SNMP is the capability to generate unsolicited notifications from an SNMP agent. These trap notifications are messages alerting the SNMP manager to conditions on the network. Traps are considered an agent-to-manager function and a request for confirmation of receipt from the SNMP manager is not required.
DHCP Server Trap Notifications
DHCP server trap notifications are sent to the SNMP manager for the following events:
- Address utilization for a subnet has risen above or fallen below a configurable threshold.
- Address utilization for an address pool has risen above or fallen below a configurable threshold.
- A lease limit violation is detected. The lease limit configuration allows you to control the number of subscribers per interface.
- The DHCP server has started or stopped.
- A duplicate IP address is detected.
The DHCP Server MIB feature does not send the same type of trap notification back-to-back for the same threshold event. For example, if the low threshold value for available free addresses becomes equal to or less than the configured value, a free address low event trap notification on the subnet or pool is generated. This same trap notification will not be resent until the value for the available free addresses has exceeded the value of the free high threshold and vise versa. This threshold control mechanism applies to all trap notifications concerning thresholds in addition to the trap notifications for the DHCP server start and stop time and the lease limit violation. The duplicate IP address trap notification is not subject to this threshold control mechanism.
Tables and Objects in the DHCP Server MIB
The DHCP Server MIB consists of the following tables and objects. The first character of a row in the table begins with “c” (Cisco) and is mapped to the object defined in the IETF draft RFC, Dynamic Host Configuration Protocol for IPv4 Server MIB. If the information is not currently available in Cisco IOS software, the value in the second column is displayed as 0 (zero).
- cDhcpv4SrvSystemsObjects (see Table 7)--System description and object IDs
- cBootpHCCounterObjects (see Table 8)--BOOTP counter information
- cDhcpv4HCCounterObjects (see Table 9)--DHCPv4 counter information
- cDhcpv4ServerSharedNetTable (see Table 10)--DHCP address pool information
- cDhcpv4ServerSubnetTable (see Table 11)--Additional DHCP address pool subnet information including secondary subnet information
- cDhcpv4SrvExtSubnetTable (see Table 12)--Additional DHCP address pool subnet information
- cDhcpv4ServerNotifyObjectsGroup (see Table 13)--This objects group is used by the cDhcpv4ServerNotificationsGroup notifications group.
- cDhcpv4ServerNotificationsGroup (see Table 14)--This notifications group consists of all traps defined in the Cisco IOS DHCP server.
- cDhcpv4SrvExtNotifyGroup (see Table 15)--This notifications group consists of all traps not defined in the draft DHCPv4 Server MIB RFC.
| Table 1 cDhcpv4SrvSystemsObjects and Descriptions |  | 
|---|---|
| Name | Description | 
|---|---|
| cDhcpv4SrvSystemDescr | Contains a textual description of the server (full name and version identification). | 
| cDhcpv4SrvSystemObjectID | Cisco experiment node for the DHCP Server MIB. For example, 1.3.6.1.4.1.9.10.102... | 
| Table 2 cBootpHCCounterObjects and Descriptions |  | 
|---|---|
| Name | Description | 
|---|---|
| cBootpHCCountRequests | The number of packets received that do contain a BOOTREQUEST message type in the first octet. | 
| cBootpHCCountInvalids | 0 | 
| cBootpHCCountReplies | The number of packets received that contain a BOOTREPLY message type in the first octet. | 
| cBootpHCCountDroppedUnknown Clients | 0 | 
| cBootpHCCountDroppedNotServingSubnet | 0 | 
| Table 3 cDhcpv4HCCounterObjects and Descriptions |  | 
|---|---|
| Name | Description | 
|---|---|
| cDhcpv4HCCountDiscovers | The number of DHCPDISCOVER packets received. | 
| cDhcpv4HCCountOffers | The number of DHCPOFFER packets sent. | 
| cDhcpv4HCCountRequests | The number of DHCPREQUEST packets sent. | 
| cDhcpv4HCCountDeclines | The number of DHCPDECLINE packets sent. | 
| cDhcpv4HCCountAcks | The number of DHCPACK packets sent. | 
| cDhcpv4HCCountNaks | The number of DHCPNACK packets sent. | 
| cDhcpv4HCCountReleases | The number of DHCPRELEASE packets sent. | 
| cDhcpv4HCCountInforms | The number of DHCPINFORM packets sent. | 
| cDhcpv4HCCountForcedRenews | 0 | 
| cDhcpv4HCCountInvalids | The number of DHCP packets received whose DHCP message type is not understood or handled by the DHCP server. | 
| cDhcpv4HCCountDropUnknownClient | 0 | 
| cDhcpv4HCCountDropNotServingSubnet | 0 | 
| Table 4 cDhcpv4ServerSharedNetTable and Descriptions |  | 
|---|---|
| Name | Description | 
|---|---|
| cDhcpv4ServerSharedNetName | The DHCP address pool name. | 
| cDhcpv4ServerSharedNetFreeAddr LowThreshold | This entry value corresponds to the utilization mark high command in DHCP pool configuration mode multiplied by the total pool addresses then divided by 100. | 
| cDhcpv4ServerSharedNetFreeAddrHighThreshold | This entry value corresponds to the utilization mark low command in DHCP pool configuration mode multiplied by the total subnet addresses then divided by 100. | 
| cDhcpv4ServerSharedNetFree Addresses | The number of IPv4 addresses that are available within this shared network. | 
| cDhcpv4ServerSharedNetReserved Addresses | The number of IP addresses that are reserved for the pool (not available for assignment). This entry corresponds to the ip dhcp excluded-address global configuration command. The value is zero if no excluded addresses are defined for the pool. | 
| cDhcpv4ServerSharedNetTotal Addresses | The number of IP addresses that are available within this shared network. | 
| Table 5 cDhcpv4ServerSubnetTable and Descriptions |  | 
|---|---|
| Name | Description | 
|---|---|
| cDhcpv4ServerSubnetAddress | The IP address of the subnet entry in the table. | 
| cDhcpv4ServerSubnetMask | The subnet mask of the subnet. | 
| cDhcpv4ServerSubnetSharedNetworkName | The DHCP address pool name to which the subnet belongs. | 

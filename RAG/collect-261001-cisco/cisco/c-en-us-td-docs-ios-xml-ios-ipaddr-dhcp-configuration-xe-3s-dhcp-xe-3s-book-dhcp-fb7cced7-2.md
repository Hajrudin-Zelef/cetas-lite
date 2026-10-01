---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7-2
title: "c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7.md
source_anchor: ""
source_lines: [80, 157]
sha256: 337c8a7c16ebb0841feb91945d1e081f5bb2e369b34172e1ba30fdbe303e167a
---

# c-en-us-td-docs-ios-xml-ios-ipaddr-dhcp-configuration-xe-3s-dhcp-xe-3s-book-dhcp-fb7cced7

| cDhcpv4ServerSubnetFreeAddrLowThreshold | This entry value corresponds to the override utilization high command in DHCP pool secondary subnet configuration mode multiplied by the total subnet addresses then divided by 100. | 
| cDhcpv4ServerSubnetFreeAddrHighThreshold | This entry value corresponds to the override utilization low command in DHCP pool secondary subnet configuration mode multiplied by the total subnet addresses then divided by 100. | 
| cDhcpv4ServerSubnetFree Addresses | The number of free IP addresses that are available in the subnet. | 
| Table 6 cDhcpv4SrvExtSubnetTable and Descriptions |  | 
|---|---|
| Name | Description | 
|---|---|
| cDhcpv4ServerDefaultRouterAddress | The entry corresponds to the override default-router command in DHCP pool secondary subnet configuration mode. | 
| cDhcpv4ServerSubnetStartAddress | The first subnet IP address. | 
| cDhcpv4ServerSubnetEndAddress | The last subnet IP address. | 
| Table 7 cDhcpv4ServerNotifyObjectsGroups and Descriptions |  | 
|---|---|
| Name | Description | 
|---|---|
| cDhcpv4ServerNotifyDuplicateIpAddr | The IP address is found to be a duplicate. Duplicates are detected by servers who send a PING before offering an IP address lease or by a client sending a gratuitous ARP message reported through a DHCPDECLINE message. | 
| cDhcpv4ServerNotifyDuplicateMac | The offending MAC address that caused a duplicate IPv4 address to be detected, if captured by the server, otherwise set to 00-00-00-00-00-00. | 
| cDhcpv4ServerNotifyClientOrServerDetected | This object is set by the server to client if the client used DHCPDECLINE to mark the offered address as in use, or to server if the server discovered that address was in use by a client before offering it. | 
| cDhcpv4ServerNotifyServerStart | The date and time when the server began operation, which is controlled by the service dhcp command. | 
| cDhcpv4ServerNotifyServerStop | The date and time when the server ceased operation, which is controlled by no service dhcp command. | 
| Table 8 cDhcpv4ServerNotificationsGroup and Descriptions |  | 
|---|---|
| Name | Description | 
|---|---|
| cDhcpv4ServerFreeAddressLow | This notification signifies that the number of available IP addresses for a DHCP address pool has fallen below the defined low threshold. This notification corresponds to the snmp-server enable traps dhcp global configuration command. | 
| cDhcpv4ServerFreeAddressHigh | This notification signifies that the number of available IP addresses for a DHCP address pool has risen above the defined high threshold. This notification corresponds to the snmp-server enable traps dhcp global configuration command. | 
| cDhcpv4ServerStartTime | This notification signifies that the server has started. This notification corresponds to the service dhcp and snmp-server enable traps dhcp timeglobal configuration commands. | 
| cDhcpv4ServerStopTime | This notification signifies that the server has stopped normally. This notification corresponds to the no service dhcp and snmp-server enable traps dhcp timeglobal configuration commands. | 
| cDhcpv4ServerDuplicateAddress | This notification signifies that a duplicate IP address has been detected. This notification corresponds to the snmp-server enable traps dhcp duplicateglobal configuration command. | 
| Table 9 cDhcpv4SrvNotifyGroup and Descriptions |  | 
|---|---|
| Name (not in the RFC draft) | Description | 
|---|---|
| cDhcpv4ServerIfLeaseLimitExceeded | This notification signifies that a per interface lease limit is exceeded. This notification corresponds to the snmp-server enable traps dhcp interfaceglobal configuration command. | 
| cDhcpv4ServerSubnetFreeAddressLow | This notification signifies that the number of available IP addresses for a subnet has fallen below the defined low threshold. This notification corresponds to the snmp-server enable traps dhcp subnetglobal configuration command. | 
| cDhcpv4ServerSubnetFreeAddressHigh | This notification signifies that the number of available IPv4 addresses for a subnet has risen above the defined high threshold. This notification corresponds to the snmp-server enable traps dhcp subnetglobal configuration command. | 
How to Enable DHCP Trap Notifications
Configuring the Router to Send SNMP Trap Notifications About DHCP
DHCP trap notifications are disabled by default. The trap notification is disabled if the corresponding trap configuration is not enabled.
1.   
      
          
            enable
          
        
2.   
      
          
            configure
          
          
            terminal
          
        
3.   
      
          
            snmp-server
            enable
            traps
            dhcp
          
          
            duplicate
          ]
 
[interface]
 
[pool]
 
[subnet]
 
[time
4.   
      
          
            end
          
        

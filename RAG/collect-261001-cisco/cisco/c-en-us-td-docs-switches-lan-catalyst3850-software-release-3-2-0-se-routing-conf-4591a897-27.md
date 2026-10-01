---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-27
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1438, 1496]
sha256: 00ba9e3f882f6ed0a6079a83523d5eefcac33d452dca08453d70605fac6805c9
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Step 6 | import map  				route-map Example:  Device(config-vrf)# import map importmap1  | (Optional) Associates a route map with the VRF. | 
| Step 7 | interface  				interface-id Example:  Device(config-vrf)# interface gigabitethernet 1/0/1  | Specifies the Layer 3 interface to be associated with the VRF, and enter interface configuration mode. The interface can be a routed port or SVI. | 
| Step 8 | ip vrf forwarding 						vrf-name Example:  Device(config-if)# ip vrf forwarding vpn1  | Associates the VRF with the Layer 3 interface. | 
| Step 9 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 10 | show ip vrf [brief \|  				detail \|  				interfaces] [vrf-name] Example:  Device# show ip vrf interfaces vpn1  | Verifies the configuration. Displays information about the configured VRFs. | 
| Step 11 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Note | When ip vrf forwarding is enabled in the Management Interface, the access point does not join. | 
Configuring VRF-Aware Services
Configuring VRF-Aware Services for ARP
Configuring VRF-Aware Services for Ping
For complete syntax and usage information for the commands, see the switch command reference for this release and the Cisco IOS Switching Services Command Reference, Release 12.4.
Configuring VRF-Aware Services for SNMP
For complete syntax and usage information for the commands, refer to the switch command reference for this release and the Cisco IOS Switching Services Command Reference, Release 12.4.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | snmp-server trap 				  authentication vrf Example:  Device(config)# snmp-server trap authentication vrf  | Enables SNMP traps for packets on a VRF. | 
| Step 3 | snmp-server 				  engineID remote  				host  				vrf  				vpn-instance engine-id 				  string Example:  Device(config)# snmp-server engineID remote 172.16.20.3 vrf vpn1 80000009030000B064EFE100  | Configures a name for the remote SNMP engine on a switch. | 
| Step 4 | snmp-server 				  host  				host  				vrf  				vpn-instance  				traps  				community Example:  Device(config)# snmp-server host 172.16.20.3 vrf vpn1 traps comaccess  | Specifies the recipient of an SNMP trap operation and specifies the VRF table to be used for sending SNMP traps. | 
| Step 5 | snmp-server 				  host  				host  				vrf  				vpn-instance  				informs  				community Example:  Device(config)# snmp-server host 172.16.20.3 vrf vpn1 informs comaccess  | Specifies the recipient of an SNMP inform operation and specifies the VRF table to be used for sending SNMP informs. | 
| Step 6 | snmp-server 				  user  				user group  				remote  				host  				vrf  				vpn-instance security 				  model Example:  Device(config)# snmp-server user abcd remote 172.16.20.3 vrf vpn1 priv v2c 3des secure3des  | Adds a user to an SNMP group for a remote host on a VRF for SNMP access. | 
| Step 7 | end Example:  Device(config-if)# end   | Returns to privileged EXEC mode. | 
Configuring VRF-Aware Servcies for uRPF
uRPF can be configured on an interface assigned to a VRF, and source lookup is done in the VRF table.
For complete syntax and usage information for the commands, refer to the switch command reference for this release and the Cisco IOS Switching Services Command Reference, Release 12.4.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | interface interface-id Example: Device(config)# 			 interface gigabitethernet 1/0/1 		   | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 3 | no 				  switchport Example:  Device(config-if)# no switchport  | Removes the interface from Layer 2 configuration mode if it is a physical interface. | 
| Step 4 | ip vrf 				  forwarding  				vrf-name Example:  Device(config-if)# ip vrf forwarding vpn2  | Configures VRF on the interface. | 
| Step 5 | ip address  				ip-address Example:  Device(config-if)# ip address 10.1.5.1  | Enters the IP address for the interface. | 
| Step 6 | ip verify unicast 				  reverse-path Example:  Device(config-if)# ip verify unicast reverse-path  | Enables uRPF on the interface. | 
| Step 7 | end Example:  Device(config-if)# end   | Returns to privileged EXEC mode. | 
Configuring VRF-Aware RADIUS
To configure VRF-Aware RADIUS, you must first enable AAA on a RADIUS server. The switch supports the ip vrf forwarding vrf-name server-group configuration and the ip radius source-interface global configuration commands, as described in the Per VRF AAA Feature Guide.
Configuring VRF-Aware Services for Syslog
For complete syntax and usage information for the commands, refer to the switch command reference for this release and the Cisco IOS Switching Services Command Reference, Release 12.4.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | logging 				  on Example:  Device(config)# logging on  | Enables or temporarily disables logging of storage router event message. | 
| Step 3 | logging 				  host  				ip-address  				vrf  				vrf-name Example:  Device(config)# logging host 10.10.1.0 vrf vpn1  | Specifies the host address of the syslog server where logging messages are to be sent. | 
| Step 4 | logging 				  buffered  				logging buffered size  				debugging Example:  Device(config)# logging buffered critical 6000 debugging  | Logs messages to an internal buffer. | 
| Step 5 | logging trap 				  debugging Example:  Device(config)# logging trap debugging  | Limits the logging messages sent to the syslog server. | 
| Step 6 | logging 				  facility  				facility Example:  Device(config)# logging facility user  | Sends system logging messages to a logging facility. | 
| Step 7 | end Example:  Device(config-if)# end   | Returns to privileged EXEC mode. | 
Configuring VRF-Aware Services for Traceroute
For complete syntax and usage information for the commands, refer to the switch command reference for this release and the Cisco IOS Switching Services Command Reference, Release 12.4.
Configuring VRF-Aware Services for FTP and TFTP
So that FTP and TFTP are VRF-aware, you must configure some FTP/TFTP CLIs. For example, if you want to use a VRF table that is attached to an interface, say E1/0, you need to configure the ip tftp source-interface E1/0 or the ip ftp source-interface E1/0 command to inform TFTP or FTP server to use a specific routing table. In this example, the VRF table is used to look up the destination IP address. These changes are backward-compatible and do not affect existing behavior. That is, you can use the source-interface CLI to send packets out a particular interface even if no VRF is configured on that interface.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | ip ftp 				  source-interface  				interface-type 				  interface-number Example:  Device(config)# ip ftp source-interface gigabitethernet 1/0/2  | Specifies the source IP address for FTP connections. | 
| Step 3 | end Example:  Device(config)#end  | Returns to privileged EXEC mode. | 
| Step 4 | configure 				  terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 5 | ip tftp 				  source-interface  				interface-type 				  interface-number Example:  Device(config)# ip tftp source-interface gigabitethernet 1/0/2  | Specifies the source IP address for TFTP connections. | 
| Step 6 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 

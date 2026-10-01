---
id: collect-261001-general-networking/general-networking/2016-12-dhcp-servers-and-relays-bfaba450
title: "DHCP servers and relays"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/2016-12-dhcp-servers-and-relays-bfaba450.md
source_anchor: ""
source_lines: [1, 76]
sha256: 482d4cd6622a35b2b35ebb2b88f5177e0b13561b7eaaf27180dfa61bad10d0fc
---

# DHCP servers and relays

*Source : https://www.fortinetguru.com/2016/12/dhcp-servers-and-relays/*

DHCP servers and relays
Note that DHCP server options are not available in transparent mode.
A DHCP server provides an address to a client on the network, when requested, from a defined address range. An interface cannot provide both a server and a relay for connections of the same type (regular or IPsec).
However, you can configure a Regular DHCP server on an interface only if the interface is a physical interface with a static IP address. You can configure an IPsec DHCP server on an interface that has either a static or a dynamic IP address.
You can configure one or more DHCP servers on any FortiGate interface. A DHCP server dynamically assigns IP addresses to hosts on the network connected to the interface. The host computers must be configured to obtain their IP addresses using DHCP.
If an interface is connected to multiple networks via routers, you can add a DHCP server for each network. The IP range of each DHCP server must match the network address range. The routers must be configured for DHCP relay.
You can configure a FortiGate interface as a DHCP relay. The interface forwards DHCP requests from DHCP clients to an external DHCP server and returns the responses to the DHCP clients. The DHCP server must have appropriate routing so that its response packets to the DHCP clients arrive at the unit.
DHCP Server configuration
To add a DHCP server, go to System > Network > Interface. Edit the interface, and select Enable for the DHCP Server row.
DHCP Server IP This appears only when Mode is Relay. Enter the IP address of the DHCP
server where the FortiGate unit obtains the requested IP address.
Address Range
By default, the FortiGate unit assigns an address range based on the address of the interface for the complete scope of the address. For example, if the interface address is 172.20.120.230, the default range cre- ated is 172.20.120.231 to 172.20.120.254. Select the range and select Edit to adjust the range as needed, or select Create New to add a dif- ferent range.
Netmask Enter the netmask of the addresses that the DHCP server assigns.
Default Gateway
Select to either use the same IP as the interface or select Specify and enter the IP address of the default gateway that the DHCP server assigns to DHCP clients.
DNS Server Select to use the system’s DNS settings or select Specify and enter the IP address of the DNS server.
Advanced... (expand to reveal more options)
Mode Select the type of DHCP server the FortiGate unit will be. By default, it is a server. Select Relay if needed. When Relay is selected, the above con- figuration is replaced by a field to enter the DHCP Server IP address.
Type Select to use the DHCP in regular or IPsec mode.
MAC Address Access Con- trol List
Select to match an IP address from the DHCP server to a specific client or device using its MAC address.
In a typical situation, an IP address is assigned ad hoc to a client, and that assignment times out after a specific time of inactivity from the client, known as the lease time. To ensure a client or device always has the same IP address, that is, there is no lease time, use IP reservation.
Add from DHCP Client List If the client is currently connected and using an IP address from the DHCP server, you can select this option to select the client from the list.
DHCP in IPv6
You can use DHCP with IPv6 using the CLI. To configure DHCP, ensure IPv6 is enabled by going to System > Config > Features and enable IPv6. Use the CLI command
config system dhcp6 server
For more information on the configuration options, see the CLI Reference.
Service
On low-end FortiGate units, a DHCP server is configured, by default on the Internal interface:
IP Range 192.168.1.110 to 192.168.1.210
Netmask 255.255.255.0
Default gateway 192.168.1.99
Lease time 7 days
DNS Server 1 192.168.1.99
These settings are appropriate for the default Internal interface IP address of 192.168.1.99. If you change this address to a different network, you need to change the DHCP server settings to match.
Alternatively, after the FortiGate unit assigns an address, you can go to System > Monitor > DHCP Monitor, locate the particular user. Select the check box for the user and select Add to Reserved.
Lease time
The lease time determines the length of time an IP address remains assigned to a client. Once the lease expires, the address is released for allocation to the next client request for an IP address The default lease time is seven days. To change the lease time, use the following CLI commands:
config system dhcp server
edit <server_entry_number>
set lease-time <seconds>
end
To have an unlimited lease time, set the value to zero.
DHCP options
When adding a DHCP server, you have the ability to include DHCP codes and options. The DHCP options are BOOTP vendor information fields that provide additional vendor-independent configuration parameters to manage the DHCP server. For example, you may need to configure a FortiGate DHCP server that gives out a separate option as well as an IP address. For example, an environment that needs to support PXE boot with Windows images.
The option numbers and codes are specific to the particular application. The documentation for the application will indicate the values to use. Option codes are represented in a option value/HEX value pairs. The option is a value 1 and 255.
You can add up to three DHCP code/option pairs per DHCP server.
To configure option 252 with value http://192.168.1.1/wpad.dat – CLI
config system dhcp server edit <server_entry_number>
set option1 252 687474703a2f2f3139322e3136382e312e312f777061642e646174
For detailed information about DHCP options, see RFC 2132, DHCP Options and BOOTP Vendor Extensions.
Exclude addresses in DHCP a range
If you have a large address range for the DHCP server, you can block a range of addresses that will not be included in the available addresses for the connecting users. To do this, go to the CLI and enter the commands:
config exclude-range
edit <sequence_number> set start-ip <address> set end-ip <address>
end end
DHCP Monitor
To view information about DHCP server connections, go to System > Monitor > DHCP Monitor. On this page, you can also add IP address to the reserved IP address list.
Breaking an address lease
Should you need to end an IP address lease, you can break the lease using the CLI. This is useful if you have limited addresses, longer lease times where leases are no longer necessary. For example, with corporate visitors.
To break a lease enter the CLI command:
execute dhcp lease-clear <ip_address>
Assigning IP address by MAC address
To prevent users in the from changing their IP addresses and causing IP address conflicts or unauthorized use of IP addresses, you can bind an IP address to a specific MAC address using DHCP.
Use the CLI to reserve an IP address for a particular client identified by its device MAC address and type of connection. The DHCP server then always assigns the reserved IP address to the client. The number of reserved addresses that you can define ranges from 10 to 200 depending on the FortiGate model.
After setting up a DHCP server on an interface by going to System > Network > Interface, select the blue arrow next to Advanced to expand the options. If you know the MAC address of the system select Create New to add it, or if the system has already connected, locate it in the list, select its check box and select Add from DHCP Client List.
You can also match an address to a MAC address in the CLI. In the example below, the IP address 10.10.10.55 for User1 is assigned to MAC address 00:09:0F:30:CA:4F.
config system dhcp reserved-address edit User1
set ip 10.10.10.55
set mac 00:09:0F:30:CA:4F
set type regular end

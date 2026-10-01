---
id: collect-261001-general-networking/general-networking/dhcp-servers-and-relays-1
title: "dhcp-servers-and-relays"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/dhcp-servers-and-relays.md
source_anchor: ""
source_lines: [1, 159]
sha256: 04860056a47138bbfb3b8534330cbeb16d5c3026b553ea788daed39ff60ee1ae
---

# dhcp-servers-and-relays

**DHC****P servers and relays**

Note that DHCP server options are not available in transparent mode.

A DHCP server provides an address to a client on the network, when requested, from a defined address range. An interface cannot provide both a server and a relay for connections of the same type (regular or IPsec).

However, you can configure a Regular DHCP server on an interface only if the interface is a physical interface with a static IP address. You can configure an IPsec DHCP server on an interface that has either a static or a dynamic IP address.

You can configure one or more DHCP servers on any FortiGate interface. A DHCP server dynamically assigns IP addresses to hosts on the network connected to the interface. The host computers must be configured to obtain their IP addresses using DHCP.

If an interface is connected to multiple networks via routers, you can add a DHCP server for each network. The IP range of each DHCP server must match the network address range. The routers must be configured for DHCP relay.

You can configure a FortiGate interface as a DHCP relay. The interface forwards DHCP requests from DHCP clients to an external DHCP server and returns the responses to the DHCP clients. The DHCP server must have appropriate routing so that its response packets to the DHCP clients arrive at the unit.


**DHC****P Server configuration**

To add a DHCP server, go to **S****ys****t****e****m > Network > Interface**. Edit the interface, and select **E****n****a****b****l****e** for the **DHC****P Server** row.


**DHC****P Server IP**                         This appears only when **Mod****e** is **R****e****l****ay**. Enter the IP address of the DHCP

server where the FortiGate unit obtains the requested IP address.


**A****dd****r****es****s Range**

By default, the FortiGate unit assigns an address range based on the address of the interface for the complete scope of the address. For example, if the interface address is 172.20.120.230, the default range cre- ated is 172.20.120.231 to 172.20.120.254. Select the range and select **E****d****i****t** to adjust the range as needed, or select **C****r****ea****t****e New** to add a dif- ferent range.


**N****e****t****m****as****k**      Enter the netmask of the addresses that the DHCP server assigns.


**D****e****f****a****u****l****t Gateway**

Select to either use the same IP as the interface or select **S****p****ec****i****f****y** and enter the IP address of the default gateway that the DHCP server assigns to DHCP clients.


**DN****S Server**    Select to use the system’s DNS settings or select **S****p****ec****i****f****y** and enter the IP address of the DNS server.


**A****d****va****n****ce****d****..****. (expand to reveal more options)**


**M****od****e**    Select the type of DHCP server the FortiGate unit will be. By default, it is a server. Select **R****e****l****a****y** if needed. When **R****e****l****a****y** is selected, the above con- figuration is replaced by a field to enter the **DHC****P Server IP** address.


**T****y****p****e**   Select to use the DHCP in regular or IPsec mode.


**M****A****C Address Access Con- trol List**

Select to match an IP address from the DHCP server to a specific client or device using its MAC address.


In a typical situation, an IP address is assigned ad hoc to a client, and that assignment times out after a specific time of inactivity from the client, known as the lease time. To ensure a client or device always has the same IP address, that is, there is no lease time, use IP reservation.


**A****d****d from DHCP Client List**      If the client is currently connected and using an IP address from the DHCP server, you can select this option to select the client from the list.


**DHC****P in IPv6**

You can use DHCP with IPv6 using the CLI. To configure DHCP, ensure IPv6 is enabled by going to **S****ys****t****e****m > Config > Features** and enable **I****P****v6**. Use the CLI command

config system dhcp6 server

For more information on the configuration options, see the CLI Reference.


**S****e****r****v****i****c****e**

On low-end FortiGate units, a DHCP server is configured, by default on the Internal interface:


**I****P Range**                                     192.168.1.110 to 192.168.1.210

**N****e****t****m****as****k**                                     255.255.255.0

**D****e****f****a****u****l****t gateway**                         192.168.1.99

**L****eas****e time**                                 7 days

**DN****S Server 1**                             192.168.1.99

These settings are appropriate for the default Internal interface IP address of 192.168.1.99. If you change this address to a different network, you need to change the DHCP server settings to match.

Alternatively, after the FortiGate unit assigns an address, you can go to **S****ys****t****e****m > Monitor > DHCP Monitor**, locate the particular user. Select the check box for the user and select **A****d****d to Reserved**.


**L****eas****e time**

The lease time determines the length of time an IP address remains assigned to a client. Once the lease expires, the address is released for allocation to the next client request for an IP address The default lease time is seven days. To change the lease time, use the following CLI commands:

config system dhcp server

edit <server_entry_number>

set lease-time <seconds>

end


To have an unlimited lease time, set the value to zero.


**DHC****P options**

When adding a DHCP server, you have the ability to include DHCP codes and options. The DHCP options are BOOTP vendor information fields that provide additional vendor-independent configuration parameters to manage the DHCP server. For example, you may need to configure a FortiGate DHCP server that gives out a separate option as well as an IP address. For example, an environment that needs to support PXE boot with Windows images.

The option numbers and codes are specific to the particular application. The documentation for the application will indicate the values to use. Option codes are represented in a option value/HEX value pairs. The option is a value 1 and 255.

You can add up to three DHCP code/option pairs per DHCP server.


**T****o configure option 252 with value** **h****tt****p****:****//****192****.****168****.****1****.****1****/****w****p****a****d****.****d****a****t** **– CLI**

config system dhcp server edit <server_entry_number>

set option1 252 687474703a2f2f3139322e3136382e312e312f777061642e646174

end


For detailed information about DHCP options, see RFC 2132, DHCP Options and BOOTP Vendor Extensions.


**E****xc****l****ud****e addresses in DHCP a range**

If you have a large address range for the DHCP server, you can block a range of addresses that will not be included in the available addresses for the connecting users. To do this, go to the CLI and enter the commands:

config system dhcp server edit <server_entry_number>

config exclude-range

edit <sequence_number> set start-ip <address> set end-ip <address>

end end

end


**DHC****P Monitor**

To view information about DHCP server connections, go to **S****ys****t****e****m > Monitor > DHCP Monitor**. On this page, you can also add IP address to the reserved IP address list.


**B****r****eak****i****n****g an address lease**

Should you need to end an IP address lease, you can break the lease using the CLI. This is useful if you have limited addresses, longer lease times where leases are no longer necessary. For example, with corporate visitors.


**T****o break a lease enter the CLI command:**

execute dhcp lease-clear <ip_address>


**A****ss****i****gn****i****n****g IP address by MAC address**


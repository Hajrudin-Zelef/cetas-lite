---
id: collect-261001-huawei/huawei/2019-09-advanced-dhcp-server-9dbc4ad4
title: "2019-09-advanced-dhcp-server-9dbc4ad4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "parameters"]
source: docs/RAG/collect-261001-huawei/2019-09-advanced-dhcp-server-9dbc4ad4.md
source_anchor: ""
source_lines: [1, 92]
sha256: c24b552a2ab40365e13a6b32c86df7281e42bbcb0549db1f0e38a9c3a207ef14
---

# 2019-09-advanced-dhcp-server-9dbc4ad4

**DHCP server**

A DHCP server provides an address from a defined address range to a client on the network, when requested.

You can configure one or more DHCP servers on any FortiGate interface. A DHCP server dynamically assigns IP addresses to hosts on the network connected to the interface. The host computers must be configured to obtain their IP addresses using DHCP.

You can configure a FortiGate interface as a DHCP relay. The interface forwards DHCP requests from DHCP clients to an external DHCP server and returns the responses to the DHCP clients. The DHCP server must have appropriate routing so that its response packets to the DHCP clients arrive at the unit.

**Configure DHCP on the FortiGate**

**To add a DHCP server on the GUI:**

1. Go to *Network > Interfaces* .
2. Edit an interface.
3. Enable the *DHCP Server* option and configure the settings.

**To add a DHCP server on the CLI:**

config system dhcp server edit 1 set dns-service default set default-gateway 192.168.1.2 set netmask 255.255.255.0 set interface “port1” config ip-range edit 1 set start-ip 192.168.1.1 set end-ip 192.168.1.1

next edit 2 set start-ip 192.168.1.3 set end-ip 192.168.1.254

next

end set timezone-option default set tftp-server “172.16.1.2”

next end

**DHCP options**

When adding a DHCP server, you can include DHCP codes and options. The DHCP options are BOOTP vendor information fields that provide additional vendor-independent configuration parameters to manage the DHCP server. For example, you might need to configure a FortiGate DHCP server that gives out a separate option as well as an IP address, such as an environment that needs to support PXE boot with Windows images.

The option numbers and codes are specific to the application. The documentation for the application indicates the values to use. Option codes are represented in a option value/HEX value pairs. The option is a value between 1 and 255.

You can add up to three DHCP code/option pairs per DHCP server.

**To configure option 252 with value http://192.168.1.1/wpad.dat using the CLI:**

config system dhcp server edit <server_entry_number>

set option1 252 687474703a2f2f3139322e3136382e312e312f777061642e646174 end

For detailed information about DHCP options, see RFC 2132, DHCP Options and BOOTP Vendor Extensions.

**Option-82**

DHCP option 82, also known as the DHCP relay agent information option, helps protect FortiGate against attacks such as spoofing (forging) of IP addresses and MAC addresses, and DHCP IP address starvation.

FG3H1E5818900749 (1) # show config reserved-address edit 1 set type option82 set ip 100.100.100.12 set circuit-id-type hex set circuit-id “00010102” set remote-id-type hex set remote-id “704ca5e477d6”

next

end

| FG3H1E5818900749 (1) # set |  | 
| type | DHCP reserved-address type. | 
| *ip | IP address to be reserved for the MAC address. | 
| circuit-id-type | DHCP option type. | 
| circuit-id | Option 82 circuit-ID of the client that will get the reserved IP address. | 
| remote-id-type | DHCP option type. | 
| remote-id | Option 82 remote-ID of the client that will get the reserved IP address. | 
| description | Description. | 

FortiGate-140D-POE (1) # set type

mac Match with MAC address. option82 Match with DHCP option 82.

FortiGate-140D-POE (1) # set circuit-id-type hex DHCP option in hex. string DHCP option in string.

FortiGate-140D-POE (1) # set remote-id-type hex DHCP option in hex. string DHCP option in string.

**Option-42**

This option specifies a list of the NTP servers available to the client by IP address.

FortiGate-140D-POE # config system dhcp server

FortiGate-140D-POE (server) # edit 2

FortiGate-140D-POE (2) # set ntp-service local IP address of the interface the DHCP server is added to becomes the client’s NTP server IP address. default Clients are assigned the FortiGate’s configured NTP servers. specify Specify up to 3 NTP servers in the DHCP server configuration.

FortiGate-140D-POE (2) # set ntp-service

FortiGate-140D-POE (2) # set ntp-server1

<class_ip> Class A,B,C ip xxx.xxx.xxx.xxx

FortiGate-140D-POE (2) # set ntp-server1 1.1.1.1

FortiGate-140D-POE (2) # set ntp-server2 2.2.2.2 FortiGate-140D-POE (2) # set ntp-server3 3.3.3.3

FortiGate-140D-POE (2) # end

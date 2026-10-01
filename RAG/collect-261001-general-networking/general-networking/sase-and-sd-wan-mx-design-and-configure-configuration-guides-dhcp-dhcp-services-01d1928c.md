---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-dhcp-dhcp-services-01d1928c
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-dhcp-dhcp-services-01d1928c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-dhcp-dhcp-services-01d1928c.md
source_anchor: ""
source_lines: [1, 49]
sha256: 6d03a488606488ab1fd42d3c37c22469a598b3c80624db5e0e1dda6e5f871350
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-dhcp-dhcp-services-01d1928c

DHCP Services
Click 日本語 for Japanese
Configuration
The Meraki MX Security Appliance and Z-series Teleworker Gateway both provide a fully-featured DHCP service when configured in Routed mode on the Security & SD-WAN/Teleworker Gateway > Configure > Addressing & VLANs page. You can enable and configure the DHCP service on each VLAN individually, or for the whole network if VLANs are not enabled. This is done on the Security & SD-WAN/Teleworker Gateway > Configure > DHCP page. The configuration options include:
- Client addressing
    
  - Choose Run a DHCP server to enable DHCP services on that particular VLAN
- Lease time
    
  - Length of time before the DHCP lease expires
- DNS nameservers
    
  - DNS servers that the DHCP server will instruct the clients to use
- Boot options:
    
  - Enables or disables DHCP options used for network booting (PXE boot)
- Boot next-server
    
  - Determines the server that network boot clients will download the boot file from. This can be an IP address or hostname. If a hostname is used, the network boot client must be able to resolve the hostname via DNS in order for network booting to work properly.
- Boot filename
    
  - Determines the full path and name of the boot file that network boot clients will use.
- DHCP options
    
  - Add custom DHCP options here. These options are usually required by VoIP phones or other special networking devices.
- Reserved IP ranges
    
  - IP ranges that are reserved and therefore will not be assigned to clients. Lists of Reserved IP ranges can be imported from a CSV file.
- Fixed IP assignments
    
  - IP addresses that are allocated to specific devices by MAC address to ensure that these devices always get the same IP address when they make a DHCP request. Lists of Fixed IPs can be imported from a CSV file.
Note that the largest DHCP pool the MX will serve is equivalent in size to a /19 subnet, even on a VLAN configured with a larger subnet
For Fixed IP assignments, client description can only contain letters, numbers, spaces, and simple punctuation (commas not supported).
The MX can also provide DHCP for statically routed subnets. This requires that the layer 3 switch or router between the MX and the DHCP clients have DHCP relay functionality configured to relay DHCP requests for that subnet to the MX. This is not to be confused with DHCP relay on the MX itself, which is described at the bottom of this page.
DHCP for static routes can be configured in the same manner as DHCP for VLANs. However, the IP that is relayed to on the switch is the interface IP on the MX, as static routes do not have interfaces. In addition to the other configuration options mentioned above, you will be required to specify the Gateway IP that clients on that subnet should be using. This will generally be the IP in that subnet that is assigned to the layer 3 switch or router between the MX and the DHCP clients.
If you are using another node as a relay agent, the agent must not use port 68 as the source port when sending DHCP traffic to the MX
DHCP Relay
If you want to forward DHCP requests for a configured subnet or VLAN to another DHCP server rather than serving DHCP on the MX, you can do so by choosing the Relay DHCP to another server option for Client addressing and entering the IP address of the DHCP server you wish to forward requests to. The DHCP relay server must be reachable in one of the following three ways:
- The DHCP server is in a local VLAN configured on the MX
- The DHCP server is in a subnet reachable via Meraki AutoVPN
- The DHCP server is in a subnet for which a static LAN route is configured on the MX
Live Monitoring
The DHCP leases displayed on the Security & SD-WAN > Monitor > Appliance status > DHCP page show the active DHCP leases on the MX. When new devices request DHCP addresses, leases are created and added to the table. This feature also indicates the time remaining on any given DHCP lease.
The following clients will not appear in the DHCP leases tool:
- Clients with statically assigned addresses
- Clients with fixed DHCP assignments
- Clients receiving an address from another DHCP server
- Clients that have not successfully received an address from the MX
- Clients connected over Client VPN

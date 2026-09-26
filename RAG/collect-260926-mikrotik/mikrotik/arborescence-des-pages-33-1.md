---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-33-1
title: "Description"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-33.md
source_anchor: ""
source_lines: [1, 143]
sha256: 6ec0671bd962e9704a579a0808436b8d8d31d3e3b6b4cbc7128055d5262a3112
---

# Description

RouterOS allows to create multiple Virtual Routing and Forwarding instances on a single router. This is useful for BGP-based MPLS VPNs. Unlike BGP VPLS, which is OSI Layer 2 technology, BGP VRF VPNs work in Layer 3 and as such exchange IP prefixes between routers. VRFs solve the problem of overlapping IP prefixes and provide the required privacy (via separated routing for different VPNs).

It is possible to set up vrf-lite setups or use multi-protocol BGP with VPNv4 address family to distribute routes from VRF routing tables - not only to other routers, but also to different routing tables in the router itself.

# Configuration

VRF table is created in **`/ip vrf`** menu. After the VRF config is created routing table mapping is added (a dynamic table with the same name is created). Each active VRF will always have a mapped routing table.

[admin@arm-bgp] /ip/vrf> print 
Flags: X - disabled; * - builtin 
 0  * name="main" interfaces=all  
[admin@arm-bgp] /routing/table> print 
Flags: D - dynamic; X - disabled, I - invalid; U - used 
 0 D   name="main" fib 

 Note that the order of the added VRFs is significant. To properly match which interface will belong to the VRF care must be taken to place VRFs in the correct order (matching is done starting from the top entry, just like firewall rules).



Let's look at the following example:

[admin@arm-bgp] /ip/vrf> print 
Flags: X - disabled; * - builtin 
 0  * name="main" interfaces=all 
 1    name="myVrf" interfaces=lo_vrf  

 Since the first entry is matching all the interfaces, the second VRF will not have any interfaces added. To fix the problem order of the entries must be changed.

[admin@arm-bgp] /ip/vrf> move 1 0
[admin@arm-bgp] /ip/vrf> print 
Flags: X - disabled; * - builtin 
 0    name="myVrf" interfaces=lo_vrf  
 1  * name="main" interfaces=all   

 Connected routes from the interfaces assigned to the VRF will be installed in the right routing table automatically.

For example, let's make an SSH service to listen for connections on the interfaces belonging to the VRF:

[admin@arm-bgp] /ip/service> set ssh vrf=myVrf 
[admin@arm-bgp] /ip/service> print 
Flags: X, I - INVALID
Columns: NAME, PORT, CERTIFICATE, VRF
#   NAME     PORT  CERTIFICATE  VRF     
0   telnet     23               main    
1   ftp        21                       
2   www        80               main    
3   ssh        22               myVrf
4 X www-ssl   443  none         main    
5   api      8728               main    
6   winbox   8291               main    
7   api-ssl  8729  none         main  

 Adding routes to the VRF is as simple as specifying the routing-table parameter when adding the route and specifying in which routing table to resolve the gateway by specifying @name after the gateway IP:

/ip route add dst-address=192.168.1.0/24 gateway=172.16.1.1@myVrf routing-table=myVrf

 Traffic leaking between VRFs is possible if the gateway is explicitly set to be resolved in another VRF, for example:

# add route in the myVrf, but resolve the gateway in the main table
/ip route add dst-address=192.168.1.0/24 gateway=172.16.1.1@main routing-table=myVrf
# add route in the main table, but resolve the gateway in the myVrf
/ip route add dst-address=192.168.1.0/24 gateway=172.16.1.1@myVrf

 # Supported features

Different services can be placed in specific VRF on which the service is listening for incoming or creating outgoing connections. By default, all services are using the `main` table, but it can be changed with a separate `vrf` parameter or by specifying the VRF name separated by "@" at the end of the IP address.

Below is the list of supported services.

| Feature | Support | Comment | 
|---|---|---|
| **BGP** | + |  /routing bgp template add name=bgp-template1 vrf=vrf1 /routing bgp vpls add name=bgp-vpls1 site-id=10 vrf=vrf1 /routing bgp vpn add label-allocation-policy=per-vrf vrf=vrf1 | 
| **E-mail** | + |  /tool e-mail set address=192.168.88.1 vrf=vrf1 | 
| **IP Services** | + | VRF is supported for `telnet` ,`www` ,`ssh` ,`www-ssl` ,`api` ,`winbox` ,`api-ssl` services. The`ftp` service does not support changing the VRF.  /ip service set telnet vrf=vrf1 | 
| **L2TP Client** | + |  /interface l2tp-client add connect-to=192.168.88.1@vrf1 name=l2tp-out1 user=l2tp-client  | 
| **MPLS** | + |  | 
| **Netwatch** | + |  /tool netwatch add host=192.168.88.1@vrf1 | 
| **NTP** | + |  /system ntp client set vrf=vrf1 /system ntp server set vrf=vrf1 | 
| **OSPF** | + |  /routing ospf instance add disabled=no name=ospf-instance-1 vrf=vrf1 | 
| **ping** | + |  /ping 192.168.88.1 vrf=vrf1 | 
| **RADIUS** | + |  /radius add address=192.168.88.1@vrf1 /radius incoming set vrf=vrf1 | 
| **RIP** | + |  /routing rip instance add name=rip-instance-1 vrf=vrf1 | 
| **RPKI** | + |  /routing rpki add vrf=vrf1 | 
| **SNMP** | + |  | 
| **EoIP** | + |  /interface eoip add remote-address=192.168.1.1@vrf1 | 
| **IPIP** | + |  /interface ipip  add remote-address=192.168.1.1@vrf1 | 
| **GRE** | + |  /interface gre  add remote-address=192.168.1.1@vrf1 | 
| **SSTP-client** | + |  /interface sstp-client  add connect-to=192.168.1.1@vrf1 | 
| **OVPN-client** | + |  /interface ovpn-client add connect-to=192.168.1.1@vrf1 | 
| **L2TP-ether** | + |  /interface l2tp-ether add connect-to=192.168.2.2@vrf | 
| **VXLAN** | + |  /interface vxlan add vni=10 vtep-vrf=vrf1 | 
| **Fetch** | + |  /tool/fetch address=10.155.28.236@vrf1 mode=ftp src-path=my_file.pcap user=admin password="" | 
| **DNS** | + Starting from RouterOS v7.21 | Specifies in which VRF router will listen DNS queries Specifies which VRF is used to contact upstream servers.  /ip dns set servers=8.8.8.8@vrf1 | 
| **DHCP-Relay** | + Starting from RouterOS v7.15 |  /ip dhcp-relay set dhcp-server-vrf=vrf1*If dhcp-client is in vrf - special parameter in* *"ip dhcp-relay" configuration is not needed* | 
| **Remote logging** | + Starting from RouterOS v7.19 |  /system logging action add name=remote1 remote=192.168.1.1 target=remote vrf=vrf1 | 

# VRF interfaces in firewall



Started from version 7.14 when interfaces are added in VRF - virtual VRF interface is created automatically. If it is needed to match traffic which belongs to VRF interface, VRF virtual interface should be used in firewall filters, for example:

/ip vrf add interfaces=ether5 name=vrf5
/ip firewall filter add chain=input in-interface=vrf5 action=accept

 If there are several interfaces in one VRF but it is needed to match only one of these interfaces - marks should be used. For example:

/ip vrf add interface=ether15,ether16 vrf=vrf1516
/ip firewall mangle
add action=mark-connection chain=prerouting connection-state=new in-interface=ether15 new-connection-mark=input_allow passthrough=yes 
/ip firewall filter
add action=accept chain=input connection-mark=input_allow

 # Examples

## Simple VRF-Lite setup

Let's consider a setup where we need two customer VRFs that require access to the internet:


/ip address
add address=172.16.1.2/24 interface=public
add address=192.168.1.1/24 interface=ether1
add address=192.168.2.1/24 interface=ether2
/ip route
add gateway=172.16.1.1
# add VRF configuration
/ip vrf
add name=cust_a interface=ether1 place-before 0
add name=cust_b interface=ether2 place-before 0
# add vrf routes
/ip route
add gateway=172.16.1.1@main routing-table=cust_a
add gateway=172.16.1.1@main routing-table=cust_b
# masquerade local source
/ip firewall nat add chain=srcnat out-interface=public action=masquerade

 It might be necessary to ensure that packets coming in the "public" interface can actually reach the correct VRF. 

This can be solved by marking new connections originated by the VRF customers and steering the traffic by routing marks of incoming packets on the "public" interface.


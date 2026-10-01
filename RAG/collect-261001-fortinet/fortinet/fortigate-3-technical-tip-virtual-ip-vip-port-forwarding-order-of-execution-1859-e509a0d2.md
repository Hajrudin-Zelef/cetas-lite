---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-virtual-ip-vip-port-forwarding-order-of-execution-1859-e509a0d2
title: "fortigate-3-technical-tip-virtual-ip-vip-port-forwarding-order-of-execution-1859-e509a0d2"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-virtual-ip-vip-port-forwarding-order-of-execution-1859-e509a0d2.md
source_anchor: ""
source_lines: [1, 4]
sha256: f7372defdb188d465e78367e26d858fadae05e55ed9149cd2864dd5208f87be0
---

# fortigate-3-technical-tip-virtual-ip-vip-port-forwarding-order-of-execution-1859-e509a0d2

Technical Tip: Virtual IP (VIP) port forwarding order of execution
| Description | This article describes the order of execution of Virtual IPs port forwarding, and how to change that order. | 
| Scope | FortiGate. | 
| Solution | In FortiGate Virtual IP (VIP) port forwarding priority goes from top to bottom and the Firewall Policy order to which these VIPs are applied does not matter. VIPs will only be checked if they are applied on at least one firewall policy. In the first example, FortiGate is forwarding WAN (DMZ) interface ports 442-444 to Loopback Interface 2 port 443 and port 443 to Loopback Interface 3 port 443. However, since the Loopback Interface 2 VIP comes first this will be triggered first. FortiGate Loopback Interfaces below:   FortiGate VIP and Firewall Policy are below:  The second example includes a VIP forwarding all ports and another VIP forwarding only a specific port. Although Lpk3 is more specific Lpk2 is triggered first because of VIP order. This order can be changed but only through CLI. The next image shows how to change the VIP order, in regards to the first example.  After using the command it can be observed the successful order change, over both CLI and GUI.    Virtual Servers: Although configured in a different GUI section, Virtual Servers are part of the VIP list. They can be moved relative to other Virtual IP addresses in the same way.    To verify the processing order of Virtual Servers, view the VIP list in the CLI.  VAN_DNAT # show firewall vip config firewall vip edit "Lpk3"         set extip 10.5.20.59 set mappedip "3.3.3.3" set extintf "dmz" set portforward enable set extport 443 set mappedport 443     next edit "Lpk2"         set extip 10.5.20.59 set mappedip "2.2.2.2" set extintf "dmz" set portforward enable set extport 442-444 set mappedport 443 set portmapping-type m-to-n     next edit "Virtual Server"         set type server-load-balance set server-type http set extip 10.5.20.59 set extintf "any" set extport 444 config realservers edit 1                     set ip 4.4.4.4 set port 443 next end next end   VAN_DNAT # config firewall vip VAN_DNAT (vip) # move "Virtual Server" before Lpk2 VAN_DNAT (vip) # end VAN_DNAT # Earlier Firmware Versions. In v7.0 and earlier, it was possible but not recommended to have overlapping virtual IP addresses. In these firmware versions, when Central NAT is disabled, it is not possible to re-order VIPs using the move command.  VAN_DNAT # move "Virtual Server" before Lpk2 VIP entry cannot be moved when central-nat is disabled. Command fail. Return code -651 Moving virtual IP objects in these firmware versions is only possible by removing the virtual IP objects from every firewall policy and reconfiguring them in the intended order.  Check the routing table. In cases where there is more than one ISP or more than one WAN interface, it is important to review the routing table to confirm that the selected WAN interface from the Virtual IP is active. Depending on the case, the selected interface can be as inactive or in a standby state awaiting the other interface to fail.  VAN_DNAT # get router info routing-table all VAN_DNAT # get router info routing-table database  Related articles:  |

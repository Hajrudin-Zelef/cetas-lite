---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-ipsec-tunnel-up-but-no-traffic-is-passing-though-2ff3edba
title: "fortigate-3-troubleshooting-tip-ipsec-tunnel-up-but-no-traffic-is-passing-though-2ff3edba"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-ipsec-tunnel-up-but-no-traffic-is-passing-though-2ff3edba.md
source_anchor: ""
source_lines: [1, 4]
sha256: 169f0f860f826ec741617f486f294c9ca1966d7692fd6a84635b11c4e5284d5e
---

# fortigate-3-troubleshooting-tip-ipsec-tunnel-up-but-no-traffic-is-passing-though-2ff3edba

Troubleshooting Tip: IPsec Tunnel up but no traffic is passing though the tunnel
| Description | This article describes how to handle a scenario where the IPsec Tunnel is up and traffic seems to be leaving FortiGate but is not reaching the remote end.  This article applies to all the possible scenarios mentioned below:  | 
| Scope | FortiGate. | 
| Solution | Follow these steps:    diagnose vpn ike gateway list name <tunnel_name> diagnose vpn tunnel list name <tunnel_name>    config vpn ipsec phase1-interface edit "VPN-Phase1" set nattraversal forced end  Make sure NAT-Traversal is also enabled on the remote end on a Third-party device.    diagnose vpn tunnel flush <tunnel_name> diagnose vpn ike gateway flush name <tunnel_name>  Or:  diagnose vpn ike gateway clear name <tunnel_name>   get router info routing-table details <destination-ip>  If the static route is configured using the 'Named Address' and in the routing table it is showing via a Physical interface, try configuring it using the specific subnet.  For guidance on how to disable NPU offloading, refer to the following article: Technical Tip: Useful filters for sniffer packet capture.  Host X (x.x.x.x) -> FGT-A (IPsec VPN) FGT-B -> (y.y.y.y) Host Y. For debug flow, run the following commands:  diagnose debug reset diagnose debug console timestamp enable diagnose debug flow filter addr x.x.x.x y.y.y.y and diagnose debug flow show iprope enable diagnose debug flow show function-name enable diagnose debug flow trace start 1000 diagnose debug enable To stop debug:  diagnose debug disable diagnose debug reset  For packet capture, run the command:  diagnose sniffer packet any "host x.x.x.x and host y.y.y.y" 4 0 l  To stop the capture, press Ctrl + C.  After initiating the above commands on the SSH session, try to initiate the traffic from source IP x.x.x.x to destination IP y.y.y.y.  Always try to take a packet capture for the destination network: Take the sniffer for the destination address. In this setup, the destination address is the SSL VPN IP after connecting to the VPN.  diagnose sniffer packet any " host y.y.y.y " 4 0 l   y.y.y.y    destination ip  Useful commands:  Diagnose VPN tunnel list: get router info routing-table database get router info routing-table details x.x.x.x <----- Where x.x.x.x is the source IP address. get router info routing-table details y.y.y.y <----- Where y.y.y.y is the destination IP address, or remote server IP that is being accessed. get router info routing-table all  Related articles: Troubleshooting Tip: First steps to troubleshoot connectivity problems to or through a FortiGate with sniffer, debug flow, session list, routing table Troubleshooting Tip: Issue with traffic not flowing through previously working IPsec VPN tunnel Troubleshooting Tip: Troubleshooting traffic not flowing through a previously working IPsec VPN tunnel |

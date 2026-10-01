---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-troubleshooting-vip-port-forwarding-97405-0369a9ef
title: "fortigate-3-technical-tip-troubleshooting-vip-port-forwarding-97405-0369a9ef"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2019-08-16"]
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-troubleshooting-vip-port-forwarding-97405-0369a9ef.md
source_anchor: ""
source_lines: [1, 84]
sha256: 6c86e7496c27a9a0d79d04aabf461457f4a12703704ff7c83ead67a2a4cf001e
---

# fortigate-3-technical-tip-troubleshooting-vip-port-forwarding-97405-0369a9ef

Technical Tip: Troubleshooting VIP (port forwarding)
Description
This article describes a step-by-step guide on how to verify and troubleshoot VIP port forwarding on the FortiGate.
Scope
FortiGate.
Solution
The following is a step-by-step guide providing details on useful debug commands that will help troubleshoot the VIP.
Step 1: Verify that the traffic is arriving at the FortiGate:
diagnose sniffer packet (portname) 'host 20.20.20.20 and port 23' 4 0 a
In this example, IP 10.10.10.10 is the public-facing interface of the FortiGate, and IP 20.20.20.20 is the public IP from which the client connects. The internal server is 192.168.1.1, and the FortiGate internal interface is internal with IP 192.168.1.99. The forwarded port is port 23.
diagnose sniffer packet wan1 'host 20.20.20.20 and port 23' 4 0 a 
interfaces=[wan1]
filters=[host 20.20.20.20 and port 23]
2019-08-16 06:50:11.124423 wan1 -- 20.20.20.20.53178 -> 10.10.10.10.23: syn 3664790935
The sniffer correctly sees the TCP SYN packet arriving at the FortiGate. 
If no traffic appears on the sniffer, it is either a traffic problem that cannot get to the FortiGate or a mistake in the filter used. It is possible that the client is not sending to the expected port. Wireshark can be used on the client PC to determine the destination port used by the service to which one is trying to connect.
Step 2: Verify that the FortiGate is accepting and forwarding the traffic by using the debug flow:
diagnose debug flow filter saddr 20.20.20.20
diagnose debug flow filter dport 23
diagnose debug flow show function-name enable
diagnose debug flow trace start 10
diagnose debug enable
What to look for:
Verify whether the traffic matches the VIP policy:
No match (FortiGate uses the default gateway):
id=20085 trace_id=1 func=vf_ip_route_input_common line=2574 msg="find a route: flag=80000000 gw-10.10.10.10 via root"
Match (FortiGate matches the VIP policy):
id=20085 trace_id=5 func=fw_pre_route_handler line=185 msg="VIP-192.168.1.1:23, outdev-wan1"
Verify whether the FortiGate is performing NAT correctly.
(Destination NAT performed from 10.10.10.10 to 192.168.1.1):
id=20085 trace_id=5 func=__ip_session_run_tuple line=3268 msg="DNAT 10.10.10.1 0:23->192.168.1.1:23"
(Traffic routed out from the internal FortiGate interface on the same subnet as 192.168.1.1):
id=20085 trace_id=5 func=vf_ip_route_input_common line=2574 msg="find a route: flag=00000000 gw-192.168.1.1 via internal"
Verify whether the traffic matches the firewall policy.
No match:
Example of a debug flow on traffic that does not match a VIP:
id=20085 trace_id=1 func=print_pkt_detail line=5347 msg="vd-root received a packet(proto=6, 20.20.20.20:53236->10.10.10.10:23) from wan1. flag [S], seq 3447622355, ack 0, win 64240"
id=20085 trace_id=1 func=init_ip_session_common line=5506 msg="allocate a new session-024df3d7"
id=20085 trace_id=1 func=vf_ip_route_input_common line=2574 msg="find a route: flag=80000000 gw-10.10.10.10 via root"
id=20085 trace_id=1 func=fw_local_in_handler line=402 msg="iprope_in_check() check failed on policy 0, drop"
Example of traffic matching a VIP:
id=20085 trace_id=5 func=print_pkt_detail line=5347 msg="vd-root received a packet(proto=6, 20.20.20.20:53301->10.10.10.10:23) from wan 1. flag [S], seq 1588647149, ack 0, win 64240"
id=20085 trace_id=5 func=init_ip_session_common line=5506 msg="allocate a new session-024dfa5e"
id=20085 trace_id=5 func=fw_pre_route_handler line=185 msg="VIP-192.168.1.1:23, outdev-wan1"
id=20085 trace_id=5 func=__ip_session_run_tuple line=3268 msg="DNAT 10.10.10.1 0:23->192.168.1.1:23"
id=20085 trace_id=5 func=vf_ip_route_input_common line=2574 msg="find a route: flag=00000000 gw-192.168.1.1 via internal"
id=20085 trace_id=5 func=fw_forward_handler line=743 msg="Allowed by Policy-2:"
Step 3: Verify that the VIP destination is sending traffic back.
Sometimes the FortiGate is correctly configured, and traffic is passing through. But the VIP destination might be configured to use a different gateway than the FortiGate, so the traffic is lost in the reply direction. This debug can also highlight some issues with the service on the destination host if the service is not listening on the correct port or is not running at all. 
To verify that the traffic is coming back from the VIP destination, use the sniffer:
diagnose sniffer packet internal 'host 192.168.1.1 and port 23' 4 0 a
interfaces=[internal]
filters=[host 192.168.1.1 and port 23]
2019-08-16 07:40:49.282710 internal -- 192.168.1.99.51606 -> 192.168.1.1.23: syn 3308869475
2019-08-16 07:40:49.282794 internal -- 192.168.1.1.23 -> 192.168.1.99.51606: syn 2526230265 ack 3308869476
In this short packet trace, the FortiGate internal interface (192.168.1.99) is sending the syn packet to the internal host 192.168.1.1, and the internal host is replying with the SYN ACK to the FortiGate internal interface. If the SYN packet sent from the FortiGate interface to the destination is not answered, there might be a problem with the destination host.
Step 4: Verify whether there is other traffic.
If there are still problems with the VIP connectivity after going through this guide, there might be some additional ports that might need to be opened for the service behind the VIP to work correctly. Both 'External Service port' and 'Mapped to Ipv4 ports' must be allowed on the VIP firewall policy. Sometimes the documentation for the service can provide an overview of the additional ports required. It is also possible to use the sniffer and debug flow tool to see if there is any additional traffic being dropped by the FortiGate. 
Sniffer: 
diagnose sniffer packet wan1 'host 20.20.20.20' 4 0 a
Example output:
2019-08-16 07:51:16.328998 wan1 -- 20.20.20.20.50287 -> 10.10.10.10.8080: syn 462174303
What can be observed here is that the client IP is sending a SYN packet to the destination port 8080.
The same process with debug flow:
diagnose debug flow filter saddr 20.20.20.20
diagnose debug flow show function-name enable
diagnose debug flow trace start 20
diagnose debug enable
id=20085 trace_id=17 func=print_pkt_detail line=5347 msg="vd-root received a packet(proto=6, 20.20.20.20:50314->10.10.10.10:8080) from wan1. flag [S], seq 2446654668, ack 0, win 64240"
id=20085 trace_id=17 func=init_ip_session_common line=5506 msg="allocate a new session-024e463e"
id=20085 trace_id=17 func=vf_ip_route_input_common line=2574 msg="find a route: flag=80000000 gw-10.10.10.10 via root"
id=20085 trace_id=17 func=fw_local_in_handler line=402 msg="iprope_in_check() check failed on policy 0, drop"
The disabled debug flow:
diagnose debug disable
To verify the port status for 10.10.10.10, run a telnet test from the FortiGate using the command below.
execute telnet 10.10.10.10 8080
In cases with a port forwarding configuration: to access HTTPS internal service with an external port that is not standard, the web browser may close the connection with an 'error 404' message.
If port 443 cannot be used as an external service port in the port-forwarding VIP, verify which port is supported by the browser. The alternative HTTPS ports 9433 or 8443 should be supported in most cases.
Note:
For mapped-addr set to FQDN in VIP, make sure FortiGate resolves FQDN to an IP address.
Related documents:
Troubleshooting Tip: Troubleshooting VIP issues
Troubleshooting Tip: How to Determine Virtual IP ID
Troubleshooting Tip: Debug flow show - "failed to match vid-1" on VIP configuration

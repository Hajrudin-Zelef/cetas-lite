---
id: collect-261001-general-networking/general-networking/support-forum-92-troubleshooting-ipsec-vpn-connection-attempts-83837-e3f82e37
title: "support-forum-92-troubleshooting-ipsec-vpn-connection-attempts-83837-e3f82e37"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2010-10-11"]
keywords: []
source: docs/RAG/collect-261001-general-networking/support-forum-92-troubleshooting-ipsec-vpn-connection-attempts-83837-e3f82e37.md
source_anchor: ""
source_lines: [1, 3]
sha256: 8399d2dfdd72d2d8c4dc3f5575563787378dbfd70b86bae80b8089bfe84d24c8
---

# support-forum-92-troubleshooting-ipsec-vpn-connection-attempts-83837-e3f82e37

Question
Troubleshooting IPsec-VPN connection attempts
Hello all,  I am a new to fortigate and I have came into a dead end in my attempts to establish a successful ipsec vpn connection. I would really appreciate any help.    On the fortigate unit an ipsec connection is configured as interface mode dialup-server, with certificate based authentication.  From the client side forticlient is used with proper certificates installed and matched configuration.    This is the output of the connection test:    Now install tunnels into kernel: 1  sys_get_local_gwy() called: [in] remote gw: (gateway ip). [in] Next hop: 0.0.0.0  sys_get_local_gwy() called: [in] remote gw: (gateway ip). [in] Next hop: 0.0.0.0  sys_get_local_gwy() called: [in] remote gw: (gateway ip). [in] Next hop: 0.0.0.0  sys_get_local_gwy() called: [in] remote gw: (gateway ip). [in] Next hop: 0.0.0.0  End installing tunnels  Got a kernel message  Detect local gateway for peer: (gateway ip)  sys_get_local_gwy() called: [in] remote gw: (gateway ip). [in] Next hop: 0.0.0.0  sys_get_local_gwy() called: [in] remote gw: (gateway ip). [in] Next hop: 10.0.0.138  Get sa_connect message...10.0.0.78->(gateway ip):500, natt_mode=0  Using new connection...natt_mode=0  Set connection name = (connection name).  Tunnel 10.0.0.78 ---> (gateway ip):500,natt_en=1 is starting negotiation  Will negotiate a normal SA    (client' s certificate name)  cert (client' s certificate name) found  set_phase1_id(): RSA -> ID_DER_ASN1_DN    and it holds on that last line..    This is the log messages from the fortigate unit:    Message meets Alert condition  date=2010-10-11 time=12:47:27 devname=(devname) device_id=(device_id) log_id=0101037128 type=event subtype=ipsec pri=error fwver=040004 vd=" root"  msg=" progress IPsec phase 1"  action=" negotiate"  rem_ip=(client' s public ip) loc_ip=(gateway ip) rem_port=885 loc_port=500 out_intf=" wan2"  cookies=" 23585904a0094a6a/0000000000000000"  user=" N/A"  group=" N/A"  xauth_user=" N/A"  xauth_group=" N/A"  vpn_tunnel=" IPSec-VPN"  status=failure init=remote mode=main dir=inbound stage=1 role=responder result=ERROR    Message meets Alert condition  date=2010-10-11 time=12:47:27 devname=(devname) device_id=F(device_id) log_id=0101037124 type=event subtype=ipsec pri=error fwver=040004 vd=" root"  msg=" IPsec phase 1 error"  action=" negotiate"  rem_ip=(client' s public ip) loc_ip=(gateway ip) rem_port=885 loc_port=500 out_intf=" wan2"  cookies=" 23585904a0094a6a/0000000000000000"  user=" N/A"  group=" N/A"  xauth_user=" N/A"  xauth_group=" N/A"  vpn_tunnel=" IPSec-VPN"  status=negotiate_error error_reason=peer SA proposal not match local policy peer_notif=N/A      Any ideas for the possible reasons for this?    How can I further, more granularly troubleshoot this? I know there are diagnose cli commands but they are not in the cli guide and from what I did try - I couldn' t get any output.

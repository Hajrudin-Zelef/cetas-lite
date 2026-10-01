---
id: collect-261001-general-networking/general-networking/t5-support-forum-bgp-neighbor-group-not-working-but-neighbor-works-m-p-327887-ebc86920
title: "t5-support-forum-bgp-neighbor-group-not-working-but-neighbor-works-m-p-327887-ebc86920"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-support-forum-bgp-neighbor-group-not-working-but-neighbor-works-m-p-327887-ebc86920.md
source_anchor: ""
source_lines: [1, 153]
sha256: 33a9816c955addee9409e50c1083bf1307769a3cba672a78094f9f5040b2f94a
---

# t5-support-forum-bgp-neighbor-group-not-working-but-neighbor-works-m-p-327887-ebc86920

BGP neighbor group not working, but neighbor works
I am setting up a Hub and Spoke ADVPN with BGP. The VPN tunnel itself work well. I can ping both way within the tunnel. However, a BGP routing cannot be not established unless I created a neighbor instead of a neighbor group in my Hub BGP settings. The Hub BGP settings even worked with a neighbor alone and without a neighbor group.
Hub: FortiGate 60F FW: 7.2.8
Spoke: FortiGate 40F FW: 7.2.8
Hub BGP (working):
config router bgp
set as 65000
set ibgp-multipath enable
set additional-path enable
config neighbor
edit "10.0.61.4"
set soft-reconfiguration enable
set remote-as 65000
set route-reflector-client enable
next
end
config neighbor-group
edit "Globe Hub"
set soft-reconfiguration enable
set remote-as 65000
set route-reflector-client enable
next
end
config neighbor-range
edit 1
set prefix 10.0.61.4 255.255.255.255
set max-neighbor-num 100
set neighbor-group "Globe Hub"
next
end
config network
edit 1
set prefix 192.168.0.0 255.255.255.0
next
end
config network6
edit 1
set prefix6 ::/128
next
end
......
end
Hub BGP (not working):
config router bgp
set as 65000
set ibgp-multipath enable
set additional-path enable
config neighbor-group
edit "Globe Hub"
set soft-reconfiguration enable
set remote-as 65000
set route-reflector-client enable
next
end
config neighbor-range
edit 1
set prefix 10.0.61.4 255.255.255.255
set max-neighbor-num 100
set neighbor-group "Globe Hub"
next
end
config network
edit 1
set prefix 192.168.0.0 255.255.255.0
next
end
config network6
edit 1
set prefix6 ::/128
next
end
config redistribute "connected"
......
end
Hub VPN Interface:
config system interface
edit "Globe Hub"
set vdom "root"
set ip 10.0.0.1 255.255.255.255
set allowaccess ping
set type tunnel
set remote-ip 10.0.0.254 255.255.0.0
set snmp-index 24
set interface "wan1"
next
end
Spoke BGP:
config router bgp
set as 65000
set ibgp-multipath enable
set additional-path enable
config neighbor
edit "10.0.0.1"
set soft-reconfiguration enable
set remote-as 65000
next
end
config network
edit 1
set prefix 192.168.15.0 255.255.255.0
next
end
config network6
edit 1
set prefix6 ::/128
next
end
config redistribute "connected"
end   
Spoke VPN Interface:
config system interface
edit "advpn-wan"
set vdom "root"
set ip 10.0.61.4 255.255.255.255
set allowaccess ping
set type tunnel
set remote-ip 10.0.0.1 255.255.0.0
set snmp-index 16
set interface "wan"
next
end
Hub Debug Info:
BGP: bgp_ih_on_close:8 delete ipc_handler=0x7f84a68480 for sock=27
BGP: bgp_ipc_server_accept:508 create ipc_handler=0x7f84a68480 for sock=27
BGP: bgp_ih_on_read:434 request type=5 len=24 vfid=0 start=0 count=5 flags=0x0
BGP: bgp_ih_on_read:485 response type=5 len=68 vfid=0 start=0 count=1 flags=0x0 total=1 ret=76
BGP: bgp_ih_on_close:8 delete ipc_handler=0x7f84a68480 for sock=27
BGP: [RIB] Scanning BGP Network Routes for VRF 0...
BGP: [RIB] Scanning BGP RIB for VRF 0...
BGP: [NETWORK] Accept Thread: Incoming conn from host 10.0.61.4 (FD=27 VRF=0)
BGP: [NETWORK] Accept Thread: 10.0.61.4 - No such Peer configured
BGP: bgp_ipc_server_accept:508 create ipc_handler=0x7f84a68480 for sock=27
BGP: bgp_ih_on_read:434 request type=4 len=24 vfid=0 start=0 count=10 flags=0x1
BGP: bgp_ih_on_read:485 response type=4 len=116 vfid=0 start=0 count=1 flags=0x1 total=1 ret=124
BGP: bgp_ih_on_close:8 delete ipc_handler=0x7f84a68480 for sock=27
BGP: bgp_ipc_server_accept:508 create ipc_handler=0x7f84a68480 for sock=27
BGP: bgp_ih_on_read:434 request type=5 len=24 vfid=0 start=0 count=5 flags=0x0
Spoke Debug Info:
BGP: [RIB] Scanning BGP Network Routes for VRF 0...
BGP: [RIB] Scanning BGP RIB for VRF 0...
BGP: [RIB] Scanning BGP Network Routes for VRF 0...
BGP: 10.0.0.1-Outgoing [FSM] State: Active Event: 9
BGP: 10.0.0.1-Outgoing [FSM] State: Connect Event: 9
BGP: 10.0.0.1-Outgoing [NETWORK] FD=26, Sock Status: 0-Success
BGP: 10.0.0.1-Outgoing [FSM] State: Connect Event: 17
BGP: 10.0.0.1-Outgoing [ENCODE] Msg-Hdr: Type 1
BGP: 10.0.0.1-Outgoing [ENCODE] Open: Ver 4 MyAS 65000 Holdtime 180
BGP: 10.0.0.1-Outgoing [ENCODE] Open: Msg-Size 69
BGP: 10.0.0.1-Outgoing [NETWORK] FD=26, Sock Status: 107-Transport endpoint is not connected
BGP: 10.0.0.1-Outgoing [FSM] State: OpenSent Event: 18
BGP: [RIB] Scanning BGP Network Routes for VRF 0...
BGP: [RIB] Scanning BGP RIB for VRF 0...
BGP: [RIB] Scanning BGP Network Routes for VRF 0...

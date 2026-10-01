---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1000128405-d95b131e-display-bgp-peer-b2e8d08b-1
title: "Display peer information."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1000128405-d95b131e-display-bgp-peer-b2e8d08b.md
source_anchor: ""
source_lines: [1, 124]
sha256: bb1a02269c98e76dab1bcc525aba85638f505457efb29fcafa84342dad46e473
---

# Display peer information.

Enterprise
display bgp [ vpnv4 vpn-instance vpn-instance-name ] peer [ { group-name | ipv4-address } log-info | [ ipv4-address ] verbose ]
display bgp [ instance instance-name ] peer { group-name | ipv4-address } log-info | [ ipv4-address ]
display bgp vpnv4 all peer [ [ ipv4-address ] verbose ]
display bgp ipv6 peer [ verbose ]
display bgp ipv6 peer ipv6-address { log-info | verbose }
display bgp ipv6 peer ipv4-address verbose
display bgp vpnv6 all peer [ [ ipv4-address ] verbose ]
display bgp vpnv6 vpn-instance vpn-instance-name peer [ { group-name | ipv6-address } log-info | [ ipv6-address ] verbose ]
| Parameter | Description | Value | 
|---|---|---|
| vpn-instance vpn-instance-name | Specifies the name of a VPN instance. | The value is a string of 1 to 31 case-sensitive characters, spaces not supported. When double quotation marks are used around the string, spaces are allowed in the string. The value _public_ is reserved and cannot be used as the VPN instance name. | 
| ipv4-address | Specifies the IPv4 address of a peer to be displayed. | It is in dotted decimal notation. | 
| group-name | Specifies the name of a peer group. | The name is a string of 1 to 47 case-sensitive characters, with spaces not supported. When double quotation marks are used around the string, spaces are allowed in the string. | 
| log-info | Indicates to display log information of the specified peer. | - | 
| verbose | Indicates to display detailed peer information. | - | 
| vpnv4 | Indicates to display information about peers in a VPNv4 instance. | - | 
| all | Indicates to display information about peers in all VPNv4 instances. | - | 
| ipv6-address | Specifies the IPv6 address of a peer to be displayed. | The prefix is a 32-digit hexadecimal number, in the format of X:X:X:X:X:X:X:X. | 
| vpnv6 | Indicates to display BGP VPNv6 peer information. | - | 
Usage Scenario
Precautions
BGP has multiple address families. By default, the display bgp peer command displays information about BGP peers in IPv4 unicast address family only. If you want to view information about BGP peers in another address family, you need to specify its address family parameter.
To view detailed information about a BGP peer, such as information about BGP timers, the number of sent and received routes, capacities supported, the number of sent and received BGP messages, and enabled functions, specify verbose in the command.
If log-info is specified in the command, log information about a BGP peer is displayed, including information about BGP peer flapping.
# Display peer information.
<HUAWEI> display bgp peer
 BGP Local router ID : 10.2.3.4
 local AS number : 10
 Total number of peers : 2                 
 Peers in established state : 1
  Peer        V    AS  MsgRcvd  MsgSent  OutQ  Up/Down       State PrefRcv
  10.1.1.1    4   100        0        0     0 00:00:07        Idle       0
  10.2.5.6    4   200       32       35     0 00:17:49 Established       0
| Item | Description | 
|---|---|
| BGP Local router ID | Indicates the ID of the BGP local router. | 
| local AS number | Indicates the local AS number. | 
| Total number of peers | Indicates the nuber of the peer. | 
| Peers in established state | Indicates the nuber of the peer in established state. | 
| Peer | Indicates the IP address of the peer. | 
| V | Indicates the BGP version used on the peer. | 
| AS | Indicates the AS number. | 
| MsgRcvd | Indicates the number of received messages. | 
| MsgSent | Indicates the number of sent messages. | 
| OutQ | Indicates the message to be sent to the specified peer. | 
| Up/Down | Indicates the period of time during which a BGP session keeps the current state. | 
| State | Status of the peer:  | 
| PrefRcv | Indicates the number of route prefixes sent from the peer. | 
# Display detailed information about the peer 10.2.2.9.
<HUAWEI> display bgp peer 10.2.2.9 verbose
         BGP Peer is 10.2.2.9,  remote AS 100
         Type: IBGP link
         BGP version 4, Remote router ID 10.1.1.1
         Update-group ID: 1 
         BGP current state: Established, Up for 00h57m53s
         BGP current event: RecvKeepalive
         BGP last state: Established
         BGP Peer Up count: 1
         Received total routes: 0
         Received active routes total: 0
         Advertised total routes: 2
         Port: Local - 42796        Remote - 179
         Configured: Active Hold Time: 180 sec   Keepalive Time:60 sec
         Received  : Active Hold Time: 180 sec
         Negotiated: Active Hold Time: 180 sec   Keepalive Time:60 sec
         Peer optional capabilities:
         Peer supports bgp multi-protocol extension
         Peer supports bgp route refresh capability
         Peer supports bgp 4-byte-as capability
         Address family IPv4 Unicast: advertised and received
 Received: 
                  Total messages                 60
                  Update messages                1
                  Open messages                  1
                  KeepAlive messages             58
                  Notification messages          0
                  Refresh messages               0
 Sent: 
                  Total messages                 61
                  Update messages                2
                  Open messages                  1
                  KeepAlive messages             58
                  Notification messages          0
                  Refresh messages               0
 Authentication type configured: None
 Last keepalive received: 2012/03/06 19:17:37
 Last keepalive sent    : 2012/03/06 19:17:37
 Last update    received: 2012/03/06 19:17:43
 Last update    sent    : 2012/03/06 19:17:37
 Minimum route advertisement interval is 15 seconds
 Optional capabilities:
 Route refresh capability has been enabled
 4-byte-as capability has been enabled 
 Listen-only has been configured
 Peer's BFD has been enabled
 Peer Preferred Value: 0 
 Routing policy configured:
 No routing policy is configured  
| Item | Description | 
|---|---|
| Type | Indicates the BGP link type, which can be IBGP or EBGP. | 
| BGP version | Indicates the BGP version. | 
| remote router ID | Indicates the router ID of the peer. | 
| Update-group ID | Indicates the ID of the update-group to which the peer belongs. | 
| BGP current state | Current state of BGP:  | 
| BGP current event | Indicates the current BGP event. | 
| BGP last state | Indicates the last BGP status, which may be Idle, Connect, Active, OpenSent, OpenConfirm, or Established. | 
| BGP Peer Up count | Indicates the flapping count of a BGP peer in a specified period of time. | 
| Received total routes | Indicates the number of received route prefixes | 
| Received active routes total | Indicates the number of received active route prefixes. | 
| Advertised total routes | Indicates the number of sent route prefixes. | 
| Port | Indicates the port number.  | 
| Configured | Indicates locally configured timers.  | 
| Received : Active Hold Time | Indicates the hold time on the peer. | 
| Negotiated : Active Hold Time | Indicates the hold time agreed between the BGP peers after capability negotiation. | 
| Address family IPv4 Unicast | Indicates the IPv4 unicast address family. | 
| Received | Indicates the number of packets received from a peer.  | 
| Sent | Indicates the number of messages sent to a peer.  | 
| Authentication type configured | Authentication type configured | 
| Last keepalive received | Indicates the time when the Keepalive packet is received last time. It can be in the following formats:  | 
| Last keepalive sent | Indicates the time when the Keepalive packet is sent last time. It can be in the following formats:  | 
| Last update received | Indicates the time when the Update packet is received last time. It can be in the following formats:  | 
| Last update sent | Indicates the time when the Update packet is sent last time. It can be in the following formats:  | 

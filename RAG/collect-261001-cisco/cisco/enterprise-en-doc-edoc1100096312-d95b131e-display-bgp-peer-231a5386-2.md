---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100096312-d95b131e-display-bgp-peer-231a5386-2
title: "Display peer information."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100096312-d95b131e-display-bgp-peer-231a5386.md
source_anchor: ""
source_lines: [139, 168]
sha256: c536d4daf0a7fc71f366721182ccc41bf0ce74264f2584349caf6ed32f4da9e5
---

# Display peer information.

 local AS number : 100
  Total number of peers : 1                 Peers in established state : 1
  Peer            V    AS  MsgRcvd  MsgSent  OutQ  Up/Down       State PrefRcv
  fe80::21          4   200       17       19     0 00:09:59 Established       3
# Display detailed information about IPv6 peers.
<HUAWEI> display bgp ipv6 peer fc00:1::1 verbose
                                                                                
        BGP Peer is fc00:1::1,  remote AS 65009                                    
        Type: IBGP link                                                         
        BGP version 4, Remote router ID 10.2.2.2                                 
        Update-group ID: 1                                                      
        BGP current state: Established, Up for 00h01m13s                        
        BGP current event: KATimerExpired                                       
        BGP last state: OpenConfirm                                             
        BGP Peer Up count: 1                                                    
        Received total routes: 1                                                
        Received active routes total: 0                                         
        Advertised total routes: 1                                              
        Port:  Local - 49152    Remote - 179                                    
        Configured: Connect-retry Time: 32 sec                                  
        Configured: Min Hold Time: 0 sec                                        
        Configured: Active Hold Time: 180 sec   Keepalive Time:60 sec           
        Received  : Active Hold Time: 180 sec                                   
        Negotiated: Active Hold Time: 180 sec   Keepalive Time:60 sec           
        Peer optional capabilities:                                             
        Peer supports bgp multi-protocol extension                              
        Peer supports bgp route refresh capability                              
        Peer supports bgp 4-byte-as capability                                  
        Address family IPv6 Unicast: advertised and received                    
 Received: Total 4 messages

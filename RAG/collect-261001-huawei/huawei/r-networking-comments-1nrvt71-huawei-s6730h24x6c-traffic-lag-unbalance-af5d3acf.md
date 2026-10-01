---
id: collect-261001-huawei/huawei/r-networking-comments-1nrvt71-huawei-s6730h24x6c-traffic-lag-unbalance-af5d3acf
title: "r-networking-comments-1nrvt71-huawei-s6730h24x6c-traffic-lag-unbalance-af5d3acf"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution", "full-duplex", "throughput"]
source: docs/RAG/collect-261001-huawei/r-networking-comments-1nrvt71-huawei-s6730h24x6c-traffic-lag-unbalance-af5d3acf.md
source_anchor: ""
source_lines: [1, 58]
sha256: 5244aa0084de1c001d121b2e7044a5f2353a911f90334f81a642cef941d0250f
---

# r-networking-comments-1nrvt71-huawei-s6730h24x6c-traffic-lag-unbalance-af5d3acf

Huawei S6730-H24X6C Traffic LAG Unbalance 
        
        
        
    
    
    
      Hii all,
I have a pair of Huawei S6730-H24X6C switches running VRP (R) Software, Version 5.170 (V200R022C00SPC500), connected via a trunk link using a 2x10G LAG. MPLS services are running on these switches.
    
I noticed that inbound and outbound traffic is not balanced across both interfaces in the LAG, which causes one of the ports to become fully utilized. I have tried several load-balancing hash algorithms I found online, but the traffic just shifts back and forth between the two links without achieving proper distribution.
      I would really appreciate any suggestions or best practices to achieve a better load balance.
Below is the configuration of the LAG ports and the hashing algorithms I have tested on both switches:
    
      [Cable Pair]
LAG Port
SW-1 XGE0/0/21 <> SW-2 XGE0/0/24
SW-1 XGE0/0/22 <> SW-2 XGE0/0/23
    
      [Switch-1]Interface PHY Protocol InUti OutUti inErrors outErrorsEth-Trunk2 up up 5.65% 46.74% 0 0XGigabitEthernet0/0/21 up up 5.64% 0% 0 0XGigabitEthernet0/0/22 up up 5.66% 93.48% 0 0
    
      interface Eth-Trunk2port link-type trunkundo port trunk allow-pass vlan 1port trunk allow-pass vlan 99 980 to 981 2889 3269 3287 4015mode lacpload-balance enhanced profile LB-PROFILE
    
      load-balance-profile LB-PROFILEmpls field top-label sip dip
    
      [Switch-2]InUti/OutUti: input utility/output utilityInterface PHY Protocol InUti OutUti inErrors outErrorsEth-Trunk0 up up 46.24% 5.62% 0 0XGigabitEthernet0/0/23 up up 92.47% 5.60% 0 0XGigabitEthernet0/0/24 up up 0% 5.65% 0 0
    
      interface Eth-Trunk0port link-type trunkundo port trunk allow-pass vlan 1port trunk allow-pass vlan 99 980 to 981 2889 3269 3287 4015mode lacpload-balance enhanced profile LB-PROFILE
    
      load-balance-profile LB-PROFILEmpls field top-label sip dip
    
Section des commentaires
LAG is primarily a redundancy mechanism. Extra bandwidth is a happy coincidence most of the time, but if you really need that throughput, upgrade your links.
Is your DIP and SIP identical for all flows because they are encapsulated and you are trying to hash between LDP neighbor IPs?
Yes, the DIP and SIP are the MPLS LSR loopback addresses used for LDP sessions
This doesn't make much sense to me. The encapsulated packet, subject to a load-balance decision, would not know about the endpoints of the LDP session. These addresses are only known in the LIB (Label information base, control plane).
The data plane which handles the forwarding and load balancing, doesn't know about LDP sessions. The LFIB only holds the outer (or top) label.
I would disagree here. LAG is there to provide additional bandwidth (though of course it does it in the lane-adding way, not the faster rate way) in addition to redundancy. It's even in the current 802.1AX standard:
Link Aggregation allows the establishment of full-duplex point-to-point links that have a higher aggregate bandwidth than the individual links that form the aggregation
This issue was solve after i used alogithm load balancing this:
load-balance-profile LB-PROFILE
ipv4 field dip protocol
mpls field top-label 2nd-label sip
what's the output for
display eth-trunk 0 load-balance?
something like
Maybe the number of combinations is low for the majority of packets. In that case you won't get additional bandwidth out of your LAG
Output
Eth-Trunk0's load-balance information:
Load-balance Configuration : SIP-XOR-DIP
Load-balance options used per-protocol:
L2 : Source XOR Destination MAC address, Vlan ID, Ethertype, Ingress-port
IPv4: Source XOR Destination IP address, Source XOR Destination TCP/UDP port
IPv6: Source XOR Destination IP address, Source XOR Destination TCP/UDP port
MPLS: Source XOR Destination IP address, Source XOR Destination TCP/UDP port
That seems to be Switch-2, IMHO the output doesn't match the config you posted.
What MPLS applications are in use? VPLS services do a cool thing to hashing and a whole standard was made to fix it. If it's application related and not local hash failures, check out
control word.

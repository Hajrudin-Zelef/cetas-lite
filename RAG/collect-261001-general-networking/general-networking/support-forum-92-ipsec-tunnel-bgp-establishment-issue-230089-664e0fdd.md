---
id: collect-261001-general-networking/general-networking/support-forum-92-ipsec-tunnel-bgp-establishment-issue-230089-664e0fdd
title: "support-forum-92-ipsec-tunnel-bgp-establishment-issue-230089-664e0fdd"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/support-forum-92-ipsec-tunnel-bgp-establishment-issue-230089-664e0fdd.md
source_anchor: ""
source_lines: [1, 35]
sha256: 7f9028042986bf10c673423bc006ea5c94a6550392904f0a1a6046827526565b
---

# support-forum-92-ipsec-tunnel-bgp-establishment-issue-230089-664e0fdd

Ipsec Tunnel - Bgp establishment issue
FortiGate-A ↔ FortiGate-B IPsec Tunnel – BGP TCP/179 Return Traffic Not Decrypting
Hi Fortinet Community,
I am troubleshooting a BGP connectivity issue over an IPsec tunnel.
Environment:
- FortiGate-A: FortiGate 60F
- FortiOS: 7.4.11 GA
- FortiGate-B: Remote FortiGate
- IPsec: IKE/IPsec with NAT-T
- BGP: TCP/179
Issue
The BGP session is not establishing over the IPsec tunnel.
Troubleshooting performed
- The IPsec tunnel is UP and DPD status is OK.
- NAT-T is enabled and UDP/4500 traffic is observed in both directions.
- FortiGate-A generates the BGP TCP/179 SYN and sends it through the IPsec tunnel.
- Flow debug confirms that the packet enters the IPsec interface and is encrypted successfully.
- On FortiGate-B, the SYN-ACK is generated and confirmed to be taking the correct return path toward FortiGate-A.
- On FortiGate-A, the return UDP/4500 packets from FortiGate-B are visible on the WAN interface.
- However, the IPsec dec:pkts counter does not increase when the return traffic is generated.
- The decrypted inner TCP/179 packet is not observed on the IPsec interface.
- NPU offload was disabled at the IPsec tunnel level, but the issue persists. The tunnel currently shows npu_flag=00 .
- MTU was also reduced during testing, but this did not resolve the issue.
Key observation
1FortiGate-A2   |3   | BGP SYN TCP/1794   ↓5IPsec encryption6   ↓7UDP/45008   ↓9FortiGate-B10   |11   | SYN-ACK TCP/17912   ↓13UDP/450014   ↓15FortiGate-A WAN interface16   |17   X18IPsec inbound decryption19   |20   X dec:pkts not increasing21   |22   X Inner TCP/179 not observed23
The current diagnose vpn tunnel list name <tunnel> output shows the IPsec SA is established, but the inbound dec:pkts counter remains unchanged while the outer UDP/4500 return packets are visible.
Current suspicion
The investigation is now focused on the inbound IPsec processing/decryption path, specifically:
- SPI/SA matching
- IPsec integrity/authentication
- Replay protection
- ESP-in-UDP/NAT-T processing
- Inbound IPsec SA/decryption handling
Has anyone experienced a similar issue on FortiGate 60F running FortiOS 7.4.11, where the remote FortiGate sends the expected SYN-ACK and UDP/4500 is visible on the local WAN interface, but the local IPsec dec:pkts counter does not increase?
Any relevant Fortinet KB/TAC references or recommended debug commands would be appreciated.

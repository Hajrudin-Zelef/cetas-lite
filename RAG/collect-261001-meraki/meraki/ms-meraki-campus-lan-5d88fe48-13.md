---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-13
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [425, 484]
sha256: 61c2011b1dbacc610ff4fbfc6156a5e1bdae042ce5384e6dcb13f93fd4bc6b87
---

# ms-meraki-campus-lan-5d88fe48

  - Core/Collapsed Core = 4096*
  - Distribution = 16384
  - Access = 61440
- All access ports on MS390 running 12.28 and higher will have Portfast enabled by default
- Designate the switch with the minimal changes (configuration, links up/down, etc) as the root bridge
- Root bridge should be in your distribution/core layer
- STP priority on your root bridge should be set to 4096*
- Ideally, the switch designated as the root should be one which sees minimal changes (config changes, link up/downs etc.) during daily operation
- Enable BPDU Guard on all access ports (Including ports connected to MR30H/36H)
- Enable Root Guard on your distribution/core switches on ports facing your access switches
- Enable Loop Guard on trunk ports connecting switches within the same layer (e.g. Trunk between two access switches that are both uplinked to the distribution/core layer)
- It is also recommended that Loop Guard be paired with Unidirectional Link Detection (UDLD)
- Keep your STP domain diameter to 7 hops as maximum
- With each hop, increase your STP priority such that it is less preferred than the previous hop
- It is recommended to couple STP with UDLD
- Remember that UDLD must be supported and enabled on both ends
- UDLD is supported on MS390 with firmware 15.3 and above (Please check firmware changelog for more info)
- UDLD is run independently on a per-switch basis, regardless of any stacking involved
- The Meraki implementation is fully interoperable with the one implemented in traditional Cisco switches
- UDLD can either be configured in Alert only or Enforce
- UDLD is fully compatible with Cisco switches (more info here)
- It is also recommended that Loop Guard be paired with Unidirectional Link Detection
* While it is acceptable to set the STP priority on the root bridge to 0, it allows for no room for modification when replacing a switch or changing the topology temporarily. Thus, setting it to 4096 gives you that flexibility
Further Guidance on Cisco interoperability and UDLD
- 
    Traditional Cisco equipment supports 'aggressive' and 'normal' UDLD modes. Meraki is able to implement similar functionality using just the 'normal' mode
- 
    In Alert only mode, the Meraki implementation generates a Dashboard alert and Event Log entry. Traffic is still forwarded when a UDLD-error state is seen while configured in Alert only mode
- 
    In Enforce mode, Meraki behavior is mostly comparable to Cisco's aggressive mode. Similar to 'disabling' the port, Meraki blocks all traffic, similar to the STP blocking state. This does not physically bring the link down, though, as in a traditional Cisco 'aggressive' configuration
STP in a Hybrid LAN
A hybrid LAN is a Wired LAN which consists of multi-vendor platforms. In many cases, each vendor has its own implementation of the STP protocol. In fact, some vendors will even slightly deviate from the protocol standard.
Since STP is all about preventing network loops and electing a root bridge by exchanging BPDUs across the network, it is vital that this process does not get interrupted across the different platforms in a hybrid environment. For instance, if a bridge is sending out BPDUs in a VLAN that is not allowed on the trunk connecting between the two bridges it may very well lead to a problem.
As such, it is very important to understand how STP operates on each switch and revise the vendor's documentation to understand the specifics of STP.
Meraki MS (except MS390) supports standard based 802.1W RSTP
Meraki MS390 supports standard based 802.1S MSTP in a single instance (Instance 0)
General Guidance for STP in a Hybrid LAN
- Please ensure that all your VLANs in your PVST+/Rapid-PVST domain are running STP
- All these VLANs should be allowed on all trunks
- Ensure that all VLANs have the same root bridge
- When using PVST/PVST+, ensure that the root bridge is in the Presides on a PVST switch
- Consider using Native VLAN 1 on all switches for the best results (Otherwise ensure Native VLAN consistency everywhere in your STP domain)
- When adding VLANs, be wary of the order of change that might push some ports into inconsistent state. As a rule of thumb, start from Core working your way downstream
- Do not leave your PVST+/Rapid-PVST Bridge priority to their default values and ensure consistency of the root location across all VLANs
- Bridge Priority 4096 can facilitate a root migration as opposed to priority 0
- With each switch hop increase the STP priority to be higher than last hop
- Where possible, avoid using default priority 32768
- Use STP guard Root Guard to protect your Root
- Use STP guard BPDU Guard to protect your STP domain from the access edge
- It is highly recommended to run MSTP in a Hybrid LAN where possible as this will reduce misconfiguration and eliminate chances of falling back to legacy STP (802.1D)
The following table provides some further guidelines on STP interoperability in a hybrid LAN network. Please follow the recommended design options and/or the guidance provided based on your specific implementation.
It is highly recommended to run the same STP protocol across all switches in your network where possible. The below design guidelines can help you to achieve better integration and performance results where running the same protocol is not possible however it requires that you understand the caveats and implications of each of the design options
| Non Meraki Switches | All MS Platforms (Except MS390) - 802.1W RSTP | MS390 - 802.1S MSTP (Instance 0 / Region 1 / Revision 1) | 
|---|---|---|
| STP | Not Recommended Option | Not Recommended Option | 
| RSTP | Recommended Option Behavior:  Guidelines:  conf t rstp single end  | Recommended Option Behavior:  Guidelines:  conf t rstp single end  | 
| PVST/PVST+ | Not Recommended Option Behavior:  Guidelines:  To avoid any issues with STP, it is recommended to convert the Cisco Catalyst environment to single instance MSTP. This will ensure maximum compatibility in the STP environment. | Not Recommended Option Behavior:  Guidelines:   If you choose to configure a native VLAN other than VLAN 1, please be wary of the order of change on your switches as this might cause the MS390 to go offline (e.g. WAN Edge - PVST+(a) - MS390 - PVST+(b), changing the native VLAN on PVST+(b) first can cause the MS390 to go offline) To avoid any issues with STP, it is recommended to convert the Cisco Catalyst environment to single instance MSTP. This will ensure maximum compatibility in the STP environment. Don't forget to configure your non-MS390 switches with the correct MSTP settings: spanning-tree mode mst spanning-tree mst configuration name region1 revision 1 spanning-tree mst 0 priority {stp priority value} It might be required to bounce ports on the non-MS390 switches after migrating to MSTP to ensure that the port type is correct (P2p) | 
| Rapid-PVST | Not Recommended Option Behavior:  Guidelines:  conf t rstp single end To avoid any issues with STP, it is recommended to convert the Cisco Catalyst environment to single instance MSTP. This will ensure maximum compatibility in the STP environment. | Recommended Option Behavior:  Guidelines:  To avoid any issues with STP, it is recommended to convert the Cisco Catalyst environment to single instance MSTP. This will ensure maximum compatibility in the STP environment. | 
| MSTP | Recommended Option Behavior:  Guidelines:   | BEST Recommended Option Behavior:  Guidelines:  spanning-tree mode mst spanning-tree mst configuration name region1 revision 1 spanning-tree mst 0 priority {stp priority value}  | 
To illustrate the behavior of the different switching platforms in a Hybrid STP domain, please refer to the following diagram which explains for a given topology the operational behavior and the considerations that need to be taken into account when designing your STP domain.

---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-12
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [388, 424]
sha256: 1e772e62574caa939af0a4c222567e52bec8eec329c93360fdf0b5e2cd3253e3
---

# ms-meraki-campus-lan-5d88fe48

- RSTP is enabled by default and should always be enabled. Disable only after careful consideration (Such as when the other side is not compatible with RSTP)
- MS switches will automatically place all access interfaces into EDGE mode. This will cause the interface to immediately transition the port into STP forwarding mode upon linkup (Please note that the port still participates in STP)
- Configure other switches in your network (where possible) in RSTP mode. Otherwise, please plan carefully for interoperability issues
You must set allow VLAN 1 on the trunk between MS switches and other switches as this is required for RSTP
- For a traditional multi-layer Campus LAN set the STP priority as follows:
    
  - Core/Collapsed Core = 4096*
  - Distribution = 16384
  - Access = 61440
- Designate the switch with the minimal changes (configuration, links up/down, etc) as the root bridge
- Root bridge should be in your distribution/core layer
- STP priority on your root bridge should be set to 4096*
- Ideally, the switch designated as the root should be one which sees minimal changes (config changes, link up/downs etc.) during daily operation
- Enable BPDU Guard on all access ports (Including ports connected to MR30H/36H)
- Enable Root Guard on your distribution/core switches on ports facing your access switches
- Enable Loop Guard on trunk ports connecting switches within the same layer (e.g. Trunk between two access switches that are both uplinked to the distribution/core layer)
- It is also recommended that Loop Guard be paired with Unidirectional Link Detection (UDLD)
- Keep your STP domain diameter to 7 hops as maximum
- With each hop, increase your STP priority such that it is less preferred than the previous-hop
- If you are running a routed access layer, it is recommended to set the uplink ports as access and keep STP enabled as a failsafe
- It is recommended to couple STP with UDLD
- Remember that UDLD must be supported and enabled on both ends
- UDLD is supported on the following MS platforms: MS22, MS42, MS120, MS125, MS210, MS220, MS225, MS250, MS320, MS350, MS355, MS390 (Running MS15.3 and above), MS400 series (Running MS10.10 and above)
- UDLD is run independently on a per-switch basis, regardless of any stacking involved
- The Meraki implementation is fully interoperable with the one implemented in traditional Cisco switches
- UDLD can either be configured in Alert only or Enforce
- UDLD is fully compatible with Cisco switches (more info here)
- It is also recommended that Loop Guard be paired with Unidirectional Link Detection
* While it is acceptable to set the STP priority on the root bridge to 0, it allows for no room for modification when replacing a switch or changing the topology temporarily. Thus, setting it to 4096 gives you that flexibility
MS390 Specific Guidance
- MS390s support MST in instance 0 / region 1 / revision 1
- MST is enabled by default and should always be enabled. Disable only after careful consideration (Such as when the other side is not compatible with MST)
- MS90 switches with 12.28.1+ supports portfast
- Please ensure that other switches in the STP domain are configured with MST (where possible) or alternatively with RSTP since it's backward compatible.
- It's required to have the same native VLAN configured for all switches in a STP domain as a switch will only send (or listen to) backward compatible BPDUs (e.g. PVST, PVST+) on its native VLAN (which is VLAN 1 by default). More information about Hybrid LAN in the below section.
- For a traditional multi-layer Campus LAN set the STP priority as follows:
    

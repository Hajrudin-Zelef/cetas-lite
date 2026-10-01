---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-33402927617047-unifi-switch-settings-1ca2d294-2
title: "hc-en-us-articles-33402927617047-unifi-switch-settings-1ca2d294"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-33402927617047-unifi-switch-settings-1ca2d294.md
source_anchor: ""
source_lines: [79, 104]
sha256: f9a159ae62c2b171cc95e83fe1d8e7313eaafd84b2cf819b2919c633e6726f9c
---

# hc-en-us-articles-33402927617047-unifi-switch-settings-1ca2d294

- Spanning Tree Protocol: Enables STP/RSTP on a specific port. Inherits global setting if not overridden.
- Egress Rate Limit: Caps outbound bandwidth from the port.
- LLDP-MED: Extends LLDP info for devices like VoIP phones. Enabled by default and recommended to leave on.
- Voice VLAN: Assigns a VLAN to VoIP phones via LLDP-MED while defaulting other traffic to the port’s native VLAN.
- QoS: Assigns traffic to specific switch queues for prioritization. Automatically set by ProAV profiles, or can be customized manually.
Device Settings
Location: Devices > [Switch] > Settings tab
These settings apply to the switch hardware itself. Availability may vary by model.
- IP Settings: Assign a static IP or use DHCP for management.
- Network Override: Manually set the management VLAN for a device. Ensure the correct VLAN is tagged or native on the upstream port; misconfiguration may require a physical reset.
- Override Global Switch Settings: Configure Jumbo Frames, Flow Control, 802.1X Control, and Spanning Tree Protocol independently for this switch.
- Priority: Set STP/RSTP bridge priority. This is important for loop prevention and performance.
- SNMP: Set SNMP Location and Contact. SNMP monitoring is enabled elsewhere in Settings.
Additional Management Actions
- Adjust LCM display settings, brightness, and night mode.
- Enable or disable Rack Multi-Screen Synchronization.
- Replace, locate, restart, or remove the switch.
- Load saved configuration, set a replacement device, or manually update firmware.
Best Practices
- Use VLANs to logically separate traffic, then configure switch behavior per VLAN.
- Enable DHCP Guarding to block unauthorized DHCP servers.
- Leave STP or RSTP enabled and assign STP priorities to manage loops.
- Only enable Jumbo Frames if all devices in the traffic path support them.
- Use Port Isolation for untrusted ports (e.g., IoT or guest VLANs).
- Monitor performance using the Insights tab.
- Use Port Profiles to apply consistent configurations across multiple switches.

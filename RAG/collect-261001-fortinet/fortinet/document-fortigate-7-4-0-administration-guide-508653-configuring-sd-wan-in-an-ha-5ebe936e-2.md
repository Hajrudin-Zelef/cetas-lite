---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-508653-configuring-sd-wan-in-an-ha-5ebe936e-2
title: "document-fortigate-7-4-0-administration-guide-508653-configuring-sd-wan-in-an-ha-5ebe936e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-508653-configuring-sd-wan-in-an-ha-5ebe936e.md
source_anchor: ""
source_lines: [68, 145]
sha256: 5b7b6207403e5f39abefb2f9efc13afbdcb68280a5a1afd83a007f79d3f2c6d3
---

# document-fortigate-7-4-0-administration-guide-508653-configuring-sd-wan-in-an-ha-5ebe936e

- 
                                                    Designate an interface as the trunk port from the CLI: config system interface edit lan1 set trunk enable next end
To configure the virtual VLAN switch, VLANs and interface settings in the CLI:
config system global
set virtual-switch-vlan enable
end
config system virtual-switch
edit "ISP1"
set physical-switch "sw0"
set vlan 10
config port
edit "wan"
next
end
next
edit "ISP2"
set physical-switch "sw0"
set vlan 20
config port
edit "a"
next
end
next
end
config system interface
edit "ISP1"
set vdom "root"
set ip 192.168.10.99 255.255.255.0
set allowaccess ping
set type hard-switch
next
edit "ISP2"
set vdom "root"
set ip 192.168.20.99 255.255.255.0
set allowaccess ping
set type hard-switch
next
end
config system interface
edit lan1
set trunk enable
next
end
To configure SD-WAN in the GUI:
- 
                                                    On the FortiGate, go to Network > SD-WAN, select the SD-WAN Zones tab, and click Create New > SD-WAN Member.
- 
                                                    In the Interface dropdown, select ISP1.
- 
                                                    Leave SD-WAN Zone set to virtual-wan-link.
- 
                                                    Enter the Gateway address 192.168.10.1.
- 
                                                    Click OK.
- 
                                                    Repeat these steps to add the second interface (ISP2) with the gateway 192.168.20.1.
- 
                                                    Click Apply.
- 
                                                    Create a health check: 
  - 
                                                            Go to Network > SD-WAN, select the Performance SLA tab, and click Create New.
  - 
                                                            Set Name to GW_HC.
  - 
                                                            Set Protocol to Ping and Servers to 8.8.8.8.
  - 
                                                            Set Participants to All SD-WAN Members.
  - 
                                                            Enable SLA Target and leave the default values.
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Create SD-WAN rules as needed. The SLA health check can be used to determine when the ISP connections are in or out of SLA, and to failover accordingly.
- 
                                                    Configure your firewall policy as needed for the virtual-wan-link SD-WAN zone.

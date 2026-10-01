---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-509828-vrf-routing-support-6512db43-2
title: "get router info bgp summary"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-509828-vrf-routing-support-6512db43.md
source_anchor: ""
source_lines: [192, 226]
sha256: d433cd3e999acfdebcff927d9b9d731d08f8b9253c550738a61cfac9a609bd9c
---

# get router info bgp summary

To add a VRF ID in a static route in the CLI:
- 
                                                    Configure the interface: config system interface
    edit port2
        set vrf 10
    next
end
- 
                                                    Add a static route to the VRF. For example, using blackhole: config router static
    edit 3
        set dst 0.0.0.0/0
        set blackhole enable
        set vrf 10
    next
end
A static route can also be added to the VRF when using an IPsec interface by enabling VPN ID with IPIP encapsulation. See SD-WAN segmentation over a single overlay for more information.
To add a static route to the VRF when using IPsec:
config vpn ipsec phase1-interface
    edit "vpn1"
        set interface "port2"
        set auto-discovery-receiver enable
        set encapsulation vpn-id-ipip
        set remote-gw 1.1.101.1
        set psksecret ******
    next
end
config router static
    edit 1
        set dst 10.32.0.0 255.224.0.0
        set device "vpn1"
        set vrf 10 
    next
end
                                            To check the routing table:
# get router info routing-table static

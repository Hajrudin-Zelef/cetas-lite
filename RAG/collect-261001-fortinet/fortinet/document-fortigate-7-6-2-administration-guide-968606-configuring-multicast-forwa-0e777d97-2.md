---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-2-administration-guide-968606-configuring-multicast-forwa-0e777d97-2
title: "Configuring multicast forwarding"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-2-administration-guide-968606-configuring-multicast-forwa-0e777d97.md
source_anchor: ""
source_lines: [173, 209]
sha256: 25efd3cb56a6943b0e737bb1200e733bdad4129b3d16dc420866db2a4ebde0d7
---

# Configuring multicast forwarding

When using multi-VDOM mode, it is important to avoid causing a multicast network loop by creating an  all-to-all multicast policy. By default, on models that support NPU virtual links, changing the `vdom-mode` to `multi-vdom` will create a pair of npu0_vlink0 and npu0_vlink1 interfaces in the same root VDOM. By virtue of the all-to-all multicast policy and the fact the npu0_vlink interfaces are virtually connected, it forms a multicast network loop.

Therefore, when using multi-VDOM mode:

1. 
                                                    Ensure there is no existing all-to-all multicast policy before changing to multi-VDOM mode.
2. 
                                                    If an all-to-all multicast policy must be defined, ensure that no two connected interfaces (such as npu0_vlink0 and npu0_vlink1) belong in the same VDOM.

###### This configuration will result in a multicast loop:

```
config system global
    set vdom-mode multi-vdom
end
config firewall multicast-policy
    edit 1
        set logtraffic enable
        set srcintf "any"
        set dstintf "any"
        set srcaddr "all"
        set dstaddr "all"
    next
end
show system interface
config system interface
    edit "npu0_vlink0"
        
```
**set vdom "root"**
        set type physical
    next
    edit "npu0_vlink1"
        **set vdom "root"**
        set type physical
    next
end

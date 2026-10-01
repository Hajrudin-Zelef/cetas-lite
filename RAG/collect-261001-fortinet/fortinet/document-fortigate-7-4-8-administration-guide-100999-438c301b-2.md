---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-administration-guide-100999-438c301b-2
title: "Hardware switch"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-administration-guide-100999-438c301b.md
source_anchor: ""
source_lines: [191, 231]
sha256: b0292138fc97ea502f51b4cc12ff9fb82110a3a32df26ec634e04a14f78b1d81
---

# Hardware switch

1. 
                                                    Configure a virtual switch to use port3 and port5: ```
config system virtual-switch
    edit "hw1"
        set physical-switch "sw0"
        config port
            edit "port3"
            next
            edit "port5"
            next
        end
    next
end
```
2. 
                                                    Enable STP for the virtual switch: ```
config system interface
    edit "hw1"
        set vdom "vdom1"
        set ip 6.6.6.1 255.255.255.0
        set allowaccess ping https ssh
        set type hard-switch
        
```
**set stp enable** set device-identification enable
        set lldp-transmission enable
        set role lan
        set snmp-index 55
        set ip-managed-by-fortiipam disable
    next
end
3. 
                                                    Disable STP on port5 by enabling it as an STP edge port: ```
config system interface
    
```
**edit "port5"** set vdom "vdom1"
        set type physical**set stp-edge enable** set snmp-index 9
    next
end
Port5 is enabled as an edge port with STP disabled. Port3 remains enabled for STP.

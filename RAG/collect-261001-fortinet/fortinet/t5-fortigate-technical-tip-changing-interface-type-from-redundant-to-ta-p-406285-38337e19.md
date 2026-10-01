---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-changing-interface-type-from-redundant-to-ta-p-406285-38337e19
title: "t5-fortigate-technical-tip-changing-interface-type-from-redundant-to-ta-p-406285-38337e19"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-changing-interface-type-from-redundant-to-ta-p-406285-38337e19.md
source_anchor: ""
source_lines: [1, 3]
sha256: 0a95c2e585ebb3ffc25a2723f089bca1cbfbb5e76f726687f0dff9947844150c
---

# t5-fortigate-technical-tip-changing-interface-type-from-redundant-to-ta-p-406285-38337e19

| Description | This article will describe the way to change the interface type from redundant to the interface type aggregate. | 
| Scope | FortiGate. | 
| Solution | It is not possible to change an interface type after creation from the GUI or CLI in FortiGate.  The interface can be deleted and recreated, but this may be a time-consuming task if the interface has are many references and dependencies.   CLI: **config system interface** edit "INT-TST" set vdom "root"         **set type redundant** set member "port6" "port7" "port8" "port9" "port10" next end  Instead of deleting and recreating the interface, the type can be changed by loading a modified configuration backup, but this requires downtime and should be considered a maintenance task. As part of restoring the configuration, the unit will reboot. Reboot times are platform-dependent.        **config system interface** edit "INT-TST" set vdom "root"         **set type aggregate** set member "port6" "port7" "port8" "port9" "port10" next end    **Related articles:** Technical Tip: Interface type which can be member of LACP Technical Tip: Link aggregation limitation and maximum supported number of interfaces Technical Tip: Unable to change IPsec tunnel type and getting the -9999 error |

---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-04-resources-dc-evpn-ext-format-html-b2f7ba4e
title: "Enable the new SD-WAN tunnel encapsulation format."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-04-resources-dc-evpn-ext-format-html-b2f7ba4e.md
source_anchor: ""
source_lines: [1, 16]
sha256: c1db71d1f0ffbe39b007c35cc2f4f723e638a3084b596ef5dc4ae40c3f1a7221
---

# Enable the new SD-WAN tunnel encapsulation format.

The evpn ext-format enable command enables the new SD-WAN tunnel encapsulation format.
The undo evpn ext-format enable command restores the default configuration.
By default, the new SD-WAN tunnel encapsulation format is not used.
Format
evpn ext-format enable
undo evpn ext-format enable
Parameters
None
Views
System view
Default Level
2: Configuration level
Usage Guidelines
In SD-WAN scenarios, you can run this command to enable the new SD-WAN tunnel encapsulation format, in which there is R=0 in the GRE encapsulation header.
Example
# Enable the new SD-WAN tunnel encapsulation format.

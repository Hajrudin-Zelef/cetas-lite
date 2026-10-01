---
id: collect-261001-huawei/huawei/nessus-compliance-checks-reference-content-huaweisyntax-htm-1583000b
title: "nessus-compliance-checks-reference-content-huaweisyntax-htm-1583000b"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/nessus-compliance-checks-reference-content-huaweisyntax-htm-1583000b.md
source_anchor: ""
source_lines: [1, 11]
sha256: a97f812d2ddd76d8cfdd650623b76b6cc717e83da5f5af92791ea873d0de9682
---

# nessus-compliance-checks-reference-content-huaweisyntax-htm-1583000b

Huawei VRP Syntax
The syntax for this plugin and an audit are as follows:
<custom_item>
description: "Huawei: Set super password"
info: "Set super password for management levels of 3-15."
solution: "In system view, run the following command to configure super
password super password level <level> encryption-type cipher
<password>"
reference: "SANS-CSC|10,PCI|2.2.4,COBIT5|BAI10.01,800-53|CM-2"
expect: "^super password level ([3-9]|1[0-5]) cipher"
</custom_item>

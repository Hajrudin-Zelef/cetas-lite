---
id: collect-261001-ia-llm/ia-llm/document-fortigate-7-2-0-administration-guide-899762-snmp-agent-54f75e6c
title: "SNMP agent"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention"]
source: docs/RAG/collect-261001-ia-llm/document-fortigate-7-2-0-administration-guide-899762-snmp-agent-54f75e6c.md
source_anchor: ""
source_lines: [1, 25]
sha256: 6454a4e00d35e67988a39ac3eb7f26067ab1fdd1e390ed00afc268e982a1e1c8
---

# SNMP agent

The SNMP agent sends SNMP traps originating on the FortiGate to an external monitoring SNMP manager defined in a SNMP community. The SNMP manager can monitor the FortiGate system to determine if it is operating properly, or if any critical events occurring.

The description, location, and contact information for this FortiGate system will be part of the information that the SNMP manager receives. This information is useful if the SNMP manager is monitoring many devices, and enables faster responses when the FortiGate system requires attention.

#### To configure the SNMP agent in the GUI:

1. Go to *System > SNMP* .
2. Enable *SNMP Agent* .
3. Enter a description of the agent.
4. Enter the location of the FortiGate unit.
5. Enter a contact or administrator for the SNMP Agent or FortiGate unit.
6. Click *Apply* .

#### To configure the SNMP agent in the CLI:

```
config system snmp sysinfo
    set status enable
    set description <string>
    set contact-info <string>
    set location <string>
end
```

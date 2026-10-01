---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-switches-lan-c9000-lyr2-fwd-flowcontrol-stormcontrol-flow-contro-39252bf2-2
title: "c-en-us-td-docs-switches-lan-c9000-lyr2-fwd-flowcontrol-stormcontrol-flow-contro-39252bf2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-switches-lan-c9000-lyr2-fwd-flowcontrol-stormcontrol-flow-contro-39252bf2.md
source_anchor: ""
source_lines: [101, 162]
sha256: 1f45720b28e9ad5e381cb2e62046f8a68fb1512981213f7dae2f3f10456ecc7b
---

# c-en-us-td-docs-switches-lan-c9000-lyr2-fwd-flowcontrol-stormcontrol-flow-contro-39252bf2

interface port-channelnumber
Example:
Device(config)# interface Port-channel1
Specifies the Port-channel interface to be configured, and enters interface configuration mode.
Configures the unknown unicast storm control level as a rising threshold percentage of the interface bandwidth or as a suppression
level in bits per second.
Step 8
storm-control action {shutdown | trap}
Example:
Device(config-if)# storm-control action shutdown
Example:
Device(config-if)# storm-control action trap
Specifies the action to take when the traffic threshold is exceeded.
Note
On Cisco C9550 Series Smart Switches, the storm-control {broadcast | multicast | unicast | unknown-unicast} level command supports only a single rising threshold value (0.00 to 100.00) or the bps option. Falling thresholds are not supported.
Configure Storm Control on a Physical Interface
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Step 2
configure terminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
interface interface-id
Example:
Device(config)# interface GigabitEthernet1/0/1
Specifies the physical interface to be configured, and enters interface configuration mode.
Configures the unknown unicast storm control level as a rising threshold percentage of the interface bandwidth or as a suppression
level in bits per second.
Step 8
storm-control action {shutdown | trap}
Example:
Device(config-if)# storm-control action shutdown
Example:
Device(config-if)# storm-control action trap
Defines the action the interface takes when the traffic threshold is exceeded. In this case, it sends an SNMP trap notification.
Configuring Storm Control Percentage or bps Options
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Step 2
configure terminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
interface interface-id
Example:
Device(config)# interface GigabitEthernet1/0/1
Specifies the interface on which to configure Storm Control, and enters interface configuration mode.

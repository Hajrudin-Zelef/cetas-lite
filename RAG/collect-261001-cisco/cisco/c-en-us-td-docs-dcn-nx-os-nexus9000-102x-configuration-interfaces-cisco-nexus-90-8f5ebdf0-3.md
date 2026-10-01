---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0-3
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0.md
source_anchor: ""
source_lines: [137, 151]
sha256: 3c2c79aac8257811bd279f7aa7de11070db7ed882fff5b4f0d14c6451c8b36c6
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0

                                       If STP is used without Bridge Assurance and if LACP is not used, use UDLD in normal mode on vPC orphan ports.
Peer-Keepalive Links and Messages
The Cisco NX-OS software uses the peer-keepalive link between the vPC peers to transmit periodic, configurable keepalive messages. You must have Layer 3 connectivity between the peer devices to transmit these messages; the system cannot bring up the vPC Peer-Link unless the peer-keepalive link is already up and running.
| Note | We recommend that you associate the vPC peer-keepalive link to a separate VRF mapped to a Layer 3 interface in each vPC peer device. If you do not configure a separate VRF, the system uses the management VRF and management ports by default. Do not use the vPC Peer-Link itself to send and receive vPC peer-keepalive messages. | 
Failure Detection and Keepalive Timers
If one of the vPC peer devices fails, the vPC peer device on the other side of the vPC Peer-Link senses the failure by not receiving any peer-keepalive messages. The default interval time for the vPC peer-keepalive message is 1 second, and you can configure the interval between 400 milliseconds and 10 seconds.
You can configure a hold-timeout value with a range of 3 to 10 seconds; the default hold-timeout value is 3 seconds. This timer starts when the vPC Peer-Link goes down. During this hold-timeout period, the secondary vPC peer device ignores vPC peer-keepalive messages, which ensures that network convergence occurs before a vPC action takes place. The purpose of the hold-timeout period is to prevent false-positive cases.
You can also configure a timeout value with a range of 3 to 20 seconds; the default timeout value is 5 seconds. This timer starts at the end of the hold-timeout interval. During the timeout period, the secondary vPC peer device checks for vPC peer-keepalive hello messages from the primary vPC peer device. If the secondary vPC peer device receives a single hello message, that device disables all vPC interfaces on the secondary vPC peer device.
Hold-Timeout vs. Timeout Parameters
The difference between the hold-timeout and the timeout parameters is as follows:
- 
                                    					
                                    During the hold-timeout, the vPC secondary device does not take any action based on any keepalive messages received, which prevents the system taking action when the keepalive might be received just temporarily, such as if a supervisor fails a few seconds after the vPC Peer-Link goes down.
- 
                                    					

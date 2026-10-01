---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546-15
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546.md
source_anchor: ""
source_lines: [676, 705]
sha256: 892f51fb5a2765f752057273d9016959bfead00c8faf63ab5a6332636a56c4c3
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546

                                 Auto enablement: Automatically enables trunk configuration on the authenticator switch, allowing user traffic from multiple VLANs coming from supplicant switches. Configure the cisco-av-pair as device-traffic-class=switch at the ISE. (You can configure this under the group or the user settings.)
| 1 | Workstations (clients) | 2 | Supplicant switch (outside wiring closet) | 
| 3 | Authenticator switch | 4 | Cisco ISE | 
| 5 | Trunk port |  |  | 
| Note | The switchport nonegotiate command is not supported on supplicant and authenticator switches with NEAT. This command should not be configured at the supplicant side of the topology. If configured on the authenticator side, the internal macros will automatically remove this command from the port. | 
Voice Aware 802.1x Security
| Note | To use voice aware IEEE 802.1x authentication, the switch must be running the LAN base image. | 
You use the voice aware 802.1x security feature to configure the switch to disable only the VLAN on which a security violation occurs, whether it is a data or voice VLAN. In previous releases, when an attempt to authenticate the data client caused a security violation, the entire port shut down, resulting in a complete loss of connectivity.
You can use this feature in IP phone deployments where a PC is connected to the IP phone. A security violation found on the data VLAN results in the shutdown of only the data VLAN. The traffic on the voice VLAN flows through the switch without interruption.
Common Session ID
Authentication manager uses a single session ID (referred to as a common session ID) for a client no matter which authentication method is used. This ID is used for all reporting purposes, such as the show commands and MIBs. The session ID appears with all per-session syslog messages.
The session ID includes:
- 
                                 The IP address of the Network Access Device (NAD)
- 
                                 A monotonically increasing unique 32 bit integer
- 
                                 The session start time stamp (a 32 bit integer)
This example shows how the session ID appears in the output of the show authentication command. The session ID in this example is 160000050000000B288508E5:
Device# show authentication sessions
Interface  MAC Address     Method   Domain   Status         Session ID
Fa4/0/4    0000.0000.0203  mab      DATA     Authz Success  160000050000000B288508E5
This is an example of how the session ID appears in the syslog output. The session ID in this example is also160000050000000B288508E5:
1w0d: %AUTHMGR-5-START: Starting 'mab' for client (0000.0000.0203) on Interface Fa4/0/4
AuditSessionID 160000050000000B288508E5
1w0d: %MAB-5-SUCCESS: Authentication successful for client (0000.0000.0203) on Interface
Fa4/0/4 AuditSessionID 160000050000000B288508E5
1w0d: %AUTHMGR-7-RESULT: Authentication result 'success' from 'mab' for client
(0000.0000.0203) on Interface Fa4/0/4 AuditSessionID 160000050000000B288508E5
The session ID is used by the NAD, the AAA server, and other report-analyzing applications to identify the client. The ID appears automatically. No configuration is required.

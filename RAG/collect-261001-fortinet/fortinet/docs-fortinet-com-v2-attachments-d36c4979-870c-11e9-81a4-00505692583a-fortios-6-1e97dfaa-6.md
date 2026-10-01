---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6-1e97dfaa-6
title: "docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["copyright", "warrants"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa.md
source_anchor: ""
source_lines: [656, 778]
sha256: 42cc2b9a524284d1eb0df95319939bfef94c501f341dc5a67f41a9ac1e99dca4
---

# docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa

Security best practices 23
Disable unused interfaces
To disable an interface from the GUI, go to Network > Interfaces. Edit the interface to be disabled and set Interface State
to Disabled.
From the CLI, to disable the port21 interface:
config system interface
edit port21
set status down
end
Disable unused protocols on interfaces
You can use the config system interface command to disable unused protocols that attackers may attempt to
use to gather information about a FortiGate unit. Many of these protocols are disabled by default. Using the config
system interface command you can see the current configuration of each of these options for the selected interface
and then choose to disable them if required.
config system interface
edit <interface-name>
set dhcp-relay-service disable
set pptp-client disable
set arpforward disable
set broadcast-forward disable
set l2forward disable
set icmp-redirect disable
set vlanforward disable
set stpforward disable
set ident-accept disable
set ipmac disable
set netbios-forward disable
set security-mode none
set device-identification disable
set lldp-transmission disable
end
Option Description
dhcp-relay-service Disable the DHCP relay service.
pptp-client Disable operating the interface as a PPTP client.
arpforward Disable ARP forwarding.
broadcast-forward Disable forwarding broadcast packets.
l2forward Disable layer 2 forwarding.
icmp-redirect Disable ICMP redirect.
vlanforward Disable VLAN forwarding.
stpforward Disable STP forwarding.
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 24
Option Description
ident-accept Disable authentication for this interface. The interface will not respond to a
connection with an authentication prompt.
ipmac Disable IP/MAC binding.
netbios-forward Disable NETBIOS forwarding.
security-mode Set to none to disable captive portal authentication. The interface will not
respond to a connection with a captive portal.
device-identification Disable device identification.
lldp-transmission Disable link layer discovery (LLDP).
Use local-in policies to close open ports or restrict access
You can also use local-in policies to close open ports or otherwise restrict access to FortiOS.
Close ICMP ports
Use the following command to close all ICMP ports on the WAN1 interface. The following example blocks traffic that
matches the ALL_ICMP firewall service.
config firewall local-in-policy
edit 1
set intf wan1
set scraddr all
set dstaddr all
set action deny
set service ALL_ICMP
set schedule always
end
Close the BGP port
Use the following command to close the BGP port on the wan1 interface. The following example blocks traffic that
matches the BGP firewall service.
config firewall local-in-policy
edit 1
set intf wan1
set scraddr all
set dstaddr all
set action deny
set service BGP
set schedule always
end
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Optional settings 25
Optional settings
This section describes settings that you can turn off so that you do not send any statistics to FortiGuard.
Collecting security statistics helps to enhance some FortiGuard services. The security statistics might be important to
customers who want to report to senior management. The collected information shows how your security is trending and
how your security ranks against industry peers.
Send malware statistics to FortiGuard
By default FortiOS periodically sends encrypted malware statistics to FortiGuard. The malware statistics record
Antivirus, IPS, or Application Control events. This data is used to improved FortiGuard services. The malware statistics
that FortiOS sends do not include any personal or sensitive customer data. The information is not shared with any
external parties and is used in accordance with Fortinet's Privacy Policy.
Sending the statistics to FortiGuard can be disabled. This will prevent FortiGate from sending malware statistics even if
the FortiGate receives updates from FortiManager instead of FortiGuard.
To disable sending malware statistics to FortiGuard:
config system global
set fds-statistics disable
end
Send Security Rating statistics to FortiGuard
Security Rating is a Fortinet Security Fabric feature that allows customers to audit their Security Fabric and find and fix
security problems. As part of the feature, FortiOS sends your security rating to FortiGuard every time a security rating
test runs.
For more information, see the white paper Proactive, Actionable Risk Management with the Fortinet Security Rating
Service at https://www.fortinet.com/content/dam/fortinet/assets/white-papers/wp-security-rating-service.pdf and the
security ratings updates at https://fortiguard.com/updates/secrating.
If you want, you can opt out of submitting Security Rating scores to FortiGuard. If you opt out, you won't be able to see
how your organization's scores compare with the scores of other organizations. Instead, an absolute score is shown.
To disable FortiGuard Security Rating result submission:
config system global
set security-rating-result-submission disable
end
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Copyright© 2022 Fortinet, Inc. All rights reserved. Fortinet®, FortiGate®, FortiCare® and FortiGuard®, and certain other marks are registered trademarks of Fortinet, Inc., in the
U.S. and other jurisdictions, and other Fortinet names herein may also be registered and/or common law trademarks of Fortinet. All other product or company names may be
trademarks of their respective owners. Performance and other metrics contained herein were attained in internal lab tests under ideal conditions, and actual performance and
other results may vary. Network variables, different network environments and other conditions may affect performance results. Nothing herein represents any binding
commitment by Fortinet, and Fortinet disclaims all warranties, whether express or implied, except to the extent Fortinet enters a binding written contract, signed by Fortinet’s
General Counsel, with a purchaser that expressly warrants that the identified product will perform according to certain expressly-identified performance metrics and, in such
event, only the specific performance metrics expressly identified in such binding written contract shall be binding on Fortinet. For absolute clarity, any such warranty will be
limited to performance in the same ideal conditions as in Fortinet’s internal lab tests. In no event does Fortinet make any commitment related to future deliverables, features or
development, and circumstances may change such that any forward-looking statements herein are not accurate. Fortinet disclaims in full any covenants, representations, and
guarantees pursuant hereto, whether express or implied. Fortinet reserves the right to change, modify, transfer, or otherwise revise this publication without notice, and the most
current version of the publication shall be applicable.

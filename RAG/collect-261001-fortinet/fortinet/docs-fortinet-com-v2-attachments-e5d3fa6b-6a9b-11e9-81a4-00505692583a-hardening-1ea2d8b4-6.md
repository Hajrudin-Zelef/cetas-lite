---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening-1ea2d8b4-6
title: "docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["copyright", "warrants"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4.md
source_anchor: ""
source_lines: [677, 744]
sha256: 6fa4cbf949b42f711d2fb32d0815f731c0fc57eb0096d2a3b87b454c88080523
---

# docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4

Security best practices 25
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
ident-accept Disable authentication for this interface. The interface will not respond to a
connection with an authentication prompt.
ipmac Disable IP/MAC binding.
netbios-forward Disable NETBIOS forwarding.
security-mode Set tonone to disable captive portal authentication. The interface will not
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
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 26
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
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Copyright© 2019 Fortinet, Inc. All rights reserved. Fortinet®, FortiGate®, FortiCare® and FortiGuard®, and certain other marks are registered trademarks of Fortinet, Inc., in
the U.S. and other jurisdictions, and other Fortinet names herein may also be registered and/or common law trademarks of Fortinet. All other product or company names may be
trademarks of their respective owners. Performance and other metrics contained herein were attained in internal lab tests under ideal conditions, and actual performance and
other results may vary. Network variables, different network environments and other conditions may affect performance results. Nothing herein represents any binding
commitment by Fortinet, and Fortinet disclaims all warranties, whether express or implied, except to the extent Fortinet enters a bindingwritten contract, signed by Fortinet’s
GeneralCounsel, with a purchaser that expressly warrants that the identified product will perform according to certain expressly-identified performance metrics and, in such
event, only the specific performance metrics expressly identified in such bindingwritten contract shall be bindingon Fortinet. For absolute clarity, any such warranty will be
limited to performance in the same ideal conditions as in Fortinet’s internal lab tests. In no event does Fortinet make any commitment related to future deliverables, features or
development, and circumstances may change such that any forward-looking statements herein are not accurate. Fortinet disclaims in full any covenants, representations, and
guarantees pursuant hereto, whether express or implied. Fortinet reserves the right to change, modify, transfer, or otherwise revise this publication without notice, and the most
current version of the publication shall be applicable.

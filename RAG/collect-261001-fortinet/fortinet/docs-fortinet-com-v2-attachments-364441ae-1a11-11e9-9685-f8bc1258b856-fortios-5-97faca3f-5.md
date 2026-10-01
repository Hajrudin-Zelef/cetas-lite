---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5-97faca3f-5
title: "docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["copyright", "warrants"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f.md
source_anchor: ""
source_lines: [594, 741]
sha256: 0cde8cfd3dd015e266c1bd2d9b55f649735e68e4f0997e5ec538bf8298a501dc
---

# docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f

Disable auto USB installation Security best practices
config system global
set fds-statistics disable
end
Disable auto USB installation
If USB installation is enabled, an attacker with physical access to a FortiGate could load a new configuration or
firmware on the FortiGate using the USB port. You can disable USB installation by entering the following from the
CLI:
config system auto-install
set auto-install-config disable
set auto-install-image disable
end
Set system time by synchronizing with an NTP server
For accurate time, use an NTP server to set system time. Synchronized time facilitates auditing and consistency
between expiry dates used in expiration of certificates and security protocols.
From the GUI go to System > Settings > System Time and select Synchronize with NTP Server. By default,
this causes FortiOS to synchronize with Fortinet's FortiGuard secure NTP server.
From the CLI you can use one or more different NTP servers:
config system ntp
set type custom
set ntpsync enable
config ntpserver
edit 1
set server <ntp-server-ip>
next
edit 2
set server <other-ntp-server-ip>
end
Enable password policies
Go to System > Settings > Password Policy, to create a password policy that all administrators must follow.
Using the available options you can define the required length of the password, what it must contain (numbers,
upper and lower case, and so on) and an expiry time.
Use the password policy feature to make sure all administrators use secure passwords that meet your
organization's requirements.
22 Hardening
Fortinet Technologies Inc.

Security best practices Configure auditing and logging
Configure auditing and logging
For optimum security go to Log & Report > Log Settings enable Event Logging. For best results send log
messages to FortiAnalyzer or FortiCloud.
From FortiAnalyzer or FortiCloud, you can view reports or system event log messages to look for system events
that may indicate potential problems. You can also view system events by going to FortiView > System Events.
Establish an auditing schedule to routinely inspect logs for signs of intrusion and probing.
Encrypt logs sent to FortiAnalyzer/FortiManager
To keep information in log messages sent to FortiAnalyzer private, go to Log & Report > Log Settings and when
you configure Remote Logging to FortiAnalyzer/FortiManager select Encrypt log transmission.
From the CLI.
config log {fortianalyzer | fortianalyzer2 | fortianalyzer3} setting
set enc-algorithm high
end
Disable interfaces that not used
To disable an interface from the GUI, go to Network > Interfaces. Edit the interface to be disabled and set
Interface State to Disabled.
From the CLI, to disable the port21 interface:
config system interface
edit port21
set status down
end
Disable unused protocols on interfaces
You can use the config system interface command to disable unused protocols that attackers may
attempt to use to gather information about a FortiGate unit. Many of these protocols are disabled by default. Using
the config system interface command you can see the current configuration of each of these options for
the selected interface and then choose to disable them if required.
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
Hardening
Fortinet Technologies Inc.
23

Use local-in policies to close open ports or restrict access Security best practices
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
security-mode Set to none to disable captive portal authentication. The interface will not respond
to a connection with a captive portal.
device-identification Disable device identification.
lldp-transmission Disable link layer discovery (LLDP).
Use local-in policies to close open ports or restrict access
You can also use local-in policies to close open ports or otherwise restrict access to FortiOS.
Close ICMP ports
Use the following command to close all ICMP ports on the WAN1 interface. The following example blocks traffic
that matches the ALL_ICMP firewall service.
config firewall local-in-policy
24 Hardening
Fortinet Technologies Inc.

Security best practices Use local-in policies to close open ports or restrict access
edit 1
set intf wan1
set srcaddr all
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
set srcaddr all
set dstaddr all
set action deny
set service BGP
set schedule always
end
Hardening
Fortinet Technologies Inc.
25

Copyright© 2022 Fortinet, Inc. All rights reserved. Fortinet®, FortiGate®, FortiCare® and FortiGuard®, and certain other marks are registered trademarks of Fortinet,
Inc., in the U.S. and other jurisdictions, and other Fortinet names herein may also be registered and/or common law trademarks of Fortinet. All other product or company
names may be trademarks of their respective owners. Performance and other metrics contained herein were attained in internal lab tests under ideal conditions, and
actual performance and other results may vary. Network variables, different network environments and other conditions may affect performance results. Nothing herein
represents any binding commitment by Fortinet, and Fortinet disclaims all warranties, whether express or implied, except to the extent Fortinet enters a binding written
contract, signed by Fortinet’s General Counsel, with a purchaser that expressly warrants that the identified product will perform according to certain expressly-identified
performance metrics and, in such event, only the specific performance metrics expressly identified in such binding written contract shall be binding on Fortinet. For
absolute clarity, any such warranty will be limited to performance in the same ideal conditions as in Fortinet’s internal lab tests. In no event does Fortinet make any
commitment related to future deliverables, features, or development, and circumstances may change such that any forward-looking statements herein are not accurate.
Fortinet disclaims in full any covenants, representations, and guarantees pursuant hereto, whether express or implied. Fortinet reserves the right to change, modify,
transfer, or otherwise revise this publication without notice, and the most current version of the publication shall be applicable.

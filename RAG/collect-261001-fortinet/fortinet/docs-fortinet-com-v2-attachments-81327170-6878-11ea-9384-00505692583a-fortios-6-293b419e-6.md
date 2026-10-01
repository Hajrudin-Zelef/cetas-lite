---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6-293b419e-6
title: "docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e.md
source_anchor: ""
source_lines: [607, 755]
sha256: 758498843cefa09f1c3074f6a0edb34d31fd615819f362e095eb783f523e8ddb
---

# docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e

Security best practices 22
Disable auto USB installation
If USB installation is enabled, an attacker with physical access to a FortiGate could load a new configuration or firmware
on the FortiGate using the USB port. You can disable USB installation by entering the following from the CLI:
config system auto-install
set auto-install-config disable
set auto-install-image disable
end
Set system time by synchronizing with an NTP server
For accurate time, use an NTP server to set system time. Synchronized time facilitates auditing and consistency
between expiry dates used in expiration of certificates and security protocols.
From the GUI go to System > Settings > System Time and select Synchronize with NTP Server. By default, this causes
FortiOS to synchronize with Fortinet's FortiGuard secure NTP server.
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
Disable the maintainer admin account
Administrators with physical access to a FortiGate appliance can use a console cable and a special administrator
account called maintainer to log into the CLI.The maintainer account allows you to log into a FortiGate if you have lost all
administrator passwords.
Once you have logged in with the maintainer account you can:
l Change the password of the admin administrator account (if it exists).
l Reset the FortiGate to the factory default configuration using the execute factoryreset command. This is the
only way to get access to the FortiGate if you have deleted the admin administrator account.
See the Fortinet knowledge base or Resetting a lost Admin password for details about using the maintainer account to
regain access to your FortiGate if you have lost all administrator account passwords.
The methodology for using the maintainer account is publicly available. As long as someone with physical access to the
device has the serial number of the device, which is labeled on the device, they can change the admin administrator
account password and access the FortiGate. This may be an unacceptable risk in some circumstances, especially
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 23
where the hardware is not physically secured. As an added security measure, the maintainer account can be disabled
using the following setting:
config system global
set admin-maintainer disable
end
If you disable this feature and lose your administrator passwords you will no longer be able to
log into your FortiGate.The only way to access your FortiGate will be to start over with a new
firmware installation and default configuration file. All of your settings will be lost.
Enable password policies
Go to System > Settings > Password Policy, to create a password policy that all administrators must follow. Using the
available options you can define the required length of the password, what it must contain (numbers, upper and lower
case, and so on) and an expiry time.
Use the password policy feature to make sure all administrators use secure passwords that meet your organization's
requirements.
Configure auditing and logging
For optimum security go to Log & Report > Log Settings enable Event Logging. For best results send log messages to
FortiAnalyzer or FortiCloud.
From FortiAnalyzer or FortiCloud, you can view reports or system event log messages to look for system events that
may indicate potential problems. You can also view system events by going to FortiView > System Events.
Establish an auditing schedule to routinely inspect logs for signs of intrusion and probing.
Encrypt logs sent to FortiAnalyzer/FortiManager
To keep information in log messages sent to FortiAnalyzer private, go to Log & Report > Log Settings and when you
configure Remote Logging to FortiAnalyzer/FortiManager select Encrypt log transmission.
From the CLI.
config log {fortianalyzer | fortianalyzer2 | fortianalyzer3} setting
set enc-algorithm high
end
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 24
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

Security best practices 25
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
FortiOS Hardening your FortiGate Fortinet Technologies Inc.


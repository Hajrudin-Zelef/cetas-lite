---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5-97faca3f-3
title: "docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f.md
source_anchor: ""
source_lines: [362, 474]
sha256: 57fa9d0a85ef26e8c158591bc944392b5257e389a133132b9fecc0964f310427
---

# docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f

Install the FortiGate unit in a physically secure location Security best practices
Security best practices
This chapter describes some techniques and best practices that you can use to improve FortiOS security.
Install the FortiGate unit in a physically secure location
A good place to start with is physical security. Install your FortiGate in a secure location, such as a locked room or
one with restricted access. A restricted location prevents unauthorized users from getting physical access to the
device.
If unauthorized users have physical access, they can disrupt your entire network by disconnecting your FortiGate
(either by accident or on purpose). They could also connect a console cable and attempt to log into the CLI. Also,
when a FortiGate unit reboots, a person with physical access can interrupt the boot process and install different
firmware.
Register your product with Fortinet Support
You need to register your Fortinet product with Fortinet Support to receive customer services, such as firmware
updates and customer support. You must also register your product for FortiGuard services, such as up-to-date
antivirus and IPS signatures. Register your product by visiting https://support.fortinet.com.
Keep your FortiOS firmware up to date
Always keep FortiOS up to date. The most recent version is the most stable and has the most bugs fixed and
vulnerabilities removed.
Fortinet periodically updates the FortiGate firmware to include new features and resolve important issues. After
you register your FortiGate or FortiOS VM, download firmware updates from the support web site,
https://support.fortinet.com.
Before you install any new firmware, be sure to follow these steps:
l Review the Release Notes for the latest firmware release.
l Review the Upgrade Paths Tool to determine the best path to take from your current version of FortiOS to the latest
version.
l Back up the current configuration.
Only FortiGate administrators who have read and write privileges can upgrade the FortiOS firmware.
16 Hardening
Fortinet Technologies Inc.

Security best practices System administrator best practices
System administrator best practices
This section describes a collection of changes you can implement to make administrative access to the GUI and
CLI more secure.
Disable administrative access to the external (Internet-facing) interface
When possible, don’t allow administration access on the external (Internet-facing) interface.
To disable administrative access, go to Network > Interfaces, edit the external interface and disable HTTPS,
PING, HTTP, SSH, and TELNET under Administrative Access.
From the CLI:
config system interface
edit <external-interface-name>
unset allowaccess
end
Allow only HTTPS access to the GUI and SSH access to the CLI
For greater security never allow HTTP or Telnet administrative access to a FortiGate interface, only allow HTTPS
and SSH access. You can change these settings for individual interfaces by going to Network > Interfaces and
adjusting the administrative access to each interface.
From the CLI:
config system interface
edit <interface-name>
set allowaccess https ssh
end
Require TLS 1.2 for HTTPS administrator access
Use the following command to require TLS 1.2 for HTTPS administrator access to the GUI:
config system global
set admin-https-ssl-versions tlsv1-2
end
TLS 1.2 is currently the most secure SSL/TLS supported version for SSL-encrypted administrator access.
Re-direct HTTP GUI logins to HTTPS
Go to System > Settings > Administrator Settings and enable Redirect to HTTPS to make sure that all
attempted HTTP login connections are redirected to HTTPS.
From the CLI:
config system global
set admin-https-redirect enable
end
Hardening
Fortinet Technologies Inc.
17

System administrator best practices Security best practices
Change the HTTPS and SSH admin access ports to non-standard ports
Go to System > Settings > Administrator Settings and change the HTTPS and SSH ports.
You can change the default port configurations for HTTPS and SSH administrative access for added security. To
connect to a non-standar port, the new port number must be included in the collection request. For example:
l If you change the HTTPS port to 7734, you would browse to https://<ip-address>:7734.
l If you change the SSH port to 2345, you would connect to ssh admin@<ip-address>:2345
To change the HTTPS and SSH login ports from the CLI:
config system global
set admin-sport 7734
set admin-ssh-port 2345
end
If you change to the HTTPS or SSH port numbers, make sure your changes do not conflict with ports used for
other services.
Maintain short login timeouts
Set the idle timeout to a short time to avoid the possibility of an administrator walking away from their management
computer and leaving it exposed to unauthorized personnel.
To set the administrator idle timeout, go to System > Settings and enter the amount of time for the Idle timeout.
A best practice is to keep the default time of 5 minutes.
To set the administrator idle timeout from the CLI:
config system global
set admintimeout 5
end
You can use the following command to adjust the grace time permitted between making an SSH connection and
authenticating. The range can be between 10 and 3600 seconds, the default is 120 seconds (minutes). By
shortening this time, you can decrease the chances of someone attempting a brute force attack a from being
successfull. For example, you could set the time to 30 seconds.
config system global
set admin-ssh-grace-time 30
end
Restrict logins from trusted hosts
Setting up trusted hosts for an administrator limits the addresses from where they can log into FortiOS. The
trusted hosts configuration applies to all forms of administrative access including HTTPS, SSH, ping, and SNMP.
When you identify a trusted host for an administrator account, FortiOS accepts that administrator’s login only from
one of the trusted hosts. A login, even with proper credentials, from a non-trusted host is dropped.
To identify trusted hosts, go to System > Administrators, edit the administrator account, enable Restrict login
to trusted hosts, and add up to ten trusted host IP addresses.
To add two trusted hosts from the CLI:
config system admin
edit <administrator-name>
set trustedhost1 172.25.176.23 255.255.255.255
18 Hardening
Fortinet Technologies Inc.


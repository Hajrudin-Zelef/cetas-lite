---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c-5
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c.md
source_anchor: ""
source_lines: [499, 659]
sha256: 32a452eed3a749d5e5a50ea95915226cd152fe387f3f0f6fbc7a093197e031bc
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c

associated with major translation types :
Established TCP sessions: 24 hours
UDP flow: 5 minutes
ICMP flow: 1 minute
The default timeout values are adequate to address the timeout requirements in most of the deployment scenarios. However,
these values can be adjusted/fine-tuned as appropriate. It is recommended not to configure very small timeout values (less
than 60 seconds) as it could result in high CPU usage. Refer the Best Practices for NAT Configuration section for more information.
Based on your configuration, you can change the timeouts described in this section.
If you need to quickly free your global IP address for a dynamic configuration, configure a shorter timeout than the default
timeout, by using the ip nat translation timeout command. However, the configured timeout should be longer than the other timeouts configured using commands specified in
the following steps.
If a TCP session is not properly closed by a finish (FIN) packet from both sides or during a reset, change the default TCP
timeout by using the ip nat translation tcp-timeout command.
SUMMARY STEPS
enable
configure terminal
ip nat translationseconds
ip nat translation udp-timeout seconds
ip nat translation tcp-timeoutseconds
ip nat translation finrst-timeoutseconds
ip nat translation icmp-timeout seconds
ip nat translation syn-timeout seconds
end
DETAILED STEPS
Command or Action
Purpose
Step 1
enable
Example:
Switch> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configure terminal
Example:
Switch# configure terminal
Enters global configuration mode.
Step 3
ip nat translationseconds
Example:
Switch(config)# ip nat translation 300
(Optional) Changes the amount of time after which NAT translations time out.
The default timeout is 24 hours, and it applies to the aging time for half-entries.
Step 4
ip nat translation udp-timeout seconds
Example:
Switch(config)# ip nat translation udp-timeout 300
(Optional) Changes the UDP timeout value.
Step 5
ip nat translation tcp-timeoutseconds
Example:
Switch(config)# ip nat translation tcp-timeout 2500
(Optional) Changes the TCP timeout value.
The default is 24 hours.
Step 6
ip nat translation finrst-timeoutseconds
Example:
Switch(config)# ip nat translation finrst-timeout 45
(Optional) Changes the finish and reset timeout value.
finrst-timeout—The aging time after a TCP session receives both finish-in (FIN-IN) and finish-out (FIN-OUT) requests or after
the reset of a TCP session.
Step 7
ip nat translation icmp-timeout seconds
Example:
Switch(config)# ip nat translation icmp-timeout 45
(Optional) Changes the ICMP timeout value.
Step 8
ip nat translation syn-timeout seconds
Example:
Switch(config)# ip nat translation syn-timeout 45
(Optional) Changes the synchronous (SYN) timeout value.
The synchronous timeout or the aging time is used only when a SYN request is received on a TCP session. When a synchronous
acknowledgment (SYNACK) request is received, the timeout changes to TCP timeout.
Step 9
end
Example:
Switch(config-if)# end
Exits interface configuration mode and returns to privileged EXEC mode.
Use SDM templates to configure system resources to optimize support for NAT.
After you set the template and the system reboots, you can use the show sdm prefer privileged EXEC command to verify the new template configuration. If you enter the show sdm prefer command before you enter the reload privileged EXEC command, the show sdm prefer command shows the template currently in use and the template that will become active after a reload.
Follow these steps to set the SDM template to maximize NAT usage:
SUMMARY STEPS
configure terminal
sdm prefer nat
end
write memory
reload
DETAILED STEPS
Command or Action
Purpose
Step 1
configure terminal
Example:
Switch# configure terminal
Enters global configuration mode.
Step 2
sdm prefer nat
Example:
Switch(config)# sdm prefer nat
Specifies the SDM template to be used on the switch.
This template is available under the network-advantage license.
Step 3
end
Example:
Switch(config)# end
Returns to the privileged EXEC mode.
Step 4
write memory
Example:
Switch# write memory
Save the current configuration before reload.
Step 5
reload
Example:
Switch# reload
Reloads the operating system.
Using Application-Level Gateways with NAT
NAT performs translation services on any TCP/UDP traffic that does not carry source and destination IP addresses in the application
data stream. Protocols that do not carry the source and destination IP addresses include HTTP, TFTP, telnet, archie, finger,
Network Time Protocol (NTP), Network File System (NFS), remote login (rlogin), remote shell (rsh) protocol, and remote copy
(rcp).
NAT Application-Level Gateway (ALG) enables certain applications that carry address/port information in their payloads to
function correctly across NAT domains. In addition to the usual translation of address/ports in the packet headers, ALGs take
care of translating the address/ports present in the payload and setting up temporary mappings.
Best Practices for NAT Configuration
In cases where both static and dynamic rules are configured, ensure that the local addresses specified in the rules do not
overlap. If such an overlap is possible, then the ACL associated with the dynamic rule should exclude the corresponding addresses
used by the static rule. Similarly, there must not be any overlap between the global addresses as this could lead to undesired
behavior.
Do not employ loose filtering such as permit ip any any in an ACL associated with NAT rule as this could result in unwanted packets being translated.
Do not share an address pool across multiple NAT rules.
Do not define the same inside global address in Static NAT and Dynamic Pool. This action can lead to undesirable results.
Exercise caution while modifying the default timeout values associated with NAT. Small timeout values could result in high
CPU usage.
Exercise caution while manually clearing the translation entries as this could result in the disruption of application sessions.
ALG packets traversing a NAT enabled interface will get punted to CPU, regardless of the packets being translated or not.
Therefore, it is recommended to use dedicated interface(s) just for NAT traffic. For all other types of traffic that does
not require NAT translation, use a different interface(s).
Troubleshooting NAT
This section explains the basic steps to troubleshoot and verify NAT.
Clearly define what NAT is supposed to achieve.
Verify that correct translation table exists using the show ip nat translation command.
Verify that timer values are correctly configured using the show ip nat translation verbose command.
Check the ACL values for NAT using the show ip access-list command
Check the overall NAT configuration using the show ip nat statistics command.
Use the clear ip nat translation command to clear the NAT translational table entires before the timer expires.
Use debug nat ip and debug nat ip detailed commands to debug NAT configuration.
Feature Information for Network Address Translation
The following table provides release information about the feature or features described in this module. This table lists
only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise,
subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco
Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
Table 1. Feature Information for NAT
Feature Name
Releases
Feature Information
Support for NetworkAddress Translation
Cisco IOS XE Gibraltar 16.10.1
This feature was introduced.
Network Address Translation (NAT) enables private IP networks that uses unregistered IP address to connect to the internet.

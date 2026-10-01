---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-ip-addressing-b-ip-addressing-m-sla-icmp-2405a64d-2
title: "c-en-us-td-docs-routers-ios-config-17-x-ip-addressing-b-ip-addressing-m-sla-icmp-2405a64d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-ip-addressing-b-ip-addressing-m-sla-icmp-2405a64d.md
source_anchor: ""
source_lines: [159, 230]
sha256: 366807b47e5837e88efffb8c5c368d2d489b277fecd2778cb1fd14245cb6a252
---

# c-en-us-td-docs-routers-ios-config-17-x-ip-addressing-b-ip-addressing-m-sla-icmp-2405a64d

is enabled, each operation response is checked for corruption. Use the
verify-data command with caution during normal operations because it generates unnecessary overhead.
Use the
debugipslatrace and
debugipslaerror commands to help troubleshoot issues with an IP SLAs operation.
What to Do Next
To add proactive threshold conditions and reactive triggering for generating traps (or for starting another operation) to
an IP Service Level Agreements (SLAs) operation, see the “Configuring Proactive Threshold Monitoring” section.
Configuration Examples for IP SLAs ICMP Path Jitter Operations
The following example shows the output when the ICMP Path Jitter operation is configured. Because the path jitter operation
does not support hourly statistics and hop information, the output for the
showipslastatistics command for the path jitter operation displays only the statistics for the first hop.
The following example shows the output when the ICMP Path Jitter operation is configured.
Device# configure terminal
Device(config)# ip sla 15011
Device(config-sla-monitor)# path-jitter 10.222.1.100 source-ip 10.222.3.100 num-packets 20
Device(config-sla-monitor-pathJitter)# frequency 30
Device(config-sla-monitor-pathJitter)# exit
Device(config)# ip sla schedule 15011 life forever start-time now
Device(config)# exit
Device# show ip sla statistics 15011
Round Trip Time (RTT) for Index 15011
Latest RTT: 1 milliseconds
Latest operation start time: 15:37:35.443 EDT Mon Jun 16 2008
Latest operation return code: OK
---- Path Jitter Statistics ----
Hop IP 10.222.3.252:
Round Trip Time milliseconds:
Latest RTT: 1 ms
Number of RTT: 20
RTT Min/Avg/Max: 1/1/3 ms
Jitter time milliseconds:
Number of jitter: 2
Jitter Min/Avg/Max: 2/2/2 ms
Packet Values:
Packet Loss (Timeouts): 0
Out of Sequence: 0
Discarded Samples: 0
Operation time to live: Forever
The Cisco Support and Documentation website provides online resources to download documentation, software, and tools. Use
these resources to install and configure the software and to troubleshoot and resolve technical issues with Cisco products
and technologies. Access to most tools on the Cisco Support and Documentation website requires a Cisco.com user ID and password.
Feature Information for IP
SLAs ICMP Path Jitter Operations
The following table provides release information about the feature or features described in this module. This table lists
only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise,
subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco
Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
Table 1. Feature Information for IP
SLAs ICMP Path Jitter Operations
Feature
Name
Releases
Feature
Information
IP SLAs
Path Jitter Operation
The Cisco
IOS IP SLAs Internet Control Message Protocol (ICMP) path jitter operation
allows you to measure hop-by-hop jitter (inter-packet delay variance).
IPSLA 4.0 -
IP v6 phase2
Support was
added for operability in IPv6 networks.
The
following commands are introduced or modified:
path- jitter,
show ip sla
configuration,
show ip sla
summary.

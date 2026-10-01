---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9-8
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9.md
source_anchor: ""
source_lines: [898, 982]
sha256: 8d1e24d5bb7b7a3ff9c9d02591be3bd448d2d273893316867735c6caf0547820
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9

Example: Configuring Standalone Mode on EtherChannel
This example shows how to configure the Standalone mode or Independent mode on an Port channel:
Device(config)# interface port-channel 1
Device(config-if)# no port-channel standalone-disable
Device(config-if)# end
This example shows how to verify the configuration of the Standalone mode on a Port
Channel interface:
Device# show running-config interface port-channel 1
Building configuration...
Current configuration:
!
interface Port-channel1
no ip address
no switchport
no port-channel standalone-disable
end
Example: Configuring Auto LAG
This example shows how to configure Auto-LAG on a switch
Device> enable
Device# configure terminal
Device(config)# port-channel auto
Device(config-if)# end
Device# show etherchannel auto
This example shows the summary of EtherChannel that was created automatically.
Device# show etherchannel auto
Flags: D - down P - bundled in port-channel
I - stand-alone s - suspended
H - Hot-standby (LACP only)
R - Layer3 S - Layer2
U - in use f - failed to allocate aggregator
M - not in use, minimum links not met
u - unsuitable for bundling
w - waiting to be aggregated
d - default port
A - formed by Auto LAG
Number of channel-groups in use: 1
Number of aggregators: 1
Group Port-channel Protocol Ports
------+-------------+-----------+-----------------------------------------------
1 Po1(SUA) LACP Gi1/0/45(P) Gi2/0/21(P) Gi3/0/21(P)
This example shows the summary of auto EtherChannel after executing the port-channel 1 persistent command .
Device# port-channel 1 persistent
Device# show etherchannel summary
Switch# show etherchannel summary
Flags: D - down P - bundled in port-channel
I - stand-alone s - suspended
H - Hot-standby (LACP only)
R - Layer3 S - Layer2
U - in use f - failed to allocate aggregator
M - not in use, minimum links not met
u - unsuitable for bundling
w - waiting to be aggregated
d - default port
A - formed by Auto LAG
Number of channel-groups in use: 1
Number of aggregators: 1
Group Port-channel Protocol Ports
------+-------------+-----------+-----------------------------------------------
1 Po1(SU) LACP Gi1/0/45(P) Gi2/0/21(P) Gi3/0/21(P)
Additional References for EtherChannels
Related Documents
Related Topic
Document Title
For complete syntax and usage information for the commands used in this chapter.
See the Layer 2/3 Commands section of theCommand Reference (Catalyst 9300 Series Switches)
Feature History for EtherChannels
This table provides release and related information for features explained in this module.
These features are available on all releases subsequent to the one they were introduced in, unless noted otherwise.
Release
Feature
Feature Information
Cisco IOS XE Everest 16.5.1a
EtherChannels
EtherChannel provides fault-tolerant high-speed links between switches, routers, and servers.
Cisco IOS XE Amsterdam 17.3.1
LACP 1:1 Redundancy and Dampening
The LACP 1:1 Redundancy feature supports an EtherChannel configuration with one active link and fast switchover to a hot-standby
link.
The LACP 1:1 Hot Standby Dampening feature configures a timer that delays switchover back to the higher priority port after
it becomes active.
Cisco IOS XE Dublin 17.10.1
Standalone Mode on Layer 3 Etherchannels
Support for configuring Standalone mode or Independent mode on Layer 3 EtherChannels was introduced in this release.
Use Cisco Feature Navigator to find information about platform and software image support. To access Cisco Feature Navigator,
go to http://www.cisco.com/go/cfn.

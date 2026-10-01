---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-11-configuration-g-ba637895-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-11-configuration-g-ba637895"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-11-configuration-g-ba637895.md
source_anchor: ""
source_lines: [95, 141]
sha256: eb9d6fe6db90ebf0a35e527d4c525936529355a4235661a5d8a77d08ac758bcc
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-11-configuration-g-ba637895

Use this command only after you initiate the traffic. show policy-map interface is updated with WRED configuration only after a traffic is sent.
You can add overlapping threshold pairs into the WRED configuration pairs.
Policy-map P1
Class CS
Random-detect dscp-based
Random-detect dscp CS1 percent 10 20 // WRED pair 1
Random-detect dscp CS2 percent 20 30 // WRED pair 2
Random-detect dscp CS3 percent 30 40 // WRED pair 3
Random-detect dscp CS4 percent 30 40 ==> belongs to WRED pair 3
Random-detect dscp CS5 percent 20 30 ==> belongs to WRED pair 2
Class-map match-any CS
match cs1
match cs2
match cs3
match cs4 >>
match cs5 >>
Default WRED pairs
If less than three WRED pairs are configured, any class-map filter participating WRED gets assigned to the third default WRED
pair with maximum threshold (100, 100).
Policy-map P1
Class CS
Random-detect dscp-based
Random-detect dscp CS1 percent 10 20 // WRED pair 1
Random-detect dscp CS2 percent 20 30 // WRED pair 2
Class-map match-any CS
match CS1
match CS2
match CS3
match CS4
In this case, classes CS3 and CS4 are mapped to WRED pair 3 with threshold (100, 100).
Rejection of Mismatched Configuration
If you configure random-detect without matching filters in a class-map, the policy installation is rejected.
This table provides release and related information for features explained in this module.
These features are available on all releases subsequent to the one they were introduced in, unless noted otherwise.
Release
Feature
Feature Information
Cisco IOS XE Everest 16.5.1a
Weighted Random Early Detection mechanism
WRED is a mechanism to avoid congestion in networks. WRED reduces the chances of tail drop by selectively dropping packets
when the output interface begins to show signs of congestion, thus avoiding large number of packet drops at once. You can
configure WRED to act based on any of the following values:
Differentiated Service Code Point
IP Precedence
Class of Service
Use Cisco Feature Navigator to find information about platform and software image support. To access Cisco Feature Navigator,
go to https://cfnng.cisco.com/.

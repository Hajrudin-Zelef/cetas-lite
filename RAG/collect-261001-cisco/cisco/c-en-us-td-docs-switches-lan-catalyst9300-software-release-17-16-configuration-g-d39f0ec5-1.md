---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5.md
source_anchor: ""
source_lines: [1, 76]
sha256: 42149681ef1f4a54c10be4bb4aabdeab6c9df7e36f20e41954c2e57dc48dbb8b
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
Multiprotocol Label Switching (MPLS) Virtual Private Network (VPN) services have been configured between the Provider Edge
(PE) routers and the customer edge (CE) routers at the customer sites.
Restrictions for EIGRP Prefix Limit Support
This feature is supported only under the IPv4 VRF address family and can be used only to limit the number of prefixes that
are accepted through a VRF.
The EIGRP Prefix Limiting Support feature is enabled only under the IPv4 VRF address-family. A peer that is configured to
send too many prefixes or a peer that rapidly advertises and then withdraws prefixes can cause instability in the network.
This feature can be configured to automatically reestablish a disabled peering session at the default or user-defined time
interval or when the maximum-prefix limit is not exceeded. However, the configuration of this feature alone cannot change
or correct a peer that is sending an excessive number of prefixes. If the maximum-prefix limit is exceeded, you will need
to reconfigure the maximum-prefix limit or reduce the number of prefixes that are sent from the peer.
Information About EIGRP Prefix Limit Support
The EIGRP Prefix Limit Support feature introduces the capability to limit the number of prefixes per VRF that are accepted
from a specific peer or to limit all prefixes that are accepted by an Enhanced Interior Gateway Routing Protocol (EIGRP) process
through peering and redistribution. This feature is designed to protect the local device from external misconfiguration that
can negatively impact local system resources; for example, a peer that is misconfigured to redistribute full Border Gateway
Protocol (BGP) routing tables into EIGRP. This feature is enabled under the IPv4 VRF address family and can be configured
to support the MPLS VPN Support for EIGRP Between Provider Edge and Customer Edge feature.
The EIGRP Prefix Limit Support feature provides the ability to configure a limit on the number of prefixes that are accepted
from EIGRP peers or learned through redistribution. This feature can be configured on per-peer or per-process basis and can
be configured for all peers and processes. This feature is designed to protect the local device from misconfigured external
peers by limiting the amount of system resources that can be consumed to process prefix updates.
Misconfigured VPN Peers
In MPLS VPNs, the number of routes that are permitted in the VPN routing and forwarding instance (VRF) is configured with
the maximumroutes VRF configuration command. However, limiting the number routes permitted in the VPN does not protect the local device from
a misconfigured peer that sends an excessive number of routes or prefixes. This type of external misconfiguration can have
a negative effect on the local device by consuming all available system resources (CPU and memory) in processing prefix updates.
This type of misconfiguration can occur on a peer that is not within the control of the local administrator.
Protecting the Device from External Peers
This feature can be configured to protect an individual peering session or protect all peering sessions. When this feature
is enabled and the maximum-prefix limit has been exceeded, the device will tear down the peering session, clear all routes
that were learned from the peer, and then place the peer in a penalty state for the default or user-defined time period. After
the penalty time period expires, normal peering will be reestablished.
Limiting the Number of Redistributed Prefixes
This feature can be configured to limit the number of prefixes that are accepted into the EIGRP topology table through redistribution
from the Routing Information Base (RIB). All sources of redistribution are processed cumulatively. When the maximum-prefix
limit is exceeded, all routes learned through redistribution are discarded and redistribution is suspended for the default
or user-defined time period. After the penalty time period expires, normal redistribution will occur.
Protecting the Device at the EIGRP Process Level
This feature can be configured to protect the device at the EIGRP process level. When this feature is configured at the EIGRP
process level, the maximum-prefix limit is applied to all peering sessions and to route redistribution. When the maximum-prefix
limit is exceeded, all sessions with the remote peers are torn down, all routes learned from remote peers are removed from
the topology and routing tables, all routes learned through redistribution are discarded, and redistribution and peering are
suspended for the default or user-defined time period.
Warning-Only Mode
The EIGRP Prefix Limit Support feature has two modes of operation. This feature can control peering and redistribution per
default and user-defined values or this feature can operate in warning-only mode. In warning-only mode the device will monitor
the number of prefixes learned through peering and/or redistribution but will not take any action when the maximum-prefix
limit is exceeded. Warning-only mode is activated only when the warning-only keyword is configured for any of the maximum-prefix limit commands. Only syslog messages are generated when this mode of
operation is enabled. Syslog messages can be sent to a syslog server or printed in the console. These messages can be buffered
or rate limited per standard Cisco IOS system logging configuration options.
Restart Reset and Dampening Timers and Counters
The EIGRP Prefix Limit Support feature provides two user-configurable timers, a restart counter, and a dampening mechanism.
When the maximum-prefix limit is exceeded, peering and/or redistribution is suspended for a default or user-defined time period.
If the maximum-prefix limit is exceeded too often, redistribution and/or peering will be suspended until manual intervention
is taken.
Restart Timer
The restart timer determines how long the router will wait to form an adjacency or accept redistributed routes from the RIB
after the maximum-prefix limit has been exceeded. The default restart-time period is 5 minutes.
Restart Counter
The restart counter determines the number of times a peering session can be automatically reestablished after the peering
session has been torn down or after the a redistributed routes have been cleared and relearned because the maximum-prefix
limit has been exceeded. The default restart-count limit is three.
Note
After the restart count limit has been crossed, you will need to enter the neighbor, or cleareigrpaddress-familyneighbor command to restore normal peering and redistribution.
Reset Timer
The reset timer is used to configure the device to reset the restart count to 0 after the default or configured reset-time
period has expired. This timer is designed to provide administrator with control over long-and medium-term accumulated penalties.
The default reset-time period is 15 minutes.
Dampening Mechanism
The dampening mechanism is used to apply an exponential decay penalty to the restart-time period each time the maximum-prefix
limit is exceeded. The half-life for the decay penalty is 150 percent of the default or user-defined restart-time value in
minutes. This mechanism is designed to identify and suppress unstable peers. It is disabled by default.
How to Configure the Maximum-Prefix Limit
Note

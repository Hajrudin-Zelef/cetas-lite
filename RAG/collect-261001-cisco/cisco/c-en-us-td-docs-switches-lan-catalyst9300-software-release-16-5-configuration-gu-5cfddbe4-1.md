---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "latency"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4.md
source_anchor: ""
source_lines: [1, 125]
sha256: ee120915855020bbfa25b14d826133bfda132afb561dd082a44419b605fffada
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
Your software release may not support all the features documented in this module. For the latest caveats and feature information,
see Bug Search Tool and the release notes for your platform and software release. To find information about the features documented
in this module, and to see a list of the releases in which each feature is supported, see the feature information table at
the end of this module.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature
Navigator, go to http://www.cisco.com/go/cfn. An account on Cisco.com is not required.
Restrictions on SLAs
This section lists the restrictions on SLAs.
The following are restrictions on IP SLAs network performance measurement:
The device does not support VoIP service levels using the gatekeeper registration delay operations measurements.
Only a Cisco IOS device can be a source for a destination IP SLAs responder.
You cannot configure the IP SLAs responder on non-Cisco devices and Cisco IOS IP SLAs can send operational packets only to
services native to those devices.
Cisco IOS IP SLAs send data
across the network to measure performance between multiple network locations or
across multiple network paths. They simulate network data and IP services and
collect network performance information in real time. Cisco IOS IP SLAs
generate and analyze traffic either between Cisco IOS devices or from a Cisco
IOS device to a remote IP device such as a network application server.
Measurements provided by the various Cisco IOS IP SLA operations can be used
for troubleshooting, for problem analysis, and for designing network
topologies.
Depending on the specific
Cisco IOS IP SLA operations, various network performance statistics are
monitored within the Cisco device and stored in both command-line interface
(CLI) and Simple Network Management Protocol (SNMP) MIBs. IP SLA packets have
configurable IP and application layer options such as source and destination IP
address, User Datagram Protocol (UDP)/TCP port numbers, a type of service (ToS)
byte (including Differentiated Services Code Point [DSCP] and IP Prefix bits),
Virtual Private Network (VPN) routing/forwarding instance (VRF), and URL web
address.
Because Cisco IP SLAs are Layer 2 transport
independent, you can configure end-to-end operations over disparate networks to
best reflect the metrics that an end user is likely to experience. IP SLAs
collect and analyze the
following performance metrics:
Delay (both round-trip and
one-way)
Jitter (directional)
Packet loss (directional)
Packet sequencing (packet
ordering)
Path (per hop)
Connectivity (directional)
Server or website download
time
Because Cisco IOS IP SLAs is
SNMP-accessible, it can also be used by performance-monitoring applications
like Cisco Prime Internetwork Performance Monitor (IPM) and other third-party
Cisco partner performance management products.
Using IP SLAs can provide the
following benefits:
Service-level agreement
monitoring, measurement, and verification.
Network performance
monitoring
Measurement of jitter,
latency, or packet loss in the network.
Continuous, reliable, and
predictable measurements.
IP service network health
assessment to verify that the existing QoS is sufficient for new IP services.
Edge-to-edge network
availability monitoring for proactive verification and connectivity testing of
network resources (for example, shows the network availability of an NFS server
used to store business critical data from a remote site).
Network operation
troubleshooting by providing consistent, reliable measurement that immediately
identifies problems and saves troubleshooting time.
Multiprotocol Label Switching
(MPLS) performance monitoring and network verification (if the
device
supports MPLS).
Network Performance Measurement with Cisco IOS IP SLAs
You can use IP SLAs to
monitor the performance between any area in the network—core, distribution, and
edge—without deploying a physical probe. It uses generated traffic to measure
network performance between two networking devices.
Figure 1. Cisco IOS IP SLAs
Operation. The following
figure shows how IP SLAs begin when the source device sends a generated packet
to the destination device. After the destination device receives the packet,
depending on the type of IP SLAs operation, it responds with time-stamp
information for the source to make the calculation on performance metrics. An
IP SLAs operation performs a network measurement from the source device to a
destination in the network using a specific protocol such as UDP.
Related Tasks Implementing IP SLA Network Performance MeasurementRelated References Restrictions on SLAs
IP SLA Responder and IP SLA Control Protocol
The IP SLA responder is a
component embedded in the destination Cisco device that allows the system to
anticipate and respond to IP SLA request packets. The responder provides
accurate measurements without the need for dedicated probes. The responder uses
the Cisco IOS IP SLA Control Protocol to provide a mechanism through which it
can be notified on which port it should listen and respond.
Note
The IP SLA responder can be a
Cisco IOS Layer 2, responder-configurable
device. The responder does not need to
support full IP SLA functionality.
The following figure shows
where the Cisco IOS IP SLA responder fits in the IP network. The responder
listens on a specific port for control protocol messages sent by an IP SLA
operation. Upon receipt of the control message, it enables the specified UDP or
TCP port for the specified duration. During this time, the responder accepts
the requests and responds to them. It disables the port after it responds to
the IP SLA packet, or when the specified time expires. MD5 authentication for
control messages is available for added security.
Figure 2. Cisco IOS IP
SLAs Operation
You do not need to enable the
responder on the destination device for all IP SLA operations. For example, a
responder is not required for services that are already provided by the
destination router (such as Telnet or HTTP).
Switches, controllers, and routers can take tens of milliseconds to process incoming packets due to other high priority processes.
This delay affects the response times because the test-packet reply might be in a queue while waiting to be processed. In
this situation, the response times would not accurately represent true network delays. IP SLAs minimize these processing delays
on the source device as well as on the target device (if the responder is being used) to determine true round-trip times.
IP SLA test packets use time stamping to minimize the processing delays.
When the IP SLA responder is enabled, it allows the target device to take time stamps when the packet arrives on the interface
at interrupt level and again just as it is leaving, eliminating the processing time. This time stamping is made with a granularity
of sub-milliseconds (ms).
Figure 3. Cisco IOS IP SLA Responder Time Stamping. The following figure demonstrates how the responder works. Four time stamps are taken to make the calculation for round-trip
time. At the target router, with the responder functionality enabled, time stamp 2 (TS2) is subtracted from time stamp 3 (TS3)
to produce the time spent processing the test packet as represented by delta. This delta value is then subtracted from the

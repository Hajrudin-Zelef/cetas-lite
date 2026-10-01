---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-ip-addressing-b-ip-addressing-m-sla-icmp-2405a64d-1
title: "c-en-us-td-docs-routers-ios-config-17-x-ip-addressing-b-ip-addressing-m-sla-icmp-2405a64d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-ip-addressing-b-ip-addressing-m-sla-icmp-2405a64d.md
source_anchor: ""
source_lines: [1, 158]
sha256: 2b6b8d9752f55bcaf66ec1e5c7462623b514bfb081971fc1493f7b1546b6601f
---

# c-en-us-td-docs-routers-ios-config-17-x-ip-addressing-b-ip-addressing-m-sla-icmp-2405a64d

IP Addressing Configuration Guide, Cisco IOS XE 17.x
Bias-Free Language
Bias-Free Language
The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
This document describes how to configure an IP Service Level Agreements (SLAs) Internet Control Message Protocol (ICMP) Path
Jitter operation to monitor hop-by-hop jitter (inter-packet delay variance). This document also demonstrates how the data
gathered using the Path Jitter operations can be displayed and analyzed using Cisco commands.
Before configuring any IP SLAs application, you can use the showipslaapplication command to verify that the operation type is supported on your software image.
In contrast with other IP SLAs operations, the IP SLAs Responder does not have to be enabled on either the target device or
intermediate devices for Path Jitter operations. However, the operational efficiency may improve if you enable the IP SLAs
Responder.
Restrictions for ICMP Path Jitter Operations
IP SLAs - ICMP Path Jitter is ICMP-based. ICMP-based operations can compensate for source processing delay but cannot compensate
for target processing delay. For more robust monitoring and verifying, we recommend that you use the IP SLAs UDP Jitter operation.
The jitter values obtained using IP SLAs - ICMP Path Jitter are approximates because ICMP does not provide the capability
to embed processing times on devices in the packet. If the target device does not place ICMP packets as the highest priority,
then the device will not respond properly. ICMP performance also can be affected by the configuration of priority queueing
on the device and by ping response.
A path jitter operation does not support hourly statistics and hop information.
Unlike other IP SLAs operations, the ICMP Path Jitter operation is not supported in the RTTMON MIB. Path jitter operations
can only be configured using Cisco commands and statistics can only be returned using the
showipsla commands.
IP SLAs - Path Jitter does not support the IP SLAs History feature (statistics history buckets) because of the large data
volume involved with jitter operations.
The following commands, available in path jitter configuration mode, do not apply to path jitter operations:
historybuckets-kept
historydistributions-of-statistics-kept
historyenhanced
historyfilter
historyhours-of-statistics-kept
historylives-kept
historystatistics-distribution-interval
samples-of-history-kept
lsr-path
tos
threshold
verify-data
Information About IP SLAs ICMP Path Jitter Operations
IP SLAs - ICMP Path Jitter provides hop-by-hop jitter, packet loss, and delay measurement statistics in an IP network. Path
jitter operations function differently than the standard UDP Jitter operation, which provides total one-way data and total
round-trip data.
An ICMP Path Jitter operation can be used a supplement to the standard UDP Jitter operation. For example, results from a
UDP Jitter operation may indicate unexpected delays or high jitter values; an ICMP Path Jitter operation could then be used
to troubleshoot the network path and determine if traffic is bottlenecking in a particular segment along the transmission
path.
The operation first discovers the hop-by-hop IP route from the source to the destination using a traceroute utility, and
then uses ICMP echoes to determine the response times, packet loss and approximate jitter values for each hop along the path.
The jitter values obtained using IP SLAs - ICMP Path Jitter are approximates because ICMP only provides round trip times.
ICMP Path Jitter operations function by tracing the IP path from a source device to a specified destination device, then
sending
N number of Echo probes to each hop along the traced path, with a time interval of
T milliseconds between each Echo probe. The operation as a whole is repeated at a frequency of once every
F seconds. The attributes are user-configurable, as shown here:
Path Jitter Operation Parameter
Default
Configured Using:
Number of echo probes (N )
10 echos
path-jitter command,
num-packets option
Time between Echo probes, in milliseconds (T )
20 ms
path-jitter command,
interval option
Note
The operation’s frequency is different than the operation’s interval.
The frequency of how often the operation is repeated (F )
once every 60 seconds
frequency command
How to Configure the IP SLAs ICMP Path Jitter Operation
Configuring the IP SLAs Responder on a Destination Device
Note
An IP SLAs Responder is not required on either the target device or intermediate devices for path jitter operations. However,
operational efficiency may improve if you enable the IP SLAs Responder.
Before you begin
The networking device to be used as the responder must be a Cisco device and you must have connectivity to that device through
the network.
SUMMARY STEPS
enable
configureterminal
ipslaresponder
exit
DETAILED STEPS
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
ipslaresponder
Example:
Example:
Device(config)# ip sla responder
(Optional) Temporarily enables IP SLAs Responder functionality on a Cisco device in response to control messages from source.
Control is enabled by default.
Step 4
exit
Example:
Device(config)# exit
(Optional) Exits global configuration mode and returns to privileged EXEC mode.
Configuring an ICMP Path Jitter Operation on the Source Device
Perform only one of the following procedures in this section:
Enters IP SLA Path Jitter configuration mode for configuring an ICMP Path Jitter operation.
Step 5
frequencyseconds
Example:
Device(config-ip-sla-pathJitter)# frequency 30
(Optional) Sets the rate at which a specified IP SLAs operation repeats.
Step 6
end
Example:
Device(config-ip-sla-pathJitter)# end
Exits to privileged EXEC mode.
Example
In the following example, the
targetOnly keyword is used to bypass the hop-by-hop measurements. With this version of the command, echo probes will be sent to the
destination only.
(Optional) Sets the protocol data size in the payload of an IP SLAs operation's request packet.
Step 8
tagtext
Example:
Device(config-ip-sla-pathJitter)# tag TelnetPollServer1
(Optional) Creates a user-specified identifier for an IP SLAs operation.
Step 9
timeoutmilliseconds
Example:
Device(config-ip-sla-pathJitter)# timeout 10000
(Optional) Sets the amount of time an IP SLAs operation waits for a response from its request packet.
Step 10
vrfvrf-name
Example:
Device(config-ip-sla-pathJitter)# vrf vpn-A
(Optional) Allows monitoring within Multiprotocol Label Switching (MPLS) Virtual Private Networks (VPNs) using IP SLAs operations.
Step 11
end
Example:
Device(config-ip-sla-pathJitter)# end
Exits to privileged EXEC mode.
Scheduling IP SLAs
Operations
Before you begin
All IP Service Level
Agreements (SLAs) operations to be scheduled must be already configured.
The frequency of all
operations scheduled in a multioperation group must be the same.
The list of one or more
operation ID numbers to be added to a multioperation group must be limited to a
maximum of 125 characters in length, including commas (,).
If the IP Service Level Agreements (SLAs) operation is not running and not generating statistics, add the
verify-data command to the configuration (while configuring in IP SLA configuration mode) to enable data verification. When data verification

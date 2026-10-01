---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-11-configuration-g-ba637895-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-11-configuration-g-ba637895"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-11-configuration-g-ba637895.md
source_anchor: ""
source_lines: [1, 94]
sha256: e1e94a958abf3cee4519fa1601823dbd718862fc3b0ea7f712b91bb3864241e9
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-11-configuration-g-ba637895

Quality of Service Configuration Guide, Cisco IOS XE Dublin 17.11.x (Catalyst 9300 Switches)
Bias-Free Language
Bias-Free Language
The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
Heterogeneous networks include different protocols used by applications, giving rise to the need to prioritize traffic in
order to satisfy time-critical applications while still addressing the needs of less time-dependent applications, such as
file transfer. If your network is designed to support different traffic types that share a single data path between devices
in a network, implementing congestion avoidance mechanisms ensures fair treatment across the various traffic types and avoids
congestion at common network bottlenecks. Congestion avoidance mechanism is achieved through packet dropping.
Random Early Detection (RED) is a commonly used congestion avoidance mechanism in a network.
Tail Drop
Tail drop treats all traffic equally and does not differentiate within a class of service. When the output queue is full and
tail drop is in effect, packets are dropped until the congestion is eliminated and the queue is no longer full.
Weighted Random Early Detection
The RED mechanism takes advantage of the congestion control mechanism of TCP. Packets are randomly dropped prior to periods
of high congestion. Assuming the packet source uses TCP, it decreases its transmission rate until all the packets reach their
destination, indicating that the congestion is cleared. You can use RED as a way to cause TCP to slow down transmission of
packets. TCP not only pauses, but also restarts quickly and adapts its transmission rate to the rate that the network can
support.
WRED is the Cisco implementation of RED. It combines the capabilities of RED algorithm with IP Precedence or Differentiated
Services Code Point (DSCP) or Class of Service (COS) values.
WRED reduces the chances of tail drop by selectively dropping packets when the output interface begins to show signs of congestion.
WRED drops some packets early rather than waiting until the queue is full. Thus it avoids dropping large number of packets
at once and minimizes the chances of TCP global synchronization.
Approximate Fair Drop (AFD) is an Active Queue Management (AQM) algorithm that determines the packet drop probability. The
probability of dropping packets depends upon the arrival rate calculation of a flow at ingress and the current queue length.
AFD based WRED is implemented on wired network ports.
AFD based WRED emulates the preferential dropping behavior of WRED. This preferential dropping behavior is achieved by changing
the weights of AFD sub-classes based on their corresponding WRED drop thresholds. Within a physical queue, traffic with larger
weight incurs less drop probability than that of smaller weight.
Each WRED enabled queue has high and low thresholds.
A sub-class of higher priority has a larger AFD weight.
The sub-classes are sorted in ascending order, based on lowest of WRED minThreshold.
WRED Weight Calculation
AFD weight is calculated using low and high threshold values; AFD is an adjsuted index of the average of WRED high and WRED
low threshold values.
When a packet arrives at an interface, the following events occur:
The drop probability is calculated. The drop probability increases as the AFD weight decreases. That means, if the average
of low and high threshold values is less, the drop probability is more.
WRED consideres the priority of packet flows and the threshold values before deciding to drop the packet. The CoS, DSCP or
IP Precedence values are mapped to the specified thresholds. Once these thresholds are exceeded, packets with the configured
values that are mapped to these thresholds are eligible to be dropped. Other packets with CoS, DSCP or IP Precedence values
assigned to the higher thresholds are en-queued. This process keeps the higher priority flows intact and minimizes the latency
in packet transmission.
If packets are not dropped using WRED, they are tail-dropped.
Limitations for WRED Configuration
Weighted Tail Drop (WTD) is enabled by default on all the queues.
WRED can be enabled / disabled per queue. When WRED is disabled, WTD is adapted on the target queue. Policy-map with WRED
profile is configured only on physical ports as output policy.
WRED is supported only in network port queues and is not supported on internal CPU queues and stack queues.
Each WRED physical queue can support three different threshold pairs. Each pair is for one QoS tag value.
Ensure that you configure bandwidth or shape in the policy-map along with WRED.
Specify all the WRED thresholds only in percentage mode.
Map the WRED threshold pairs by mapping class-map filter with corresponding match filters.
We recommend the class-map with match “any” filter.
WRED for priority traffic is not supported.
WRED and queue limit are not supported for the same policy.
Wired ports support a maximum of eight physical queues, of which you can configure WRED only on four physical queues, each
with three threshold pairs. The remaining queues are configured with WTD. Policies with more than four WRED queues are rejected.
Usage Guidelines for WRED
To configure AFD based WRED feature, specify the policy map and add the class. Use the random-detect command to specify the method (using the dscp-based / cos-based / precedence-based arguments) that you want WRED to use
to calculate the drop probability.
Note
You can modify the policy on the fly. The AFD weights are automatically recalculated.
WRED can be configured for any kind of traffic like IPv4/IPv6, Multicast, and so on. WRED is supported on all 8 queueing classes.
Consider the following points when you are configuring WRED with random-detect command:
With dscp-based argument, WRED uses the DSCP value to calculate drop probability.
With cos-based argument, WRED uses the COS value to calculate drop probability.
By default, WRED uses the IP Precedence value to calculate drop probability. precedence-based argument is the default and it is not displayed in the CLI.
Note
show run policy-map policy-map command does not display “precedence” though precedence is configured with random-detect command.
The dscp-based and precedence-based arguments are mutually exclusive.
Each of the eight physical queues can be configured with different WRED profiles.
Specifies the minimum and maximum thresholds, in percentage.
Step 8
interface interface-name
Example:
device(config)#interface HundredGigE1/0/2
Enters the interface configuration mode.
Step 9
service-policy outputpolicy-map
Example:
device(config-if)#service-policy output pwred
Attaches the policy map to an output interface.
WRED Configuration Example
WRED Support with Hierarchical QoS
Hierarchical QoS allows you to specify QoS behavior at multiple policy levels, which provides a high degree of granularity
in traffic management.
For HQoS, WRED is allowed only on the child policy and not on the parent policy. You can have the shaping configured on the
parent policy and WRED on the child.
The following example configures the parent policy pwred-parent with traffic shaped on the basis of 10 percent of the bandwidth, that applies to its child, pwred-child configured for DSCP-based WRED.
The following example shows how to display WRED AFD Weights, WRED Enq (in Packets and Bytes), WRED Drops (in Packets and Bytes),
Configured DSCP labels against the Threshold pairs:
Note

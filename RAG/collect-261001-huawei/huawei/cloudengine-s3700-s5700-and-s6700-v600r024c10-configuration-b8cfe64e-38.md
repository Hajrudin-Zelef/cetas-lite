---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-38
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [5306, 5426]
sha256: d2d6c87f19a8d62ba57ef73ec3aa96e01b4900adc720125ec60fc4a2bcbf1ca1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Procedure
         Step 1 Check whether packets enter correct queues on the outbound interface.
                    Run the display qos queue statistics interface interface-type interface-number
                    command to check whether packets enter correct queues on the outbound
                    interface.
                    ●   If packets enter incorrect queues, locate the fault. For details, see 8.10.1
                        Packets Enter Incorrect Queues.
                    ●   If packets enter correct queues, go to step 2.
         Step 2 Check whether the priority types trusted by the inbound and outbound interface
                are correct.
                    Run the display this command in the inbound or outbound interface view to
                    check whether the priority type trusted using the trust command is correct.
                    ●   If the trusted priority type is incorrect, run the trust command to specify the
                        correct priority type.
                    ●   If the trusted priority type is correct, go to step 3.
         Step 3 Check whether the priority mappings in the DiffServ domain bound to the
                inbound or outbound interface are correct.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               94
QoS Configuration
QoS Configuration                                                        8 Priority Mapping Configuration


                    Run the display this command in the inbound or outbound interface view to
                    check whether the trust upstream command is configured.
                    Run the display diffserv domain [ name ds-domain-name | brief ] command to
                    check whether the mappings between internal priorities/colors and external
                    priorities are correct.
                    ●   If the priority mappings are incorrect, run the ip-dscp-outbound, or 8021p-
                        outbound command to configure the mappings between internal priorities/
                        colors and external priorities.
                    ●   If the priority mappings are correct, go to step 4.
         Step 4 Check whether there are configurations affecting priority mapping on the inbound
                and outbound interfaces.
                    On an interface:
                    ●   If the qos phb marking dscp enable command is not configured, the system
                        does not perform mapping between PHBs and DSCP values for outgoing
                        packets on the interface.
                    ●   If the qos phb marking 8021p disable command is configured, the system
                        does not perform mapping between PHBs and 802.1p values for outgoing
                        packets on the interface.
                    ●   If the trust upstream none command is configured, the system does not
                        perform priority mapping for outgoing packets on the interface.
                    ●   If packets match a traffic policy that is applied to the inbound or outbound
                        direction (using the traffic-policy command) and contains the action of
                        remark 8021p or remark dscp, the packet priority is the re-marked external
                        priority.
                    Run the display this command in the inbound or outbound interface view to
                    check whether there are any configurations affecting priority mapping. If such
                    configurations are found, delete or modify them.

                    ----End




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             95
QoS Configuration                                          9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                          based Rate Limiting Configuration




          9
          Traffic Policing, Traffic Shaping, and
  Interface-based Rate Limiting Configuration

                    9.1 Overview of Traffic Policing, Traffic Shaping, and Interface-based Rate Limiting
                    9.2 Understanding Traffic Policing, Traffic Shaping, and Interface-based Rate
                    Limiting
                    9.3 Configuration Precautions for Traffic Policing, Traffic Shaping, and Interface-
                    based Rate Limiting
                    9.4 Default Settings for Traffic Policing, Traffic Shaping, and Interface-based Rate
                    Limiting
                    9.5 Configuring Traffic Policing
                    9.6 Configuring Traffic Shaping
                    9.7 Configuring Interface-based Rate Limiting
                    9.8 Configuring a QoS Profile
                    9.9 Maintaining Traffic Policing, Traffic Shaping, and Interface-based Rate Limiting


9.1 Overview of Traffic Policing, Traffic Shaping, and
Interface-based Rate Limiting
Definition
                    Traffic policing, traffic shaping, and interface-based rate limiting monitor and
                    control traffic rates and resource usage.

                    ●   Traffic policing: monitors the rate of traffic entering a network and discards
                        excess traffic to ensure the incoming traffic rate remains within a specified
                        range, conserving network resources.
                    ●   Traffic shaping: adjusts the rate at which traffic is sent to reduce traffic bursts,
                        thereby ensuring a stable transmission rate and preventing congestion on the
                        downstream device.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                  96
QoS Configuration                                         9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                         based Rate Limiting Configuration


                    ●   Interface-based rate limiting: controls the rate at which packets are sent or
                        received on an interface. This mechanism is useful for limiting the rate of all
                        traffic on an interface, regardless of packet types.

Purpose
                    Network congestion may occur when the transmit rate on an upstream device is
                    higher than the receive rate on a downstream device or when the interface rate on
                    a downstream device is lower than the interface rate on an upstream device. If
                    users are allowed to send traffic at an unlimited rate, continuous traffic bursts
                    from many users may result in a congested network. Therefore, user traffic must
                    be rate-limited to ensure services remain stable even in scenarios where network
                    resources are limited. As such, traffic policing, traffic shaping, and interface-based
                    rate limiting can be used to control the traffic rate to improve network resource
                    utilization and service stability.


9.2 Understanding Traffic Policing, Traffic Shaping, and
Interface-based Rate Limiting


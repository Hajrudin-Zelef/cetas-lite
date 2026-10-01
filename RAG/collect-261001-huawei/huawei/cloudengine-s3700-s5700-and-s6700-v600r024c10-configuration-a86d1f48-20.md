---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-20
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [1984, 2147]
sha256: bdbf44feb65c6a7b93616ff5c083b33e5937407a04bf5e5e2960948074994b22
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

3.7.3 Verifying the Configuration
Procedure
                  ●      Run the display cpu-defend policy [ policy-name ] command to check the
                         attack defense policy configuration.
                  ●      Run the display auto-defend attack-source [ slot slot-id | history [ slot slot-
                         id ] | trace-type { source-mac | source-ip | source-portvlan } [ slot slot-id ] ]
                         command to check attack source information.
                  ●      Run the display auto-defend configuration [ cpu-defend policy policy-
                         name | slot slot-id ] command to check the configuration of attack source
                         tracing in the attack defense policy.
                  ●      Run the display auto-defend whitelist slot slot-id command to check the
                         whitelist information configured for attack source tracing.
                  ----End

3.7.4 Example for Configuring Attack Source Tracing
Networking Requirements
                  In Figure 3-4, users on different network segments access the Internet through
                  DeviceA. Because there are a large number of access users, DeviceA often
                  processes a large number of ARP packets, leading to a high CPU usage and hence
                  affecting services.
                  The administrator requires that the device analyze the ARP packets sent to the
                  CPU, identify the packets whose rate exceeds the threshold as attack packets, find

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             30
Security Configuration
Security Configuration                                                         3 Local Attack Defense Configuration


                  out the attack source user or source interface, and send logs and alarms to notify
                  the administrator so that the administrator can take security measures to protect
                  the CPU. Users on Net2 are fixed authorized users, so the administrator needs to
                  ensure that ARP packets of these users can be sent to the CPU.

                  Figure 3-4 Networking diagram of attack source tracing
                          NOTE

                         In this example, interface 1, interface 2, interface 3, and interface 4 represent 10GE1/0/1,
                         10GE1/0/2, 10GE1/0/3, and 10GE1/0/4, respectively.




Procedure
         Step 1 Configure an attack defense policy.
                  # Create an attack defense policy.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] cpu-defend policy test1

                  # Enable attack source tracing.
                  [DeviceA-cpu-defend-policy-test1] auto-defend enable

                  # Set the attack source tracing threshold to 100 pps.
                  [DeviceA-cpu-defend-policy-test1] auto-defend threshold 100

                  # Set the sampling ratio for attack source tracing to 7. That is, one packet is
                  sampled in every seven packets.
                  [DeviceA-cpu-defend-policy-test1] auto-defend attack-packet sample 7

                  # Set the packet type for attack source tracing to ARP packets.
                  [DeviceA-cpu-defend-policy-test1] auto-defend protocol arp

                  # Set the attack source tracing mode to source tracing based on source MAC
                  addresses and source IP addresses.
                  [DeviceA-cpu-defend-policy-test1] auto-defend trace-type source-mac source-ip

                  # Enable the event reporting function for attack source tracing.
                  [DeviceA-cpu-defend-policy-test1] auto-defend alarm enable

                  # Configure a punishment action for attack source tracing: When the device is
                  attacked, it discards packets from the attack source for 360s. Configure the

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            31
Security Configuration
Security Configuration                                                               3 Local Attack Defense Configuration


                  punishment action when you confirm that the device is under attack; otherwise,
                  the device may discard a lot of valid protocol packets, affecting services.
                  [DeviceA-cpu-defend-policy-test1] auto-defend action deny timeout 360
                  [DeviceA-cpu-defend-policy-test1] quit

                  # Configure a whitelist for attack source tracing.
                  [DeviceA] acl number 2001
                  [DeviceA-acl-basic-2001] rule permit source 10.2.2.0 0.0.0.255
                  [DeviceA-acl-basic-2001] quit
                  [DeviceA] cpu-defend policy test1
                  [DeviceA-cpu-defend-policy-test1] auto-defend whitelist 1 acl 2001
                  [DeviceA-cpu-defend-policy-test1] quit

         Step 2 Apply the attack defense policy.
                  [DeviceA] cpu-defend-policy test1

                  ----End


Verifying the Configuration
                  # Display the configuration of attack source tracing.
                  [DeviceA] display auto-defend configuration cpu-defend policy
                  test1
                  -----------------------------------------------------------------------
                   Name : test1
                   Related slot : ***
                   auto-defend                     : enable
                   auto-defend threshold                 : 100 (pps)
                   auto-defend attack-packet sample : 7 (pps)
                   auto-defend alarm                   : enable
                   auto-defend alarm threshold              : 128 (pps)
                   auto-defend action                 : deny timer: 360 (second)
                   auto-defend trace-type                : source-mac source-ip
                   auto-defend protocol                 : arp
                   auto-defend whitelist 1              : acl number 2001
                  -----------------------------------------------------------------------



Configuration Scripts
                  DeviceA
                  #
                  sysname DeviceA
                  #
                  cpu-defend policy test1
                   auto-defend enable
                   auto-defend threshold 100
                   auto-defend attack-packet sample 7
                   auto-defend protocol arp
                   auto-defend trace-type source-mac source-ip
                   auto-defend alarm enable
                   auto-defend action deny timeout 360
                   auto-defend whitelist 1 acl 2001
                  #
                  cpu-defend-policy test1
                  #
                  acl number 2001
                   rule 5 permit source 10.2.2.0 0.0.0.255
                  #
                  return




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          32
Security Configuration
Security Configuration                                             3 Local Attack Defense Configuration




3.8 Configuring Defense Against Malformed Packet
Attacks

3.8.1 Understanding Defense Against Malformed Packet
Attacks
                  A malformed packet attack is a type of attack in which malformed IP packets are
                  sent to a target device, causing the device to encounter an error or even crash
                  when processing such packets and ultimately impacting services running on the
                  device. Defense against malformed packet attacks enables a device to detect and
                  discard malformed packets in real time to protect the device.

                  Malformed packet attacks are classified into the following types.


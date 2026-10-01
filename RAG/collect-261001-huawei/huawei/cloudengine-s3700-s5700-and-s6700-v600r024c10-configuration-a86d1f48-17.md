---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-17
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [1573, 1715]
sha256: 1015923d8bc906baf875324a3c42dea6f79dcc7250ca2971959a0637ef51f8c1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

         Step 6 (Optional) Set the protocol packet sampling ratio for port attack defense.
                  auto-port-defend sample sample-value

         Step 7 (Optional) Set the aging time for port attack defense.
                  auto-port-defend aging-time aging-time

         Step 8 (Optional) Configure a whitelist for port attack defense.
                  auto-port-defend whitelist whitelist-id { acl acl-number | acl ipv6 ipv6-acl-number | interface interface-
                  type interface-number }

                  By default, no whitelist is configured for port attack defense.

                          NOTE

                         Note the following when configuring an ACL-based whitelist for port attack defense:
                         ● Before referencing an ACL in a whitelist, create the ACL and configure rules.
                         ● The ACL referenced can be a basic ACL, advanced ACL, Layer 2 ACL, basic ACL6, or
                           advanced ACL6.
                         ● All packets matching an ACL referenced by a whitelist are considered legitimate and are
                           not processed based on port attack defense, regardless of whether the ACL rule is
                           permit or deny.
                         ● If an ACL has no rule, the whitelist that references the ACL does not take effect.
                         ● If an ACL rule is defined by a protocol, ensure that the port attack defense function
                           supports this protocol.
                         ● Insufficient ACL resources may lead to an ineffective whitelist.

         Step 9 (Optional) Enable the function of reporting port attack defense events.
                  undo auto-port-defend alarm disable

                  By default, the function of reporting port attack defense events is enabled.

        Step 10 Return to the system view.
                  quit

        Step 11 Apply the attack defense policy.
                  ●      Configure attack defense policies in batches.
                         cpu-defend-policy policy-name batch slot { start-slot [ to end-slot ] } &<1-12>

                  ●      Configure an attack defense policy separately.
                         cpu-defend-policy policy-name [ slot slot-id | mcu ]


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                               23
Security Configuration
Security Configuration                                                      3 Local Attack Defense Configuration


                  After an attack defense policy is created, you must apply the policy in the system
                  view. Otherwise, the policy does not take effect.

                  ----End

Example
                  Enable port attack defense in the attack defense policy named test. Use the
                  default rate threshold, set the sampling ratio to 7 and aging time to 200s, and add
                  10GE 1/0/1 to a whitelist so that port attack defense will not be applied to this
                  interface.
                  <HUAWEI> system-view
                  [HUAWEI] cpu-defend policy test
                  [HUAWEI-cpu-defend-policy-test] auto-port-defend enable
                  [HUAWEI-cpu-defend-policy-test] auto-port-defend sample 7
                  [HUAWEI-cpu-defend-policy-test] auto-port-defend aging-time 200
                  [HUAWEI-cpu-defend-policy-test] auto-port-defend whitelist 1 interface 10ge 1/0/1
                  [HUAWEI-cpu-defend-policy-test] quit
                  [HUAWEI] cpu-defend-policy test


3.5.3 Verifying the Configuration
Procedure
                  ●      Run the display cpu-defend policy [ policy-name ] command to check the
                         attack defense policy configuration.
                  ●      Run the display cpu-defend auto-port-defend configuration [ slot slot-id ]
                         command to check the configuration of port attack defense.
                  ●      Run the display cpu-defend auto-port-defend attack-source [ slot slot-id ]
                         command to check source tracing information for port attack defense.
                  ●      Run the display cpu-defend auto-port-defend whitelist slot slot-id
                         command to check the whitelist information configured for port attack
                         defense.
                  ●      Run the display cpu-defend auto-port-defend statistics [ slot slot-id ]
                         command to check packet statistics for port attack defense.
                  ----End


3.6 Configuring User-Level Rate Limiting
                          NOTE

                         Only the S6780-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2, S6750-H, S5755-H, and
                         S5732-H-V2 support user-level rate limiting.


3.6.1 Understanding User-Level Rate Limiting
                  User-side hosts are vulnerable to virus attacks, and infected hosts can flood
                  devices with protocol packets. This overloads device CPUs and hinders their
                  performance, ultimately impacting services. To prevent this, the network
                  administrator can configure user-level rate limiting, which identifies users based
                  on MAC addresses and rate-limits only the user who initiates an attack, ensuring

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  24
Security Configuration
Security Configuration                                                         3 Local Attack Defense Configuration


                  other normal users are unaffected. User-level rate limiting (based on users) is
                  more precise and has less of an impact on normal users than CPCAR (based on
                  devices) and port attack defense (based on ports).
                  The process of user-level rate limiting is as follows:
                  1.     When receiving users' protocol packets, the device performs a hash
                         calculation on the source MAC addresses and places packets from different
                         source MAC addresses into different buckets.
                  2.     When the number of packets placed in a bucket within the unit time exceeds
                         the rate limit, the bucket discards the excess packets. The device counts the
                         number of discarded packets every 10 minutes, and reports a packet discard
                         log for the bucket when the number of discarded packets exceeds 2000 within
                         10 minutes. If the number of discarded packets exceeds 2000 in multiple
                         buckets, the device records the packet discard logs for the top 10 buckets.

3.6.2 Configuring User-Level Rate Limiting

Context
                  User-level rate limiting can be configured to accurately limit the rate based on
                  user MAC addresses, reducing the impact on common users.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable user-level rate limiting globally.
                  cpu-defend host-car enable

         Step 3 Set the user-level rate limit.
                  cpu-defend host-car [ mac-address mac-address | car-id car-id ] pps pps-value

         Step 4 Specify the packet types to which user-level rate limiting is applied.
                  cpu-defend host-car { { 8021x | arp | dhcp-request | dhcpv6-request | nd | mac-miss } * | all }

         Step 5 Enter the interface view.
                  interface interface-type interface-number

         Step 6 Enable user-level rate limiting on an interface.
                  undo host-car disable


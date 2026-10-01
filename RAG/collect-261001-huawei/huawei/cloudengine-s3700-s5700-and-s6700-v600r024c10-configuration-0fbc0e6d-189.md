---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-189
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [27430, 27560]
sha256: 7b479fd614541517f883df40a2a3f3291a69c4e08f708b99e06d944a6cc1bafc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 For an Ethernet interface, you also need to run the undo portswitch command to switch the
                 interface to Layer 3 mode.
                 Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S, S6750E-S,
                 S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from Layer 2 mode to Layer 3
                 mode using the undo portswitch command.Determine whether to run this command based on
                 the current interface mode.

                 To implement link protection, specify the link parameter. Otherwise, only node
                 protection is implemented.
                 After Auto FRR is enabled globally, mpls te auto-frr default is automatically
                 configured for all MPLS TE-enabled interfaces on the device. To disable Auto FRR
                 on some interfaces, run the mpls te auto-frr block command on these interfaces.
                 After the mpls te auto-frr block command is run on an interface, the interface
                 does not have the Auto FRR capability, regardless of whether Auto FRR is enabled
                 or re-enabled globally.
                 To enable an automatic bypass tunnel to dynamically select node protection or
                 link protection based on network conditions, specify self-adapting.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               458
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


                         NOTE

                        ● If the mpls te auto-frr default command is run, the Auto FRR capability status of an
                          interface is the same as the global Auto FRR capability status.
                        ● After node protection is enabled, if a bypass tunnel for node protection fails to be
                          created because the topology does not meet requirements, the penultimate hop of the
                          primary tunnel attempts to create link protection, and other nodes do not degrade to
                          create link protection.
                        ● If self-adapting is not specified and node protection is enabled, the penultimate hop of
                          the primary tunnel attempts to create link protection when a bypass tunnel fails to be
                          created because the topology does not meet requirements. Other nodes do not degrade
                          to create link protection.

         Step 6 (Optional) Configure the device to consider the Node protection flag sent by the
                tunnel ingress during FRR protection.
                 quit
                 mpls
                 mpls te plr apply node-protection-flag

                 After this command is run, if node protection is desired by the ingress, the PLR
                 prioritizes node protection. Otherwise, the PLR selects link protection.

                         NOTE

                        ● If Huawei devices interwork, you are advised to run the mpls te frr { node-protection |
                          no-node-protection } command on the ingress and the mpls te plr apply node-
                          protection-flag command on the PLR of a tunnel.
                        ● After the mpls te auto-frr command is run in the system view, the mpls te auto-frr
                          default command is automatically run for all MPLS TE-enabled interfaces on the device.
                          If the mpls te auto-frr command is not run in the view of a specific interface, node
                          protection is used by an automatic bypass tunnel on the interface by default.
                           –    However, if the mpls te plr apply node-protection-flag command is run, only the
                                received Node protection flag is considered when an automatic bypass tunnel is
                                established.
                           –    If the non-default mpls te auto-frr configuration is used on an interface, the TE
                                FRR protection mode is selected for an automatic bypass tunnel based on the
                                interface configuration, regardless of whether the mpls te plr apply node-
                                protection-flag command is run.

                 ----End

4.27.3 Enabling TE FRR and Configuring Attributes for an
Automatic Bypass Tunnel

Context
                 After TE FRR is enabled for a primary tunnel, the system can automatically
                 establish a bypass tunnel.

                 Perform the following configuration on the ingress of a primary MPLS TE tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      459
MPLS Configuration
MPLS Configuration                                                                           4 MPLS TE Configuration


         Step 2 Enter the tunnel interface view of a primary MPLS TE tunnel.
                 interface tunnel tunnel-number

         Step 3 Enable MPLS TE FRR.
                 mpls te fast-reroute [ bandwidth ]

                 If bandwidth protection is required, the bandwidth parameter must be configured.
         Step 4 (Optional) Configure bandwidth and priority constraints for the automatic bypass
                tunnel.
                 mpls te bypass-attributes [ bandwidth bandwidth ][ priority setup-priority [ hold-priority ]]

                         NOTE

                        ● The bandwidth attribute can be configured for a bypass tunnel only after mpls te fast-
                          reroute bandwidth is configured for the primary tunnel.
                        ● The bandwidth of an automatic bypass tunnel cannot be greater than that of the
                          primary tunnel.
                        ● If no attribute is configured for an automatic bypass tunnel, the default bandwidth of
                          the automatic bypass tunnel is the same as that of the primary tunnel.
                        ● The setup priority of a bypass tunnel cannot be higher than its holding priority, and
                          neither of them can be higher than the corresponding priority of the primary tunnel.
                        ● After TE FRR is disabled, the attributes of a bypass tunnel are automatically deleted.

         Step 5 (Optional) Configure the affinity and link administrative group attributes for the
                automatic bypass tunnel.
                 1.     Configure the affinity attribute for the bypass tunnel in the tunnel interface
                        view of the primary tunnel.
                        mpls te bypass-attributes affinity property properties [ mask mask-value ]

                 2.     Configure the affinity and link administrative group attributes for the bypass
                        tunnel in the interface view of the link that the bypass tunnel passes through.
                        quit
                        interface interface-type interface-number
                        mpls te link administrative group value
                        mpls te auto-frr attributes affinity property properties [ mask mask-value ]

                         NOTE

                        If an automatic bypass tunnel that satisfies the specified affinity cannot be established, the
                        system will select a manual bypass tunnel that satisfies the specified affinity for the primary
                        tunnel.


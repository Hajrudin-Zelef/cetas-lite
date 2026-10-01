---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-179
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [25984, 26129]
sha256: 60bc73f3c977127bee6a4406db327cad06232c130c2488a9c23260d2ea1dcb0d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        By default, when a Huawei device functions as the ingress of a tunnel in facility backup
                        mode, the device always sets the Node protection flag to 1 in sent messages, indicating
                        that node protection is desired.
                        If link protection is required, run the mpls te frr no-node-protection command on the
                        ingress. If a PLR is a Huawei device, you also need to run the mpls te plr apply node-
                        protection-flag command on the PLR.
                        If node protection is required, you do not need to run this command on the ingress or run
                        the mpls te plr apply node-protection-flag on the PLR.
                        After the mpls te frr { node-protection | no-node-protection } command is run in the
                        MPLS view, the configuration takes effect on all P2P RSVP-TE tunnel interfaces for which
                        the mpls te frr { node-protection | no-node-protection } command is not run but TE FRR
                        is enabled.

                 ----End

4.26.3 Configuring a Bypass Tunnel

Context
                 Before configuring a bypass tunnel, plan the links or nodes to be protected and
                 ensure that the bypass tunnel does not pass through the protected links or nodes.
                 Otherwise, the bypass tunnel cannot protect them.

                         NOTE

                        TE FRR does not support multi-point failures. That is, if FRR switching occurs, data is
                        switched from the primary tunnel to the bypass tunnel. During data forwarding through the
                        bypass tunnel, the bypass tunnel must remain up. If the bypass tunnel goes down during
                        this period, the protected data cannot be forwarded through MPLS. As a result, traffic may
                        be interrupted and FRR fails. Even if the bypass tunnel goes up again, traffic cannot be
                        forwarded through the bypass tunnel. Traffic can be forwarded through the primary tunnel
                        only after the primary tunnel recovers or is re-established.

                 Perform the following steps on the PLR of a bypass tunnel:


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the view of the bypass tunnel interface.
                 interface tunnel tunnel-number

         Step 3 Use either of the following methods to configure an IP address for the bypass
                tunnel interface:
                 ●      Configure an IP address for the tunnel interface.
                        ip address ip-address { mask | mask-length } [ sub ]

                        Configure a primary IP address before you can add a secondary IP address for
                        a tunnel interface.
                 ●      Configure the tunnel interface to borrow the IP address of another interface.
                        ip address unnumbered interface interface-type interface-number


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        434
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


                             NOTE

                            A TE tunnel can be established on a tunnel interface without an IP address. However,
                            an IP address must be configured for the tunnel interface before it can forward traffic
                            over the TE tunnel. An MPLS TE tunnel is unidirectional and does not involve peer
                            address configuration. Therefore, you are advised to specify the ingress LSR ID as the
                            IP address of the tunnel interface, instead of configuring a unique IP address for the
                            interface.

         Step 4 Configure MPLS TE as the tunneling protocol.
                 tunnel-protocol mpls te

         Step 5 Configure the LSR ID of an MP as the destination address of the bypass tunnel.
                 destination ip-address

         Step 6 Configure a bypass tunnel ID.
                 mpls te tunnel-id tunnel-id

         Step 7 Configure the bypass tunnel.
                 mpls te bypass-tunnel

                 After a bypass tunnel is configured, the system automatically enables the route
                 recording function to record detailed path information about the tunnel.

                         NOTE

                        The mpls te fast-reroute and mpls te bypass-tunnel commands cannot be configured on
                        the same tunnel interface.

         Step 8 Specify the interface to be protected by the bypass tunnel.
                 mpls te protected-interface interface-type interface-number

         Step 9 (Optional) Configure an explicit path for the bypass tunnel.
                 mpls te path explicit-path path-name [ secondary ]

                 Before configuring an explicit path for a bypass tunnel, run the explicit-path
                 command to create the explicit path. Note that the physical link of the bypass
                 tunnel cannot overlap that of the primary tunnel.
        Step 10 (Optional) Configure the bandwidth for the bypass tunnel.
                 mpls te bandwidth ct0 ct0-value

        Step 11 (Optional) Configure a delay for FRR switching.
                 quit
                 mpls
                 mpls te frr-switch-delay value

                 On a network where both hot standby and FRR are configured and the PLR of FRR
                 is the ingress of a tunnel, when the primary LSP fails, traffic may be switched to
                 the FRR path and then to the hot-standby path. To prevent secondary switching,
                 you can configure a delay for FRR switching so that FRR entry delivery is delayed
                 and traffic is directly switched to the hot-standby path.
        Step 12 (Optional) Set the interval at which the TE FRR binding relationship is refreshed.
                 quit
                 mpls
                 mpls te timer fast-reroute [ weight ]

                 After TE FRR is configured, the PLR periodically refreshes the TE FRR binding
                 relationship. By default, the system searches for an optimal bypass tunnel among

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                       435
MPLS Configuration
MPLS Configuration                                                                     4 MPLS TE Configuration


                 all manually configured bypass tunnels for each primary tunnel at an interval of
                 1s and binds the bypass tunnel to the primary tunnel.

        Step 13 (Optional) Configure the device to consider the Node protection flag sent by the
                tunnel ingress during FRR protection.
                 quit
                 mpls
                 mpls te plr apply node-protection-flag

                 After this command is run, if node protection is desired by the ingress, the PLR
                 prioritizes node protection. Otherwise, the PLR selects link protection.

                         NOTE

                        If Huawei devices interwork, you are advised to run the mpls te frr { node-protection | no-
                        node-protection } command on the ingress of a tunnel and the mpls te plr apply node-
                        protection-flag command on the PLR of the tunnel.

                 ----End

4.26.4 Verifying the Configuration


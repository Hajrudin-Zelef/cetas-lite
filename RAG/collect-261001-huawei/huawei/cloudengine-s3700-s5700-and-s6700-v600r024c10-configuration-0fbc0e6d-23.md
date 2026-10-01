---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-23
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [2612, 2772]
sha256: 3e320c65fc9088b73070d5f16c3af881528cb46ecd2fbd9f3bbfc1261d5d3baf
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

3.6.3 Configuring LDP Sessions
Context
                 You can configure the following types of MPLS LDP sessions:
                 ●      Local LDP sessions (interface-based)
                        Generally, local LDP sessions need to be configured when MPLS LDP services
                        are deployed.
                 ●      Remote LDP sessions
                        Remote LDP sessions need to be configured if non-neighboring LSRs need to
                        communicate.
                 Both a local LDP session and a remote LDP session can be configured between
                 two LSRs. If this is the case, ensure that the configurations for the local and
                 remote LDP sessions are consistent.

Procedure
                 ●      Configure a local LDP session (interface-based).
                        Perform the following configuration on two neighboring LSRs.
                        a.       Enter the system view.
                                 system-view
                        b.       Enter the view of an interface on which an LDP session is to be
                                 established.
                                 interface interface-type interface-number
                        c.       Switch the interface working mode to Layer 3.
                                 undo portswitch

                                 Determine whether to perform this step based on the current interface
                                 working mode.

                                         NOTE

                                        If multiple interfaces need to be switched to the Layer 3 mode, run the undo
                                        portswitch batch interface-type { interface-number1 [ to interface-number2 ] }
                                        &<1-10> command in the system view to switch these interfaces to the Layer 3
                                        mode in a batch.
                        d.       Enable MPLS on the interface.
                                 mpls

                                 By default, MPLS is disabled on an interface.
                        e.       Enable MPLS LDP on the interface.
                                 mpls ldp


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                      45
MPLS Configuration
MPLS Configuration                                                              3 MPLS LDP Configuration


                             By default, MPLS LDP is disabled on an interface.
                 ●      Configure a remote LDP session.

                        Perform the following configuration on two non-neighboring LDP LSRs that
                        need to communicate.

                        a.   Enter the system view.
                             system-view

                        b.   Create a remote MPLS LDP peer and enter the remote MPLS-LDP peer
                             view.
                             mpls ldp remote-peer remote-peer-name

                        c.   Configure a description for the remote MPLS LDP peer.
                             description description-value

                        d.   Configure an IP address for the remote MPLS LDP peer.
                             remote-ip ip-address

                             By default, no IP address is configured for a remote LDP peer.

                             The IP address of a remote peer must be its LSR ID. If the LSR ID of the
                             LDP instance and the LSR ID of a node are different, use the LSR ID of
                             the LDP instance.



                                 NOTICE

                             ● Modifying or deleting the IP address of a remote peer leads to the
                               deletion of a remote LDP session, interrupting MPLS services.
                             ● After an IP address is specified for a remote peer using the remote-ip
                               ip-address command, do not use the IP address as a local interface's IP
                               address. Otherwise, the remote session is torn down, and MPLS
                               services are interrupted.

                        e.   (Optional) Disable label distribution to remote peers as required.

                             ▪    Disable label distribution to a specified remote LDP peer.
                                  remote-ip ip-address pwe3

                                  To enable a device to distribute labels to a specified remote MPLS
                                  LDP peer, run the clear remote-ip pwe3 command.

                             ▪    Disable label distribution to any remote LDP peers.
                                  quit
                                  mpls ldp
                                  remote-peer pwe3

                                  By default, the device is permitted to distribute labels to all remote
                                  LDP peers.

                 ----End


Result
                 Run the display mpls ldp [ all | all verbose ] command to check LDP information.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 46
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration


3.6.4 Example for Configuring Local LDP Sessions

Networking Requirements
                 As shown in Figure 3-7, LSRA, LSRB, and LSRC function as core or edge devices on
                 the backbone network. Configure local LDP sessions for MPLS LDP services. The
                 LSRs can then exchange labels to establish LDP LSPs.


                 Figure 3-7 Configuring local LDP sessions
                         NOTE

                        Interfaces 1 and 2 in this example represent VLANIF100 and VLANIF200, respectively.




Precautions
                 During the configuration, note the following:

                 ●      LSR IDs must be set before other MPLS commands are run.
                 ●      LSR IDs can only be manually configured, and do not have default values.
                 ●      Using the IP address of a reachable loopback interface on an LSR as the LSR
                        ID is recommended.


Configuration Roadmap
                 The configuration roadmap is as follows:

                 1.     Assign an IP address to each interface and configure OSPF to advertise the
                        route to the network segment to which each interface is connected and the
                        host route to each LSR ID.
                 2.     Enable MPLS and MPLS LDP globally on each LSR.
                 3.     Enable MPLS on the interfaces of each LSR.
                 4.     Enable MPLS LDP on the interfaces of both ends of each local LDP session.


Data Preparation
                 To complete the configuration, prepare the following data:

                 ●      IP address of each interface on each LSR (as shown in Figure 3-7), OSPF
                        process ID, and area ID
                 ●      LSR ID of each node


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    47
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration



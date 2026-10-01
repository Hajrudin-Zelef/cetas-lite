---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-217
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [31586, 31714]
sha256: 6a82aea65e89a83c5d956b9e57b882148b4c7e018b88a7a2e15bcad78376da51
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      CSPF fails to calculate a path.
                 ●      RSVP is not configured on a device along the TE tunnel.
                 ●      The physical state of an interface is not up.
                 ●      Devices along the TE tunnel fail to send or receive packets.

Procedure
         Step 1 Check whether the tunnel configuration is complete.

                 Run the display current-configuration command on a tunnel node to check
                 whether the following commands are executed on the tunnel interface. If any
                 command is missing, the configuration is incomplete.
                 interface Tunnel1
                  tunnel-protocol mpls te
                  destination 10.4.4.9
                  mpls te tunnel-id 1

                 ●      If the tunnel configuration is incomplete, add missing command
                        configurations.
                 ●      If the tunnel configuration is complete, go to Step 2.

         Step 2 Check whether a path is calculated using CSPF.

                 Method 1: Run the display mpls te cspf destination ip-address command on the
                 TE tunnel node to check whether CSPF-based path calculation is successful. To
                 check information about the paths that satisfy certain constraints, specify the
                 corresponding parameters in the command. For example, specify the explicit-path
                 path-name parameter to check information about the path that has a specified
                 path name.

                 Method 2: Run the display mpls te tunnel-interface last-error [ tunnel-name ]
                 command on the TE ingress to check the latest errors that occur on the tunnel
                 interface of the ingress. Determine whether CSPF-based path calculation is
                 successful based on the command output.

                 ●      If CSPF-based path calculation is unsuccessful, run the display ip routing-
                        table command to check whether there is a route to the destination address
                        of the tunnel.
                        –   If no such route exists, refer to the ping troubleshooting method to rectify
                            faults.
                        –   If such a route exists, go to Step 3.
                 ●      If CSPF-based path calculation is successful, go to Step 3.

         Step 3 Check whether RSVP is configured for the devices along the tunnel.

                 Determine the interfaces along the tunnel by running the display mpls te tunnel
                 path command or the interfaces along the LSP by checking the network topology.
                 Run the display this command on each of the interfaces to check whether MPLS,
                 MPLS TE, and RSVP-TE are enabled on the interfaces residing on the path to the
                 destination.

                 ●      If they are not enabled, run the mpls, mpls te, and mpls rsvp-te commands
                        in the interface view.
                 ●      If they are enabled, go to Step 4.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                          522
MPLS Configuration
MPLS Configuration                                                                 4 MPLS TE Configuration


         Step 4 Check whether the physical state of an interface is up.
                 Run the display interface interface-type interface-number command to check the
                 physical state of the interface.
                 ●      If the interface state is not up, restart the interface. Specifically, run the
                        shutdown and then undo shutdown commands in the interface view.
                        Alternatively, run the restart command in the interface view.
                 ●      If the interface state is up, go to Step 5.
         Step 5 Check whether packets are properly sent and received between devices along the
                tunnel.
                 Run the display mpls te tunnel-interface command on each TE tunnel node to
                 check the Ingress LSR ID, Primary LSP ID, and Session ID fields in the command
                 output. Assume that LSRA, LSRB, and LSRC are located on the tunnel path, as
                 identified in Step 3.
                 Perform the following operations to check whether RSVP Path messages and RSVP
                 Resv messages are correctly sent and received.
                 ●      Check whether RSVP Path messages are correctly sent and received (Path
                        message sending direction: LSRA -> LSRB -> LSRC).
                        Run the display mpls rsvp-te psb-content command on each node along the
                        path.
                        –   If the command displays information properly on each node, RSVP Path
                            messages are correctly sent and received between these nodes. In this
                            case, refer to the ping troubleshooting method to rectify faults.
                        –   If no information is displayed on a node, RSVP Path messages fail to be
                            sent or received between the local node and its upstream node.
                 ●      Check whether RSVP Resv messages are correctly sent and received (Resv
                        message sending direction: LSRC -> LSRB -> LSRA).
                        Run the display mpls rsvp-te rsb-content command on each node along the
                        path.
                        –   If the command displays information properly on each node, RSVP Resv
                            messages are correctly sent and received between these nodes.
                        –   If no information is displayed on a node, RSVP Resv messages fail to be
                            sent or received between the local node and its upstream node. In this
                            case, go to Step 6.
         Step 6 If the fault persists, contact technical support and provide the following
                information:
                 ●      Results of the preceding procedure
                 ●      Configuration file, log, and alarm information of the device

                 ----End

4.33.2 An MPLS TE Tunnel's State Changes from Up to Down
Suddenly
Fault Symptom
                 An MPLS TE tunnel in the up state suddenly goes down.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                523
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


Possible Causes
                 ●      The configuration of the TE tunnel is deleted using a command.
                 ●      A physical interface on the TE tunnel goes down.
                 ●      The transmission of RSVP-TE messages times out.


Procedure
         Step 1 Check whether the tunnel configuration is deleted using a command.

                 Run the display this command on the ingress of the tunnel to check whether the
                 following commands are configured:

                 ●      shutdown
                 ●      undo interface Tunnel tunnel-number
                 ●      If one of the preceding commands is configured, restore the deleted
                        configurations.
                 ●      If none of the preceding commands is configured, go to Step 2.


---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-218
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [31715, 31792]
sha256: 8c442563398d2d7d9f8219942558da94d7105ff41ea81b2b616719e5d545f2a6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

         Step 2 Check whether any physical interface along the TE tunnel is down.
                 1.     Run the display mpls te cspf destination ip-address explicit-path pri-path-
                        name command on a tunnel node. The IP addresses displayed in the
                        command output indicate the nodes that the tunnel traverses. Record the
                        nodes that the tunnel traverses.
                 2.     Check alarms on each node along the tunnel. Check whether the active/
                        linkDown alarm is generated on the NMS. Record the alarm time, such as T1.
                        Run the display interface interface-type interface-number command to check
                        the value of the Last line protocol up time field. Record the time displayed
                        in the field, such as T2.
                 ●      If T1 is larger than T2, the TE tunnel went down because the physical
                        interface on the TE tunnel is faulty. Refer to the methods for troubleshooting
                        physical interface faults.
                 ●      If T1 is smaller than T2, go to Step 3.

         Step 3 Check whether the transmission of RSVP-TE messages times out.

                 Check the P2PTE/2/mplsTunnelDown log on the ingress of the tunnel to determine
                 the time when the tunnel went down, such as T1. Then, run the display mpls
                 rsvp-te last-error command on each node along the tunnel determined in Step 2
                 to check last-error information. Check whether the following information is
                 generated within 10 minutes before T1:
                 ●      PATH TIME OUT
                 ●      RESV TIME OUT

                 ●      If the information exists, refer to the ping troubleshooting method to rectify
                        faults.
                 ●      If the information does not exist, go to Step 4.

         Step 4 Contact technical support.
                 ●      Results of the preceding procedure

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            524
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


                 ●      Configuration file, log, and alarm information of the device

                 ----End

4.33.3 A Loop Occurs on an MPLS TE Tunnel
Fault Symptom
                 A loop occurs on an MPLS TE tunnel.

Possible Causes
                 ●      Path messages are looped.
                 ●      Resv messages are looped.

Procedure
         Step 1 Check whether Path messages are looped.
                 Run the display mpls rsvp-te last-error command on an RSVP-TE-enabled node.
                 If path loop information is displayed, Path messages are looped.
                 Run the display mpls te cspf destination ip-address explicit-path pri-path-name
                 command on a tunnel node. The IP addresses displayed in the command output
                 and the LSR IDs indicate the nodes that the tunnel traverses.
                 On the node where a loop occurs, run the tracert ip-address command. Set the ip-
                 address to each hop of the tunnel to check whether an IP address conflict occurs.
                 ●      If an IP address conflict occurs, delete or change the IP address.
                 ●      If no conflicting IP address exists, go to Step 2.
         Step 2 Check whether Resv messages are looped.
                 Run the display mpls rsvp-te last-error command on an RSVP-TE-enabled node.
                 If Resv loop information is displayed, Resv messages are looped.
                 The troubleshooting operations are the same as that in Step 1.
                 ●      If an IP address conflict occurs, delete or change the IP address.
                 ●      If no conflicting IP address exists, go to Step 3.
         Step 3 Collect the following information and contact technical support personnel:
                 ●      Results of the preceding procedure
                 ●      Configuration file, log, and alarm information of the device

                 ----End




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          525


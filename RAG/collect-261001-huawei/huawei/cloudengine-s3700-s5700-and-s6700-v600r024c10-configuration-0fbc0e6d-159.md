---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-159
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [23009, 23144]
sha256: d637fb56e67158769ae40b94dfd698b9807ca742051d00c98f7327bd41c22bad
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      Best-effort path: A temporary CR-LSP is created when both the primary and
                        backup CR-LSPs fail. Service traffic is switched to the best-effort path.
                        On the network shown in Figure 4-34, the primary CR-LSP is set up over the
                        path PE1 -> P1 -> P2 -> PE2, and the backup CR-LSP is set up over the path
                        PE1 -> P3 -> PE2. When both CR-LSPs fail, PE1 sets up a best-effort path PE1
                        -> P4 -> PE2 to take over traffic.

                        Figure 4-34 Best-effort path




                             NOTE

                           A best-effort path has no bandwidth reserved for traffic, but has an affinity and a hop
                           limit configured to control the nodes it passes through.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                     383
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration


Manual Switchback and Switchback Policy of CR-LSP Hot Standby
                 In CR-LSP hot standby scenarios, traffic can be switched in automatic or manual
                 mode:
                 ●      Automatic switching: Traffic is switched to a hot-standby CR-LSP when the
                        primary CR-LSP goes down. When the primary CR-LSP goes up, traffic is
                        automatically switched back to the primary CR-LSP. You can determine
                        whether to switch traffic back to the primary CR-LSP and set a switchback
                        delay.
                 ●      Manual switching: You can manually trigger traffic switching. Forcibly switch
                        traffic from the primary CR-LSP to a hot-standby CR-LSP before some devices
                        on a primary CR-LSP are upgraded or primary CR-LSP parameters are
                        adjusted. After the required operations are complete, manually switch traffic
                        back to the primary CR-LSP.

Path Overlapping
                 The path overlapping function can be configured for hot-standby CR-LSPs. It
                 means that the paths of a hot-standby CR-LSP and primary CR-LSP can overlap
                 while being disjointed as much as possible. This ensures that the hot-standby CR-
                 LSP provides maximum protection to the primary CR-LSP.

Comparison with Other Features
                 1. Differences between CR-LSP backup and TE FRR
                 ●      CR-LSP backup provides end-to-end path protection for an entire CR-LSP.
                 ●      FRR is a partial protection mechanism used to protect a link or node on a CR-
                        LSP. It is a temporary protection measure that can respond to faults rapidly,
                        but it has strict requirements on the switching time.
                        For details about FRR, see TE FRR.

                 2. Synchronization between CR-LSP hot standby and TE FRR

                 TE FRR is a temporary local protection mechanism used when the ingress does not
                 detect a fault. It is a supplement to CR-LSP hot standby. Once the ingress of the
                 primary CR-LSP detects the fault, traffic is switched to the hot-standby CR-LSP.

                 After synchronization (FRR in use) is enabled, if the primary CR-LSP fails but the
                 backup CR-LSP is up, traffic is switched to the TE FRR bypass tunnel and then
                 immediately switched to the backup CR-LSP. At the same time, the system
                 attempts to restore the primary CR-LSP. If the backup CR-LSP is down, traffic is still
                 transmitted through the bypass CR-LSP.

                 3. Synchronization between CR-LSP ordinary backup and TE FRR
                 ●      Synchronization is not enabled.
                        If a protected link or node fails, a PLR switches traffic to a bypass tunnel. Only
                        after both the primary and bypass CR-LSPs fail, the ingress of the primary CR-
                        LSP attempts to establish an ordinary backup CR-LSP and switches traffic to
                        this CR-LSP.
                 ●      Synchronization is enabled (FRR in use).
                        If a protected link or node fails, a PLR switches traffic to a bypass tunnel.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               384
MPLS Configuration
MPLS Configuration                                                                           4 MPLS TE Configuration


                        –    If the PLR is the ingress of the primary CR-LSP, it attempts to set up an
                             ordinary backup CR-LSP. After the ordinary backup CR-LSP is set up,
                             traffic is switched to the ordinary backup CR-LSP.
                        –    If the PLR is a transit node of the primary CR-LSP, it sends fault
                             information to the ingress of the primary CR-LSP through RSVP-TE
                             signaling, triggering the ingress to set up an ordinary backup CR-LSP.
                             When the ordinary backup CR-LSP is set up successfully, the ingress
                             switches traffic to the ordinary backup CR-LSP.
                        If the ordinary backup CR-LSP fails to be set up, traffic keeps traveling
                        through the bypass CR-LSP.

4.23.2 Configuring CR-LSP Hot Standby
Prerequisites
                 Before configuring CR-LSP hot standby, complete the following task:
                 ●      Configure a dynamic MPLS TE tunnel.
                 ●      Enable MPLS, MPLS TE, and RSVP-TE globally and on interfaces on each node
                        of the backup CR-LSP. For details, see 4.7.1 Enabling MPLS TE and RSVP-TE.

Context
                 CR-LSP backup provides end-to-end protection for traffic on a CR-LSP. If a primary
                 CR-LSP fails, traffic is rapidly switched to a backup CR-LSP, ensuring uninterrupted
                 traffic transmission.
                 Perform the following configuration on the ingress of an MPLS TE tunnel.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS TE tunnel interface view.
                 interface tunnel tunnel-number

         Step 3 Configure the hot standby CR-LSP establishment mode.
                 mpls te backup hot-standby { mode { revertive [ wtr interval ] | non-revertive } | wtr interval }

                 Select the following parameters as needed to enable sub-functions:
                 ●      mode revertive [ wtr interval ]: enables a device to switch traffic back to a
                        primary CR-LSP.
                 ●      mode non-revertive: disables a device from switching traffic back to a
                        primary CR-LSP.
                 ●      wtr interval: sets the time before a traffic switchback can be performed.

                         NOTE

                        The bypass and backup tunnels cannot be configured on the same tunnel interface.
                        Specifically, the mpls te bypass-tunnel and mpls te backup commands cannot be
                        configured on the same tunnel interface. Also, the mpls te protected-interface and mpls
                        te backup commands cannot be configured on the same tunnel interface.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                           385
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration



---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-295
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [43617, 43735]
sha256: 9eff9f2bff463019d3b5497504d59d29c0935b9cd55ef8837f5da265f80fdd75
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          699
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration


                             NOTE

                            The forwarding priority takes effect only when PW redundancy operates in master/
                            slave mode. In master/slave mode, PE1 instructs PE2 and PE5 to adjust the PW
                            forwarding status. In independent mode, the forwarding status notified by the remote
                            PEs determines the PW forwarding status on the local PE.
                    ●   Active/Standby: defines the PW forwarding status and cannot be configured.
                        Only PWs in the active state are used to forward traffic, and PWs in the
                        standby state can only be configured to receive traffic.
                             NOTE

                            In some documents, Huawei uses active/inactive or primary/backup to describe the
                            PW status. These terms have the same meaning as the term active/standby defined in
                            drafts. They all indicate the PW forwarding status.


Implementation
                    When the PW redundancy protection mechanism is introduced, to retain the
                    original forwarding behavior, PEs on both ends must use the same PW in a PW
                    protection group to transmit service data. This ensures that only one PE in the PW
                    protection group works as the primary PW and the others work as secondary PWs.
                    To achieve this goal, the PW protection group must use a signaling mechanism.

                    Relevant standards in LDP PW signaling specify the PW Status TLV to transmit the
                    PW forwarding status. The PW Status TLV, a 32-bit status code field, is carried in a
                    Notification message. PW redundancy introduces a new PW status code of
                    0x00000020 indicating PW forwarding standby, which means that a PW is a
                    secondary PW.

                    Forwarding priorities (primary or secondary) must be configured for PWs that
                    back up each other. The PW with the highest priority is selected as the primary
                    PW to forward traffic. The remaining PWs function as the secondary PWs to
                    protect the primary PW.

                    The forwarding status (active or standby) of a PW determines whether the PW is
                    used to forward traffic. The PW forwarding status depends on:

                    ●   Local and remote PW signaling statuses: A PE monitors its local PW signaling
                        status and uses an LDP Notification message to obtain the remote PW
                        signaling status from a remote PE.
                    ●   PW redundancy mode: The master/slave or independent mode is specified on
                        PE1.
                    ●   PW forwarding priority: determines whether a PW is the primary or secondary
                        PW and is specified on PE1.

                    In Figure 6-26, VPLS PW redundancy is configured on PE1. In normal cases, the
                    local and remote PW signaling statuses on PE1 are both up. PEs at the two ends
                    of a PW in different VPLS PW redundancy modes use different methods to select
                    the same PW for transmitting user packets.

                    ●   In master/slave mode, PE1 determines the local PW forwarding status based
                        on configured forwarding priorities and notifies the remote PE of the PW
                        forwarding status. The remote PE then determines its PW forwarding status
                        based on the received PW primary/ secondary status.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                    700
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration


                    ●   In Independent mode, PE1 determines the local PW forwarding status based
                        on the forwarding status notified by the remote PE. In this case, the remote
                        PE determines its master/backup status based on signaling (such as VRRP)
                        and notifies the local PE of the forwarding status.

                    In both master/slave and independent modes, if the primary PW is faulty, its
                    forwarding status becomes inactive and the forwarding status of the secondary
                    PW becomes active. PW-side faults do not affect the AC status. If AC-side faults
                    occur (for example, a PE or AC link is faulty), the PW primary/secondary status in
                    Independent mode will change, as the status is determined by signaling on dual-
                    homing devices. In master/slave mode, if an AC-side fault occurs, the PW primary/
                    secondary status remains unchanged, as the status is determined by the
                    configured forwarding priorities.

                         NOTE

                        VPLS PW redundancy is similar to VPWS PW redundancy. The major difference is that a
                        VPLS VSI has multiple PWs to different PEs. These PWs may form multiple PW redundancy
                        groups. PW switchover in a PW redundancy group does not affect other PW redundancy
                        groups.
                        When configuring VPLS PW redundancy, you are advised to configure consistent parameters
                        for the primary and secondary PWs. Otherwise, the secondary PW may fail to take over
                        services when the primary PW fails, leading to service interruption.


Derivative Functions
                    In addition to real-time protection against network faults, VPLS PW redundancy
                    allows users to manually switch traffic between PWs in a PW protection group
                    during network operation and maintenance. For example, after a PW protection
                    group is configured, if a device on the primary PW needs to be maintained, you
                    can switch traffic to the secondary PW and switch traffic back to the primary PW
                    after the maintenance.

Application Scenario
                    VPLS PW redundancy can be used on HVPLS networks and VPWS accessing VPLS
                    networks. These two types of networks can carry all services.

                    VPLS PW redundancy can also be used to improve existing network reliability. On
                    the VPLS network shown in Figure 6-26, CE1 can communicate with CE2, CE3, and
                    CE4 through PWs that are established between a VSI (on PE1) and each of PE2,
                    PE3, and PE4. On one hand, as services develop, services between CE1 and CE2,
                    and between CE1 and CE3 require high reliability. On the other hand, services
                    between CE1 and CE4 do not require high reliability. To meet the reliability
                    requirements, PE5 and PE6 are deployed on the VPLS network to provide VPLS PW
                    redundancy protection for PE2 and PE3. In addition, multiple PW protection
                    groups are configured in one VSI on PE1 to access remote PEs. Links between CE1
                    and CE4 remain unchanged. VPLS PW redundancy protects services against
                    failures on the network, ACs, and PEs without affecting existing services, ensuring
                    high network reliability.

                         NOTE


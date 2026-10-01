---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-296
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [43736, 43898]
sha256: d73e306e3975bbd044193d1108da5635fcb62574528ddf3b6d32f33d5ce32cad
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        VPLS PW redundancy can be configured for desired services without affecting services on
                        other PWs, reducing costs and maximizing profits.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    701
VPN Configuration
VPN Configuration                                                                       6 VPLS Configuration


6.11.2 Configuring a PW

Context
                    The following describes how to configure a PW in two networking modes: HVPLS
                    and VPLS accessing VPWS. Currently, only LDP PWs can be configured.

Procedure
                    ●     Configure PWs in HVPLS networking. For configuration details, see
                          Configuring LDP HVPLS.
                    ●     Configure a PW in VPLS accessing VPWS networking.
                          –         Configure UPEs. For configuration details, see Configuring LDP HVPLS.
                          –         Configure NPEs. For details, see Configuring PW Redundancy in
                                    Master/Slave Mode.
                                    To implement VPLS PW redundancy in master/slave mode, you must
                                    configure a bypass PW.

                    ----End

6.11.3 Adding PWs to a PW Protection Group

Context
                    After configuring two VSI PWs on a UPE, add them to a PW protection group,
                    specify their priorities, and configure PW protection group parameters for the two
                    PWs to work in backup mode.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the created VSI.
                    vsi vsi-name

         Step 3 Enter the VSI-LDP view.
                    pwsignal ldp

         Step 4 Configure a VSI ID.
                    vsi-id vsi-id

         Step 5 Create a PW protection group and enter its view.
                    protect-group group-name

         Step 6 Specify the PW redundancy mode of the current PW protection group.
                    protect-mode pw-redundancy { master | independent }

                    By default, a PW protection group does not have a PW redundancy mode.

                    In HVPLS networking, the master mode must be configured. In VPLS accessing
                    VPWS networking, if the master mode is configured, a bypass PW must be
                    configured between NPEs.

Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                        702
VPN Configuration
VPN Configuration                                                                               6 VPLS Configuration


                          NOTE

                    After a PW protection group is configured, you must specify the PW redundancy mode before
                    configuring other parameters. Deleting the PW redundancy mode of a PW protection group will
                    clear all the configurations of the group.

         Step 7 Add a PW to the PW protection group, and specify the priority of the PW. A
                smaller value indicates a higher priority. Of the two PWs added to a PW
                protection group, the one with a higher priority functions as the primary PW.
                    peer peer-address [ negotiation-vc-id vc-id ] preference preference-value

                    By default, no PW is added to a PW protection group.
                    Add PWs to a PW protection group in the descending order of PW priorities. Do
                    not add the PW with a lower priority to the PW protection group first.
         Step 8 (Optional) Configure both the primary and secondary PWs in the PW protection
                group to receive packets.
                    stream-dual-receiving

                    By default, only the primary PW can receive packets.

                          NOTE

                    When PW redundancy is deployed on a network, you are advised to run this command to
                    prevent packet loss during PW switchback. However, after this command is run, a loop may
                    occur.
                    This command can be run only when both the primary and secondary PWs exist.

         Step 9 (Optional) Configure a switching delay for the PW protection group in master/
                slave PW redundancy mode.
                    holdoff holdoffTime

                    By default, no switching delay is configured for a PW protection group in master/
                    slave PW redundancy mode. In this case, if a fault occurs on the primary PW,
                    traffic is immediately switched from the primary PW to the secondary PW.

                          NOTE

                         A PW protection group in independent PW redundancy mode does not support delayed
                         switching.
                         On a VPLS network that uses BFD for fault detection, traffic is immediately switched from
                         the primary PW to the secondary PW after BFD detects a fault on the primary PW. It is
                         recommended that you configure BFD or delayed switching based on your requirements.
                         After you configure a switching delay (using holdoffTime), traffic forwarded during the
                         delay will be interrupted if the primary PW fails to recover before the delay expires.

        Step 10 (Optional) Configure a switchback policy for the PW protection group in master/
                slave PW redundancy mode.
                    reroute { delay delay-time | immediately | never }

                    By default, the switchback policy for a PW protection group in master/slave PW
                    redundancy mode is delayed switchback, with a default delay of 30s; for a PW
                    protection group in independent PW redundancy mode, the switchback policy is
                    immediate switchback.
                    The switchback policy affects traffic switchback after the primary PW recovers.

                    ----End

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       703
VPN Configuration
VPN Configuration                                                                                  6 VPLS Configuration


6.11.4 (Optional) Associating Spoke PW Status with Hub PW
Status
Context
                    Figure 6-27 shows a VPLS PW redundancy scenario. The UPE is connected to SPE1
                    over PW1, a primary PW, and is connected to SPE2 over PW2, a secondary PW.
                    PW1 and PW2 work in backup mode. SPEs and NPEs are fully meshed over hub
                    PWs.

                    Figure 6-27 VPLS PW redundancy networking




                    In normal cases, if all hub PWs connected to SPE1 go down but PW1 is up, uplink
                    traffic still travels along PW1, the primary PW. As a result, traffic is lost. To prevent
                    this problem, associate the spoke PW status with the hub PW status on the SPE
                    where the primary spoke PW resides. Then, after all the hub PWs of the SPE where
                    the primary spoke PW resides go down, the SPE instructs the UPE to switch traffic
                    to the secondary spoke PW.
                    Perform the following steps on the SPE:

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the created VSI.
                    vsi vsi-name

         Step 3 Enter the VSI-LDP view.
                    pwsignal ldp

         Step 4 Configure a VSI ID.
                    vsi-id vsi-id


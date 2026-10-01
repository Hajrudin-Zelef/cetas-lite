---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-72
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [10222, 10369]
sha256: cec837f0700311fbc489bd3df557f7483b6e5d434710c8862860ff4a90a36322
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          186
QoS Configuration
QoS Configuration                                                                      12 MPLS QoS Configuration


                                   NOTE

                                 ● The DiffServ mode can be used only on PEs.
                                 ● After exp-value is specified in the diffserv-mode { pipe { mpls-exp exp-value
                                   | domain ds-name } | short-pipe [ mpls-exp exp-value ] domain ds-name }
                                   command, the EXP value in the MPLS private network label is set to the
                                   specified value.
                    ●   Configuring a DiffServ mode for L2VPN (VLL networking mode)
                              NOTE

                             Only the S6780-H, S6750-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2, S5755-S,
                             S5755E-H, S5755-H and S5732-H-V2 series support this configuration.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the AC-side interface view.
                             interface interface-type interface-number

                        c.   Change the interface working mode to Layer 3.
                             undo portswitch

                             Only the S6780-H, S6750-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-
                             V2, S5755-S, S5755E-H, S5755-H and S5732-H-V2 series support this step.

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Configure a DiffServ mode for VLL.
                             diffserv-mode { pipe { mpls-exp exp-value | domain ds-name } | short-pipe [ mpls-exp exp-
                             value ] domain ds-name | uniform [ domain ds-name ] }

                             By default, the DiffServ mode supported by VLL is uniform.

                                   NOTE

                                 ● The DiffServ mode can be used only on PEs.
                                 ● After exp-value is specified in the diffserv-mode { pipe { mpls-exp exp-value
                                   | domain ds-name } | short-pipe [ mpls-exp exp-value ] domain ds-name }
                                   command, the EXP value in the MPLS private network label is set to the
                                   specified value.
                    ●   Configuring a DiffServ mode for L2VPN (VPLS networking mode)
                              NOTE

                             Only the S6780-H, S6750-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2, S5755-S,
                             S5755E-H, S5755-H and S5732-H-V2 series support this configuration.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the VSI view.
                             vsi vsi-name

                        c.   Configure a DiffServ mode for VPLS.
                             diffserv-mode { pipe { mpls-exp exp-value | domain ds-name } | short-pipe [ mpls-exp exp-
                             value ] domain ds-name | uniform [ domain ds-name ] }

                             By default, the DiffServ mode supported by VPLS is uniform.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          187
QoS Configuration
QoS Configuration                                                                   12 MPLS QoS Configuration


                                    NOTE

                                   ● The DiffServ mode can be used only on PEs.
                                   ● After exp-value is specified in the diffserv-mode { pipe { mpls-exp exp-value
                                     | domain ds-name } | short-pipe [ mpls-exp exp-value ] domain ds-name }
                                     command, the EXP value in the MPLS private network label is set to the
                                     specified value.

                    ----End


12.6 Configuring Priority Mapping
Context
                    The process of configuring the priority mapping function is as follows:

                    1.   Configure a DiffServ domain to determine mappings between EXP values and
                         internal priorities so that the device can provide differentiated services based
                         on internal priorities.
                    2.   Apply the DiffServ domain to an object so that the mappings in the DiffServ
                         domain take effect.

12.6.1 Configuring a DiffServ Domain

Context
                    When the device is used as an edge node connecting a DiffServ domain and
                    another network, you need to configure mappings between internal and external
                    priorities.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a DiffServ domain and enter the DiffServ domain view.
                    diffserv domain ds-domain-name

                    By default, the default domain exists. It can be modified but not deleted. For
                    details about the default priority mapping, see Table 12-2 and Table 12-3 in
                    "Default Settings for MPLS QoS."

                    You can also create a DiffServ domain and define the priority mapping.

         Step 3 Define the priority mapping of the device based on actual requirements.
                    ●    Map EXP values of incoming MPLS packets on an interface to PHBs and
                         colors.
                         mpls-exp-inbound exp-value phb service-class color

                    ●    Map PHBs and colors of outgoing MPLS packets on an interface to EXP
                         values.
                         mpls-exp-outbound service-class color map exp-value

                    ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   188
QoS Configuration
QoS Configuration                                                                        12 MPLS QoS Configuration


12.6.2 Applying a DiffServ Domain
Context
                    You can bind a DiffServ domain to the inbound interface of packets so that the
                    device maps priorities of incoming packets to corresponding PHBs and colors
                    according to the mappings defined in the DiffServ domain.
                    You can bind a DiffServ domain to the outbound interface of packets so that the
                    device maps PHBs and colors of outgoing packets to corresponding priorities
                    according to the mappings defined in the DiffServ domain.

                         NOTE

                        ● When MPLS is disabled globally, no default configuration exists and this function does
                          not take effect.
                        ● After MPLS is enabled globally, the default configuration exists. This function takes
                          effect only when it is configured before a public network tunnel is established.
                        ● If this function is configured after a public network tunnel is established and needs to
                          take effect immediately, you must restart the MPLS LDP session.
                        ● If the DiffServ mode of an interface, VSI, or VPN instance or the global DiffServ domain
                          is changed and the configuration needs to take effect immediately, you must re-
                          establish an MPLS LDP session.



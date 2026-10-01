---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-96
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["agent", "copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [11937, 12098]
sha256: a48e85791c51df3f9daf67b349725ee477ae7ddefd85d4242b3b3a73b832b8ea
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Purpose
                  Point-to-Point Protocol over Ethernet (PPPoE) provides access services for hosts on
                  an Ethernet through a remote access device, and implements access control and
                  accounting on each connected host. It uses the client/server model, in which a
                  PPPoE client sends a connection request to the PPPoE server and, through client-
                  server negotiation, the server provides functions such as access control and
                  authentication for the client.
                  Although PPPoE provides good authentication and security mechanisms, it has
                  some limitations. For example, the PPPoE server authenticates a user only using
                  the user name and password. If the account is compromised, it can be used
                  elsewhere to access the network and pass the PPPoE authentication, causing
                  RADIUS service embezzlement. To prevent this, PPPoE+ is introduced.


11.2 PPPoE+ Fundamentals
                  PPPoE involves three stages: Discovery, Session, and Terminate. PPPoE+ is applied
                  mainly in the Discovery and Session stages. The following figure shows the PPPoE
                  + working process.

                  Figure 11-2 PPPoE+ working process




                  1.     A user host (PPPoE client) sends a PPPoE Active Discovery Initiation (PADI)
                         packet.
                  2.     The Device obtains the PADI packet, adds to this packet a PPPoE+ tag
                         containing information about the interface connected to the PPPoE client
                         (such as the slot ID/subcard ID/interface number, VLAN ID, and MAC address),
                         and forwards the packet to the BRAS (PPPoE server).
                  3.     After receiving the tagged PADI packet, the BRAS sends a PPPoE Active
                         Discovery Offer (PADO) packet to the user host.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              219
Security Configuration
Security Configuration                                                           11 PPPoE+ Configuration


                  4.     After receiving the PADO packet, the user host sends a PPPoE Active
                         Discovery Request (PADR) packet.
                  5.     After obtaining the PADR packet, the Device adds to this packet a PPPoE+ tag
                         and sends the packet to the BRAS.
                  6.     After receiving the tagged PADR packet, the BRAS generates a unique PPP
                         session ID that identifies the session with the user host, and sends a PPPoE
                         Active Discovery Session-confirmation (PADS) packet to the user host. If no
                         error occurs, the BRAS and user host enter the Session stage.
                  7.     At the Session stage, PPP negotiation is performed and PPP packets are
                         transmitted between the user host and the BRAS. After PPP negotiation is
                         complete, the BRAS encapsulates a PPPoE+ tag in the NAS-Port-ID attribute
                         of RADIUS packets and sends the packets to the RADIUS server. Based on the
                         value of this attribute, the RADIUS server authenticates the user account
                         together with the access interface.
                  8.     After a PPPoE session is established, the PPPoE client or PPPoE server can
                         send a PPPoE Active Discovery Terminate (PADT) packet at any time to
                         terminate the session.


11.3 Configuration Precautions for PPPoE+

11.4 Default Settings for PPPoE+
                  Table 11-1 describes the default settings for PPPoE+.

                  Table 11-1 Default settings for PPPoE+
                   Parameter                                   Default Setting

                   Global PPPoE+                               Disabled

                   Trusted interface                           None

                   Policy for processing original              replace
                   information fields in user-side PPPoE
                   packets

                   Format and contents of information          circuit-id and remote-id in common
                   fields added to PPPoE packets               format

                   Vendor ID added to PPPoE packets            2011

                   Policy for processing original              No processing
                   information fields in PPPoE reply
                   packets from a PPPoE server




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             220
Security Configuration
Security Configuration                                                                    11 PPPoE+ Configuration




11.5 Configuring PPPoE+
Prerequisites
                  You have configured PPPoE on downstream hosts and the upstream PPPoE server.
                  PPPoE authentication can be performed successfully.

                          NOTE

                         After PPPoE+ is configured, the device encapsulates user-side network information into
                         PPPoE packets sent by users for PPPoE authentication and RADIUS authentication. When
                         the network configuration on the user side changes, if the PPPoE+ configuration on the
                         device is not updated in a timely manner, PPPoE users may fail to go online.


11.5.1 Enabling PPPoE+

Context
                  To prevent user accounts from being compromised, you can configure the PPPoE+
                  function. You must enable PPPoE+ globally before configuring functions related to
                  PPPoE+.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable PPPoE+ globally.
                  pppoe intermediate-agent information enable

                  After the command is executed in the system view, PPPoE+ is enabled on all
                  interfaces.

                  By default, PPPoE+ is disabled globally.

                  ----End

11.5.2 Configuring a PPPoE+ Trusted Interface

Context
                  To prevent spoofing of a PPPoE server and the security risk caused by PPPoE
                  packets being forwarded to non-PPPoE service interfaces, you can configure the
                  interface connecting a device to a PPPoE server as a trusted interface. This
                  configuration ensures that PPPoE packets are forwarded to the PPPoE server
                  through the trusted interface only. It also ensures that only the PPPoE packets
                  received on the trusted interface are forwarded to a PPPoE client.

                          NOTE

                         A trusted interface controls protocol packets at the PPPoE Discovery stage only; it does not
                         control service packets at the PPPoE Session stage.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       221
Security Configuration
Security Configuration                                                                      11 PPPoE+ Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the interface view.
                  interface interface-type interface-number

         Step 3 Configure the interface as a trusted interface.
                  pppoe uplink-port trusted

                  ----End


---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-97
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["agent", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [12099, 12220]
sha256: aeab915c40c4dcd08c4080822e11699465f62c147b4df43e88c103aed8892831
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

11.5.3 Configuring the Policy for Processing User-Side PPPoE
Packets
Context
                  After you configure the policy for processing user-side PPPoE packets, the device
                  can add to PPPoE packets information about the interface connected to a user
                  host. This makes it possible to authenticate user accounts together with the access
                  interfaces, thereby preventing accounts from being compromised.
                  You can configure the policy for processing original information fields in user-side
                  PPPoE packets in the system view or interface view. The configuration in the
                  system view takes effect for all interfaces. An interface preferentially uses the
                  policy that is configured on the interface.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure the policy for processing original information fields in user-side PPPoE
                packets.
                  ●      Configure the policy globally in the system view.
                         This configuration takes effect for all interfaces.
                         pppoe intermediate-agent information policy { drop | keep | replace }

                         By default, an interface processes original information fields in received user-
                         side PPPoE packets in replace mode.
                                 NOTE

                             –    drop: If a received PPPoE packet contains a tag, the device removes the original
                                  information fields from the packet. If a received PPPoE packet does not contain a
                                  tag, the device does not process the packet.
                             –    replace: If a received PPPoE packet contains a tag, the device replaces the original
                                  information fields in the PPPoE packet according to the specified field format. If a
                                  received PPPoE packet does not contain a tag, the device adds a tag to the PPPoE
                                  packet according to the specified field format.
                             –    keep: If a received PPPoE packet contains a tag, the device does not process the
                                  packet. If a received PPPoE packet does not contain a tag, the device adds a tag to
                                  the packet according the specified field format.
                  ●      Configure the policy in the interface view.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       222
Security Configuration
Security Configuration                                                                        11 PPPoE+ Configuration


                         a.   Enter the interface view.
                              interface interface-type interface-number

                         b.   Configure the policy for processing original information fields in user-side
                              PPPoE packets on the interface.
                              pppoe intermediate-agent information policy { drop | keep | replace }

                         c.   (Optional) Configure the formats of information fields to be added to
                              PPPoE packets.
                              pppoe intermediate-agent information [ vlan vlan-id ] [ ce-vlan cevlan-id ] format { circuit-
                              id | remote-id } { common | extend | user-defined text }

                              By default, the device adds information fields circuit-id and remote-id of
                              the common format to PPPoE packets.
                              If this command is configured in both the interface view and system view,
                              the configuration in the interface view takes effect preferentially.
                         d.   Return to the system view.
                              quit

         Step 3 (Optional) Configure the formats and contents of information fields used to
                replace the original information fields in user-side PPPoE packets when the policy
                for processing these packets is set to replace.
                  1.     Configure the content of information fields to be added to PPPoE packets.
                         pppoe intermediate-agent information encapsulation { circuit-id | remote-id } *

                  2.     Configure the formats of information fields to be added to PPPoE packets.
                         pppoe intermediate-agent information format { circuit-id | remote-id } { common | extend | user-
                         defined text }

                         By default, the device adds information fields circuit-id and remote-id of the
                         common format to PPPoE packets.
         Step 4 (Optional) Configure the vendor ID to be added to PPPoE packets.
                  pppoe intermediate-agent information vendor-id vendor-id

                  By default, the device adds vendor ID 2011 to PPPoE packets.

                          NOTE

                         A vendor ID identifies a vendor. After PPPoE+ is enabled, the device can perform PPP
                         negotiation with the PPPoE server using only PPPoE packets that contain a specified vendor
                         ID. By default, the device adds vendor ID 2011 to PPPoE packets. If the device connects to a
                         non-Huawei PPPoE server and a different vendor ID (for example, 3561) is required, you
                         can run the pppoe intermediate-agent information vendor-id vendor-id command to
                         change the vendor ID.

                  ----End

11.5.4 (Optional) Configuring the Policy for Processing Server-
Side PPPoE Packets
Context
                  Generally, the device does not need to process PPPoE reply packets from a PPPoE
                  server and instead transparently transmits them to PPPoE clients. However, in
                  cases where the clients cannot identify such packets, the device processes these
                  packets to ensure successful establishment of PPPoE sessions. The PPPoE reply
                  packets are processed as follows: The PPPoE reply packets are processed as
                  follows:

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                         223
Security Configuration
Security Configuration                                                                 11 PPPoE+ Configuration


                  ●      When the policy for processing original information fields in PPPoE packets is
                         replace or keep:
                         –   If PPPoE reply packets from the PPPoE server do not contain information
                             fields, the device transparently transmits these packets.
                         –   If PPPoE reply packets from the PPPoE server contain information fields,
                             the device determines whether these fields have the same formats and
                             contents as those added to the user-side PPPoE packets. If they do, the
                             device removes the information fields from the PPPoE reply packets
                             before forwarding them. If they do not, the device transparently transmits
                             them.
                  ●      When the policy for processing original fields in PPPoE packets is drop, the
                         device transparently transmits PPPoE reply packets from the server.

                         NOTE


---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-41
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [5172, 5286]
sha256: 5b99f534c0a59249380946e3cdc5913425da276a3757d4041d9d18c0b309a533
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   Dynami    When port security    After the device           If users frequently
                   c         is enabled on an      restarts or the            change their access
                   secure    interface, dynamic    interface goes down,       locations, you can
                   MAC       MAC address entries   the entries are lost       configure the dynamic
                   address   that have been        and need to be             secure MAC address
                             learned by the        relearned.                 function on the user-
                             interface are         Dynamic secure MAC         side interface of the
                             deleted, and MAC      addresses do not age       device. This function
                             address entries       by default. They age       ensures security and
                             learned               only when the set          deletes bound MAC
                             subsequently are      aging time is reached.     address entries
                             converted into                                   immediately when
                             dynamic secure        The number of MAC          they age.
                             MAC address           addresses can be
                             entries.              limited.
                                                   Protection actions can
                                                   be configured for the
                                                   interface, including:
                                                   discarding the packets,
                                                   reporting alarms, or
                                                   shutting down the
                                                   interface.




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                             92
Security Configuration
Security Configuration                                                              6 Port Security Configuration


                   Type          Description                Characteristic               Scenario

                   Sticky        When the sticky            The entries are not          If users seldom
                   MAC           MAC address                lost after the device        change access
                   address       function is enabled        restarts or the              locations, you can
                                 following port             interface goes down.         configure the sticky
                                 security, the              Sticky MAC address           MAC address function
                                 dynamic MAC                entries will not age.        on the user-side
                                 address entries that                                    interface of the
                                 have been                  The number of MAC            device. This function
                                 converted into             addresses can be             ensures security
                                 dynamic secure             limited.                     without losing bound
                                 MAC address entries        Protection actions can       MAC address entries.
                                 will be converted          be configured for the
                                 into sticky MAC            interface, including:
                                 address entries, and       discarding the packets,
                                 the newly learned          reporting alarms, or
                                 dynamic MAC                shutting down the
                                 address entries will       interface.
                                 be directly
                                 converted into
                                 sticky MAC address
                                 entries.

                   Static        Static secure MAC          The entries are not          If users are few and
                   secure        addresses refer to         lost after the device        seldom change their
                   MAC           MAC addresses that         restarts or the              access locations, you
                   address       are manually               interface goes down.         can configure the
                                 configured on an           Static secure MAC            static secure MAC
                                 interface where port       address entries will         address function on
                                 security is enabled.       not age.                     the user-side interface
                                                                                         of the device to
                                                                                         manually bind the
                                                                                         MAC address entries.




                          NOTE

                         After the sticky MAC address function is disabled on an interface, sticky MAC address
                         entries on the interface are converted into dynamic secure MAC address entries. After port
                         security is disabled on an interface, existing dynamic secure MAC address entries on the
                         interface are deleted. The interface then learns dynamic MAC address entries again.


Aging of Secure MAC Addresses
                  Only dynamic secure MAC addresses will age when the set aging time is reached.
                  ●      Absolute aging time: If the absolute aging time is 5 minutes, the system
                         checks whether there is traffic from the MAC address every 5 minutes. If not,
                         the dynamic secure MAC address is aged out immediately.
                  ●      Relative aging time: If the relative aging time is 5 minutes, the system checks
                         whether there is traffic from the MAC address every 1 minute. If not, the
                         dynamic secure MAC address is aged out 5 minutes later.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       93
Security Configuration
Security Configuration                                                               6 Port Security Configuration


                  ●      Forcible aging time: The system calculates the lifetime of each MAC address
                         once every minute. For example, if the forcible aging time is set to 5 minutes
                         and the lifetime of a dynamic secure MAC address is greater than or equal to
                         5 minutes, this MAC address ages out immediately.

Port Security Protection Actions
                  After port security is enabled on an interface, if the interface receives packets
                  sourced from a MAC address not existing in the MAC address table, the device
                  considers that the packets are sent from an unauthorized user, regardless of
                  whether the destination MAC address of packets exists, and takes the protection
                  actions.

                          NOTE

                         If the static MAC address flapping detection function is enabled on the device, it also takes
                         protection actions against the invalid packets (packets whose source MAC address is a static
                         MAC address entry on another interface) it detects.


                  Table 6-2 Port security protection actions
                   Action               Description


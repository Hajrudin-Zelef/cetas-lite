---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-42
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [5287, 5448]
sha256: 57d4cbbc4b4a330b17a86afe712334626c8ce9cd613ae66f86f9daaff8205206
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   restrict             Discards invalid packets and reports an alarm. This action is
                                        recommended.

                   protect              Discards invalid packets without reporting an alarm.

                   error-down           Discards invalid packets, sets the interface status to error-down,
                                        and reports an alarm.
                                        By default, an interface in error-down state can be restored only
                                        after the restart command is run in the interface view.
                                        To enable an interface in error-down state to automatically go
                                        up after a period of time, run the error-down auto-recovery
                                        cause portsec-reachedlimit interval interval-value command in
                                        the system view before the interface status becomes error-down.
                                        In this command, interval-value specifies the period of time
                                        after which an interface can automatically go up.




6.3 Configuration Precautions for Port Security

6.4 Default Settings for Port Security
                  Table 6-3 describes the default settings for port security.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        94
Security Configuration
Security Configuration                                                               6 Port Security Configuration


                  Table 6-3 Default settings for port security

                   Parameter                                              Default Setting

                   Port security                                          Disabled

                   Maximum number of secure MAC                           1
                   addresses on an interface

                   Port security protection action                        restrict

                   Aging time of dynamic secure MAC                       Disabled
                   addresses




6.5 Configuring Port Security
Prerequisites
                  The interface has had Smart Link, SEP, and ERPS disabled. Otherwise, loops
                  cannot be prevented.

                  The MAC address learning restriction function has been disabled on the interface.

6.5.1 Configuring the Dynamic Secure MAC Address Function

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the interface view.
                  interface { interface-name | interface-type interface-number }

         Step 3 Change the interface working mode from Layer 3 to Layer 2.
                  portswitch

                  Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                  S6750E-S, S5732-H-V2, S5755-S , S5755-H , S5755E-H series can be switched from
                  Layer 3 mode to Layer 2 mode using the portswitch command.

         Step 4 Enable the port security function and set the maximum number of MAC addresses
                that the interface can learn.
                  port-security enable [ maximum max-number ]

                          NOTE

                         After the port security function is enabled, the dynamic secure MAC address function takes
                         effect.

         Step 5 (Optional) Configure a protection action for port security on an interface.
                  port-security protect-action { protect | restrict | error-down }

         Step 6 (Optional) Set the aging time for dynamic secure MAC address entries.
                  port-security aging-time time [ type { absolute | inactivity } ]


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      95
Security Configuration
Security Configuration                                                                    6 Port Security Configuration


                  If the aging time is too short (for example, 1 minute), dynamic secure MAC
                  address entries will age fast, which may cause traffic forwarding failures.
         Step 7 Exit the interface view.
                  quit

         Step 8 (Optional) Delete the dynamic secure MAC address entries.
                  undo mac-address security { [ interface-type interface-number | interface-name ] | [ vlan vlanId ] } *

                  ----End

Verifying the Configuration
                  ●      Run the display port-security [ interface { interface-typeinterface-number |
                         interface-name } ] command to check port security information.
                  ●      Check the port security-related alarms by running the display trapbuffer
                         command or through the NMS.

6.5.2 Configuring the Sticky MAC Address Function
Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the interface view.
                  interface { interface-name | interface-type interface-number }

         Step 3 Change the interface working mode from Layer 3 to Layer 2.
                  portswitch

                  Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                  S6750E-S, S5732-H-V2, S5755-S , S5755-H , S5755E-H series can be switched from
                  Layer 3 mode to Layer 2 mode using the portswitch command.
         Step 4 Enable the port security function and set the maximum number of MAC addresses
                that the interface can learn.
                  port-security enable [ maximum max-number ]

         Step 5 Enable the sticky MAC address function.
                  port-security mac-address sticky

                          NOTE

                         After the sticky MAC address function is enabled on an interface, existing dynamic secure
                         MAC address entries and MAC address entries learned subsequently on the interface are
                         converted into sticky MAC address entries.
                         After the sticky MAC address function is enabled on an interface, sticky MAC address
                         entries will not age even if the port-security aging-time command is configured.
                         When new sticky MAC address entries are configured or existing sticky MAC address entries
                         are changed, you need to run the save command to ensure that the new or changed MAC
                         address entries take effect after the device restarts.
                         Sticky MAC address entries are saved into a .dtbl, .ztbl, or .ctbl file by running the save
                         command. This ensures that sticky MAC address entries will not be lost when the device
                         restarts. The file name must be the same as the name of the system configuration file. For
                         example, if the name of the system configuration file is test.cfg, the name of the sticky
                         MAC address entry file must be test.ctbl. Otherwise, sticky MAC address entries will fail to
                         recover after the device restarts.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                96
Security Configuration
Security Configuration                                                                  6 Port Security Configuration


         Step 6 (Optional) Configure a protection action for port security on an interface.
                  port-security protect-action { protect | restrict | error-down }

         Step 7 (Optional) Add a sticky MAC address entry manually.
                  port-security mac-address sticky mac-address vlan vlan-id


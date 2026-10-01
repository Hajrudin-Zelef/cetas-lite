---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-43
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [5449, 5604]
sha256: 3868572b6798c4de131d658dd034ebf79a9e88d17baa31352854a216c27311c6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

         Step 8 Exit the interface view.
                  quit

         Step 9 (Optional) Delete a sticky MAC address entry.
                  undo mac-address sticky { [ portType portNum | portName ] | [ vlan vlanId ] } *

                  ----End

Verifying the Configuration
                  ●      Run the display port-security [ interface { interface-typeinterface-number |
                         interface-name } ] command to check port security information.
                  ●      Check the port security-related alarms by running the display trapbuffer
                         command or through the NMS.

6.5.3 Configuring the Static Secure MAC Address Function
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

         Step 5 Configure a static secure MAC address entry.
                  port-security mac-address mac-address vlan vlan-id

                          NOTE

                         Static secure MAC address entries will not age even if the port-security aging-time
                         command is configured.

         Step 6 Exit the interface view.
                  quit

         Step 7 (Optional) Delete the static secure MAC address entries.
                  undo mac-address sec-config { [ portType portNum | portName ] | [ vlan vlanId ] } *

                  ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       97
Security Configuration
Security Configuration                                                               6 Port Security Configuration


Verifying the Configuration
                  ●      Run the display port-security [ interface { interface-typeinterface-number |
                         interface-name } ] command to check port security information.
                  ●      Check the port security-related alarms by running the display trapbuffer
                         command or through the NMS.

6.5.4 Configuring Static MAC Address Flapping Detection

Context
                  Assume that a static secure MAC address is configured on interface B of the device
                  by running the mac-address static command, enabling a user to access the
                  device. When the user is disconnected from interface A and then connected to
                  interface B, interface A will receive packets with a source MAC address that is the
                  static MAC address of interface B. The device considers this as static MAC address
                  flapping and discards the packets. As a result, the user cannot access the device.

                  To configure the device to report an alarm when static MAC address flapping
                  occurs, enable the static MAC address flapping detection function on the device
                  and port security on interface A. In this way, interface A will perform port security
                  protection actions like reporting an alarm when static MAC address flapping
                  occurs.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable static MAC address flapping detection.
                  port-security static-flapping protect

         Step 3 Enter the interface view.
                  interface { interface-name | interface-type interface-number }

         Step 4 Change the interface working mode from Layer 3 to Layer 2.
                  portswitch

                  Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                  S6750E-S, S5732-H-V2, S5755-S , S5755-H , S5755E-H series can be switched from
                  Layer 3 mode to Layer 2 mode using the portswitch command.

         Step 5 Enable the port security function and set the maximum number of MAC addresses
                that the interface can learn.
                  port-security enable [ maximum max-number ]

         Step 6 (Optional) Configure a protection action for port security on an interface.
                  port-security protect-action { protect | restrict | error-down }

         Step 7 Exit the interface view.
                  quit

                  ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    98
Security Configuration
Security Configuration                                                                6 Port Security Configuration


Verifying the Configuration
                  Check the port security-related alarms by running the display trapbuffer
                  command or through the NMS.

6.5.5 Configuring the Function of Allowing Users to Go Online
from Different Broadcast Domains in the Scenario Where Port
Security and Authentication Are Both Deployed
Prerequisites
                  Port security is enabled, and the action for port authentication is set to Err-Down.

Context
                  In the user authorization change scenario, the sticky MAC address is not deleted
                  when a user goes offline from VLAN 10. When the user with the same MAC
                  address goes online again in VLAN 20, the device checks the MAC address validity
                  and finds that the MAC address already exists in VLAN 10. If the MAC address is
                  invalid, the port will go down and an alarm will be triggered. In this case, you can
                  enable the function of allowing users to go online from different broadcast
                  domains in the scenario where port security and authentication are both deployed
                  to ensure that only one MAC address entry matches each online user. If a user
                  goes online again from a different broadcast domain, the device deletes the
                  original MAC address entry of the user. After this function is enabled, port security
                  entries are not affected when a user goes online again from a different broadcast
                  domain. Therefore, the port will not go down and no alarm will be triggered.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable the function of allowing users to go online from different broadcast
                domains in the scenario where port security and authentication are both deployed.
                That is, users on the same port are matched based on MAC addresses.
                  port-security mac-address convert-from-dynamic multi-authen enhance enable

                          NOTE

                         This function does not take effect for the Muti-Share mode configured for authentication.
                         This function does not take effect for the static MAC address entries that are configured
                         using commands and have configuration files, except sticky MAC address entries.
                         After this function is enabled, security and sticky MAC address entries will be deleted
                         globally and users are triggered to go online again.

                  ----End


---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-112
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [14033, 14166]
sha256: 332f4e6ad6b6372a5fcfdacf38db59c45318fb59f86834e273abe44e988249a3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         By default, no source interface is configured for an SSH server.
                  ●      Configure the SSH server to use all valid interfaces as the source interface.
                         ssh server-source all-interface

                  ●      Configure a source IPv4 address for the SSH server.
                         ssh server-source -a ip-address [ -vpn-instance vpnName ]

                  ●      Configure a source IPv6 address for the SSH server.
                         ssh ipv6 server-source -a ipv6-address [ -vpn-instance vpn-instance-name ]

                         By default, no source interface or source IPv6 address is specified for an SSH
                         server.
                  ●      Configure the SSH server to use all valid interfaces using IPv6 addresses as
                         the IPv6 source interface.
                         ssh ipv6 server-source all-interface

                          NOTE

                         After the ssh server-source all-interface or ssh ipv6 server-source all-interface command
                         is run, users can log in to the SSH server from all valid interfaces, posing system security
                         risks. Exercise caution while running these two commands.

         Step 9 Configure the maximum number of connections that can be established between
                a single IP address and the SSH server.
                  ssh server ip-limit-session limit-session-num

                  By default, a maximum of 256 connections can be established on the SSH server
                  using a single IP address.

        Step 10 Enable keyboard-interactive authentication on the SSH server.
                  ssh server authentication-type keyboard-interactive enable

                  By default, keyboard-interactive authentication is enabled on an SSH server.

                  If an SSH user logs in to a device using password card authentication, keyboard-
                  interactive authentication must be enabled for this user.

        Step 11 (Optional) Enable the client IP address locking function on the SSH server.
                  undo ssh server ip-block disable

                  By default, the client IP address locking function is enabled on an SSH server.

                  ●      When this function is enabled, locked IP addresses fail to pass authentication
                         and are displayed in the display ssh server ip-block list command output.
                  ●      If this function is disabled, the display ssh server ip-block list command does
                         not display any previously locked IP addresses or IP addresses that fail to pass
                         authentication.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      258
Security Configuration
Security Configuration                                                                          13 SSH Configuration


                          NOTE

                         If an SSH user fails authentication six consecutive times within 5 minutes, the user's IP
                         address will be locked for 5 minutes. To unlock the IP address before the locking period
                         elapses, run the activate ssh server ip-block ip-address ip-address [ vpn-instance vpn-
                         name ] command.

        Step 12 (Optional) Configure alarm generation and clearance thresholds as a number of
                SSH server login failures within a specified period.
                  ssh server login-failed threshold-alarm upper-limit report-times lower-limit resume-times period period-
                  time

                  By default, an alarm is generated when 30 or more login failures occur within 5
                  minutes, and the alarm is cleared when the number of login failures within 5
                  minutes falls below 20.
        Step 13 (Optional) Configure the DSCP priority of SSH packets.
                  ssh server dscp value

                  By default, the DSCP priority of SSH packets is 48.
        Step 14 (Optional) Enable the local port forwarding service on the SSH server.
                  ssh server tcp forwarding enable

                  By default, the local port forwarding service is disabled on the SSH server.

                  ----End

Querying the Configuration
                  ●      Run the display dsa peer-public-key command to view detailed information
                         about the DSA public key.
                  ●      Run the display ecc peer-public-key command to check detailed information
                         about the ECC public key.
                  ●      Run the display rsa peer-public-key command to view detailed information
                         about the RSA public key.
                  ●      Run the display sm2 peer-public-key command to view detailed information
                         about the SM2 public key.
                  ●      Run the display sftp-client command to check the configuration of the SFTP
                         client.

13.5.2 Configuring the VTY User Interface to Support SSH
Context
                  Before logging in to the device through SSH, configure the VTY user interface that
                  will be used to support SSH.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the VTY user interface view.
                  user-interface vty first-ui-number [ last-ui-number ]


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                          259
Security Configuration
Security Configuration                                                                       13 SSH Configuration


                  By default, no user interface view is displayed.

         Step 3 Configure the AAA authentication mode for the VTY user interface.
                  authentication-mode aaa

                  By default, no authentication mode is used on a VTY user interface.

                  Configure the AAA authentication mode on the VTY user interface to enable SSH
                  support. If the AAA authentication mode is not set, the protocol inbound ssh
                  command does not take effect.

         Step 4 Configure the VTY user interface to support SSH.
                  protocol inbound { all | ssh }

                  By default, the VTY user interface supports all protocols, including SSH.

                  ----End

13.5.3 Configuring an SSH User


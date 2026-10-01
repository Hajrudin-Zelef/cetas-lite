---
id: collect-261001-huawei/huawei/network-ptmngsys-web-tsrev-s-en-content-s-42-edesk-mac-authentication-unsuccessf-28bea404
title: "network-ptmngsys-web-tsrev-s-en-content-s-42-edesk-mac-authentication-unsuccessf-28bea404"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/network-ptmngsys-web-tsrev-s-en-content-s-42-edesk-mac-authentication-unsuccessf-28bea404.md
source_anchor: ""
source_lines: [1, 57]
sha256: b5a9f0c188f75a4268fe067693bc57b59c33296b5089e594936e8f32453b780d
---

# network-ptmngsys-web-tsrev-s-en-content-s-42-edesk-mac-authentication-unsuccessf-28bea404

After you run the test-aaa command, the system displays the message Account test succeed!.
<HUAWEI> test-aaa user1 huawei123 radius-template huawei
Info: Account test succeed!
The message indicates that services on the switch are accessible to the RADIUS server.
After you run the test-aaa command, the system displays the message Account test time out..
<HUAWEI> test-aaa user1 huawei123 radius-template huawei
Info: Account test time out.
The message indicates that the switch does not receive response packets from the RADIUS server.
Run the ping command to check whether there are reachable routes between the switch and the RADIUS server.
Run the display radius-server configuration [ template template-name ] command to check whether the port number configured in the RADIUS server template is the same as that of the RADIUS server, and whether NAS-IP-Address in the template is the same as that of the RADIUS server.
<HUAWEI> display radius-server configuration template shiva
  ------------------------------------------------------------------------------
  Server-template-name          :  shiva
  Protocol-version              :  standard
  Traffic-unit                  :  B
  Shared-secret-key             :  %^%#O09i(W[^YT4g#Z37Nct9$IK#TH(-B6-1|<;q|D)"%^%#
  Group-filter                  :  class
  Timeout-interval(in second)   :  5
  Retransmission                :  2
  EndPacketSendTime             :  0
  Dead time(in minute)          :  5
  Domain-included               :  YES
  NAS-IP-Address                :  -
  Calling-station-id MAC-format :  xxxx-xxxx-xxxx
  Called-station-id MAC-format  :  XX.XX.XX.XX.XX.XX
  Service-type                  :  - 
  NAS-IPv6-Address              :  ::
  Server algorithm              :  master-backup 
  Detect-interval(in second)    :  60 
  Authentication Server 1       :  10.7.66.66     Port:1812  Weight:80  [UP]
                                   Vrf:- LoopBack:NULL
                                   Source IP: ::
  Authentication Server 2       :  10.7.66.67     Port:1812  Weight:80  [UP]
                                   Vrf:- LoopBack:NULL
                                   Source IP: ::
 Check whether the port is occupied by another program. If the port is occupied, exit the program.
If the Agile Controller functions as the RADIUS server, run the netstat -nao | findstr 1812 and netstat -nao | findstr 1813 commands on the server.
Check whether the IP address of the switch added on the RADIUS server is the same as the source IP address configured using the radius-server authentication command.
[HUAWEI] radius-server template controller
[HUAWEI-radius-controller] radius-server authentication 192.168.1.2 1812 source ip-address 192.168.1.11
If the source IP address is not configured, the switch uses the IP address of the outbound interface as the source IP address to send RADIUS packets to the RADIUS authentication server.
Check whether the switch and RADIUS server are on different virtual private networks (VPNs).
If they are on different VPNs, configure a VPN instance on the switch.
[HUAWEI] radius-server template controller
[HUAWEI-radius-controller] radius-server authentication 192.168.1.2 1812 vpn-instance vpntest
After you run the test-aaa command, the system displays a message indicating that the account test fails.
<HUAWEI> test-aaa user1 huawei123 radius-template huawei
Info: Account test failed.
Run the display radius-server configuration [ template template-name ] command to check whether a shared key and an IP address have been configured in the RADIUS server template.
If a shared key and an IP address have been configured, make sure that they are the same as those on the RADIUS server.
[HUAWEI] radius-server template controller
[HUAWEI-radius-controller] radius-server authentication 192.168.1.2 1812
[HUAWEI-radius-controller] radius-server shared-key cipher Huawei@2012
After you run the test-aaa command, the system displays a message indicating that the user name or password is incorrect.
<HUAWEI> test-aaa user1 huawei123 radius-template huawei
Info: User name or password is wrong.
The link between the switch and the RADIUS server is normal, but the user name or password is incorrect. Therefore, you need to change the user name or password to be the same as that on the RADIUS server. After you run the test-aaa command, the test is passed, but the user cannot pass authentication when logging in to the switch.

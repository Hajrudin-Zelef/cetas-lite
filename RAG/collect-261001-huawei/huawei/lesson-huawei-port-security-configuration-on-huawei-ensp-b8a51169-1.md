---
id: collect-261001-huawei/huawei/lesson-huawei-port-security-configuration-on-huawei-ensp-b8a51169-1
title: "lesson-huawei-port-security-configuration-on-huawei-ensp-b8a51169"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-huawei-port-security-configuration-on-huawei-ensp-b8a51169.md
source_anchor: ""
source_lines: [1, 128]
sha256: df782f6f8d3682e878e0fd89d9aff4f1e6c6385478e32a75eba6db15bfe7dd9f
---

# lesson-huawei-port-security-configuration-on-huawei-ensp-b8a51169

In this Huawei Port Security Configuration Example, we will configure Port Security on a Huawei switch. For our Huawei Port Security Configuration, we will use the below topology consist of one switch, one hub and four PCs. You can also view Cisco Port Security Configuıration Lesson.
You can download all Huawei Configuration Examples on Huawei eNSP Labs Page.
Our Huawei Port Security Configuration steps will be like below:
1. We will configure switch’s Gigabitethernet 0/0/1 interface for Port Security.
2. Firstly, we will check Protection action “Shutdown”.
3. Then, we will check Protection action “Restrict”.
4. After that, we will check Protection action “Protect”.
Now, let’s configure all these steps one by one.
Table of Contents
At the beginning, only two PCs wil be connected to the hub.
Here, firstly, we will enable Port Security under the interface GigabitEthernet0/0/1. Then, we will set the Protect Action as “Shutdown”. After that, we will set the maximum allowed MAC address number in this port. Lastly, we will enable dynamic MAC learning on the interface with “sticky” keyword for this Huawei Port Security Configuration example.
<Huawei>system-view
[Huawei] interface GigabitEthernet0/0/1
[Huawei-GigabitEthernet0/0/1] port-security enable
[Huawei-GigabitEthernet0/0/1] port-security protect-action shutdown
[Huawei-GigabitEthernet0/0/1] port-security max-mac-num 3
[Huawei-GigabitEthernet0/0/1] port-security mac-address sticky
Now let’s check our interface GigabitEthernet0/0/1 with “display interface brief” command.
<Huawei>display interface brief
PHY: Physical
*down: administratively down
(l): loopback
(s): spoofing
(b): BFD down
(e): ETHOAM down
(dl): DLDP down
(d): Dampening Suppressed
InUti/OutUti: input utility/output utility
Interface PHY Protocol InUti OutUti inErrors outErrors
GigabitEthernet0/0/1 up up 0% 0% 0 0
GigabitEthernet0/0/2 *down down 0% 0% 0 0
GigabitEthernet0/0/3 down down 0% 0% 0 0
GigabitEthernet0/0/4 down down 0% 0% 0 0
GigabitEthernet0/0/5 down down 0% 0% 0 0
GigabitEthernet0/0/6 down down 0% 0% 0 0
GigabitEthernet0/0/7 down down 0% 0% 0 0
GigabitEthernet0/0/8 down down 0% 0% 0 0
GigabitEthernet0/0/9 down down 0% 0% 0 0
GigabitEthernet0/0/10 down down 0% 0% 0 0
GigabitEthernet0/0/11 down down 0% 0% 0 0
GigabitEthernet0/0/12 down down 0% 0% 0 0
....
Then, we will send a ping from PC1 to PC2. With this ping, these two PC’s MAC addresses will be recorded to the MAC Table of the switch.
We can check switch MAC Table, with “display mac-address” command.
<Huawei>display mac-address
MAC address table of slot 0:
-------------------------------------------------------------------------------
MAC Address VLAN/ PEVLAN CEVLAN Port Type LSP/LSR-ID
VSI/SI MAC-Tunnel
-------------------------------------------------------------------------------
5489-988a-1211 1 - - GE0/0/1 sticky -
5489-987d-46bb 1 - - GE0/0/1 sticky -
-------------------------------------------------------------------------------
Total matching items on slot 0 displayed = 2
As you can see above, MAC addresses of these two connected PCs are in the switch’s MAC table.
Now, let’s connect the third PC, PC3 to the switch. After this connection, when we send a ping from PC3 to PC1, PC3 ‘s MAC address will be recored to the swicth’s MAC Table. Because, our Port Security maximum allowed MAC address on this port is 3 and still it is not violated.
<Huawei>display mac-address
MAC address table of slot 0:
-------------------------------------------------------------------------------
MAC Address VLAN/ PEVLAN CEVLAN Port Type LSP/LSR-ID
VSI/SI MAC-Tunnel
-------------------------------------------------------------------------------
5489-988a-1211 1 - - GE0/0/1 sticky -
5489-9848-4550 1 - - GE0/0/1 sticky -
5489-987d-46bb 1 - - GE0/0/1 sticky -
-------------------------------------------------------------------------------
Total matching items on slot 0 displayed = 3
Now, it is time to action! Let’s connect PC4 to the swicth. After connecting PC4 to the switch, when we send a ping from PC4 to one of the other PCs, for example PC1, our port wil go down. This is because our Port Security maximum allowed MAC address number is 3 and with the fourth PC, it is violated. And our violation action is “Shutdown”. In other words, we are saying to switch that, “if more than 3 MAC address tries to connect this port, go down!”
And on the switch screen, we will see the below violation error for this case. This error says that, the port has gone down becaue of the Port Security violation.
Apr 7 2019 17:03:43-08:00 Huawei L2IFPPI/4/PORTSEC_ACTION_ALARM:OID 1.3.6.1.4.1
.2011.5.25.42.2.1.7.6 The number of MAC address on interface (6/6) GigabitEthern
et0/0/1 reaches the limit, and the port status is : 3. (1:restrict;2:protect;3:s
hutdown)
Apr 7 2019 17:03:43-08:00 Huawei %%01PHY/1/PHY(l)[57]: GigabitEthernet0/0/1:
change status to down
Apr 7 2019 17:03:43-08:00 Huawei %%01IFNET/4/IF_STATE(l)[58]:Interface Vlanif1
has turned into DOWN state.
And on the MAC Table of the switch, there will be notting about PC4. It is not recored to the MAC Table of the switch.
Lastly, our interface. AS you can see below, it is down.
<Huawei>display int brief
PHY: Physical
*down: administratively down
(l): loopback
(s): spoofing
(b): BFD down
(e): ETHOAM down
(dl): DLDP down
(d): Dampening Suppressed
InUti/OutUti: input utility/output utility
Interface PHY Protocol InUti OutUti inErrors outErrors
GigabitEthernet0/0/1 *down down 0% 0% 0 0
GigabitEthernet0/0/2 *down down 0% 0% 0 0
GigabitEthernet0/0/3 down down 0% 0% 0 0
GigabitEthernet0/0/4 down down 0% 0% 0 0
GigabitEthernet0/0/5 down down 0% 0% 0 0
GigabitEthernet0/0/6 down down 0% 0% 0 0
GigabitEthernet0/0/7 down down 0% 0% 0 0
GigabitEthernet0/0/8 down down 0% 0% 0 0
GigabitEthernet0/0/9 down down 0% 0% 0 0
GigabitEthernet0/0/10 down down 0% 0% 0 0
GigabitEthernet0/0/11 down down 0% 0% 0 0
GigabitEthernet0/0/12 down down 0% 0% 0 0
GigabitEthernet0/0/13 down down 0% 0% 0 0
GigabitEthernet0/0/14 down down 0% 0% 0 0
GigabitEthernet0/0/15 down down 0% 0% 0 0
GigabitEthernet0/0/16 down down 0% 0% 0 0
...
Now, for the second case of Huawei Port Security Configuration, let’s delete the link of PC4 and “no shut” the GigabitEthernet 0/0/1.
[Huawei]interface GigabitEthernet 0/0/1
[Huawei-GigabitEthernet0/0/1]undo shutdown 
[Huawei-GigabitEthernet0/0/1]
Apr 7 2019 17:05:43-08:00 Huawei %%01PHY/1/PHY(l)[59]: GigabitEthernet0/0/1:
change status to up
Apr 7 2019 17:05:43-08:00 Huawei %%01IFNET/4/IF_STATE(l)[60]:Interface Vlanif1
has turned into UP state.
Now, let’s set our Port Security Protect Action as “Restrict”. After that, we will save our configuration with “save” command.
[Huawei-GigabitEthernet0/0/1]port-security protect-action restrict 
[Huawei-GigabitEthernet0/0/1]quit 
[Huawei]quit 
<Huawei>save
Now, let’s connect PC4 again and send a ping from PC4 to PC1.
After the ping, a violation alarm will appear on the screen again as below. This violation alarm is “restrict” alarm. The port will be still up but the packets to the PC4 will be dropped.
Apr 7 2019 17:09:07-08:00 Huawei L2IFPPI/4/PORTSEC_ACTION_ALARM:OID 1.3.6.1.4.1
.2011.5.25.42.2.1.7.6 The number of MAC address on interface (6/6) GigabitEthern
et0/0/1 reaches the limit, and the port status is : 1. (1:restrict;2:protect;3:shutdown)
 
Now, let’s delete the conection of PC4 again and set the Port Security violation as “Protect” on GigabitEthernet 0/0/1 interface.
 

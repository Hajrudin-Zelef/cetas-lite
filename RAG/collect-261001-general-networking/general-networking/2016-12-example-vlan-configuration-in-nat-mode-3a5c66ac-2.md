---
id: collect-261001-general-networking/general-networking/2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac-2
title: "2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac.md
source_anchor: ""
source_lines: [181, 389]
sha256: bbdfaca986a1e55e5bb56753faa3b0a37975479876c89d9eed6987018cb34268
---

# 2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac

**T****o add the firewall addresses – CLI**

config firewall address edit VLAN_100_Net

set type ipmask

set subnet 10.1.1.0 255.255.255.0 next

edit VLAN_200_Net set type ipmask

set subnet 10.1.2.0 255.255.255.0

end


**A****d****d the security policies**

Once you have assigned addresses to the VLANs, you need to configure security policies for them to allow valid packets to pass from one VLAN to another and to the Internet.

You can customize the Security Policy display by including some or all columns, and customize the column order onscreen. Due to this feature, security policy screenshots may not appear the same as on your screen.

If you do not want to allow all services on a VLAN, you can create a security policy for each service you want to allow. This example allows all services.


**T****o add the security policies – web-based manager**

**1****.** Go to **P****o****li****c****y & Objects > Policy > IPv4** or **P****o****li****c****y & Objects > Policy > IPv6** and select **C****r****ea****t****e New**.

**2****.** Leave the **P****o****li****c****y Type** as **F****i****r****e****w****a****l****l** and the **P****o****li****c****y Subtype** as **A****dd****r****ess**.

**3****.** Enter the following information and select **O****K**:

**I****n****c****o****m****i****n****g Interface**                   VLAN_100

**S****ou****r****c****e Address**                        VLAN_100_Net

**O****u****t****go****i****n****g Interface**                   VLAN_200

**D****es****t****i****n****a****t****i****o****n Address**                 VLAN_200_Net

**S****c****h****e****du****le**                                    Always

**S****e****r****v****i****c****e**                                       ALL

**A****c****t****i****o****n**                                         ACCEPT

**E****n****a****b****l****e NAT**                                Enable

**4****.** Select **C****r****ea****t****e New**.

**5****.** Leave the **P****o****li****c****y Type** as **F****i****r****e****w****a****l****l** and the **P****o****li****c****y Subtype** as **A****dd****r****ess**.

**6****.** Enter the following information and select **O****K**:

**I****n****c****o****m****i****n****g Interface**                   VLAN_200

**S****ou****r****c****e Address**                        VLAN_200_Net

**O****u****t****go****i****n****g Interface**                   VLAN_100

**D****es****t****i****n****a****t****i****o****n Address**                 VLAN_100_Net

**S****c****h****e****du****l****e**                                    Always

**S****e****r****v****i****c****e**                                       ALL

**A****c****t****i****o****n**                                         ACCEPT

**E****n****a****b****l****e NAT**                                Enable

**7****.** Select **C****r****ea****t****e New**.

**8****.** Leave the **P****o****li****c****y Type** as **F****i****r****e****w****a****l****l** and the **P****o****li****c****y Subtype** as **A****dd****r****ess**.

**9****.** Enter the following information and select **O****K**:

**I****n****c****o****m****i****n****g Interface**                   VLAN_100

**S****ou****r****c****e Address**                        VLAN_100_Net

**O****u****t****go****i****n****g Interface**                   external

**D****es****t****i****n****a****t****i****o****n Address**                 all

**S****c****h****e****du****l****e**                                    Always

**S****e****r****v****i****c****e**                                       ALL

**A****c****t****i****o****n**                                         ACCEPT

**E****n****a****b****l****e NAT**                                Enable

**10****.** Select **C****r****ea****t****e New**.

**11****.** Verify the **P****o****li****c****y Type** is **F****i****r****e****w****a****l****l** and the **P****o****li****c****y Subtype** is **A****dd****r****ess**.

**12****.** Enter the following information and select **O****K**:

**I****n****c****o****m****i****n****g Interface**                   VLAN_200

**S****ou****r****c****e Address**                        VLAN_200_Net

**O****u****t****go****i****n****g Interface**                   external

**D****es****t****i****n****a****t****i****o****n Address**                 all

**S****c****h****e****du****l****e**                                    Always

**S****e****r****v****i****c****e**                                       ALL

**A****c****t****i****o****n**                                         ACCEPT

**E****n****a****b****l****e NAT**                                Enable


**T****o add the security policies – CLI**

config firewall policy or Config firewall policy6 edit 1

set srcintf VLAN_100

set srcaddr VLAN_100_Net set dstintf VLAN_200

set dstaddr VLAN_200_Net set schedule always

set service ALL set action accept set nat enable

set status enable next

edit 2

set srcintf VLAN_200

set srcaddr VLAN_200_Net set dstintf VLAN_100

set dstaddr VLAN_100_Net set schedule always

set service ALL set action accept set nat enable

set status enable next

edit 3

set srcintf VLAN_100

set srcaddr VLAN_100_Net set dstintf external

set dstaddr all

set schedule always set service ALL

set action accept set nat enable

set status enable next

edit 4

set srcintf VLAN_200

set srcaddr VLAN_200_Net set dstintf external

set dstaddr all

set schedule always set service ALL

set action accept set nat enable

set status enable

end


**C****on****f****i****gu****r****e the VLAN switch**

On the Cisco Catalyst 2950 Catalyst VLAN switch, you need to define VLANs 100 and 200 in the VLAN database, and then add a configuration file to define the VLAN subinterfaces and the 802.1Q trunk interface.

One method to configure a Cisco switch is to connect over a serial connection to the console port on the switch, and enter the commands at the CLI. Another method is to designate one interface on the switch as the management interface and use a web browser to connect to the switch’s graphical interface. For details on connecting and configuring your Cisco switch, refer to the installation and configuration manuals for the switch.

The switch used in this example is a Cisco Catalyst 2950 switch. The commands used are IOS commands. Refer to the switch manual for help with these commands.


**T****o configure the VLAN subinterfaces and the trunk interfaces**

Add this file to the Cisco switch:

!

interface FastEthernet0/3 switchport access vlan 100

!

interface FastEthernet0/9 switchport access vlan 200

!

interface FastEthernet0/24

switchport trunk encapsulation dot1q switchport mode trunk

!

The switch has the configuration:

**P****o****r****t 0/3**                         VLAN ID 100

**P****o****r****t 0/9**                         VLAN ID 200

**P****o****r****t 0/24**                       802.1Q trunk


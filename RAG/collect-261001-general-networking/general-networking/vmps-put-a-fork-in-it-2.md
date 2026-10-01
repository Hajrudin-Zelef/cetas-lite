---
id: collect-261001-general-networking/general-networking/vmps-put-a-fork-in-it-2
title: "vmps-put-a-fork-in-it"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-general-networking/vmps-put-a-fork-in-it.md
source_anchor: ""
source_lines: [51, 79]
sha256: 3f2b1333d1d3f65b91208c17f906c2d364b87477dc2544fe4534b9bf74f645d9
---

# vmps-put-a-fork-in-it

Either method is easily defeated by knowledgeable attackers. Using a static IP address will bypass DHCP lease management handily. ARP poisoning is a bit stronger, but on Windows hosts, using the built-in “arp -a” command will create a static ARP mapping. The tricky part is getting the network peer–a router, for example–to know what your real MAC address is. Constantly sending out directed ARP responses is one solution. Therefore, while DHCP lease management and ARP poisoning have their uses as interim enforcement methods during a NAC pilot, the better enforcement methods such as VLAN steering or 802.1X should be preferred. Exceptions to this rule would be those cases where nothing else works well, such as when the infrastructure is unmanaged or it’s too costly to deploy in-line enforcement.

Cisco Layer-3 switches have DHCP snooping and Dynamic ARP Inspection features that can help keep tabs on these LAN abuses. I often recommend these to my customers that these features be enabled even if a NAC solution is not being deployed.

The best NAC approaches perform the assessment of the end-point and check that state before the client is allowed to connect to the network. The stronger NAC solutions provide integration with the LAN switching infrastructure and use 802.1X to achieve this tight control of the edge of the network. However, when organizations have hubs it makes 802.1X infrastructure solutions difficult to deploy. These hubs will need to be removed in order to be able to deploy a NAC solution.

NAC solutions that use 802.1X don’t typically have a lot of remediation capabilities and therefore must be combined with other techniques for remediation. IEEE 802.1X is merely just an access control methodology. Worst-case solutions require manual fixing of computers once on the guest VLAN. The better solutions can automate the remediation of the hosts.

NAC systems that are in-band can potentially suffer from the large amount of bandwidth coming from Gigabit Ethernet desktops. Some of the in-band NAC solutions use high performance ASICs and some even have 10GE interfaces for connection between distribution and core network switches. Since many organization’s network today consist of mostly 10/100/1000 access ports bandwidth is of less concern. Some of the in-band solutions operate with IEEE 802.1Q on the NAC layer-2 solutions. This could be a work solution because the aggregated bandwidth for an in-band solution would be acceptable.

While there are several solutions exist as viable alternatives for VMPS the most popular replacement options for VMPS is for organizations to use Identity Based Network Services (IBNS) 802.1X MAC Authentication Bypass (MAB) or a more complete NAC solution like Cisco’s NAC Appliance. Even Cisco recommends VMPS and Cisco Secure User Registration Tool (CSURT) users migrate to 802.1X MAB.

802.1X MAB allows devices without 802.1X supplicants that connect to the access switch to have their MAC address used in the RADIUS authentication request. The RADIUS server looks up the MAC address in its database and determines if the MAC address is allowed and what VLAN it should be assigned. The RADIUS server responds to the authentication request by the access switch with that information to allow the computer access to the network. The trick is getting the MAC addresses and VLAN mapping into the RADIUS server or into an LDAP database for the RADIUS server to query.

If you are curious here is the link to the Cisco 6500 CatOS 8.7 configuration guide for 802.1X. Furthermore, here is the Cisco 6500 Cat IOS 12.2(SXH) configuration guide for 802.1X.

Here is a sample configuration for an interface on a Cat IOS switch that has been placed into 802.1X MAB mode. interface FastEthernet1/6 switchport access vlan 20 switchport mode access dot1x mac-auth-bypass dot1x pae authenticator dot1x port-control auto spanning-tree portfast spanning-tree bpduguard enable

The configuration commands have changed to a new syntax using the “authentication” command in newer versions of 12.2 on 6500s. Here is an example of a newer MAB configuration for Cat IOS 12.2(33)SXI.

Router(config)# interface FastEthernet 1/6 Router(config-if)# authentication port-control auto Router(config-if)# authentication event no-response action authorize vlan 20 Router(config-if)# dot1x pae authenticator Router(config-if)# dot1x timeout supp-timeout 3 Router(config-if)# dot1x timeout tx-period 15 Router(config-if)# dot1x pae authenticator Router(config-if)# mab [eap]

Here are some useful Cisco IOS show commands for determining if your MAB configuration is working as you expect. show dot1x all [summary] show dot1x interface show dot1x interface fastethernet1/6 details show authentication [ registrations | interface |method ] show authentication sessions [handle handle] [interface interface] [mac mac] [method method] [session-id session-id] show vlan group [group-name group-name] clear dot1x interface fastethernet 1/6 clear authentication sessions interface fastethernet 1/6 method mab

To learn more about the newest features of 802.1X in Cat IOS 12.2(33)SXI check out Jamey Heary’s blog on the topic.

If you are a current user of VMPS I wish you all the best of luck migrating to 802.1X (with or without MAB) or a NAC appliance solution.

Scott

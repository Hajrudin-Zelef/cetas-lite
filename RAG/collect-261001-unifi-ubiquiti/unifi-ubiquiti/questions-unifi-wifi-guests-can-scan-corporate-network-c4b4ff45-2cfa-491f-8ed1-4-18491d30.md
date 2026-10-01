---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-unifi-wifi-guests-can-scan-corporate-network-c4b4ff45-2cfa-491f-8ed1-4-18491d30
title: "questions-unifi-wifi-guests-can-scan-corporate-network-c4b4ff45-2cfa-491f-8ed1-4-18491d30"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-unifi-wifi-guests-can-scan-corporate-network-c4b4ff45-2cfa-491f-8ed1-4-18491d30.md
source_anchor: ""
source_lines: [1, 34]
sha256: 98c6ce333d8496612e0425d1d27053ed073b7ac1e3438535e564720ca8d7ffb7
---

# questions-unifi-wifi-guests-can-scan-corporate-network-c4b4ff45-2cfa-491f-8ed1-4-18491d30

@UI-Team
Hello, I have a problem with securing my corporate network against scanning by wifi clients. Unifi settings are - apply guest policy, no guest portal, restrict network 10.0.0.0/8 and still I can scan my corp. network. I´m using zAnti for Android
www.zimperium.com/products.php
and I can see all computers with their IP and MAC addresses, hostnames and all open ports on DHCP server. Wifi clients are isolated, shared folders are unavailable, as well as open ports on other PCs. Maybe I have overlooked something in Unifi settings, but I would like to separate wifi from wire network.
Thank you for your help.
Standa
PS. sorry for my english
That is to be expected and has been reported here often. If you want better isolation setup VLANs.
youre correct... unifi is isolated and restricted. the ap doesnt restrict ICMP protocol.
thats why im anticipating 3.0.0, give the public there own dhcp range and be done with it.
About VLAN.. can you post link to some guide please or shortly describe setting procedure? I am not very experienced in VLAN. I have only one switch with vlan support. Do I need all of them with VLAN?
Thank you very much.
This link will help you with your vlan setup
wiki.ubnt.com/UniFi_and_switch_VLAN_configuration
yes, vlans are the way to go. 
While VLANs will solve a lot of issues, for MY application it is completely useless, since I cannot install extra equipment on all the networks where I want Unifi APs. The way it has been advertised to work, with whatever additional filtering necessary is the way to go. What I see as an advantage is that the APs do NOT need to be on a single corporate network, but can adapt to many different networks.
So my vote is for "better isolation", good enough to convince a network owner.
thats why im waiting for 3.0.0
a colleague has the cisco wifi controller... easy work around is to assign the guest ssid with its own WAP DHCP and let employees to pass straight to the network. thats what he does. hes kind of in the same boat about adding more hardware.
Ja, that's how we do it. Cisco can tunnel the guest WLAN through without VLANs.
Standa80:I have a problem with securing my corporate network against scanning by wifi clients. 
Your Corporate and WiFi network should NOT be on the same network !!!!!!
Your corporate network should be behind a NAT router and firewall, and the WiFi can be in front of that. That way NO ONE outside the firewall can look into the Corporate network.
If you need a wireless access for just your employees, put in a separate wireless AP, JUST FOR THEM. (Secured and encrypted)
You do not want some hacker pulling into your parking lot, and playing around getting all your corporate information or your guests personal and credit info !!!
KEEP THE TWO NETWORKS SEPARATE !!!
It would be very difficult to give general VLAN setup advice, so just take some time, dig around, and learn how to setup a VLAN for your specific hardware- you'll be glad you did.
In my single experience with UniFi hardware, I ended up forgoing (turning off) ubnt's guest feature-set. I did tag the traffic using the SSID VLAN ID setting, and then deferred the rest of the guest SSID's isolation to a Mikrotik firewall using a VLAN and a couple extra firewall rules. It works perfectly.
.....Your corporate network should be behind a NAT router and firewall, and the WiFi can be in front of that. That way NO ONE outside the firewall can look into the Corporate network.... 
So VLANing a network wouldn't/doesn't cut the mustard?
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

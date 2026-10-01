---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-how-to-enable-intervlan-routing-on-the-unifi-usg-cf4a4213-6c6c-4002-8d-8b85c77d
title: "questions-how-to-enable-intervlan-routing-on-the-unifi-usg-cf4a4213-6c6c-4002-8d-8b85c77d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-how-to-enable-intervlan-routing-on-the-unifi-usg-cf4a4213-6c6c-4002-8d-8b85c77d.md
source_anchor: ""
source_lines: [1, 25]
sha256: f7154e9dcc03d59a80fa6cc1bcdf2823255bcf4ddbea58706b188bf4e30efc78
---

# questions-how-to-enable-intervlan-routing-on-the-unifi-usg-cf4a4213-6c6c-4002-8d-8b85c77d

@UI-Team
Hi All,
How do we enable interVLAN routing on the USG 3P?
What I have is two vlans coming out the lan 1 to a managed switch and two devices on the switch on different vlans. Pretty darn sure the tagging on my switch is correct because I have interVLAN routing working on Juniper and Cisco routers and firewalls. The USG can ping each device but the devices can not ping each other.
For resolutions, I have tried all default firewall rules with only the vlans created, and I have tried opening up the network with allow rules on the lan in and out. Please help!
Thank you
Ínter vlan routing is enabled as default, as long as both networks are set to corporate type networks. Is one of yours guest? If not I would delete all your custom firewall rules and reboot the USG to see if something was stuck in wrong config.
If not, you can post screensshots of your networks and firewall for us to look for issues. Remove all personal info if you do so.
I've got the same issue, just tried setting this up and cannot communicate between VLANs. All are set as corporate, I have NO firewall rules defined other than the default ones that USG creates when adding the VLAN on LAN-IN/OUT
VLAN1 - 192.168.10.1/24
VLAN50 - 192.168.50.1/24
VLAN100 - 192.168.100.1/24
I have a UniFi switch that I set ports 3 and 4 for VLAN 50 and 100, all other ports are default as ANY
Devices get an IP address in the appropriate range depending on what switch port they are plugged in to.
Devices on the 10 network cannot ping any device on either VLAN 50 or 100.
Devices in the 50 network can ping the 10 network, but not the 100 network
Devices in the 100 network can ping the 10 network but not the 50 network.
All devices can ping the x.x.x.1 gateways for each x.x.10.1, x.x.50.1 and x.x.100.1 subnets
I've been going around in circles with this, my understanding is that this SHOULD work out of the box, then you set rules to deny it, however I'm not sure if my understanding is wrong, or there is a bug in the software.
Sooooo, worked out what my problem was, Windows Firewall!
The three devices I was using on each VLAN were Win10 machines, and in each case, disabling the firewall on them let me ping them. Far out, was tearing my hair out thinking UniFi was the problem... 😀
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

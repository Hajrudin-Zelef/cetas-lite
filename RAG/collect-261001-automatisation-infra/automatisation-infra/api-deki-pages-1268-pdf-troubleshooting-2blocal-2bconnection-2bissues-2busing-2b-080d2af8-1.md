---
id: collect-261001-automatisation-infra/automatisation-infra/api-deki-pages-1268-pdf-troubleshooting-2blocal-2bconnection-2bissues-2busing-2b-080d2af8-1
title: "api-deki-pages-1268-pdf-troubleshooting-2blocal-2bconnection-2bissues-2busing-2b-080d2af8"
domain: automatisation-infra
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-automatisation-infra/api-deki-pages-1268-pdf-troubleshooting-2blocal-2bconnection-2bissues-2busing-2b-080d2af8.md
source_anchor: ""
source_lines: [1, 93]
sha256: 522ca4c88b94e02da59baabb2fa6d3d378759405f1b87d567df77d682e7f2969
---

# api-deki-pages-1268-pdf-troubleshooting-2blocal-2bconnection-2bissues-2busing-2b-080d2af8

Troubleshooting Meraki AP Cloud and Gateway Connectivity
Click日本語 for Japanese
Overview
This articleexplains how to diagnose and resolve common connectivity issuesencounteredwith Cisco Meraki MR access points (APs). It provides step-by-step
instructions toidentifyand fix problems that prevent an AP from contacting the Meraki Cloud,operatingas a gateway, or being configured locally.
The guide covers four related issue categories:
• An AP reporting difficulty contacting the Meraki Cloud.
• A gateway APswitchingto repeater mode.
• A gateway APfailingto rejoin as a gateway after reconnection.
• Connecting to a Meraki AP locally using its default SSID when it cannot reach the cloud.
Environment
• Hardware: Cisco Meraki MR series access points(including MR46 and other Wi-Fi 6 and newer APs)
• Network services: DHCP, DNS, ARP, outboundfirewallaccess to the Meraki Cloud
• Local configuration tools: ap.meraki.com andmy.meraki.com local status pages
Troubleshooting Meraki Cloud connectivity issue
An alert may appear on your AP detail pagestating:
"This device is having difficulty contacting the Meraki Cloud. Please make sure your wired network allows outgoing connections tox.x.x.xandx.x.x.xon ports
443, 7734, 7351 and 7752."
When this happens, the AP icon (locatedin dashboardunderWireless > Monitor > Access Points) turns yellow, the connectivity graph stays green, and the
AP does not download the latest firmware or configuration from the Meraki Cloud.
Possible causes
• The AP sits behinda firewallthat blocks outbound access to the Meraki Cloud.
• The DNS servers the AP uses are unreachable, do not respond, or send invalid DNS responses.
• The upstream ISP-provided modem has traffic inspection security enabled.
• Dashboardis experiencing a temporary outage.
1

Troubleshooting steps
1. Check yourfirewalland confirm itallowsoutbound access to the cloudon ports443,7734,7351, and7752.
2. If thefirewallsettings are correct, investigate DNSand perform following steps:
• Change the DNS servers used by your AP to a public server (such as Google Public DNS)
• Confirm yourfirewallallows outbound DNS traffic (UDP port53). If your AP uses a static IP address, refer to
thestatic DNS settingsdocumentation.
3. If the issuepersistsperform following steps:
• Enabletraffic inspection securitysettingon your ISP-supplied modem or router.
• Rebootthe modem or router, mostlyfixesthe issue temporarily.
• Comcast modems and SMC-manufactured routers ship with theGateway Smart Packet Detection feature
enabled by default, which is known to cause this condition. Contact your ISP or modem/router
manufacturerforhelp disabling thissecurity feature.
Expected outcome
The AP iconreturns togreen, the AP downloads the latest firmware and configuration, and the alert no longer appears on the AP detail page.
Troubleshooting gateway AP repeater-mode issue
A gateway AP is an access point with a wired interface configured with an IP address, connected to the LAN, and with a route to the Internet.
If the Internet is unavailable and the SSID allows LAN access, the AP continues to act as a gateway because it still holds a valid IP and can reach the local
router orfirewall.
Possible causes
The AP converts to a repeater only when one of these conditions is true:
• The AP cannot receive an ARP reply packet from the default gateway on the LAN (usually a localfirewallor router).
• The AP cannot obtain a valid IP address via DHCP.
If the AP uses a static IP, it advertises"<ssid name>-bad gateway" when the default gateway is unavailable. With no other APs to mesh with that provide a
route to the Internet, the APremainsoffline. If it finds another AP with a route to the Internet, it acts as a repeater, but the dashboard reports an invalid IP
configuration.
Troubleshooting steps
If an AP joinsthedashboard as a repeater but you expect it to be a gateway, complete these verifications:
1. Connect a laptop withWiresharkinstalled to the switch port where the AP connects. Confirm the laptop receives an IP by DHCP and can ping its
gateway.
2. If the laptop does not receive an IP by DHCP, troubleshoot until it successfully gets an IP address. Then reconnect the AP; it should jointhe dashboard
as a gateway.
2

3. If the laptop receives an IP by DHCP butdoesnotreceiveARP replies from its default gateway, troubleshoot the default gateway accordingly.
4. If the AP needs a static IP address because no DHCP server is available on site, uselocal status pageto configure it.
Expected outcome
The AP obtains a valid IP and reaches its default gateway, then joins dashboard as a gateway rather than a repeater.
Troubleshooting gateway AP reconnection issue
All Meraki access points dynamicallymonitortheir uplink port for Ethernet connectivity. In a few scenarios, an AP that was once a gateway will not become a
gateway again after youreconnect its Ethernet cable. When this happens, the AP signal LEDs scan back and forth, and an SSID appended with"-
scanning" may appear in your wireless network list.
Possible causes
• A bad cable run or a cable run that is too long.
• A faulty PoE injector.
• The AP has no IP address
• A Layer 1, 2, or 3problemon the switch port.
• The switch port is bad or administratively shut down.
Troubleshooting steps
1. Check the cable run:
1. Verifythe cable is securely connected on both ends.
2. Replace the current cablerun, andconfirmthe total run is100mor less (a physical limitation of CAT5/
Ethernet cables).
2. Check the PoE injector. Swap the PoE injector in use with a known good PoE injector.
3. Check if AP is missing the IP address. If the AP is set to obtain an IP address automatically, the DHCP server may not be responding, may be
unreachable, or may be out of IP addresses.
1. Verifythe DHCP server is running and reachable.
2. Verifythe DHCP pool has addresses available for lease.
3. Configure the AP with a static IP address to see if it becomes a gateway; successindicatesa problem with
the DHCP service on the LAN.
4. Verify Layer 1, 2, and 3 on the switch port using a laptop.
1. Disconnect the AP from the switch port.
2. Plug the laptop into the same switchportthe AP used.
3. Verify the laptop obtains a DHCP address and can ping hosts on the Internet.
5. Check the switch port. If the port is bad or administratively shut down, connect the AP to a different port on the switch.
3


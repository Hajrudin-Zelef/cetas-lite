---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-l2tp-ipsec-vpn-on-udm-not-working-3a50de8d-39d1-47e1-a5df-84690aa0bfbf-45b3686a
title: "questions-l2tp-ipsec-vpn-on-udm-not-working-3a50de8d-39d1-47e1-a5df-84690aa0bfbf-45b3686a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-l2tp-ipsec-vpn-on-udm-not-working-3a50de8d-39d1-47e1-a5df-84690aa0bfbf-45b3686a.md
source_anchor: ""
source_lines: [1, 34]
sha256: 68ab19c5140b2cddee95f440ccd4e46d6025cc9811daf821747992b73921b728
---

# questions-l2tp-ipsec-vpn-on-udm-not-working-3a50de8d-39d1-47e1-a5df-84690aa0bfbf-45b3686a

@UI-Team
Hey Guys,
my Dream Machine is Running UnIFI OS Version 1.12.38 and Network Version 7.2.97.
When trying to set up my L2TP VPN im running into some issue:
Here are the firewall rules that were automatically generated:
I would appreciate some help. I tried changing the release channel to update to the relase candidate version but that didnt seem to actually work either.
Kind of off topic, but also:
When trying to enable the "teleport" feature and generating an invitation link i get the following error: "Remote access disabled. Enable your remote access and try again." even though the feature is already enabled....
I really hope you can help me out here!
Thanks guys!
Hi, have you tried to manually set your WAN1 IP address with your Static IP Address from the ISP, for the Unifi Gateway IP?
Is Windows 10/11 fully updated. Also make sure MS Chap is enabled on the Windows machines.
From devices manager try deleting all the mini port drivers and let them re-add themselves after a scan for new hardware or a reboot.
@Nick.Franklin wrote:
Hey Nick!
Thanks for commenting but i already tried that, i got a different error but in the end it didnt work either :(
I just tried connecting via MacOS and that also didnt work (L2TP Server not responding)
I have now tested the whole thing again with a Linux Mint Machine. Same problem... The VPN connection to my UDM simply cannot be established.
@Nettcode wrote:
What is your connection to the internet from the UDM. Do you have another router from your ISP, FTTP with an ONT etc. I presume you are not trying to connect from to your VPN from your internal network.
Welcome to my world: https://community.ui.com/questions/L2TP-VPN-stopped-working-at-an-unknown-point-in-the-past-now-showing-NOPROPOSALCHOSEN/5bc91061-d6aa-4584-a4cd-f0fd2e4ad865
Though sadly I never received a feedback from anyone at UI.
Hey @Nick.Franklin
From the ONT, a LAN cable goes directly into my UDM. And on the UDM, my Internet access is established with PPPoE credentials. Afterwards I get my static IP address assigned and am online.
I am currently trying out VPN access from work. Cell phone hotspot or generally other locations with other internet providers do not work either.
OK, does your ISP allow VPN connections. Some don't. Worth checking.
Yes, 100%. I had a UDR before my UDM. And with the UDR I could use the WireGuard VPN without any problems.
I had to switch to the UDM because the UDR has a too weak processor for my internet connection.
I have a UDM currently running the latest EA release 2.4.27 having upgraded from 1.12.x version. Each time the VPN has worked and continues to work. Your rules look good as well and compare to mine on my UDM.
On your VPN configuration though I see you have it set to Auto. Set it to manual. define the User Access List, gateway and subnet. Add in your preferred Name Servers if required but enable Require Strong Authentication and Weak Ciphers and test again.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

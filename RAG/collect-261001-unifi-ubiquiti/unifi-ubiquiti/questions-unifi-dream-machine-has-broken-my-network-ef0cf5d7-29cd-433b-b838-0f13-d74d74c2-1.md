---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-unifi-dream-machine-has-broken-my-network-ef0cf5d7-29cd-433b-b838-0f13-d74d74c2-1
title: "questions-unifi-dream-machine-has-broken-my-network-ef0cf5d7-29cd-433b-b838-0f13-d74d74c2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-unifi-dream-machine-has-broken-my-network-ef0cf5d7-29cd-433b-b838-0f13-d74d74c2.md
source_anchor: ""
source_lines: [1, 35]
sha256: bb25d36e0b810799fed9f74921413df086d05500c3eb92f2c7eea77a3a3195c1
---

# questions-unifi-dream-machine-has-broken-my-network-ef0cf5d7-29cd-433b-b838-0f13-d74d74c2

@UI-Team
I have a reliable SOHO network and wanted to upgrade. I chose the Unifi and the Dream Machine seemed a great choice for our simple needs. It has never worked reliably; drops connections, either completely or temporarily, both through its own SSID and APs with wired connections; won't reliably allow control panel connection etc.
Having spent months with support, I RMAd and swapped in a new device. Same results.
I am following completely standard setup.
What am I missing? Might it be the VDSL modem?? or another reason that two different UDMs fail to function reliably? (i'm not even talking advanced features... just reliable routing).
Many thanks for your help (I can't face the black hole of "support" again)
Without a list of hardware/firmware details along with the exact issues and steps you have taken so far, we can't help.
Thanks gregorio
I have FTTC connected to a Technicolor TG588v2, configured as a modem and connected to the UDM's WAN port. It is the "pill bottle" UDM, not the Pro, running firmware v.1.5.6. Before connecting these I used a FRITZ!Box 7530 and the LAN and WAN worked fine. Other devices in the LAN include a TP-LINK RE450 configured as an AP and Netgear GS308 un-managed switch.
All the faults are intermittent. The most common one is failing to provide any WAN connection e.g. web pages get a browser error saying the server's DNS can't be found, before moving to an error saying there is no internet connection. This happens a little more often with the UDM's built-in AP than with the other APs (all have a wired connection back to the UDM). Line speed tests are also slower; previously I reliably got 72Mbps down and 20Mbps up, now usually ~50 and ~13 respectively. We experience these faults across multiple client devices spanning various OSs (Windows, Chrome, iOS, Android, other smart devices e.g. Sky TV set top box) and many manufacturers.
With the previous UDM I have re-booted, re-set and gone through multiple firmwares; adjusted radio frequency and band widths; ensured smart queues were off, all following instructions from support. I can go through the support logs for the specifics if you think that will be fruitful.
Again, many thanks for your help. I'm certainly no network engineer, but creating a vanilla setup on a SOHO device is normally well within my capabilities, and I'd really love to get this working (next step: get the Unifi PoE switch in, and then their APs).
Well, I am certainly sympathetic. Failures like this are frustrating. I would break the problems into two categories; stability and performance. The first is going to be the hardest to nail down but I feel the most important.
Hard to know exactly what is happening. Browser connection messages are pretty generic and can be caused by any number of issues. Is the ISP link dropping? Is there any indication in the Technicolor device? Could there be a configuration problem? Is the UDM not routing? Is the problem restricted to wired or wireless devices or both? Does it happen to all devices at once?
Have you tried the new RC firmware 1.7.0, 1.7.1, 1.7.2?
Smart queues is probably enabled since your internet speed is less than 300mbps. That would explain your slightly reduced speeds. Overall with your isp speed smart queues is ideal to keep the internet flowing evenly. As far as the DNS issue have your run a DNS benchmark test and figured out what server is best for your location? Doing that and changing your dns server could fix your issue. Is your isp modem/router in bridge mode? Perhaps it’s a double nat issue as well.
Have you tried taking the Technicolor modem out of the equation altogether and just used the UDM on it's own?
Thanks for the suggestions - I really appreciate it, nice to not feel alone with this! Responding to those that I can now; I've had to take the UDM out of the network today (my wife and I are both running our businesses from home), so will attempt changes this evening.
Again, thank you - guidance gratefully received, this is starting to drive me nuts!
With my UDM (in Australia) I was provided with a TP-Link Archer VR1600v modem and plugged the UDM into that. Had a few issues and decided to take the Archer out. The UDM worked much better without the Archer in bridge mode. Had to set up a VLAN id. Once I worked that all out ... it all worked fine and has been great ever since. I thought the UDM was a modem also. Not sure what is required with Zen wherever you are though. Good luck with it.
@Zarn73 do you know what technology is used for the OLT (at your ISP side) If it is Huawe you maybe be ok with a UFiber Nano (https://www.ui.com/ufiber/ufiber-nano-g/) I am using it here in Spain and I love it.POE and rock solid. You need to know your PLOAM password though ...
I had some reliability problems with my UDM Pro that I tracked down to it losing DHCP connection to the modem every ten minutes. This caused all internet connections (TCP) to get dropped every ten minutes. But it would restart immediately, so you'd only really notice it for downloads that took more than 10 minutes. If I bypassed the UDM Pro and connected directly to my mode, no problem.
Anyhow, what finally made it clear what was wrong and helped me fix it was tcpdump. From my UDM Pro, I ran "tcpdump -vnes0 -i eth8 port 68" and I could see that DHCP renewal wasn't working. It ended up being a config problem where my local network and the modem's local network overlapped. (I have a bad AT&T configured modem whose bridge mode is somewhat flawed).
My suggestion -- try tcpdump to debug your DNS traffic. (you'll have to figure out which ports, etc.) You can capture both on the device side and the modem side. I'll bet you'll figure out what is going on. It might take a little time, but I bet it'll give you the answer you are looking for.
PS: I would try upgrading my UDM firmware first.
Again, thank you for your help and suggestions. I'm still working through them, but wanted to also respond to some of the questions
1) My WAN connection is fibre to the cabinet, then the last ~200m is twisted copper pair; so I don't think I can use the UFiber Nano?
2) Firmware 1.5.6 (my current) is the latest listed for the UDM; 1.7.x is listed as being for the UDM Pro - any views on applying it to UDM?
Thanks!
The Technicolor modem is fine - you can't use a UFiber Nano product, they are only for GPON networks (FTTP).
Given you've switched from a Fritzbox to a Technicolor + UDM combo, it could also be an issue with the Technicolor modem.
As a first step in troubleshooting, you configure the WAN port on the UDM to be DHCP, and plug the Fritzbox back in instead of the Technicolor. This will give you double NAT temporarily, but at least it will allow you to test with a known good "modem". If you are still having issues at that point, it's definitely UDM related, and getting onto the latest firmware would be the best first step.
*EDIT* Latest firmware is 1.8.x - currently in "RC" I think.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.

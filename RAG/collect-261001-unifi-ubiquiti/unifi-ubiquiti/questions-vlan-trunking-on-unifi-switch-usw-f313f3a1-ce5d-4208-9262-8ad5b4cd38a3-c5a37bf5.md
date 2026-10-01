---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-vlan-trunking-on-unifi-switch-usw-f313f3a1-ce5d-4208-9262-8ad5b4cd38a3-c5a37bf5
title: "questions-vlan-trunking-on-unifi-switch-usw-f313f3a1-ce5d-4208-9262-8ad5b4cd38a3-c5a37bf5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-vlan-trunking-on-unifi-switch-usw-f313f3a1-ce5d-4208-9262-8ad5b4cd38a3-c5a37bf5.md
source_anchor: ""
source_lines: [1, 24]
sha256: c434c9c09f165e4f9d521793f510cf8dfa2d44d6694a0041dcbf38161130b53f
---

# questions-vlan-trunking-on-unifi-switch-usw-f313f3a1-ce5d-4208-9262-8ad5b4cd38a3-c5a37bf5

@UI-Team
This is probably obvious/well documented, but caused me a few days of pain, so thought I would share;
When setting up VLANs and using multiple switches you need to ensure the port on BOTH ends of the "uplink" are configured for all VLANS (or at least all the VLANs you need available).
I have a switch in building A linked to a switch in building B linked to a switch in building C. I have 2 VLANs, one for normal traffic, then another for my HDMI video senders/receivers.
Till last week all ports on switches B and C were tagged with the "normal traffic" VLAN only. The HDMI senders/receivers were all connected to switch A where the relevant individual ports were tagged with the HDMI VLAN. Everything was working nicely.
This week I wanted to add a HDMI receiver on switch C. I set the port on switch A (which links to switch B) to all VLANs and set the port on switch B (which links to switch C) to all VLANs, then finally set the relevant port on switch C (with the HDMI device plugged in) to the HDMI VLAN. Unfortunately it didn't work.
Eventually I figured out that the ports on both ends of each uplink/downlink needed to be set to all VLANs.
I did a bit of a facepalm, but hope this post may help someone else in future!
As a matter of good design practice, I ensure that switch port profiles to meet all use-cases are created, and applied without fail to all switch ports.
The "ALL" option is never used for trunk switch ports, all VLANs are explicitly declared. Whilst LAN1 exists, it is not used. All connected devices are classified and use an appropriate VLAN.
Ports connecting switches to APs are trunk ports as they pass the management VLAN, and the associated VLANs for each of the SSIDs being broadcast.
When classifying user/device types, be sure to keep the goats from the sheep. Goat Guest networks should isolate guests from the rest of the network, and other guests.
Don't let anything "smart" share a VLAN with unrelated smart devices.
Manage video and voice traffic on their own VLANs.
Disable PoE on ports which do not require it, and disable ports which are not in use.
Proactive port management using profiles helps maintain a predictable and reliable network.
And yet they ship them with ALL set to every port.... First thing I do is change that.....
@gregorio wrote:
You guys are better than most. I have to admit, ALL is so much easier. 😁
Just don't do that in the convention world.... When you deal with 1,000s of guests, especially if you have a hacking convention in house - you better not miss something or they will show you 😡 😂 Thankfully I don't have to deal with the operational side of that anymore...👍
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

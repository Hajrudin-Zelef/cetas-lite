---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-unifi-dream-machine-pro-se-internal-bottleneck-01a00a97-0414-48f4-9478-89bce008
title: "questions-unifi-dream-machine-pro-se-internal-bottleneck-01a00a97-0414-48f4-9478-89bce008"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-unifi-dream-machine-pro-se-internal-bottleneck-01a00a97-0414-48f4-9478-89bce008.md
source_anchor: ""
source_lines: [1, 31]
sha256: 6d360dc60c06e605ab7680fe1290c3d2179e5b80f28e37ee99e1548e88a2c30f
---

# questions-unifi-dream-machine-pro-se-internal-bottleneck-01a00a97-0414-48f4-9478-89bce008

@UI-Team
Hello Everyone,
Could someone tell me if there more recent Unifi Dream Machine doe still have the internal 1Gbit bottleneck?
Depends on your definition of bottleneck. Technically this is a normal switch with each port supporting full duplex 1Gbps connectivity with also a 1Gbps uplink to the UDM`s CPU. Any other low end (no Enterprise or Pro) Unifi switch would offer the same, nothing more and nothing less in terms of capacity. In terms of functions there is some L2 functionality missing like Spanning Tree & Link Aggregation; not sure why that is not supported.
@RobbieH wrote:
He's talking about the fact there's only a 1gbps to the backplane CPU for all ports, therefore causing a bottleneck.
So what is the difference to any other lower end L2 Unifi switch ?
They all have this 1Gbps upload to the CPU of the UDM except for the higher end L3 Enterprise or Pro switches that have 10Gbps via SFP+ and completely different price levels......they are more expensive than the UDM`s themselves.
Just to clarify to who ever cares.
On the UDMP SE you have the 1G ports on the left. They between each other are line rate - you have an 8 port line rate switch. If something on those ports wants to go to WAN or LAN ports they are limited to a 1G pipe. So if all 8 ports are trying to talk to the WAN at 1G at the same time you are trying to shove 8G in to a 1G pipe and things drop.
The WAN/LAN ports are line rate to each other for the most part. If your goal is to have things on your network speak line rate to each other then:
UDMP-SE - LAN - Linerate switch with your clients on it.
If you are are going between VLANS then you are routing via the UDM and you are limited to whatever the configuration is on the UDM - which is around 3.5G.
If you have the enterprise switches you can move the routing to the switch and that should be linerate between the SVIs on the switch.
@sahafeez wrote:
Long text to say what I am saying......the switch part of the UDM`s has the same behavior as any other Unifi switch except for the more expensive Pro & Enterprise switches. On switch level this switch can handle full duplex 1Gbps on any of the 8 ports meaning 16Gbps switching capacity. Now we`re not talking of the lack of functionality like STP & LACP.
No, that's the problem. The UDM line has only a 1gbps backplane for all the RJ-45 connections. Not 16gbps.
https://evanmccann.net/blog/unifi-routers-overview
Not true.
The switching chipset is the Realtek 8370 that has full duplex 1Gbps on each port. Since the proof is in the pudding plse have a look at this video. There is so much confusion about this.
Technically it has full, 18Gbit (8 LAN ports plus the uplink), non-blocking throughput across the backplane. It can saturate each port in both directions simultaneously. The limitation is that it is just like any other 9 port Gigabit switch. It has a single gigabit uplink to the router CPU. Worse, it is caught somewhere between being a dumb switch (closer to this) and a fully managed one. This limitation is why it cannot be taken seriously.
@RobbieH I think the bottleneck mentioned in the OP is the switch to the routing engine, not between the 8 ports, they should be switch speed on same L2.
This is why, if you have a lot of VLAN traffic (routing locally), use the 10 Gbps port to a 10 Gbps switch which will also give you RSTP and other switch features.
This is different from the EdgeMAX like the ER4 which has 1 Gbps per port to the routing engine (but no switch0)
@gregorio wrote:
That is because of missing functionality like spanning tree and link aggregation but it does do VLAN and IGMP snooping. Where confusion kicks in, is on this backplane. On a pure switch level, it is capable of switching 2 x 8 = 16Gbps where often this is explained as 2 x 1 Gbps shared amongst the ports which is simply not true so no bottleneck there. The 1Gbps upload to CPU for internet & inter-VLAN traffic is of course a bottleneck if you compare it to Enterprise and Pro level switches. That is no fair comparison because of completely different price levels.
So in a setup where there is no risk of loops and you are not in desperate need of aggregation, this is just as any normal 8 port switch.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

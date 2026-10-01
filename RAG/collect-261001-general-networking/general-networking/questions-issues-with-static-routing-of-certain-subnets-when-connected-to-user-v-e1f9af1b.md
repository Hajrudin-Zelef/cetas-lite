---
id: collect-261001-general-networking/general-networking/questions-issues-with-static-routing-of-certain-subnets-when-connected-to-user-v-e1f9af1b
title: "questions-issues-with-static-routing-of-certain-subnets-when-connected-to-user-v-e1f9af1b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-issues-with-static-routing-of-certain-subnets-when-connected-to-user-v-e1f9af1b.md
source_anchor: ""
source_lines: [1, 13]
sha256: d54b0db88d2fc3da2bf1cd23e18d89a3054a5e5e6af8b7c2078d2774cb78683b
---

# questions-issues-with-static-routing-of-certain-subnets-when-connected-to-user-v-e1f9af1b

@UI-Team
I have a network with two routers, each with its own static WAN. One is the UDM-Pro, one SonicWall. The routers are linked with their own subnet so I can set a destination for the static routes. I have successfully set up static routes and firewall rules such that certain subnets on the UDM-Pro can access certain subnets on the SoinicWall, and static routes and firewall rules are set up on the SonicWall such that it knows where to send replies.
It works flawlessly when I'm connecting to, say, a network printer on the SonicWall from the wireless network hosted by the UDM-Pro. But, when I'm connected remotely over the VPN on the UDM-Pro, I cannot access the same printer.
I use this configuration because the SonicWall is more managed by the software vendor, though I can make some changes, but I have more control of other hardware, and thus it's easier to set up other networks on other hardware. Otherwise, I might just host all the networks on the SonicWall.
Below is a diagram, without my real WAN addresses. A static route on the USG-Pro is set to send traffic destined for 192.168.2.0/24 via 172.31.131.1. A static route on the SonicWall sends traffic destined for 172.18.80.0/24 via 172.31.131.2. Works flawlessly. But if I do the same static route/firewall setup on the SonicWall for the Unifi VPN network (not shown) at 10.100.200.0/24 to also go via 172.31.131.2, I can't connect to the printer and other network devices.
Now, I do know the Unifi Teleport VPN also uses the 192.168.2.0/24 subnet, and, even when Teleport is disabled, that subnet is reserved in some places in the USG, and I don't know how to change that subnet, nor can I easily change that subnet on the SonicWall because it is also tied into a site to site VPN that I cannot change the other end of; changing that subnet would break critical access to the remote server, and is not an option.
How do I get this to work, or is it even possible?
I'm connected to the Unifi VPN with a Mac, and I do have "send all traffic over this VPN" set.
Never mind. I guess the SonicWall and/or UDM-Pro was delayed in updating its internal routing table or policies. It works now. I will leave this post active for informational purposes.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

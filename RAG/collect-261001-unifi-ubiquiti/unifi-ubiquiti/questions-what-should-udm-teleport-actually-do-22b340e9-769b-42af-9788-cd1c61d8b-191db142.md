---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-what-should-udm-teleport-actually-do-22b340e9-769b-42af-9788-cd1c61d8b-191db142
title: "questions-what-should-udm-teleport-actually-do-22b340e9-769b-42af-9788-cd1c61d8b-191db142"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-what-should-udm-teleport-actually-do-22b340e9-769b-42af-9788-cd1c61d8b-191db142.md
source_anchor: ""
source_lines: [1, 17]
sha256: c10eebf7cbacc9650fdea335b64d7826d670bb8f5d474c47ff5e99b6004b1ab6
---

# questions-what-should-udm-teleport-actually-do-22b340e9-769b-42af-9788-cd1c61d8b-191db142

@UI-Team
Stupid question perhaps. I have a UDM on firmware 1.12.22 and was excited to see Teleport supported. I pressed the recommended buttons to enable on the UDM as guided and nothing else. Remote access is enabled.
Next, I configured Teleport on an iPhone 13.
But over 4G cellular Wifiman doesn't stay connected. It connects for perhaps 5-10 seconds then always drops and won't reconnect. Over WiFi on the LAN itself it works fine.
I haven't used a VPN before seriously so was expecting this to "just work" and mean that the iPhone could browse e.g. LAN NAS storage or use Plex etc. as if were connected to the LAN directly when using 4G instead of WiFi. Is that a reasonable expectation?
If so, I am guessing there is there something getting in the way with my 4G cellular carrier (UK ISP EE)? Or what gives?
For reasons I don't understand, my UDM isn't working anymore with IPv6 (only IPv4). Could that be a cause of woe here (e.g. 4G CGNAT?)
I'd obviously like to try out Teleport via another WiFi network too, to help isolate as an obvious next step but haven't had the chance yet.
Any clarity greatly appreciated!
Yes, Teleport is a Zero-config VPN solution, so it is reasonable to expect you'd be able to access devices on your network remotely. It does use a random subnet range, so if you have firewall rules in place that are very restrictive you may need to poke holes in them for Teleport's IP range.
That would only affect accessing devices on your network though, not make the connection in Teleport itself unstable.
It may well have something to do with either your mobile carrier or broadband provider. I've used Teleport successfully on O2 4G in the UK with the UDR I was using for testing being connected through Double NAT from a TalkTalk FTTC Broadband which is IPv4 Dynamic.
Trying Teleport on another WiFI Network may help rule out if it's something with your UDM/Broadband or Mobile Carrier is the issue.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-unifi-guest-net-without-vlan-304d724b-4a04-44de-8507-4a4fcb786982-477ebe97
title: "questions-unifi-guest-net-without-vlan-304d724b-4a04-44de-8507-4a4fcb786982-477ebe97"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-unifi-guest-net-without-vlan-304d724b-4a04-44de-8507-4a4fcb786982-477ebe97.md
source_anchor: ""
source_lines: [1, 27]
sha256: 3b18717344872996eed89ca31b13d685a214a1f64de41a358e801daf5219eb45
---

# questions-unifi-guest-net-without-vlan-304d724b-4a04-44de-8507-4a4fcb786982-477ebe97

@UI-Team
Is there any chance to set up a wireless guest network without VLAN ?
I am new to Ubiquiti, doing mods to the original portal files and to test them.
However, after installation of my first UAP, I get "VLAN required" when trying to define a guest network with a portal.
And I have no VLAN switch.
So I do not care about any security issued, as this setup is only for testing my special portal files.
I hav Unifi COntroller 4.6.6 right now, may be, an older version would be appropiate for my requirements ?
Yeah, just create the new SSID, and check the "enable guest settings / isolation" box. Then go into the guest site settings to choose how they'll access it (e.g. TOS only, simple password, vouchers, external portal ...)
Thanx a lot, that works.
So simply not to define a special network for the guests. OK.
Now next problem showed up:
For first testing, I just defined password auth for the guests, with redirect to "promotional site". Using browser, I can enter pwd, but then a redirect
http://192.168.20.121:8881/redirect?ec=j7ij1X2uwT801k....
is displayed in the URL of the brwoser, and "connection refused".
The IP is the IP of my PC running the controller (W7), and there I do not see an active process listening on port 8881.
Only on 8080.
I guess, there should show up the IP of AP (192.168.1.127) instead.
However, connection to internet is granted, so the proplem only is the wrong redirect somewhere.
Check that you didn't set port 8881 in the config by mistake (or typo it somewhere along the way).
My mistake again: Did not know, that the portal files are served from the controller, which has to be connected to the AP, by cable obviously to use the portal. But this connection was broken, only WiFi from controller-PC to AP active. But this does not seem to be sufficient.
As a newbee in Ubiquiti, I have to say, a bit strange.
Why is permanent cable conn necessary between controller-PC and AP ? Or did I do wrong something else ?
Hey Reiner. Like you mentioned earlier, the hotspot files are served from the controller. In a normal enviroment (home or small business) the controller does not have to run 24x7 because most people will not be running a guest portal. The only time a controller needs to run 24x7 is when guest portal is enabled.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-unifi-site-to-site-vpn-5c173eb9-b88d-408c-a8f6-d1e5cb3ae8bb-5eee5294
title: "questions-unifi-site-to-site-vpn-5c173eb9-b88d-408c-a8f6-d1e5cb3ae8bb-5eee5294"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-unifi-site-to-site-vpn-5c173eb9-b88d-408c-a8f6-d1e5cb3ae8bb-5eee5294.md
source_anchor: ""
source_lines: [1, 13]
sha256: bbe5202b5672ecf1c6a8b089dc6aedb72ba83d9c62bf702cef3da2e5f43d7d77
---

# questions-unifi-site-to-site-vpn-5c173eb9-b88d-408c-a8f6-d1e5cb3ae8bb-5eee5294

@UI-Team
Hello, I am trying to setup a site to site VPN using 1 USG, 1 USG pro and a cloud key gen 2. The USG pro is on one site adopted and is up and running. I have a second site created and have the USG plugged into the same local network as the USG pro (with a static lan and the DHCP server off). The USG shows up in the new site but when i click adopt it says adoption failed. Is this a problem with the USG or am I going about this the wrong way? I am new to using a cloud key to manage multiple sites. :)
@jpglanz What's the output of the command info from the USG that's failing adoption?
@jpglanz Looks like the USG doesn't have an address on it's WAN. Is that USG connected to your default LAN of the functional USG? If so, and you left your default LAN as 192.168.1.0/24, that would be an IP conflict as that's how the 2nd USG's LAN is currently configured. You'd need to connect to a different network (anything other than 192.168.1.0/24).
The Lan on the USG is plugged into the LAN of the other USG Pro but that network is a 192.168.0.0/24.
@jpglanz If you connect another device to the cable that's plugged in to the USG's WAN, can it obtain an IP address? Have you tried rebooting this USG?
Can you see any DHCP Discovers on your primary USG? Run the following from your primary USG's CLI:
sudo tcpdump -nevvi eth0 port 67 or port 68
@UI-JSee I DMed you a screenshot of the output of that command. I also rebooted the USG several times.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

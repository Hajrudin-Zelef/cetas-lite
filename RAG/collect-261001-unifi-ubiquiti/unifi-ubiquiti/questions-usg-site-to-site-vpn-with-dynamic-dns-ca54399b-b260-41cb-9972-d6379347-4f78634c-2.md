---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-usg-site-to-site-vpn-with-dynamic-dns-ca54399b-b260-41cb-9972-d6379347-4f78634c-2
title: "questions-usg-site-to-site-vpn-with-dynamic-dns-ca54399b-b260-41cb-9972-d6379347-4f78634c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-usg-site-to-site-vpn-with-dynamic-dns-ca54399b-b260-41cb-9972-d6379347-4f78634c.md
source_anchor: ""
source_lines: [81, 97]
sha256: 6e36e94b553a7156d46f9993c3c786612724c80eb4cf72e79ad4c447a6f8563e
---

# questions-usg-site-to-site-vpn-with-dynamic-dns-ca54399b-b260-41cb-9972-d6379347-4f78634c

There seems to be a general misunderstanding here about how this is used, people seem to think that being able to use a hostname (dyndns) makes the other end able to reach the device with Dynamic IP and stablish the VPN when in reality when something other than an IP is configured as the Gateway ID for the remote the device does not attempt to establish the tunnel and just waits for the other end to do it, I believe for security reasons. Watchguard for example even lets you specify the remote IP and Identifier separately, and gives you the option to select "Dynamic IP Address" instead of the remote IP which basically makes it listen only as explained above. So the problem isn't not being able to use a dyndns hostname, but not being able to use anything other than an IP...
The problem with this workaround (or using any other invalid or wrong IP) is that it will make the other end think that is indeed the IP of the remote (unless it has an option like the Watchguard has) and start trying to connect to it...
Seems stupid to me that people have been complaining about this for so many years and UBNT still forces you to use an IP as the identifier when pretty much all other vendors give you the option to use domains, hostnames and even x500 records...
And... add another UI customer to this thread, iv got two sites, both with dynamic IP's, a Draytek at one end and a USG Pro at the other... just tried creating a IPSec VPN between them, and PEER IP / Local WAN IP... :(
Looks like I just spent $130 down the drain. :(
This is so stupid, just fix it ubnt!
so still the same :'( no dynmaic ip because of the old strongswarm version?
what is wrong guys?
old Hardware in the USGs? just talk to us!
Yep, another issue here due to lack of DNS name support.
The ERL has always had it, so it really can't be that difficult...
Still no solution? bought 2 UDM Pro to setup a s2s to my company just to find out it doesnt work with dyndns. big disappointment.
Almost all home & business routers have this capability. Has there been any progress or updates since the last post above which was 10 months ago?
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

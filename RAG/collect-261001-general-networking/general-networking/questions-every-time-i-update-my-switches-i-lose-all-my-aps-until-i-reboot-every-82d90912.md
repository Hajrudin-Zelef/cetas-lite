---
id: collect-261001-general-networking/general-networking/questions-every-time-i-update-my-switches-i-lose-all-my-aps-until-i-reboot-every-82d90912
title: "questions-every-time-i-update-my-switches-i-lose-all-my-aps-until-i-reboot-every-82d90912"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-every-time-i-update-my-switches-i-lose-all-my-aps-until-i-reboot-every-82d90912.md
source_anchor: ""
source_lines: [1, 18]
sha256: 37298f1e14c6b61d9f94467c0f8201b8f6dfcbdae9149228057b926e7170af44
---

# questions-every-time-i-update-my-switches-i-lose-all-my-aps-until-i-reboot-every-82d90912

@UI-Team
Hello:
I have a Fiber, 4 switches and 12 APs.
Every time I update a switch, I get stuck in a loop where the APs try to adopt, show "action required," and then "adoption failed." If I unplug everything and plug it back in, I get everything back.
I need to update all of my switches but I'm dreading it because it will take a while before I'm back online again. Any idea how to avoid this?
Thanks!
Looks like your hardware is configured to obtain an address by DHCP. Does setting Static IP for the Unifi hardware solve the issue?
It would be helpful to all to know the exact hardware involved, what firmware they are running, your network app version, how the APs are configured for network, as well as any log entries of interest.
Does a POE cycle of the switch or a reboot with POE cycle solve it?
Good question: Yes, a power cycle of the switches does solve it. Sometimes an AP will stay offline, and I can cycle the port and get it back.
Here is some of my hardware info:
Do you mean assign a static IP for all of the switches and the APs?
Yeah. Unifi equipment has a nasty habit of falling back to 192.168.1.20
If they do that en-mass, then they will be too busy fighting for that IP to notice that DHCP is back.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

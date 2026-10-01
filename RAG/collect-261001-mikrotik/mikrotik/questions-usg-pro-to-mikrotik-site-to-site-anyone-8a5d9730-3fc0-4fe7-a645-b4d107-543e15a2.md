---
id: collect-261001-mikrotik/mikrotik/questions-usg-pro-to-mikrotik-site-to-site-anyone-8a5d9730-3fc0-4fe7-a645-b4d107-543e15a2
title: "questions-usg-pro-to-mikrotik-site-to-site-anyone-8a5d9730-3fc0-4fe7-a645-b4d107-543e15a2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-mikrotik/questions-usg-pro-to-mikrotik-site-to-site-anyone-8a5d9730-3fc0-4fe7-a645-b4d107-543e15a2.md
source_anchor: ""
source_lines: [1, 30]
sha256: 3e9f0ab2ea4021f44a2bcdbe71b810c1f66da16c5e3666afb67b9954e60b2cd8
---

# questions-usg-pro-to-mikrotik-site-to-site-anyone-8a5d9730-3fc0-4fe7-a645-b4d107-543e15a2

@UI-Team
Hello
I am trying to setup a site to site VPN between an USG and a Mikrotik firewall and I must confess some difficulties.
I understand the L2TP would be the way to go but I can't seem to have both router talking at all (if someone has it working please share !)
I have managed to bring up a PPTP vpn (I know, not ideal, although in this case there are no major security concerns) but if subnet A can reach sunet B without issue I can only reach subnet A gateway (the USG at 172.16.107.254) from subnet B.
From a host on subnet B:
C:\Windows\System32>tracert 172.16.107.254
Tracing route to 172.16.107.254 over a maximum of 30 hops
  1    <1 ms    <1 ms    <1 ms  172.16.100.254
  2     5 ms     5 ms     5 ms  172.16.107.254
Trace complete.
C:\Windows\System32>tracert 172.16.107.200
Tracing route to 172.16.107.200 over a maximum of 30 hops
  1    <1 ms    <1 ms    <1 ms  172.16.100.254
  2     5 ms     5 ms     5 ms  172.16.255.100
  3     *        *        *     Request timed out.
Any suggestion / pointer most welcome !
Well anything that works :)
As mentioned I have tried IPSec - the tunnel comes up but not traffic either way.
I have then moved to PPTP out of expediency - it is working better but as described not 100%
I might give a try to OpenVPN... but from experience the Mikrotik (did you notice that I have a different brand at the other end) is somewhat "delicate". But worth trying I guess
@soleroit wrote:
Did you made firewall rules to allow traffic from subnets?
Did you used the next to troubleshoot https://help.ui.com/hc/en-us/articles/360002668854?
Cheers,
Mike
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

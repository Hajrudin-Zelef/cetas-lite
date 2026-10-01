---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1q5t3z7-localin-policy-not-applying-1a94beeb
title: "r-fortinet-comments-1q5t3z7-localin-policy-not-applying-1a94beeb"
domain: fortinet
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1q5t3z7-localin-policy-not-applying-1a94beeb.md
source_anchor: ""
source_lines: [1, 35]
sha256: 64c13137ade9500652e150d9f7a8b5f3780d0f41d0b107df79e56f3e0eaa841c
---

# r-fortinet-comments-1q5t3z7-localin-policy-not-applying-1a94beeb

Local-in policy not applying? 
        
    Fortigate 40F running 7.2.11
I have the following local-in policy where I am trying to prevent any access to the fortigate from IPs that are blasting our ssl vpn. For some reason, I am still seeing "sslvpn_login_permission_denied" messages in the logs. It was my impression that creating this policy would stop any access to the fortigate where the IPs were in that defined address group.
I initially tried "set service "SSLVPN"" (we have a service configured with that name), but it wasn't working either. Am I wrong in my thinking that this is where I should configure things? I can't geo-block b/c many of the IPs are coming from US hosting companies.
I read through a few guides and on here, and it looks like this should work.
config firewall local-in-policy
edit 5
set intf "any"
set srcaddr "grp_summarized_blocklist_16"
set dstaddr "all"
set service "ALL"
set schedule "always"
set comments "Deny SSL VPN from blacklist IPs"
next
Section des commentaires
I am yet to see a case where local-in policy would not work as expected, so:
- Make sure this rule is top-most, as being rule 5 means there are other rules, possibly above that may or may not allow the very same traffic.
- Make sure the targeted SSL VPN IP sits on the Fortigate itself, not routed or a VIP as then it would not work.
- By default, Local-in policy hits are not logged, you have to set in Log Settings → Log All for denied packets to be logged. The logs are in Local Traffic section.
Ah...SSL is using a VIP.
Thanks for all the feedback and suggestions. I think we'll move to using a loopback interface as most noted. Eventually we'll do IPSec VPN once I get the 40Fs replaced
What are rules above this rule? Perhaps rules above are allowing in ?
No rules above that one when I do a show. There are 4 in total, but they are specific to other src/dest
Is your SSL vpn terminating directly to the fortigate interface IP?
or do you have loopback / VIP setup?
Different topic, but you probably want to be moving away from SSL VPN.
Yep, already working away from it 👍🏻
Nice one.
Is the default action deny? Check show full config in this rule
The default is indeed deny: https://docs.fortinet.com/document/fortigate/7.6.5/cli-reference/185227842/config-firewall-local-in-policy
From IPs you want to block?
Assuming your other local-in policies aren't messing something up the policy you posted will block any attempts to connect to your FortiGate from the IPs in "grp_summarized_blocklist_16".
You can move the sslvpn to a loopback and use a vip which allows you to make normal policies if you are having a hard time with local-in policies. Gives you access to isdb objects and threat feeds as well for your ssl vpn policies. It might effect offloading though, but I can't recall for sure.
I‘d recommend to use loopback and block bad ip‘s using the ISDB. https://community.fortinet.com/t5/FortiGate/Technical-Tip-SSL-VPN-connection-to-a-Loopback-Interface-using/ta-p/328376

---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-16zoxpd-network-goes-offline-after-any-firewall-policy-ca613691
title: "r-fortinet-comments-16zoxpd-network-goes-offline-after-any-firewall-policy-ca613691"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-16zoxpd-network-goes-offline-after-any-firewall-policy-ca613691.md
source_anchor: ""
source_lines: [1, 53]
sha256: fc516d2f46061d37e6efd34f7c8389178c3a20ff5aa190ab61db564710462d2a
---

# r-fortinet-comments-16zoxpd-network-goes-offline-after-any-firewall-policy-ca613691

Network Goes Offline After Any Firewall Policy Change 
        
    Hello everyone,
I am experiencing a weird problem with my FortiGate and wanted to see if anyone has encountered this issue before or has any troubleshooting suggestions.
Background:
FortiGate HA Active-Passive Pair (601F)
Version 7.0.12
Issue:
Whenever we make any edits to any existing firewall policy, or create any new policy, the entire network goes down. All internal users cannot reach the internet and VPN-connected users cannot access internal resources.
Once the firewall policy changes are reverted, network connectivity is restored. During the outage, system resources are stable and there are no logs indicating any problem taking place on the FortiGate. This change can be as simple as adding a service or IP to a policy, either way, the network goes down once the config is applied.
Has anyone experienced an issue like this or have any troubleshooting suggestions? We have opened a case with TAC but cannot do any testing with them until late tonight as any sort of live testing will bring down the network during production hours.
Section des commentaires
Thank you everyone for your insight. We are working to schedule an outage window so we can test some of these solutions. Additionally, during this outage window, we will have TAC on to hopefully diagnose the issue. I will report back with an update in the next few days.
Man, that really sounds like a bug. I suspect it will require live testing and debugging to isolate a possible cause. Best of luck on that one. Definitely strange!
For anyone still following this, we still do not have a solution.
Due to the critical nature of this site, it is hard to find planned downtime windows. We will be troubleshooting with TAC next Friday night.
Guys, I am sorry for the delay. We did get this issue resolved.
This customer had a large SDWAN configuration consisting of 5 different ISP circuits. All of these circuits were FIber, one of which was an ethernet ISP circuit. The SDWAN was configured so it load balanced all traffic equally across all circuits, including the ethernet circuit. During packet captures at the time of the issue, we would see the traffic egressing the ethernet circuit. We removed this circuit from the configuration as a test and have not seen the issue since, so it was specifically a problem with that ISP circuit.
Never seen this before but from the other replies. I would say if you are using IPsec tunnel. Would be wise to implement the following blackhole routes:
config router staticedit 60set dst10.0.0.0255.0.0.0set distance 254set blackhole enablenextedit 61set dst172.16.0.0255.240.0.0set distance 254set blackhole enablenextedit 62set dst192.168.0.0255.255.0.0set distance 254set blackhole enablenextend
So if those tunnel(s) go down traffic won't be routed through the default route. But instead be dropped until the tunnel comes back up.
This!
Please explain?
It's exactly what is written. If the tunnel is down RFC1918 traffic without a specific route would take the default route, which is a bad use of resources.
Maybe check bug 843554.
It doesn't affect cli changes, so if gui takes down traffic but cli works properly, try the workaround for this bug.
I had something that sounds quite related, back when 1500D was new. after migrating/importing the policy, nothing worked. we had to enter every single policy in the GUI and confirm it to make it work ...
some more debugging would be helpful in your case, interfaces down? policies not matching?
No interfaces down. Hopefully I will have an update once we are able to have a session with TAC.
don't know your contract, but usually would recommend to first debug with your fortinet partner. good luck and looking forward to an update :)
I have experienced this. Same fix. Reboot or revert change. What I found was the logs showed everything was getting denied by policy 0.
Never got around to taking it up with TAC.
By any chance are you using fortimanager and had a number of changes pent up before you installed to the gate? That was the circumstance I thought might have triggered it.
Thanks for the insight. No, this site is not managed by FortiManager.
This is what I have experienced. Suddenly stuff being dropped by the rule 0. It looks like for some reason making some policy or routing changes causes the sessions to not match on the original rule that was allowing traffic. If you flush the session table they re-establish correctly and it starts passing traffic again. I have had this happen several times. It is usually traffic that is routed across a tunnel and the tunnel drops. The traffic then starts trying to go the default route then won't return to the tunnel when the tunnel comes back up.
Really weird, please update later.
I had this on 7.2.5 on 200F HA pair, push from FMG and all traffic stops, local web admin interface policy view shows 0 traffic bytes on all policy rules.
Editing the policy locally e.g. disable a rule and re-enable and suddenly all traffic starts flowing again.
Easily repeatable.
Reboot resolved it...for now.
TAC have been no help
Bit worrying, never seen anything like it
Did you clear all sessions after traffic stopped? Seems whenever I make changes any traffic that loses it normal route will creat bogus sessions that break a bunch of stuff. Clearing the session table resolves this.
This is a strange behavior. I have only seen this type of behavior when the firewall rule table is way large and changing policy orders push a traffic redistribution. On my case this was a device with 35k policies badly implemented
#Offtopic; You are the one who is being mentioned in the release notes known/fixed bugs. Always was wondering who has 35k+ policies. Is it even possible to view those in the GUI?
Is this a new issue in 7.0.12? This is very disruptive so you must have noticed this before for any policy change?
Does it happen if you create a new policy that is added to the end of the list?
I believe when there is a change to a policy, the FGT needs to re-evaluate all existing sessions against the new policy list since it doesnt know if the new change may block or allow the traffic.
How many policies and active sessions do you have? Try clearing all sessions like other ppl said to see if that help.
Pair of 200F firewalls in a virtual cluster. Had the same issue as @stevo1666 where all interfaces appear to stop passing traffic after a configuration change, the changes are very basic object or policy updates and there is no way the change could cause this issue. The other issue that occurs is Fortinet support have no idea how to resolve, troubleshoot or do anything to assist in this issue.
Up to 10 outages since 2023 now, upgraded from 7.2.7 > 7.2.8 > 7.4.7 > 7.4.9
Clearly not related to a FortiOS bug, may have to RMA these firewalls
anyone else seen and had this issue resolved by a config or design change?

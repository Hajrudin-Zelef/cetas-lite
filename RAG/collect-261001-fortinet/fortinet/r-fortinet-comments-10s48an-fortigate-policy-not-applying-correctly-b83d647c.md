---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-10s48an-fortigate-policy-not-applying-correctly-b83d647c
title: "r-fortinet-comments-10s48an-fortigate-policy-not-applying-correctly-b83d647c"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2023-06-02"]
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-10s48an-fortigate-policy-not-applying-correctly-b83d647c.md
source_anchor: ""
source_lines: [1, 39]
sha256: 5b6dea955cd305d10a965dee57a9b6476f71ada999f3b5f8fee522a519ad2387
---

# r-fortinet-comments-10s48an-fortigate-policy-not-applying-correctly-b83d647c

Fortigate Policy not applying correctly?
Been scratching my head over this and can't figure it out.
I have a Fortigate 80E that remote users connect to via SSL VPN. To improve security, I created another VPN user group with a separate IP range intended for IT. I then created/applied policies to grant this IT user group and users the same access as the regular user group and more. However, its not able to access IPs it should
E.G. I have a subnet range 10.59.1.0/24. The Fortigate has an interface in the same subnet. . If I connect to the VPN using the IT user, I can ping devices on that subnet. I can connect to the web interface for a server. I can not connect to the Fortigate web interface but can ping it.
Policy lookup / iprope returns policy ID 0, aka implicit deny. So i do some research, verify settings, but everything looks correct. I then tried adding the IT user group / ip range to a policy that allows access to the internet and was already being applied to the existing VPN user group. However, it returns policy ID 0 and doesn't work either.
      Firewall Policy in CLI
Firewall Addresses in CLI
Firewall iprope lookup result
    
Hoping someone can point me in the right direction or tell me what's wrong. Both user groups (regular and IT) have their own portal and portal mapping. The only difference there between the two is they have different Source IP Pools.
Edit: 02/03/2023Ok so I'm currently in the firewall and have been running
diagnose debug flow trace
      to test. From what I can see, attempts to access the firewall interface over HTTPS are being dropped by a different policy. The problem is that policy is applied to a different everything; src/dst.port and src/dst.addr.
Debug flow for HTTPS attempt
Policy 4 (everything is different that debug flow)
    
Edit 2: 02/06/2023
r/3rid pointed out that local-in policies do infact have a policy ID number in the CLI. Went through them and saw policy 4 was an implicit deny all. Added the new VPN group IP range to the currently allowed groups and can now access the HTTPS interface of the firewall without issue. Mystery solved.
Section des commentaires
Try creating separate firewall policies for each group/vpn subnet instead of combining both in the same policy.
Same result. Separate policy, same source/destination interface.
Do you know how to use "diag debug flow"? Try using this tool and paste the output from it.
https://docs.fortinet.com/document/fortigate/6.2.12/cookbook/54688/debugging-the-packet-flow
I did, see my first edit in post. Note how the policy that dropped it is completely unrelated.
We have a forticare contract for this firewall, so I've opened a ticket with them and hopefully can get someone soon to help me out.
The debug flow screenshot is unreadable for me. Are you sure there is no conflict for new IP range? Maybe firewall is expecting it on another interface? Can you check if there are any logs in gui regarding this flow?
Sorry if the screenshot isn't clear. I attempted to access 10.59.1.254 via HTTPS via my web browser from 10.0.101.1 . The debug flow shows it failing a check on policy 4 and dropping the packet. From the debug flow, you can see the traffic came in from the ssl.root interface, which the SSL VPN connection flows through.
I've verified there is no conflict with the new IP Range. Policy 4 has a different source and destination interface. The "to" and "from" ip addresses for that policy are both a /32 not in the 10.0.101.0/24 range.
Interestingly enough, in "Log & Report > Forward Traffic" there are no hits for policy 4. The only hits for source ip 10.0.101.1 are from an hour earlier when i tried deleting the allow policy, tested pings, then recreated the policy. Those all hit the implicit deny policy.
No problem. Did you check the local-in policy 4?
Did you add a static routes for the IP ranges going out the SSL VPN interface?
Check your routing table.
These are already in place and working correctly for the existing SSL VPN group. That said, I went back through and verified them once more. Given they are still working for the existing VPN users, I don't believe the issue lies there.
I get that but you stated you created a new IP range and user group.
So did you create...
A new authentication rule in the sslvpn settings A new route configuration for the new IP range
Commentaire supprimé par le membre
Unfortunately not. See my first edit. There is some strange fuckery going on.
This doesn’t seem quite right but I don’t know enough about the happenings of cephalopod lairs to dispute it.

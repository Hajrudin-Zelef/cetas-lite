---
id: collect-261001-fortinet/fortinet/questions-943714-is-it-possible-to-have-name-resolution-from-fortigate-and-local-46d44d9c
title: "questions-943714-is-it-possible-to-have-name-resolution-from-fortigate-and-local-46d44d9c"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-943714-is-it-possible-to-have-name-resolution-from-fortigate-and-local-46d44d9c.md
source_anchor: ""
source_lines: [1, 25]
sha256: d7f65675f95ade05cd8e6dd6570f1a867f344c3c0bd60482fdf03f3e27a0ae94
---

# questions-943714-is-it-possible-to-have-name-resolution-from-fortigate-and-local-46d44d9c

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
3
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Can you advise on moving to a hybrid DNS?
Currently, all our LAN machines receive their IP address from our Fortigate 60D (each machine is either allocated an IP address from the Fortigate DHCP, or has a static IP address set in the Fortigate).
Our DNS records are currently managed from fortiddns.com.
Can I create a local DNS server, that will perform name-resolution for some of our LAN machines?
That being said you can certainly can set up DNS internally for it to resolve the locals hosts. Each time you add a host you will want to make sure their DNS entry is added to the server if you don't have a way for them to register themselves (active directory usually takes care of this in a Windows environment). I would advise against using .local since it may cause you pain down the road and has been advised against using for local networks. In AD environments you typically see a company with a site MyCompany.com with a public DNS server for all things web. Internally you would create a subdomain for your DNS server like office.MyCompany.com which you wouldn't actually publish to your public DNS and it would be valid internally only. Your hosts would end up looking like computer1.office.mycompany.com and since your internal DNS server only knows to serve the office subdomain it would look externally to resolve anything else. This is the more proper way to do it. If you wanted to keep some reference to .local you could use local.mycompany.com. for your resolvers you would set the hostname to just be computer1 and have it search your configured DNS server for local.mycompany.com
Its not exactly clear what you are trying to achive. I think you mean that you have some hosts which are provisioned via DHCP which is linked to a DNS server, thereafter it gets a bit confused.
Do you also have hosts considered to be in the same zone but which are not provisioned with DHCP which you want to add to the DNS zone?
If this is the case, and the other hosts are on the same subnet then the simplest option is to add static mappings in the DHCP server for these hosts
If the other hosts are on different subnets, then it's possible this configuration won't work and/or won't solve the problem - my first guess at a solutionn would be to use DNSmasq which can acts as a DNS forwarder proxy but can override specific records
Are all the the hosts in scope provisioned via DHCP but the linked DNS server is outwith your current DNS resolver provision?
Simply define the zone and DNS server to your existing DNS forwarder service
if it'd be possible, I'd continue to use *.local as the FQDN
It's probably possible but not advisable. This is the best time to change it.
I can see this is an old question, but I have an alternative idea you can use.
Let's say fortiddns.com is hosting your main domain example.com and you want the FQDN server.example.com to resolve to a local IP address and any other lookup to the domain example.com will resolve to the official record defined in fortiddns.com.
You could then install Bind 9 on a local server where you use Response policy zone to override any lookups when the request matches the query server.example.com.
This is basically how Lancache.net uses DNS to cache installation files for games and windows updates on local LAN, which is completely transparent to the end user.
I can update the answer with how I use Response policy zones when I get home from work.
.localas the domain in our LAN (e.g.,server1.local), so, if it'd be possible, I'd continue to use *.local as the FQDN. None of our LAN machines are exposed across our firewall, so that shouldn't be a problem.

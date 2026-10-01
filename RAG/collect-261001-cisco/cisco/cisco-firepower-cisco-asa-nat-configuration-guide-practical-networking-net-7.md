---
id: collect-261001-cisco/cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net-7
title: "cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net"
domain: cisco
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net.md
source_anchor: ""
source_lines: [886, 1056]
sha256: 0dd3b5204e10e711699d869cc65c566d56c85da10078a2e14401bd24b1f20ee4
---

# cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net

In such cases, you would configure a NAT Exemption on the ASA for traffic from `10.1.1.0/24` to `10.2.2.0/24`. As mentioned before, the NAT Exemption configuration involves configuring an Identity NAT, i.e. translating traffic to itself:

object network SEATTLE
  subnet 10.1.1.0 255.255.255.0
  nat (inside,outside) dynamic 72.3.3.77
object network DENVER
  subnet 10.2.2.0 255.255.255.0
nat (inside,outside) source static SEATTLE SEATTLE destination static DENVER DENVER

With the configuration above, the Manual NAT statement would appear in Section 1 and take precedence over the Auto NAT statement which would appear in Section 2:

asa98# **show nat detail | exclude hits**
Manual NAT Policies **(Section 1)**
1 (inside) to (outside) source static SEATTLE SEATTLE destination static DENVER DENVER
 **Source - Origin: 10.1.1.0/24, Translated: 10.1.1.0/24**
 **Destination - Origin: 10.2.2.0/24, Translated: 10.2.2.0/24**
Auto NAT Policies **(Section 2)**
1 (inside) to (outside) source dynamic SEATTLE 72.3.3.77
 **Source - Origin: 10.1.1.0/24, Translated: 72.3.3.77/32**

Traffic from Seattle to Denver would match the Manual NAT statement and *not* be translated (Identity NAT), and traffic from Seattle to anywhere else on the Internet would match the Auto NAT statement and be translated using Dynamic PAT to `72.3.3.77`.

## Cisco ASA NAT – Summary

The Cisco ASA and Cisco ASA-X firewalls provides nearly infinite flexibility in so far as their NAT configuration. From the modularity of using objects, to the simplicity of configuring Auto NAT, to the granularity of Manual NAT, to the precision of NAT precedence — the ASA can do it all.

This article covers each of these concepts in detail, explaining what they mean, when to use them, and how to apply them. If you read it from start to finish, and were able to follow with the examples and illustrations, then you can decidedly consider yourself an address translation expert on the Cisco ASA platform.

Awesome explanation!!

Thank you so much.

You’re welcome. =)

Checkpoint and fortigate are so much more friendly to the admin with NAT config, at least their naming does exactly what it sound like

I’m sure its an exposure thing. I’ve gotten very used to the ASA’s configuration of NAT. Specially since 8.4.

But, if you like the way Checkpoint or Fortigate does it and/or names their iterations of NAT, please let me know what words they use and I’ll link them on the NAT Terminology Disambiguation page.

Thank you very much for your explanations. They help to understand everything in a very easy way. They are spectacular and, above all, very helpful.

Glad they helped, Sergio =)

What a fantastic guide. I know from CCNA 2, a few weeks ago, NAT is my weakest link. This will help, Thank you.

Glad you enjoyed it! Definitely check out the NAT article series if you are looking for a rundown on NAT as a technology (not tied to a specific vendor): pracnet.net/nat.

This is very helpful, thanks so much!

never seen such a great documentation before..Thanks a lot

Mate , you genuinely saved my life

Hi Karol. In that case, I’m really glad you found the article =)

Very well written article on NAT , thank you !

Some really terrific work on behjalf oof the owner of tthis site, deadd outstanding content.

This was exactly what I needed. Thank you so much for writing this! Using this new knowledge to set up my GNS3 lab.

Glad you enjoyed it, Jack =) Hope it helps with your GNS3 studies.

How would one generalize the Google DNS redirect example for any external DNS server? Does 0.0.0.0 act as a wildcard?

object network ANY-DNS

subnet 0.0.0.0 0.0.0.0

object (Inside,Outside) source dynamic All_Int_Nets DPAT-IP-DNS destination static Any_DNS CORP_DNS service UDP53 UDP53

Something like that? Thanks

Hello Ed Harmoush,

please can you also add a doc for the NAT with IPv6 ?

NAT64

NAT46

NAT66

NAT66 PAT

DNS64

ref.:

https://www.cisco.com/c/en/us/td/docs/security/asa/asa910/configuration/firewall/asa-910-firewall-config/nat-reference.html#concept_5FBE69B32F8E4A499276904DF6A2BB21

Without doubt this is the best guide on the web for using NAT on ASAâs. Many thanks.

Thanks for the very informative read through and your efforts are appreciated to untangle the various dubious NAT concepts.

I would like to request for a additional explanation of “route-lookup” attachment at the end of the NAT statement. Which is quite confusing and so far I have not found a very promising explanation.

Thanks once again!

Hi techkludge. I thought of adding

`route-lookup`but its behavior has changed with different code versions. In order for this document to apply to as many code versions it could, I opted to leave those details out. Maybe it will become a future article. Who knows.
Glad you enjoyed the article though =)

Thank you Ed for this very useful document that explains a lot.

I am not very clear only about the precedence if doable, when we use the same Static NAT on the source for Identity NAT on two different destinations

1. nat (inside,outside) source static INSIDE66 OUTSIDE66 destination static HOST45 HOST45

2. nat (inside,outside) source static INSIDE66 OUTSIDE66 destination static HOST55 HOST55

Thanks again!

Dear Ed,

thank you for your brief explan of NAT.

would you mind if add about CGN and how ipv6 works regarding with NAT ?

Still remain nice,

No doubt the best explanation of Cisco ASA post-8.3 NAT!

However, I have some doubts: in all examples and syntax highlights with regard to destination NAT, you stated something like

“… destination static “, but Cisco’s syntax (8.4, 9.x) says:

nat [ ( real_ifc , mapped_ifc ) ] [ line | { after-object [ line ]}] source static real_ob [ mapped_obj | interface ] [ destination static { mapped_obj | interface } real_obj ] [ service real_src_mapped_dest_svc_obj mapped_src_real_dest_svc_obj ] [ dns ] [ no-proxy-arp ] [ inactive ] [ description desc ]

This basically says that with destination NAT, first the MAPPED address is to be stated, and then the REAL, which is quite the opposite from the syntax given here. Perhaps I’m missing something?

Thanks

I second this. The Cisco documentation is a reverse of what you’ve stated here. Can you please clarify further on the destination side of things? Thank you

The previous comment did not render properly, so instead ââ¦ destination static â should be:

“… destination static REAL-DST MAPPED-DST”

I guess that less than and greater than signs make page not display comment right way ð

I can’t thank you enough. I finally understand NAT properly.

Awesome…. Keep it up. Very much helpful

This is by far the best explanation I’ve ever read on ASAs. Not just the ASA, any posts on practicalnetworking.net are the best explanations I’ve ever read of networking concepts.

As they say, simplicity is the virtue of genius. Thank you, Ed! Please keep adding more and more to the website.

Very very good explanation I watched too many videos also raid too many blogs but not a crystal clear explanation like you. Good job. But Cisco nat is à¤µ complex than other Utm or network devices.

Hi Ed ,

Your blog is really helpful and has helped cleared a lot of doubts i previously had .I look forward to your articles and my only grievance is that you dont update in frquesntly enough ð

Ipsec topic says coming soon for a long time now , Still patiently waiting

hey everyone, thanks for the article it’s well explained and so informative.

For OUTBOUND traffic I tried to use a Manual NAT with static PAT, and I wanted to use the outside interface for mapping and I got an error:

i have created a network object with the same IP address of the outside interface

ERROR: Address 172.16.0.1 overlaps with outside interface address.

ERROR: NAT Policy is not downloaded

Other wise I must use another IP address or AUTO NAT if I want to use the IP address of the outside interface.


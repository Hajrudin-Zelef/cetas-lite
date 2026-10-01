---
id: collect-261001-cisco/cisco/r-cisco-comments-8nkgge-everything-you-need-to-know-about-nat-c9c8ce19-1
title: "r-cisco-comments-8nkgge-everything-you-need-to-know-about-nat-c9c8ce19"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-cisco-comments-8nkgge-everything-you-need-to-know-about-nat-c9c8ce19.md
source_anchor: ""
source_lines: [1, 55]
sha256: e36e055271461df353e9ba1666af8aaaa23d907162587599510ec0960e14f9ae
---

# r-cisco-comments-8nkgge-everything-you-need-to-know-about-nat-c9c8ce19

Everything you need to know about NAT configuration on Cisco ASA and ASA-X Firewalls (v8.4+)
I just released the ultimate Cisco ASA NAT configuration guide -- 100% free, no sign up, no paywall, no mandatory e-mail subscription, no nothing =).
The article is incredibly thorough and covers pretty much everything you would need to know to really understand how to configure Auto NAT or Manual NAT to accommodate just about any address translation scenario imaginable.
The article is available here:
pracnet.net/asanat
To make it easy to check out a specific section, here are all the topics covered in the Configuration Guide:
I sincerely believe that if you read the whole guide, start to finish, you could go from having no exposure to Cisco ASA's NAT functionality, to being an expert at it. I'm open to hearing any feedback which might agree or disagree with that.
Hope it helps =)
Edit: remove the small note about EoL announcements, the goal of the post was to announce the configuration guide not discuss the EoL/EoS
Section des commentaires
Thank you, but can you tell us which EoL announcement you're referring to?
ASA's are End of Life 03/2018 and End of Support 05/2023
ASAx's are End of Life 12/2017 and End of Support 05/2023
More info here: https://www.cisco.com/c/en/us/products/security/asa-firepower-services/eos-eol-notice-listing.html
Well to be fair those are just the 5585-X firewalls with FirePower, not every ASA.
Based on your previous articles i know this is gonna be a simple yet thorough read :-)
I appreciate the kind words =)
Pretty comprehensive indeed, great work.
However there are 2 things I noticed:
This sentence seems to be misleading: "NAT statements are written from the perspective of outbound traffic (traveling from inside to outside)." It's perfectly fine to create (outside,inside) PATs / NATs, one is not restricted to (X,outside) pairs. Personally, I only ever do (X,outside), but understand why some of my colleagues might prefer the more "natural" order "outside,inside" if they think about allowing access to Outlook Web App FROM outside TOWARDS Exchange Server on the inside...
With Manual NAT, the order of real and mapped is swapped for the destination side of the equation! The mapped object comes first, followed by the real object. Same for the destination service (port). See Cisco documentation
Example: I've got an inside host 192.168.10.3 behind an ASA whose HTTPS web GUI I want to access from outside, from one specific external IP 10.10.10.10. However, I need to present the external administration client to the inside device as an L2 adjacent device (let's say I don't want to configure any routing at all on that device, or that the security posture only allows mgmt access from a directly attached device on the same subnet). Additionally, TCP 443 is already in use on the outside IP of the ASA, so I'm going to use 1443 instead and translate to 443.
So we've got: Static PAT from outside interface (intf IP X.X.X.X of ASA) port 1443 to inside 192.168.10.3 port 443 and at the same time static NAT from outside 10.10.10.10 to inside 192.168.10.95
The following 2 lines achieve the intended effect and are functionally identical: (wordily named objects to simplify illustration)
OR
Thanks for taking the time to write that out. I can see how it can be misleading. I struggled with going into more detail about the order (real/mapped) but I opted to omit that topic from the (already very long) guide. So let's talk about it here.
The difference between your two NAT statements have to do with Proxy ARP.
So I'll start off acknowledging, the translation effect between both the statements you provided works identical. If the traffic gets to the Firewall, either application of Manual NAT will have the same translation.
The difference will be in how (if) the packets get to the Firewall.
In your example, your Static PAT is using the Firewall's Interface IP address. But for the sake of explaining the nuance, let's use this image, where the translation is NOT using the Firewall's Interface IP:
http://www.practicalnetworking.net/wp-content/uploads/2017/01/arp-media-proxy-arp-nat-topology.png
The image indicates a Static NAT, but the idea will apply to the Static PAT as well. If we were to configure this Static PAT using Manual NAT, we would have the following:
The upstream router (C) is directly connected to the 72.3.4.0/24 network. If Router C receives a packet with a destination IP of 72.3.4.55, it will issue an ARP Request to the 72.3.4.0/24 network.
Due to the configuration above, the ASA would respond and provide its own MAC address to Router C. Then Router C can then create the L2 header to get the packet to the Firewall.
Here is the key...the ASA will only answer ARP requests based upon NAT on the Mapped interface for the Mapped IP addresses.
Which means, if you inverse the NAT statement, the ASA would only answer ARP requests on the Inside interface:
This isn't 100% proper syntax -- you can't do a destination only clause, but you get the idea
So the traffic from Router C would never make its way to the ASA, unless you manually instructor the Router using a Static Route to point the 72.3.4.55 IP address to the Firewall's actual interface IP.
There is more details of this here: http://www.practicalnetworking.net/series/arp/proxy-arp/#proxy-arp-nat
In your example, the reason it works is because the Router already knows the Firewall's Interface IP's MAC address. So when you do your Static PAT statement in the inverse direction, the upstream router already knows how to get the packet to the Firewall.
If you had a Static PAT for another IP address, that wasn't your Firewall's interface IP address, but was on the same IP network as your Firewall's interface IP... then you'll find that configuring it in the order of "Mapped/Real" would not work.
Thanks for reminding me why I normally don't do the (outside,inside) path. I didn't think about Proxy ARP honestly.
However you didn't address my second point, or misunderstood. Let's go back to your guide:
At the very beginning of the document you explain the difference between Real (R) and Mapped (M): "Another way to remember it is the mapped attributes only exist because the ASA created them, whereas the real attributes exist despite any configuration on the ASA."
However, this seems to be at odds with the Twice NAT example:
Your "plain English" explanation of the line is correct, however it doesn't match up with your original R/M explanation at the beginning of the document, because in this case CORP-DNS is the device that really exists and GOOGLE-DNS is the mapped address that the ASA makes up.
Also consider my own example from above (the inside,outside variant) :
nat (inside,outside) source static targetserver_real-192.168.10.3 interface destination static mgmtclient_mapped-192.168.10.95 mgmtclient_real-10.10.10.10 service src-https src-1443
As you can see with my deliberate naming, this sentence of yours is wrong: "In all cases, the real attributes are being translated to their mapped counterparts. The order of the items in the manual NAT statement remains constant: Always real, then mapped."
Unfortunately, the Manual NAT order is src_R src_M dst_M dst_R port_R port_M
It makes sense (because this allows your plain English sentence to map this well to the Manual NAT entry), but kills your sentence and your colored highlighting for Manual NATs, unfortunately.
Or am I missing something again?
What is supposed to replace the ASA? Hopefully not meraki...
Palo Alto if you're smart :-O
I'm wrapping up my CCNP and going to dig into some sort of security appliance next. I see a bunch of job postings with security experience required and Palo Alto gets mentioned a lot. Seems like they're the new hot thing. You think it lasts? Fortinet has great resources towards their certification track and that's what I was leaning towards. Is PA better?

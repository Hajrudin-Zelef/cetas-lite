---
id: collect-261001-huawei/huawei/questions-242731-seeking-advice-guidance-on-configuring-vpn-to-ubiquiti-usg-behi-478cdd55
title: "questions-242731-seeking-advice-guidance-on-configuring-vpn-to-ubiquiti-usg-behi-478cdd55"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-242731-seeking-advice-guidance-on-configuring-vpn-to-ubiquiti-usg-behi-478cdd55.md
source_anchor: ""
source_lines: [1, 23]
sha256: 69b6bd74ef32979ab733f666562a8e011214f45193e6351384397e446d9138cf
---

# questions-242731-seeking-advice-guidance-on-configuring-vpn-to-ubiquiti-usg-behi-478cdd55

Abstract:
Access the USG network via VPN, through a Meraki MX84.
Summary:
I am currently working on setting up some new hardware and re-configuring a network that I am responsible for.
The objective is to have an individual VPN into the USG network.
The USG is on it's own network behind a Meraki MX84.
The USG is connected to the MX84 via a VLAN configured port (configured within the Meraki Dashboard).
The USG Network can be accessed when connected to the local network that the Meraki is connected to, and can also be accessed when VPNing to the Meraki,
but
I seek to allow a user to VPN directly to the USG network through the MX84, without having to VPN to the Meraki first (without having to double VPN).
Ramblings:
I must admit that this is new to me. I have a general understanding of networking, but a lot of this in unfamiliar grounds.
On the USG side, there are two settings for a VPN (well, three actually, but one doesn't work with this): Remote VPN and Site-2-site VPN.
I have two different thoughts about this,
I could setup the USG with a Remote VPN and have those seeking to connect with the USG be pushed/ported through via rules on the Meraki,
or
I could link the Meraki and USG via a site-2-site VPN connection via a VLAN.
My Troubles:
I'm not sure which is the best approach.
Also, all of the VPN clients I have used ask for an IP or Name Server (DNS, Domain Name, why is there no standard?), yet, from my understanding, we, via our ISP, are only given one IP address, and the Meraki itself can be given a Domain Name that matches that IP address (as it is set in the settings of the Dashboard), so I am not sure how a remote user can actually get to the USG via a VPN client if the USG isn't actually visible on the internet.
Another thing, once a remote user connects to the Meraki, how does the Meraki know which data to send to the USG VLAN if all the data is coming from one source without anything distinguishing it from the other data? I recognize that there is a chain in the way data moves through routers and the such, but from my position, how am I to tell my simple VPN client "Go to [IP], and THEN go to [IP]."
I assume there are rules that can be set within the Meraki that will sort all of this out.
I assume I am thinking too hard about this, and/or don't have enough experience.

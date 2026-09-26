---
id: collect-260926-rattrapage/rattrapage/r-mikrotik-comments-tylqlp-are-fastrack-and-wireguard-mutually-exclusive-4f4883f7
title: "r-mikrotik-comments-tylqlp-are-fastrack-and-wireguard-mutually-exclusive-4f4883f7"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/lot-rattrapage/servers-reviews/r-mikrotik-comments-tylqlp-are-fastrack-and-wireguard-mutually-exclusive-4f4883f7.md
source_anchor: ""
source_lines: [1, 22]
sha256: f38048447761f6265335a88672c6d44e7828ba65e43e85bda809ee10693e2e4d
---

# r-mikrotik-comments-tylqlp-are-fastrack-and-wireguard-mutually-exclusive-4f4883f7

Are Fastrack and Wireguard mutually exclusive? 
        
    I have 1G/35M at home, but have a relative with symmetrical 2G. I duplicated my QNAP box and hooked it into a hEX at their place, then set up two WG tunnels between it and my 5009; traffic to .124/24 on my end ends up on their end, traffic to .126/24 on their end ends up on mine.
But I also expected to do QNAP RTRR backups and NFS mounts from them to me and only getting ~175Mbits/sec thruput, and was surprised at such a slowdown, also verified via iPerf3. One thing I noticed in /ip/firewall/connections is there's no Fastrack on either side of this tunnel. If I dst-nat a hole from my side's external IP to the iPerf server on their side (i.e., bypassing the local-network tunnel) speeds are in the ~750Mbit range, as expected, so it's not a thruput issue per se.
So are FT and WG mutually exclusive? Is that kind of limited speed expected, and is that due to the WG encryption, or lack of FT? Both boxes are hovering at > 50% when I'm sending a lot of traffic
Section des commentaires
Wireguard does not have offloading support, and probably won't without a hardware upgrade. Wireguard requires ChaCha20 for encryption, while mikrotik hardware acceleration is limited to various flavors of AES. If you terminate the wireguard tunnel on the QNAP, speed will most likely improve.
Thanks everyone. One thing I've done is put in a dstnat rule on the remote hEX for my home's external IP on the RTRR port to the QNAP, which gives me ~750Mbit/sec to my home QNAP. I don't care for that workaround (i.e., if my home's WAN IP changes) as it seems hacky, but it'll do for now.
I'd like to do the same thing for Plex, but that appears to be a different issue, as the remote Plex is seeing my server at home via the tunnel somehow (which I'd think wouldn't be happening, as it's two different networks) so when I stream from their box it goes over the tunnel (when I want it to go over the WAN).
Yeah, QNAP security is "submarine-screen-door" quality- that thing's totally behind the firewall (with the exception of Plex, listening on another port) and I don't dare expose anything else.
In case you haven't discovered it yet (and for the benefit of others finding this), you can add a dns entry to an address list and it will resolve and keep updated. You can use cloud dns (or any other ddns service) for this, and make rules that keep working even if your ISP changes your IP.
hows the cpu usage looking on the hEX while you're maxing out what you can get from the tunnel ? I'm guessing it's hitting 100% long before the 5009?
The hEX has IPsec hardware acceleration ~450 Mbps- I'm not sure if wireguard can leverage that?
Maybe you can try create firewall rule to fasttrack the wireguard traffic? AFAIK all that fasttrack does is bypass the firewall, so the mangle rules etc. I don't think wire/switch speeds are possible with any VPN because it needs to be processed by the CPU. wireguard is probably the fastest in general because it's implemented at kernel level. But maybe with the hEX you could get more performance using IPsec thanks to the built-in hardware acceleration for that? I'd be interested to know also, as I'm deploying hEX's for remote CCTV locations and was also planning on using wireguard for the vpn side of things
you can try overclock the hEX in routerboard settings, which should also increase throughput. it should be safe but obviously you're on your own at that point
hopefully someone more clued up about the fasttrack question can shed some more light here
I haven't checked but I wouldn't be surprised. I may try and get another 5009 next time I'm at their place (it's ~1500 mi away).
I don't think I could set that up reliably, though. I don't really know how it works.
How are you getting 175 mbps with WG on a HEX? IIRC, I was only getting ~40 mbps with WG between a HEX and a CCR2004.
With GRE/IPsec, I'm only able to get 100 mbps.
Commentaire supprimé par un membre de l’équipe de modération
Why not use ipsec?

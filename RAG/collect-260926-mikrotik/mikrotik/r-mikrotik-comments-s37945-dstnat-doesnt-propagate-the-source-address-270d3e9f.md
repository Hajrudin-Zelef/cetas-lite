---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-s37945-dstnat-doesnt-propagate-the-source-address-270d3e9f
title: "r-mikrotik-comments-s37945-dstnat-doesnt-propagate-the-source-address-270d3e9f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/r-mikrotik-comments-s37945-dstnat-doesnt-propagate-the-source-address-270d3e9f.md
source_anchor: ""
source_lines: [1, 32]
sha256: 041c56bf580760a613582889a22eaf31e6f3996ec0557b168d166f6982558fd8
---

# r-mikrotik-comments-s37945-dstnat-doesnt-propagate-the-source-address-270d3e9f

Dstnat doesn't propagate the source address 
        
    Hello,
      I've configured my NAT Firewall configuration with dst-nat :
    
chain=dstnat action=dst-nat to-addresses=192.168.***.*** to-ports=80 protocol=tcp dst-address=176.***.***.*** dst-port=80 log=yes log-prefix="NAT-NAS - "
      All internet connexion to my public ip on port 80  and 443  (not shown) are forwarded to my NAS hosting different applications. My NAS is set has a reverse proxy based on the subdomain source of the request to redirect through the right application.
    
      The problem I have is that all packet from internet that my NAS received are from my router 192.168.**.1 . And when my application want to ban an IP (for wrong password for example) it'll ban my own local router IP.
    
      I don't understand why this happening (maybe it's caused by something else) because from the mikrotik wiki the dstnat should only edit the destination address and not the source address :
    
      destination NAT or dstnat. This type of NAT is performed on packets that are destined to the natted network. It is most comonly used to make hosts on a private network to be acceesible from the Internet. A NAT router performing dstnat replaces the destination IP address of an IP packet as it travel through the router towards a private network.
Thanks a lot for your help
Section des commentaires
Dst-nat doesn't do that and I am pretty sure this is not some random bug.
You almost certainly unintentionally trigger another rule with src-nat as well. If you share rest of your firewall NAT rules, (obviously redact private details) it should be easy to point out
Thanks I made and export :
I was trying to debug and try to move the hairpinning (first item right now) to the top. It was after my
NAS - SSH / SFTP. I also disable all my VPN configuration to try.
It's working right now but I'm not sure was. In the precedent configuration it's possible it was juste something like <public_ip> → router → change src → <nas>
Could you enlighten me why just moving my hairpinning to the top help to fix the issue ?
Have you added the field source address for some hairpin nat rules? I think it masquerades all packets from the internet because they were matched by some hairpin nat rules without specific source address.
Post your config with /export hide-sensitive. But it sounds like you have a masquerade rule that also matches.
Thanks, it could be my hairpinning that I just move to the top the issue. I post my config here https://www.reddit.com/r/mikrotik/comments/s37945/comment/hsjcqnu/?utm_source=share&utm_medium=web2x&context=3
Another reason you should not be using hairpin NAT.
What is the best way to connect to LAN device using external IP (via DNS) ?
... not doing it
split horizon dns is your friend ( and with a reverse proxy it's even easier)
assuming the dns ip used by your LAN devices is the routeros device, add a static entry ( /ip/dns/static ) pointing the "public" sites to the reverse proxy internal ip address.
external visitors will resolve your host names to your public IP and will reach your reverse proxy though the dstnat rule
LAN devices will resolve to the host names to the internal ip of your reverse proxy

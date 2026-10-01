---
id: collect-261001-general-networking/general-networking/questions-737834-cant-rdp-to-work-desktop-over-vpn-by-name-e31789fe
title: "questions-737834-cant-rdp-to-work-desktop-over-vpn-by-name-e31789fe"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-737834-cant-rdp-to-work-desktop-over-vpn-by-name-e31789fe.md
source_anchor: ""
source_lines: [1, 17]
sha256: 4ad0152cf36f81a3cb8d9d915c9b4e5ea927e530de46205a52d0d9c30ca6d1b8
---

# questions-737834-cant-rdp-to-work-desktop-over-vpn-by-name-e31789fe

I have a couple of users that are having the same issue. They can VPN in to our local site but once the VPN is established they can only RDP to there desktop using the IP address. I have had a user test this and when they ping their work desktop from home, after the VPN connection is established, by IP address she gets a good return, but by desktop name she gets a "Request timed out." return and I noticed the address it has in brackets next to the name is nowhere near correct [198.105.254.63]. We have our local network set up as 10.101.x.x. Any idea why this could be and how I can fix this? The DNS IP's she is getting, after the VPN is connected, are correct.
1 Answer 1
I guess your VPN clients are not using your Internal Office DNS servers when connected, and that's why can not resolve the internal hostnames correctly.
You need to make sure that you configure your VPN server so that the clients, when connected also recieve the internal namservers (those you use in Office lan). Then they will be able to resolve the names properly.
The hostnames that are used in LAN environment are not globally resolvable as they used some private ip addresses which are also not globally routable. The host/ipaddress information are configured and stored in local nameservers and that's why when using some other DNS server outside the LAN they are not resolvable.
- 
        I have had her check this with an ipconfig /all command in her cmd prompt and it appears the DNS server she is getting from the VPN connection is correct.Eddie Studer– Eddie Studer2015-11-20 16:04:44 +00:00Commented Nov 20, 2015 at 16:04
- 
            
            
- 
        Can you also chek if they are used?Diamond– Diamond2015-11-20 16:28:48 +00:00Commented Nov 20, 2015 at 16:28
- 
        Sorry for the late response but yes you were correct. We were able to get her into her machine by IP.Eddie Studer– Eddie Studer2015-12-16 21:39:59 +00:00Commented Dec 16, 2015 at 21:39
- 
        
nslookup IPADDRESSwhere IPADDRESS is the IP address of the computer you attempting to remote into?

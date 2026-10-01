---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1pxzy0n-best-methods-to-troubleshoot-ipsec-vpn-connections-92fa5478
title: "r-fortinet-comments-1pxzy0n-best-methods-to-troubleshoot-ipsec-vpn-connections-92fa5478"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1pxzy0n-best-methods-to-troubleshoot-ipsec-vpn-connections-92fa5478.md
source_anchor: ""
source_lines: [1, 23]
sha256: e1babb07e44b91da844c8a27420184197972f7792e9b36eb037bdd9aed1dd9b7
---

# r-fortinet-comments-1pxzy0n-best-methods-to-troubleshoot-ipsec-vpn-connections-92fa5478

Best methods to troubleshoot IPsec VPN connections to diagnose remote access user problems? 
        
        
        
    
    
    We have a Remote Access IPsec VPN along with FortiClient EMS which is what users use to remotely access our network. We have about 5 or so consistently remote users and then everyone is just a random mix of on-site and remote throughout the year. Most users have no issues but every once and a while I have a random user that starts complaining that they keep getting disconnected from the VPN. Most of the time it ends up being their home Internet or Wi-Fi router or other environmental thing on their side, it just usually takes a while to prove it to them.
This particular person I was helping last week has been disconnected from the VPN every 10-20 minutes or so. It happened to them about 5-10 times two days last week. I tried to get them to connect their laptop via physical cable but they said their Wi-Fi router was in the basement and that they use a booster to get the sigal to the rest of the house. Of course that's sign #1 that the issue is on their end.
Anyway, from my side, I'm just wondering what are some common CLI commands I can run to diagnose the exact connection issues (if possible) and what all to look for.
Currently I have been using these commands from my notes:
Debug IPsec VPN issues: (IP Address is the remote user's IP)
diagnose vpn ike log filter rem-addr4 [IP ADDRESS]
diagnose debug app ike -1
diagnose debug enable
Troubleshooting SAML:
diagnose debug console timestamp enable
diagnose debug application samld -1
diagnose debug enable
diag sniffer packet any ' host [IP ADDRESS] ' 4 0 l
Sniffer:
diag sniffer packet any ' host [IP ADDRESS] ' 4 0 l
The difficulty is that there's a lot of output and I'm not sure what kinds of things to look for.
Section des commentaires

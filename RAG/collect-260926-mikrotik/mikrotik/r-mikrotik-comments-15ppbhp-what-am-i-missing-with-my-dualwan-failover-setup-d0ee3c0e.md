---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-15ppbhp-what-am-i-missing-with-my-dualwan-failover-setup-d0ee3c0e
title: "r-mikrotik-comments-15ppbhp-what-am-i-missing-with-my-dualwan-failover-setup-d0ee3c0e"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-15ppbhp-what-am-i-missing-with-my-dualwan-failover-setup-d0ee3c0e.md
source_anchor: ""
source_lines: [1, 39]
sha256: 23124f62b8e37b7cfea16c04b9c84804acff248ee00c4dc3bdf88e444ab42df9
---

# r-mikrotik-comments-15ppbhp-what-am-i-missing-with-my-dualwan-failover-setup-d0ee3c0e

What am I missing with my dual-WAN failover setup? 
        
    Hi, I want to have dual-WAN failover (main fiber on ether1, secondary LTE modem on ether5) and found a script which does exactly what I need, however I'm not sure I set it up right: https://raw.githubusercontent.com/loiklo/mikrotik/main/failover-dhcp
I started out with the default single-WAN setup on ether1.
- 
      I removed ether5 from the bridge
- 
      I added a DHCP client on ether5 (the failover) and checked that it got a lease
- 
      I removed the ether1 cable, then ran the script but the failover did not happen
- 
      I tried adding ether5 to the WAN interface list, which worked but I think it only worked because it was considered to be like ether1. So no proper failover but another route to the internet (purely guessing here).
- 
      I undid that and tried disabling "add default route" on both ether1 and ether5 dhcp connections and running the script but not much luck either
- 
      The only other thing I did was add a similar masquerade rule to NAT where srcnat: local lan, out interface. ether5, but still no luck.
I'm not sure what the other steps involved for this setup to work are. If there are other firewall rules, routing rules I need to add before this can work.
Any tips are welcome! Thanks!
Section des commentaires
I used the netwatch tool in routerOS.
I created a route for a public ip, say clouding 1.1.1.1 to use the gateway on the fiber wan only with admin distance 1. Now traffic for this 1.1.1.1 ip will only use this route over the fiber.
Add your two default routes, add a comment of fiber for the fiber route and lte for the second route. If you want the lte route to only be used when fiber fails, then give the lte route a higher administrative distance.
Go to netwatch under tools, add the following new rule
I used the icmp ping command and put the host 1.1.1.1, the one I had going only on the fiber route above. Timeout of 1 second.
Under the up section, this is when the ping is successful and fiber route is active I put
/ip route enable [find comment=fiber]
Similarly, in the down section, this is when the ping fails so I know internet access on fiber isn't possible
/ip route disable [find comment=fiber]
This will disable the lower admin distance fiber route and let the higher admon distance lte route take over
I ended up going with your solution, thank you!
I’ve published my script on here earlier:
https://www.reddit.com/r/mikrotik/comments/115k4g8/part_2_script_for_rb5009_initial_success_with/?utm_source=share&utm_medium=ios_app&utm_name=ioscss&utm_content=2&utm_term=1
I found this helpful:
https://www.youtube.com/watch?v=iA3yDMDZ-20
No scripting required in this tutorial. Pure UI stuff.
Have a look at https://thebadjr.com/ Customer mikrotiks that has that functionality
Commentaire supprimé par un membre de l’équipe de modération
I followed this video https://youtu.be/eTmpBAAW_pQ on router os 7 it worked great without any problems
This video looks great but it seems that you know the gateway IP for setting up the routes beforehand, and mine can change

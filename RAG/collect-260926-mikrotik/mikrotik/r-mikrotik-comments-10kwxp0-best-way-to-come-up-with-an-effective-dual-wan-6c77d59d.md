---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-10kwxp0-best-way-to-come-up-with-an-effective-dual-wan-6c77d59d
title: "r-mikrotik-comments-10kwxp0-best-way-to-come-up-with-an-effective-dual-wan-6c77d59d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/r-mikrotik-comments-10kwxp0-best-way-to-come-up-with-an-effective-dual-wan-6c77d59d.md
source_anchor: ""
source_lines: [1, 16]
sha256: e9671853fde5515a8f35b65ebb48db3eaa4930380b7633ffca03058271d29e70
---

# r-mikrotik-comments-10kwxp0-best-way-to-come-up-with-an-effective-dual-wan-6c77d59d

Best way to come up with an effective dual wan failover 
        
    Hi,
My MK device still runs ROS V6, but I want to upgrade to V7 shortly.
It seems that a few things have changed with V7, and I am worried that my failover setup isn't going to work anymore.
So, I need to make some changes apparently. I don't know which method is the best to come up with an effective failover.
on the Mikrotik wiki, I found this one which uses mangle :
https://help.mikrotik.com/docs/pages/viewpage.action?pageId=26476608
but I see a different way in this youtube video as well. it is very similar to the old one in my ROS v6
https://www.youtube.com/watch?v=eTmpBAAW_pQ
In your opinion, which one is the best? Thanks
Section des commentaires
the easiest one is using recursive routes, I've had it working on ROS7 (but I don't have that config anymore). Migration was easy, I had to change scope/target scope for some routes to have it working after upgrade (EDIT: I think this was because I was doing LB too). There is new functionality in netwatch worth checking too, you can combine it with mangle to configure your own custom failover.
I gave the second way (that one in the youtube video above) a go in a virtual environment (GNS3) and it works perfectly.
I'd like to try the first one too, but I need to learn and understand better how mangle rules along with the route setup work first.
Thanks

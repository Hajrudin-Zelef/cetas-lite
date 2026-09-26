---
id: collect-260926-mikrotik/mikrotik/questions-928501-how-can-i-access-webfig-of-my-mikrotik-routerboard-57358ebc
title: "questions-928501-how-can-i-access-webfig-of-my-mikrotik-routerboard-57358ebc"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2015-05-18"]
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-928501-how-can-i-access-webfig-of-my-mikrotik-routerboard-57358ebc.md
source_anchor: ""
source_lines: [1, 26]
sha256: 49c00868a39fb0ee289ee79de828cd3dc843a2cc190e3c3e3ec1f52595b5aa2f
---

# questions-928501-how-can-i-access-webfig-of-my-mikrotik-routerboard-57358ebc

To connect to web config on mikrotik routerboard I need the ip address of my routerboard but the problem is that I don't know the ip address of my routerboard. Is there anyway I can get the ip address of my routerboard in other to connect to the webfig?
- 
        1Look into the docs, there should be a default address be listed (possibly 192.168.88.1).Sven– Sven2015-05-18 14:31:14 +00:00Commented May 18, 2015 at 14:31
- 
            
            
- 
        Have tried 192.168.88.1 n I checked d docs but nothing has happenedToye_Brainz– Toye_Brainz2015-05-18 14:40:34 +00:00Commented May 18, 2015 at 14:40
- 
        How is your PC's network interface configured and how do you connect it to the device?Sven– Sven2015-05-18 14:41:49 +00:00Commented May 18, 2015 at 14:41
- 
        I connect to the device through an application called winbox using ethernet port(cable). The winbox searches for the routerboard device then I click connect. A new app will pop up which allows me to configure my routerboard. The first thing I did on the routerboard was to reset the configuration. I don't know if this has in anyway affect the ip address of my routerboardToye_Brainz– Toye_Brainz2015-05-18 14:49:36 +00:00Commented May 18, 2015 at 14:49
                    
                        Add a comment
                    
                 | 
            
                
            
        
         
    1 Answer 1
You can find out the IP address with the program winbox,if you are connected to a port where discovery is enabled.
http://www.mikrotik.com/download (under "Useful tools and utilities")
Look at the tab "neighbors" - (in the older version of winbox - version 2 - it's the button with three dots)
If the IP is in a different subnet then you can connect via mac address

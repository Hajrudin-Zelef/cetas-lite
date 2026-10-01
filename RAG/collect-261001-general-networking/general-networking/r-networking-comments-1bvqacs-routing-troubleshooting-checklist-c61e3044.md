---
id: collect-261001-general-networking/general-networking/r-networking-comments-1bvqacs-routing-troubleshooting-checklist-c61e3044
title: "r-networking-comments-1bvqacs-routing-troubleshooting-checklist-c61e3044"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/r-networking-comments-1bvqacs-routing-troubleshooting-checklist-c61e3044.md
source_anchor: ""
source_lines: [1, 21]
sha256: 37178a6a7906152f290edb0b0448ef5b6ad4b3485a0969d75fe0a74a3d27d0db
---

# r-networking-comments-1bvqacs-routing-troubleshooting-checklist-c61e3044

Routing Troubleshooting Checklist 
        
        
        
    
    
    Hi all,
has anyone come across on the internet to any checklist that can help the networking troubleshooting process become smoother? Especially to those who do not have to deal with that on a daily basis?
Something like the main culprits to watch for, for specific routing protocols.
When troubleshooting…
BGP, look for this first, then this… and/or not necessarily in order RIP, look for these first… OSPF…
etc…
Section des commentaires
IMHO the best aid to troubleshooting is to know the network. Do you have a map? Can you follow the path of data flow from one device to the other? Example can you follow the path of a word document from the sub folder in the file share on the server, to the workstation, to the print server and finally to the network printer? Do you know the IP address of each device in that path? What switch is that server connected to? Whats the IP address of the switch? What the port number is the server using on that switch? Same for the printer and the workstation. To me if you don't understand the network you can't know how find what is broken. So I always start with a network map.
Smoother than what? The complexity and "focus" of networking issues varies wildly between networks. A general checklist would include so many steps that might never be used. Like what would this list include? 50 show commands for every single protocol that the troubleshooting person with the list wouldn't understand anyway? Troubleshooting is something that is often based on pure experience with the affected technology. But generally it's a good idea to work the OSI model bottom to top while even that can be a huge waste of time. Just trying things due to lacking experience is always not smooth 😅
If you don't work on it on a routine basis then you lose the core fundamentals of troubleshooting. It's hand in glove.
At a high level: Can you ping peer interfaces? Are interfaces set to passive (not advertising and able to form an adjacency). Are networks and interfaces actually being advertised? Are the timers the same for both ends. Are the interfaces actually up? Are the correct with compatible IP and Subnet Mask?
If you are green on routing protocols this is about all that I can offer. If you need to get past that you need to become educated in the protocols you will be working with.
The book "Troubleshooting IP Routing Protocols" is starting to show it's age, but is still good. "Troubleshooting BGP: A Practical Guide to Understanding and Troubleshooting BGP" is newer. The CCNP used to have a separate troubleshooting focused test. Any of the old TSHOOT study guides are still good and will cover more than routing protocols.
I've got one in my head. You'll benefit and retain more if you build it yourself. Make up scenarios in your head. Lab them up if needed
One example... Two BGP neighbors aren't forming. Think about what would prevent that from happening. First thing I might do is look in the logs. I might then compare the configs. Id then verify the peers can ping each other. Maybe then id confirm nothing is blocking TCP 179, and so on

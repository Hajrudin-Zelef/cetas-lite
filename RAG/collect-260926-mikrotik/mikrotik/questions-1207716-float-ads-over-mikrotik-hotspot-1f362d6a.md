---
id: collect-260926-mikrotik/mikrotik/questions-1207716-float-ads-over-mikrotik-hotspot-1f362d6a
title: "questions-1207716-float-ads-over-mikrotik-hotspot-1f362d6a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2017-05-10"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-1207716-float-ads-over-mikrotik-hotspot-1f362d6a.md
source_anchor: ""
source_lines: [1, 16]
sha256: 45cb5ed8d505c1b04c1fa6bb0397a9863c6e9eeb1a11f715db185eb192defeed
---

# questions-1207716-float-ads-over-mikrotik-hotspot-1f362d6a

i have recently installed mikrotik hotspot with my own WiFi portal for ads purpose and i am successful to preview ad on splash page(alogin.html) by injecting my ads script but is it possible to show float advertisement in every webpage which user visits?
- 
        1While you might be able to do that did you think about the user experience?Seth– Seth2017-05-10 05:51:11 +00:00Commented May 10, 2017 at 5:51
                    
                        Add a comment
                    
                 | 
            
                
            
        
         
    1 Answer 1
No, you can't do this.
Major providers are moving to HTTPS to prevent this kind of behaviour - the information your router sees is encrypted - thus the large sites most people spend most time engaged in will lock this kind of behavior out.
And if you think about it, thats the way it should be - an end user should not trust the connection they go through, and thus people serving data will try and protect against the man in the middle.

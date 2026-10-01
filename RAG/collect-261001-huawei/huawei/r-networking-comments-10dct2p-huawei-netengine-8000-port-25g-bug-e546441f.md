---
id: collect-261001-huawei/huawei/r-networking-comments-10dct2p-huawei-netengine-8000-port-25g-bug-e546441f
title: "r-networking-comments-10dct2p-huawei-netengine-8000-port-25g-bug-e546441f"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/r-networking-comments-10dct2p-huawei-netengine-8000-port-25g-bug-e546441f.md
source_anchor: ""
source_lines: [1, 21]
sha256: a0488d26981e173ff4668e30dc70dac1a2870eaac48692542f3f51aeae799650
---

# r-networking-comments-10dct2p-huawei-netengine-8000-port-25g-bug-e546441f

Huawei NetEngine 8000 Port 25G Bug 
        
        
        
    
    
    Helly Guys, the first time I have experience with Huawei NetEngine which is seemed to be abnormal.
The 25G slot and 25G module is no issue during testing service, but once we shut down the service, we saw that there is something wrong with one port which is Status inside switch is UP/UP and also the physical slot show the green light. The strange thing is with ETH-TRUNK showed both port is down.
Have anyone experience the kind of issue? or this is a bug with version of Huawei? Thanks guys.
      Solved: Inside the interface there was a cmd "hold-up" meaning it will force the interface to up/up even though there is no connectivity.
I simply just use "undo hold-up" and everything is great.
    
Section des commentaires
You should open a case with support. In my experience they are very fast in helping.
I think I remember some similar issue in a Comware switch, but I can not really remember. I think it was a firmware issue.
I found the issue and solution.
This seem to be hardware issue or software issue as you have experience?
Have you turned it off and on agane?
I am sure it will be solved, but before I shut/unshut I need to understand why it occurred.
Hope this isn't in the USA.
Hope it does not in any part of DC, gladly it's not yet launch into production.

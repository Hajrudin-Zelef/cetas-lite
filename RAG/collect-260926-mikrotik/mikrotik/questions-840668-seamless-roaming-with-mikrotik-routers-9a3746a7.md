---
id: collect-260926-mikrotik/mikrotik/questions-840668-seamless-roaming-with-mikrotik-routers-9a3746a7
title: "questions-840668-seamless-roaming-with-mikrotik-routers-9a3746a7"
domain: mikrotik
role: reference
task: reference
actors: ["Apple", "Intel"]
dates: []
keywords: ["ethernet", "intel"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-840668-seamless-roaming-with-mikrotik-routers-9a3746a7.md
source_anchor: ""
source_lines: [1, 29]
sha256: f7614d40630aa6da8fca2f126fbbf13034181415d9675d982daee122652365ee
---

# questions-840668-seamless-roaming-with-mikrotik-routers-9a3746a7

I need to cower broad area (>1x1km) with WiFi connection that won't be interrupted while clients are moving from one AP to another. And wiring APs with ethernet is strongly undesirable (AP setup is AP+Solar PSU).
Client's TCP connection should be still alive after switching to new AP.
Is Mikrotik MESH capable of solving my issue?
- 
        2True seamless wireless roaming is very difficult to implement. Even in situations where one has a professionally-designed high-end wireless system and full control of all client configuration, roaming still does not happen reliably 100% of the time. Your best bet is to detect connection issues on the client side and retry when appropriate.EEAA– EEAA2017-03-26 15:05:11 +00:00Commented Mar 26, 2017 at 15:05
1 Answer 1
Wireless Roaming happens on the client. The client's device decides when and if it will switch over to a better AP. The AP doesn't play much role in it.
So the level of stability during a switchover to another AP is at the mercy of the client's device and drivers.
I strongly suggest you wire all the APs with ethernet.
You can't have a stable setup by using the same air space for both your clients and your 'backbone' between all the APs. Especially on 2.4GHz which is pretty much full of noise.
Or if you absolutely need to do this without cables you should use a different band for the backbone (ie: 5ghz) and different for the clients (ie: 2.4GHz). That means that the APs you will use must have dual radios to be able to transmit on different bands at the same time.
But, the most reliable and cheapest way to do this would be using cables.
Regarding the TCP connections, as long as you don't do any NAT or any other connection tracking stuff between the AP and the router, and you don't change the client's IP, the TCP connections should continue just fine during transition from one AP to another.  On an iPhone I just tested this and it simply switches over to the new/closest AP. Running a ping while moving it shows zero to 1-3 packets lost during the transition.
At the same time on a Windows laptop with an Intel wifi card it doesn't even migrate from one AP to another unless I force it by disabling/enabling the wifi interface.
So as I said, the wireless roaming is done on the client. The AP can't do much about it.
One thing that Mikrotik can do is set a threshold of minimum signal before forcing a disconnect on a client that way forcing a re-connect on another AP potentially. But this method will probably cause more lost packets. Again it all depends on the client implementation.
I wouldn't use the WDS approach. I just don't see any benefit. Especially if you have to use the same frequency on all APs. That would be kind of disastrous in terms of generated noise between the APs.
It's better to have each AP on a different frequency to utilize the whole spectrum without interfering much with each other.
- 
        Cha0s, my clients is very low power IEEE802.11 b/g/n (2.4GHz SISO,20MHz) ~60Mbit IOT modules, they unlikely to support 802.11r, will seamless roaming work for them or they will drop connections?The Architect– The Architect2017-03-29 22:45:37 +00:00Commented Mar 29, 2017 at 22:45
- 
            
            
- 
        According to this clients will drop connections...The Architect– The Architect2017-03-29 22:50:06 +00:00Commented Mar 29, 2017 at 22:50
- 
        And can I use APs with SINGLE radio but DUAL BAND to separate clients and 'backbone' traffic?The Architect– The Architect2017-03-29 23:20:04 +00:00Commented Mar 29, 2017 at 23:20
- 
        You can't use a single radio device to broadcast on both bands at the same time. It's physically impossible.Cha0s– Cha0s2017-03-31 15:20:22 +00:00Commented Mar 31, 2017 at 15:20

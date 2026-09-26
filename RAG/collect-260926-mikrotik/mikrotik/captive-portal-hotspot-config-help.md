---
id: collect-260926-mikrotik/mikrotik/captive-portal-hotspot-config-help
title: "captive-portal-hotspot-config-help"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "exploit"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/captive-portal-hotspot-config-help.md
source_anchor: ""
source_lines: [1, 57]
sha256: 6ecb1819202d2e797030c1e67b1ebc305b46628a7f44d6d5078ea3a4cb98bf68
---

# captive-portal-hotspot-config-help

Hi, I’m trying to configure the captive portal how I need, but I either can’t work it out, or it doesn’t seem to be supported.

I’m trying to achieve a configuration where each client has a weekly time (and maybe data) limit.

I don’t want to manage user accounts; I want each MAC address to be accounted separately.

So a user should be identified by MAC address, their session terminated after a specified time limit, and reset the counters weeky.

There do seem to be session time limits, but the user can log back in immediately after the session time expires and I can’t see any way to set a delay or session reset timeout.

What is that even for? What’s the point of a session timeout if they can log back in immediately?

Possible? I can’t work it out.

Thanks!

             
            
           
          
            
            
              Anyone? I have a very low-power site where 'tik equipment is in place; I’d rather not install separate equipment to manage a captive portal when the mikrotik solution is so close. It’s just bizarrely missing a couple of key details, or I just don’t know how to configure this. The docs are very thin on this matter…

             
            
           
          
            
            
              Devices can use randomise mac addresses while using a hotspot, you might need to use a login username and password

             
            
           
          
            
            
              That’s fine. It’s not intended to be fool-proof, it’s intended to be simple and minimum friction.

If someone can work out how to exploit the limiters, that’s fine, and power to them! I’m working with a serious low-iq community here.

             
            
           
          
            
            
              Bump!

I still really need help with this; does anybody know anything much about Mikrotik’s captive portal? I gets very little attention…

The thing I just can’t get my head around is how it’s useful for sessions to terminate after some limit conditions (like timeout), but allow the user to re-establish and begin a new session IMMEDIATELY.

How can this be used to limit users access? I feel like I’ve missed something really obvious.

A terminated session should be terminated… and the counters that caused the session to terminate should only reset after some reset condition; ie, given time period, or at a particular moment (midnight daily/weekly)…?

---
id: collect-260926-mikrotik/mikrotik/wireless-isp
title: "wireless-isp"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/wireless-isp.md
source_anchor: ""
source_lines: [1, 81]
sha256: c4e9852bda0ebe52b66654bd9d422902182be21c087a34a632fb15f03caa2cac
---

# wireless-isp

Hi everyone,

I want to start a Wireless ISP and I wanted to know what devices do I need, I have to cover an area of 5-6 km 360 degrees.

I hope that someone will give me some advice what I should do, money in not in question, I just need to know what devices anthenas etc I should use.

Thanks in Advance.

             
            
           
          
            
            
              where r u from? wht kind of landscape r u looking into. how many users u want to add.

there r so many things u need to think before u start a wisp

             
            
           
          
            
            
              The landscape is like a valley, and i’m in the middle of it, so everyone will be able to see me here is a pic of the area i’m planing  to cover.

I’m planing to have about 100-150 users.

             
            
           
          
            
            
              If u want to use 2.4 ghz freq. band. u need 3 radios 433ah with r52h mini PCI, 3 sectorial antenna 120 degree preferably 18dbi a network switch, a bandwidth mgmnt server, a UPS for backup,some cat5 cables and pigtails for interfaces and ofcourse internet bandwidth  

             
            
           
          
            
            
              Thanks mahnet, for the bandwidth manager should I use Radius server like free radius or something similar like that?

             
            
           
          
            
            
              I wouldn’t run the 3 2.4GHz cards on the same board. I have done that and once you have enough customers on each sector you will run into self interference issues.

3 RB411A boards in separate enclosures would work better.

             
            
           
          
            
            
              In such populated area as you show us 2,4Ghz is probably a bad choice for deploying a WISP. Unless you want to serve laptops etc. with a hotspot. For a fixed network it is not the best choice.

Loads of people will have 2,4Ghz wifi routers nowadays and many other devices (mobiles, smartphones, pad’s, remote controls, wireless phones etc. etc) uses the 2,4Ghz band too. So there is a good change your clients will suffer a lot from interferences meaning poor quality connections and many disconnections. Specially because you already have to use each of the basically only 3 available free channels (the others are overlapping) for yourself.

If this Wisp is to be build from zero and you serve fixed (houses, offices etc.) clients you’d better start in 5Ghz (802.11a/n) and with mimo straight away. The investment might be somewhat higher but your network will much better be able to expand and you are able to serve more speed to the clients. Therefore it is much more future proof.

I would start with one tower and 3 sectors, each with a good sector antenna, short cable (losses!) and high power cards.

Then at the clients also units with high power cards and good high gain directional antennas.

Use for each antenna a separate radio with its own routerboard. 3 radio’s in one board is possible but there is some change these cards will create self interferences in the box.

Also use preferably metal boxes for the routerboards and radio’s at both your AP as the client. It prevents interferences of unwanted signals…

If you are buying antenna’s, take a good look at the specs before you decide what you want. Not all antenna’s do deliver what they at first sight promise. You have to look at the spectral diagrams and gains etc.

Cables and connectors are important too. Use good stuff and use the properly and protect them against environmental influences. (Salt, rain, moisture, wind etc.)

Also, depending on several factors, a routerboard can only handle 30 to 50 clients per radio max. We are talking a rb800 in the last example! So be aware that you have room and planned to upgrade you system as soon as you reach saturation levels on the AP’s

Careful planning is a must or you might end up having spend a lot of money and you’re out of business soon again… (the competition is lurking around the corner!)

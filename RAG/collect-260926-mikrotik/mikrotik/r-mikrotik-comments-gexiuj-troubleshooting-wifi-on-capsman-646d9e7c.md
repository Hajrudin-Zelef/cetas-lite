---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-gexiuj-troubleshooting-wifi-on-capsman-646d9e7c
title: "r-mikrotik-comments-gexiuj-troubleshooting-wifi-on-capsman-646d9e7c"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/r-mikrotik-comments-gexiuj-troubleshooting-wifi-on-capsman-646d9e7c.md
source_anchor: ""
source_lines: [1, 44]
sha256: 068ef9c08b86daf21da7a7ac92dc88c5a9ec4a4738b7237915246ce1a708924e
---

# r-mikrotik-comments-gexiuj-troubleshooting-wifi-on-capsman-646d9e7c

Troubleshooting WiFi on Capsman
Hi everyone,
I am trying to troubleshoot an issue on my network that appears to have arisen due to my switch to Capsman (at least whatever channels I have chosen seem to have put it in this state, maybe?).
Behavior: My laptop randomly loses connection or I get packet loss out of nowhere.
Setup: 2x CAP AC (POE Injector) 1x HAP AC2 (capsman) 1x WAP AC (POE Injector)
All 4 are statically set on separate 40mhz (Ce) wide channels
Upon troubleshooting I am noticing 2 of the networks randomly drop signal abruptly from my monitoring tool, my guess is this is what is causing the issue. I dont see anything in capsman logs that would indicate why this is happening, and not all APs have this behavior.
See attached images of the channel graph + signal graphs:
Anyone have some potential guidance on what I could try?
EDIT: Some additional info:
- 
      The 2 devices that show the connection drops are: a CAP AC and a WAP AC.
- 
      I also noticed that the connection drops appear almost at the exact same time, so weird!
EDIT2: My config https://gist.github.com/bryantlee/9d33653f6d6c9d2dc5b8820ae3c6bad7
Section des commentaires
Hi everyone - so I restarted these devices (which I feel like I did multiple times), and this signal drop has gone away. In general things are more stable. I am trying to explain some routine latency spikes that happen when pinging my router/hotspots that do NOT happen on devices that are hard wired into my network.
Every 10 seconds or so I see latency spikes with no load on the connection. Any thoughts here? Happens on all hotspots.
Having the same issue and going nuts already. What made me look into it was that whenever I was trying to view a video i.e. gopro on my NAS, from time to time the video would stutter. Apparently that stutter is caused by such latency spikes.
Looked into any available thread/discussion etc. and tried lots of things. Nothing worked.
I'm on a mac. Eventually the spikes were caused by the location services. If you have such feature in your OS, check it out.
Can you post your setting here? On capsman and on cap
Here it is:
https://gist.github.com/bryantlee/9d33653f6d6c9d2dc5b8820ae3c6bad7
Just a note, I know my 2.4g network is not set up correctly, for now I am just trying to focus on getting 5g network to be reliable
Are you using DFs channels?
I used the image below to choose channels, I am pretty confident I am not:
https://www.ekahau.com/wp-content/uploads/2019/04/Detailed-5-GHz-Channel-Allocations.jpg
Also confirmed here: https://en.wikipedia.org/wiki/List_of_WLAN_channels
If it’s at the same time , seems to me that is radar interference and you are using DFS channels. I will take a look at your conf
I used this following image to avoid DFS channels (I hope), would love confirmation.
https://www.ekahau.com/wp-content/uploads/2019/04/Detailed-5-GHz-Channel-Allocations.jpg
Would recommend configuring the lower 4/5 channels and using them for a while.
I had my CapsMan up in the DFS range, and it would disconnect frequently because of RADAR signals - best to avoid.
I used the image below to choose channels, I am pretty confident I am not:
https://www.ekahau.com/wp-content/uploads/2019/04/Detailed-5-GHz-Channel-Allocations.jpg
The table you reference also would support that the channel I am using should be OK. One note is that one of the networks that are dropping signal are indeed in one of the lower channels that you mention. I dont think this is DFS related but would appreciate confirmation.
Here is the output from my CapsMan channel config. It may help - particularly with the side(extension)-channels you have configured currently.
The above settings have been stable for me for more than 3 months now, so can confirm that they function as they should. Note that some of the configured channels ARE DFS ones, so best to avoid them.
Here, I use channels 52 and 149 to avoid neighbour interferance, and like I said - works well here.
Hope this helps.
Edit: Clarification
Commentaire supprimé par le membre
The tool I’m using shows other networks on the same channel and I think I’m okay here, there’s not much going on if you look at the other networks. That’s what I think at least

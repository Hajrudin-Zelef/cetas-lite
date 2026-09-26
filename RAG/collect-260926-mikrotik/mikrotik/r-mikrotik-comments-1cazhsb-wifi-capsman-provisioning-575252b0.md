---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-1cazhsb-wifi-capsman-provisioning-575252b0
title: "r-mikrotik-comments-1cazhsb-wifi-capsman-provisioning-575252b0"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/r-mikrotik-comments-1cazhsb-wifi-capsman-provisioning-575252b0.md
source_anchor: ""
source_lines: [1, 17]
sha256: 2718da17454fc998bdc98193824630972ddcc843bbc509a76b1b15edd96be505
---

# r-mikrotik-comments-1cazhsb-wifi-capsman-provisioning-575252b0

WiFi / CAPSMAN / Provisioning 
        
    I have some questions regarding the wireless operation of my Mikrotik setup.
I have a RB5009 which is my main router and is the "master" for CAPSMAN. Attached to it are 3 AP's. A CAP AX, a HAP AX2 and a HAP AX3.
When I access my RB5009 through Winbox and go to the registration tab of the WiFi configuration I see the list with MAC addresses and to which AP they are connected. For some reason the CAP AX, although devices are successfully connected to it does not show any data in the Tx and Rx columns. Any ideas on why that is the case?
      The WiFi tab in the WiFi config looks like this:
I limited the 5GHz frequencies because some devices had trouble with frequencies above 5700MHz so I disabled those.
    
I use provisioning to configure the different WiFi networks I have at home.
When I access the CAP AX through Winbox the WiFi tab on that AP looks like this.
There are a few devices that for some reason fail to connect to the CAP Ax although I can't find a particular reason. One of the devices is the work laptop from my wife which is an HP laptop with a Microsoft WiFi adapter. It connects fine to the HAP Ax2 or HAP Ax3 on 5GHz but not to the CAP Ax. My work laptop, a Dell Latitude Rugged connects to all AP's just fine. After I disabled frequencies above 5700MHz my main PC connected just fine to all AP's and so did the Sonos speakers (which seem to have issues with frequencies above 5700MHz as well).
And yes, I am looking into removing 80MHz bandwith channels on 5GHz and 40MHz bandwith on 2.4GHz but as long as it works I would like to keep it this way. I am wondering if disabling DFS channels might be a good idea as well.
Section des commentaires
Do you have matching firmware and wifi package versions on all devices?
There were some peculiarities introduced through the transition to multiple wifi packages and running "old" vs "new" capsman on devices without their own radios... might be related to that. If you're not on latest, an update might fix it - I'm guessing that the missing data is due to some kind of compatibility bug from the cap reporting its info back to capsman.
As for some devices not connecting to that ap, there's basic troubleshooting to go through first - like if you run a wifi scanner app on the affected devices, can you actually see that device broadcasting? What happens if you manually select channels for all of them?
Considering all three devices use the exact same IPQ6010 SoC, they're really the same device as far as main processor and wireless is concerned, so this seems likely to be a configuration or software issue rather than a problem with the cap that isn't present on the other two.

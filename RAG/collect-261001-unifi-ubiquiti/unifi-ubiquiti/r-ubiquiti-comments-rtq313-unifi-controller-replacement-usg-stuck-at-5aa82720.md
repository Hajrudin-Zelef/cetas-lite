---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-rtq313-unifi-controller-replacement-usg-stuck-at-5aa82720
title: "r-ubiquiti-comments-rtq313-unifi-controller-replacement-usg-stuck-at-5aa82720"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-rtq313-unifi-controller-replacement-usg-stuck-at-5aa82720.md
source_anchor: ""
source_lines: [1, 36]
sha256: cc1b52650f8aafe3d3c7f575eb9f232803a853135e781ded49e73e59e97af5a5
---

# r-ubiquiti-comments-rtq313-unifi-controller-replacement-usg-stuck-at-5aa82720

Unifi Controller Replacement; USG stuck at "Adoption Failed" 
        
        
        
    
    
    Followup from my post a few days ago. Couldn't upgrade my CloudKey, so decided to launch a fresh 6.5.55 controller via Docker and manually migrate settings and devices. All the Unifi APs and switch were adopted by the new controller no problem, but the USG failed.
I remember this same problem last time, and my work-around was to reset the USG to factory defaults by holding the reset pin for 10 seconds, then console in and change the LAN (eth1) interface from 192.168.1.1 to a valid IP for my LAN.
I've done that, rebooted the USG a couple times, and also manually set the inform URL, but no luck:
ubnt@ubnt:~$ show interfaces ethernet 
Codes: S - State, L - Link, u - Up, D - Down, A - Admin Down
Interface    IP Address                        S/L  Description                 
---------    ----------                        ---  -----------                 
eth0         100.77.77.78/30                   u/u                              
eth1         192.168.249.1/24                  u/u                              
eth2         -                                 A/D 
ubnt@ubnt:~$ unifi info 
Model:       UniFi-Gateway-3
Version:     4.4.56.5449062
MAC Address: f0:9f:c2:xx:xx:xx
IP Address:  100.77.77.78
Hostname:    ubnt
Uptime:      1317 seconds
Status:      Not Adopted (http://192.168.249.218:8080/inform)
The Unifi console just showed the USG-3P as "failed adoption" with a 192.168.1.1 IP address. There's no option to force re-adoption or modify settings.
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
You have to adopt it on 192.168.1.1 FIRST and then switch to your desired subnet. Won’t work otherwise.
How would the USG communicate with a controller that isn't on 192.168.1.0/24?
You have to move the controller to the 192.168.1.0/24 subnet, adopt, then switch back.
Hi u/greenlakejohnny Thanks for your patience. Please start a LiveChat or a support ticket at help.ui.com so our team can collect more information to properly review and assist.
Thanks, but have decided to sell the USG and put the money towards a new FortiGate. The USG was nice as a visibility tool, but not great a far as bang for the buck.

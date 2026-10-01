---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-18w4okp-very-confused-on-how-to-setup-a-vpn-so-i-can-3e1892d2
title: "r-ubiquiti-comments-18w4okp-very-confused-on-how-to-setup-a-vpn-so-i-can-3e1892d2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-18w4okp-very-confused-on-how-to-setup-a-vpn-so-i-can-3e1892d2.md
source_anchor: ""
source_lines: [1, 37]
sha256: 313bfd5ada630ff21a80e63f3d391c66b94921f90e33d9bd1f0855794e3407fe
---

# r-ubiquiti-comments-18w4okp-very-confused-on-how-to-setup-a-vpn-so-i-can-3e1892d2

Very confused on how to setup a vpn so i can remote desktop to a machine on my ubiquity network from anywhere. 
        
        
        
    
    
    I tried going into the console of my USG and clicked 'vpn' but it keeps warning me that 'wireguard or teleport' are both faster and 'better'.
What I want to do is remote desktop to the one of the machines on my nettwork while i'm on the road. Both machines are using Windows 11.
The documentation on the unifi website assumes I know everything about vpns and doesn't make it easy. I only want to access one machine from another, not setup an 'enterprise' vpn.
Is there a guide somewhere? thank you and happy new year!
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Did you try downloading wifiman and use teleport? Teleport there is no config. Baked into gateway and apps
Thank you. I think I'm a little bit silly because I tried that but it says that it's for mobile? Does it mean we activate it through the mobile phone and then use it on a PC?
You can activate it and then use an rdp or similar client. Or just use TeamViewer for free and install on each machine and have remote control over the machines when you want.
There is a new-for-Windows WiFiMan client out: WiFiman Desktop 0.3.1 for Windows.
Load it on your PC. Set up Teleport (the instructions are easy as you just send yourself a link.)
Connect.
VERY straightforward and it takes zero expertise to set up.
You can also check on YouTube...Crosstalk Solutions has all kinds of useful videos on "how to", one of which is a VPN setup. Also not difficult. (He also has one on setting up WiFiMan.)
Oh very cool! Thank you !
If you have a USG, you can only use L2TP/PPTP. Newer consoles support things like my preference of WireGuard or teleport (which is just WireGuard with a simplified UI)
Ahhh... That's why
Would definitely recommend Tailscale running somewhere within your network. Much easier.
Thank you. I'll give that a try
In the Classic interface you can setup a L2TP/IPSEC VPN Server
Then you can use the default Windows VPN client to connect back
https://lazyadmin.nl/network/unifi-vpn/
Teleport seems to be an "app" that functions like a "Privacy" VPN and not a Remote Access VPN as there is no pure desktop client that you would normally use in business.
That looks like the cleanest way to do it
It really is.
I hope you get VPN working, but want to offer another option: an SSH tunnel.
I have just one port open to an OpenBSD device (runs well on old HW) that I use with WSL2 client on my Win11 laptop and that's all I need for RDP, Home Assistant, etc. On Android, JuiceSSH is the client. It's a simple solution.

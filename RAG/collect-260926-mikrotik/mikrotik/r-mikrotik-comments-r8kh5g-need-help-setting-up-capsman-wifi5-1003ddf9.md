---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-r8kh5g-need-help-setting-up-capsman-wifi5-1003ddf9
title: "dec/04/2021 13:48:24 by RouterOS 6.49.1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/r-mikrotik-comments-r8kh5g-need-help-setting-up-capsman-wifi5-1003ddf9.md
source_anchor: ""
source_lines: [1, 87]
sha256: ff7c122e08e6d8a0354d05d5747273f5def919a8ac12e252b4aff0165b5929cd
---

# dec/04/2021 13:48:24 by RouterOS 6.49.1

Need help setting up CAPsMAN WIFI5 
        
    Hi Everyone!
A bit about my setup:
- 
      ISP Uplink: 1Gbps
- 
      Main Router/CAPsMAN controller - RB3011UiAS
- 
      APs: 2x cAP AC, 1x wAP AC
- 
      PoE: TP Link Gigabit PoE Switch
My CAPsMAN config is as follows:
[bhanu@ORNG-RB3011] > /caps-man channel print
 0 name="WIFI-4" frequency=2462 control-channel-width=20mhz band=2ghz-g/n extension-channel=disabled tx-power=10 skip-dfs-channels=yes 
 1 name="WIFI-5" frequency=5180,5200,5220,5240 control-channel-width=20mhz band=5ghz-n/ac extension-channel=XXXX tx-power=20 reselect-interval=1h 
   skip-dfs-channels=no 
[bhanu@ORNG-RB3011] > /caps-man provisioning print
Flags: X - disabled 
 0   radio-mac=00:00:00:00:00:00 hw-supported-modes=gn identity-regexp="" common-name-regexp="" ip-address-ranges="" action=create-dynamic-enabled 
     master-configuration=ORNG-W4 slave-configurations="" name-format=prefix-identity name-prefix="" 
 1   radio-mac=00:00:00:00:00:00 hw-supported-modes=ac,an identity-regexp="" common-name-regexp="" ip-address-ranges="" action=create-dynamic-enabled 
     master-configuration=ORNG-W5 slave-configurations="" name-format=prefix-identity name-prefix="" 
[bhanu@ORNG-RB3011] > /caps-man config print
 0 name="ORNG-W4" mode=ap ssid="ORNG" country=india installation=any security=ORNG_SG1 datapath=DP1 datapath.bridge=BrCapDP channel=WIFI-4 
   channel.tx-power=0 channel.skip-dfs-channels=yes rates=No B Rates 
 1 name="ORNG-W5" mode=ap ssid="ORNG-5G" country=india installation=any security=ORNG_SG1 datapath=DP1 datapath.bridge=BrCapDP channel=WIFI-5 
   channel.skip-dfs-channels=yes 
[bhanu@ORNG-RB3011] > /caps-man rates print
 0 name="No B Rates" basic=12Mbps supported=12Mbps,24Mbps,48Mbps,54Mbps [bhanu@ORNG-RB3011] > /caps-man export hide-sensitive
# dec/04/2021 13:48:24 by RouterOS 6.49.1
# software id = Z2LV-J30B
#
# model = RB3011UiAS
# serial number = E14E0D75AD86
/caps-man channel
add band=2ghz-g/n control-channel-width=20mhz extension-channel=disabled frequency=2462 name=WIFI-4 skip-dfs-channels=yes tx-power=10
add band=5ghz-n/ac control-channel-width=20mhz extension-channel=XXXX frequency=5180,5200,5220,5240 name=WIFI-5 reselect-interval=1h skip-dfs-channels=\
    no tx-power=20
/caps-man datapath
add bridge=BrCapDP client-to-client-forwarding=yes local-forwarding=yes name=DP1
/caps-man rates
add basic=12Mbps name="No B Rates" supported=12Mbps,24Mbps,48Mbps,54Mbps
/caps-man security
add authentication-types=wpa2-psk disable-pmkid=yes encryption=aes-ccm group-encryption=aes-ccm group-key-update=1h name=ORNG_SG1
/caps-man configuration
add channel=WIFI-4 channel.skip-dfs-channels=yes channel.tx-power=0 country=india datapath=DP1 datapath.bridge=BrCapDP installation=any mode=ap name=\
    ORNG-W4 rates="No B Rates" security=ORNG_SG1 ssid=ORNG
add channel=WIFI-5 channel.skip-dfs-channels=yes country=india datapath=DP1 datapath.bridge=BrCapDP installation=any mode=ap name=ORNG-W5 security=\
    ORNG_SG1 ssid=ORNG-5G
/caps-man manager
set ca-certificate=auto certificate=auto enabled=yes upgrade-policy=suggest-same-version
/caps-man provisioning
add action=create-dynamic-enabled hw-supported-modes=gn master-configuration=ORNG-W4 name-format=prefix-identity
add action=create-dynamic-enabled hw-supported-modes=ac,an master-configuration=ORNG-W5 name-format=prefix-identity
Now, Coming to the problem: I only get around 200/200mbps over Wifi 5.
I've tried using a netgear router directly connected to ISP uplink. it gives me ~650mbps over its wifi 5. I'd expect the same or better from mikrotik but sadly, that's not happening in my case.
connecting to the router directly gives me the full bandwidth (~950mbps) so its safe to say the router is not causing the throttle. I've checked, and checked again and then checked once again my entire config against the recommended best practices for CAPsMAN and haven't found a config that allows would work for my case.
From my reading of the mikrotik forum, what I've understood is that there is some driver/kernel level issue in 1st gen Wifi 5 product that mikrotik hasn't solved which causes poor performance but I'm not sure how much ground that holds.
I'd appreciate if someone could basically do some hand-holding with CAPsMAN wifi 5 configuration to get me to a more bearable speed. I'm seriously at a point where I have no clue about what is causing the issue.
Thanks in advance!
Ps:I'm seriously considering replacing CAPsMAN for some other solution but due to the ongoing global shipping issues and chip shortage, all my local vendors prices have skyrocketed. I'm waiting for the prices to stablize in the market before committing to an expensive wifi upgrade.
Section des commentaires
Try to use only ac in wifi5
Just did, sadly didn't make a dent in my situation. :(
Troubleshooting device TX/RX-RATE in registration table
I have no clue about those but here is a screenshot if it helps point out something.
link
I ran into the same issue as you before...
here are my findings:
https://forum.mikrotik.com/viewtopic.php?t=161282
IMHO: Do not use CAPsMan for 2 APs......
I do not use CAPsMan event for 30 APs, because the users complain about bad performance and unstable connections. If I run a standalone config with the same config, the users are happy.
The standalone configs drawback is, that you can not see the connected devices in one place.
My rationale for using CAPsMAN or any other Central management solution is that I want to make this system easily extendable. As of today, I have just 3 APs but I Plan to deploy more. Assuming that I don't find a solution for my woes, I'm confident I'll just switch over to ubiquiti or aruba when the market stablizes.
I am NOT a Mikrotik pro, not at all, but back when I was labbing with CAPSMAN, I found this video extremely useful:
https://www.youtube.com/watch?v=JRbAqie1_AM&ab_channel=MikroTik
EDIT: If you ever finetune it to your liking, please share it with the group :)
Good luck mate!
I've been recommended that video more than once. It talks a lot about Wifi 4 (2.4Ghz) and I really don't have any issues there.
The problem is with Wifi 5 (5Ghz) and it is still broken. The lack of information about 5Ghz best practices and troubleshooting is the motivation behind my post.
And if I find a solution, I'll definitely share it here. That's for sure.
I could get up to 340Mb out of my cAP, following the simple CAPsMAN setup from mikrotik wiki, make sure to enable local forwarding.
I stopped chasing maximum speed on wifi as the clients who actually need that sort of speed are anyways hardwired, I suggest you do the same; also my expereience with mikrotik wifi is ..weird see here.
Commentaire supprimé par un membre de l’équipe de modération
I just added to OP
Commentaire supprimé par un membre de l’équipe de modération

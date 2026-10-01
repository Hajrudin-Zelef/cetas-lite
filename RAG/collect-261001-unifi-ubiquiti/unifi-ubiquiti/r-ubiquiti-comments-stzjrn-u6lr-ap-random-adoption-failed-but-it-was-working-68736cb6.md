---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-stzjrn-u6lr-ap-random-adoption-failed-but-it-was-working-68736cb6
title: "r-ubiquiti-comments-stzjrn-u6lr-ap-random-adoption-failed-but-it-was-working-68736cb6"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-stzjrn-u6lr-ap-random-adoption-failed-but-it-was-working-68736cb6.md
source_anchor: ""
source_lines: [1, 78]
sha256: 7a8323d59a5a858a690ca86c856c9350cbe4e99131da5b93de94be971a5c811c
---

# r-ubiquiti-comments-stzjrn-u6lr-ap-random-adoption-failed-but-it-was-working-68736cb6

U6-LR AP random "adoption failed", but it was working fine for 3 weeks prior 
        
        
        
    
    
    I need some input as to why...aside from 'maybe it's just broken'. Here is the situation at my office:
- 
      Ubiquiti setup has been in place since 2017 with regular hardware updates 
  - 
      Data Switch, phone switch, Cloud Key gen2+, AP, Watchguard firewall
- 
      
- 
      Due to recent "lag" I replaced UAP-AC-Pro with U6-LR (U6 pro was out of stock, until the day after I ordered!)
- 
      U6-LR was adopted on 1/16/22 with no problems, AC-pro removed from service (but not "unmanaged")
- 
      on 2/14/22 staff reported no Wifi. Console indicated "adoption failed" which is strange since I wasn't doing anything with the system. Eventually it came back online, then it happened again a short while later. I then put AC-Pro back in service on another POE port from switch, as a backup AP.
- 
      At some point, in the 'topography" the U6-LR was shown stemming from the AC-pro, but later was showing as connected to the switch and not the other AP.
- 
      All the while, the light on the LR was steady blue.
- 
      no firmware updates we in process (as far as I can tell)
Ubiquiti support ticket response seems to not understand the issue, because the person keeps giving me instruction on how to adopt an AP. Adopting the AP is not the problem. Below are the responses I got from them. thanks in advance.
****
Feb 16, 2022, 0:41 MST
Hello Alan,
The adopted AP gets into a disconnected state when the UniFi Network application does not have connectivity to the access point (check cables, network settings, and changes to topology).
I am sharing with you AP gets to Adoption Failed when I try to adopt a device.
The above article shared advanced troubleshooting methods and reasons for adoption failure.
l will go ahead and put this case in a resolved state, Thank you for choosing Ubiquiti, and have a nice day!
----
Feb 15, 2022, 11:53 MST
Hi ****,
Adoption via L2 is not the issue: the device was successfully adopted the day I received it several weeks ago. Today, for no apparent reason, it was “lost”. Is there any explanation to this situation? As far as I can tell, SSH with L3 adoption is not the issue, because I am sitting right underneath the device and the IT cabinet is 25 feet away from me. The AP should not have disconnected on its own: there was no firmware installing, and no other actions being taken on the site. Any thoughts on this?
----
Feb 15, 2022, 11:07 MST
Hello Alan,
As I can see in your shared network all your devices are adopted at this point in time.
If you see any disconnection in the network, you may ssh to the AP and troubleshoot it as shared in my last email.
----
Feb 15, 2022, 10:48 MST
Hi ****,
The device is currently showing as online and adopted (see image). However, 2x today staff reported no internet and on checking the interface it was showing as “adoption failed” even though there was no intervention pending on my part.
----
Feb 15, 2022, 10:38 MST
Hi,
Greetings for the day
I am ***** reaching out on behalf of the Ubiquiti Tech Support Team. I will be providing assistance on your case today.
If your AP is disconnected, Kindly manually hard reset your AP And readopt it as shared in UniFi Network - How to Adopt/Add new devices
Alternatively you can SSH to the AP and adopt it via UniFi - Layer 3 Adoption for Remote UniFi Network Applications
----
Alan
Feb 15, 2022, 9:42 MST
New AP "U6-LR has been disconnecting with indicator "adoption failed". Device has already been in use for a few weeks, as upgrade to AC AP Pro (now reconnected due to LR problems).
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Put your AP pro back, go to the setting then go to that AP turn off meshing.
Reset the AP LR (while not power on, hold the reset for 10 secs, after 10 secs plug the power in while still holding the rest for another 10 secs).
SSH to your AP LR and do the manual inform to your controller.
Hope that help
thanks. Will give it a try
So...the AP-LR got disconnected again yesterday and wouldn't re-adopt...I guess u/chris-itg was correct. I did the reset / SSH thing 3 weeks ago, but who knows. I restarted a support ticket with Unifi on this.
Meanwhile, I put my older AP back in service and have 2 running in the office. Now, my topology is weird. Meshing is ON, so maybe this is why. If I knew how to attach a photo I would, but basically the AP-LR is shown stemming from the other AP, rather than adjacent to it. Is this normal?
oh...NVM. I must just be dumb. I just realized that the reason I am seeing the weird topology is because I have meshing ON, which is unnecessary with the APs.
I would just factory reset it and adopt it again and monitor if it still happens.
Thanks. along the lines of what u/coolvt stated. appreciate your time...
Current firmware and controller version. There was a bug in a version that they resolved this very issue.
thx. I was thinking this might be the case. I might follow u/coolvt instructions anyway as well.
While that may work temporarily it won't permanently resolve. Again if you'll provide your controller version and AP firmware version I'll be happy to look at it for you and let you know if you're affected by the big as well as firmware revision to update.
I have exactly the same experience. One brand new U6-LR failing to adopt, and one UAP-AC-Pro that worked fine Friday and now it does not.
I can not get them to be adopted.

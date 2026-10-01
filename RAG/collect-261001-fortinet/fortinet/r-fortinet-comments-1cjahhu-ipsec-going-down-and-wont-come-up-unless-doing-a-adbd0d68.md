---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1cjahhu-ipsec-going-down-and-wont-come-up-unless-doing-a-adbd0d68
title: "r-fortinet-comments-1cjahhu-ipsec-going-down-and-wont-come-up-unless-doing-a-adbd0d68"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1cjahhu-ipsec-going-down-and-wont-come-up-unless-doing-a-adbd0d68.md
source_anchor: ""
source_lines: [1, 36]
sha256: a53ba6278b43af1240b85d97b36e58c7e86dc650cbb1268287e4c20a84d525ab
---

# r-fortinet-comments-1cjahhu-ipsec-going-down-and-wont-come-up-unless-doing-a-adbd0d68

IPSec going down, and won’t come up unless doing a reboot 
        
    Hi everyone, Having this really weird issue with a customer of mine. Decided to check with the community before opening a ticket.
Basically, every once in a while, one of the IPSec tunnels of the customer of mine is going down. Couldn’t seem to find a trigger for that. All I was able to figure out was that when ever running the following debug(whenever the issue happens again, i will try and provide with the log too):
diag vpn ike log-filter dst-addr4 x.x.x.x diag debug application ike -1
I was able to see that IKE requests being retransmitted constantly. Even tried to configure the tunnel from scratch, tried changing the IKE, made sure that DPD doesn’t cause the issues(usually DPD is on idle, tried several modes, tried disabling completely) Yet nothing helped. Only after rebooting the device(specifically a remote sites fw, and not customer’s main fw) - everything would go up.
What could be the issue? Thanks in advance!
Edit: The hardware is FGT 90E, the firmware is 7.2.4(i wanna upgrade it to 7.2.8, customer is stubborn).
Section des commentaires
Quick update - the adsl modem was limiting my encrypted traffic to 20mbps, and the line was 100 and traffic reached a bit over 50mbps(lived to know that modems can actually limit encrypted traffic, even though it does make sense). Got a different modem - everything started working perfectly. Thanks everyone for the help!
What is interesting is why rebooting the remote fortigate but not the upstream modem would resolve it. I also supsect in my case it is the upstream provider equipment. But cannot prove it.
Usually when it’s an adsl modem you can ask for a replacement. Any other case you can just destroy it and ask for new one :)
You should first check if the IKE packets are even being sent and arrive on both sides.
Seeing as how the issues seems to be the remote site the debugs and traffic there are more important.
As i have mentioned, ike requests being retransmitted. Ike packets being sent from both sides. Neither receive the packets back.
with what frequency does it go down? had issues with fortinet to cisco going down even though matched their p1 timeout. increased the timeout and problem went away.
There’s no consistency of the tunnel going down frequency. Which timeout did you increase? Key lifetime? Or what exactly
If there is no consistency - i'd disregard - i was referencing the key lifetime in our case. But it was clearly on a 24 hour basis it was going down.
I've seen behavior similar to what you describe, on occasion. I'm not sure that it is the same issue.
In my cases, the best that I have been able to determine so far is that an SA is having some sort of issue (sync?) that is preventing the re-establishment of the tunnel. I have yet to figure out what triggers this problem.
What I have found is that it was not the reboot that fixes it but rather enough time for the SAs on both sides to timeout and die. Then the tunnel is able to be re-established.
To restart the tunnels, try disabling the tunnel for longer than 60 seconds, so that the SAs timeout and die. After the SA dies, re-enable the tunnel and it comes right up and all is well for weeks to months.
Sounds very similar to my issue. Will try doing that thanks!
Yep. Been saying issues like this as well. Flashing the vpn interface on both sides usually gets it up quick.
We might be having the same or similar issues. Since October we have been having issues with some of our spokes. From time to time connectivity will be lost. Looking at the tunnels they are still up but the spokes are not getting any BGP routes. To resolve this a reboot or flushing the tunnel fixes it. There have been a number of technicians from Fortinet looking at this and everyone has said to try this and a few weeks later the issue comes back.
I'm having problems with 2-3 tunnels out of hundreds, currently flushing with automation works fine
https://community.fortinet.com/t5/FortiGate/Technical-Tip-Using-automation-stitches-to-run-debugs-or-flush/ta-p/273950
this can also be triggered from another internal monitoring with a incoming webhook on the fortigate
(fortios 7.2.7)
Do you have blackhole routes set for ipsecs?
This is a known bug I have experienced it in multiple version. Current fix is to disable the npu(network process unit). You should be able to find the Fortinet document online it’s one line command under config system. Fortinet say they fixed it on 7.2.8 but I am not sure I can’t trust them no more.
Same issue. For me , only a reboot rebuilds the tunnels. This is happening on only one dial up ipsec branch spoke out of about 70. I have tried cycling the interface, restarting ike service, dpd changes, auto neg, and flushing the tunnels. Nothing works besides a reboot of the spoke.
Commentaire supprimé par le membre
Sorry, post updated.
The hardware is FGT 90E, the firmware is 7.2.4(i wanna upgrade it to 7.2.8, customer is stubborn).
Hope they don’t have ssl vpn enabled then

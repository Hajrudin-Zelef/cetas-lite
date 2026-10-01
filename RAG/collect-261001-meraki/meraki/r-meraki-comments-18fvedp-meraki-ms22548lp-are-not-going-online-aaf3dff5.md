---
id: collect-261001-meraki/meraki/r-meraki-comments-18fvedp-meraki-ms22548lp-are-not-going-online-aaf3dff5
title: "r-meraki-comments-18fvedp-meraki-ms22548lp-are-not-going-online-aaf3dff5"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-18fvedp-meraki-ms22548lp-are-not-going-online-aaf3dff5.md
source_anchor: ""
source_lines: [1, 38]
sha256: 0ffff18236a853037e651dcab57f67baf9e9848f85b2c616a261a824365e6b76
---

# r-meraki-comments-18fvedp-meraki-ms22548lp-are-not-going-online-aaf3dff5

Meraki MS225-48LP are not going online 
        
        
        
    
    
    Hello, so we moved 2 ms switches from HQ to a branch office. When connecting all MS switches with stacking, all went online except the two which we brought from HQ. We have tried every troubleshooting possible but does not help.
We tried to disconnect all stack cables connecting to these switches and connect the uplink sfp to this switch but that did not help and the swich is in solig orange. I have removed it from org and added it again and still nothing. When using factory reset button nothing happens, it stays orange. I am holding the factory reset very long and nothing happens. Also powered all the switches off and conneted 1 by 1.
Trying to login to the MS switch with username: serial and password: blank, does not help either, I don´t get logged in.
What is the chance that to MS switches break at the same time? I do not see any other troubleshooting with this, we have opened a case for RMA.
Just thought if someone had any tips or something I might have missed.
Section des commentaires
Can you confirm that your uplink is actually working? What is on the other side of the uplink ports? On that switch can you see the switches MACs on the interface / ARP table?
Yes, the uplink is working as we are using another switch in the same stack for that and all are coming online but not on these 2 switches.
Did you have DHCP Security enabled and local status page disabled in the old network? Then factory reset is needed for sure..
Anyways, just try factory resetting the device using paperclip pressing reset button for 10-15 seconds until flashing leds stop flashing. Normally switch is receptive for factory reset as soon as it starts rainbowing...
Connect it just to a trunk with your management vlan as native and see what it does. Even try to set a fixed DHCP address so you can try to ping it from a neighboring switch..
factory reset does not work, it stays orange all the time. I hold it more then 15 seconds for sure. dhcp is open as other switches are working and getting correct mgmt network.
Yeah, but if the devices has that config from previous network and it now is not reporting in dashboard it will not update that dhcp security configpart...
I have heard of an AP power it on for 30 seconds, power off, power on for 30 sec, power off, and so on till I powered it on for 6 times and I had to let it boot the 7th time... Don't know if that is working for switches...
For sure call in with support.. do the tests they want you to, and if all is nog going well, it'll be RMA-ed...
Check firmware between the different networks. If the network they were coming from had a higher major version MS 15 or 16 they won’t downgrade without support help.
I will check, although I did connect the switch to the working fiber connection and got no connection so not sure how that will help when the switch cannot reach meraki cloud.
Hello, I understand you already mentioned some troubleshooting steps, just providing some ideas:
In Dashboard:
Create a new network
Remove the switch(es) from whichever network they are now, and move them to such a new network
Physically:
Disconnect any stack cable
Just connect 1 uplink (not the sfp for now, just an ethernet port) to the network: connect it to a port upstream which is known to be working and through which you can confirm you can receive an IP and access internet (for example, connect a laptop, make sure it can get an IP, and can browse to the internet). Give a sanity check to whichever DHCP server you have on site to confirm there are available IPs in the pool. I would suggest avoid to connect to the rest of the stack for this test.
Power on the switch
Factory reset the switch (press reset for ~15 sec) (yes I know you said you tried, I would suggest try again)
Take a packet capture on whichever port you have connected the switch to upstream. Do you see any traffic from the switch? If so, what traffic do you see? Anything coming from the switch at all? If yes, any response being sent back?
Be patient here, if you see traffic which would indicate bi-directional communication between switch and cloud - then give it 15-20 minutes, sometimes it takes a while to show green.
Did it come online? If yes, then now try to bring it online via the sfp. Doesn't work? maybe you got a faulty module there.
Still not online? Try again with the local status page (for the credentials, make sure you enter the serial number with upper-case letters with dashes)
Ultimately, if you just want to pick one step - take the packet capture upstream to see if there is any traffic from the switch or not.
Check your arp table upstream and look for any entries containing the switch MAC addresses

---
id: collect-261001-general-networking/general-networking/t5-wireless-ultra-fast-roaming-m-p-203445-abdae216
title: "t5-wireless-ultra-fast-roaming-m-p-203445-abdae216"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent", "pricing", "throughput"]
source: docs/RAG/collect-261001-general-networking/t5-wireless-ultra-fast-roaming-m-p-203445-abdae216.md
source_anchor: ""
source_lines: [1, 152]
sha256: 2c90c068dc68cf3829f384761202d3522b861492572004e26ac11472b548cc8d
---

# t5-wireless-ultra-fast-roaming-m-p-203445-abdae216

Ultra fast Roaming
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-01-2023 08:57 PM
Just posting to get some general advice from anyone experienced in tweaking Meraki wifi for warehouse enviroments.
What do you normaly set things to in relation to what I'll call ultra fast roaming. That is imagine some guy roaming around on one of those motorised picker. They are not standing still at any spot for more than a minute (more like 30 sedonds). This cause havoc with roaming.
Looking at the logs. You will see they are connected to an AP for like 30 seconds. Next they will be some where else on a totaly different AP.
Complain is "we can't scan (the barcode)" as soon as they stop in a location. But they can once their device has fixed onto the AP they are located in.
In our case they are using the Cipperlab hand scanner (basicly an android device).
Unfortunately the Meraki MR36 gear currently in place was done in a hurry (a long story / covid stock issues / what we had on hand).
I'm pushing to gett a profesional (with proper gear) to do a physical audit / scan / heat map etc.
Just after what I can tweak on Meraki Dashboard with what we got on hand.
Thanks
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-01-2023 10:54 PM
Hi
Ran into an issue while testing a setup with MR44 and Ascom VoFi Phones.
Wireless-Setup was done in this basic way:
DHCP: local
802.11r: off
802.11w: off
Auth: WPA3
Issue: Bad connection (sporadically disconnects), roaming is working bad.
-> Enabeld 802.11r -> same problem
-> switched to WPA2 -> solved the problem
Just as an example, my experience in the last 20years was:
It's nearly always an end-device problem. Got so much trouble with "ultra fast raoming" in combination with Aironet and CCX - most of the time the clients did not support the protocol ;).
So for your problem:
1. Check basic coverage, simply done with the Unify WifiMan App (Android) and check the roaming behaviour
2. Log date/time-stamps from affected clients when the disconnect happens and search Dashboard log.
3. Try to find out if it is an end-device problem of the bar-code scanner or a global problem (test with an windows notebook and walk through the area)
4. Check the APs load, if there are too many client connected/bandwith used.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-01-2023 11:23 PM
First thing I would do in a warehouse environment is reduce my min basic rate down to 6 Mbps - there are so many reflections in a warehouse that some 'cheaper' scanners can't handle a higher data rate. 
I'd also use a WPA2 PSK (as long as you have good physical security and an encrypted connection from the device to the app server). This way the device will roam easier/quicker anyway. 
I'd enable 802.11k but disable 80211r and v
lastly... I'd sell them 7Signal MobileEye which is an agent they can run on the android device to give them all the roaming/wireless data that the device can see from the OS perspective. That'll tell you ext what's going on and what's going wrong. It'll also correlate the driver versions in use and which ones work well and which ones don't. 7Signal
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-01-2023 11:31 PM
Thanks for all the advice / sugestions so far.
Yes, we are using WAP2 PSK.
Interesting on the Min rate. My thinking that you would go the other way (set a higher minimum rate) as to force clients to jump faster soon as the rate drop (they are far from the last AP).
I'll check out 7Signal MobileEye. First impresion on it is the website has no pricing and just a bunch of fluffy marketing video which doesn't tell you anything about the actual product.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-01-2023 11:39 PM
It’s the ONLY thing out there that can do what it does. They have all the patents. It’s very very good. Doesn’t work on iOS though. Apple don’t give access to the data they need. 
With regard to the MBR, it all depends on the roaming algorithm that the devices uses when roaming. I’ve seen some cheaper scanners revert to starting at the lowest end of the MCS table as they don’t need to send large amounts of data rather than working there way backwards from highest throughput to lowest. In that case, 6 Mbps is the best thing to enable as they’ll head there first.
MCS Index - this is how Wi-Fi works
Most laptops/smartphones will try to keep the highest data rate. Most VoIP phones, scanners, Sensors will just dive straight to 1SS - 6 Mbps as they send small amounts of data.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-02-2023 12:52 PM
Also if the interference isn't too bad in your area try 2.4GHz as that travels further and means fewer roams.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-02-2023 12:53 PM
Only if you can guarantee that your APs can't hear another on the same channel
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-01-2023 11:44 PM
Those barcode scanner scan like 20 characters barcode so bandwidth requirement is practicaly nothing!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-01-2023 11:44 PM
Exactly
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-03-2023 11:45 AM
If you have long aisles make sure you have single AP's that cover almost the entire aisle with an external antenna. But then also have AP's at the intersections so your device can leave the aisle while under the cover of a local AP. And if you have control over it set the roaming aggressiveness on the scanners high. To avoid the AP's in the intersections to have too far coverage and interfering each other you could use narrow beam antennas pointing straight down so your device is also enticed to roam faster to the neighboring AP.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-03-2023 03:53 PM
Thanks. That make sense.

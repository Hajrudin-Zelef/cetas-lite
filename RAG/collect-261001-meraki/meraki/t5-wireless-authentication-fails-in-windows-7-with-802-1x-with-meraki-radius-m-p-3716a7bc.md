---
id: collect-261001-meraki/meraki/t5-wireless-authentication-fails-in-windows-7-with-802-1x-with-meraki-radius-m-p-3716a7bc
title: "t5-wireless-authentication-fails-in-windows-7-with-802-1x-with-meraki-radius-m-p-3716a7bc"
domain: meraki
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-meraki/t5-wireless-authentication-fails-in-windows-7-with-802-1x-with-meraki-radius-m-p-3716a7bc.md
source_anchor: ""
source_lines: [1, 169]
sha256: e13fdfc76d38fe53273bd2d383f821c54806d7802fb5bc99608e687b3ba0729b
---

# t5-wireless-authentication-fails-in-windows-7-with-802-1x-with-meraki-radius-m-p-3716a7bc

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-09-2018 07:55 AM
Dear Colleagues,
I have an issue with the meraki radius log (with user & password) on laptops with windows 7 professional x64 bits
The problem happens with one`s try to connect to the network, at the moment appears a message of security alert of windows.
Once i proceed with "connect" appears a new message saying "the device cannot conect to the network", please see the images for more knowledge.
Really i dont know what to do
Regards!
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-13-2018 01:14 PM
On one of your Windows 7 machines, try manually updating the root certificates (I would prefer you used Windows Update instead ...).
http://dreamlayers.blogspot.co.nz/2009/12/windows-7-cant-always-automatically.html
Basically:
1. Download http://download.windowsupdate.com/msdownload/update/v3/static/trustedr/en/rootsupd.exe
2. Extract the files using the command rootsupd.exe /c /t:C:\temp\extroot
3. from c:\temp\extroot run the following 4 commands (from an elevated prompt)
updroots.exe authroots.sst
updroots.exe updroots.sst
updroots.exe -l roots.sst
updroots.exe -d delroots.sst
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-09-2018 08:03 AM
Alas I can only understand English.
But it looks like you have a certificate issue. Have you got all the Windows updates applied? You might be missing a root certificate update.
Are you using an internal RADIUS server?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-13-2018 01:04 PM
PhillipDAth,
Nice to meet you. Yes, looks like a certificate problem. On this enterprise we use Windows 7 Professional x64 Service Pack 1 But i dont know if is updated to date. For other way we dont have an internal radius server we only have the service router in cloud.
Regards.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-09-2018 08:30 AM
Try creating a manual connection to the network and deselect the option to verify the certificates under "Settings" on the security tab of the connection.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-13-2018 01:05 PM
MRCUR,
Thanks for the response. I have only one week whit this system. If you could give me more details to do that, i think could do it.
Greetings.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-13-2018 01:11 PM
This provides good instructions for Windows 7: https://documentation.meraki.com/MR/Encryption_and_Authentication/Configuring_Clients_for_802.1X_and_Meraki-hosted_RADIUS
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-19-2018 04:31 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-13-2018 01:14 PM
On one of your Windows 7 machines, try manually updating the root certificates (I would prefer you used Windows Update instead ...).
http://dreamlayers.blogspot.co.nz/2009/12/windows-7-cant-always-automatically.html
Basically:
1. Download http://download.windowsupdate.com/msdownload/update/v3/static/trustedr/en/rootsupd.exe
2. Extract the files using the command rootsupd.exe /c /t:C:\temp\extroot
3. from c:\temp\extroot run the following 4 commands (from an elevated prompt)
updroots.exe authroots.sst
updroots.exe updroots.sst
updroots.exe -l roots.sst
updroots.exe -d delroots.sst
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-19-2018 04:38 AM
Thank you,
Actually at this day. i have the last updates of windows seven on this machine.
but this link didn`t work
http://download.windowsupdate.com/msdownload/update/v3/static/trustedr/en/rootsupd.exe
i am thinking if is not a problem about my internet provider or is a meraki configuration problem
somethimes with this SSID y have microcuts. for example of 180 packets send i have 10 or 15 losts.
If you want y can take screenshots of this SSID configuration.
Regards!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-13-2018 06:29 AM
Did you ever get this working CarOneAdmin??
We are having same issues, but with Win 10 and Win 7 laptops.
Setup new SSID using 802.1x and meraki authentication. Mobile devices connect fine, but Win10 and Win7 laptops either say connecting...... for ages or come up with Certificate related issues same as those you posted.
Win7 machine fully updated as is Win 10 machine.
We want to be able to setup two Wifi SSIDS, internal and guest. Internal we'd like only those devices registered in Meraki MDM to be able to access, don't want to have to type userid or passwords etc, Just basically, if registered in MDM, then devices is allowed to connect to the Wifi SSID, not registered, no connection.
As said, this does work for Mobile (IPhones tested currently) but not on laptops.
Any help appreciated
cheers
Gary
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-09-2019 05:24 AM
I know this thread is quite old, but did anyone find a solution for Windows 7 Pro computers, Our Windows 10 Pros can connect no problem. I've installed the windows updates, checked the date and time, and installed the latest root certs. But still struggling to get this working.
Can anyone help ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-09-2019 07:16 AM
For anyone who finds this, I solved my problem by upgrading the WiFi drivers to the latest available. Lenovo's drivers were dated 2012 I upgraded using drivers downloaded directly from Intel's website. With the latest version installed the drivers were updated to 2017 and the problem is no longer happening.

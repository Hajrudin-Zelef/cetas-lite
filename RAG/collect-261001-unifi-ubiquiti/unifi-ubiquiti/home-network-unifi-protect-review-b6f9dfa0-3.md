---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/home-network-unifi-protect-review-b6f9dfa0-3
title: "home-network-unifi-protect-review-b6f9dfa0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-unifi-ubiquiti/home-network-unifi-protect-review-b6f9dfa0.md
source_anchor: ""
source_lines: [109, 137]
sha256: 5027855696efd0663d5380615c6d3358a034a43b4c9e96adfb71e91486385129
---

# home-network-unifi-protect-review-b6f9dfa0

I understand they now support 3rd party cameras albeit without smart detections (now “AI” detections as everything has to be labelled AI these days ). And then they sell a separate device to enable detections on 3rd party cameras – worth it?
I also can’t find anywhere if UniFi Protect iOS app supports PIP, I only see things about PIP notifications on Apple TV using HomeKit or Scrypted.
Sorry, took 6 months, but it’s updated finally.
I’m looking for a solution which uploads the video to a cloud storage in case of defined motion alerts within near real time. Is there any way to achive this with Unifi Protect? Anything which can be automated by coding would be ok as well.
Unifi Protect can record on movement, that is not enough? You could try to place the Unifi NVR offsite, with a site-to-site VPN connection for example. But offloading the recording with a script is not possible. You can’t even extract the video files from the record and play them on something else.
I wrote a tool for exactly this: https://github.com/ep1cman/unifi-protect-backup
Is there any way to allow a remote Ubiquiti camera on a different LAN to be part of my Unifi Protect environment on my LAN/Cloud Key Plus?
You will need to create a Site-to-Site VPN tunnel between the two sites.
Nice article. Can you do one as well for Unifi Access and how it can intergrate with other applications like a garage door. I think this may interest a lot of people out there
Is Ubiquiti going to add the option to use Windows 10 as a NVR to integrate with UniFi Protect?
You had the option for UniFi video……why take away the option? I would like to use a normal blade server for my NVR and still have remote access with Protect since UniFi Video has reached EOL.
Thanks
No probably not. No docker images either.
Hey Ruud,
I switched from UniFi Video on a Intel NUC to UniFi Protect on a CloudKey Gen2 Plus. The 2 things I miss the most, or maybe I’m looking wrong, is the option to make scheduled motion events for recording and the possibility to record and/or move the recordings to external storage like a NAS or cloud solution. Do you perhaps know if there are in Protect as well? Cause I cannot find them :S
Both are not possible, unfortunately. You can however replace the disk in your Cloudkey Gen2 plus with a larger one up to 5TB.
Hi Rudy,
The problem with local storage is when the CK is gone/stolen then so are the files with the perps on it. I hope Ubiquiti wises up. Those 2 features we’re in Unifi Video so why it’s not in Protect is a mystery to me….
What you can do is use rclone and make a backup of the folder /srv/unifi-protect/video. The file format is .ubv, so it might be a challenge to view the files. Maybe it’s possible to restore them back to another Cloudkey, but I haven’t tested that (yet).
Thanks for the useful review. Just to confirm my understanding, I’m currently running UniFi Video with the NVR software running on my Win10 PC. It looks like UniFi Protect does not allow you to use your Win10 PC as the NVR, instead it requires one of their hardware devices, is that correct?
Hi Tim,
That is correct. You will need a Cloudkey Gen2 Plus, Dream Machine Pro, or Univi NVR.
I have a basic unifi system in my home (USG, cloud key gen 2+, us-8-60, ap pro, ap lite) and am considering switching from Arlo cameras to unifi. What is the best option for me as far as additional equipment needed. I’m going to use the g3 flex cameras. I will probably have a max of 6 cameras. Should I just buy another us-8-60? I have a shop that is about a 200ft run that is wired with cat 6 where I am using the ap lite. What would I need to add to also be able to install a camera or two in my shop?
Great article. Thanks!
I would place an us-8-60 indeed in the shop, for the AC and flex (I have exactly the same setup for my garage). For the other camera’s you will need to see if it’s worth buying an extra us-8-60 or just use PoE adapters.
Great article! I am having an issue with the location/geofencing in that my status is always ‘Away’. I have gone into the UnifiProtect web interface and reset my location address and radius, restarted my app, but no change. Any other ideas of how to get it to know that I am home?
Did you set the correct permissions for the Protect app on your phone?
I’m looking to add this to my unifi network and was curious if you can have role based user logins. For example, if I want to give a friend a logon I just want them to be able to look at video feeds.
I have updated the article a bit with the roles. But in short, yes you can create (custom) role with view only and/or edit permissions per camera.

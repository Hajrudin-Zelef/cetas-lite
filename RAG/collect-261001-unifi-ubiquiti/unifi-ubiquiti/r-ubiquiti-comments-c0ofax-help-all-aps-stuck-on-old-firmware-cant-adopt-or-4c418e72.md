---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-c0ofax-help-all-aps-stuck-on-old-firmware-cant-adopt-or-4c418e72
title: "r-ubiquiti-comments-c0ofax-help-all-aps-stuck-on-old-firmware-cant-adopt-or-4c418e72"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-c0ofax-help-all-aps-stuck-on-old-firmware-cant-adopt-or-4c418e72.md
source_anchor: ""
source_lines: [1, 22]
sha256: 04dd822e98a65c46be265a9ba94e55f658c83a403fc3301721c36d3c134b0d9f
---

# r-ubiquiti-comments-c0ofax-help-all-aps-stuck-on-old-firmware-cant-adopt-or-4c418e72

Help. All APs stuck on old firmware - can't adopt or update.
SOLVED - many thanks everone. Got most of them updated via SSH, once I found which firmware was required and got used to how long it took for the APs to update and reboot. One of them got locked out of SSH for some reason so I manually 'put' the bin file onto it via TFTP.
I logged in to set up a temp guest wifi and noticed that my 4 UniFi AP-LRs and one AP were stuck on firmware 3.7.5.4969 and were not actually being managed by the controller. My wifi has been working all this time. I guess the controller updated and lost control. I am not able to adopt or update the APs via the web interface. I was not able to get it working using the SSH instructions so I ended up trying the tftp method featured here: https://help.ubnt.com/hc/en-us/articles/204910124-UniFi-TFTP-Recovery-for-Bricked-Access-Points#3
It connected and seemed to copy but the AP did not reboot. I'm at a complete loss and currently don't have any wifi. I'm on a Mac although I have access to Windows PCs if needed. Mine are the legacy APs with the amber/green lights.
Section des commentaires
What specifically does the controller say? You may be able to use the web UI to do a manual upgrade.
Separately, if you can ssh in to the AP's it's super easy. Assuming they are the original UAP-LR model, if you have something different, these will need changed.
upgrade http://dl.ubnt.com/unifi/firmware/BZ2/3.9.54.9373/BZ.ar7240.v3.9.54.9373.180913.2356.bin
That's not the latest version, but your so far behind I'd step through the upgrades, so then run:
upgrade https://dl.ubnt.com/unifi/firmware/BZ2/4.0.42.10433/BZ.ar7240.v4.0.42.10433.190518.0923.bin
The guy from Ubiquiti also said I would need to upgrade in steps. I’m going to try this in a moment. Thanks.
Solved now - see original post. Many thanks.
It doesn't sound like your APs are bricked though. Also in my experience, it's nicer using Macs you already have a shell built in.
Something like this happened to me once. I couldn't update nor adopt. What I did was I factory reset the APs and then used scp to load the firmware into the /tmp folder. It's been awhile, but I think I used these instructions:
https://help.ubnt.com/hc/en-us/articles/204910064-UniFi-Changing-the-Firmware-of-a-UniFi-Device#SSH
Thanks, I used that page and wasn’t able to get that working - which is why I went to the other method I mentioned above. However, I did not try factory resetting the APs first. I’ll give that a go next, I guess, with one of them at least.
Solved now - see original post. Many thanks.
Are you using the legacy controller or the newest version of the controller?
I think the controller has been updated to the latest version. Unifi controller v 5.10.24
Can't ssh in and run wget?
Commentaire supprimé par un membre de l’équipe de modération
Thank you! This fixed my problem!!!

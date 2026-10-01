---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-wifi-problems-with-iphones-and-ipads-after-switching-to-unifi-7623c924-04b99145
title: "questions-wifi-problems-with-iphones-and-ipads-after-switching-to-unifi-7623c924-04b99145"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-wifi-problems-with-iphones-and-ipads-after-switching-to-unifi-7623c924-04b99145.md
source_anchor: ""
source_lines: [1, 48]
sha256: 890c9f903366bcbf5e659bfc987265eb11f8fa5e2ebb99467c9b414a8e92f8fd
---

# questions-wifi-problems-with-iphones-and-ipads-after-switching-to-unifi-7623c924-04b99145

@UI-Team
Hi,
I recently installed a Unifi WiFi “upgrade” at home and have no end of troubles since.
I have since seen hundreds of threads on here regarding problems with apply products and I’m a little shocked about it TBH.
My network is:
Virgin Media Vivid 350 Broadband.
Super Hub.
US-8-150 Poe Switch with a cloud key connected.
Connected to this are some wired devices (NAS, CCTV, Hue Hub) , 1x AC-AP-Lite and 2 more switches.
Switch No.2 - Unifi Flex to iMac and Printer.
Switch No.3 - Unmanaged 8 Channel TP-Link PoE switch to 2x AC-AP-Lite + PS4 + TV
Old network was similar other brand switches but AP’s were all apple Airports and Time capsule.
Since switching over to Unifi the WiFi experience has been horrible.
Wired products are all good with no problems.
WiFi is constantly disconnecting, iPhones and iPads are showing a pop up saying WiFi network is not to the internet, but when choosing to carry on trying we get a connection. WiFi Calling is cutting out and dropping 100% of calls usually between 2-10 mins.
Social media streams are not scrolling and loading. Often fine and then suddenly stopping.
I could go on but you get the picture.
I have followed some threads on these forums and Reddit but the small changes I have been confident making have made no difference.
Can anyone help with this or make suggestions on what to try.
These are some of the log items coming up despite the dashboard showing green with good WiFi stats. 50-100 Anomalies a day showing up, all similar to this.
As a note, Super Hub has no option to change DNS settings.
Any suggestions??
Thanks in advance.
I'm going to plagiarize quote @RobbieH from this thread since it was typed up nicely
Don't allow automatic channel changes. You need to do a site survey and manually tune the radios to the proper channels and power levels. 
Disable Settings > Site > Auto-Optimize Network
Disable Settings > Wireless Networks > SSID > Advanced Options > High Performance Devices
Disable Connect High Performance clients to 5 GHz only
Disable Prefer 5G
Disable Fast Roaming
Set DTIM to 3 on both bands under WiFi > Advanced
And make sure the iPhones are on iOS 14.1.
Don't be shocked, these are Prosumer devices and are intended for installation by professionals who know how to perform site surveys, proper tuning, etc.
Also make sure your APs and switches are not on 4.3.21. If so, you need to move them to 4.3.20.
If you have more than one AP, paste a screenshot of Settings > Maintenance > Show System Config. To paste it here, use the square icon to the right of the smiley face icon.
On the iPhones & iPad you could try resetting all the network settings and also turn off “Private Address” in the WiFi settings. I have 4 iPads and 3 iPhones are working ok on 4.3.21 with AP AC Pros, USG Pro 4, Cloud Key 2+ (on 6.0.28) and PoE switches.
I second @RobbieH in that it is best to use 4.3.20 firmware. Many people have reported issues with 4.3.21. Once you have a stable configuration running, you can try 4.3.21 if you desire, but 4.3.20 is generally considered more stable (and better than chasing your tail here).
-Thanks for the replies. Is this the config page you asked for?
Is there a guide anywhere on how to dial back firmware? Thanks for the tips.
I understand that these are ‘prosumer’ equipment. I’m an electrician and occasionally install wifi upgrades. I picked these up with a view to learning the system and potentially offering it as a service in future. It’s far more buggy than expected. Being prompted in the app for an update to an unstable firmware release is an example of this.
Should I change firmware or change settings? Or both?
'
You will probably want to ssh into the APs to downgrade the firmware. Here is the process. You can download the firmware file here.
Your APs also are set to auto and they have overlapping channels... you will likely want to change them to manually define the channels so that they don't overlap, and then lower the power so that your devices want to connect to the nearest one as they move around your space rather than hanging on to a connection to an AP that is far away.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.

---
id: collect-261001-cisco/cisco/questions-780861-how-to-make-a-dep-registered-ipad-supervised-f4d3475c
title: "questions-780861-how-to-make-a-dep-registered-ipad-supervised-f4d3475c"
domain: cisco
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-780861-how-to-make-a-dep-registered-ipad-supervised-f4d3475c.md
source_anchor: ""
source_lines: [1, 17]
sha256: c88285e7d2be0246a57308931346ecee5adc0331c982724cffcd4e0e7e82309e
---

# questions-780861-how-to-make-a-dep-registered-ipad-supervised-f4d3475c

- Enter device serial number into DEP at deploy.apple.com
(You only have to do this once in the device's life).
- In Meraki, navigate to "SystemsManager>MDM>DEP". Your device should have appeared automatically.
- In Meraki>MDM>DEP, select the check-box next to your desired device and then click "Assign Settings" at the top of the device list (Settings are configured via "SystemsManager>MDM>Settings").
- From the device itself, Erase All Content and Settings (Alternatively, you can use Configurator to restore a backup at this point. This process also wipes the device).
- From the device, proceed through the initial configuration screens. One of the screens will ask if you would like to apply settings from the MDM ("SoandSo can automatically configure your iPad: Apply Configuration/Skip Configuration"). Select ApplyConfiguration and then click Next. Complete your setup.
- Everything should be happy now. Meraki should list the device as Supervised, and the device Settings>General>About should have an message stating the device is supervised just underneath the device name.
If you ever want to reconfigure the device, just push the Settings again from Meraki, wipe the device and then accept the Settings.
Short Version:
- Enter serial into DEP
- Apply settings from Meraki
- Wipe device
- Complete device config and accept settings
It has taken me quite a bit of time to get comfortable with iPad management through Meraki, but I wont go into all of the perils. I will say that things actually become more complex with Configurator (more parts, more problems). Fortunately, what you are trying to do is fairly straight forward, and it shouldn't take very long to get the hang of it.
You can also reference this page, if you haven't already: 
https://documentation.meraki.com/SM/Device_Enrollment/Using_Apple's_Device_Enrollment_Program_(DEP)_with_Systems_Manager
One more tip I can give you is that I have Meraki set to automatically install the Meraki MDM app OTA onto all newly enrolled iOS devices. Not sure if it helps at all, but it makes me feel like I'm being thorough. Good luck, fellow traveler.

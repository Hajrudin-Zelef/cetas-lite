---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-692127-unifi-ap-setup-and-configuration-903ae3fa
title: "questions-692127-unifi-ap-setup-and-configuration-903ae3fa"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-692127-unifi-ap-setup-and-configuration-903ae3fa.md
source_anchor: ""
source_lines: [1, 5]
sha256: 019b403bb84690a287b224271d58df8872ed139e2eb44cda0cf5733d9ccf9a9f
---

# questions-692127-unifi-ap-setup-and-configuration-903ae3fa

So I am trying to replace the wireless Access points in an Office. The APs they already use have static IP adresses (do not use DHCP) and are dlink dwl-2100 AP. I am supposed to replace with them the Ubiquiti UniFi AP. Sooo...since I have not done anything of the sort up til now, I need help. The incoming LAN cable is from a switch and is functioning. I have connected the new AP to that cable and have run the controller software...this is what I do not understand. It reqiures a secure SSID? It says that that will be the primary network between the APs (I need 4 altogether) but I am not sure if I need such a thing...I thought Access points were basically antennas that take a wired connection and 'transform' it into a wireless network laptops and mobile Devices can attach to. Any help with this would be great.
1 Answer 1
Don't lose hope, we use Ubiquiti UniFI APs at the office where I work and they are great!
The controller software is what you need to configure, all the APs talk to the UniFi controller and grab all their updates, configurations, IP addresses, etc, from the software. You will need to define at least one SSID. This is simply the wireless name for your WiFi. for example we have two SSIDs a "-Private" and "-Public". Both APs are secured but the "-Public" SSID has a very generic password, just to prevent any stranger from stealing WiFi.
The SSIDs are configured through the controller software, which should have a Wizard the first time you run it to set up all the APs. I also found the documentation on the website to have some useful information. UniFi Documentation

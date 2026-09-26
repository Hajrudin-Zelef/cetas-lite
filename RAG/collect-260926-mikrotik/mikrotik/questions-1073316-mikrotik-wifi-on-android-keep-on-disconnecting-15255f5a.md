---
id: collect-260926-mikrotik/mikrotik/questions-1073316-mikrotik-wifi-on-android-keep-on-disconnecting-15255f5a
title: "questions-1073316-mikrotik-wifi-on-android-keep-on-disconnecting-15255f5a"
domain: mikrotik
role: reference
task: reference
actors: ["Xiaomi"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/questions-1073316-mikrotik-wifi-on-android-keep-on-disconnecting-15255f5a.md
source_anchor: ""
source_lines: [1, 37]
sha256: c10d22aa9b51eed5fe909389f2dd010b4344991a53568821e49f1e45cc987df1
---

# questions-1073316-mikrotik-wifi-on-android-keep-on-disconnecting-15255f5a

My company recently purchased a few Mikrotik hAP Mini. These units are used for Site to Site VPN. The first unit I setup manually worked flawlessly. I then exported the configuration to the second, third and fourth units.
On these 3 units, when my Pixel 3a is connected, it receive an IP address, proclaming "no internet connection" and after one second it disconnect and reconnect to the WiFi over and over. The phone is literally on top of the AP unit and the surrounding area has at most 15 other APs. Another user who received a unit also reported the same problem on his Xiaomi phone.
On Mikrotik, the log shows "macaddr@wlan1: disconnected, received deauth: sending station leaving (3)". I believe this is just normal "user disconnected" message.
Surprisingly, if I set the DHCP server to not providing any DNS servers then the Pixel stays connected, asking what it should do when this wifi has no internet. The DNS server combination I tried are:
- Providing remote DNS server IP to the client (this server is over VPN)
- Providing 1.1.1.1 to the client
- Providing router's IP as DNS server, and
  - Set the router's upstream to 1.1.1.1, 1.0.0.1
  - Set the router's upstream to remote DNS server
  - Set the router's upstream to remote DNS server AND 1.1.1.1 as secondary
It seems like this mysterious feature of Android at blame here. Also, if I use the exact same IP configuration as DHCP (with remote DNS server), but as static IP, then I stay connected and internet is also working
In summary:
- Two Android phones from different brands can't stay connected to 3 Mikrotik devices over 5 seconds, each AP has the exact same setup and tested at point blank range.
- Providing no DNS server in DHCP does make it stay connected, but of course, no internet
- Providing any DNS server in DHCP give the same result as 1
- Using the exact same IP information provided by DHCP provide the expected result, but is a bad user experience
- Factory default config also work
- My Linux laptop works for all configuration in 1-5 except 2
The relevant Mikrotik configuration are:
/interface wireless
set [ find default-name=wlan1 ] band=2ghz-b/g/n channel-width=20/40mhz-XX \
    country=thailand disabled=no distance=indoors frequency=auto \
    installation=indoor mode=ap-bridge ssid=ssid station-roaming=enabled \
    wireless-protocol=802.11
/interface wireless security-profiles
set [ find default=yes ] authentication-types=wpa-psk,wpa2-psk mode=\
    dynamic-keys supplicant-identity=MikroTik
/ip dns
set allow-remote-requests=yes servers=10.0.0.1,1.1.1.1,1.0.0.1 use-doh-server=\
    https://cloudflare-dns.com/dns-query
/ip pool
add name=dhcp ranges=10.1.0.0/29
/ip address
set [ find comment=defconf ] address=10.1.0.1/29 interface=bridge network=10.1.0.0
/ip dhcp-server network
set [ find comment=defconf ] address=10.1.0.0/29 dns-server=10.1.0.1 gateway=10.1.0.1 netmask=29
Other settings are mostly factory default, including all default firewall configuration and DHCP optionset. The RouterOS version is 6.47.10 (longterm)

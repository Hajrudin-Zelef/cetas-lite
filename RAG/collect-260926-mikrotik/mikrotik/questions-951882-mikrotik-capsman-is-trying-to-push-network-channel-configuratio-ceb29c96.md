---
id: collect-260926-mikrotik/mikrotik/questions-951882-mikrotik-capsman-is-trying-to-push-network-channel-configuratio-ceb29c96
title: "questions-951882-mikrotik-capsman-is-trying-to-push-network-channel-configuratio-ceb29c96"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/questions-951882-mikrotik-capsman-is-trying-to-push-network-channel-configuratio-ceb29c96.md
source_anchor: ""
source_lines: [1, 35]
sha256: a75db11c9612bc84ebf51b9363d4b3b0c3511796a10d86c43bd6b9ca8b903c56
---

# questions-951882-mikrotik-capsman-is-trying-to-push-network-channel-configuratio-ceb29c96

According to information from allowed-channels command on my Access Point the allowed channels are the following:
5180/20-Ceee/ac,
5260/20-Ceee/ac/DP,
5745/20-Ceee/ac,
5750/20-Ceee/ac,
5755/20-Ceee/ac,
5760/20-Ceee/ac,
5765/20-Ceee/ac
Despite the above, the manager insists on assigning frequencies outside of this range. It's not obvious at first but sometimes in logs I can find something like this:
cap7: failed to select channel, no supported channel
cap7: failed to select channel, no supported channel
cap8: failed to select channel, no supported channel
cap1: selected channel 5300/20-eeCe/ac/DP(23dBm)
5300/20-eeCe/ac/DP is being selected like 90% of time. Very rarely it selects something else... which is also unsupported anyway.
It looks like no matter what I do with the configuration on the Caps manager, the manager is trying to provision configuration that I didn’t even define.
You can see that in my configuration I explicitly defined only the frequencies allowed by the access point. I have also explicitly set hw-supported-modes to make sure that 5G config will go only to 5G interface. The configuration itself is extremely simple:
/caps-man channel
add band=5ghz-onlyac control-channel-width=20mhz extension-channel=Ceee frequency=5180,5260,5745,5750,5755,5760,5765 name=5G-20-Ceee
add band=2ghz-g/n control-channel-width=20mhz extension-channel=Ce frequency=2412,2417,2422,2427,2432,2437,2442 name=2G-20-Ce
/caps-man datapath
add bridge=bridge1 client-to-client-forwarding=yes local-forwarding=yes name=datapath1
/caps-man security
add authentication-types=wpa-psk,wpa2-psk name="The LAN" passphrase=*************
/caps-man configuration
add channel=5G-20-Ceee comment="5G Only" datapath=datapath1 mode=ap name="The LAN11" security="The LAN" ssid="The LAN11"
add channel=2G-20-Ce comment="2G only" datapath=datapath1 mode=ap name="The LAN2" security="The LAN" ssid="The LAN2"
/caps-man provisioning
add action=create-dynamic-enabled comment="5G only" hw-supported-modes=ac master-configuration="The LAN11" name-format=identity
add action=create-dynamic-enabled comment="2G only" hw-supported-modes=gn master-configuration="The LAN2" name-format=identity
Despite the above configuration, the manager completely ignores channels configuration and selects something different
Which in the end looks like the reason why the access point is unable to setup wireless interface correctly. And this is a screenshot from the access point itself. You can see that the interfaces configuration is bogus and has nothing to do with the configuration that I’ve set on the manager:
You can see that interface is disabled nad there is no SSID - this is the symptom of provisioning channel out of allowed range.
I'm using 6.43.8 version of RouterOS (which is the latest one as of writing) with the following devices: hEX PoE Lite as router and capstan manager and multiple cAP ac access points.
I tried various different setups, disabled everything apart from basics to isolate the issue. MikroTik support responded with some general copy-paste answer and asked me to set hw-supported-modes which I did despite the fact that initially I was testing with only one interface.
Am I doing something wrong or is it a bug in RouterOS?

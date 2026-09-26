---
id: collect-260926-mikrotik/mikrotik/questions-927459-mikrotik-5ghz-ac-only-110mbps-1fe0de22
title: "questions-927459-mikrotik-5ghz-ac-only-110mbps-1fe0de22"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-927459-mikrotik-5ghz-ac-only-110mbps-1fe0de22.md
source_anchor: ""
source_lines: [1, 31]
sha256: 931dfe6dc29923e401147589cca23828e1c65910453b0f071cb33dd05dd22378
---

# questions-927459-mikrotik-5ghz-ac-only-110mbps-1fe0de22

I just bought Mikrotik hAP AC^2 and I configured it's WiFi to be controlled by CAPsMAN. The problem I have is that when I connect to it with my laptop (MacBook Pro 2016 Touchbar) I cannot achieve more than 110Mbps. On cable I can achieve 560Mbps.
Provision:
0   radio-mac=00:00:00:00:00:00 hw-supported-modes=gn identity-regexp="" common-name-regexp="" ip-address-ranges="" action=create-dynamic-enabled master-configuration=cfg1 slave-configurations="" name-format=prefix-identity 
     name-prefix="2GHz" 
1   radio-mac=00:00:00:00:00:00 hw-supported-modes=ac identity-regexp="" common-name-regexp="" ip-address-ranges="" action=create-dynamic-enabled master-configuration=cfg2 slave-configurations="" name-format=prefix-identity 
     name-prefix="5GHz" 
Configuration:
 0 name="cfg1" ssid="SSID" security=security1 datapath=datapath1 channel=channel1 
 1 name="cfg2" ssid="SSID" security=security1 datapath=datapath1 channel=channel2
Channel:
0 name="channel1" control-channel-width=20mhz band=2ghz-onlyn 
1 name="channel2" control-channel-width=20mhz band=5ghz-onlyac 
Datapath:
 0 name="datapath1" bridge=bridge 
Security:
 0 name="security1" authentication-types=wpa2-psk encryption=aes-ccm passphrase="*****"
In registration table I see 866Mbps and it's using both antennas, but when I test I get only 120Mbps. The device is 1M away from router, ALMOST clear. (there is a wired PC between them)
This is signal strength from device:
 PHY Mode:  802.11ac
  Channel:  36
  Country Code: BG
  Network Type: Infrastructure
  Security: WPA2 Personal
  Signal / Noise:   -50 dBm / -93 dBm
  Transmit Rate:    527
  MCS Index:    6
After detaching 5GHz interface from CAPsMAN the speed raised to 400Mbps. After attaching again I saw I hit the CPU limit.
cpu-used: 26%
cpu-used-per-cpu: 0%,1%,4%,100%
This one is when I'm detached from CAPsMAN.
So I'm hitting CPU limit - but why?

---
id: collect-260926-mikrotik/mikrotik/cap-ac-capsman-transmission-speed-not-as-high-as-it-can-be-1
title: "nov/16/2022 11:00:52 by RouterOS 7.6"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "throughput"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/cap-ac-capsman-transmission-speed-not-as-high-as-it-can-be.md
source_anchor: ""
source_lines: [1, 140]
sha256: 7ae6eec33583ded4b6d731373be729de82d1f96e55ef81707e3281154cbdbe7e
---

# nov/16/2022 11:00:52 by RouterOS 7.6

Hey.

i have this question, is it possible to do something to bump up the transmission speed to the maximum possible (866Mbps)?

My Macbook Pro practically always sets the connection at 390Mbps, sometimes higher values fall in. From what I’ve noticed, there’s no difference if I take the measurement at my desk (the cAP hangs behind a wall about 1m away) or under the AP itself. As for interference, it’s clean here, I live in the countryside, so it’s not a major problem. I paste below the configuration of my CAPsMAN.

I attached a screen from registration table + MacOS WiFi config




```
# nov/16/2022 11:00:52 by RouterOS 7.6
# software id = 23YS-AQUL
#
# model = RB760iGS
# serial number = D4500DDD9071
/caps-man channel
add band=2ghz-g/n control-channel-width=20mhz extension-channel=XX frequency=\
    2412,2437,2462 name=2.4Ghz-channel
add band=5ghz-onlyac control-channel-width=20mhz extension-channel=Ceee \
    frequency=5180,5260 name=5Ghz-channel skip-dfs-channels=no
/caps-man datapath
add bridge=bridge-vlan10-mgmt name=vlan10-datapath
add bridge=bridge-vlan20-lan name=vlan20-datapath vlan-id=20 vlan-mode=\
    use-tag
add bridge=bridge-vlan30-guest name=vlan30-datapath vlan-id=30 vlan-mode=\
    use-tag
add bridge=bridge-vlan40-devices name=vlan40-datapath vlan-id=40 vlan-mode=\
    use-tag
add bridge=bridge-vlan50-iot name=vlan50-datapath vlan-id=50 vlan-mode=\
    use-tag
/caps-man rates
add basic=11Mbps,6Mbps,9Mbps,12Mbps,18Mbps,24Mbps,36Mbps,48Mbps,54Mbps \
    ht-basic-mcs="" name=local-rates supported=\
    11Mbps,6Mbps,9Mbps,12Mbps,18Mbps,24Mbps,36Mbps,48Mbps,54Mbps
/caps-man security
add authentication-types=wpa2-psk eap-radius-accounting=no encryption=aes-ccm \
    name=temp-lan-cfg
add authentication-types=wpa2-psk encryption=aes-ccm name=guest-networrk-cfg
add authentication-types=wpa2-psk encryption=aes-ccm name=devices-network-cfg
add authentication-types=wpa2-psk encryption=aes-ccm name=iot-network-cfg
/caps-man configuration
add channel=2.4Ghz-channel country=poland datapath=vlan20-datapath \
    datapath.local-forwarding=no distance=dynamic hide-ssid=no installation=\
    any keepalive-frames=enabled mode=ap multicast-helper=full name=\
    2.4Ghz-main_network rates=local-rates security=temp-lan-cfg ssid=Z93-LAN
add channel=5Ghz-channel country=poland datapath=vlan20-datapath \
    datapath.local-forwarding=no distance=dynamic hide-ssid=no installation=\
    any keepalive-frames=enabled mode=ap multicast-helper=full name=\
    5Ghz-main_network rates=local-rates security=temp-lan-cfg ssid=Z93-LAN5
add channel=2.4Ghz-channel country=poland datapath=vlan30-datapath \
    installation=any name=2.4Ghz-guest_network rates=local-rates security=\
    guest-networrk-cfg ssid=Z93-GUEST
add channel=2.4Ghz-channel country=poland datapath=vlan40-datapath \
    installation=any name=2.4Ghz-devices_network rates=local-rates security=\
    devices-network-cfg ssid=Z93-DEVICES
add channel=2.4Ghz-channel country=poland datapath=vlan50-datapath hide-ssid=\
    no installation=any name=2.4Ghz-iot_network rates=local-rates security=\
    iot-network-cfg ssid=Z93-IOT
add channel=5Ghz-channel country=poland datapath=vlan40-datapath distance=\
    dynamic installation=any keepalive-frames=enabled name=\
    5GHz-devices_network rates=local-rates security=devices-network-cfg ssid=\
    Z93-DEVICES5
/caps-man access-list
add action=accept allow-signal-out-of-range=10s disabled=no signal-range=\
    -80..120 ssid-regexp=""
add action=accept allow-signal-out-of-range=10s comment="ESP - Korytarz" \
    disabled=no mac-address=8C:AA:B5:7B:F7:7B ssid-regexp=""
add action=accept allow-signal-out-of-range=10s comment="Watomierz ZMAI" \
    disabled=no mac-address=98:F4:AB:C2:1F:9E ssid-regexp=""
add action=accept allow-signal-out-of-range=10s comment=Pralka disabled=no \
    mac-address=A4:CF:12:77:88:90 ssid-regexp=""
add action=reject allow-signal-out-of-range=10s disabled=no signal-range=\
    -120..-81 ssid-regexp=""
/caps-man manager
set enabled=yes
/caps-man provisioning
add action=create-dynamic-enabled hw-supported-modes=gn master-configuration=\
    2.4Ghz-main_network name-format=prefix-identity name-prefix=2.4Ghz- \
    slave-configurations="2.4Ghz-main_network,2.4Ghz-guest_network,2.4Ghz-devi\
    ces_network,2.4Ghz-iot_network"
add action=create-dynamic-enabled hw-supported-modes=ac master-configuration=\
    5Ghz-main_network name-format=prefix-identity name-prefix=5Ghz- \
    slave-configurations=5GHz-devices_network
```

 
            
           
          
            
            
              The radio interface rate will go UP only if needed (and will obviously be upwards limited by radio channel conditions).

The problem with your configuration is use of CAPsMAN forwarding (datapath.local-forwarding=no). It adds two bottlenecks (as compared to local-forwarding=yes): one is traffic encapsulation into tunnel between CAP and CAPsMAN which causes traffic volume overhead (full-size ethernet frames have to be fragmented, also smaller frames get additional header overhead) and the other is CPU-bound encryption/decryption of the said tunnel. Both are known to reduce wireless throughput when using CAPsMAN forwarding, the magnitude depends on both cAP and CAPsMAN performance (and hEX S is not exactly speed monster).

The reduced maximum data rate (due to explained woes) mean that radio interface doesn’t have to go above observed rates.

Alternative to CAPsMAN forwarding is local forwarding, which in turn requires proper configuration of bridge on cAP device and the trunk connection between cAP and main router. This config has to be done manually on cAP itself, CAPsMAN doesn’t touch that part of CAP devices. The benefit is that it avoids all the overhead described above and usually allow for maximum possible wireless speeds.

             
            
           
          
            
            
              I always advocate for local forwarding… it makes it that much easier to change your wireless to another vendor, once reality sets in.

But as for connection speed… that’s a negotiated thing between client and AP. And remember you are using WIFi5 V1. No MU-MIMO or any of the other enhancements of the 2016 update to the wifi spec.

             
            
           
          
            
            
              
But as for connection speed… that’s a negotiated thing between client and AP


Indeed defined between client and AP, both directions are independent.

This MCS05=520Mbps  and MCS07=650Mbps used here, with 80MHz dual stream, normally means (apart from the fact that there might be no need to rise the interface rate, or that the interface rate is young and still rising) , that the finer encoding with 256-QAM did not succeed in this connection, and the speed fell back to 64-QAM.

See MCSINDEX.COM


The mechanism is described in the MT wiki (Wifi troubleshooting). The not-acked-transmission will be retried “hw-retries” times before it is counted as a failure, what will initiate a MCS rate step-down.

A successful transmission will initiate a MCS rate step-up. (Step-up can be MCS increase or going from 1S to 2S, or from 0.8µs to SGI 0.4µs.) Doing so, will converge the interface rate around it’s usable max rate.

Wall’s, reflections and interference distort transmissions, making high QAM transmissions too blurred to be decodeable.

Experience has shown that the MCS rate step-up in MT devices is slow. With a new connection, you can easily see the slow increments in the “registration table”. (Try to understand what you see, As from time to time the basic rate is show there!) Other brands are faster in step-up, and may even allow you to select the rate adjustment algoritm. (eg Draytek).

Interface rate adjustment should not depend on CAPsMAN or not, unless the CAPsMAN data link (CAPWAN tunnel) is only slowly feeding the AP. [But the CCQ would be 100% in that case, as all transmissions succeed]




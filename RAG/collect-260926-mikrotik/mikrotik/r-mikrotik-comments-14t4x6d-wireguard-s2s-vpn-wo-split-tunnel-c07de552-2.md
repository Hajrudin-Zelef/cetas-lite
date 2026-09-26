---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-14t4x6d-wireguard-s2s-vpn-wo-split-tunnel-c07de552-2
title: "r-mikrotik-comments-14t4x6d-wireguard-s2s-vpn-wo-split-tunnel-c07de552"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2023-07-07"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/r-mikrotik-comments-14t4x6d-wireguard-s2s-vpn-wo-split-tunnel-c07de552.md
source_anchor: ""
source_lines: [227, 427]
sha256: b0f29a48c411687a4bf09aa0fab00988e72c3c282b2b36f7af8afa6b81abd468
---

# r-mikrotik-comments-14t4x6d-wireguard-s2s-vpn-wo-split-tunnel-c07de552

      add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
    
      add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
    
      add action=drop chain=forward comment="defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
    
      /ip firewall mangle
    
      add action=change-mss chain=forward new-mss=1380 out-interface=WG_RA_VPN passthrough=yes protocol=tcp tcp-flags=syn tcp-mss=1381-65535
    
      /ip firewall nat
    
      add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN
    
      ...
    
      /ip ipsec identity
    
      add auth-method=eap certificate=NordVPN_CA disabled=yes eap-methods=eap-mschapv2 generate-policy=port-strict mode-config=NordVPN peer=NordVPN policy-template-group=NordVPN username=...
    
      add auth-method=eap certificate=NordVPN_CA disabled=yes eap-methods=eap-mschapv2 generate-policy=port-strict mode-config=NordVPN peer=NordVPN_Close policy-template-group=NordVPN username=...
    
      add auth-method=eap certificate=NordVPN_CA disabled=yes eap-methods=eap-mschapv2 generate-policy=port-strict mode-config=NordVPN peer=NordVPN_LowLoad policy-template-group=NordVPN username=...
    
      /ip ipsec policy
    
      add dst-address=0.0.0.0/0 group=NordVPN proposal=NordVPN_Light src-address=0.0.0.0/0 template=yes
    
      /ip route
    
      add disabled=no dst-address=192.168.21.0/24 gateway=192.168.14.2 routing-table=main suppress-hw-offload=no
    
      ...
    
      /system identity
    
      set name=RB4011.home
    
      /system resource irq rps
    
      set sfp-sfpplus1 disabled=no
    
      ...
    
      /tool graphing interface
    
      add interface=pppoe-out1
    
      /tool mac-server
    
      set allowed-interface-list=LAN
    
      /tool mac-server mac-winbox
    
      set allowed-interface-list=LAN
    
And WAP AC LTE6
      [admin@WAP_AC_LTE6] > export
    
      # 2023-07-07 14:39:19 by RouterOS 7.10.1
    
      # software id = V6I3-6BQW
    
      #
    
      # model = RBwAPGR-5HacD2HnD
    
      # serial number = ...
    
      /interface bridge
    
      add admin-mac=DC:2C:6E:D9:5F:00 auto-mac=no comment=defconf name=bridge
    
      /interface lte
    
      set [ find default-name=lte1 ] allow-roaming=no band=1,3,7,20 network-mode=lte
    
      /interface wireguard
    
      add listen-port=13232 mtu=1420 name=WG_S2S_VPN
    
      /interface list
    
      add comment=defconf name=WAN
    
      add comment=defconf name=LAN
    
      ...
    
      /ip pool
    
      add name=default-dhcp ranges=192.168.21.10-192.168.21.254
    
      /ip dhcp-server
    
      add address-pool=default-dhcp interface=bridge lease-time=10m name=defconf
    
      /interface bridge port
    
      add bridge=bridge comment=defconf interface=ether1
    
      add bridge=bridge comment=defconf interface=ether2
    
      add bridge=bridge interface=wlan_5Ghz
    
      add bridge=bridge interface=wlan_2,4GHz
    
      /ip neighbor discovery-settings
    
      set discover-interface-list=LAN
    
      /interface list member
    
      add comment=defconf interface=bridge list=LAN
    
      add comment=defconf interface=lte1 list=WAN
    
      /interface wireguard peers
    
      add allowed-address=0.0.0.0/0 endpoint-address=70.18.138.1 endpoint-port=13232 interface=WG_S2S_VPN public-key=...
    
      /ip address
    
      add address=192.168.21.1/24 comment=defconf interface=bridge network=192.168.21.0
    
      add address=192.168.14.2/24 interface=WG_S2S_VPN network=192.168.14.0
    
      /ip dhcp-client
    
      # DHCP client can not run on slave or passthrough interface!
    
      add interface=wlan_2,4GHz
    
      # DHCP client can not run on slave or passthrough interface!
    
      add interface=wlan_5Ghz
    
      /ip dhcp-server network
    
      add address=192.168.21.0/24 comment=defconf dns-server=192.168.21.1 gateway=192.168.21.1
    
      /ip dns
    
      set allow-remote-requests=yes
    
      /ip dns static
    
      add address=192.168.21.1 comment=defconf name=router.lan
    
      /ip firewall filter
    
      add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
    
      add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
    
      add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
    
      add action=accept chain=input comment="defconf: accept to local loopback (for CAPsMAN)" dst-address=127.0.0.1
    
      add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN
    
      add action=accept chain=forward comment="defconf: accept in ipsec policy" ipsec-policy=in,ipsec
    
      add action=accept chain=forward comment="defconf: accept out ipsec policy" ipsec-policy=out,ipsec
    
      add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related hw-offload=yes
    
      add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
    
      add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
    
      add action=drop chain=forward comment="defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
    
      add action=accept chain=input dst-port=13232 protocol=udp
    
      /ip firewall nat
    
      add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN
    
      /ip route
    
      add disabled=no distance=1 dst-address=192.168.10.0/24 gateway=192.168.14.1 pref-src="" routing-table=main scope=30 suppress-hw-offload=no target-scope=10
    
      add disabled=yes distance=1 dst-address=0.0.0.0/0 gateway=192.168.14.1 pref-src="" routing-table=main scope=30 suppress-hw-offload=no target-scope=10
    
      ...
    
      /system identity
    
      set name=WAP_AC_LTE6
    
      ...
    
      /tool mac-server
    
      set allowed-interface-list=LAN
    
      /tool mac-server mac-winbox
    
      set allowed-interface-list=LAN
    

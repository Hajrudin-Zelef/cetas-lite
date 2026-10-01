---
id: collect-261001-fortinet/fortinet/questions-409591-fortigate-remote-vpn-no-matching-gateway-for-new-request-e0401dca
title: "questions-409591-fortigate-remote-vpn-no-matching-gateway-for-new-request-e0401dca"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2012-07-20"]
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-409591-fortigate-remote-vpn-no-matching-gateway-for-new-request-e0401dca.md
source_anchor: ""
source_lines: [1, 63]
sha256: 6325952ebdfbf5007981f38851bd182a50a170d9cdc9f03b6a9e67cc7ab7bba2
---

# questions-409591-fortigate-remote-vpn-no-matching-gateway-for-new-request-e0401dca

I am trying to configure a Fortigate 60C to act as an IPSec endpoint for remote VPN.
I configured it like this :
SCR-F0-FGT100C-1 # diagnose vpn ike config
vd: root/0
name: SCR-REMOTEVPN
serial: 7
version: 1
type: dynamic
mode: aggressive
dpd: enable  retry-count 3  interval 5000ms
auth: psk
dhgrp:  2
xauth: server-auto
xauth-group: VPN-group
interface: wan1
distance: 1
priority: 0
phase2s:
  SCR-REMOTEVPN-PH2 proto 0 src 0.0.0.0/0.0.0.0:0 dst 0.0.0.0/0.0.0.0:0  dhgrp 5  replay  keep-alive  dhcp
policies: none
Here is the configuration:
config vpn ipsec phase1-interface
    edit "SCR-REMOTEVPN"
        set type dynamic
        set interface "wan1"
        set dhgrp 2
        set xauthtype auto
        set mode aggressive
        set proposal aes256-sha1 aes256-md5
        set authusrgrp "VPN-group"
        set psksecret ENC xxx
    next
config vpn ipsec phase2-interface
    edit "SCR-REMOTEVPN-PH2"
        set keepalive enable
        set phase1name "SCR-REMOTEVPN"
        set proposal aes256-sha1 aes256-md5
        set dhcp-ipsec enable
    next
end
But when I try to connect from a remote device (I tested with an Android Phone), the phone fail to connect and the fortinet return this error :
2012-07-20 13:08:51 log_id=0101037124 
type=event 
subtype=ipsec 
pri=error 
vd="root" 
msg="IPsec phase 1 error" 
action="negotiate" 
rem_ip=xxx
loc_ip=xxx 
rem_port=1049 
loc_port=500 
out_intf="wan1" 
cookies="xxx" 
user="N/A" 
group="N/A" 
xauth_user="N/A" 
xauth_group="N/A" 
vpn_tunnel="N/A" 
status=negotiate_error error_reason=no matching gateway for new request 
peer_notif=INITIAL-CONTACT
I tried searching on the web, but i did not find anything revelant to this.
Do you have any idea of what can be the problem ? I tried many combinaisons of settings on the fortigate without success..

---
id: collect-260926-mikrotik/mikrotik/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca-2
title: "mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2016-09-18", "2016-10-14", "2016-11-04", "2016-11-19", "2017-01-21", "2017-01-23", "2017-01-24", "2017-02-28"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca.md
source_anchor: ""
source_lines: [113, 235]
sha256: 851187d9ea1c5af56c97fe6680570d44dbd26fb230cc9621f0c350daef1acc1a
---

# mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca

Afterwards, you can use following command to get dns-name value of a local router, which will be used in configuration of script on remote router:

/ip cloud print

### IPSec Remote Address Update Script

This script is no longer needed, but it is kept here for a reference. In the past RouterOS could use only IPs in peer address configuration, so dynamically updated addresses needed to be updated – now hostnames are allowed. Additionally in the past IPSec policies required to have the sa-dst-address attribute updated with IP of remote peer as well – this is now updated automatically by RouterOS.

Following script updates IPSec peer address and policy SA destination address, if remote peer’s address has changed. Make sure to fill out correct peerid and peerhost variables:

- peerid: Identifies IPSec peer and policy records, which will get updated, by comment field value. In this example it is “vpn01”.
- peerhost: Remote router’s value of dns-name from IP Cloud setup. If different DDNS solution than MikroTik IP Cloud is used for remote site, enter remote DDNS hostname here. In this example it is “0123456789.sn.mynetname.net”. Be aware that peerhost is different for each router.

##### MikroTik router 1 and MikroTik router 2

```
/system script add name="ipsec-peer-update-vpn01" policy=read,write,test source=":local peerid    \"vpn01\"\
    \n:local peerhost  \"0123456789.sn.mynetname.net\"\
    \n:local peerip    [:resolve \$peerhost]\
    \n:local peeruid\
    \n:set peeruid     [/ip ipsec peer   find comment=\"\$peerid\" and address!=\"\$peerip/32\"]\
    \n:local policyuid\
    \n:set policyuid   [/ip ipsec policy find comment=\"\$peerid\" and sa-dst-address!=\"\$peerip\"]\
    \n:if (\$peeruid != \"\") do={\
    \n  /ip ipsec peer set \$peeruid address=\"\$peerip/32\"\
    \n  :log info \"Script ipsec-peer-update updated peer '\$peerid' with address '\$peerip'\"\
    \n}\
    \n:if (\$policyuid != \"\") do={\
    \n  /ip ipsec policy set \$policyuid sa-dst-address=\"\$peerip\"\
    \n  :log info \"Script ipsec-peer-update updated policy '\$peerid' with address '\$peerip'\"\
    \n}"
```
### Schedulers

This scheduler periodically executes enforced update of IP Cloud DDNS IP. Enforced IP Cloud update is used, because MikroTik router behind NAT does not always check public IP in required intervals. Therefore by scheduler and enforced update, this interval is controlled.

Configured intervals should reflect how promptly routers will detect and process public IP change, but also they should avoid any excessive usage. In case of often VPN connection break downs, because of public IP changes, it should be considered to use static public IPs instead and thus avoid IP changes altogether.

##### MikroTik router 1 and MikroTik router 2

/system scheduler
add disabled=yes interval=10m name=ip-cloud-forceupdate on-event="/ip cloud force-update" policy=read,write

As mentioned before, in the past, script was required to update IPSec addresses with remote DDNS IP. This is no longer the case, so creation of such scheduler is now skipped. Using this will rewrite the hostname set in peer’s configuration with IP address, so avoid setting this up if it is not necessary. The script execution scheduler looked like this:

/system scheduler
add disabled=yes interval=1m name=ipsec-peer-update-vpn01 on-event="/system script run ipsec-peer-update-vpn01" policy=read,write,test

### Netwatch and Route

Netwatch checks availability of remote MikroTik router’s LAN IP address. In case remote router is unavailable, Netwatch enables schedulers for update of IP Cloud DDNS IP and IPSec remote address (not needed anymore, so no longer present in this guide). When remote router is available again, schedulers are disabled back. Netwatch is not able to identify correct interface for remote router’s LAN IP, therefore route needs to be added as well. Make sure that gateway is set to your local router’s LAN interface or bridge.

##### MikroTik router 1

/ip route add comment="vpn01" distance=1 dst-address=10.10.20.0/24 gateway=bridge-local
/tool netwatch add comment=ipsec-peer-update-vpn01 down-script="/system scheduler enable ip-cloud-forceupdate" host=10.10.20.1 up-script="/system scheduler disable ip-cloud-forceupdate"

##### MikroTik router 2

/ip route add comment="vpn01" distance=1 dst-address=10.10.10.0/24 gateway=bridge-local
/tool netwatch add comment=ipsec-peer-update-vpn01 down-script="/system scheduler enable ip-cloud-forceupdate" host=10.10.10.1 up-script="/system scheduler disable ip-cloud-forceupdate"

## Result

Netwatch on both Mikrotik routers should detect unavailability of remote router and trigger down-script, enforcing IP Cloud update. If public IP did not change since IP Cloud has been set up. When the tunnel is up, Netwatch will detect presence of remote router via VPN tunnel and disable schedulers, until next public IP change will break the tunnel and trigger the Netwatch again. While tunnel is up, hosts in both private networks can communicate with each other.

## 77 Comments

## dabar · September 18, 2016 at 20:31

thank you for sharing!!

## george · October 14, 2016 at 07:14

I would appreciate a variation of this setup where only one of the routers uses dynamic IP while the other has known static IP (a “central office”).

## Pessoft · November 4, 2016 at 00:16

Setup mentioned in the article should work also in the case when one of the routers is connected using the static IP. In such scenario, for the router connected via static IP you don’t need IP Cloud (dynamic DNS) and ip-cloud-forceupdate scheduler and for the router connected via dynamic IP you don’t need ipsec-peer-update scheduler and temporary placeholder IP set in ipsec section ( 127.99.99.99/32 ) can be configured directly ( with the known static IP ).

## Joao · November 19, 2016 at 22:56

Hi,

i am stuck on phase1 with IPSEC error: phase1 negotiation failed due to time up……

both of routers are behind nat with ports (udp500, udp4500) opened and static ip addresses.

Pessoft, please contact me by email, i need to get this working.

Thanks

## chris · January 21, 2017 at 04:49

The best mikrotik site-site ipsec guide I have seen out there.

I have followed all the step and the connection is up but I can not ping the remote site. Can you help please?

## Pessoft · January 24, 2017 at 00:24

Thanks.

At first try to check IPSec. On both sides you should see in remote-peers

establishedconnection ( /ip ipsec remote-peers print ) and pair ofmatureinstalled SAs ( /ip ipsec installed-sa print ). Then try to ping remote Mikrotik’s internal IP and also IP of some device in remote network. If both, peers and SAs, are correct and ping still does not work via IPSec tunnel, but does locally, then it can be a routing issue. Also, I had issues with the IPSec NAT-T tunnel running on Mikrotik RouterOS 6.38 and had to upgrade to 6.38.1.
## kk · January 23, 2017 at 07:58

thanks, the script works well.

## Kevin · February 28, 2017 at 10:32

Hello，

Firstly, thanks for your doc, I have followed all the step and can get remote site public ip via IP Cloud each other but the ipsec phase I is still fail. can you help please ?

Secondary, I have one question for this solution. if I setup another Mikrotik box in one of the site and paste same config besides the ether1-gateway IP, the ipsec tunnel will be establish with which one?

## Pessoft · February 28, 2017 at 23:50

Hello,

Mikrotiks on both sites need to have IPSec traffic forwarded from their gateway routers ( in example above are these gateways called ISP routers ), so at first I suggest to check that there is 500/UDP and 4500/UDP forwarding configured on these gateways. Also check in which state are remote peers on Mikrotiks ( terminal command /ip ipsec remote-peers print ).


---
id: collect-261001-fortinet/fortinet/fortigate-web-filtering-all-you-need-to-know-6
title: "diag webfilter fortiguard cache dump"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2025-03-13"]
keywords: ["license", "sandbox"]
source: docs/RAG/collect-261001-fortinet/fortigate-web-filtering-all-you-need-to-know.md
source_anchor: ""
source_lines: [894, 1045]
sha256: d3cb55e77080f8ca2d2b2883fa648dd91d78e3906fff2e9179af514d4427f538
---

# diag webfilter fortiguard cache dump

```
(fortiguard) # get
fortiguard-anycast  : enable
fortiguard-anycast-source: fortinet
protocol            : https
port                : 443  <-- WORTH CHANGING IF IT IS BEING BLOCKED to 53, 8888
load-balance-servers: 1
auto-join-forticloud: disable
update-server-location: usa
sandbox-region      :
update-ffdb         : enable
update-uwdb         : enable
update-dldb         : enable
update-extdb        : enable
update-build-proxy  : enable
vdom                :  <-- WHICH VDOM IS USED TO CONNECT TO FORTIGUARD
auto-firmware-upgrade: enable
auto-firmware-upgrade-day:
auto-firmware-upgrade-delay: 3
auto-firmware-upgrade-start-hour: 1
auto-firmware-upgrade-end-hour: 4
FDS-license-expiring-days: 15
antispam-force-off  : disable
antispam-cache      : enable
antispam-cache-ttl  : 1800
antispam-cache-mpermille: 1
antispam-license    : Contract
antispam-expiration : Fri Jan  1 2038 <-- LIC EXPIRATION DATE
antispam-timeout    : 7
outbreak-prevention-force-off: disable
outbreak-prevention-cache: enable
outbreak-prevention-cache-ttl: 300
outbreak-prevention-cache-mpermille: 1
outbreak-prevention-license: Contract
outbreak-prevention-expiration: Fri Jan  1 2038
outbreak-prevention-timeout: 7
webfilter-force-off : disable
webfilter-cache     : enable
webfilter-cache-ttl : 3600
webfilter-license   : Contract
webfilter-expiration: Fri Jan  1 2038 <-- LICENSE EXPIRATION DATE
webfilter-timeout   : 15
anycast-sdns-server-ip: 0.0.0.0
anycast-sdns-server-port: 853
sdns-options        :
source-ip           : 0.0.0.0 <-- SOURCE IP TO USE
source-ip6          : ::
proxy-server-ip     :
proxy-server-port   : 0
proxy-username      :
proxy-password      : *
ddns-server-ip      : 0.0.0.0
ddns-server-ip6     : ::
ddns-server-port    : 443
interface-select-method: auto <-- SOURCE INT OF THE REQUESTS TO FORTIGUARD,
                                CAN BE: AUTO, SD-WAN, SPECIFY
```
If you want to disable anycast and switch to unicast when having problems with FortiGuard servers reachablity, see https://yurisk.info/2021/02/21/failed-to-connect-to-fortiguard-servers-updated/

- 
To see the latest update date and versions of all downloadable databases on Fortigate, run **diag autoupdate versions** :

# diag autoupdate versions
AV Engine
---------
Version: 7.00035 signed
Contract Expiry Date: Fri Jan  1 2038
Last Updated using manual update on Thu Nov 14 00:39:00 2024
Last Update Attempt: n/a
Result: Updates Installed
Virus Definitions
---------
Version: 1.00000 signed
Contract Expiry Date: Fri Jan  1 2038
Last Updated using manual update on Mon Apr  9 19:07:00 2018
Last Update Attempt: n/a
Result: Updates Installed
--CUT--

To force Fortigate to download updates: **execute update-now**.

To debug auto updates: **diagnose debug application update -1**, **dia deb enable**, and **exe update-now**. Excerpts of debug output:

# diagnose debug application update -1
# dia deb en
# execute update-now
upd_daemon[1838]-Received update request from pid=2307
do_update[680]-Starting now UPDATE (final try)
__update_upd_comp_by_settings[495]-Disabling NIDSDB/ISDB/MUDB components.
__update_upd_comp_by_settings[499]-Disabling APPDB/IOTDB/OTDB components.
__update_upd_comp_by_settings[507]-Disabling AVEN components.
__update_upd_comp_by_settings[511]-Disabling AVDB/FLDB/MMDB components.
upd_fds_load_default_server6[1046]-Resolve and add fds usupdate.fortinet.net
ipv6 address failed.
upd_comm_connect_fds[457]-Trying FDS 173.243.141.6:443
[116] __ssl_cert_ctx_load: Added cert /etc/cert/factory/root_Fortinet_Factory.cer,
root ca Fortinet_CA, idx 0 (default)
[497] ssl_ctx_use_builtin_store: Loaded Fortinet Trusted Certs
upd_cfg_extract_sfas_version[789]-version=07004000SFAS00000-00005.00046-2502130417
pack_obj[186]-Packing obj=Protocol=3.2|Command=Update|Firmware=FGVMA6-FW-7.04-2726|
SerialNumber=FGTAWSNXXXXXX|UpdateMethod=0|AcceptDelta=1|
DataItem=07004000DBDB00100-00003
upd_status_extract_contract_info[1330]-pending registration(255)
support acct() company() industry() <-- THIS FGT IS NOT REGISTERED FOR SUPPORT

- 
Verify that the name resolving on the Fortigate works. For live queries to Fortiguard, Fortigate uses 2 hosts - **service.fortiguard.net** (without anycast) and**globalguardservice.fortinet.net** (with anycast):

execute ping service.fortiguard.net
execute ping update.fortiguard.net
execute ping guard.fortinet.net

- 
For Fortiguard Category filtering, make sure the websites are categorized correctly by looking at the cache with **diag webfilter fortiguard cache dump** , see Category cache verification for example output.

If the cache is not up-to-date, use **dia test app urlfilter 2** to clear the cache (no downtime).

To see the whole list of Categories with their respective numbers, run **get webfilter categories**

```
FGT-Perimeter # get webfilter categories
  g01 Potentially Liable:
      1 Drug Abuse
      3 Hacking
      4 Illegal or Unethical
      5 Discrimination
--CUT--
```
- 
Logs - Web Filter logs are under Events → Security Events → Web Filter menu or on the CLI it is **exe log filter category 3** and**exe log display** .  See example at [logs_cli]
- 
Regular **dia deb flow** debug is useful as well. E.g. for an internal host 192.168.17.0:

diagnose debug reset
diagnose debug flow trace stop
diagnose debug flow filter clear
diagnose debug flow filter addr 192.168.17.0
diagnose debug flow filter port 443
diagnose debug flow show function-name enable
diagnose debug console timestamp enable
diagnose debug flow trace start 1000
dia deb enable

- 
Finally, the most verbose Web Filtering debug is **dia deb app urlfilter -1** ,
the example output:

2025-03-13 04:47:38 0(2190) Stop quota EP localvpn1, cat 52
2025-03-13 04:47:38 0(2190) action=10(ftgd-block) wf-act=3(BLOCK)
user="localvpn1" src=192.168.17.0 sport=56505 dst=4.207.247.137
dport=443 service="https" cat=52 url_cat=52 ip_cat=0
hostname="client.wns.windows.com" url="/"

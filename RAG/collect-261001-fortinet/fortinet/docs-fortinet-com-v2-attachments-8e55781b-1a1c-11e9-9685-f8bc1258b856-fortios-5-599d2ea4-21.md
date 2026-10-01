---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-21
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [2517, 2709]
sha256: 61b03d10c4d33d54f8723d8071e2ce94924edad195bac9313ed1b72f44234441
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

FGCPconfigurationexamplesandtroubleshooting SettinguptwonewFortiGates asanFGCPcluster
4. Enterandconfirmanewpassword.
5. Select OK.
6. Goto Network > Interfaces.
7. Editthe port2interfaceandchange IP/Netmaskto10.11.101.100/24.
8. Select OK.
AfterchangingtheIPaddressoftheport1interfaceyoumayhavetochangetheIP
addressofyourmanagement computerandthenreconnecttotheport1interface
usingthe172.20.120.141 IPaddress.
9. Editthe port1interfaceandchange IP/Netmaskto172.20.120.141/24.
10. Select OK.
11. Goto Network > Static Routes.
12. Changethedefault route.
Destination IP/Mask 0.0.0.0/0.0.0.0
Gateway 172.20.120.2
Device port1
Distance 10
13. Select OK.
Configuring a NAT mode active-passive cluster of two FortiGates - CLI
Usethefollowing procedurestoconfiguretwoFortiGates forNATHAoperationusingtheFortiGate CLI. These
proceduresassumeyouarestartingwithtwoFortiGates withfactorydefault settings.
To configure the first FortiGate (host name FGT_ha_1)
1. PowerontheFortiGate.
2. Connectanullmodem cabletothecommunicationsportofthemanagement computerandtotheFortiGate
Consoleport.
3. StartHyperTerminal(oranyterminal emulation program),enteranamefortheconnection,andselect OK.
4. ConfigureHyperTerminaltoconnectdirectlytothecommunicationsportonthecomputertowhichyouhave
connectedthenullmodem cableandselect OK.
5. Selectthefollowing portsettingsandselect OK.
Bits per second 9600
Data bits 8
Parity None
Stop bits 1
Flow control None
HighAvailability
Fortinet TechnologiesInc.
77

SettinguptwonewFortiGates asanFGCPcluster FGCPconfigurationexamplesandtroubleshooting
6. Press EntertoconnecttotheFortiGate CLI.
TheFortiGate CLIloginpromptappears.
If thepromptdoesnotappear,pressEnter.If itstill doesnotappear,poweroff yourFortiGate and
poweritbackon.If youareconnected,atthisstageyouwillseestartupmessagesthat willconfirm
youareconnected.Theloginpromptwillappearafterthestartuphascompleted.
7. Typeadmin andpress Entertwice.
8. RegisterandapplylicensestotheFortiGate.
9. ChangethehostnameforthisFortiGate.
config system global
set hostname FGT_ha_1
end
10. ConfigureHAsettings.
config system ha
set mode a-p
set group-name example1.com
set password HA_pass_1
end
TheFortiGate negotiatestoestablishanHAcluster.Youmaytemporarilylosenetworkconnectivity
withtheFortiGate astheHAclusternegotiatesandtheFGCPchangestheMACaddressofthe
FortiGate interfaces.TheMACaddressesoftheFortiGate interfaceschangetothefollowing virtual
MACaddresses:
l port1interfacevirtualMAC:00-09-0f-09-00-00
l port2interfacevirtualMAC:00-09-0f-09-00-01
l port3interfacevirtualMAC:00-09-0f-09-00-02
l port4interfacevirtualMAC:00-09-0f-09-00-03
Toreconnectsooner,youcanupdatetheARPtableofyourmanagement PCbydeletingtheARP
tableentryfortheFortiGate (orjustdeletingallarptableentries).Youmaybeabletodeletethearp
tableofyourmanagement PCfromacommandpromptusingacommandsimilartoarp -d.
ToconfirmtheseMACaddresschanges,youcanusetheget hardware nic (ordiagnose
hardware deviceinfo nic)CLIcommandtoviewthevirtualMACaddressofanyFortiGate
interface. Forexample,usethefollowing commandtoviewtheport1interfacevirtualMACaddress
(MAC)andtheport1permanentMACaddress(Permanent_HWaddr):
get hardware nic port1
.
.
.
Current_HAaddr 00:09:0f:09:00:00
Permanent_HWaddr  02:09:0f:78:18:c9
.
.
.
10. DisplaytheHAconfiguration(optional).
get system ha
group-id : 0
group-name : example1.com
mode : a-p
password : *
78 HighAvailability
Fortinet TechnologiesInc.

FGCPconfigurationexamplesandtroubleshooting SettinguptwonewFortiGates asanFGCPcluster
hbdev : "port3" 50 "port4" 50
session-sync-dev :
route-ttl : 10
route-wait : 0
route-hold : 10
sync-config : enable
encryption : disable
authentication : disable
hb-interval : 2
hb-lost-threshold : 20
hello-holddown : 20
arps : 5
arps-interval : 8
session-pickup : disable
update-all-session-timer: disable
session-sync-daemon-number: 1
link-failed-signal : disable
uninterruptible-upgrade: enable
ha-mgmt-status : disable
ha-eth-type : 8890
hc-eth-type : 8891
l2ep-eth-type : 8893
ha-uptime-diff-margin: 300
vcluster2 : disable
vcluster-id : 1
override : disable
priority : 128
slave-switch-standby: disable
minimum-worker-threshold: 1
monitor :
pingserver-monitor-interface:
pingserver-failover-threshold: 0
pingserver-slave-force-reset: enable
pingserver-flip-timeout: 60
vdom : "root"
11. Poweroff theFortiGate.
To configure the second FortiGate (host name FGT_ha_2)
1. PowerontheFortiGate.
2. Connectanullmodem cabletothecommunicationsportofthemanagement computerandtotheFortiGate
Consoleport.
3. StartHyperTerminal,enteranamefortheconnection,andselect OK.
4. ConfigureHyperTerminaltoconnectdirectlytothecommunicationsportonthecomputertowhichyouhave
connectedthenullmodem cableandselect OK.
5. Selectthefollowing portsettingsandselect OK.
Bits per second 9600
Data bits 8
Parity None
HighAvailability
Fortinet TechnologiesInc.
79

SettinguptwonewFortiGates asanFGCPcluster FGCPconfigurationexamplesandtroubleshooting
Stop bits 1
Flow control None
6. Press EntertoconnecttotheFortiGate CLI.
TheFortiGate CLIloginpromptappears.
7. Typeadmin andpress Entertwice.
8. RegisterandapplylicensestotheFortiGate.
9. Youcanalsoinstallanythird-partycertificatesontheprimaryFortiGate beforeforming thecluster.Oncethe
clusterisformedthird-partycertificatesaresynchronizedtothebackupFortiGate.
FortiTokenlicensescanbeaddedatanytime becausetheyaresynchronizedtoallclustermembers.
10. ChangethehostnameforthisFortiGate.
config system global
set hostname FGT_ha_2
end
11. ConfigureHAsettings.
config system ha
set mode a-p
set group-name example1.com
set password HA_pass_1
end
TheFortiGate negotiatestoestablishanHAcluster.Youmaytemporarilylosenetworkconnectivity
withtheFortiGate astheHAclusternegotiatesandbecausetheFGCPchangestheMACaddressof
theFortiGate interfaces.
Toreconnectsooner,youcanupdatetheARPtableofyourmanagement PCbydeletingtheARP
tableentryfortheFortiGate (orjustdeletingallarptableentries).Youmaybeabletodeletethearp
tableofyourmanagement PCfromacommandpromptusingacommandsimilartoarp -d.
12. DisplaytheHAconfiguration(optional).
get system ha
group-id : 0
group-name : example1.com
mode : a-p
password : *
hbdev : "port3" 50 "port4" 50
session-sync-dev :
route-ttl : 10
route-wait : 0
route-hold : 10
sync-config : enable
encryption : disable
authentication : disable
hb-interval : 2
hb-lost-threshold : 20
hello-holddown : 20
arps : 5
arps-interval : 8
session-pickup : disable
update-all-session-timer: disable
session-sync-daemon-number: 1
link-failed-signal : disable
80 HighAvailability
Fortinet TechnologiesInc.


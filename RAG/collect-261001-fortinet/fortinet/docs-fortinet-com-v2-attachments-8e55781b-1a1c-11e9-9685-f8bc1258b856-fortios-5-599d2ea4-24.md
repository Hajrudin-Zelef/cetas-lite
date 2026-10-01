---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-24
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [2995, 3176]
sha256: ce1afa7df347820b88ee990f235ffdffbf94a8c06824f3a2a889091714bab3b3
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

FGCPconfigurationexamplesandtroubleshooting Active-activeHAclusterintransparentmode
Operation Mode Transparent
Management IP/Mask 10.11.101.100/24
Default Gateway 10.11.101.2
6. Select Apply.
Theclusterswitchestooperatingintransparentmode. ThevirtualMACaddressesassignedtothe
clusterinterfacesdonotchange.
To view cluster status
Usethefollowing stepstoviewtheclusterdashboardandclustermemberslisttoconfirmthat theclusterunits
areoperatingasacluster.
Oncetheclusterisoperating,becauseconfigurationchangesaresynchronizedtoall
clusterunits, configuringtheclusteristhesameasconfiguringanindividualFortiGate.
Youcouldhaveperformedthefollowing configurationstepsseparatelyoneach
FortiGate beforeyouconnectedthem toformacluster.
1. StartInternet Explorerandbrowsetotheaddresshttps://10.11.101.100 (remembertoincludethe“s”inhttps://).
TheFortiGate Loginisdisplayed.
2. TypeadminintheNamefield andselectLogin.
TheFortiGate dashboardisdisplayed.
TheHAStatusdashboardwidgetdisplayshowlongtheclusterhasbeenoperating(Uptime)andthe
time sincethelastfailoveroccurred(StateChanged).YoucanhoverovertheStateChangedtime to
seetheeventthat causedthestatechange.YoucanalsoclickontheHAStatusdashboardwidgetto
configureHAsettingsortogetalistingofthemostrecentHAeventsrecordedbythecluster.
3. Goto System > HA toviewtheclustermemberslist.
Thelistshowsbothclusterunits, theirhostnames,theirrolesinthecluster,andtheirdevicepriorities.
Youcanusethislisttoconfirmthat theclusterisoperatingnormally.Forexample,ifthelistshows
onlyoneclusterunitthentheotherunithasleft theclusterforsomereason.
To troubleshoot the cluster configuration
If theclustermemberslistandthedashboarddonotdisplayinformation forbothclusterunits, theFortiGates are
notfunctioning asacluster.SeeTroubleshootingHAclustersonpage138totroubleshootthecluster.
To add basic configuration settings to the cluster
Usethefollowing stepstoconfigurethecluster.Notethat thefollowing areexampleconfigurationstepsonlyand
donotrepresentallofthestepsrequiredtoconfiguretheclusterforagivennetwork.
1. LogintotheclusterGUI.
2. Goto System > Administrators.
3. Edit adminandselect Change Password.
4. Enterandconfirmanewpassword.
5. Select OK.
HighAvailability
Fortinet TechnologiesInc.
89

Active-activeHAclusterintransparentmode FGCPconfigurationexamplesandtroubleshooting
Youaddedadefault gatewaywhenyouswitchedtotransparentmodesoyoudon’t
needtoaddadefault routeaspartofthebasicconfigurationoftheclusteratthis
point.
Configuring a transparent mode active-active cluster of two FortiGates - CLI
Usethefollowing procedurestoconfiguretheFortiGates fortransparentmodeHAoperationusingtheFortiGate
CLI.
To configure each FortiGate for HA operation
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
Stop bits 1
Flow control None
6. Press EntertoconnecttotheFortiGate CLI.
TheFortiGate CLIloginpromptappears.If thepromptdoesnotappear,pressEnter.If itstill doesnot
appear,poweroff yourFortiGate andpoweritbackon.If youareconnected,atthisstageyouwillsee
startupmessagesthat willconfirmyouareconnected.Theloginpromptwillappearafterthestartup
hascompleted.
7. Typeadmin andpress Entertwice.
8. RegisterandapplylicensestotheFortiGate.
9. ChangethehostnameforthisFortiGate. Forexample:
config system global
set hostname FGT_ha_1
end
10. ConfigureHAsettings.
config system ha
set mode a-a
set group-name example2.com
set password HA_pass_2
end
90 HighAvailability
Fortinet TechnologiesInc.

FGCPconfigurationexamplesandtroubleshooting Active-activeHAclusterintransparentmode
Thisistheminimum recommendedconfigurationforanactive-activeHAcluster.You
canalsoconfigureotherHAoptions,butifyouwaituntil aftertheclusterisoperating
youwillonlyhavetoconfiguretheseoptionsoncefortheclusterinsteadofseparately
foreachclusterunit.
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
group-name : example2.com
mode : a-a
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
HighAvailability
Fortinet TechnologiesInc.
91

Active-activeHAclusterintransparentmode FGCPconfigurationexamplesandtroubleshooting
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
Stop bits 1
Flow control None
6. Press EntertoconnecttotheFortiGate CLI.
TheFortiGate CLIloginpromptappears.If thepromptdoesnotappear,pressEnter.If itstill doesnot
appear,poweroff yourFortiGate andpoweritbackon.If youareconnected,atthisstageyouwillsee
startupmessagesthat willconfirmyouareconnected.Theloginpromptwillappearafterthestartup
hascompleted.
7. Typeadmin andpress Entertwice.
8. RegisterandapplylicensestotheFortiGate.
9. ChangethehostnameforthisFortiGate.
config system global
92 HighAvailability
Fortinet TechnologiesInc.


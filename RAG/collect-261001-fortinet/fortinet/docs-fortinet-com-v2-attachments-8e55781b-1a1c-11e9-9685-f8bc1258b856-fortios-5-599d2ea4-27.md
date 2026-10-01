---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-27
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "licenses"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [3489, 3628]
sha256: b9fbed1237a482bdf1c81df3c318ea5038434242bce620a805a7af904b745bd5
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

FGCPconfigurationexamplesandtroubleshooting FortiGate-5000active-activeHAclusterwithFortiClient licenses
unitsandchangetheirconfiguration, thoseconfigurationchangesarealsosynchronizedtoeachclusterunit. The
exceptiontothisisconfigurationobjectsthat arenotsynchronized,suchasthehostname, FortiClient license
andsoon.
YoucanalsomanageeachclusterunitbyloggingintotheprimaryunitCLIandusingthefollowing commandto
connecttootherclusterunits:
execute ha manage <cluster-index>
To add basic configuration settings to the cluster
Usethefollowing stepstoconfigurethecluster.
1. LogintotheclusterGUI.
Youcanlogintotheprimaryunitoranyoneoftheclusterunitsusingtheappropriatemgmt1 IP
address.
2. Goto System > Administrators.
3. Edit admin andselect Change Password.
4. Enterandconfirmanewpassword.
5. Select OK.
6. Goto Network > Interfacesandeditthe port1interface. SetthisinterfaceIPaddresstotheaddressrequiredto
connecttotheinterfacetotheInternet.
7. Edittheport2interfaceandsetitsIPtoanIPaddressfortheinternalnetwork.
To add a FortiClient license to each cluster unit
NormallyyouwouldaddFortiClient licensestotheFortiGates beforeforming thecluster.However,youcanuse
thefollowing stepstoaddFortiClient licensestoanoperatingcluster.
ContactyourresellertopurchaseFortiClient licensesforyourclusterunits. Eachclusterunitmusthaveitsown
FortiClient license.
Whenyoureceivethelicensekeysyoucanlogintohttps://support.fortinet.com andaddaFortiClient licensekey
toeachlicensedFortiGate. Then,aslongastheclustercanconnecttotheInternet thelicensekeysare
downloadedfromtheFortiGuardnetworktoalloftheFortiGates inthecluster.
Youcanalsousethefollowing stepstomanuallyaddthelicensekeystoyourclusterunitsfromtheGUI. Your
clustermustbeconnectedtotheInternet.
1. LogintotheGUIofeachclusterunitusingitsreservedmanagement interfaceIPaddress.
2. Gotothe License InformationdashboardwidgetandbesideFortiClient select Enter License.
3. Enterthelicensekeyandselect OK.
4. Confirmthat thelicensehasbeeninstalledandthecorrectnumberofFortiClients arelicensed.
5. Repeatforalloftheclusterunits.
Youcanalsousethefollowing commandtoaddthelicensekeyfromtheCLI:
execute FortiClient-NAC update-registration-license <license-number>
YoucanconnecttotheCLIsofeachclusterunitusingtheirreservedmanagement IPaddress.
YoucanalsologintotheprimaryunitCLIandusetheexecute ha manage commandtoconnecttoeach
clusterunitCLI.
HighAvailability
Fortinet TechnologiesInc.
101

FortiGate-5000active-activeHAclusterwithFortiClient licenses FGCPconfigurationexamplesandtroubleshooting
Configuring the FortiGate-5000 active-active cluster - CLI
TheseproceduresassumeyouarestartingwiththreeFortiGate-5001DboardsandtwoFortiSwitch-5003Bboards
installedinacompatible FortiGate-5000serieschassis.TheFortiSwitch-5003Bboardsareinchassisslots1and
2andtheFortiGate-5001Dboardsareinchassisslots3,4,and5andthechassisispoweredon.Alldevicesare
intheirfactorydefault configuration. NoconfigurationchangestotheFortiSwitch-5003Bboardsarerequired.
To configure the FortiGate-5005FA2 units
1. Fromtheinternalnetwork,logintotheCLIoftheFortiGate-5001Dunitinchassisslot3byconnectingtothe
mgmt1 interface.
Bydefault themgmt1 interfaceofeachFortiGate-5001DunithasthesameIP
address.TologintoeachFortiGate-5001Dunitseparatelyyoucouldeitherdisconnect
themgmt1 interfacesoftheunitsthat youdon’twanttologintoorchangethemgmt1
interfaceIPaddressesforeachunitbyconnectingtoeachunit’sCLIfromtheirconsole
port.
Youcanalsouseaconsoleconnection.
2. RegisterandapplylicensestotheFortiGate.
3. ChangethehostnameforthisFortiGate. Forexample:
config system global
set hostname 5001D-Slot-3
end
4. Enterthefollowing commandtodisplaybackplaneinterfacesontheGUI:
config system global
set show-backplane-intf enable
end
5. SettheAdministrativeStatusofthebase1andbase2interfacesto Up.
config system interface
edit base1
set status up
next
edit base2
set status up
end
6. AddanIPaddresstothemgmt1 interface.
config system interface
edit mgmt1
set ip 172.20.120.110/24
set allowaccess http https ssl ping
end
Becausemgmt1 willbecomethereservedmanagement interfacefortheclusteruniteachFortiGate-
5001Dshouldhaveadifferent mgmt1 interfaceIPaddress.Givethemgmt1 interfaceanaddressthat
isvalidfortheinternalnetwork.OnceHAwiththereservedManagement interfaceisenabledtheIP
addressofthemgmt1 interfacecanbeonthesamesubnetastheport2interface(whichwillalsobe
connectedtotheInternal network).
102 HighAvailability
Fortinet TechnologiesInc.

FGCPconfigurationexamplesandtroubleshooting FortiGate-5000active-activeHAclusterwithFortiClient licenses
7. ConfigureHAsettings.
config system ha
set mode a-a
set ha-mgmt-status enable
set ha-mgmt-interface mgmt1
set group-name example3.com
set password HA_pass_3
set hbdev base1 50 base2 50 mgmt2 50
end
TheFortiGate negotiatestoestablishanHAcluster.WhenyouselectOKyoumaytemporarilylose
connectivitywiththeFortiGate astheHAclusternegotiatesandtheFGCPchangestheMACaddress
oftheFortiGate interfaces.TheMACaddressesoftheFortiGate-5001D interfaceschangetothe
following virtualMACaddresses:
l base1interfacevirtualMAC:00-09-0f-09-00-00
l base2interfacevirtualMAC:00-09-0f-09-00-01
l fabric1interfacevirtualMAC:00-09-0f-09-00-02
l fabric2interfacevirtualMAC:00-09-0f-09-00-03
l fabric3interfacevirtualMAC:00-09-0f-09-00-04
l fabric4interfacevirtualMAC:00-09-0f-09-00-05
l fabric5interfacevirtualMAC:00-09-0f-09-00-06
l mgmt1 keepsitsoriginalMACaddress
l mgmt2 interfacevirtualMAC:00-09-0f-09-00-08
l port1interfacevirtualMAC:00-09-0f-09-00-09
l port2interfacevirtualMAC:00-09-0f-09-00-0a
Toreconnectsooner,youcanupdatetheARPtableofyourmanagement PCbydeletingtheARP
tableentryfortheFortiGate (orjustdeletingallarptableentries).Youmaybeabletodeletethearp
tableofyourmanagement PCfromacommandpromptusingacommandsimilartoarp -d.
Youcanusetheget hardware nic (ordiagnose hardware deviceinfo nic)CLI
commandtoviewthevirtualMACaddressofanyFortiGate interface. Forexample,usethefollowing
commandtoviewtheport1interfacevirtualMACaddress(Current_HWaddr)andtheport1
permanentMACaddress(Permanent_ HWaddr):
get hardware nic base1
.
.
.
Current_HWaddr 00:09:0f:09:00:00
Permanent_HWaddr 00:09:0f:71:0a:dc
.
.
.
7. RepeatthesestepsfortheFortiGate-5001Dunitsinchassisslots4and5,withthefollowing differences.
Setthemgmt1 interfaceIPaddressofeachFortiGate-5001Dunittoadifferent IPaddress.
SettheFortiGate-5001Dunitinchassisslot4hostnameto:
config system global
set hostname 5001D-Slot-4
end
HighAvailability
Fortinet TechnologiesInc.
103


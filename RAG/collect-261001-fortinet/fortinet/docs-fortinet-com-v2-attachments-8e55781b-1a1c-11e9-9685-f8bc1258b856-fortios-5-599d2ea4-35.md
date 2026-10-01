---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-35
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["licenses"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [4752, 4932]
sha256: 17d178dca012bf40a13213559afa3fbc7daaba140cc3b8eb6377b3f58836838a
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

ExampleHAandredundantinterfaces FGCPconfigurationexamplesandtroubleshooting
To connect the cluster to the network
1. Connecttheport1andport2interfacesofFGT_ha_1andFGT_ha_2toaswitchconnectedtotheInternet.
Configuretheswitchsothat theport1andport2ofFGT_ha_1makeuparedundantinterfaceand
port1andport2ofFGT_ha_2makeupanotherredundantinterface.
2. Connecttheport3andport4interfacesofFGT_ha_1andFGT_ha_2toaswitchconnectedtotheinternal
network.
Configuretheswitchsothat theport3andport4ofFGT_ha_1makeuparedundantinterfaceand
port3andport4ofFGT_ha_2makeupanotherredundantinterface.
3. Connecttheport5interfacesofFGT_ha_1andFGT_ha_2together. YoucanuseacrossoverEthernetcableor
regularEthernetcablesandaswitch.
4. Connecttheport5interfacesoftheclusterunitstogether. YoucanuseacrossoverEthernetcableorregular
Ethernetcablesandaswitch.
5. Powerontheclusterunits.
Theunitsstartandnegotiate tochoosetheprimaryunitandthesubordinateunit. Thisnegotiation
occurswithnouserinterventionandnormallytakeslessthanaminute.
Whennegotiation iscompletetheclusterisreadytobeconfiguredforyournetwork.
To view cluster status
Usethefollowing stepstoviewclusterstatusfromtheCLI.
1. LogintotheCLI.
2. Enterget system status toverifytheHAstatusoftheclusterunitthat youloggedinto.Lookforthefollowing
information inthecommandoutput.
Current HA mode: a-a,
master
Theclusterunitsareoperatingasaclusterandyouhaveconnectedtothe
primaryunit.
Current HA mode: a-a,
backup
Theclusterunitsareoperatingasaclusterandyouhaveconnectedtoa
subordinateunit.
Current HA mode:
standalone
TheclusterunitisnotoperatinginHAmode
3. Enterthefollowing commandtoconfirmtheHAconfigurationofthecluster:
get system ha status
HA Health Status: OK
Model: FortiGate- XXXX
Mode: HA A-P
Group: 0
Debug: 0
Cluster Uptime: 7 days 00:30:26
.
.
.
Youcanusethiscommandtoconfirmthat theclusterishealthyandoperatingnormally,some
information abouttheclusterconfiguration, andinformation abouthowlongtheclusterhasbeen
136 HighAvailability
Fortinet TechnologiesInc.

FGCPconfigurationexamplesandtroubleshooting ExampleHAandredundantinterfaces
operating.Information notshowninthisexampleincludeshowtheprimaryunitwasselected,
configurationsynchronizationstatus, usagestatsforeachclusterunit, heartbeatstatus, andthe
relativeprioritiesoftheclusterunits.
To troubleshoot the cluster configuration
SeeTroubleshootingHAclustersonpage138.
To add basic configuration settings and the redundant interfaces
Usethefollowing stepstoaddafewbasicconfigurationsettingsandtheredundantinterfaces.
1. Addapasswordfortheadminadministrative account.
config system admin
edit admin
set password <psswrd>
end
2. Temporarilydeletethedefault route.
Youcannotaddaninterfacetoaredundantinterfaceifanysettings(suchasthedefault route)are
configuredforit. Inthisexampletheindexofthedefault routeis1.
config router static
delete 1
end
3. Addtheredundantinterfaces:
config system interface
edit Port1_Port2
set type redundant
set member port1 port2
set ip 172.20.120.141/24
set vdom root
next
edit Port3_Port4
set type redundant
set member port3 port4
set ip 10.11.101.100/24
set vdom root
end
ThevirtualMACaddressesoftheFortiGate interfaceschangetothefollowing. Notethat port1and
port2bothhavetheport1virtualMACaddressandport3andport4bothhavetheport3virtualMAC
address:
l port1interfacevirtualMAC:00-09-0f-09-00-00
l port10interfacevirtualMAC:00-09-0f-09-00-01
l port11interfacevirtualMAC:00-09-0f-09-00-02
l port12interfacevirtualMAC:00-09-0f-09-00-03
l port13interfacevirtualMAC:00-09-0f-09-00-04
l port14interfacevirtualMAC:00-09-0f-09-00-05
l port15interfacevirtualMAC:00-09-0f-09-00-06
l port16interfacevirtualMAC:00-09-0f-09-00-07
HighAvailability
Fortinet TechnologiesInc.
137

TroubleshootingHAclusters FGCPconfigurationexamplesandtroubleshooting
l port17interfacevirtualMAC:00-09-0f-09-00-08
l port18interfacevirtualMAC:00-09-0f-09-00-09
l port19interfacevirtualMAC:00-09-0f-09-00-0a
l port2interfacevirtualMAC:00-09-0f-09-00-00 (sameasport1)
l port20interfacevirtualMAC:00-09-0f-09-00-0c
l port3interfacevirtualMAC:00-09-0f-09-00-0d
l port4interfacevirtualMAC:00-09-0f-09-00-0d (sameasport3)
l port5interfacevirtualMAC:00-09-0f-09-00-0f
l port6interfacevirtualMAC:00-09-0f-09-00-10
l port7interfacevirtualMAC:00-09-0f-09-00-11
l port8interfacevirtualMAC:00-09-0f-09-00-12
l port9interfacevirtualMAC:00-09-0f-09-00-13
4. Addthedefault route.
config router static
edit 1
set dst 0.0.0.0 0.0.0.0
set gateway 172.20.120.2
set device Port1_Port2
end
To configure HA port monitoring for the redundant interfaces
1. ConfigureHAportmonitoring fortheredundantinterfaces.
config system ha
set monitor Port1_Port2 Port3_Port4
end
TroubleshootingHA clusters
ThissectiondescribessomeHAclusteringtroubleshootingtechniques.
Ignoring hardware revisions
ManyFortiGate platforms havegonethroughmultiple hardwareversionsandinsomecasesthehardware
changespreventclusterformation. If yourunintothisproblemyoucanusethefollowing commandoneach
FortiGate tocausetheclustertoignoredifferent hardwareversions:
execute ha ignore-hardware-revision enable
ThiscommandisonlyavailableonFortiGates that havehadmultiple hardwarerevisions.
Bydefault thecommandissettopreventclusterformation betweenFortiGates withdifferent hardwarerevisions.
Youcanenterthefollowing commandtoviewitsstatus:
execute ha ignore-hardware-revision status
Usuallytheincompatibility iscausedbydifferent hardwareversionshavingdifferent harddisksandenablingthis
commanddisableseachFortiGate'sharddisks.Asaresultofdisablingharddiskstheclusterwillnotsupport
loggingtotheharddiskorWANOptimization.
138 HighAvailability
Fortinet TechnologiesInc.

FGCPconfigurationexamplesandtroubleshooting TroubleshootingHAclusters
If theFortiGates dohavecompatible hardwareversionsorifyouwanttorunaFortiGate instandalonemodeyou
canenterthefollowing commandtodisableignoringthehardwarerevisionandenabletheharddisks:
execute ha ignore-hardware-revision disable
Affectedmodelsincludebutarenotlimited to:
l FortiGate-100D
l FortiGate-300C
l FortiGate-600C
l FortiGate-800C
l FortiGate-80CandFortiWiFi-80C
l FortiGate-60C
Itspossiblethat aclusterwillnotformbecausethediskpartition sizesofthecluster
unitsaredifferent. Youcanusethediagnose sys ha checksum test |
grep storage commandtocheckthediskstoragechecksumofeachclusterunit. If
thechecksumsaredifferent thenvisittheFortinet Supportwebsiteforhelpinsetting
upcompatible storagepartitions.
Before you set up a cluster
Beforeyousetupaclusteraskyourselfthefollowing questionsabouttheFortiGates that youareplanningtouse
tocreateacluster.
1. DoalltheFortiGates havethesamehardwareconfiguration?Includingthesameharddiskconfiguration?
2. DoalloftheFortiGates havethesameFortiGuard, FortiCloud, FortiClient, VDOMandFortiOSCarrierlicensing?
3. DoalltheFortiGates havethesamefirmwarebuild?
4. ArealltheFortiGates settothesameoperatingmode(NATortransparent)?
5. ArealltheFortiGates operatinginsingleVDOMmode?
6. If theFortiGates areoperatinginmultiple VDOMmodedotheyallhavethesameVDOMconfiguration?
Insomecasesyoumaybeabletoformaclusterifdifferent FortiGates havedifferent
firmwarebuilds,different VDOMconfigurations,andareindifferent operatingmodes.
However,ifyouencounterproblemstheymayberesolvedbyinstalling thesame
firmwarebuildoneachunit, andgivethem thesameVDOMconfigurationand
operatingmode. If theFortiGates intheclusterhavedifferent licenses,theclusterwill
formbutitwilloperatewiththelowestlicensinglevel.
Troubleshooting the initial cluster configuration
Thissectiondescribeshowtocheckaclusterwhenitfirststartsuptomakesurethat itisconfiguredand
operatingcorrectly.ThissectionassumesyouhavealreadyconfiguredyourHAcluster.
To verify that a cluster can process traffic and react to a failure
1. Addabasicsecuritypolicyconfigurationandsendnetworktraffic throughtheclustertoconfirmconnectivity.
Forexample,iftheclusterisinstalledbetweentheInternet andaninternalnetwork,setupabasic
internaltoexternalsecuritypolicythat acceptsalltraffic. ThenfromaPContheinternalnetwork,
HighAvailability
Fortinet TechnologiesInc.
139


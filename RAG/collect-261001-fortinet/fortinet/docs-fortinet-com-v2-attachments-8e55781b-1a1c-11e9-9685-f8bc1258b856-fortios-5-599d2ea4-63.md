---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-63
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [8030, 8212]
sha256: 958867d1312370b7e3dec393ca11f57a6b151a33e5091825a8d26882c709898f
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

HAandfailoverprotection ClustervirtualMACaddresses
Example virtual MAC addresses
AnHAclusterwithHAgroupIDunchanged(default=0)andvirtualdomainsnotenabledwouldhavethefollowing
virtualMACaddressesforinterfacesport1toport12:
l port1virtualMAC:00-09-0f-09-00-00
l port10virtualMAC:00-09-0f-09-00-01
l port2virtualMAC:00-09-0f-09-00-02
l port3virtualMAC:00-09-0f-09-00-03
l port4virtualMAC:00-09-0f-09-00-04
l port5virtualMAC:00-09-0f-09-00-05
l port6virtualMAC:00-09-0f-09-00-06
l port7virtualMAC:00-09-0f-09-00-07
l port8virtualMAC:00-09-0f-09-00-08
l port9virtualMAC:00-09-0f-09-00-
l port11virtualMAC:00-09-0f-09-00-0a
l port12virtualMAC:00-09-0f-09-00-0b
If thegroupIDischangedto34thesevirtualMACaddresseschangeto:
l port1virtualMAC:00-09-0f-09-22-00
l port3virtualMAC:00-09-0f-09-22-03
l port4virtualMAC:00-09-0f-09-22-04
l port5virtualMAC:00-09-0f-09-22-05
l port6virtualMAC:00-09-0f-09-22-06
l port7virtualMAC:00-09-0f-09-22-07
l port8virtualMAC:00-09-0f-09-22-08
l port9virtualMAC:00-09-0f-09-22-
l port11virtualMAC:00-09-0f-09-22-0a
l port12virtualMAC:00-09-0f-09-22-0b
l port10virtualMAC:00-09-0f-09-22-01
l port2virtualMAC:00-09-0f-09-22-02
AclusterwithvirtualdomainsenabledwheretheHAgroupIDhasbeenchangedto23,port5andport6areinthe
rootvirtualdomain(whichisinvirtualcluster1),andport7andport8areinthevdom_1virtualdomain(whichisin
virtualcluster2)wouldhavethefollowing virtualMACaddresses:
l port5interfacevirtualMAC:00-09-0f-09-23-05
l port6interfacevirtualMAC:00-09-0f-09-23-06
l port7interfacevirtualMAC:00-09-0f-09-23-27
l port8interfacevirtualMAC:00-09-0f-09-23-28
Displaying the virtual MAC address
EveryFortiGate physicalinterfacehastwoMACaddresses:thecurrenthardwareaddressandthepermanent
hardwareaddress.Thepermanenthardwareaddresscannotbechanged,itistheactualMACaddressofthe
interfacehardware.Thecurrenthardwareaddresscanbechanged.Thecurrenthardwareaddressistheaddress
HighAvailability
Fortinet TechnologiesInc.
217

ClustervirtualMACaddresses HAandfailoverprotection
seenbythenetwork.ForaFortiGate notoperatinginHA,youcanusethefollowing commandtochangethe
currenthardwareaddressoftheport1interface:
config system interface
edit port1
set macaddr <mac_address>
end
end
Foranoperatingcluster,thecurrenthardwareaddressofeachclusterunitinterfaceischangedtotheHAvirtual
MACaddressbytheFGCP.Themacaddr optionisnotavailableforafunctioning cluster.Youcannotchangean
interfaceMACaddressandyoucannotviewMACaddressesfromthesystem interface CLIcommand.
Youcanusetheget hardware nic <interface_ name_str> commandtodisplaybothMACaddresses
foranyFortiGate interface. Thiscommanddisplayshardwareinformation forthespecifiedinterface. Depending
ontheirhardwareconfiguration, thiscommandmaydisplaydifferent information fordifferent interfaces.Youcan
usethiscommandtodisplaythecurrenthardwareaddressasCurrent_HWaddr andthepermanenthardware
addressasPermanent_ HWaddr.ForsomeinterfacesthecurrenthardwareaddressisdisplayedasMAC.The
commanddisplaysagreatdealofinformation abouttheinterfacesoyoumayhavetoscrolltheoutput tofindthe
hardwareaddresses.
Youcanalsousethediagnose hardware deviceinfo nic <interface_ str>
commandtodisplaybothMACaddressesforanyFortiGate interface.
BeforeHAconfigurationthecurrentandpermanenthardwareaddressesarethesame.Forexampleforoneof
theunitsinCluster_1:
FGT60B3907503171 # get hardware nic internal
.
.
.
MAC: 02:09:0f:78:18:c9
Permanent_HWaddr: 02:09:0f:78:18:c9
.
.
.
DuringHAoperationthecurrenthardwareaddressbecomestheHAvirtualMACaddress,forexampleforthe
unitsinCluster_1:
FGT60B3907503171 # get hardware nic internal
.
.
.
MAC: 00:09:0f:09:00:02
Permanent_HWaddr: 02:09:0f:78:18:c9
.
.
.
Thefollowing commandoutput forCluster_2showsthesamecurrenthardwareaddressforport1asforthe
internalinterfaceofCluster_2,indicatingaMACaddressconflict.
FG300A2904500238 # get hardware nic port1
.
.
.
MAC: 00:09:0f:09:00:02
218 HighAvailability
Fortinet TechnologiesInc.

HAandfailoverprotection ClustervirtualMACaddresses
Permanent_HWaddr: 00:09:0F:85:40:FD
.
.
.
Diagnosing packet loss with two FortiGate HA clusters in the same broadcast domain
AnetworkmayexperiencepacketlosswhentwoFortiGate HAclustershavebeendeployedinthesame
broadcastdomain. DeployingtwoHAclustersinthesamebroadcastdomaincanresultinpacketlossbecauseof
MACaddressconflicts. Thepacketlosscanbediagnosedbypingingfromoneclustertotheotherorbypinging
bothoftheclustersfromadevicewithinthebroadcastdomain. YoucanresolvetheMACaddressconflictby
changingtheHAGroupIDconfigurationofthetwoclusters.TheHAGroupIDissometimes alsocalledthe
ClusterID.
Thissectiondescribesatopologythat canresultinpacketloss,howtodetermineifpacketsarebeinglost, and
howtocorrecttheproblembychangingtheHAGroupID.
PacketlossonanetworkcanalsobecausedbyIPaddressconflicts. FindingandfixingIP
addressconflictscanbedifficult. However,ifyouareexperiencingpacketlossandyour
networkcontainstwoFortiGate HAclustersyoucanusetheinformation inthisarticleto
eliminate onepossiblesourceofpacketloss.
Changing the HA group ID to avoid MAC address conflicts
ChangetheGroupIDtochangethevirtualMACaddressofallclusterinterfaces.YoucanchangetheGroupID
fromtheFortiGate CLIusingthefollowing command:
config system ha
set group-id <id_integer>
end
Example topology
Thetopologybelowshowstwoclusters.TheCluster_1internalinterfacesandtheCluster_2port1interfacesare
bothconnectedtothesamebroadcastdomain. Inthistopologythebroadcastdomaincouldbeaninternal
network.BothclusterscouldalsobeconnectedtotheInternet ortodifferent networks.
HighAvailability
Fortinet TechnologiesInc.
219

ClustervirtualMACaddresses HAandfailoverprotection
Example HA topology with possible MAC address conflicts
Ping testing for packet loss
If thenetworkisexperiencingpacketloss,itispossiblethat youwillnotnoticeaproblemunlessyouare
constantlypingingbothHAclusters.Duringnormaloperationofthenetworkyoualsomight notnoticepacketloss
becausethelossratemaynotbesevereenoughtotimeout TCPsessions.AlsomanycommontypesifTCP
traffic, suchaswebbrowsing,maynotbegreatlyaffected bypacketloss.However,packetlosscanhavea
significant effect onrealtime protocolsthat deliveraudioandvideodata.
Totestforpacketlossyoucansetuptwoconstantpingsessions,onetoeachcluster.If packetlossisoccurring
thetwopingsessionsshouldshowalternating repliesandtimeouts fromeachcluster.
Cluster_1 Cluster_2
reply timeout
reply timeout
reply timeout
timeout reply
220 HighAvailability
Fortinet TechnologiesInc.

HAandfailoverprotection Synchronizingtheconfiguration
Cluster_1 Cluster_2
timeout reply
reply timeout
reply timeout
timeout reply
timeout reply
timeout reply
timeout reply
Viewing MAC address conflicts on attached switches
If twoHAclusterswiththesamevirtualMACaddressareconnectedtothesamebroadcastdomain(L2switchor
hub),theMACaddresswillconflictandbouncebetweenthetwoclusters.ThisexampleCiscoswitchMAC
addresstableshowstheMACaddressflappingbetweendifferent interfaces(1/0/1and1/0/4).
1 0009.0f09.0002 DYNAMIC Gi1/0/1
1 0009.0f09.0002 DYNAMIC Gi1/0/4
Synchronizingthe configuration
TheFGCPusesacombination ofincrementalandperiodicsynchronizationtomakesurethat theconfigurationof
allclusterunitsissynchronizedtothat oftheprimaryunit.
Thefollowing settingsarenotsynchronizedbetweenclusterunits:
l HAoverride.
l HAdevicepriority.
l Thevirtualclusterpriority.
l TheFortiGate hostname.
l TheHAprioritysetting forapingserver(ordeadgatewaydetection)configuration.
l ThesysteminterfacesettingsoftheHAreservedmanagement interface.
l TheHAdefault routeforthereservedmanagement interface, setusingtheha-mgmt-interface- gateway
optionoftheconfig system ha command.
Theprimaryunitsynchronizesallotherconfigurationsettings, includingtheotherHAconfigurationsettings.
AllsynchronizationactivitytakesplaceovertheHAheartbeatlinkusingTCP/703andUDP/703packets.
Disabling automatic configuration synchronization
Insomecasesyoumaywanttousethefollowing commandtodisableautomatic synchronizationoftheprimary
unitconfigurationtoallclusterunits.
config system ha
HighAvailability
Fortinet TechnologiesInc.
221


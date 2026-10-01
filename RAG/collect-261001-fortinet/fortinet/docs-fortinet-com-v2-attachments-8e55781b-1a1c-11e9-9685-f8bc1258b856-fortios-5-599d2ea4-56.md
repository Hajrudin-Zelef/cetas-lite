---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-56
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [7370, 7474]
sha256: 6d9ec18cc6f1690a1f538d91943c17fba0646484efde1920b34a6ea2aa9a2b32
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

Operatingclustersandvirtualclusters AddingadisconnectedFortiGate backtoitscluster
Youcanusethefollowing proceduresforastandardclusterandforavirtualclusteringconfiguration. Tousethe
following proceduresfromavirtualclusteryoumustbeloggedinastheadminadministrator andyoumusthave
selectedGlobalConfiguration.
WhenyoudisconnectaclusterunityoumustassignanIPaddressandnetmasktooneoftheinterfacesofthe
disconnectedunit. Youcandisconnectanyunitfromtheclustereventheprimaryunit. Aftertheunitis
disconnected,theclusterrespondsasifthedisconnectedunithasfailed. Theclustermayrenegotiateandmay
selectanewprimaryunit.
WhentheclusterunitisdisconnectedtheHAmodeischangedtostandalone.Inaddition, allinterfaceIP
addressesofthedisconnectedunitaresetto0.0.0.0 exceptfortheinterfacethat youconfigure.
Otherwisetheconfigurationofthedisconnectedunitisnotchanged.TheHAconfigurationofthedisconnected
unitisnotchangedeither(excepttochangetheHAmodetoStandalone).
To disconnect a cluster unit from a cluster - GUI
1. Goto System > HA toviewtheclustermemberslist.
2. SelecttheDisconnectfromclustericonfortheclusterunittodisconnectfromthecluster.
3. Selecttheinterfacethat youwanttoconfigure.YoualsospecifytheIPaddressandnetmaskforthisinterface.
WhentheFortiGate isdisconnected,allmanagement accessoptionsareenabledforthisinterface.
4. SpecifyanIPaddressandnetmaskfortheinterface. YoucanusethisIPaddresstoconnecttotheinterfaceto
configurethedisconnectedFortiGate.
5. Select OK.
TheFortiGate isdisconnectedfromtheclusterandtheclustermayrenegotiateandselectanew
primaryunit. TheselectedinterfaceofthedisconnectedunitisconfiguredwiththespecifiedIP
addressandnetmask.
To disconnect a cluster unit from a cluster - CLI
1. Enterthefollowing commandtodisconnectaclusterunitwithserialnumberFGT5002803033050.Theinternal
interfaceofthedisconnectedunitissettoIPaddress1.1.1.1 andnetmask255.255.255.0.
execute ha disconnect FGT5002803033050 internal 1.1.1.1 255.255.255.0
Adding a disconnectedFortiGate back to its cluster
If youdisconnectaFortiGate fromacluster,youcanre-connectthedisconnectedFortiGate totheclusterby
setting theHAmodeofthedisconnectedunittomatchtheHAmodeofthecluster.Usuallythedisconnectedunit
rejoinstheclusterasasubordinateunitandtheclusterautomatically synchronizesitsconfiguration.
YoudonothavetochangetheHApasswordonthedisconnectedunitunlesstheHA
passwordhasbeenchangedaftertheunitwasdisconnected.Disconnectingaunit
fromaclusterdoesnotchangetheHApassword.
HighAvailability
Fortinet TechnologiesInc.
201

diagnosesyshadump-bycommand Operatingclustersandvirtualclusters
Youshouldmakesurethat thedevicepriorityofthedisconnectedunitislowerthanthe
devicepriorityofthecurrentprimaryunit. Youshouldalsomakesurethat theHA
override CLIoptionisnotenabledonthedisconnectedunit. Otherwise,whenthe
disconnectedunitjoinsthecluster,theclusterwillrenegotiateandthedisconnected
unitmaybecometheprimaryunit. If thishappens,theconfigurationofthe
disconnectedunitissynchronizedtoallotherclusterunits. Thisconfigurationchange
might disrupttheoperationofthecluster.
Thefollowing procedureassumesthat thedisconnectedFortiGate iscorrectlyphysicallyconnectedtoyour
networkandtotheclusterbutisnotrunninginHAmodeandnotpartofthecluster.
Beforeyoustartthisprocedureyoushouldnotethedevicepriorityoftheprimaryunit.
To add a disconnectedFortiGate back to its cluster - GUI
1. LogintothedisconnectedFortiGate.
If virtualdomainsareenabled,loginastheadminadministrator andselectGlobalConfiguration.
2. Goto System > HA.
3. ChangeModetomatchthemodeofthecluster.
4. If required,changethegroupnameandpasswordtomatchthecluster.
5. SettheDevicePrioritylowerthanthedevicepriorityoftheprimaryunit.
6. Select OK.
ThedisconnectedFortiGate joinsthecluster.
To add a disconnectedFortiGate back to its cluster - CLI
1. LogintotheCLIoftheFortiGate tobeaddedbacktothecluster.
2. Enterthefollowing commandtoaccesstheglobalconfigurationandaddtheFortiGate backtoaclusteroperating
inactive-passivemodeandsetthedevicepriorityto50(alownumber)sothat thisunitwillnotbecomethe
primaryunit:
config global
config system ha
set mode a-p
set priority 50
end
end
Youmayhavetoalsochangethegroupname, groupidandpassword.Howeverifyouhavenot
changedthesefortheclusterortheFortiGate afteritwasdisconnectedfromtheclusteryoushould
nothavetoadjustthem now.
diagnose sys ha dump-by command
Youcanusethefollowing diagnosecommandtodisplayadataaboutacluster:
diagnose sys ha dump-by {group | vcluster | rcache | debug-zone | vdom | kernel | device |
stat | sesync}
202 HighAvailability
Fortinet TechnologiesInc.

Operatingclustersandvirtualclusters diagnosesyshadump-bycommand
kernel
ThiscommanddisplaystheHA configurationstoredbythekernel.
diagnose sys ha dump-by kernel
HA information.
group_id=88, nvcluster=2, mode=2, load_balance=0, schedule=3, ldb_udp=0.
nvcluster=2, mode=2, ses_pickup=0, delay=0, load_balance=0
schedule=3, ldb_udp=0, standalone_ ha=0, upgrade_mode=0.
vcluster 1:
FGT51E5618000206, 0, 0.
FGT51E5618000259, 1, 1.
vcluster 2:
FGT51E5618000206, 1, 1.
FGT51E5618000259, 0, 0.
stat
Thiscommanddisplayssomestatisticsabouthowwelltheclusterisfunctioning. Information includespacket
counts,memoryuse,failed linksandpingfailures.
diagnose sys ha dump-by stat
HA information.
packet count = 1, memory = 220.
check_linkfails = 0, linkfails = 0, check_pingsvrfails = 2822
bufcnt = -5, bufmem = 0
HighAvailability
Fortinet TechnologiesInc.
203


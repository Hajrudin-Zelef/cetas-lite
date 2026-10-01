---
id: collect-261001-cisco/cisco/questions-18717-moving-interface-names-on-a-cisco-asa-while-maintaining-the-rest-e74936bf
title: "questions-18717-moving-interface-names-on-a-cisco-asa-while-maintaining-the-rest-e74936bf"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-cisco/questions-18717-moving-interface-names-on-a-cisco-asa-while-maintaining-the-rest-e74936bf.md
source_anchor: ""
source_lines: [1, 49]
sha256: a5302d1ca74b7534e5c27a716ad7c7d0e6f62c27d289e3e1bd5ffaee55926997
---

# questions-18717-moving-interface-names-on-a-cisco-asa-while-maintaining-the-rest-e74936bf

This is how I did the same thing.
First I changed all of my interface names to have "_REPLACE" at the end.
This ensured that I woudl only select the actual places where the interface was present.
Several items will default to being named with default interface name at the time they are created, but will not change if you rename the interface later.  Further, you may have objects you named with the interface name yourself at some point, which could also be confused.
EG:  Arbitrary_Net interface became Arbitrary_Net_REPLACE
Then I created a Port-channel across a couple of interfaces and created VLAN Interfaces Below these, putting the VLAN ID as the subinterface number and appended to the Interface name that woudl move there.
EG:  "Arbitrary_Net" interface is on VLAN 173, so PortChannel1.123 was assigned VLAN ID 123, and given the name "Arbitrary_Net_123"
However, doing it again I would amend this process to simply have "_NEW" at the end instead of as "Arbitrary_Net_NEW" as this could allow me to make changes much more quickly.
I ensured that the Interfaces had their MAC addresses setup for HA etc and that the port channels were working and some arbitrary temporary VLAN I create don the ASA and the switch showed the ASA could ping the switch and the Switch could ping the ASA to ensure there was already connectivity.
I also changed the HTTP and SSH to use LOCAL instead of AAA until changes completed Iwhen I woudl change them back, and I added and tested management access from a host on a separate network from my normal management network so that I could move that network first then complete the changes from it after confirming I still had management access.
Then I just TFTPed down the config copy run tftp://[hostIP]/[Path]/FW_Cluster A_Primary_Active_Original.asa and opened this in Notepad++ to edit.
To be cautious I just Searched all of the "_Replace" statements and placed them in a separate file (let's call this "Main File"
I reviewed these and if any were not flush to the left side I checked that area to make sure I grabbed outer commands needed from the TFTPed file.
Note: I could have done that with AAA setups but I figured it was much easier to just to let ASDM take care of making sure to re-write the commands exactly on those as AAA has to be completely removed and readded and I had half a dozen setups.  So I just discarded the AAA items until later.
Likewise, I discarded commands such as nameif and MTU that were not changing.
This left mostly a bunch of NAT rules, the ACE to interface associations, Cyrptomaps, and some routes etc. in the Main File
The Routes, I MOVED to a separate File (call it Penultimate File), here I doubled them, put no route to the old interface and then route to the new interface and grouped them by the interface.
WebVPN and HTTP Redirect I COPIED to  Penultimate File as well, and I prefixed all of these with no  appended to them.
I added a few more odds and Ends to Penultimate File which needed to be removed and then added again such as:
no management Arbitrary_NET_REPLACE
management Arbitrary_NET_NEW
When I went back to the MAIN File and was satisfied there was none of that sort of stuff left in it
So I just did a search for _REPLACE and Replaced with _New (or if you were following along I actually did a bunch of these because I was using VLAN IDs on the interfaces)
I pulled all of the NAT Rules out into an Excel as well and Ran a sh run nat | in _REPLACE On the ASA
I put the output of that command in one set of Cells and the New versions of the commands in the other set, and reviewed to make sure they were matched up.
I then used Excel to Split off the NAT Rule Numbers from the first set of Cells and add it to the second just before the "Source" keyword
IE I put the results of the show into column B, the NATs from the Main File into Column C and then** in Column E I did a substitution using excel  =SUBSTITUTE(C1, source ,LEFT(B1," "&FIND(" ",B1)&"source "))  This took the rule numbers and stuck them in right before source which is what is needed to place them in the correct order.
(**TBH before doing the substitution in column E, I used Column D to run a compare against Column B and C looking at the NAT rules from the word source onwards but I don't feel like writing in that equation at the moment from scratching my head, and the results all 8 times were as expected so my visual inspection was accurate enough)
Then in Column A, I added this Formula to get the rule numbers =LEFT(B1,FIND(" ",B1)-1)
Then I selected Columns A through E, hit Filter, and Sorted them to reverse the order of the NAT rules so they would be applied from highest number to lowest number.
I then Copied these back into the MAIN File replacing the NATs that had been there before.
Finally, I went back to the TFTPed File and copied the Interface Commands for the Old Interfaces into a Final File
These I changed the IP on them to a network completely outside the normal range, lets say the Original networks were all 10.1.x.y, I changed these all to 10.2.x.y
I changed all of the Monitor  statements to no monitor and added a shut at the end
I then Copied all of the New Interface commands in, and put int he correct IP addresses and subnet masks, and I added monitor keywords for each.
I then paired each old interface with the new so I could do these one set at a time just in case I missed some issue.
Lastly, I copied the original commands again for these original interfaces, and only retained the interface gi0/0 command and the nameif  command and appended no  to each nameif  so they were all no nameif, and placed this in a separate file for later, just in case.
FINALLY I was Ready.
I applied the changes from the Mian file.
Now I had all my ACLss associated with both the new and old interface, and all of the NAT rules were duplicated to them as well (which is not a problem because it only mattered once traffic woudl pass the given interface.)
I could review the NAT rule sin CLI or ASDM and see the new and old rules next to each other correctly which was also very helpful in confirming no issues.
I then applied the changes from the Penultimate file and Final file ONLy for the interface which I had set up alternate access sin to for ASDM and SSH, I then confirmed all was well and logged into t SSH and ASDM there to make the remaining changes.
I now applied all of the remaining changes from Penultimate file, and then All of the Changes from the FInal File.
This left me with the original interfaces still there, the ACLs and NATs still applied to them, and all of it able to be turned back on and start working with only a minor set of changes.
Once I was satisfied a few hours later that the changes were good
I Moved all of the AAA settings overusing ASDM (thank you ASDM!) and then I ran the no nameif commands I had saved in a separate file for later, and all of the remaining entries for anything to do with those interfaces were removed from the firewall cluster.
Doing this you could feasibly do the average firewall in a day or two of work once you are comfortable with the process, but I found that I was much more cautious and only the firewalls cluster with practically no NATs and VPNs seemed "quick" at all.
In part, because you are going to want to tripel check everything and even having done so you may likely forget some small thing.
However, I can say that leaving the interface changes till the very last step allowed me to deal with little gotchas well before that step so that when traffic actually moved over there was often only ping or two lost if that, usually just a little latency.

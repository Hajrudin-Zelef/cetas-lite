---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-5
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [261, 333]
sha256: ebe95051d51ff1470b43bcc2212a403c15e94835859b76a82cb384ad2629b3a4
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community Cookbook 
Version 1.11 
Page 6 of 165 © EVE-NG LTD 
10.3.4 Boot nodes from exported config set ............................................................... 150 
10.3.5 Edit exported configurations ............................................................................ 151 
10.3.6 Set lab to boot from none ................................................................................ 151 
10.3.7 Lab config script timeout .................................................................................. 152 
11 EVE TROUBLESHOOTING ....................................................................................... 153 
11.1 CLI DIAGNOSTIC INFORMATION DISPLAY COMMANDS ................................................ 153 
11.1.1 Display full EVE Community diagnostic ........................................................... 153 
11.1.2 Display the currently installed EVE Community version: ................................. 153 
11.1.3 Display if EVEs Intel VT-x/EPT option on/off: .................................................. 153 
11.1.4 Display EVEs CPU INFO: ................................................................................ 153 
11.1.5 Display EVEs HDD utilization. ......................................................................... 153 
11.1.6 Display EVEs Bridge interface status .............................................................. 154 
11.1.7 Display EVEs system services status .............................................................. 154 
11.2 EXPAND EVES SYSTEM HDD................................................................................. 154 
11.2.1 Expand HDD on VMware Workstation ............................................................. 155 
11.2.2 Expand your HDD on ESXi .............................................................................. 155 
11.2.3 Expand your HDD on a Bare Metal EVE Server ............................................. 156 
11.3 RESET MANAGEMENT IP ........................................................................................ 156 
11.4 EVE COMMUNITY SQL DATABASE RECOVERY ........................................................ 156 
11.5 EVE LOG FILES ..................................................................................................... 156 
11.6 EVE CLI DIAGNOSTIC INFO ...................................................................................... 157 
12 IMAGES FOR EVE ..................................................................................................... 158 
12.1 QEMU IMAGE NAMING TABLE ................................................................................... 158 
12.2 HOW TO PREPARE IMAGES FOR EVE ....................................................................... 159 
12.3 HOW TO ADD CUSTOM IMAGE TEMPLATE .................................................................. 159 
12.3.1 Templates folder choice ................................................................................... 159 
12.3.2 Prepare template file ........................................................................................ 159 
12.3.3 Prepare interface format and name lines......................................................... 160 
12.3.4 Edit your new template file: .............................................................................. 162 
12.3.5 Prepare new icon for your template: ................................................................ 163 
12.3.6 Template use ................................................................................................... 163 
12.4 HOW TO HIDE UNUSED IMAGES IN THE NODE LIST ..................................................... 163 
12.4.1 Creating new config.php file ............................................................................ 163 
12.4.2 Edit config.php file ............................................................................................ 164 
13 EVE RESOURCES ..................................................................................................... 165

EVE-NG Community Cookbook 
Version 1.11 
Page 7 of 165 © EVE-NG LTD 
Preface 
When I first heard about EVE-NG I was skeptical. Back then I used to Lab mainly with ESX by 
deploying many virtual Devices and connecting them manually by separate vSwitches for Point-
to-Point connections. The Problem with that was, that it was extremely time-consuming and did 
not scale - for every new Device I had to create multiple vSwitches to interconnect them with 
the virtual Machines - a Nightmare. I was in the middle of my JNCIE-Exam-Prep when I first 
saw EVE-NG on Twitter - I downloaded the Community Edition, which was the only Edition 
back then and I was amazed how easy Labbing all of a sudden was. No more deploying of 
vSwitches to interconnect nodes and boy did it Scale… 
If you follow me on Twitter you know, that I'm one of the hardest Juniper Fanboys and of course 
my Goal was to "Juniperize" EVE. I started to get in touch with Uldis and Alain and found myself 
into the Position as one of the Juniper Test Guys. Meanwhile I added nearly all Juniper related 
Devices (including cSRX and JATP) and I still test a Lot - but now on EVE-Pro. 
The Pro-Edition was a big step forward for the Project. It added some nice Features like "hot -
add-interconnect" and the Ability to use EVE-NG with multiple Users. Especially Companies 
will love EVE as it is THE Solution for Labs and PoC's. I have succe ssfully run over 30 PoC's 
in EVE and over 100 Labs (Job -Related and Personal Labs) - and I still enjoy it every day 
thanks to EVE and the a mazing Team behind it. When the Guys asked me to write the 
Introduction I was of course honored and now this Book is finally coming out to help you on 
your Quest to Setup, Run and Manage EVE-NG in a lot of possible ways. 
 
Well - enough from my Side. I hope you enjoy this Cookbook and use it wisely for your Everyday 
EVE Work. If you have Problems there is always the EVE -Forum and Live-Helpdesk - you will 
also find me there from time to time ;) 
I wish you happy reading and if you think, that this Product is amazing feel free to support it by 
buying the PRO-Edition or Donating a bit – it helps to expand this already cool Pro duct even 
more and it also honors all the work that the Guys spent in it. 
 
Christian Scholz 
@chsjuniper


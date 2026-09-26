---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-38-2
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-38.md
source_anchor: ""
source_lines: [281, 404]
sha256: a2609d63db960e21f59a839708f73341acc7f247a89057ed19cbc3c5afd0b0b3
---

# Introduction

This script is checking if a new log entry is added to a particular buffer.

In this example we will use PPPoE logs:

/system logging action
add name="pppoe"
/system logging
add action=pppoe topics=pppoe,info,!ppp,!debug

 


Log buffer will look similar to this one:

[admin@mainGW] > /log print where buffer=pppoe 
13:11:08 pppoe,info PPPoE connection established from 00:0C:42:04:4C:EE  

 


Now we can write a script to detect if a new entry is added.

```
:global lastTime;
:global currentBuf [ :toarray [ /log find buffer=pppoe  ] ] ;
:global currentLineCount [ :len $currentBuf ] ;
:global currentTime [ :totime [/log get [ :pick $currentBuf ($currentLineCount -1) ] time   ] ];
:global message "";
:if ( $lastTime = "" ) do={ 
	:set lastTime $currentTime ; 
	:set message [/log get [ :pick $currentBuf ($currentLineCount-1) ] message];
} else={
	:if ( $lastTime < $currentTime ) do={ 
		:set lastTime $currentTime ; 
		:set message [/log get [ :pick $currentBuf ($currentLineCount-1) ] message];
	}
}
```
 


After a new entry is detected, it is saved in the "message" variable, which you can use later to parse log messages, for example, to get the PPPoE client's mac addresses.

# Allow use of ntp.org pool service for NTP

This script resolves the hostnames of two NTP servers, compares the result with the current NTP settings, and changes the addresses if they're different. This script is required as RouterOS does not allow hostnames to be used in the NTP configuration. Two scripts are used. The first defines some system variables which are used in other scripts and the second does the grunt work:

# System configuration script - "GlobalVars"
:put "Setting system globals";
# System name
:global SYSname [/system identity get name];
# E-mail address to send notifications to
:global SYSsendemail "mail@my.address";
# E-mail address to send notifications from
:global SYSmyemail "routeros@my.address";
# Mail server to use
:global SYSemailserver "1.2.3.4";
# NTP pools to use (check www.pool.ntp.org)
:global SYSntpa "0.uk.pool.ntp.org";
:global SYSntpb "1.uk.pool.ntp.org";

 
```
# Check and set NTP servers - "setntppool"
# We need to use the following globals which must be defined here even
# though they are also defined in the script we call to set them.
:global SYSname;
:global SYSsendemail;
:global SYSmyemail;
:global SYSmyname;
:global SYSemailserver;
:global SYSntpa;
:global SYSntpb;
# Load the global variables with the system defaults
/system script run GlobalVars
# Resolve the two ntp pool hostnames
:local ntpipa [:resolve $SYSntpa];
:local ntpipb [:resolve $SYSntpb];
# Get the current settings
:local ntpcura [/system ntp client get primary-ntp];
:local ntpcurb [/system ntp client get secondary-ntp];
# Define a variable so we know if anything's changed.
:local changea 0;
:local changeb 0;
# Debug output
:put ("Old: " . $ntpcura . " New: " . $ntpipa);
:put ("Old: " . $ntpcurb . " New: " . $ntpipb);
# Change primary if required
:if ($ntpipa != $ntpcura) do={
    :put "Changing primary NTP";
    /system ntp client set primary-ntp="$ntpipa";
    :set changea 1;
    }
# Change secondary if required
:if ($ntpipb != $ntpcurb) do={
    :put "Changing secondary NTP";
    /system ntp client set secondary-ntp="$ntpipb";
    :set changeb 1;
    }
# If we've made a change, send an e-mail to say so.
:if (($changea = 1) || ($changeb = 1)) do={
    :put "Sending e-mail.";
    /tool e-mail send \
        to=$SYSsendemail \
        subject=($SYSname . " NTP change") \
        from=$SYSmyemail \
        server=$SYSemailserver \
        body=("Your NTP servers have just been changed:\n\nPrimary:\nOld: " . $ntpcura . "\nNew: " \
          . $ntpipa . "\n\nSecondary\nOld: " . $ntpcurb . "\nNew: " . $ntpipb);
    }
```
 Scheduler entry:

/system scheduler add \
  comment="Check and set NTP servers" \
  disabled=no \
  interval=12h \
  name=CheckNTPServers \
  on-event=setntppool \
  policy=read,write,test \
  start-date=jan/01/1970 \
  start-time=16:00:00

 # Other scripts

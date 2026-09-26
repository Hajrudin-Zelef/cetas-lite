---
id: collect-260926-mikrotik/mikrotik/generate-backup-and-rsc-based-on-router-name-date-time
title: "generate-backup-and-rsc-based-on-router-name-date-time"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2016-08-02"]
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/generate-backup-and-rsc-based-on-router-name-date-time.md
source_anchor: ""
source_lines: [1, 214]
sha256: c014e42519ce565f3b55c3d8023b0726d42a5584e07492d22ee3a2e97dbfee59
---

# generate-backup-and-rsc-based-on-router-name-date-time

Hi,

I tried to execute the following script to get a backup and .rsc file (with name + date + time) but it gave me errors:

**:log info “STARTING BACKUP”**

**:global backupfile ([/system identity get name] . “-” . [/system clock get date] . “-” . [/system clock get time])**

**/system backup save name=$backupfile**

**:log info “DELAY 3S”**

**:delay 3s**

**:log info “GENERATING RSC”**

**:global rsc ([/system identity get name] . “-” . [/system clock get date] . “-” . [/system clock get time])**

**/export file=$rsc**

**:log info “BACKUP FINISHED”**

It gives me:

**#[in backup section]**

**Saving system configuration**

**Failed to save system configuration backup**

**action failed (6)**

**#[in rsc section]**

**file operation error**

I don’t know, may be system doesn’t accept some special character.

I was not able to copy RB-16:42:00.rsc to my local computer. I suspect it could be the colons in the name.

How do I can change the colon to dash, for example, via script?

Is there a solution to these errors?

 
                
            
           
          
            
            
              hello Tomasi and welcome aboard,

The problem occurs at line 3, when you try to give a name containing a backslash (in date format) and a column (in time format).

Your script produce a name like “RouterName-aug**/**02**/**2016-09**:**53**:**08.backup”. Underlined characters are not allowed for almost every OS not only ROS.

So, you have to make a few changes in order to remove these characters





```
:log info "STARTING BACKUP";
:local filename;
:local date [/system clock get date];
:local time [/system clock get time];
:local name [/system identity get name];
:local months ("jan","feb","mar","apr","may","jun","jul","aug","sep","oct","nov","dec");
:local varHour [:pick $time 0 2];
:local varMin [:pick $time 3 5];
:local varSec [:pick $time 6 8];
:local varMonth [:pick $date 0 3];
:set varMonth ([ :find $months $varMonth -1 ] + 1);
:if ($varMonth < 10) do={ :set varMonth ("0" . $varMonth); }
:local varDay [:pick $date 4 6];
:local varYear [:pick $date 7 11];
:set filename ($name. "-" .$varYear."-".$varMonth."-".$varDay."-".$varHour.$varMin.$varSec);
/system backup save name=$filename;
:log info "DELAY 3S"
:delay 3s;
:log info "GENERATING RSC";
:global rsc $filename;
/export file=$rsc;
:log info "BACKUP FINISHED";
```

RouterName-2016-08-02-095308.backup

RouterName-2016-08-02-095308.rsc

             
            
           
          
            
            
              Thank you for your reply   

I tested the new script in my RB2011 6.34.4, but it gave me some errors. Another issue was it did not generate the RSC.

Some idea about this?

 
            
           
          
            
            
              I’ve changed a bit your example (**all local to global**).

Now the BACKUP and RSC were sucessfully created.

Can I have problems using global instead of local?

```
:log info "STARTING BACKUP";
:global filename;
:global date [/system clock get date];
:global time [/system clock get time];
:global name [/system identity get name];
:global months ("jan","feb","mar","apr","may","jun","jul","aug","sep","oct","nov","dec");
:global hour [:pick $time 0 2];
:global min [:pick $time 3 5];
:global sec [:pick $time 6 8];
:global month [:pick $date 0 3];
:set month ([ :find $months $month -1 ] + 1);
:if ($month < 10) do={ :set month ("0" . $month); }
:global day [:pick $date 4 6];
:global year [:pick $date 7 11];
:set filename ($name. "-" .$year."-".$month."-".$day."-".$hour.$min.$sec);
/system backup save name=$filename;
:log info "DELAY 3S"
:delay 3s;
:log info "GENERATING RSC";
:global rsc $filename;
/export file=$rsc;
:log info "BACKUP FINISHED"
```

 
            
           
          
            
            
              I am not an expert on programming, but i think that a global variable consume more resources.

Because a local’s variable value discarded from memory when script execution ends but a global’s variable value is not discarded from memory and you can use it later from the same or another script.

             
            
           
          
            
            
              Maybe someone finds the below useful - a modification of known script I use to send both attachments in a single email to my Gmail account.

Please notice you need ROS version at least 6.40 to be able to attach multiple files to a single email.




```
#### Modify these values to match your requirements ####
#Your email address to receive the backups
:local toemail "nice.guy@gmail.com"
#The From address (you can use your own address if you want)
:local fromemail "MikrotikBackup"
#A mail server your machines can send through
:local emailserver "smtp.gmail.com"
############## Don’t edit below this line ##############
:local textfilename
:local backupfilename
:local allfiles
:local filedate 
:local date [/system clock get date];
:local time [/system clock get time];
:local sysname [/system identity get name];
:local months ("jan","feb","mar","apr","may","jun","jul","aug","sep","oct","nov","dec");
:local varMonth [:pick $date 0 3];
:set varMonth ([ :find $months $varMonth -1 ] + 1);
:if ($varMonth < 10) do={ :set varMonth ("0" . $varMonth); }
:local varDay [:pick $date 4 6];
:local varYear [:pick $date 7 11];
:set filedate ($varYear."-".$varMonth."-".$varDay);
:set textfilename ($"filedate" . "-" . $"sysname" . ".rsc")
:set backupfilename ($"filedate" . "-" . $"sysname" . ".backup")
:set allfiles ($"textfilename" . "," . $"backupfilename")
:execute [/export file=$"textfilename"]
:execute [/system backup save name=$"backupfilename"]
#Allow time for export to complete
:delay 2s
#email copies
:log info "Emailing backups"
/tool e-mail send to=$"toemail" from=$"fromemail" server=[:resolve $emailserver] port=587 subject="[Config Backup] $sysname" file=$"allfiles"
#Allow time to send
:delay 10s
#delete copies
/file remove $textfilename
/file remove $backupfilename
```

And of course you’d need to set your settings for /tool e-mail service, in my case I use similar to

```
       address: smtp.gmail.com
          port: 587
     start-tls: yes
          from: <nice.guy@gmail.com>
          user: nice.guy@gmail.com
      password: yourgmailpassword
```

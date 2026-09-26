---
id: collect-260926-mikrotik/mikrotik/unable-to-update-firmware-to-7-16-1
title: "unable-to-update-firmware-to-7-16-1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/troubleshooting/unable-to-update-firmware-to-7-16-1.md
source_anchor: ""
source_lines: [1, 95]
sha256: 220c02702139fac1dfe8843e64ae8fb3247a39fa98c5918f67f6ecdc12d10575
---

# unable-to-update-firmware-to-7-16-1

Hello, everyone,

I have a Mikrotik hAP ac3 router with firmware version 7.13.3.

I downloaded the latest firmware 7.16.1 (stavble), loaded it into the files and rebooted the router.

On reboot the software version is the same and in the logs I see the message: ‘missing wireless package’.

Could this be the cause of the failure to update?

The router works correctly as in AP Bridge mode.

Thank you all for your help.

             
            
           
          
            
            
              Update fails due to a missing package (wireless to be precise).

Step one: identify currently installed packages (routeros and at least wireless in your case)

Step two: download all installed packages and drop them in folder. Then reboot.

Alternatively, update through /system/packages/ automatically (if this device has Internet access).

More info on upgrading:

https://help.mikrotik.com/docs/spaces/ROS/pages/328142/Upgrading+and+installation#Upgradingandinstallation-Upgrading

             
            
           
          
            
            
              Along with routeros-7.16.1-arm.npk, have you tried uploading one of the wifi-qcom-ac-7.16.1-arm.npk or wireless-7.16.1-arm.npk files? Depending on which package you are using?

             
            
           
          
            
            
              
Hi,

In fact, in the packages I have the ‘routeros’ file and the ‘wireless’ file.

Do I therefore need to update both?

I also have another AP identical to this one (where I managed to do the update) which only has ‘routeros’ as packages.

Given the difference, what are the correct packages to have in the router?

             
            
           
          
            
            
              Just use the upgrader and dont copy around npk files manually. Or do you have a special reason for doing cumbersome way?

             
            
           
          
            
            
              
In mi configuration, the AP has no DNS since it uses the ethernet port 1 as trunk for several VLAN.

I find more usefull to upload the firmware ad proceed with a reboot.

On the help.mikrotik.com I found  this: “Starting from RouterOS 7.13, the routeros (system) package and one of the wireless packages are needed for the basic operation of a simple home router.”

So the two packages are required.

             
            
           
          
            
            
              When upgrading manually, you need base OS and the needed wifi package.

For AX3 that wifi package is wifi-qcom. Nothing else.

But the question remains: why did you upgrade manually and not using the normal system/upgrade path ?

So ?

That means your config is not ok.

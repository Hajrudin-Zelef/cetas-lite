---
id: collect-260926-mikrotik/mikrotik/the-extra-packages-in-routeros-1
title: "NAME                                                                        VERSION          SCHEDULED"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/the-extra-packages-in-routeros.md
source_anchor: ""
source_lines: [1, 201]
sha256: 012bb82de0b958a751532faad9150e35a50fd3f7de5f4cae70457d8d90089201
---

# NAME                                                                        VERSION          SCHEDULED

I've a new CRS326 switch with just the RouterOS installed (no other packages installed).

Recently I got also another switch, a used CRS125. There are more packages installed:



[admin2@CRS125] /system package> print

Flags: X - disabled

# NAME                                                                        VERSION          SCHEDULED

0   routeros-mipsbe                                                             6.46.6

1   system                                                                      6.46.6

2 X ipv6                                                                        6.46.6

3   wireless                                                                    6.46.6

4   hotspot                                                                     6.46.6

5   mpls                                                                        6.46.6

6   routing                                                                     6.46.6

7   ppp                                                                         6.46.6

8   dhcp                                                                        6.46.6

9   security                                                                    6.46.6

10   advanced-tools                                                              6.46.6

```
>
Questions:
1) I think currently I only need just the basic RouterOS package. If I uninstall the other packages, how would I install them later if needed?
2) What tools are in the "advanced-tools" package"?
3) What kind of security does the "security" package offer?
4) What is in the package "system"? Is it optional?
5) What about the "routing" package? Is this an extra/advanced routing stuff?
6) Surely there must be some documentation of/about these packages. Can someone kindly tell me the link? Thx
```

             
            
           
          
          
            
            
              
Thanks, this explains it all, indeed.

A related question: What does the following from the above link do? I assume it downgrades to the previous version. Is it also possible to specify the version to downgrade to?

/system package downgrade; /system reboot;

Btw, on my CRS326 switch the “system” package is not installed. It seems necessary(?) only if the device is a router, or in router mode.

             
            
           
          
            
            
              Don’t think you can specify a specific version. Just upgrade (to the latest) or downgrade (to previous) within the channel chosen (eg. stable/long-term/…) (and I’m not even sure on the channel-flag)

             
            
           
          
            
            
              When you put package files into proper place in device and reboot device, it will instal packages … if extra packages versions match system package … or if new system package version is newer than currently installed. If you want to install older version, you have to run */system package downgrade* - which simply disables safeguard against downgrading.

IIRC changing release channel allows downgrading without playing abovementioned games.

             
            
           
          
            
            
              
Would the following method work to install a specific version: uploading the routeros.npk file to the Files directory and then rebooting the device ?

             
            
           
          
            
            
              
So, can you confirm: do i need to do these steps for installing an older version from an .npk file? :

Step 1: /system package downgrade

Step 2: upload routeros.npk to Files directory

Step 3: reboot the device

And: how can I get which “channel” is active? I think this information is missing in the above link, isn’t it?

             
            
           
          
            
            
              I have no DHCP server configured on the above said CRS125 device, so I want to disable the “dhcp” package,

but it doesn’t work: it gets scheduled for disable, but after reboot it’s still active. How come? What to do?

             
            
           
          
            
            
              In current ROS there’s bundle of packages which is what you get by downloading “Main package”. You can not uninstall individual packages if installed with bundle, you can only disable them.

But you can “unbundle” ROS by installing different version of ROS, but taking necessary packages from ZIP downloaded as “extra packages”. Bare minimum is “system” package. And there are some dependencies, security depends on dhcp (it’s been explained why dependency, IIRC it has to do with IPsec, but I forgot the details). If you need neither DHCP server nor client, you can disable package (but has to be installed).

Note that with ROS v7, Mikrotik went away from packaged installation, currently you can only install bundle and there’s no way of unbundling the installation.

You can check configured update channel by running */system package update print*.

BTW, you can not downgrade to version lower than factory-software (displayed using */system resource print*).

             
            
           
          
            
            
              
I have been able to disable most of the packages, but some won't even disable, take a look:



/system package print

Flags: X - disabled

# NAME                                                                        VERSION        SCHEDULED

0   routeros-mipsbe                                                             6.46.6

1   system                                                                      6.46.6

2 X ipv6                                                                        6.46.6

3 X wireless                                                                    6.46.6

4 X hotspot                                                                     6.46.6

5 X mpls                                                                        6.46.6

6   routing                                                                     6.46.6

7 X ppp                                                                         6.46.6

8   dhcp                                                                        6.46.6

9   security                                                                    6.46.6

10 X advanced-tools                                                              6.46.6

```
>
But Uninstalling these already Disabled packages is not possible: it gives an error saying "Couldn't perform action - can not uninstall bundled package (6)"
Is this a bug?
.
> Note that with ROS v7, Mikrotik went away from packaged installation, currently you can only install bundle and there's no way of unbundling the installation.
>
> You can check configured update channel by running > _/system package update print_> .
>
> BTW, you can not downgrade to version lower than factory-software (displayed using > _/system resource print_> ).
I want to install a version that is >= the factory-software version.
I put the .npk into the Files dir, and then rebooted the device, but it did not replace the current version with the given older version; it just deleted the .npk from the Files dir.
So, what are the correct steps to install an older version .npk?
```

             
            
           
          
            
            
              To upgrade you can use the System > Packages window and click on “Check for Upgrades”.

From there you can also change channels (stable, long-term, rc, dev)

If you want to downgrade to a specific version, then you have to upload the packages manually to the router and then from the System > Packages menu you click Downgrade.

Simply rebooting without clicking downgrade will do nothing.

When ROS sees packages uploaded, it will only try to upgrade and not downgrade after a reboot.


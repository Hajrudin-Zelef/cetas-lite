---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-54-1
title: "Description"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-54.md
source_anchor: ""
source_lines: [1, 83]
sha256: 171dd2238c26f68d9cc0239fe5aa96e705212e810ec62d1fdf2c555e8cfbdd1d
---

# Description

The **device-mode** is a feature which sets specific limitations on a device, or limits access to specific configuration options. It helps to protect your router and network from attackers who might gain unauthorized  and use it as a gateway for attacks to other networks.  

Available device-mode are *advanced,* *home, basic and rose*. Device-mode configuration is factory pre-installed to routers (for devices with MikroTik RouterOS v7.17 or later). *Advanced (previously called enterprise)* mode is configured to CCR and 1100 series devices, *home* mode is configured to home routers and basic mode to any other type of device. For devices running versions prior to RouterOS version 7.17, all devices use the *advanced/enterprise* mode. 

The device-mode can be changed by the authorized RouterOS user, but physical access to the device required to change it. After changing the device-mode, you need to confirm it, by pressing a button on the device itself, or perform a "cold reboot" - that is, unplug the power. When the change is confirmed, regardless of confirmation mode, the **device will be rebooted**! 

# Changing mode of device-mode

If no power off or button press is performed within the specified time, the mode change is canceled. If another update command is run in parallel, both will be canceled.

There are several EOL products which do not "confirm" mode changes with a reset button press. These routers can confirm mode change only with a power cycle.

In order to protect device against attacker who might silently gain access to your router, abuse it with some scripts and simply try to wait until you will reboot your router and not even know that at that time you are accepting changes requested by some intruder, you can "update" mode only three times. There is a counter which will count how many update attempts are made and will not allow any more updates. This counter can be reset only when administrator does power-cycle the router or press a button when seeing such a warning on mode settings update attempt (same as with accepting any updates).

# Changing device-mode settings using Netinstall or FlashFig

Starting from RouterOS 7.22, it is possible to configure **device‑mode** and **protected‑routerboot** and other settings.

For more information, see the Netinstall and FlashFig documentation.

# Enabling device-mode feature

The following commands are available in the /**system/device-mode** menu: 

| Property | Description | 
|---|---|
| get | Returns value that you can assign to variable or print on the screen. | 
|  | Shows the active mode and its properties. | 
| update | Applies changes to the specified properties, see below. | 

# Available device-mode modes

There are four device modes available for configuration (mode=advanced is default one), each mode has a subset of features that are not allowed when it is used. Note that __there is no mode, which has all features enabled__. Certain features need to be enabled even if you have "advanced" mode enabled. ROSE device-mode is very similar to advanced mode, but it is made for RDS and similar devices with wider option to use disks, therefore to support container installations from the factory. See section "Feature clarification" for more details about what each option means. So, as per the below table it can be seen that "traffic-gen, container, partitions, routerboard" features are always disabled, unless specifically enabled by the admin user. 

| **Feature / Property** | **Home** | **Basic** | **Advanced** | **ROSE** | 
| **Bandwidth Test** (/tool`bandwidth-test` ) | ❌ | ❌ | ✅ | ✅ | 
| **Containers** (/`container` ) | ❌ | ❌ | ❌ | ✅ | 
| **Email** (/tool`e-mail` ) | ❌ | ✅ | ✅ | ✅ | 
| **Fetch** (/tool`fetch` ) | ❌ | ✅ | ✅ | ✅ | 
| **Hotspot** (/ip`hotspot` ) | ❌ | ❌ | ✅ | ✅ | 
| **Install Any Version** (`install-any-version` ) | ❌ | ❌ | ❌ | ❌ | 
| **IPsec** (/ip`ipsec` ) | ✅ | ✅ | ✅ | ✅ | 
| **L2TP** (/interface`l2tp` ) | ✅ | ✅ | ✅ | ✅ | 
| **Partitions** (/`partitions` ) | ❌ | ❌ | ❌ | ❌ | 
| **PPTP** (/interface`pptp` ) | ✅ | ✅ | ✅ | ✅ | 
| **Proxy** (/ip`proxy` ) | ❌ | ❌ | ✅ | ✅ | 
| **RoMon** (/tool`romon` ) | ❌ | ✅ | ✅ | ✅ | 
| **Routerboard Settings** (/system`routerboard` ) | ❌ | ❌ | ❌ | ❌ | 
| **Scheduler**  (`scheduler` ) | ❌ | ✅ | ✅ | ✅ | 
| **SMB** (/ip`smb` ) | ✅ | ✅ | ✅ | ✅ | 
| **Sniffer** (/tool`sniffer` ) | ❌ | ✅ | ✅ | ✅ | 
| **SOCKS Proxy** (/ip`socks` ) | ❌ | ❌ | ✅ | ✅ | 
| **Traffic Generator** (/tool`traffic-gen` ) | ❌ | ❌ | ❌ | ❌ | 
| **ZeroTier** (`zerotier` ) | ❌ | ❌ | ✅ | ✅ | 

# List of available properties

| Property | Description | 
|---|---|
| **scheduler, socks, fetch, pptp, l2tp, bandwidth-test, traffic-gen, sniffer, ipsec, romon, proxy, hotspot, smb, email, zerotier, container, install-any-version****, partitions, routerboard**  (*yes \| no* ) | The list of available features, which can be controlled with the **device-mode** option. See section "Feature clarification" for more details about what each option means. | 
| **activation-timeout** (default:**5m** ); | The reset button or power off activation timeout can be set in range 00:00:10 .. 1d00:00:00. If the reset button is not pressed (or cold reboot is not performed) during this interval, the update will be canceled. | 
| **flagging-enabled** (*yes \| no* ; Default:**yes** ) | Device will perform configuration analysis and if traces of suspicious code are found, flagged mode will be triggered, setting **flagged=yes** , enabling restrictions described in the**flagged=yes** . See the See the "Flagged status" paragraph. | 
| **flagged** (*yes \| no* ; Default:**no** ) | RouterOS employs various mechanisms to detect tampering with it's system files. If the system has detected unauthorized access to RouterOS, the status "flagged" is set to yes. If "flagged" is set to yes, for your safety, certain limitations are put in place. See below chapter for more information. | 
| **mode:** (basic, home, advanced; default:**advanced** ); | Allows choosing from available modes that will limit device functionality. By default, **advanced** mode allows options except **traffic-gen, container, partitions,** **install-any-version****, routerboard.** So to use these features, you will need to turn it on by performing a device-mode update. By default, **home** mode disables the following features:**scheduler, socks, fetch, bandwidth-test, traffic-gen, sniffer, romon, proxy, hotspot, email, zerotier, container,** **install-any-version****, partitions, routerboard.** | 

More specific control over the available features is possible. Each of the features controlled by device-mode can be specifically turned on or off.

For instance **scheduler** won't allow to perform any action at system scheduler. Used device-mode disables all listed features, for instance  **mode**=home is used, but **zerotier** is required for your setup, device-mode update /system device-mode update zerotier=yes will be required with the physical access to device to push the button or cut the power.

### Advanced example of changing device-mode

If the update command specifies any of the mode parameters, this update replaces the entire device-mode configuration. In this case, all "per-feature" settings will be lost, except those specified with this command. For instance:

We see that fetch = yes and email = yes is missing, as they were overriden with the mode change. However, specifying only "per-feature" settings will change only those:

If the feature is disabled, an error message is displayed for interactive commands:

However, it is possible to add the configuration to a disabled feature, but there will be a comment showing the disabled feature in the device-mode:

# Feature clarification


---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-17-3
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["nand", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-17.md
source_anchor: ""
source_lines: [136, 199]
sha256: 3623de1968255d5d928c477d774b4f32b5c71547187e4594a7d3585a03b76f8d
---

# Introduction

- Select your desired RouterOS packages and press **Install.**  Wait for the installation to finish and press "**Reboot** " (Devices without serial console have to be rebooted manually).

If you have downloaded RouterOS packages for multiple architectures, Netinstall will only display the appropriate architecture packages for your device after you have selected it. Unsupported packages will not appear in this window once a device is selected.

If the installation does not start (progress bar is not moving or no status is shown), then you can try closing the Netinstall application and opening it up again or try to put the device into Etherboot mode again. If you are still unable to get Netinstall working, then you should try using it on a different computer since there might be an operating system's issue that is preventing Netinstall from working properly.

The **"Keep old configuration"** process involves downloading the configuration database from the router, reinstalling the router (including disk formatting), and uploading the configuration files back to it. However, it's important to note that this process solely applies to the configuration itself and does not impact the files, including databases like the User Manager database, Dude database, and others.

After using Netinstall the device will be reset to defaults (unless you specified not to apply default configuration). Some devices are not accessible through **ether1** port with the default configuration for security reasons. Read more about Default configuration.

Option **"Keep branding"** allows you to retain the device's already installed branding package without reinstalling it using Netinstall.

- You're all set! Configure your device and reconnect it to your network. Your device should now be functioning correctly!

# Quick start guide for Linux

Download the tool from our download page (links not literal):

wget https://download.mikrotik.com/routeros/[VERSION]/netinstall-[VERSION].tar.gz

Extract it:

tar -xzf netinstall-[VERSION].tar.gz

Run the tool:

sudo ./netinstall-cli [-parameters] [address/interface] routeros-arm64-[package VERSION].npk

The tool requires privileged access and must be run as root, use sudo.

Make sure you have set the IP on your computer's interface to static IP address:

Then run the Netinstall version 6 (an example that resets the configuration upon reinstallation procedure):

Or run the Netinstall version 7 (an example that applies an empty configuration and discards the branding during the reinstallation procedure):

# Etherboot

Etherboot mode is a special state for a MikroTik device that allows you to reinstall your device using Netinstall. There are two types of booters available for use: the regular booter and the backup booter. It's essential to verify both options.

- To use the Regular booter press Ctrl+E to enter etherboot mode using the serial console or press the Reset button after a 1-2 second delay from when you power it on.
- To employ the backup booter, power OFF the device. Press the Reset button and power on your device (wait until the "USR" led is blinking then stable "On", and when the "USR" led is "Off" - release the Reset button) - the device is booting in bootp mode to reinstall RouterOS using Netinstall.

## Reset button

The **Reset** can be found on all MikroTik devices, this button can be used to put the device into Etherboot mode. An easy way to put a device into Etherboot mode using the **Reset** button is by powering off the device, hold the **Reset** button, power on the device while holding the **Reset** button and keep holding it until the device shows up in your **Netinstalll** window.

If you have set up a Protected bootloader, then the reset button's behavior is changed. Make sure you remember the settings you used to set up the Protected bootloader, otherwise you will not be able to use Eterboot mode and will not be able to reset your device.

## RouterOS

If your device is able to boot up and you are able to log in, then you can easily put the device into Etherboot mode. To do so, just connect to your device and execute the following command:

After that either reboot the device or do a power cycle on the device. Next time the device will boot up, then it will first try going into Etherboot mode. Note that after the first boot up, the device will not try going into Etherboot mode and will boot directly from NAND or from the storage type the device is using.

## Serial console

Some devices come with a serial console that can be used to put the device into Etherboot mode. To do so, make sure you configure your computer's serial console. The required parameters for all MikroTik devices (except for RouterBOARD 230 series) are as following:

For RouterBOARD 230 series devices the parameters are as following:

Make sure you are using a proper null modem cable, you can find the proper pinout here. When the device is booting up, keep pressing **CTRL+E** on your keyboard until the device shows that it is **trying bootp protocol**:

At this point your device is in Etherboot mode, now the device should show up in your Netinstall window.

---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-61-1
title: "System Requirements"
domain: mikrotik
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["acquisition", "aws", "compute", "license", "licenses"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-61.md
source_anchor: ""
source_lines: [1, 141]
sha256: c5e0e763f2144ef697cd97bd6f9219db8bee1c79bd6d52885e38e66d477e571c
---

# System Requirements

Cloud Hosted Router (CHR) is a RouterOS version intended for running as a virtual machine. It supports the x86 64-bit architecture and can be used on most of the popular hypervisors such as VMWare, Hyper-V, VirtualBox, KVM, and others. CHR has full RouterOS features enabled by default but has a different licensing model than other RouterOS versions.

# System Requirements

- Package version: RouterOS v6.34 or newer
- Host CPU: 64-bit with virtualization support
- RAM: 256MB or more
- Disk: 128MB or more
- RouterOS version 6: The maximum supported CHR virtual hard drive size is 16GB
- RouterOS version 7: The maximum amount of RAM and disk space is limited by the Linux kernel 5.6.3 and depends on the specific hardware.

The minimum required RAM depends on interface count and CPU count. You can get an approximate number by using the following formula:

- RouterOS v6 - RAM = 128 + [ 8 × (CPU_COUNT) × (INTERFACE_COUNT - 1) ]
- RouterOS v7 - RAM = 256 + [ 8 × (CPU_COUNT) × (INTERFACE_COUNT - 1) ]

**Note:** We recommend allocating at least 1024MiB of RAM for CHR instances.

## CHR has been tested on the following platforms:

- VirtualBox 6 on Linux and OS X
- VMWare Fusion 7 and 8 on OS X
- VMWare ESXi 6.5 and higher
- Qemu 2.4.0.1 on Linux and OS X
- Hyper-V on Windows Server 2008r2, 2012 and Windows 10 *(Only Generation 1 Hyper-V virtual machine is supported at the moment)*
- Xen Server 7.1

**Warning:** Hypervisors that provide paravirtualization are not supported.

## Usable Network and Disk interfaces on various hypervisors:

- ESX:
  - Network: vmxnet3, E1000
  - Disk: IDE, VMware paravirtual SCSI, LSI Logic SAS, LSI Logic Parallel

- Hyper-V:
  - Network: Network adapter, Legacy Network adapter
  - Disk: IDE, SCSI

- Qemu/KVM:
  - Network: Virtio, E1000, vmxnet3 (optional)
  - Disk: IDE, Sata, Virtio

- VirtualBox
  - Network: E1000, rtl8193
  - Disk: IDE, Sata, SCSI, SAS

**Note:** SCSI controller Hyper-V and ESX are usable just for secondary disks, system image must be used with IDE controller!

**Warning:** We do not recommend using the E1000 network interface if better synthetic interface options are available on a specific Hypervisor!

## How to Install a virtual RouterOS system with CHR images


We provide 4 different virtual disk images to choose from. Note that they are only disk images, and you can't simply run them.

- RAW disk image (.img file)
- VMWare disk image (.vmdk file)
- Hyper-V disk image (.vhdx file)
- VirtualBox disk image (.vdi file)

**Steps to install CHR**

1. Download the virtual disk image for your hypervisor from the Cloud Hosted Router section.
2. Create a guest virtual machine
3. Use the previously downloaded image file as a virtual disk drive
4. Start the guest CHR virtual machine
5. Log in to your new CHR. The default user is 'admin', without a password

Please note that running CHR systems can be cloned and copied, but the copy will be aware of the previous trial period, so you cannot extend your trial time by making a copy of your CHR. However, you are allowed to license both systems individually. To make a new trial system, you need to make a fresh installation and reconfigure RouterOS.

**Installing CHR guides**

- VMWare Fusion/Workstation, ESXi 6.5 and higher
- VirtualBox
- Hyper-V
- Amazon Web Services (AWS)
- Hetzner Cloud Installation
- Linode
- Google Compute Engine
- ProxMox
- Vultr

# CHR Licensing

The CHR (Cloud Hosted Router) has 4 license levels:

- ***free***
- ***p1****perpetual-1* ($45)
- ***p10****perpetual-10* ($95)
- ***p-unlimited****perpetual-unlimited* ($250)

The 60-day free trial license is available for all paid license levels. To get the free trial license, you have to have an account on MikroTik.com as all license management is done there.

Perpetual is a lifetime license (buy once, use forever). It is possible to transfer a perpetual license to another CHR instance. A running CHR instance will indicate the time when it has to access the account server to renew its license. If the CHR instance will not be able to renew the license it will behave as if the trial period has run out and will not allow an upgrade of RouterOS to a newer version, or package changes (such as disabling or enabling packages).

After licensing a running trial system, you **must** manually run the */system license renew* function from the CHR to make it active. Otherwise, the system will not know you have licensed it in your account. If you do not do this before the system deadline time, the trial will end and you will have to do a complete fresh CHR installation, request a new trial, and then license it with the license you had obtained.

| License | Speed limit | Price | 
|---|---|---|
| Free | 1Mbit | FREE | 
| P1 | 1Gbit | $45 | 
| P10 | 10Gbit | $95 | 
| P-Unlimited | Unlimited | $250 | 

## Paid licenses

**p1**

*p1* (perpetual-1) license level allows CHR to run indefinitely. It is limited to 1Gbps upload per interface. All the rest of the features provided by CHR are available without restrictions. It is possible to upgrade from P1 to P10 or P-Unlimited. Once the upgrade is purchased at the full price, the former license will become available for later use on your account.

**p10**

*p10* (perpetual-10) license level allows CHR to run indefinitely. It is limited to 10Gbps upload per interface. All the rest of the features provided by CHR are available without restrictions. It is possible to upgrade from P10 to P-Unlimited. Once the upgrade is purchased at the full price, the former license will become available for later use on your account.

**p-unlimited**

The *p-unlimited* (perpetual-unlimited) license level allows CHR to run indefinitely. It is the highest-tier license and it has no enforced limitations.

## Free licenses

There are several options to use and try CHR free of charge.

**free**

The *free* license level allows CHR to run indefinitely. It is limited to 1Mbps upload per interface. All the rest of the features provided by CHR are available without restrictions. To use this, all you have to do is download the disk image file from our download page and create a virtual guest.

**60-day trial**

In addition to the limited Free installation, you can also test the increased speed of P1/P10/PU licenses with a 60 trial.

You will have to have an account registered on MikroTik.com. Then you can request the desired license level for trial from your router that will assign your router ID to your account and enable the purchase of the license from your account. All the paid license equivalents are available for trial. A trial period is 60 days from the day of acquisition after this time passes, your license menu will start to show "Limited upgrades", which means that RouterOS can no longer be upgraded or change packages (disabling or enabling packages).

If you plan to purchase the selected license, you should do it within 60 days of the trial end date. If your trial ends, and there are no purchases within 2 months after it ended, the device will no longer appear in your MikroTik account. You will have to make a new CHR installation to make a purchase within the required time frame.

To request a trial license, you must run the command "**/system license renew**" from the CHR device command line. You will be asked for the username and password of your mikrotik.com account.

**before**you request a trial license. Note that this feature must be used only while CHR is running on a free type of RouterOS license. If you have already obtained a paid or trial license, do not use the regenerate feature since you will not be able to update your current key anymore

An expired CHR license means the CHR instance failed to renew its license before the "deadline-at" date by contacting the MikroTik server or that the 60-day trial period has ended. While the router continues operating at the same tier, software updates and package changes are disabled.


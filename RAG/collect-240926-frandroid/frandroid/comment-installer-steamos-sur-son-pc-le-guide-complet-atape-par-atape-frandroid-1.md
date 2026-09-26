---
id: collect-240926-frandroid/frandroid/comment-installer-steamos-sur-son-pc-le-guide-complet-atape-par-atape-frandroid-1
title: "comment-installer-steamos-sur-son-pc-le-guide-complet-atape-par-atape-frandroid"
domain: frandroid
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "ethernet", "gpus", "intel", "nvidia"]
source: docs/RAG/clean_en/frandroid/comment-installer-steamos-sur-son-pc-le-guide-complet-atape-par-atape-frandroid.md
source_anchor: ""
source_lines: [1, 78]
sha256: 0c38bda75254e8fec50460ce4fc5bf73acb42b81804fb1ddea247ccec7eb603d
---

# comment-installer-steamos-sur-son-pc-le-guide-complet-atape-par-atape-frandroid

<!-- source: https://www.frandroid.com/comment-faire/tutoriaux/tutoriels-pc/3155683_installer-steamos-sur-son-pc-le-guide-complet-etape-par-etape -->

Turning a PC into a living-room console the Steam Deck way now takes about a quarter of an hour. The process resembles a standard Linux installation, with two or three settings that make all the difference between a machine that boots on the first try and a frustrating black screen. We walk through everything, leaving nothing to chance. If you don't yet have the machine, we explain exactly how to build your own Steam Machine in 8 components, before installing its system. If you're still torn between building one or buying Valve's, our Steam Machine vs. €1,000 PC comparison sets the scene.

## What configuration to run SteamOS?

The decisive point in configuring such a gaming PC is the graphics card. SteamOS relies on AMD and Mesa open-source drivers: a Radeon RX 6000 (RDNA 2) or RX 7000 (RDNA 3) offers the best support, since these are the architectures of the Steam Deck and the Steam Machine. Desktop Intel Arc cards, such as the B580, now work too, but via a community workaround: the official installer doesn't handle them yet, and the setup documented by testers involves using an AMD card before swapping in the Arc. GeForce cards, on the other hand, are not yet supported — we'll come back to that below. For the rest of the machine, aim for an NVMe SSD dedicated to SteamOS, 16 GB of RAM to be comfortable, and keep in mind that a recent card requires a recent image, or you'll miss the appropriate Mesa drivers.

For a base that matches the Steam Machine without breaking the bank, our selection of components for a home-built Steam Machine serves as a ready-made starting point.

A vocabulary detail to set expectations: the official Intel support introduced with SteamOS 3.8 is primarily aimed at handheld consoles, such as the MSI Claw, not desktop Arc graphics cards. Running SteamOS on a desktop PC with an Arc card remains, for now, the domain of tinkerers, and Valve has announced no date for native support of dedicated Arc cards.

    To go further

            Our tips for building your own Steam Machine: 8 components, cheaper, and one trap
            

## What you need before you start

- A USB drive of at least 8 GB, 16 GB to be safe. Its contents will be erased.
- The SteamOS recovery image, available from Valve's official support page.
- A writing tool: Rufus on Windows, Balena Etcher on macOS and Linux, or the dd command for regulars.
- A keyboard and mouse for the installation screens, even if you'll end up on a controller.
- An Ethernet connection preferably, more reliable than Wi-Fi during installation.
- A backup of your files: the target disk will be completely rewritten.

If you've never touched Linux, there's nothing insurmountable: the logic of a bootable USB drive and a disk to prepare is the same everywhere. Our tutorial on installing Linux on a PC details these basics, useful before diving into SteamOS.

## Step 1: create the installation USB drive

Download the image from the SteamOS installation and repair page, the only official source.

Mouse, keyboard, and gaming headset: Logitech G's G3 series combines precision, comfort, and RGB customization. Plus, G HUB lets you tailor every setting to your style and fully enjoy your games.

The file weighs several gigabytes, so a stable connection helps. Plug in the drive, open Rufus or Balena Etcher, select the image then the drive, and start writing. Make sure you're pointing at the USB drive and not another disk — that's the classic mistake that erases the wrong data.

The procedure is identical to that of a standard Linux installation drive, if you've done it before. Once writing is complete, safely eject the drive.

    To go further

            How to install Linux on a Windows 10 PC: the complete guide
            

## Step 2: prepare the BIOS

This is the step not to rush. Restart and enter the BIOS, usually via the Del or F2 key. Disable Secure Boot: SteamOS isn't signed with Microsoft's keys, so it would refuse to boot otherwise. Stay in UEFI mode and turn off CSM if the option exists. If you can, physically unplug the machine's other disks during installation — it's the safest way not to overwrite the wrong one. Save the changes, then exit.

## Step 3: start the installation

1. On startup, open the boot menu (often F12, F11, Esc, or Del depending on the motherboard) and choose the USB drive.
2. The installer loads a few lines of Linux, then displays a KDE desktop. Launch the option that erases the disk and installs SteamOS.
3. Confirm erasing the target disk, checking one last time that it's the right one. Allow about a quarter of an hour.
4. At the end, remove the drive before rebooting, so as not to relaunch the installer.
5. SteamOS boots into its initial setup: language, region, keyboard, then network, and finally connecting to your Steam account.

If the installer refuses to start, double-check that Secure Boot is indeed disabled and change USB ports — front panels sometimes cause problems.

If you get a black screen after installation, check the boot order and make sure the right SSD is selected. If a recent card displays nothing, start over from a SteamOS 3.8 or newer image.

## Step 4: configure the system

First reflex before anything else: update. Go to Settings then System, and install the latest version of SteamOS. The machine then boots directly into the Steam Deck's controller interface, with shader pre-compilation that smooths out stutters.

To go further, switch to Desktop mode via the power menu: there you'll find a complete KDE Plasma environment, to install applications, emulators or a browser.

A few settings improve gaming comfort. In the compatibility settings, enable Steam Play for all titles, so that Proton runs the vast majority of Windows games. Add Decky Loader if you want extensions to game mode, and a launcher like Heroic or Lutris to hook up your Epic, GOG or other libraries. Before buying or launching a game, a glance at ProtonDB reveals its compatibility, and Are We Anti-Cheat Yet indicates whether its anti-cheat works under Linux.

## What you need to know before getting started

- **No official dual boot**: the installation takes the whole disk. Valve says it is working on an installer capable of coexisting with Windows, but it's not there yet. To keep both, the cleanest option remains a separate disk for each system.
- **Kernel-level anti-cheats block**: Valorant, Call of Duty, Battlefield or EA Sports games refuse to launch under Linux. An entire category of competitive games remains inaccessible.
- **No HDMI-CEC**: you won't be able to turn on the TV or control it with the controller as on a real console.
- **Secure Boot remains disabled**: keep this in mind if you repurpose the machine for another use later.

## What if you have an Nvidia card?

This is the main limitation at the moment. SteamOS does not yet support Nvidia GPUs, whose drivers remain proprietary where those of AMD and Intel are open and integrated into the Linux kernel. On an immutable system like SteamOS, Valve cannot simply add these closed drivers, it must work hand in hand with Nvidia. The publisher has indeed confirmed that it is collaborating very closely with Nvidia, with a dedicated team that is growing. The timeline remains distant: this rapprochement with Nvidia should not come to fruition before the end of 2026, more likely 2027.


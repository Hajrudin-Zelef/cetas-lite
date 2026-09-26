---
id: collect-260926-mikrotik/mikrotik/eworm-de-routeros-scripts-blob-head-readme-md-f042a0fe-1
title: "eworm-de-routeros-scripts-blob-head-readme-md-f042a0fe"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "distribution", "license"]
source: docs/RAG/lot-mikrotik/RouterOS/eworm-de-routeros-scripts-blob-head-readme-md-f042a0fe.md
source_anchor: ""
source_lines: [1, 112]
sha256: c7314177d314de2a245df517d8fb845ce88fac512bf88c4559b6d2ddae76396a
---

# eworm-de-routeros-scripts-blob-head-readme-md-f042a0fe

a collection of scripts for MikroTik RouterOS
RouterOS 
Use at your own risk, pay attention to license and warranty, and disclaimer on external links!
Latest version of the scripts require recent RouterOS to function properly. Make sure to install latest updates before you begin. This is supposed to work flawlessly with these channels:
- stable - the latest version considered stable for daily use, including
new features
- long-term - a version considered rock-solid, usually one minor version
behindstable (7.(n-1) )
New functionality or breaking changes in RouterOS are adopted fairly quick.
These changes are pushed for general availability once a version of
RouterOS supporting this had been released to the long-term channel a
reasonable time ago.
At any time you should have at least two minor versions and their bugfix releases to choose from. Often way older versions of RouterOS work just fine.
On the other hand in seldom cases and for good reasons specific scripts
may require an even newer RouterOS version, so only stable is supported
temporarily.
💡️ Hint: If in doubt have a look at the badge at the top of each page showing the minimum version required:
ℹ️ Info: The main branch is now RouterOS v7 only. If you are still
running RouterOS v6 switch to routeros-v6 branch!
The
device-mode scheduler and fetch at least, specific scripts may require additional
features.
RouterOS packages increase in size with each release. This becomes a problem for devices with 16MB storage and below, those with an ARM CPU are specifically affected.
Huge configuration and lots of scripts give an extra risk. Take care!
If you know how things work just copy and paste the
initial commands. These also support fixing an
existing but broken installation. Remember to edit and rerun
global-config-overlay!
💡️ Hint: First time users should take the long way in detail below.
Want to see it in action? I've had a presentation Repository based
RouterOS script distribution 
⚠️ Warning: Some details changed. So see the presentation, then follow the steps below for up-to-date commands.
The update script does server certificate verification, so first step is to establish trust.
RouterOS comes with a builtin trust store with several CA certificates. If you intend not to trust this store jump to download and import certificate now.
Select the fetch command to trust these builtin certificates at
least, but make sure not to drop other targets:
/certificate/settings/set builtin-trust-store=fetch;
You can skip the steps regarding download and import certificate and jump to installation of scripts now.
If you intend to download the scripts from a different location (for example from github.com) install the corresponding certificate chain.
/tool/fetch "https://rsc.eworm.de/main/certs/Root-YE.pem" dst-path="root-ye.pem";
ℹ️ Info: Note that the command above does not verify server
certificate, so if you want to be safe download with your workstations's
browser from CA's website and transfer the file to your MikroTik device:
Let's Encrypt / ISRG Root YE 
↗️ 
Then we import the certificate.
/certificate/import file-name="root-ye.pem" passphrase="";
Do not worry that the command is not shown - that happens because it contains a sensitive property, the passphrase.
For basic verification we rename the certificate and print it by subject key identifier (skid). Make sure exactly this one certificate ("Root-YE") is shown.
/certificate/set name="Root-YE" [ find where common-name="Root YE" ];
/certificate/print proplist=name,common-name,skid where skid="A3C8265A8EA14CD03563FC9B23C83AAE56F34F56";
Always make sure there are no certificates installed you do not know or want!
All following commands will verify the server certificate. For validity the certificate's lifetime is checked with local time, so make sure the device's date and time is set correctly!
Now let's download the main scripts and add them in configuration on the fly.
:foreach Script in={ "global-config"; "global-config-overlay"; "global-functions" } do={ /system/script/add name=$Script owner=$Script source=([ /tool/fetch check-certificate=yes-without-crl ("https://rsc.eworm.de/main/" . $Script . ".rsc") output=user as-value ]->"data"); };
And finally run configuration and functions. This will also add the scheduler for loading at system startup automatically.
/system/script { run global-config; run global-functions; };
💡️ Hint: You see complaints regarding syntax errors? Most likely the RouterOS on your device is too old. Check for updates!
The last step is optional: Add this scheduler only if you want the scripts to be updated automatically!
/system/scheduler/add name="ScriptInstallUpdate" start-time=startup interval=1d on-event=":global ScriptInstallUpdate; \$ScriptInstallUpdate;";
The configuration needs to be tweaked for your needs. Edit
global-config-overlay, copy relevant configuration from
global-config (the one without -overlay).
Save changes and exit with Ctrl-o.
/system/script/edit global-config-overlay source;
Additionally creating configuration snippets is supported. The script name
of these snippets has to start with global-config-overlay.d/ to make them
being loaded automatically. This allows to split off parts of the
configuration.
To apply your changes run global-config, which will automatically load
the overlay as well:
/system/script/run global-config;
This last step is required when ever you make changes to your configuration.
ℹ️ Info: It is recommended to edit the configuration using the command
line interface. If using Winbox on Windows OS, the line endings may be
missing. To fix this run:
/system/script/set source=[ :tocrlf [ get global-config-overlay source ] ] global-config-overlay;
To update existing scripts just run function $ScriptInstallUpdate. If
everything is up-to-date it will not produce any output.
$ScriptInstallUpdate;
If the update includes news or requires configuration changes a notification is sent - in addition to terminal output and log messages.
To add a script from the repository run function $ScriptInstallUpdate with
a comma separated list of script names.
$ScriptInstallUpdate check-certificates,check-routeros-update;
Most scripts are designed to run regularly from
scheduler check-routeros-update, so let's run it daily to make sure not to
miss an update.
/system/scheduler/add name="check-routeros-update" interval=1d start-time=startup on-event="/system/script/run check-routeros-update;";
Some events can run a script. If you want your DHCP hostnames to be available
in DNS use dhcp-to-dns with the events from dhcp server. For a regular
cleanup add a scheduler entry.
$ScriptInstallUpdate dhcp-to-dns,dhcpv4-server-lease;
/ip/dhcp-server/set lease-script="dhcpv4-server-lease" [ find ];
/system/scheduler/add name="dhcp-to-dns" interval=5m start-time=startup on-event="/system/script/run dhcp-to-dns;";
There's much more to explore... Have fun!
- Find and remove access list duplicates (accesslist-duplicates )
- Upload backup to Mikrotik cloud (backup-cloud )
- Send backup via e-mail (backup-email )
- Save configuration to fallback partition (backup-partition )
- Upload backup to server (backup-upload )
- Download packages for CAP upgrade from CAPsMAN (capsman-download-packages )
- Run rolling CAP upgrades from CAPsMAN (capsman-rolling-upgrade )
- Renew locally issued certificates (certificate-renew-issued )
- Renew certificates and notify on expiration (check-certificates )
- Notify about health state (check-health )
- Notify on LTE firmware upgrade (check-lte-firmware-upgrade )
- Check perpetual license on CHR (check-perpetual-license )
- Notify on RouterOS update (check-routeros-update )
- Collect MAC addresses in wireless access list (collect-wireless-mac )
- Use wireless network with daily psk (daily-psk )
- Comment DHCP leases with info from access list (dhcp-lease-comment )
- Create DNS records for DHCP leases (dhcp-to-dns )

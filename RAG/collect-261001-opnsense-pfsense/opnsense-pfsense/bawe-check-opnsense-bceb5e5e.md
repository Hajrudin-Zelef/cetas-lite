---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/bawe-check-opnsense-bceb5e5e
title: "bawe-check-opnsense-bceb5e5e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/bawe-check-opnsense-bceb5e5e.md
source_anchor: ""
source_lines: [1, 140]
sha256: a7f8942e6b4700d6f57cbcce66e1a37c743448a9cf5156a89e80b3bdf89809a4
---

# bawe-check-opnsense-bceb5e5e

Icinga check command for OPNsense firewall monitoring

This check command depends on the following python modules:

- enum
- requests
- argparse

**Installation on Debian / Ubuntu**

```
apt install python3 python3-requests
```
**Installation on Rocky / Alma Linux 9**

```
yum install python3 python3-requests
```
**Installation on FreeBSD**

```
pkg install python3 py39-requests
```
Add a check command definition and a service to Icinga2.

Use `./check_opnsense.py -h` to get instructions:

```
usage: check_opnsense.py [-h] -H HOSTNAME [-p PORT] --api-key API_KEY --api-secret API_SECRET [-k] -m {updates,ipsec,interfaces,services,wireguard,disk,memory}
                         [-w TRESHOLD_WARNING] [-c TRESHOLD_CRITICAL] [-v] [-f FILTER]
Check command OPNsense firewall monitoring
options:
  -h, --help            show this help message and exit
API Options:
  -H, --hostname HOSTNAME
                        OPNsense hostname or ip address
  -p, --port PORT       OPNsense https-api port
  --api-key API_KEY     API key (See OPNsense user manager)
  --api-secret API_SECRET
                        API key (See OPNsense user manager)
  -k, --insecure        Don't verify HTTPS certificate
Check Options:
  -m, --mode {updates,ipsec,interfaces,services,wireguard,disk,memory}
                        Mode to use.
  -w, --warning TRESHOLD_WARNING
                        Warning treshold for check value
  -c, --critical TRESHOLD_CRITICAL
                        Critical treshold for check value
  -v, --verbose         Enable verbose Output max -vvv
  -f, --filter FILTER   String that can be used in multiple modes to exclude unwanted items from the output or exit code calculation. Example: 'Disk 1, Disk 2'.
```
Go to the user manager and select the user you want to use for API access. Click the `+` icon in the `API keys` section to add a new API key, which triggers a download of a tex file containing the key and secret.

This file should look similar to this one:

```
key=w86XNZob/8Oq8aC5r0kbNarNtdpoQU781fyoeaOBQsBwkXUt
secret=XeD26XVrJ5ilAc/EmglCRC+0j2e57tRsjHwFepOseySWLM53pJASeTA3
```
For further information have a look at the opnsense documentation.

**Check for updates**

```
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET>  -m updates
[CRITICAL] There are 43 updates available, total download size is 199.1MiB. This update requires a reboot.|upgrade_packages=42 reinstall_packages=1 remove_packages=0 available_updates=43
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET>  -m updates
[WARNING] There are 14 updates available, total download size is 64.8MiB.|upgrade_packages=14 reinstall_packages=0 remove_packages=0 available_updates=14
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET>  -m updates
[OK] - System up to date|upgrade_packages=0 reinstall_packages=0 remove_packages=0 available_updates=0
```
**Check ipsec tunnel status**

```
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET>  -m ipsec
[WARNING] IPsec tunnels not connected: headquarter
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET>  -m ipsec
[OK] IPsec tunnels connected: remote-office, headquarter
```
**Check wireguard tunnel status**

Options:

- `--carp-backup-ok` exits with OK if this node is a CARP backup. WireGuard tunnels are expected to be down on the passive node.

```
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET>  -m wireguard
[OK] 2/2 Wireguard peers are online
[OK] Peer host1 is online (8.8.8.4:35376)
[OK] Peer host2 is online (8.8.8.5:34376)
```
```
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET>  -m wireguard --carp-backup-ok
[OK] Node is CARP backup. WireGuard is expected to be down.
```
**Check available disk space**

Options:

- `-w` and`-c` define maximum disk usage i.e.`-w 80` will warn if disk usage exceeds 80%
- `-f <mountpoint>` will not check`<mountpoint>` i.e.`-f /` will not check the root filesystem.

```
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET> -m disk
[OK] Disk space is ok | /=2%;80.0;90.0;0;100
[OK] / has 201G of 222G (98.0%) free disk space
```
```
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET> -m disk - w 1 -c 2.5
[WARNING] Disk space is low on 1 disk(s) | /=2%;1.0;2.5;0;100
[WARNING] / has only 201G of 222G (98.0%) free disk space
```
```
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET> -m disk - w 1 -c 2
[CRITICAL] Disk space is critically low on 1 disk(s) | /=2%;1.0;2.0;0;100
[CRITICAL] / has only 201G of 222G (98.0%) free disk space
```
```
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET> -m disk - w 1 -c 2 -f '/'
[UNKNOWN] No disks found
```
**Check memory**

Options:

- `-w` and`-c` define maximum memory usage i.e.`-w 80` will warn if memory usage exceeds 80%

Opnsense systems **without ZFS**, not using ARC:

```
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET> -m memory
[OK] Memory usage is 34% | memory=34%;80.0;90.0;0;100;
```
Opnsense systems **with ZFS**, using ARC:

```
./check_opnsense.py -H <OPNSENSE_HOSTNAME> --api-key <API_KEY> --api-secret <API_SECRET> -m memory -w 35 -c 50
[WARNING] Memory usage is 39% | memory=39%;35.0;50.0;0;100; arc_size=199MB;
Additional memory used for ARC: 199MB
```

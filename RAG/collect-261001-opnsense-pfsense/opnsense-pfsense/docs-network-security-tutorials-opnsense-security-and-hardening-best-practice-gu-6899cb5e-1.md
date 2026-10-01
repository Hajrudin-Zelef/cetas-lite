---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e-1
title: "docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e.md
source_anchor: ""
source_lines: [1, 133]
sha256: bd3c47aa8ba83fa38c8b7799e96dfd472eb23486b661ca69c60d09703dff2558
---

# docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e

Firewalls require security and hardening to work successfully as the first line of protection against cyber attacks. Firewalls safeguard the network from external threats such as viruses, hacking attempts, and illegal access. Hardening the firewall entails putting in place a variety of security measures to prevent unwanted access and lower the risk of exploitation. Regular software upgrades, removing unnecessary services, establishing access limits, and adopting robust authentication mechanisms are all part of this.

OPNsense is a secure operating system based on HardenedBSD, which provides a strong foundation for security. It includes features like packet filtering, stateful firewall, intrusion detection and prevention, vpn, and etc. While OPNsense is secure by default, you can further enhance its security.

In this article, we outline the importance of firewall security hardening and how you can increase the security of your firewall by applying the best practices for the OPNsense platform. The topics covered in this writing are as follows:

- 
Updating and Upgrading OPNsense
- 
Changing Default Passwords
- 
Enabling Two-Factor Authentication
- 
Configuring Country Blocking
- 
Disabling SSH Connections
- 
Taking Regular Backups and Protecting Backup Files
- 
Configuring OPNsense for high availability
- 
Disabling root access For the WebGUI
- 
Restricting access to Management Portal
- 
Enabling IPS/IDS
- 
Creating Notifications for Changes

After applying the best practices explained in this article on your OPNsense firewall you may read and deploy the best practices of FreeBSD security as well to enhance hardening.

## Why are Security and Hardening Important?

Systems hardening is a set of technologies, approaches, and best practices designed to decrease vulnerability in technology applications, systems, infrastructure, firmware, and other domains. Systems hardening aims to decrease security risk by removing potential attack vectors and reducing the attack surface of the system. By deleting unnecessary programs, account functions, apps, ports, permissions, and access, attackers, and malware have fewer possibilities to penetrate your IT infrastructure.

System hardening necessitates a thorough strategy for auditing, identifying, closing, and controlling any security vulnerabilities throughout the enterprise. There are several sorts of system-hardening activities that are listed below::

- 
Operating system hardening
- 
Server hardening
- 
Endpoint hardening
- 
Database hardening
- 
Network hardening

Although the fundamentals of system hardening are universal, particular tools and procedures differ depending on the type of hardening being performed. System hardening is required throughout the whole technology lifespan, from initial installation through decommissioning at the end of its useful life. Systems hardening is a requirement of regulations such as PCI DSS and HIPAA, and cyber insurers are increasingly demanding it.

Firewall hardening refers to the practice of protecting a firewall by decreasing potential vulnerabilities through configuration adjustments and specialized actions.

### What are the Benefits of Firewall Hardening?

Firewall hardening provides the following benefits:

- 
**Improved Security:** Hardening the firewall reduces the likelihood of potential security breaches and unauthorized network access.
- 
**Compliance:** Hardening the firewall can help firms comply with PCI DSS, HIPAA, and other security and compliance laws.
- 
**Improved Performance:** Optimizing the firewall's performance and reducing the danger of overload which results in sluggish network speeds or outages is possible by fortifying it.
- 
**Increased Visibility:** Hardening the firewall increases network visibility, which aids in detecting and preventing attacks in real time.
- 
**Better Risk Management:** Hardening the firewall facilitates risk management by minimizing possible threats and weaknesses.
- 
**Increased Reliability:** Reliability is increased by lowering the risk of an outage or data loss by fortifying the firewall.

## Updating and Upgrading OPNsense

Maintaining an update of a network security appliance is essential. It is even more necessary to upgrade a network firewall because smartphones, web browsers, and desktop operating systems are constantly updated. IT experts should constantly assess and enhance their enterprise's firewall architecture and deployment to enhance network security.

OPNsense should be patched according to the severity and danger of vulnerabilities. Evaluate each patch based on your organization's inherent dangers. You may easily update your OPNsense firewall via Web UI. The OPNsense project provides many tools to rapidly patch the system, revert a package to a prior (earlier version), or revert the kernel:

- `opnsense-update` : The`opnsense-update` utility provides integrated kernel and base system upgrades through remotely acquired binary sets, in addition to package upgrades via`pkg` . The syntax of the `opnsense-update` command is given below:

```
opnsense-update [-BbCdefikPpRsuVvz] [-A mirror_abi] [-a abi_hint]
[-D device] [-l directory] [-m mirror_url]
[-n mirror_dir] [-N crypto_lib] [-r release] [-t type]
```
```
opnsense-update [-cf [-bkpt]]
```
```
opnsense-update [-LSTU [-bkp]]
```
```
opnsense-update [-GKMO]
```
The `-b`, `-k`, `-p`, and `-t` options may be stacked to generate selected updates using a minor update sequence. `-bkp` is often used to update all presently installed components simultaneously. Major upgrades are triggered with `-u` instead of `-bkp`, unless `-bkp` is explicitly given.

- `opnsense-revert` : The`opnsense-revert` utility provides the option to safely install prior versions of packages included in an OPNsense release, provided that the specified mirror caches such release. The`automatic` and`vital` package flags will be reset to their anticipated settings. The options of`opnsense-revert` utility are as follows:
  - `-i` : Ignore the outcome of the signature verification.
  - `-l` : Respect active locking mechanisms. The default behavior is to remove them and continue with the reversion.
  - `-r release` : Choose the version of the package to install. Note that the release is not the version of the particular package, but rather the version included in the OPNsense release.
  - `-z` : Utilize the snapshot directory that does not utilize minor versioning.

For example, to revert `Unbound` package to the version found in OPNsense 22.7.1, you may run the following command:

```
opnsense-revert -r 22.7.1 unbound
```
To bring `Unbound` package back to the current version, you may run the next command:

```
opnsense-revert unbound
```
- `opnsense-patch` : The`opnsense-patch` function accepts all parameters as commit hashes from the upstream git repository, downloads them, and then applies them sequentially. Patches can be reversed by reapplying them, however in order to succeed, many patches must be applied in reverse. The syntax of the`opnsense-patch` utility is given below:

```
opnsense-patch [-defiNV] [-c repo_default] commit_hash ... 
```
```
opnsense-patch [-defiNV] [-a account] [-P prefix_dir] [-p patch_level] [-r repository] [-s site] commit_hash ...
```
```
opnsense-patch -l [-c repo_default]
```
```
opnsense-patch -l [-r repository]
```
## Changing Default Passwords

Passwords are one of the most crucial security measures employed today. It is essential that the administrator and all users have secure, difficult-to-guess passwords. Having a strong password is the most crucial security measure you can take.

To improve security, we recommend that you change the default password. Failure to do so exposes the client to the danger of being hacked.

### How to change the default password on OPNsense?

You may change the default OPNsense password via CLI by following the steps below:


---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-nils-ost-ansible-collection-opnsense-collection-to-streamline-my-opnsense-configura
title: "github-nils-ost-ansible-collection-opnsense-collection-to-streamline-my-opnsense-configurations"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-nils-ost-ansible-collection-opnsense-collection-to-streamline-my-opnsense-configurations.md
source_anchor: ""
source_lines: [1, 54]
sha256: febba3e0b3fb04a7be86c02d395cb3e4f3999f9d5fe5fd75db36ab6fc3133740
---

# github-nils-ost-ansible-collection-opnsense-collection-to-streamline-my-opnsense-configurations

This repository contains the `nils_ost.opnsense` Ansible Collection.

I like using OPNsense as a router for my different environments.
Usually I use the same set of functions/services and don't like to write a custom playbook everytime.
Therefore this collection (or more specific the role within) contains all configurations I've ever used, parameterized by structured variables.
Now I just have to define the variables of services I like to configure, use the role `nils_ost.opnsense.configure`.
And thats it: A collection for my purposes but available for everyone who find a need in using it ;)

- python:
  - httpx *>=0.28.1*
  - prettytable *>=3.17.0*
- httpx 
- collections:
  - community.general *>=12.2.0*
  - oxlorg.opnsense *>=25.7.8*
- community.general 

```
pip install httpx prettytable
ansible-galaxy collection install community.general --upgrade
ansible-galaxy collection install oxlorg.opnsense
```
| Name | Description | 
|---|---|
| nils_ost.opnsense.configure | configures OPNsense-Services by group or host variables | 

`ansible-galaxy collection install nils_ost.opnsense`
You can also include it in a `requirements.yml` file and install it via
`ansible-galaxy collection install -r requirements.yml` using the format:

```
collections:
  - name: nils_ost.opnsense
```
To upgrade the collection to the latest available version, run the following command:

`ansible-galaxy collection install nils_ost.opnsense --upgrade`
You can also install a specific version of the collection, for example, if you
need to downgrade when something is broken in the latest version (please report
an issue in this repository). Use the following syntax where `X.Y.Z` can be any
available version:

`ansible-galaxy collection install nils_ost.opnsense:==X.Y.Z`
See Ansible Using Collections for more details.

See the changelog.

This collection is mainly intended to be used by myself. Therefor I'm just developing the stuff I need for my current projects on a irregular basis. But if you find some benefit in this collection, feel free to use it. If you like to have some features added feel free to create a pull-request or write an issue with a feature-request and I'm going to see if I can make it happen.

Currently I imagine a more guided setup of tailscale, as this is currently quite rudimentary. If this bothers me enough I'm going to tackle it...

GNU General Public License v3.0 or later.

See LICENSE to see the full text.

---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/nils-ost-ansible-collection-unifinetworkapplication-aaab4498
title: "nils-ost-ansible-collection-unifinetworkapplication-aaab4498"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/nils-ost-ansible-collection-unifinetworkapplication-aaab4498.md
source_anchor: ""
source_lines: [1, 37]
sha256: e7eb270aa6d24ed7d11f4210d5f86fe6d1993abd274fa1c9759a6711d17dad9c
---

# nils-ost-ansible-collection-unifinetworkapplication-aaab4498

This repository contains the nils_ost.unifinetworkapplication Ansible Collection. For installing and initial configuration of a Unifi Network Application instance.
This collection has been tested against following Ansible versions: >=2.13.9.
For collections that support Ansible 2.9, please ensure you update your network_os to use the
fully qualified collection name (for example, cisco.ios.ios).
Plugins and modules within a collection may be tested with only specific Ansible versions.
A collection may contain metadata that identifies these versions.
PEP440 is the schema used to describe the versions of Ansible.
- python:
  - requests >=2.32.5
As this collection is intended to do it's dependend module calls delegate_to: localhost it's enough to apply the following installs on the ansible-controller.
pip install requests
| Name | Description | 
|---|---|
| nils_ost.unifinetworkapplication.inform_host | override inform_host for device adoption | 
| nils_ost.unifinetworkapplication.login | create una API session (login) | 
| nils_ost.unifinetworkapplication.setup | fetch npm API token (login) | 
| Name | Description | 
|---|---|
| nils_ost.unifinetworkapplication.install_with_docker | installs Unifi Network Application within docker | 
| nils_ost.unifinetworkapplication.init | initial configuration of Unifi Network Application | 
ansible-galaxy collection install nils_ost.unifinetworkapplication
You can also include it in a requirements.yml file and install it via
ansible-galaxy collection install -r requirements.yml using the format:
collections:
  - name: nils_ost.unifinetworkapplication
To upgrade the collection to the latest available version, run the following command:
ansible-galaxy collection install nils_ost.unifinetworkapplication --upgrade
You can also install a specific version of the collection, for example, if you
need to downgrade when something is broken in the latest version (please report
an issue in this repository). Use the following syntax where X.Y.Z can be any
available version:
ansible-galaxy collection install nils_ost.unifinetworkapplication:==X.Y.Z
See Ansible Using Collections for more details.
See the changelog.
This collection is mainly intended to be used by myself. Therefor I'm just developing the stuff I need for my current projects on a irregular basis. The main focus of this collection is to do the initial setup of Application, and from then on execute the deeper configuration by hand. This is caused by the fact, that I use UNA just for managing Unifi Network AccesPoints, provision them with WiFi SSIDs and VLANs...
GNU General Public License v3.0 or later.
See LICENSE to see the full text.

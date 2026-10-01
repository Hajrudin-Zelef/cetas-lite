---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-o-x-l-ansible-opnsense-ansible-collection-to-manage-opnsense-firewalls-using-their--1
title: "latest version:"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/github-o-x-l-ansible-opnsense-ansible-collection-to-manage-opnsense-firewalls-using-their-api.md
source_anchor: ""
source_lines: [1, 42]
sha256: 34d8604ea4cc0a8c50d574b384f376462244795f7d3fc95784660ad23858ed72
---

# latest version:

**Functional Tests**: - (no resources)

The httpx python module is used for API communications!

`python3 -m pip install --upgrade httpx`
Then - install the collection itself:

```
# latest version:
ansible-galaxy collection install git+https://github.com/O-X-L/ansible-opnsense.git
# stable/tested version:
ansible-galaxy collection install git+https://github.com/O-X-L/ansible-opnsense.git,26.1.11
## OR
ansible-galaxy collection install oxlorg.opnsense
```
See: Docs

If you DO NOT want to use Ansible - this fork provides you with a raw Python3 interface.

Support the Open-Source projects that make these modules possible:

- Donate to OPNsense or Buy the Business-Edition
- Contact the ansible-collection maintainer for support

Feel free to contribute to this project using pull-requests, issues and discussions!

See also: Contributing

Currently, these modules have no multi-version support!

That means each release of the collection is only compatible with the same OPNsense version!

If you need to use these modules - ensure that there is a supported release BEFORE upgrading your OPNsense firewalls!

I currently try to create a stable release once or twice a year - as this takes 10-20h of work each time for implementing API-fixes.

As this project is unfunded we do not actively check for API-changes - if you find missing functionalities you need/want to have please report it! Not all features will be implemented!

**Development States**:

not implemented => development => testing => unstable (*practical testing*) => stable


---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/github-opnsense-core-opnsense-gui-api-and-systems-backend
title: "make package CORE_NAME=my_new_name"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "open source"]
source: docs/RAG/collect-261001-opnsense-pfsense/github-opnsense-core-opnsense-gui-api-and-systems-backend.md
source_anchor: ""
source_lines: [1, 53]
sha256: 76a70b7ade02f3d07636954a1fcd4484a709ed8d038b3ea8121f87002ff093ee
---

# make package CORE_NAME=my_new_name

The OPNsense project invites developers to start contributing to the code base. For your own purposes or – even better – to join us in creating the best open source firewall available.

The build process has been designed to make it easy for anyone to build and write code. The main outline of the new codebase is available at:

Our aim is to gradually evolve to a new codebase instead of using a big bang approach into something new.

To create working software like OPNsense you need the sources and the tools to build it. The build tools for OPNsense are freely available.

Notes on how to build OPNsense can be found in the tools repository:

You can contribute to the project in many ways, e.g. testing functionality, sending in bug reports or creating pull requests directly via GitHub. Any help is always very welcome!

You can learn more about contributing on CONTRIBUTING.md.

OPNsense is and will always be available under the 2-Clause BSD license:

Every contribution made to the project must be licensed under the same conditions in order to keep OPNsense truly free and accessible for everybody.

The repository offers a couple of targets that either tie into tools.git build processes or are aimed at fast development.

A package of the current state of the repository can be created using this target. It may require several packages to be installed. The target will try to assist in case of failure, e.g. when a missing file needs to be fetched from an external location.

Several OPTIONS exist to customise the package, e.g.:

- CORE_DEPENDS: a list of required dependencies for the package
- CORE_DEPENDS_ARCH: a list of special -required packages
- CORE_ORIGIN: sets a FreeBSD compatible package/ports origin
- CORE_COMMENT: a short description of the package
- CORE_MAINTAINER: email of the package maintainer
- CORE_WWW: web url of the package
- CORE_NAME: sets a package name

Options are passed in the following form:

```
# make package CORE_NAME=my_new_name
```
In general, options are either set to sane defaults or automatically detected at runtime.

Update will pull the latest commits from the current branch from the upstream repository.

Upgrade will run the package build and replace the currently installed package in the system.

Fetch changes from the running system for all known files.

Run several syntax checks on the repository. This is recommended before issuing a pull request on GitHub.

Run the PSR12 and PEP8 style checks on MVC PHP code and Python, respectively.

For easier development you may want to use an OPNsense VM and install
the `os-debug` plugin that will offer the necessary tools.

Run several automatic sanitizers on the code base.

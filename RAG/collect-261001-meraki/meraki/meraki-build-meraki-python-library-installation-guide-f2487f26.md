---
id: collect-261001-meraki/meraki/meraki-build-meraki-python-library-installation-guide-f2487f26
title: "meraki-build-meraki-python-library-installation-guide-f2487f26"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-build-meraki-python-library-installation-guide-f2487f26.md
source_anchor: ""
source_lines: [1, 37]
sha256: 85dfb8614ee9fb3860719a92da8470077fc4e1f231f79f93ed7bd7f4b2010d54
---

# meraki-build-meraki-python-library-installation-guide-f2487f26

Meraki Python Library Installation Guide
The Meraki Dashboard API is a powerful way of interacting with your network. To streamline your development, a Python module has been written with all of the available API calls. This allows for instant interaction with API without the overhead of writing the request calls.
This guide will get you up and running with the Meraki Python library module.
For macOS users
- Install Homebrew first, if not already present, in terminal:
> ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
- Install Python 3.6
> brew install python3
- Install the requests module
> pip3 install requests
- Install Dashboard API module
> pip3 install merakiIf previously installed, upgrade > pip3 install meraki --upgrade
- To make sure your installation is successful, launch Python
> python3
- Using your API key, run the following commands to import the meraki and list the organizations
>>> from meraki import meraki
>>> meraki.myorgaccess('42a2e94ea6092c11e323bc750f368c9ede56c8ba')
- You should see the following returned
[{'id': 640799, 'name': 'Meraki API Lab'}]
For Windows users
- Download the latest 3.6 version from https://www.python.org/downloads/
- Be sure to select the “Add Python 3.6 to PATH” checkbox, then “install Now”.
- Install the requests module, in command prompt (cmd.exe)
> pip install requests
- Install Dashboard API module
> pip install merakiIf previously installed, upgrade > pip install meraki --upgrade
- To make sure your installation is successful, launch Python
> python
- Using your API key, run the following commands to import the meraki and list the organizations
>>> from meraki import meraki
>>> meraki.myorgaccess('42a2e94ea6092c11e323bc750f368c9ede56c8ba')
- You should see the following returned:
[{'id': 640799, 'name': 'Meraki API Lab'}]
Additional Resources
Text editors
- Download a developer text editor if you don’t already have one. Some options include Sublime, Atom, Visual Studio Code, and Notepad++.
- Native TextEdit on macOS or Notepad/WordPad on Windows will not suffice! In addition, watch out for these editors altering certain characters like quotes to non-ASCII encoding, which will cause problems.

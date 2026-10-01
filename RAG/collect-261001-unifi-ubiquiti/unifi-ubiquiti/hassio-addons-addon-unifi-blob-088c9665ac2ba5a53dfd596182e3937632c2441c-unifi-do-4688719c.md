---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hassio-addons-addon-unifi-blob-088c9665ac2ba5a53dfd596182e3937632c2441c-unifi-do-4688719c
title: "hassio-addons-addon-unifi-blob-088c9665ac2ba5a53dfd596182e3937632c2441c-unifi-do-4688719c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["copyright", "liability", "license", "memory", "mit license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hassio-addons-addon-unifi-blob-088c9665ac2ba5a53dfd596182e3937632c2441c-unifi-do-4688719c.md
source_anchor: ""
source_lines: [1, 89]
sha256: de39724d4c3198b6666fb31a18cb47e092e4fa7b2617e66758a6f88dd11957f0
---

# hassio-addons-addon-unifi-blob-088c9665ac2ba5a53dfd596182e3937632c2441c-unifi-do-4688719c

This add-on runs Ubiquiti Networks' UniFi Network Application software, which allows you to manage your UniFi network via the web browser. The add-on provides a single-click installation and run solution for Home Assistant, allowing users to get their network up, running, and updated, easily.
The installation of this add-on is pretty straightforward and not different in comparison to installing any other Home Assistant add-on.
- 
Click the Home Assistant My button below to open the add-on on your Home Assistant instance.
- 
Click the "Install" button to install the add-on.
- 
Check the logs of the "UniFi Network Application" to see if everything went well.
- 
Click the "OPEN WEB UI" button, and follow the initial wizard.
- 
After completing the wizard, log in with the credentials just created.
- 
Go to the settings (gears icon in the bottom left) -> System -> Advanced.
- 
Next to the Inform Host label, click the checkbox option forOverride so that is now "checked".
- 
Change the Host for Inform to match the IP or hostname of
the device running Home Assistant.
- 
Hit the "Apply Changes" button to activate the settings.
- 
Ready to go!
Note: Remember to restart the add-on when the configuration is changed.
Example add-on configuration, with all available options:
log_level: info
memory_max: 2048
memory_init: 512
Note: This is just an example, don't copy and paste it! Create your own!
The log_level option controls the level of log output by the addon and can
be changed to be more or less verbose, which might be useful when you are
dealing with an unknown issue. Possible values are:
- trace : Show every detail, like all called internal functions.
- debug : Shows detailed debug information.
- info : Normal (usually) interesting events.
- warning : Exceptional occurrences that are not errors.
- error : Runtime errors that do not require immediate action.
- fatal : Something went terribly wrong. Add-on becomes unusable.
Please note that each level automatically includes log messages from a
more severe level, e.g., debug also shows info messages. By default,
the log_level is set to info, which is the recommended setting unless
you are troubleshooting.
This option allows you to change the amount of memory the UniFi Network Application is allowed to consume. By default, this is limited to 256 MB. You might want to increase this, in order to reduce CPU load or reduce this, in order to optimize your system for lower memory usage.
This option takes the number of Megabyte, for example, the default is 256.
This option allows you to change the amount of memory the UniFi Network Application will initially reserve/consume when starting. By default, this is limited to 128MB.
This option takes the number of Megabyte, for example, the default is 128.
The UniFi Network Application ships with an automated backup feature. This feature works but has been adjusted to put the created backups in a different location.
Backups are created in /backup/unifi. You can access this folder using
the normal Home Assistant methods (e.g., using Samba, Terminal, SSH).
Alternatively to setting up a custom inform address (installation steps 7-9) you can manually adopt a device by following these steps:
- SSH into the device using ubnt as username andubnt as password
- $ mca-cli
- $ set-inform http://<IP of Hassio>:<controller port (default:8080)>/inform  - for example $ set-inform http://192.168.1.14:8080/inform
- for example 
- 
The AP seems stuck in "adopting" state: Please read the installation instructions carefully. You need to change some controller settings in order for this add-on to work properly. Using the Ubiquiti Discovery Tool, or SSH'ing into the AP and setting the INFORM after adopting will resolve this. (see: Manually adopting a device)
- 
The following error can show up in the log, but can be safely ignored:   INFO: I/O exception (java.net.ConnectException) caught when processing
  request: Connection refused (Connection refused)
This is a known issue, however, the add-on functions normally.
- 
Due to security policies in the UniFi Network Application software, it is currently impossible to add the UniFI web interface to your Home Assistant frontend using a panel_iframe .
- 
The broadcast feature of the EDU type APs are currently not working with this add-on. Due to a limitation in Home Assistant, is it currently impossible to open the required "range" of ports needed for this feature to work.
- 
This add-on cannot support Ingress due to technical limitations of the UniFi software.
- 
During making a backup of this add-on via Home Assistant, this add-on will temporary shutdown and start up after the backup has finished. This prevents data corruption during taking the backup.
This repository keeps a change log using GitHub's releases functionality. The format of the log is based on Keep a Changelog.
Releases are based on Semantic Versioning, and use the format
of MAJOR.MINOR.PATCH. In a nutshell, the version will be incremented
based on the following:
- MAJOR : Incompatible or major changes.
- MINOR : Backwards-compatible new features and enhancements.
- PATCH : Backwards-compatible bugfixes and package updates.
Got questions?
You have several options to get them answered:
- The Home Assistant Community Add-ons Discord chat server for add-on support and feature requests.
- The Home Assistant Discord chat server for general Home Assistant discussions and questions.
- The Home Assistant Community Forum.
- Join the Reddit subreddit in /r/homeassistant
You could also open an issue here GitHub.
The original setup of this repository is by Franck Nijhof.
For a full list of all authors and contributors, check the contributor's page.
MIT License
Copyright (c) 2018-2024 Franck Nijhof
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-b290fa57-1
title: "platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-b290fa57"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-b290fa57.md
source_anchor: ""
source_lines: [1, 40]
sha256: eb9b6e1b4107c0c664854d1cdd346698bb46179c71a7d9e146dd163e06fb4a67
---

# platform-management-sm-endpoint-management-operate-and-maintain-apps-and-softwar-b290fa57

Installing Custom Apps on Windows and Mac Devices
One of the most powerful features of Systems Manager is its ability to remotely install and uninstall applications from your devices. The Software Installer's flexibility allows for one-click deployment of applications across all of your devices, regardless of operating system.
To see information on how to deploy scripts to Windows and Mac devices, see this article.
This feature is not available for Legacy Systems Manager users.
Prerequisite: Installing applications on Mac and Windows devices requires that the agent be installed. Read more about agent installation here.
Note: macOS devices can install apps using the device's native Mac App Store in conjunction with Apple Volume Purchase Program via your Apple School Manager or Apple Business Manager account and Meraki Systems Manager. To do this, enroll the macOS device with the Meraki Management enrollment profile. Detailed steps for Mac App Store app deployments within Systems Manager can be found here.
Installing Applications on Devices
Using Systems Manager, suites of applications can very easily be deployed to end user devices.
The following instructions outline how to deploy a new application, as well as overview additional installation options:
- Navigate to Systems Manager > Manage > Apps
- Click on the + Add app button and select an appropriate App platform:
 
- Once an App platform Custom app is selected, fill in the following information about the application:
- Name (required): The name of the application as it will appear on the end device (e.g. Firefox, FileZilla, etc). If the application is already installed on managed clients, it will be listed in the drop-down menu. This name can be changed at any time.
The application name must exactly match the application name that is detected by the end device.
- Identifier (optional): The vendor specified app identifier.
- Icon URL (optional): A URL containing the app icon which will be used for the custom app.
- Vendor (optional): The vendor of the specified application.
- Version (optional): The version number of the application to be installed.
Note: If no version is entered, Dashboard defaults to "0", and the version is then required when editing the app.
- Description (optional): A brief description of the application.
- Source (required): This option denotes whether the application's installer will be hosted in Dashboard, or on an external server (e.g. Dropbox).
    
  - Upload to the Meraki Cloud: Select this option and click Browse to upload the installer file directly to Dashboard.
  - Specify a URL: Select this option if the installer file is hosted on an external server or file share site, and specify the URL for the hosted file. The URL field must point to the direct download link to a publicly hosted file or an internally hosted server accessible by the end user's device.
Note: The installer must be silent (require no user interaction) in order for the application to install correctly in the background. Windows applications can be installed in the foreground in order to prompt user interaction.
Note: The following installer file types are supported:
Windows: .exe, .msi
macOS:  .dmg, .pkg
Note: macOS may install an .app but it must be encapsulated inside a .dmg image or within a .pkg package. For more information, review Apple's Developer Documentation.
- Keep app up to date: Specifies if the app can automatically keep app up to date with the latest version available. When a targeted device checks in and an app update is available, an app upgrade request will be automatically sent to the device.
- Auto-install: By default, the application will be auto-installed the moment the Save button is clicked. If auto-install is disabled, the application won't be installed until it's manually pushed by navigating to Systems manager > Manage > Apps, and selecting Push > Push to all. Dashboard users can manually queue install requests via the device page, or by pressing the Repush buttons on this page.
- Install in foreground (Windows-only): Allow the installer to show user prompts, instead of silently installing in the background.
- Installation arguments (optional): If the installer must be run with specific command line arguments, they can be listed here.
 The following list shows the actual install command run on the end-device, where [arguments] refers to the contents of the Installation arguments field in Dashboard:
  - EXE: application_installer.exe [arguments]
  - MSI: msiexec /quiet /i application_installer.msi [arguments]
  - PKG: /usr/sbin/installer [arguments] -pkg application_installer.pkg -target /
- Command line (optional): Specifies a command to run after installation has completed. This is commonly used to reboot the machine after installation, using the following commands:
    

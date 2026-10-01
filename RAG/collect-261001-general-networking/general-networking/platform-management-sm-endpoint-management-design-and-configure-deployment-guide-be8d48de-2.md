---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de-2
title: "platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Samsung"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de.md
source_anchor: ""
source_lines: [47, 83]
sha256: fe1baeb46b5732fbb2a297c30279dea31d9ca9738c95212ef5947260fa9ee482
---

# platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de

3. Provision the container - Once authenticated, the app will go through a quick loading screen and will be setup for Android Enterprise. A successful BYOD enrollment will result in icons on the device with an orange badge or a "Work" folder being created on the device home screen. You can uninstall the non-badged copy of the Systems Manager app, if desired, as all functionality now takes place within the badged copy of the app. To control which badged apps are provisioned, see the below section.
A sample of the screens are shown below:
Device Owner Enrollment
Device Owner mode is designed for institutionally owned devices with additional restrictions and control. Enrollment follows a similar process to BYOD, but everything is done in one step after a factory reset of the device. This step behaves slightly different depending on if a Google Managed or Meraki Managed domain is bound to Dashboard.
Device Owner mode can only be enabled after the factory reset of a device, and by default will disable all system apps unless configured otherwise. See the following section on controlling system apps.
If you have a Lollipop device (Android 5.0+) please reference this article for how to enable device owner mode.
- Google Managed - After a factory reset, follow the steps on screen until prompted for a Google Account. Sign in with an account that belongs to the bound Google Domain. This will prompt the installation of the SM app and automatically enroll the device in Dashboard.
- Meraki Managed - After factory reset and on the very first screen displayed at startup, tap six times in a row anywhere on the screen (six times in one spot). This should launch the camera in a QR code scan mode, and force the device into its Android Enterprise enrollment setup. Scan the QR code from the SM > Add devices page for the device to trigger its Android Enterprise enrollment into the desired network. The app will prompt the end user for authentication, if enabled, and finish setup.
As shown in the last image, enabling Device Owner mode removes all non-essential apps from the device.
Controlling Native System Apps
By default, all apps will be disabled when enrolling in Device Owner mode, including the default SMS and phone dialing apps. In Work Profile mode, Systems Manager will automatically create a work version of default apps, indicated with the orange briefcase, into the work profile. The applications that are installed by default or treated as 'system apps' will vary by device manufacturer - for example, Samsung devices use different dialer, camera, and SMS apps from Google Nexus or Pixel devices.
To customize which default Android apps are provisioned into Device Owner mode, or duplicated into the managed work profile, see the Controlling Android System Apps article.
Note: Stacking multiple Android System Apps payloads on a single device is not a supported configuration.
Troubleshooting Enrollment
To verify whether a client device is enrolled, check the client page by navigating to Systems Manager > Monitor > Devices. Select the client from the list and check the Management section in the left-hand column near the top of the client details page. If the organization is successfully enrolled/synced, there will be a field called Android for Work Account. If the device is enrolled in Android Enterprise, it will say Yes. If this field does not exist, then it is likely that the organization is not enrolled in Android Enterprise correctly yet.
Troubleshooting Device Playstore Account
Verify the device has successfully set up a Playstore account by launching the SM app and confirming that 'AFW account enabled' has a green checkmark. If you see a warning icon here instead of a green checkmark, tap the icon to have the SM app reprovision the local Android Enterprise Playstore account on the device. This will be downloaded at next check-in, and should install after a few minutes. The below images show a device enrolled through Work Profile mode that has been successfully enrolled with the AFW account enabled.
It is important to have a unique owner on every device for System Manager's app installation tag scoping to function as designed. When using Meraki Managed (in Org > MDM) the Owner of the device generates a local Playstore account on the device. Each device requires a unique owner, so they can have a unique Playstore account installed. If you are unsure if your device has a unique Playstore account, perform the following steps:
- Click on the device in Systems Manager > Devices.
- If the device already has an owner set, click on "Edit details" and then "Clear owner". If the device doesn't have an Owner currently set move to the next step.
- Add a new Owner. This Owner needs to be an Owner that is new, and not currently the Owner of another Android device.
- When the device performs its next check-in to Dashboard, it will automatically install a new Playstore account on the device.
Note: to create new owners and assign them to devices on a mass scale, use Owner .csv importing to assist.
Using the Android Restriction or Android Device Owner setting "Disable modifying accounts" will prevent the AFW account from being enabled/installed. Please make sure Android devices do not have this restriction enabled, or the Android device will be unable to provision the local PlayStore account onto the device.
Enable Device Restrictions
Device restrictions for Android Enterprise enabled devices can be found in Systems Manager > Manage > Settings by searching for 'Android Restrictions'. Some other Android settings include: App permissions, Restrictions, Device Owner, Kiosk Mode.
- App permissions - This setting allows for custom application permissions. Examples include denying an application access to the device's contacts, saved payment methods and even network access. Application permissions vary from app to app and a list of relevant permissions can be found using the "Fetch permissions" button that appears once an app has been selected.
- 
    Restrictions - These are general settings that can apply to all devices using Android Enterprise, both BYOD and Device Owner mode.
- Device Owner - These are a special set of restrictions that can only be applied to Android devices that are provisioned in Device Owner mode.
- Kiosk Mode - Kiosk mode allows an administrator to lock a device into a particular application. This can only be used with Android 6+ devices in Device Owner mode. See more info here.
- App permissions are not to be confused with App Settings. More about App Settings for Android Enterprise devices can be found here.
- The general Restrictions (not the one found under More Android) only apply to KNOX devices using the older version of Systems Manager.
Recommended Android Settings
In Android Device Owner and Work Profile enrollments, many things are disabled by default. To re-enable these settings, There are various Android Settings which can be applied, so please experiment with these settings to obtain the desired configurations for your organization. However, there are a few Android Settings which may result in unexpected behavior if the device does not have them applied. Adding an empty Android Device Owner payload, an empty Android Restrictions payload, and an Android Systems App adding only a single app to the block list (so all other System apps are allowed) can setup the Android devices as desired.
An empty Android Device Owner payload, like in the above example, can be important to uncheck the "Disable modifying accounts" restriction and others. After devices are setup, feel free to lock them down further if desired.
An empty Android Restrictions payload, like in the above example, will allow all other System apps to show on the device(s). Without this, Android Device Owner mode may block all system apps by default.

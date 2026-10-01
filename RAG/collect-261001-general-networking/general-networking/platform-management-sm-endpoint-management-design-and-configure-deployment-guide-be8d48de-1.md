---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de-1
title: "platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de.md
source_anchor: ""
source_lines: [1, 46]
sha256: d8db6746279438aeface0c6282dd0796376cb23b3c210a9bab29308cf251c8e5
---

# platform-management-sm-endpoint-management-design-and-configure-deployment-guide-be8d48de

Android Enterprise Deployment Guide
Introduction
This article provides deployment guidance for Android Enterprise (formerly Android for Work) with Cisco Meraki's System Manager. Android Enterprise is a platform for devices running on the Google Android mobile operating system that allows IT to manage and secure business applications using a work-specific profile. Android Enterprise comes in two different types of deployments:
- Work Profile or BYOD. In BYOD mode administrators only have control over work managed applications and settings. Systems Manager will containerize all corporate data and represent it using an orange badge icon. An administrator will have complete control over these applications, but have no visibility or control over personal applications on the device.
- Device Owner mode. In Device Owner mode administrators have complete control of the device. This type of deployment is primarily used on institutionally owned devices and include special features such as kiosk mode. (Device Owner mode can be thought of as the "Supervised" state for those familiar with iOS.)
Requirements
Managing Android Enterprise devices through Systems Manager requires:
- A bound domain: either a Gmail address used for administration for Meraki-managed domain, or a G Suite account for Google-managed domain. The following section describes these two in more depth.
- Android Android 10 (Quince Tart) +
- Device support for Google Play Services version 21.42.58+ (as of this writing)
- Device support for Google Mobile Services (GMS) especially the device_admin and managed_users feature flags. The latest requirements defined by Google can be found here.
Your devices must also support the work profile and work managed modes. If you are using an OEM device that is not listed in the following catalog, or is on an older version of Android, it may not support the full Android Enterprise suite of features available through Systems Manager.
For device compatibility recommendations, see Google's official Enterprise Device Catalog
For more info on enrollment options for Android devices, reference our article here. More information about Android Enterprise can be found here: https://www.android.com/work
The email address that is used to bind a work domain to an account is considered an admin email. Changing this email by going to Organization > Configure > MDM > Android Enterprise and clicking on unenroll, and then introducing a new admin email, will require the re-enrollment of all Android devies into the existing SM network.
Deployment Considerations
There are 5 main stages in an Android Enterprise deployment on Systems Manager:
- Determine and Bind a Work Domain
- Enable Authentication as a part of Enrollment
- Enroll a Device
- Enable Device Restrictions
- Push Applications
Determine and Bind a Work Domain
There are two flavors of Android Enterprise: Google Managed Domain and Meraki Managed Domain.
- Google Managed Domain - This is an Android Enterprise deployment that capitalizes on existing Google services. If services such as Gmail, Google Calendar, Google Docs, etc. are being used, it is likely a Google Managed Domain. This can be enabled in the Google Admin Console as a super administrator. Navigate to Devices > Mobile & endpoints > Settings > Third-party integrations and copy the token. This will be entered in the first step of the process. Check the "Enforce EMM policies on Android Devices" to require SM be installed on the device in order to access Google services.
If the free Android Enterprise subscription has not already been added to the Google Domain please reference the following article to enable it. The section that states "If you are a G Suite customer" provides more information about enabling the free subscription: https://support.google.com/work/andr.../6174046?hl=en
- Meraki Managed Domain - If no Google services are currently being used, Meraki can generate a Managed Domain for your Android Enterprise deployment, which may be preferable to setting up a G Suite domain that otherwise may not be used. All that is needed is a Google supported administrative email address (i.e. any @gmail.com account). In Google documentation this is referred to as an Android Enterprise account (as these accounts can only be used for Android Enterprise).
More about this can be read here: https://support.google.com/googlepla..._topic=7042018
Google Managed Domain
To bind an existing Google Managed Domain navigate to Organization > MDM, enter the domain name (e.g. 'meraki.com'), followed by the token copied from the Google Admin Console and click "Enroll Domain."
Meraki Managed Domain
To bind a Meraki Managed Domain navigate to Organization > MDM and click "Get signup URL".
Next click the URL generated that appears in step 2 and it will redirect to the "Bring Android to Work" page. Click through the form to complete and create a Meraki Managed account. If possible, it is recommended to use a Gmail account associated with your organization and not a personal account.
Once the "Complete Registration" button has been clicked, return to the Meraki Dashboard. Under Organization > MDM, there should now be a bound domain associated to the email used to complete the "Bring Android to Work" page.
Enable Authentication as a part of Enrollment
Adding authentication is a necessary step in order to associate a user to the Android Enterprise profile placed onto a device. To enable authentication in Systems Manager, navigate to Systems Manager > General and select an option in the section labeled User authentication settings.
If a Google Managed Domain was used SM will automatically authenticate (via O-auth) against the associated Google domain. However if a Meraki Managed Domain was used, please select "Managed: User Meraki hosted accounts." If no user accounts have been created, click on the Configure Meraki hosted user, after clicking Save. The username and password entered as a Meraki Owner is what SM will authenticate against.
Enrolling a Device
As mentioned earlier there are two ways to deploy Android Enterprise: BYOD mode or Device Owner mode. Each of these modes have different enrollment paths detailed below. Additional details and recommendations on choosing between the two for your deployment can be found in this article.
Google requires that Android 5.0+ devices be encrypted when using Android Enterprise. This is important for both general device security as well as application specific data security.
BYOD Enrollment
Enrolling a BYOD device into Systems Manager is a simple 2-step process:
1. Install the Systems Manager app - This can be done two ways. Using a Google Managed domain, simply add a Google account in the bound domain and it will prompt the user to install the SM app. Alternatively, a Meraki Managed domain can download the SM app from the Google Play Store. The app can be found here: https://play.google.com/store/apps/d...=com.meraki.sm or downloaded directly here if the Google Play Store is unavailable. Regardless of the domain type, once the app is installed, follow the steps provided on the device to complete enrollment.
2. Sign in / Authenticate - When the app is opened two options will appear: Google and Meraki. These refer to the domain types that were bound to Dashboard.
- If Google is selected, it will prompt the user to login with their Google domain credentials, or select an account that has already signed into the device. The app will then automatically enroll in the correct Dashboard network.
- If Meraki is chosen, it will prompt to enter an enrollment code (this can be found in Dashboard under Systems Manager > Manage > Add Devices > Android Tab) and subsequently ask for a username and password.

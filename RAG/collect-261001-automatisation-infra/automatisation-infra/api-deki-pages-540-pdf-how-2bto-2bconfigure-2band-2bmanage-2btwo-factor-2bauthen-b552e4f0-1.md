---
id: collect-261001-automatisation-infra/automatisation-infra/api-deki-pages-540-pdf-how-2bto-2bconfigure-2band-2bmanage-2btwo-factor-2bauthen-b552e4f0-1
title: "api-deki-pages-540-pdf-how-2bto-2bconfigure-2band-2bmanage-2btwo-factor-2bauthen-b552e4f0"
domain: automatisation-infra
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-automatisation-infra/api-deki-pages-540-pdf-how-2bto-2bconfigure-2band-2bmanage-2btwo-factor-2bauthen-b552e4f0.md
source_anchor: ""
source_lines: [1, 81]
sha256: 4ffdc079f8e5d7e4afa766fc29ff386ae56a0f23d9ec63fac4561ef9f5693a9b
---

# api-deki-pages-540-pdf-how-2bto-2bconfigure-2band-2bmanage-2btwo-factor-2bauthen-b552e4f0

How to Configure and Manage Two-Factor Authentication
Click 日本語 for Japanese
Overview
This article explains how to set up, manage, and recover Two-Factor Authentication (TFA) for Cisco Meraki dashboard access and Client VPN authentication.
TFA is also called 2FA, two-step verification, or multi-factor authentication (MFA). It adds an extra layer of security to user verification. It uses a security identifier
in addition to a username and password.
The identifier stays separate from the original login method. Examples include phone apps or keyfobs.
When TFA is not enabled, the user sees a banner. It states: "Two-factor Authentication is not currently enabled on your Meraki dashboard account. For an extra
layer of security, we recommend enabling it at your earliest convenience." Selecting the [x] button does not permanently remove the banner.
Duo Mobile on an iOS or Android device provides two-factor authentication regardless of SMS service. It also lets users back up Duo-protected accounts.
Recovery is possible to the same device or to a new device.
Prerequisites
• A valid Meraki dashboard username and password.
• A smartphone (iOS or Android) for the Duo Mobile app, when using app-based 2FA.
• A third-party two-factor authentication solution for Client VPN. Client VPN does not natively support two-factor authentication.
• Organization admin access with Full permissions, for organization-wide enforcement.
Step-by-step instructions
Set up Duo Mobile in Dashboard
1. Open your smartphone's mobile app store. Download the Duo Mobile app.
Dashboard organizations should always have at least two organization admins with full permissions. This is best practice if one account is locked out.
It also helps if access to that account's email address is lost.
1

2. Log in to dashboard. Navigate to the My Profile page on the top right.
3. Scroll to the section labeled Two-factor authentication.
4. Select Set up two-factor authentication.
5. On the next page, under Set up app, follow the listed steps. This adds your dashboard account to Duo Mobile as a token.
6. Enter the current, active token into the Code field under Verify your device on dashboard. The token changes every 30 seconds.
On the Phone On Dashboard
2

7. After verification, select Continue. Then select OK to turn on two-factor authentication.
Manage one-time backup codes
After you set up two-factor authentication, note the eight backup codes provided. These codes serve only as a last-resort recovery method. Use them if your
primary method is temporarily unavailable, for example, due to a lost or damaged device.
To find your backup codes:
1. Log in to the dashboard with a valid username and password.
2. Locate the My Profile option on the top right corner of the screen.
3. Select My Profile.
4. Scroll down to the Two-factor authentication section. Find your One-time codes (a numbered list, one through eight).
You can use each code only once. Select Generate a new set of one-time codes at any time to refresh the eight codes. Any unused codes from
the previous set become void at that time.
Avoid using backup codes for regular sign-ins. Keep your primary method active and accessible. Store the codes securely for emergencies. Rely on your main
method whenever possible.
The dashboard account logs out once you select OK. At the next login, dashboard prompts you for the active verification code from the
authenticator. Enable Duo Restore (iOS, Android). This allows easy account recovery to the same device or a new device. To change the
authenticator app after setup, first disable two-factor authentication. Then start the configuration again.
3

Use two-factor authentication with Client VPN
Meraki Client VPN includes several methods for authenticating users. Authentication happens before users join the network. For an additional level of security,
Client VPN also supports third-party two-factor authentication solutions. These require users to complete a second authorization step.
Client VPN does not natively support two-factor authentication. A third-party solution is required. Consult your two-factor authentication solution's documentation
for additional information and troubleshooting.
You can incorporate two-factor authentication in one of two ways:
1. As part of the authentication. Users enter a username and password as normal. They must also provide additional information required by the third-
party solution. For example, they append a key to the password.
2. As a push notification. An agent on a RADIUS server holds an accept message. The user selects an 'accept' button or equivalent to proceed.
By default on the Meraki platform, the RADIUS session times out after a short period. This may be too short for some solutions. Contact Meraki Support
to extend this timeframe.
Both methods are compliant under the PCI DSS 3.0 standard as two-factor security for remote access.
Disable two-factor authentication
1. Log in to the dashboard with a valid username and password.
2. Locate the My Profile option on the top right corner of the screen.
3. Select My Profile.
4. Scroll down to the Two-factor authentication section.
5. Select Turn off two-factor authentication.
6. Verify your password to finalize the process.
7. Optionally, select Remove next to your previously established phone number(s). This stops it from being saved for future configuration.
Verification
• Log out, then log back in. Dashboard prompts you for the active verification code from your authenticator. The correct code confirms TFA is active.
• For backup codes, confirm the Two-factor authentication section displays your eight one-time codes (numbered one through eight).
• For Client VPN, a successful second authorization step confirms two-factor authentication works. This is a key appended to the password or an accepted
Client VPN does not support xauth. Two-factor authentication solutions that use xauth are not supported.
The organization-wide security configuration Force users to set up and use two-factor authentication overrides an individual's ability to disable
TFA. Disabling this configuration does not disable TFA for any users. Re-enabling it forces everyone in the organization to follow the policy. This
includes new administrators and existing administrators who temporarily disabled their individual TFA.
You may lose access to the original or current phone number. If so, Meraki Support may assist. First follow the account recovery steps in the
Troubleshooting section.
4


---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-27130425853079-configuring-touch-pass-in-unifi-access-44e73df1-2
title: "hc-en-us-articles-27130425853079-configuring-touch-pass-in-unifi-access-44e73df1"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "China", "Google", "United States"]
dates: []
keywords: ["pricing"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-27130425853079-configuring-touch-pass-in-unifi-access-44e73df1.md
source_anchor: ""
source_lines: [13, 117]
sha256: a456b865e14868b8e2ffc925353b44afa4b4a6acc2b0b48f78d0338fef7340ba
---

# hc-en-us-articles-27130425853079-configuring-touch-pass-in-unifi-access-44e73df1

  - G6 Pro Entry (UVC-G6-Pro-Entry) 5.2.84 or later*
  - G6 Entry (UVC-G6-Entry) 5.1.210 or later*
  - G3 Intercom (UA-G3-Intercom) 1.7.29 or later
  - G3 Reader Pro (UA-G3-Pro) 1.10.30 or later
  - G3 Reader (UA-G3) 3.14.3.0 or later
  - Reader Flex (UA-G3-Flex) 1.3.7 or later
* To enable door unlock functions, G6 Entry devices must be adopted in the Protect application and directly connected to Access Control Hubs.
Remote Access enabled
Navigate to your local console (https://192.16.1.1) and go to Settings > Control Plane > Console > Remote Access. Only the Console Owner can perform this action.
Smart Door Access enabled
User device support
For iOS
- iPhone or Apple Watch compatible with Apple Pay
- Running iOS version 16.0 or later (update iPhone | Apple Watch)
- Set up Face ID, Touch ID, or passcode on their iPhone
- UniFi Endpoint mobile app
  - You have invited users to UniFi Endpoint
  - UniFi Endpoint iOS 3.5.0 or later
- Touch Pass is not available when using a federated Apple account.
For Android
- Android phone compatible with Google Wallet (currently not available on Wear OS)
- Running Android version 9 or later (update Android phone)
- UniFi Endpoint mobile app
  - You have invited users to UniFi Endpoint
  - UniFi Endpoint Android 3.5.1 or later
Apple Pay/Google Wallet is regionally supported
- For iOS: Ensure Apple Pay is available in your region.
- For Android: Ensure Google Wallet is available in your region.
Touch Pass is not available for kids under 13 due to Apple Pay and Google Wallet policies.
Note: Touch Pass is not supported in China, Kosovo, Vietnam, Nepal and any country or region subject to the sanctions of the United States government and other relevant governments.
Assigning Touch Passes
UniFi's one-year free trial includes a one-time grant of 10 complimentary Touch Passes for each console that supports UniFi Access. This offer becomes available the first time a compatible reader (e.g., G3 Intercom, G3 Reader Pro, or G3 Reader) is adopted in the UniFi Access application. For information on purchasing additional Touch Passes, see Purchasing New Touch Passes.
Touch Pass can be assigned using one of the two methods detailed below.
Via the Touch Pass Page
- Navigate to Access application > Settings > Touch Pass, then do one of the following:
  - Select a pass and Click to Assign it to a user.
  - Click Multi-Assign for bulk assignment.
- Select users and click Assign. Users who have already been assigned a pass will not appear in the list.
- Once a pass is assigned, an invitation email will be sent automatically to the user if an email address is provided.
  - If no email address is available, click Invite User to Activate Touch Pass to add an email or copy the link to send it to the user.
  - If Require a Verification Code When Loading a Credential is enabled in Settings > Identity > Identity Credentials, please send the invitation from the People page to ensure the verification code is included in the email.
Via the People Page
- Navigate to Access application > People > select a user > Settings > Credentials > Touch Pass > +.
  - If there are available passes, one will be automatically assigned to the user.
  - If no passes are available but auto-scaling is enabled, a pass will be automatically assigned to the user. The associated fee will be charged to the credit card on file and billed to the Console Owner.
  - If no passes are available and auto-scaling is disabled, you will be prompted to either purchase more passes or contact the Console Owner.
- Click Apply Changes.
- Once a pass is assigned, go to Overview > Send Invitation (or Invite Again) to send an invitation email or link to the user.
Unlocking Doors with Touch Pass
Touch Pass allows users to unlock doors, simply by tapping their phone on the UniFi reader. By default, users can unlock doors even if their phone is locked. To enhance security, you can disable this feature by navigating to Access application > Settings > Touch Pass > Express Mode.
For more information, see our detailed Touch Pass User's Guide.
Purchasing New Touch Passes
Only the Console Owner can purchase passes. Our 1-year free trial comes with 10 Touch Passes. To purchase more:
- Navigate to Access application > Settings > Touch Pass.
- Click the Purchase New Touch Pass button.
- Select the pass count and payment method and click Purchase.
Enabling Auto-Scaling for Fast Purchases
To streamline the process, enable auto-scaling, which automatically charges you for Touch Passes as they are assigned to users. Only the Console Owner can manage this feature.
- Navigate to Access application > Settings > Touch Pass.
- Click the Auto-Renewal & Auto-Scaling button.
- Tick the Auto-Scaling checkbox.
- Enter the Service Owner name, which is usually the credit card owner in account.ui.com.
- The pricing per user/year will be displayed.
- Select the Payment Method. If no method has been set up, add it to account.ui.com first.
- Agree to the Terms and Conditions and the Privacy Policy checkbox.
- Click Set Up.
Enabling Auto-Renewal for Touch Passes
Enabling auto-renewal ensures the Touch Pass payment is auto-renewed yearly. If not renewed, an expired pass cannot be used to unlock doors. Only the Console Owner can manage this feature.
- Navigate to Access application > Settings > Touch Pass.
- Click the Auto-Renewal & Auto-Scaling button.
- Tick the Auto-Renewal checkbox.
- Click Set Up.
Managing Touch Passes
You can view the status and manage all Touch Passes in Access application > Settings > Touch Pass. In addition to monitoring usage, you can:
- 
Unbind: Disconnect the Touch Pass from the user's current device.
  - The user can no longer use that device to unlock doors with the Touch Pass.
  - The Touch Pass remains owned by the same user and linked to their iCloud/Google account.
  - The user can activate the Touch Pass on another device using the same account.
  - Important: Unbinding does not make the Touch Pass available for another user. To assign the Touch Pass to a different user, please contact Technical Support.
- Suspend: Temporarily disable a Touch Pass so it can no longer be used to unlock doors. Suspended Touch Passes will not be billed in the next billing cycle, even if auto-renewal is enabled.
- Resume: Reactivate a suspended pass.
Restoring Touch Passes
Restore previously assigned Touch Passes when replacing consoles.
- Navigate to Access application > Settings > Touch Pass.
- Click the upper-right Restore Touch Pass.
- Choose a console to restore the passes from and confirm.
- Review and reassign passes as needed.
Note: Restored passes will be removed from the original console and will need to be reassigned after.
FAQs
Can I use a single Touch Pass for my smartphone and smartwatch?
- For iOS: Yes, a single Touch Pass can be used on both an iPhone and an Apple Watch, provided both devices are signed into the same iCloud account.
- For Android: No, Touch Pass is currently only available on Android phones.
How many iCloud or Google accounts can be assigned a Touch Pass?
Can a single smartphone have multiple Touch Passes?
Can a Touch Pass that has been activated by a user be unassigned?
Will a suspended Touch Pass be billed in the next auto-renewal cycle?
When does the one-year free trial for a Touch Pass start, and does it expire if unused?
What happens to Touch Pass when upgrading to or downgrading from UniFi Fabric?
When upgrading to UniFi Fabric for a multi-site setup
- All Touch Passes on your consoles are migrated to UniFi Fabric.
- After migration, users need to re-add their Touch Pass in the UniFi Endpoint mobile app before they can use it again.
When dissolving UniFi Fabric and downgrading to a single-site setup
- Users do not need to re-add their Touch Pass; they can use it directly after the downgrade.
- If you have only one console, all Touch Passes will be automatically restored to that console.

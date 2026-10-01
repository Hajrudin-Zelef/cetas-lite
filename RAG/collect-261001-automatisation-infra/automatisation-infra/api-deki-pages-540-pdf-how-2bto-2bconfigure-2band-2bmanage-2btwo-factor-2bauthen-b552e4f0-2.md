---
id: collect-261001-automatisation-infra/automatisation-infra/api-deki-pages-540-pdf-how-2bto-2bconfigure-2band-2bmanage-2btwo-factor-2bauthen-b552e4f0-2
title: "api-deki-pages-540-pdf-how-2bto-2bconfigure-2band-2bmanage-2btwo-factor-2bauthen-b552e4f0"
domain: automatisation-infra
role: reference
task: reference
actors: ["Apple", "United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/api-deki-pages-540-pdf-how-2bto-2bconfigure-2band-2bmanage-2btwo-factor-2bauthen-b552e4f0.md
source_anchor: ""
source_lines: [82, 120]
sha256: 059b89ea4a3788d60c8817e297563cf094fd64b8659b5b3489f60e38d3ff25c2
---

# api-deki-pages-540-pdf-how-2bto-2bconfigure-2band-2bmanage-2btwo-factor-2bauthen-b552e4f0

push notification.
Troubleshooting
Recover access to a TFA-protected account
Meraki provides one-time codes to use in place of a TFA code. These are available for MFA and Duo Mobile. Treat one-time codes and Duo Restore as the
primary recovery steps. They temporarily bypass two-factor authentication to regain access.
Enable Duo Restore (iOS, Android). This allows easy account recovery to the same device or a new device.
If those solutions are not possible, Meraki Support can disable the account's TFA configuration. TFA is an important security mechanism. Meraki Support will not
disable it without first positively identifying the account owner.
2FA removal requests cannot be resolved via Support phone lines. Please Create a case with support using the Meraki Support Home page ,you can use
the No dashboard Access? option to do that as in the screenshot below:
Verify account ownership for recovery
Method 1
Use this method if a second organization administrator with full access does exist.
1. Open a support case from the Cisco Meraki Support website. Using the email address of the account TFA is to be disabled on. Include the full name of
the organization the account resides in.
1. Respond to the automated case creation email to confirm the TFA reset request.
The organization-wide security configuration Force users to set up and use two-factor authentication overrides Meraki Support's ability to disable
TFA. To complete the process, disable this configuration from every organization the account is associated with. Disabling it does not disable TFA for
any users. It only lets Meraki Support manually disable TFA for the locked-out account.
5

2. A second organization administrator must comment on the case through dashboard to approve the disable. Email or phone approval is not acceptable.
Approval must come as a case comment.
An organization administrator with Full access may grant this. A SAML administrator with Full access may also grant it. Approval by network administrators or
read-only administrators is not accepted.
Method 2
Use this method if a second organization administrator with full access does not exist or is unavailable.
1. Open a case from the Cisco Meraki Support website. Use the email address of the account TFA is to be disabled on. Include the full name of the
organization the account resides in.
2. The Support Operations Specialist requests information about the organization's specifics to verify ownership.
3. After verification, a DocuSign email is sent to the organization administrator. Fill it out, sign it digitally, and update the case once completed.
Additional resources
• DuoSecurity RADIUS documentation
• RSA SecurID
• Meraki employee IT support
Cisco Meraki US Government Region Administrators must and comment with approval via a government case and reference the commercial case #
created in the previous step.
For more on handling 2FA unlock requests, refer to the documentation on Support Policies and Exceptions.
6

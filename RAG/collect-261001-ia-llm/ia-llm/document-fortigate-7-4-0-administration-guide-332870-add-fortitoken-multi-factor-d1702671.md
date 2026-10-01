---
id: collect-261001-ia-llm/ia-llm/document-fortigate-7-4-0-administration-guide-332870-add-fortitoken-multi-factor-d1702671
title: "document-fortigate-7-4-0-administration-guide-332870-add-fortitoken-multi-factor-d1702671"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-ia-llm/document-fortigate-7-4-0-administration-guide-332870-add-fortitoken-multi-factor-d1702671.md
source_anchor: ""
source_lines: [1, 25]
sha256: 2fcd3b2426dfff6c491d51a06e323452c20e234de07b5184d81f8bdd894200ee
---

# document-fortigate-7-4-0-administration-guide-332870-add-fortitoken-multi-factor-d1702671

Add FortiToken multi-factor authentication
Add FortiToken multi-factor authentication
This configuration adds multi-factor authentication (MFA) to the FortiClient dialup VPN configuration (FortiClient as dialup client). It uses one of the two free mobile FortiTokens that is already installed on the FortiGate.
To configure MFA using the GUI:
- Edit the user:
  - Go to User & Authentication > User Definition and edit local user vpnuser1.
  - Enable Two-factor Authentication.
  - For Authentication Type, click FortiToken and select one mobile Token from the list.
  - Enter the user's Email Address.
  - Enable Send Activation Code and select Email.
  - Click Next and click Submit.
- Activate the mobile token.
  - When a FortiToken is added to user vpnuser1, an email is sent to the user's email address. Follow the instructions to install your FortiToken mobile application on your device and activate your token.
To configure MFA using the CLI:
- Edit the user and user group:config user local
    edit "vpnuser1"
        set type password
        set two-factor fortitoken
        set fortitoken <select mobile token for the option list>
        set email-to <user's email address>
        set passwd <user's password>
    next
end
- Activate the mobile token.
  - When a FortiToken is added to user vpnuser1, an email is sent to the user's email address. Follow the instructions to install your FortiToken mobile application on your device and activate your token.

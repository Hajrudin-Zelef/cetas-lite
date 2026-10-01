---
id: collect-261001-ia-llm/ia-llm/document-fortigate-7-4-12-administration-guide-458581-set-up-fortitoken-multi-fa-a770fb5c
title: "document-fortigate-7-4-12-administration-guide-458581-set-up-fortitoken-multi-fa-a770fb5c"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-ia-llm/document-fortigate-7-4-12-administration-guide-458581-set-up-fortitoken-multi-fa-a770fb5c.md
source_anchor: ""
source_lines: [1, 28]
sha256: 97d1afe7cfdca1b5f4bd796ffa7303a82bfdd81ae1dbfc7da4e59e12b89ce0cc
---

# document-fortigate-7-4-12-administration-guide-458581-set-up-fortitoken-multi-fa-a770fb5c

Set up FortiToken multi-factor authentication
Set up FortiToken multi-factor authentication
This configuration adds multi-factor authentication (MFA) to the split tunnel configuration (SSL VPN split tunnel for remote user). It uses one of the two free mobile FortiTokens that is already installed on the FortiGate.
To configure MFA using the GUI:
- Configure a user and user group:
  - Go to User & Authentication > User Definition and edit local user sslvpnuser1.
  - Enable Two-factor Authentication.
  - For Authentication Type, click FortiToken and select one mobile Token from the list.
  - Enter the user's Email Address.
  - Enable Send Activation Code and select Email.
  - Click Next and click Submit.
- Activate the mobile token.When a FortiToken is added to user sslvpnuser1, an email is sent to the user's email address. Follow the instructions to install your FortiToken mobile application on your device and activate your token.
To configure MFA using the CLI:
- Configure a user and user group:config user local
    edit "sslvpnuser1"
        set type password
        set two-factor fortitoken
        set fortitoken <select mobile token for the option list>
        set email-to <user's email address>
        set passwd <user's password>
    next
end
config user group
    edit "sslvpngroup" 
        set member "sslvpnuser1"
    next 
end
- Activate the mobile token.When a FortiToken is added to user sslvpnuser1, an email is sent to the user's email address. Follow the instructions to install your FortiToken mobile application on your device and activate your token.

---
id: collect-261001-meraki/meraki/ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-api-key-leak-re-bcec91ca-1
title: "ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-api-key-leak-re-bcec91ca"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["incident", "training"]
source: docs/RAG/collect-261001-meraki/ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-api-key-leak-re-bcec91ca.md
source_anchor: ""
source_lines: [1, 151]
sha256: ff80ba1389f3cc5e4f03dc2cb098abeb2141795477db45451b206261835c1ea0
---

# ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-api-key-leak-re-bcec91ca

If a Meraki Dashboard API key has been exposed or compromised, immediate action is required to secure your organization. API keys grant access to your network configurations and data, so treating a leak with urgency is critical.
This guide provides step-by-step procedures for responding to an API key leak, investigating the incident, and preventing future exposures.
⚠️ Important Disclaimer
This guide is not legal advice and is not intended to override or replace your organization's internal security practices, policies, or procedures.
This guide represents general industry best practices and Meraki-specific recommendations for responding to API key leaks. However, you must first and foremost follow your own organization's security breach procedures, incident response plans, and legal requirements. Your organization may have specific policies, compliance obligations, or regulatory requirements that supplement or supersede the guidance provided here.
Before taking any action, consult with:
- Your organization's security team or Security Operations Center (SOC)
- Your legal or compliance department
- Your incident response team or designated security contacts
- Any applicable industry-specific regulatory requirements (HIPAA, PCI-DSS, GDPR, etc.)
This guide should be used as a reference and starting point only, not as a replacement for your organization's established procedures or as legal counsel.
First, determine who owns the compromised API key. Each API key is associated with a specific admin identity (email address).
If the key belongs to your own email address:
- Proceed directly to Step 2
If you recognize the key but do not own it:
- Record the email address and name of the admin who owns the key
- Contact the owner immediately if they are a current employee
- If the owner has left the organization or is unresponsive, proceed with the admin removal process (see Step 3)
If you cannot identify the key owner but believe the key has access to your dashboard organization:
- Contact your organization's security team
- Review the list of admins in Dashboard (Organization > Admins)
- Check API activity logs to correlate the key with recent API calls
For your own API key:
- Sign in to Meraki Dashboard: https://dashboard.meraki.com
- Navigate to Organization >API & Webhooks
- Select API keys and access from the top tabs
- Locate the compromised API key in the list
- Click Revoke next to the compromised key
- Confirm the revocation
⚠️ Critical: Once revoked, the API key will immediately stop working. Any applications or scripts using this key will fail. Have a replacement key ready before revoking if you need to maintain service continuity.
If the API key belongs to a different admin who has access to your organization:
- Sign in to Meraki Dashboard: https://dashboard.meraki.com
- Navigate to Organization >Admins
- Locate the admin account associated with the compromised key
- Review all organizations to which this admin has access
- Remove the admin from each organization:
  - Click on the admin's name
  - Select Delete or Remove from organization
  - Confirm the removal
- Document the admin's email address, name, and removal date for your records
Note: Removing an admin from an organization immediately revokes their access to that organization regardless of which API key(s) they have. If they have access to multiple organizations, remove them from each one.
Use Meraki's API analytics to check for unauthorized or unexpected API activity:
- 
Sign in to Meraki Dashboard: https://dashboard.meraki.com
- 
Navigate to Organization >API & Webhooks
- 
Select API usage orAPI logs from the tabs
- 
Review recent API calls for suspicious activity: 
  - Look for:
    - API calls from unfamiliar IP addresses
    - Calls during unusual hours (nights, weekends)
    - High-volume requests (potential data exfiltration)
    - Calls to sensitive endpoints (configuration changes, deletions)
    - Failed authentication attempts (401/403 errors)
  - Filter by:
    - The compromised admin's identity
    - Date range covering the exposure window
    - Specific API operations (GET, POST, PUT, DELETE)
- Look for:
- 
Document any suspicious activity: 
  - Timestamp of suspicious calls
  - Source IP addresses
  - API endpoints accessed
  - Operations performed (reads vs. writes)
- 
Review activity for other admins as well to ensure no lateral movement or additional compromises
If you need to restore API access after revoking a compromised key:
- Sign in to Meraki Dashboard using the admin account that will provide the new API key: https://dashboard.meraki.com
- Navigate to Organization >API & Webhooks
- Select API keys and access from the top tabs
- Click Generate new API key
- Copy the new API key immediately (it will only be displayed once)
- Store the new key securely in your environment variables or secrets manager
Update all applications, scripts, and services that use the API key:
- 
Identify all systems using the compromised key: 
  - Development, staging, and production environments
  - Automated scripts and cron jobs
  - CI/CD pipelines
  - Third-party integrations
  - Documentation and runbooks
- 
Deploy the new API key to each system
- 
Test each application to verify the new key works
- 
Monitor for errors or failed authentication attempts
After containing the immediate threat, investigate how the API key was leaked and what damage may have occurred.
Common sources of API key leaks include:
- 
Version Control Systems 
  - Check git commit history for hardcoded keys
  - Search GitHub, GitLab, Bitbucket for public repositories containing your key
  - Use tools like git-secrets or truffleHog to scan repositories
- 
Logs and Monitoring Systems 
  - Review application logs for accidentally logged keys
  - Check centralized logging systems (Splunk, ELK, CloudWatch)
  - Search CI/CD build logs
- 
Documentation and Wikis 
  - Search internal documentation for example code with real keys
  - Check shared note-taking apps (Confluence, Notion, OneNote)
  - Review training materials and runbooks
- 
Communication Channels 
  - Search Slack, Teams, or email for keys shared in messages
  - Check support tickets or help desk systems
  - Review screenshare recordings or recorded meetings
- 
Public Exposure 
  - Search paste sites (Pastebin, GitHub Gists)
  - Check Stack Overflow and developer forums
  - Review any public-facing applications or API proxies
This list is illustrative but not exhaustive. Consult your organization's incident response team for additional avenues of exposure.
If you have removed an admin's access or identified a leak from another team member's key:
- Contact the key owner to inform them of the incident
- Explain what happened and the actions you've taken
- Ask them to:
  - Confirm which systems were using the key
  - Help identify how the key was exposed
  - Review their security practices
- Provide guidance on preventing future leaks (see Prevention section below)
Determine what unauthorized actions, if any, were taken using the compromised key:
- 
Review API logs for the entire exposure window
- 
Identify unauthorized operations: 
  - Data access: Were any sensitive network configurations or device details accessed?
  - Configuration changes: Were any networks, devices, or settings modified?
  - Data exfiltration: Was there high-volume data retrieval?
  - Destructive actions: Were any resources deleted or disabled?
- 
Document findings: 
  - Timeline of unauthorized activity
  - Scope of data accessed or modified
  - Business impact assessment
  - Compliance implications (GDPR, HIPAA, PCI-DSS, etc.)
- 
Remediate unauthorized changes: 
  - Revert configuration changes if necessary
  - Review and validate network settings
  - Check for backdoors or persistence mechanisms
Depending on your organization's regulatory requirements, you may need to:
- 
Report the incident internally: 
  - Notify your security team
  - Inform management and stakeholders

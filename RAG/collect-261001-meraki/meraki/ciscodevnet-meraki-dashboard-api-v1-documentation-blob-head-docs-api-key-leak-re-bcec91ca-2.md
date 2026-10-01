---
id: collect-261001-meraki/meraki/ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-api-key-leak-re-bcec91ca-2
title: "ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-api-key-leak-re-bcec91ca"
domain: meraki
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["aws", "cyber", "incident", "training"]
source: docs/RAG/collect-261001-meraki/ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-api-key-leak-re-bcec91ca.md
source_anchor: ""
source_lines: [152, 324]
sha256: 36385ff18fc68aa945258ec8ee3cd3bc7905b6181eee99317906f4bada41ea5a
---

# ciscodevnet-meraki-dashboard-api-v1-documentation-blob-head-docs-api-key-leak-re-bcec91ca

  - Document the incident for security records
- 
Report the incident externally (if required): 
  - Notify affected customers if their data was accessed
  - File breach reports with regulatory bodies (GDPR requires notification within 72 hours)
  - Inform cyber insurance provider
- 
Preserve evidence: 
  - Save API logs showing unauthorized access
  - Document the timeline of the incident
  - Retain copies of communications related to the incident
Implement these practices to prevent future API key leaks:
Never hardcode API keys in source code. Always use one of these secure storage methods:
- 
Environment Variables (recommended for local development) 
  - Store keys in MERAKI_DASHBOARD_API_KEY environment variable
  - Add .env files to.gitignore
  - Use tools like python-dotenv ordirenv to manage environment variables
- Store keys in 
- 
Secrets Management Systems (recommended for production) 
  - HashiCorp Vault
  - AWS Secrets Manager
  - Azure Key Vault
  - Google Cloud Secret Manager
  - 1Password Secrets Automation
  - Doppler
- 
CI/CD Secret Stores 
  - GitHub Secrets
  - GitLab CI/CD Variables
  - Jenkins Credentials
  - CircleCI Environment Variables
- 
Add API keys to .gitignore :.env
.env.local
config/secrets.yml
**/api_keys.txt
- 
Use pre-commit hooks: 
  - Install git-secrets
  - Configure pre-commit with secret detection
- 
Scan repositories regularly: 
  - Use truffleHog to scan for secrets
  - Enable GitHub's secret scanning feature
  - Use GitGuardian or similar tools for continuous monitoring
- 
If you accidentally commit a key: 
  - Revoke the key immediately (don't wait to clean git history first)
  - Remove the key from git history using git filter-branch or BFG Repo-Cleaner
  - Force push the cleaned history (coordinate with team)
  - Remember: public commits may be cached by search engines or GitHub's API
Establish a routine API key rotation schedule. These are only suggestions; consult your organization's security team for any specific requirements that may override or replace these suggestions.
- 
Recommended rotation frequency: 
  - Every 30-90 days is a common baseline.
  - After any personnel changes, if the admin ever had access to the API key, and the API key didn't belong to their admin user, then rotate the key immediately. If the admin never had access to the key, or if the key belonged to that admin, then deleting the admin from an org prevents the key from working in that org.
- 
Use Meraki's two-key system for zero-downtime rotation: 
  - Week 1: Generate second API key
  - Week 2: Deploy new key to all development/staging systems
  - Week 3: Deploy new key to production systems
  - Week 4: Verify new key functionality and revoke old key
- 
Automate rotation where possible: 
  - Use secrets managers with automatic rotation capabilities
  - Script the key generation and deployment process
  - Set calendar reminders for manual rotation
- 
Principle of Least Privilege: 
  - Create admin accounts with only the permissions required
  - Use read-only admin roles when write access isn't needed
  - Consider using OAuth 2.0 instead of API keys for applications requiring specific, limited permissions
- 
Multi-Factor Authentication: 
  - Enable 2FA for all admin accounts that can generate API keys
  - Use authenticator apps instead of SMS whenever possible
- 
Admin Lifecycle Management: 
  - Document all API keys and their purposes
  - Revoke API keys immediately when admins leave the organization
  - Review admin access quarterly
  - Maintain an inventory of active API keys
Set up monitoring to detect potential API key compromises:
- 
Establish baselines: 
  - Normal API call volume per hour/day
  - Typical API operations performed
  - Expected source IP addresses
  - Standard usage patterns (time of day, day of week)
- 
Configure alerts for anomalies: 
  - API calls from new geographic locations
  - High-volume requests (potential data exfiltration)
  - Failed authentication attempts (401/403 errors)
  - Calls to sensitive endpoints (DELETE operations, admin changes)
  - API usage during off-hours
- 
Regular log reviews: 
  - Review API logs weekly or monthly
  - Look for unusual patterns or unexpected activity
  - Investigate any rate limit breaches
- 
Rate limiting awareness: 
  - Meraki enforces 5 calls per second per organization
  - Unusual rate limiting errors may indicate unauthorized usage
  - Monitor for 429 (Too Many Requests) responses
Consider whether API keys are the best authentication method for your use case:
| Feature | OAuth 2.0 Grants | API Keys | 
|---|---|---|
| Best for | Third-party applications, organization-wide automation | Personal scripts, admin-specific tasks | 
| Scope | App-scoped with granular permissions | Admin-scoped based on user role | 
| Permissions | Configurable per application | Inherits from admin's role | 
| Identity | Application identity | Admin identity | 
| Management | Organization level | Individual admin level | 
| Token lifetime | 60 minutes (auto-refresh) | Permanent until revoked | 
| Security | Better isolation, shorter-lived tokens | Simpler but higher risk if exposed | 
Recommendation: For third-party applications or organization-wide automation, prefer OAuth 2.0 grants over API keys. OAuth tokens are short-lived (60 minutes) and have configurable, granular permissions. See the OAuth documentation for more information.
- 
Security training: 
  - Train developers on API key security best practices
  - Include security in onboarding for new team members
  - Conduct regular security awareness sessions
- 
Incident response drills: 
  - Practice API key leak response procedures
  - Test your team's ability to detect and respond quickly
  - Update procedures based on lessons learned
- 
Documentation: 
  - Maintain runbooks for key rotation and incident response
  - Document where API keys are used in your infrastructure
  - Keep contact information current for security incidents
Use this checklist when responding to an API key leak:
- Identify the owner of the compromised API key
- Revoke the compromised API key via Dashboard
- Remove admin access if owner has left the organization
- Generate a replacement API key (if needed)
- Notify your security team
- Audit API activity logs for suspicious behavior
- Identify source of the leak (git history, logs, etc.)
- Determine exposure window (when was key created vs. exposed)
- Assess damage and unauthorized actions taken
- Contact the key owner to discuss the incident
- Document findings and timeline
- Update all applications and services with new API key
- Revert any unauthorized configuration changes
- Test all systems to verify new key works
- Monitor API logs for continued suspicious activity
- File compliance reports if required
- Implement preventive measures (pre-commit hooks, secrets scanning)
- Update security procedures and documentation
- Conduct team training on API key security
- Set up monitoring and alerting for future incidents
- Schedule regular API key rotation
- Review and update access controls
- Authorization Documentation - Learn about API keys vs OAuth 2.0
- OAuth Documentation - Implementing OAuth for better security
- Getting Started Guide - Best practices for new API users
- Rate Limiting - Understanding API rate limits
- Meraki Developer Portal - Official Meraki API documentation
- Meraki Community - Get help from the Meraki community
If you need assistance responding to an API key leak:
- 
For urgent security incidents: 
  - Contact your organization's security team immediately and follow their instructions.
  - If you're a Meraki customer with support, open an urgent support case and provide information about which steps you've already taken.
- 
For general questions: 
  - Visit the Meraki Community
  - Review the official API documentation
  - Consult your organization's security policies and procedures

---
id: collect-240926-datacamp/datacamp/github-copilot-enterprise-spaces-et-api-des-metriques-dusage-1
title: "github-copilot-enterprise-spaces-et-api-des-metriques-dusage"
domain: datacamp
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2025-11-01"]
keywords: ["copilot", "cost", "governance", "license", "licenses"]
source: docs/RAG/clean_en/datacamp/github-copilot-enterprise-spaces-et-api-des-metriques-dusage.md
source_anchor: ""
source_lines: [1, 163]
sha256: d3dcb6625a24951e298fd6ef4747067b945e6bb195eacc6b071cdf12fa60b839
---

# github-copilot-enterprise-spaces-et-api-des-metriques-dusage

<!-- source: https://www.datacamp.com/fr/blog/github-copilot-enterprise -->

Course

You've deployed GitHub Copilot Enterprise across the organization, assigned licenses, configured policies, and your developers are already using it in their IDEs. Now you need to answer the hard questions:

- How do you optimize Copilot so it better learns your company's specific engineering context?
- How do you measure Copilot's value? Which departments are adopting it successfully and which are ignoring it entirely?

This is where GitHub Copilot Spaces and the usage metrics API come in. Spaces allows Copilot to ingest your organization's technical knowledge. The usage metrics API helps administrators measure adoption, retention, and productivity trends across the enterprise.

In this article, we'll cover:

- What GitHub Copilot Enterprise includes
- How Copilot Spaces work
- How to configure Spaces at scale
- The GitHub Copilot usage metrics API endpoints
- Authentication and reporting workflows
- Concrete strategies for measuring ROI

If you're not comfortable with GitHub organizations, pull requests, and permission models, the Intermediate GitHub Concepts course covers these fundamentals. If you're also new to Copilot, our How to Use GitHub Copilot tutorial introduces the basic features that this guide builds on.

## Strengthen your data privacy and governance

Ensure compliance and protect your business with DataCamp for Business. Specialized courses and centralized tracking to protect your data.

## What is GitHub Copilot Enterprise?

GitHub Copilot Enterprise sits at the top of GitHub's Copilot offerings.

Compared to GitHub Copilot Business or Pro+, Enterprise emphasizes governance, organizational context, and measurement capabilities. It's designed for enterprises managing large engineering environments, rather than individual developers or small teams.

Two capabilities matter most in practice:

1. Custom organizational context via **Spaces**
2. Organization-wide telemetry via the **usage metrics API**

These two features transform Copilot from simple "smart autocomplete" into something resembling a true internal AI-powered engineering platform.

Enterprises that get the most value from GitHub Copilot Enterprise integrate it as a key element of their internal infrastructure. They curate organizational context, continuously measure adoption, and adjust policies based on usage data, not assumptions.

For a broader overview of the GitHub ecosystem, we recommend our Introduction to GitHub Products guide.

### How Enterprise differs from Business and Pro+

GitHub Copilot Enterprise extends the Business offering with:

- Organization-level usage metrics
- Enhanced governance controls
- Enterprise-wide policy inheritance
- Higher quotas for premium requests (1,000 vs. 300 on Business)
- Additional model access and management

Enterprise requires GitHub Enterprise Cloud in addition to the Copilot Enterprise subscription. This adds an additional per-user cost, so make sure your organization actually needs enterprise-level governance, telemetry, and administration.

| **Feature** | **Pro+** | **Business** | **Enterprise** | 
| Individual use | Yes | No | No | 
| Centralized license management | No | Yes | Yes | 
| Audit logs | No | Yes | Yes | 
| File exclusions | No | Yes | Yes | 
| Spaces support | Yes, with Copilot | Yes, limited admin visibility | Yes, full enterprise-level management | 
| Usage metrics API | No | Organization level | Enterprise + organization level | 
| Enterprise policy inheritance | No | No | Yes | 

**Note:** Business subscribers access the usage metrics API at the organization level (`/orgs/{org}/…`). Enterprise subscribers additionally have access to aggregated reports at the enterprise level (`/enterprises/{enterprise}/…`) covering all organizations in a single view.

### Who GitHub Copilot Enterprise is for

GitHub Copilot Enterprise is aimed at organizations with mature GitHub environments.

Typical Enterprise customers:

- Large engineering organizations
- Regulated industries
- Multi-team platform teams
- Enterprises with internal development standards
- Organizations requiring centralized governance

Note that this doesn't inherently improve Copilot's performance. This distinction is important: many teams initially oversize their purchase thinking that Enterprise = "better Copilot," when Enterprise mainly adds governance and measurement tools.

## Copilot Spaces: custom context for your organization

Copilot Spaces solves one of the main limitations of general-purpose code assistants.

By default, Copilot has a good grasp of public programming knowledge. It doesn't automatically understand your internal APIs, architectural decisions, code conventions, deployment workflows, or onboarding documentation.

Spaces provides curated organizational context that Copilot can leverage in conversation and for coding assistance.

In practice, Spaces helps Copilot answer questions such as:

- "How do we structure our API handlers internally?"
- "Which authentication library does our platform team recommend?"
- "Which deployment workflow should this microservice use?"
- "What naming conventions does our backend team follow?"

### What Spaces support

Spaces covers a broader range of organizational content than the old Knowledge Bases system.

Supported content types:

- Code files
- Markdown documentation
- JSON files
- Uploaded files
- Images
- GitHub Issues
- Pull requests

Each content type provides a different kind of value.

Code files help Copilot understand implementation patterns. Markdown files detail architecture and onboarding. Pull requests expose review discussions and past engineering decisions. Together, they improve understanding of your organization's development practices.

A subtle but important point: Spaces is not just a vector store backed by GitHub. It includes sharing controls and governance workflows designed for the enterprise.

### The end of Knowledge Bases

GitHub retired the old Copilot Knowledge Bases feature on November 1, 2025.

Spaces replaces Knowledge Bases with:

- Broader content support
- Better sharing controls
- Improved administration
- More flexible organization-level management

You will still find outdated documentation and blog posts mentioning Knowledge Bases. Be careful with old tutorials: many endpoints and workflows changed between 2025 and 2026.

## Creating and configuring Copilot Spaces

On the administration side, creating a Copilot Space is fairly straightforward. The challenge is managing dozens, or even hundreds, of them across teams.

The initial structure tends to persist. I've seen organizations inadvertently create a "documentation jungle" in Spaces due to a lack of clear ownership rules from the start.

Anyone can create a Copilot Space, so let's test it in a personal repository. The steps are similar at the Enterprise level, with a few different pages.

### Configuring a Space

Creation generally follows this workflow:

1. Go to the Copilot Spaces page in the Enterprise admin area
2. Create a new Space

1. Select repositories and content sources, including MCPs and other useful tools

1. Add sources via the "+ Add sources" button on the right

1. Choose to share the Space or set sharing settings at this stage

1. Verify that Copilot can reference the content during conversations

Note for Enterprise users: your administrator may disable sharing of personal Spaces. If you are using your own account, this may limit sharing of a Copilot Space that does not use the company's repositories.

After configuration, administrators should test the Space with concrete prompts.

For example:

`How does our authentication middleware handle token refresh logic?`
Or:

`Show me an example of how our backend services structure database migrations.`
If Copilot cannot answer accurately, the cause is usually:


---
id: vague2-datacamp/datacamp/what-is-github
title: "Qu'est-ce que GitHub ? Le guide ultime"
domain: datacamp
role: reference
task: article
actors: ["AWS"]
dates: ["2026-09-23"]
keywords: ["agent", "aws", "license", "open source"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/what-is-github.md
source_anchor: ""
source_lines: [1, 69]
sha256: 0c4d50a8997d3309648aadbfa622b796299f126118457f05a826aede7f6959ed
---

# Qu'est-ce que GitHub ? Le guide ultime

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/what-is-github
- **Site** : DataCamp
- **Type** : Guide / Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp guide by Samuel Shaibu explains what GitHub is and how to use it for version control and team collaboration, aimed especially at data professionals. It opens with a relatable scenario: a bug appears, you want to revert to the last working version but can't recall all changes, or merging contributions is a headache. GitHub solves these problems.

Version control is a system that records the history of file changes over time, enabling multiple people to work together while maintaining an accurate record of changes. Without it, tracking code changes becomes chaotic, especially in teams.

GitHub's uses extend beyond version control: creating a project portfolio (public profile to showcase skills to recruiters), collaborating (teamwork, code sharing, peer review), and contributing to open source (exploring and contributing to data science projects).

How GitHub works: key components include repositories (folders storing project files and version history, with a unique URL), forks (personal copies of another user's repo), pull requests (formal proposals for changes for review and merge), issues (tracking tasks, bugs, improvements), branches (parallel versions of a repo), and merging (combining changes into the original project).

Git vs GitHub: Git is a distributed version control system (DVCS) that tracks changes and enables branches, with a staging area and commit history. GitHub is a web platform built on Git adding access control, bug tracking, task management, wikis, and integrations. A comparison table outlines definition, purpose, features, and benefits.

The practical guide walks through: creating an account (sign up, choose a plan — the free tier suffices for beginners), creating a repository (New repository, name/description, public/private, optional README/.gitignore/license), creating branches (Branch:main → New branch), making commits (edit file, Commit Changes with a descriptive message), creating a pull request (compare branches, create PR, add reviewers), and merging branches (Merge pull request → Confirm merge). It notes most actions can be done via Git CLI (`git pull`).

Alternatives to GitHub are profiled: GitLab (DevOps platform with CI/CD, issue tracking, Auto DevOps, container registry, security), Bitbucket (Atlassian, Git/Mercurial, Jira integration, branch permissions), SourceForge (early open-source hosting, SVN/Git), AWS CodeCommit (managed AWS service, IAM, CI/CD integration, encryption), and Cursor Origin (AI-agent-first Git forge by Anysphere, GitHub mirroring, Cursor integration). The article closes with an FAQ.

## Key points

- Version control records file change history, enabling collaboration and accurate tracking.
- GitHub uses: project portfolio, collaboration, open-source contribution.
- Core components: repositories, forks, pull requests, issues, branches, merging.
- Git is a DVCS; GitHub is a web platform built on Git adding collaboration/project-management features.
- Creating a GitHub account is free; the free tier suffices for beginners.
- Repositories can be public (visible to all) or private (you and invited collaborators).
- Workflow: create repo → create branch → commit changes → open pull request → merge.
- Most GitHub actions can also be performed via the Git command line.
- Alternatives: GitLab, Bitbucket, SourceForge, AWS CodeCommit, Cursor Origin.
- Avoid storing sensitive data (passwords, API keys) in public repositories.

## Technical data / figures

| Category | Git | GitHub |
|----------|-----|--------|
| Definition | Distributed version control system | Web platform built on Git |
| Purpose | Manage code, track changes, create branches | Host Git repos, provide collaboration tools |
| Features | Staging area, commit history, branches, merge | Access control, bug tracking, task management, wikis, integrations |
| Benefit | Collaborative work, detailed change tracking | Improved collaboration, project management, code review |

| Alternative | Key features |
|-------------|--------------|
| GitLab | Integrated CI/CD, issue tracking, Auto DevOps, container registry, security |
| Bitbucket | Atlassian suite, Git/Mercurial, Jira integration, branch permissions, Pipelines |
| SourceForge | Open-source hosting, download stats, forums, SVN and Git support |
| AWS CodeCommit | Managed AWS service, IAM, CodePipeline/CodeBuild/CodeDeploy, encryption |
| Cursor Origin | AI-agent-first Git forge, GitHub mirroring, Cursor integration, PR review |

| Workflow step | Action |
|---------------|--------|
| Account | Sign up, choose plan (free tier for beginners) |
| Repository | + → New repository → name, description, public/private, README/.gitignore/license |
| Branch | Branch:main → New branch → name → Create |
| Commit | Edit file → Commit Changes → descriptive message |
| Pull request | Pull requests → New pull request → compare → Create → add reviewers |
| Merge | Open approved PR → Merge pull request → Confirm merge |

## Why this source matters for the RAG

This guide provides a complete, practical introduction to GitHub covering concepts, the Git vs GitHub distinction, hands-on workflows, and platform alternatives, making it a strong retrieval source for beginner and collaboration queries. Its step-by-step instructions and comparison tables offer directly reusable content for the knowledge base.

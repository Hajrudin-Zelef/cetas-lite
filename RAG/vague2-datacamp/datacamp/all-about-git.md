---
id: vague2-datacamp/datacamp/all-about-git
title: "Qu'est-ce que Git ? - le guide du débutant sur le contrôle de version Git"
domain: datacamp
role: reference
task: article
actors: ["Microsoft"]
dates: ["2026-09-23"]
keywords: ["open source"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/all-about-git.md
source_anchor: ""
source_lines: [1, 73]
sha256: 75d7598933d2f2cc356c28055724cd1950f9b8b0f8713ded896758de566b540a
---

# Qu'est-ce que Git ? - le guide du débutant sur le contrôle de version Git

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/all-about-git
- **Site** : DataCamp
- **Type** : Article / Guide
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp guide by Summer Worsley introduces Git, the world's most widely used version control system, used by over 90% of professional developers. Git is a distributed version control system (dVCS) that manages and tracks different versions of a project, recording changes to files so specific iterations can be retrieved later. VCSs are also called source code management (SCM) or revision control systems (RCS).

The guide explains the benefits of version control: attributable changes (each change linked to a team member), fine traceability and easy rollback, better organization and communication (commit messages), concurrency (parallel work with limited conflicts), and branching/merging (separate branches merged into the main branch). Git is not the only VCS — others include Fossil, Mercurial, and Subversion — but it is the most popular. The key distinction is centralized vs distributed systems: centralized systems have a central server, while distributed systems (like Git) give each member a local clone of the full project history, enabling offline work.

The history section explains Git was developed in 2005 by Linus Torvalds (creator of the Linux kernel) to replace the proprietary BitKeeper after BitMover stopped providing free service to the Linux community. A stable version was released just months after development began. The name "Git" is a self-deprecating nod (or "Global Information Tracker"). A VCS timeline: SCCS (1972), RCS (1982), CVS (1986), Perforce (1995), Subversion/BitKeeper (2000), and Git (2005).

The article distinguishes Git (a version control system) from GitHub (a cloud platform hosting Git repositories, created in 2008, owned by Microsoft since 2018, freemium). Competitors include GitLab and Bitbucket. Git is useful beyond code (documentation, music projects). It is popular for being free/open source, fast, offering fine traceability, offline work, ubiquity, and collaboration.

The technical basics cover: repository (`git init`), commits (SHA-1 hash), staging area (`git add`), branches (`git branch`, default `master`), merging (`git merge`), remote repositories (`git remote`), push/pull, fetch, clone, conflict resolution, and log (`git log`). The article closes with an FAQ.

## Key points

- Git is a distributed version control system (dVCS) used by over 90% of professional developers.
- Version control benefits: attributable changes, traceability/rollback, communication, concurrency, branching/merging.
- Distributed systems give each member a full local clone; centralized systems rely on one central server.
- Git was created in 2005 by Linus Torvalds to replace the proprietary BitKeeper.
- Other VCSs: Fossil, Mercurial, Subversion (SVN), CVS, Perforce, RCS, SCCS.
- Git (VCS) and GitHub (cloud hosting platform, 2008, Microsoft-owned) are complementary.
- Git is free/open source, fast, supports offline work, and enables collaboration.
- Core concepts: repository, commits (SHA-1), staging area, branches, merge, remotes, push/pull, fetch, clone, conflict resolution, log.
- The default branch is traditionally `master` (now often `main`).
- Git is used beyond software: documentation, design, writing, even music.

## Technical data / figures

| Concept | Command / Detail |
|---------|------------------|
| Initialize repository | `git init` |
| Stage changes | `git add` |
| Commit | Commit with unique SHA-1 hash |
| Create branch | `git branch <branch-name>` |
| Default branch | `master` |
| Merge | `git merge` |
| Manage remotes | `git remote` |
| Push / pull | `git push` / `git pull` |
| Fetch | `git fetch` (retrieve without auto-merge) |
| Clone | `git clone` |
| View history | `git log` |

| VCS timeline | Event |
|--------------|-------|
| 1972 | SCCS created by Bell Labs |
| 1982 | RCS developed (Purdue) |
| 1986 | CVS developed (first centralized repo) |
| 1995 | Perforce developed |
| 2000 | Subversion (SVN) and BitKeeper |
| 2005 | Git created by Linus Torvalds |
| 2008 | GitHub launched |
| 2018 | Microsoft acquires GitHub |

| Detail | Value |
|--------|-------|
| Developer usage | >90% use Git |
| GitHub founding | 2008 |
| GitHub owner | Microsoft (since 2018) |
| Git author | Linus Torvalds (2005) |

## Why this source matters for the RAG

This guide provides a foundational, well-structured explanation of Git, its history, benefits, and core concepts, making it an excellent retrieval source for version-control fundamentals and Git vs GitHub distinctions. Its concrete command references and historical timeline add both practical and contextual value to the knowledge base.

---
id: vague2-datacamp/datacamp/git-merge-vs-git-rebase
title: "Git Merge vs Git Rebase : Avantages, inconvénients et meilleures pratiques"
domain: datacamp
role: reference
task: article
actors: []
dates: ["2026-09-23"]
keywords: ["cost"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/git-merge-vs-git-rebase.md
source_anchor: ""
source_lines: [1, 65]
sha256: 290b71f65bf6061a9c6a83a492334cf81746f3676d02110e4e027bd33136f697
---

# Git Merge vs Git Rebase : Avantages, inconvénients et meilleures pratiques

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/git-merge-vs-git-rebase
- **Site** : DataCamp
- **Type** : Guide / Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp guide by Dario Radečić compares Git's two main branch-integration strategies: merge and rebase. The TL;DR: `git merge` preserves the full development history by creating new commits that combine branches without altering existing commits; `git rebase` rewrites history by replaying a branch's commits onto another, creating a linear narrative but changing commit SHA hashes. Use merge for collaboration and audit trails, rebase for clean, private development.

**Core concepts**: `git merge` is a non-destructive integration mechanism creating a new commit combining both branches. It uses a three-way merge algorithm comparing the common ancestor with each branch's current state, creating merge commits as junction points. It is ideal for collaborative environments and audit-sensitive projects (financial software, medical apps, regulated industries requiring full traceability). Trade-off: preserves valuable context but increases visual complexity.

`git rebase` rewrites commit history by taking a branch's commits and replaying them on top of another branch, moving the feature branch to start from the target branch's latest commit. Git temporarily removes commits, updates the base, and reapplies changes one by one, giving each a new SHA. It handles conflicts with finer granularity (commit-by-commit). It works well for private branches but rebasing public branches others have based work on causes serious coordination problems and duplicated commits.

**Workflow implications**: merge-centric teams emphasize feature isolation and periodic integration ("integration days"), creating complex commit graphs but preserving full context and enabling easy rollback of entire features. Rebase-oriented teams practice continuous rebasing, emphasize squashing commits and keeping a clean history; the Linux kernel project is cited as a large-scale example, with Linus Torvalds advocating rebasing feature branches before submission.

**Conflict resolution**: merge presents all conflicts in a single session, preserving context and creating a permanent record of reconciliation, but can be overwhelming for many conflicts. Rebase forces iterative, commit-by-commit resolution, breaking complex problems into manageable pieces and keeping each commit coherent, but can be tedious for long-lived branches (same conflict resolved repeatedly).

**History preservation vs readability**: merge preserves temporal authenticity and collaboration context (merge commits as timestamps), at the cost of narrative coherence. Rebase builds a logical sequence, removing experimental noise and easing debugging, but introduces historical revisionism that can hide how decisions were made.

**Hybrid approaches**: many high-performing teams rebase locally to clean up commits, then merge for team integration. Tooling (GitKraken, SourceTree, GitHub network graph, VS Code/IntelliJ conflict editors, CI/CD) influences strategy. Platform features matter: GitHub's "Squash and merge" and GitLab's rebase options.

The article concludes that the choice depends on team size, project phase, and compliance requirements, with an FAQ.

## Key points

- `git merge` preserves full history, creates merge commits, doesn't alter existing commits (non-destructive).
- `git rebase` rewrites history by replaying commits, creating linear history but changing SHA hashes.
- Merge suits collaboration and audit trails; rebase suits clean, private development.
- Merge resolves all conflicts in one session; rebase resolves iteratively, commit by commit.
- Never rebase public branches others have based work on — causes coordination problems and duplicated commits.
- Merge preserves temporal authenticity and collaboration context; rebase offers narrative clarity.
- Merge-centric teams do periodic integration; rebase-oriented teams rebase continuously and squash commits.
- The Linux kernel project is a large-scale example of effective rebase use.
- Hybrid workflows (rebase locally, merge for team integration) offer the best of both.
- Choice depends on team size, project phase, and compliance requirements.

## Technical data / figures

| Aspect | Git Merge | Git Rebase |
|--------|-----------|------------|
| History | Preserved, non-destructive | Rewritten, linear |
| Commit hashes | Unchanged | New SHA per commit |
| Merge commits | Yes | No |
| Conflict resolution | All at once | Commit by commit |
| Best for | Collaboration, audit trails | Private branches, clean history |
| Risk | Complex visual history | Coordination issues on public branches |
| Example command | `git merge feature` | `git rebase main` / `git rebase -i main` |

| Command | Purpose |
|---------|---------|
| `git checkout main && git merge feature` | Merge feature into main |
| `git checkout feature && git rebase main` | Rebase feature onto main |
| `git rebase -i main` | Interactive rebase (squash, reorder) |
| `git add foo.js && git rebase --continue` | Continue after resolving a rebase conflict |
| `git rebase --onto` | Advanced rebase for minimizing disruptions |

## Why this source matters for the RAG

This guide provides a thorough, practical comparison of Git merge and rebase, including mechanisms, conflict handling, workflow implications, and hybrid strategies, making it an excellent retrieval source for Git workflow and history-management queries. Its concrete commands and decision guidance offer high-value, actionable content for the knowledge base.

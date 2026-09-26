---
id: collect-240926-datacamp/datacamp/git-merge-vs-git-rebase-avantages-inconvenients-et-meilleures-pratiques-1
title: "On main, bring in feature-xyz:"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean_en/datacamp/git-merge-vs-git-rebase-avantages-inconvenients-et-meilleures-pratiques.md
source_anchor: ""
source_lines: [1, 117]
sha256: fa03fde2ef896dfe166a9eca10a7eab920ac4f121c3974cb08c3b6f679a3513a
---

# On main, bring in feature-xyz:

<!-- source: https://www.datacamp.com/fr/blog/git-merge-vs-git-rebase -->

Course

If you've ever looked at a messy Git history and wondered which commits really matter, you need to learn the difference between rebase and merge.

Git's two main branch integration strategies - merge and rebase - have fundamentally different goals in your development workflow. Your choice between these strategies has a direct impact on code review efficiency, debugging sessions, and the long-term maintainability of the project. A poor approach can turn your delivery history into an unreadable mess or remove important collaboration context that your team needs.

In this guide, you'll learn when to use each strategy, how they handle conflicts differently, and how to implement hybrid workflows that give you the best of both worlds.

## TL;DR: Git Rebase vs Git Merge

Looking for a quick answer?

- Git merge preserves the full development history by creating new commits that combine branches without modifying existing commits.
- Git rebase rewrites history by replaying commits from one branch onto another, creating a linear narrative but changing the SHA hashes of the commits.

Use merge for collaboration and audit trails, and rebase for clean, private development.

Keep reading to learn more!

## Understanding the basics

Before choosing the right integration strategy, you need to understand how each method works at a fundamental level. Let's break down the basic mechanics of these two approaches.

### Understanding git merge

Git merge is a non-destructive integration mechanism that preserves the full development history of your project. When you merge branches, Git creates a new commit that combines the changes from both branches without modifying existing commits.

This approach is ideal in collaborative environments where multiple developers are working on the same codebase at the same time. You can see exactly when features were developed, who worked on what, and how the different branches evolved over time.

Audit-sensitive projects can benefit from the transparency behind `git merge`. Financial software, medical applications, and other regulated industries often require complete traceability of code changes for compliance purposes.

The command uses a three-way merge algorithm that compares the common ancestor of both branches with the current state of each branch. This creates merge commits that serve as junction points in your commit graph, showing where the branches converged.

In short, it preserves valuable context, but it also increases visual complexity. increases the visual complexity of the project's history.

```
# On main, bring in feature-xyz:
git checkout main
git merge feature
# ➜ Creates a merge commit:
#    *   Merge branch 'feature'
#    |\
#    | * feature change 1
#    | * feature change 2
#    * | hotfix on main
#    |/
#    * initial commit
```
Image 1 - Git project history after a `git merge` operation.

> To learn more about Git merge, feel free to read our complete guide.*read our complete guide.*

### Understanding git rebase

Git rebase fundamentally rewrites your commit history by taking the commits from one branch and replaying them on top of another branch. Instead of creating merge commits, rebase moves your entire feature branch to start from the latest commit on your target branch.

This process involves a historical revision where Git temporarily removes your commits, updates the base branch, and then reapplies your changes one by one. Each delivery gets a new SHA hash, which creates new deliveries containing the same changes but with different parent relationships.

Rebasing handles conflicts with finer granularity than merging. Instead of resolving all conflicts at once, you'll encounter and resolve conflicts commit by commit as Git replays your changes.

This approach works exceptionally well for private branches where you're the only developer making changes. You can clean up your delivery history, group related deliveries together, and present your team with a well-defined sequence of changes.

However, rebasing public branches that others have based their work on can create serious coordination problems and duplicate commits in your shared history.

```
# Bring your feature branch up to date:
git checkout feature
git rebase main
# If you want to squash or reorder:
git rebase -i main
# → opens editor with:
#   pick abc123 Feature commit 1
#   pick def456 Feature commit 2
# Change to:
#   pick abc123 Feature commit 1
#   squash def456 Feature commit 2
```
Image 2 - Git project history after a `git rebase` operation.

*> If you're a beginner, this Git rebase guide will get you up and running.*

## Workflow and branch management implications

Your choice between merge and rebase determines how your entire team collaborates and manages code integration. Each strategy creates distinct development patterns that affect everything from daily workflows to long-term project maintenance.

Let's look at these two aspects in more detail.

### Merge-centric development patterns

Teams that use merge-intensive workflows typically emphasize feature isolation and periodic integration cycles. Developers create feature branches, work independently for days or weeks, then merge their completed work into the main branch during a major update.

This model encourages developers to focus on completing entire features before integration. You will often see teams schedule regular "integration days" during which multiple feature branches are merged simultaneously. This approach works well for teams that prefer a clear separation between development phases and integration phases.

Merge-centric workflows tend to create more complex commit graphs as team size increases, with multiple merge commits creating a web-like structure that can make it difficult to trace the evolution of specific features. However, this complexity comes with the following advantage: it preserves the full context of feature development and integration.

The branching model creates natural checkpoints that allow you to easily revert entire features without affecting other development work. This safety net appeals to teams working on critical applications where stability takes precedence over development speed.

> Did you know that you could revert merge commits? *Our Git Revert guide, packed with examples, will show you how to do it.*

### Rebase-oriented development models

Rebase-focused teams adopt continuous rebasing practices where developers regularly update their feature branches with the latest changes from the main branch. Instead of waiting for feature completion, team members rebase their work several times a day to stay current with ongoing developments.

These teams place strong emphasis on squashing commits and preserving history. Before merging any work, developers combine related commits, rewrite commit messages for clarity, and ensure that their branch tells a coherent story of feature development.

The Linux kernel project is the perfect example of effective large-scale rebasing. With thousands of contributors worldwide, the project maintains a remarkably clean commit history through strict rebasing standards. Linus Torvalds himself advocates rebasing feature branches before submission, arguing that a preserved history makes debugging and code review significantly more efficient.

Rebase-oriented teams often report faster code review cycles, because reviewers can follow a logical progression of changes rather than deciphering the chronological chaos of collaborative development.

**> Are you a beginner looking to keep your repositories t***idy? The Git clean command is all you need.*

## Learn the basics of Git today

## Conflict resolution dynamics


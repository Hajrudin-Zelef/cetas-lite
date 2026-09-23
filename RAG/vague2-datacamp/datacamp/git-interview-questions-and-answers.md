---
id: vague2-datacamp/datacamp/git-interview-questions-and-answers
title: "Les 25 meilleures questions d'entretien Git et leurs réponses, pour tous niveaux"
domain: datacamp
role: reference
task: article
actors: []
dates: ["2026-09-23"]
keywords: []
source: docs/RAG/Collect RAG Vague 2/02_datacamp/git-interview-questions-and-answers.md
source_anchor: ""
source_lines: [1, 69]
sha256: ef3ec3d47fbb68a51c9f5333a3df8ff810ed2aa460cf7440370d72efc0b57f67
---

# Les 25 meilleures questions d'entretien Git et leurs réponses, pour tous niveaux

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/git-interview-questions-and-answers
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp article by Kurtis Pykes prepares readers for technical interviews with 25 Git interview questions and answers, from beginner to advanced. It notes Git was created by Linus Torvalds in 2005 to support Linux kernel development and is now the backbone of countless software projects.

**Beginner questions**: what a Git repository is (stores files and revision history, local or online); how Git works (records changes as successive snapshots, supports branches, merges, reverting); `git add` (stages changes for the next commit); `git push` (sends local commits to a remote repo); `git status` (shows current repo state — modified, staged, untracked files); what a commit is (a snapshot with a unique ID and descriptive message); branching (diverging from main to work on features/fixes); conflicts (incompatible changes to the same file portion, marked with `<<<<<<<`, `=======`, `>>>>>>>`, resolved manually); and merging (combining changes from multiple branches into one).

**Intermediate questions**: what a remote is (a repo hosted on a server for collaboration, e.g., GitHub/GitLab/Bitbucket); how to undo a pushed public commit (`git revert <commit-hash>`, safer than reset/amend); `git stash` (temporarily stores uncommitted local changes); `git reflog` (shows HEAD pointer changes and recent actions, useful for recovering lost commits); and tracking a local branch with an existing remote branch (`git branch --set-upstream-to` / `-u`).

**Advanced questions**: managing multiple configurations (`git config --global/--system/--local`, `includeIf`); handling large files (Git LFS); `git submodule` (managing external dependencies, with update steps); `git cherry-pick` (applying a specific commit from one branch to another, e.g., backporting fixes); `git bisect` (binary-search debugging to find the commit that introduced a bug); and Git hooks (client-side like `pre-commit`/`commit-msg`, server-side like `pre-receive`, stored in `.git/hooks/`).

**Commonly confused concepts**: `git fetch` vs `git pull` (fetch retrieves without merging; pull fetches then merges); `git reset` (`--soft`, `--mixed`, `--hard`, with warnings about rewriting shared history); `git push --force-with-lease` vs `--force` (safer, checks remote hasn't changed); `git rebase` vs `git merge` (linear history vs merge commits, never rebase shared branches); and `git clone` vs `git fork` (local copy vs server-side copy for open-source contribution).

The article closes with preparation tips (master fundamentals, practice real cases, learn troubleshooting, do mock interviews).

## Key points

- 25 Git interview questions across beginner, intermediate, advanced, and commonly confused concepts.
- Git was created by Linus Torvalds in 2005 for Linux kernel development.
- `git add` stages changes; `git commit` creates a snapshot with a unique ID.
- Conflicts are marked with `<<<<<<<`, `=======`, `>>>>>>>` and resolved manually.
- `git revert` safely undoes public commits without rewriting history (unlike reset/amend).
- `git stash` temporarily saves uncommitted changes; `git reflog` recovers lost commits.
- `git fetch` retrieves without merging; `git pull` fetches then merges.
- `git reset` modes: `--soft` (keeps staged), `--mixed` (unstages), `--hard` (discards — dangerous on shared branches).
- `git push --force-with-lease` is safer than `--force` (checks remote state).
- `git rebase` gives linear history but rewrites it; never rebase shared branches.
- `git clone` makes a local copy; `git fork` makes a server-side copy for open-source contributions.
- Advanced: Git LFS for large files, submodules, cherry-pick, bisect, and hooks.

## Technical data / figures

| Command | Purpose |
|---------|---------|
| `git add` | Stage changes for the next commit |
| `git status` | Show repo state (modified, staged, untracked) |
| `git commit` | Snapshot with unique ID and message |
| `git push` / `git pull` | Send / fetch+merge remote changes |
| `git fetch` | Retrieve remote changes without merging |
| `git revert <hash>` | Undo a public commit without rewriting history |
| `git stash` | Temporarily store uncommitted changes |
| `git reflog` | Show HEAD/reference history; recover lost commits |
| `git branch -u <remote>/<branch>` | Track a remote branch |
| `git reset --soft/--mixed/--hard` | Move HEAD; soft keeps staged, hard discards |
| `git push --force-with-lease` | Safer forced push |
| `git cherry-pick <hash>` | Apply a specific commit to another branch |
| `git bisect` | Binary-search for the bug-introducing commit |
| `git clone` / `git fork` | Local copy / server-side copy |
| Git LFS | Large file storage |

| Conflict markers | Meaning |
|------------------|---------|
| `<<<<<<<` | Start of conflicting changes |
| `=======` | Separator between versions |
| `>>>>>>>` | End of conflicting changes |

## Why this source matters for the RAG

This Q&A reference provides a structured, level-graded inventory of Git concepts and commands, including advanced operations and commonly confused distinctions, making it an excellent retrieval source for Git interview preparation and command reference. Its concrete syntax examples and safety warnings deliver high-value, actionable content for the knowledge base.

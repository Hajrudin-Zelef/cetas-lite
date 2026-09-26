---
id: collect-240926-datacamp/datacamp/les-25-meilleures-questions-d-entretien-git-et-leurs-reponses-pour-tous-niveaux-2
title: "Git extracts an intermediate commit; you test, then:"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/clean_en/datacamp/les-25-meilleures-questions-d-entretien-git-et-leurs-reponses-pour-tous-niveaux.md
source_anchor: ""
source_lines: [124, 249]
sha256: 65a0981aaea9e7edd363b8aefc6bcb00a537f46626061d61e3c33402bbcc077d
---

# Git extracts an intermediate commit; you test, then:

### How do you manage multiple configurations depending on the project in Git?

Use `git config` with the `--global`, `--system`, or `--local` flags to adjust settings at different levels. You can also use `includeIf` in the Git configuration to include specific settings based on the repository path.

### How do you handle large files with Git?

Large files can weigh down the repository and degrade performance. Use Git LFS to store these files outside the Git repository while keeping lightweight pointers in the history. This reduces the repository size and improves performance. Git LFS supports various storage providers and integrates naturally with Git workflows.

### What is git submodule for and how do you update one?

The `git submodule` command allows you to manage external dependencies within a Git repository. It lets you include external repositories as submodules in your main repository, which is handy for integrating third-party code while keeping it separate from your codebase.

To update a submodule:

1. 
Navigate to the submodule's directory in the main repository.
2. 
Use `git fetch` to retrieve the latest changes from the submodule's remote repository.
3. 
To advance to the latest commit of the branch tracked by the submodule, use `git pull` .
4. 
Otherwise, to target a specific commit or branch, use `git checkout` with the desired hash or branch name.
5. 
Once in the desired state, commit in the main repository to record the submodule's new revision.

### What is git cherry-pick and when should you use it?

`git cherry-pick` applies a specific commit from one branch onto another, without merging the entire branch.

`git cherry-pick <commit-hash>``main` but you also need the fix on a `release` branch: you can retrieve only that commit rather than merging the entire `main` branch into `release`.
Also useful when a commit was made by mistake on the wrong branch: cherry-pick it onto the right one, then revert it from the one where it doesn't belong.

### What is git bisect and what is it used for?

`git bisect` is a debugging tool that uses binary search to find the commit that introduced a bug. Rather than testing commits one by one, you tell Git a "good" commit (without the bug) and a "bad" commit (with the bug); Git will then check out intermediate commits, halving the search space until it finds the culprit.

```
git bisect start
git bisect bad                # the current commit contains the bug
git bisect good <commit-hash> # this older commit was healthy
# Git extracts an intermediate commit; you test, then:
git bisect good   # or git bisect bad
# repeat until the first bad commit is identified
git bisect reset  # return to the initial state
```
This is much faster than manual testing in a large repository.

### What are Git hooks and how do you use them?

Git hooks are scripts that run automatically at key moments in the Git workflow. They live in a repository's `.git/hooks/` directory and can be written in any scripting language.

There are two types:

- 
**Client-side**: run locally — for example `pre-commit` (before a commit is created) or `commit-msg` (validating the commit message format).
- 
**Server-side**: run on the remote repository — for example `pre-receive` (before accepting pushed commits).

A common use case is a `pre-commit` hook that automatically runs a linter or a test suite before allowing a commit, in order to enforce quality standards.

Note that hooks are not copied during a clone; teams therefore share them via a dedicated script or a tool like `pre-commit` (the Python package).

## Questions about commonly confused Git concepts

### What is the difference between git fetch and git pull?

The main difference between git fetch and git pull lies in their effect on the local repository.

`git fetch` retrieves changes from a remote repository and updates the remote-tracking branches (e.g. origin/master) without modifying your working directory or merging anything into the current branch. This lets you examine new changes without impacting your work.

`git pull` also retrieves changes, but goes further: it chains a fetch followed by a merge into your current branch, directly integrating the remote updates.

### What is git reset for?

The `git reset` command repositions HEAD to a given state. It allows you to undo changes, remove files from the index, or move HEAD to another commit. There are three main modes:

- `--soft`: moves HEAD to a specific commit while keeping changes in the index. The files remain modified and ready to be re-committed.

- `--mixed`: moves HEAD and removes changes from the index. The files remain modified in the working tree, but are no longer staged.

- `--hard`: moves HEAD and deletes all modifications in the working tree and the index. Use with caution: uncommitted changes are permanently lost.

**Important:** never use `git reset --hard` on commits that have already been pushed to a shared branch. It rewrites history and will cause serious problems for your colleagues. Prefer `git revert` for public commits.

### Why prefer git push --force-with-lease over git push --force?

`git push --force-with-lease` is a more cautious way to force a push than `git push --force` because it avoids inadvertently overwriting someone else's work on the remote repository.

With `git push --force`, you force the update without checking whether the remote branch has been modified since your last fetch, which can erase other developers' work.

Conversely, `git push --force-with-lease` checks that the remote branch has not changed since your last fetch. If it has, the push is rejected, preventing the unintentional overwriting of others' changes.

### What is git rebase and how does it differ from git merge?

git rebase and `git merge` integrate changes from one branch into another, but in different ways.

- 
`git merge` combines the histories of two branches by creating a new "merge commit." This preserves the full history of divergences and reunions, which is useful for auditing and team transparency.
- 
`git rebase` "replays" the commits of one branch on top of another to obtain a linear history, without merge commits. The log is more readable, but the history is rewritten. Golden rule: *never rebase a branch that others are working on*.

### What is the difference between git clone and git fork?

**Cloning** creates a local copy of a remote repository on your machine. You remain connected to the original repository and can push changes to it (with the necessary permissions).

`git clone https://github.com/user/repo.git`
**Forking** creates a server-side copy of someone else's repository under your own account — usually on GitHub or GitLab. You own that fork and can push to it freely. When your changes are ready, you open a pull request to the original repository.

Forking is the standard workflow for contributing to open source projects when you don't have direct write access to the original repository.

## Preparing well for a Git interview

Highlighting your Git knowledge and experience in an interview is crucial to demonstrate your mastery of collaboration workflows and development tooling.

Here are a few tips for preparing your technical interview and effectively presenting your Git skills:

### Master the fundamentals of Git

Make sure you understand the fundamentals: repositories, branches, merges, commits, and basic commands like `pull`, `push`, `clone`, and `commit`. This foundation will structure your discussions in the interview. It is also helpful to have a good grasp of the key principles of version control and the differences between Git and other systems.

Finally, familiarize yourself with Git methodologies such as Git Flow, GitHub Flow, and GitLab Flow. Assess their advantages and limitations, and know when to apply them.

Our comprehensive guide to Git is a good starting point for reviewing the fundamentals.

### Practice on real-world cases


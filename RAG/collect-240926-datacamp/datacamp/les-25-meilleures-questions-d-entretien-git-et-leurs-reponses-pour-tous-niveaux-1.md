---
id: collect-240926-datacamp/datacamp/les-25-meilleures-questions-d-entretien-git-et-leurs-reponses-pour-tous-niveaux-1
title: "Git extracts an intermediate commit; you test, then:"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean_en/datacamp/les-25-meilleures-questions-d-entretien-git-et-leurs-reponses-pour-tous-niveaux.md
source_anchor: ""
source_lines: [1, 123]
sha256: 13e97fc88e37471417712a6b1f9a58560bf05c9013407bcd7a70cc765898e17d
---

# Git extracts an intermediate commit; you test, then:

<!-- source: https://www.datacamp.com/fr/blog/git-interview-questions-and-answers -->

Course

Git is an essential tool in the toolbox of modern developers, recognized for its powerful version control capabilities. Created by Linus Torvalds in 2005 to support the development of the Linux kernel, Git has since become the backbone of countless software projects around the world. Its efficiency and flexibility in version management, combined with solid support for collaboration, make it indispensable for teams of all sizes.

This article aims to prepare you for technical interviews by covering the top 20 Git interview questions, from beginner to advanced level. Whether you are new to Git or looking to deepen your knowledge, these questions and answers will help you demonstrate your mastery and succeed in your interview.

## Become a data engineer

## Git questions for beginners

If you are new to Git, some interview questions will likely cover basic concepts and simple use cases. To review these fundamentals, check out DataCamp's Introduction to Git course.

### What is a Git repository?

A Git repository stores a project's files and the history of their revisions, and facilitates version control by tracking changes over time. It can be hosted locally in a folder on your machine or online on a platform like GitHub. This allows you to collaborate, revert to earlier versions, and efficiently manage development through commands like commit, push, and pull.

### How does Git work?

Git records changes made to a project's files and directories as successive snapshots. You can track modifications, create branches to develop in parallel, merge branches, and revert to previous states if necessary. It promotes collaboration and ensures efficient version control in software projects.

### What is git add?

The `git add` command is used to stage changes for the next commit. It prepares additions, deletions, or modifications made in the working space so that they are included in the next snapshot. Note that it does not validate the changes: it simply places them in the index.

### What is git push?

The `git push` command sends the contents of the local repository to a remote repository. It transfers commits from the local repository to a remote server, typically GitHub or GitLab. This command promotes collaboration by sharing your changes with other project members.

To learn more about git push and git pull, check out our dedicated tutorial.

### What is git status?

The `git status` command displays the current state of the repository. It indicates which files have been modified, which are ready to be committed, and which are untracked. It helps track the progress of work and identify what needs to be added to the index or committed.

### What is a commit in Git?

A commit represents a snapshot of the changes made to a repository's files at a given moment. When you commit your modifications, you save the current state of your files and can add a descriptive message explaining the changes (highly recommended).

Each commit has a unique identifier, which makes it possible to trace the repository's history. Commits are essential to version control: they allow you to go back, review the history of modifications, and collaborate by sharing updates.


*Check out DataCamp's Git cheat sheet to prepare for your interview*

### What is branching in Git?

Branching involves diverging from the main line of development (usually called `main`, formerly `master`) in order to work on new features, fixes, or experiments without impacting the main codebase. This allows multiple lines of development to coexist in parallel within the same repository.

Each branch follows its own line of commits, which allows multiple developers to work simultaneously on distinct topics. Branching facilitates collaboration, experimentation, and organization: once the changes are finished and tested, they can be merged into the main branch.

### What is a conflict in Git?

Conflicts occur when incompatible changes are made to the same portion of a file by different contributors, usually during a merge or rebase. Git cannot resolve them automatically and requires manual intervention.

To resolve a conflict, open the relevant file: Git marks the conflicting sections with the `<<<<<<<`, `=======`, and `>>>>>>>` markers. Edit the file to keep the correct version, remove the markers, then:

```
git add <resolved-file>
git commit
```
`git mergetool` make this process more visual and easier to follow.
### What is merging in Git?

Merging is a fundamental operation that facilitates collaboration and the integration of changes between different branches of a project. It involves combining the modifications from several branches into one, usually the main branch (e.g., master or main).

A merge integrates the changes from one branch into another and creates a new commit that brings together the histories of both branches. To learn more about resolving merge conflicts, check out our dedicated tutorial.

## Get certified for the Data Engineer job of your dreams

Our certification programs help you stand out and prove to potential employers that your skills are suited to the job.

## Intermediate Git questions

### What is a remote in Git?

A remote is a repository hosted on a server or another computer for collaborating and sharing code. It serves as a central point where developers can push their local changes and pull those of others.

Remotes are generally configured on platforms like GitHub, GitLab, or Bitbucket. They enable distributed development and facilitate teamwork by providing a common location to store and synchronize code among multiple contributors.

### How do you undo a commit that has already been pushed and made public?

The `git revert <commit-hash>` command allows you to undo a commit that has already been pushed and made public.

Step-by-step procedure:

1. Identify the commit to undo by finding its hash. Use `git log` to browse the history and retrieve the desired hash.

2. Once you have the hash, run the git revert command followed by that hash to create a new commit that undoes the changes introduced by the targeted commit. For example:

`git revert <commit-hash>`
3. Git opens an editor for you to enter the revert message. Modify it if necessary, then save and close.

4. After saving, Git creates a new commit that effectively undoes the changes introduced by the original commit. This new commit is added to the history.

5. Finally, push the new commit to the remote repository to make the revert public:

`git push origin <branch-name>` With `git revert`, you create an undo commit without rewriting history. This is safer than `git reset` or `git amend`, which can modify history and disrupt collaborators who have already pulled the changes.

### What is git stash?

`git stash` temporarily stores local modifications that are not ready to be committed. It allows you to save your work in progress without committing it to the repository.

Stash is useful when you switch branches without wanting to commit or lose your changes. Later, you can apply the stash to your working directory or pop it to pick up where you left off.

### What is git reflog?

git reflog displays reference logs, which record changes to the HEAD pointer and the history of commits viewed in the repository. It provides a chronological list of recent actions (commits, checkouts, merges, resets).

The reflog is valuable for recovering lost commits or branches and understanding the sequence of actions carried out in the repository.

### How do you make a local branch track an existing remote branch?

To make a local branch track a remote branch, use `git branch` with the `--set-upstream-to` or `-u` option, followed by the name of the remote branch.

The syntax is as follows:

`git branch --set-upstream-to=<remote-name>/<branch-name>`
or

`git branch -u <remote-name>/<branch-name>`
## Advanced Git Questions


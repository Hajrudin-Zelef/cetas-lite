---
id: collect-240926-datacamp/datacamp/les-25-meilleures-questions-d-entretien-git-et-leurs-reponses-pour-tous-niveaux
title: "Git extracts an intermediate commit; you test, then:"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["open source", "reasoning"]
source: docs/RAG/clean_en/datacamp/les-25-meilleures-questions-d-entretien-git-et-leurs-reponses-pour-tous-niveaux.md
source_anchor: ""
source_lines: [1, 272]
sha256: 552625b0761f595a9aefdf40f298670ab5d3076a653d32203da20d5de5b44a22
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

The more you use Git, the more your knowledge becomes ingrained. Regular practice makes you comfortable with commands and work habits. Integrate Git into your daily routine, experiment with creating and merging branches, and practice resolving conflicts.

If you lack project ideas, contributing to open source on GitHub is an excellent way to expose yourself to the tools and collaboration workflows used in the industry.

### Know Common Issues and Their Troubleshooting

You will inevitably encounter problems with Git: merge conflicts, detached HEAD state, reverting changes, recovering lost commits, etc. Diagnosing these situations strengthens your troubleshooting skills and your understanding of Git's internal mechanisms.

By actively analyzing error messages and practicing, you will become more efficient at identifying and resolving incidents, reduce risks, and gain confidence in managing version workflows.

### Practice with Mock Interviews

Mock interviews help you target your areas for improvement, both in Git knowledge and communication.

They also offer the opportunity to practice realistic Git-related scenarios and coding exercises. This practice builds confidence and improves your ability to clearly explain your reasoning on the day.

## Conclusion

Git is a powerful version control system, widely used to manage code changes, collaborate, and preserve project history. Mastering Git is essential in technical interviews: it proves your comfort with essential tools and workflows, your ability to collaborate, and to manage code effectively as a team.

Understanding Git concepts and commands allows you to adopt effective versioning practices, ensuring code integrity, project continuity, and smooth development processes. These skills are valuable for engineers and developers seeking successful interviews and a fulfilling career.

To go further, consult:

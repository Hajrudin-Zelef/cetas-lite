---
id: collect-240926-datacamp/datacamp/comment-apprendre-git-en-2026-guide-complet-pour-les-debutants-2
title: "comment-apprendre-git-en-2026-guide-complet-pour-les-debutants"
domain: datacamp
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["open source"]
source: docs/RAG/clean_en/datacamp/comment-apprendre-git-en-2026-guide-complet-pour-les-debutants.md
source_anchor: ""
source_lines: [105, 226]
sha256: e9bc3d8c8389a07c61e91d8436c7084664eb91ffdaf6a58222f04254fad9d6ef
---

# comment-apprendre-git-en-2026-guide-complet-pour-les-debutants

#### Record changes made to the repository

Please record even minor changes in order to keep snapshots of the changes. GitHub will notably keep the following information:

- The state of your folders
- Newly created files
- Files modified on stage
- Staged and unstaged changes

#### View commit history

Since you will frequently need to review recorded changes, it is important to learn how to view the history of your commits. In this way, you will not only be informed of the progress of your work, but you will also be able to see:

- Who made the changes
- When the changes were made
- What changes were made

To do this, please use the `git log` command. 

#### Undo changes

Git does not have the traditional "Undo" feature that allows you to reverse your last action. This is why it is quite complex to undo changes in Git, which can lead to significant losses.

It is therefore necessary to first examine the commits and determine what went wrong. For example, you might have committed too quickly or made a mistake in your commit message. It is also possible to accidentally stage a file. Some actions being irreversible, this skill must be mastered with caution.

The elements you need to learn include:

- To determine which validation changes you want to undo, the `git log` command can be useful.
- Unstage a staged file: Please use different commands such as `git restore --staged file-to-unstage`.
- Undo changes with Git restore: the `git revert` and `git reset` commands are used for this purpose.
- Undo local commits: the `git reset --hard` command allows you to remove the desired commits and reset them to their previous state.

#### Learn how to tag

Tagging allows you to mark important points in your project's history, such as released versions. To do this, it is recommended that you learn how to use the `git tag` command to list all tags, create lightweight and annotated tags, and push them to a remote repository.

### 3. Master intermediate Git and GitHub skills.

When it comes to intermediate Git and GitHub skills, you can never learn enough. However, I have highlighted some of the most important intermediate skills that can add value:

#### Branching

As a data professional, you spend most of your time experimenting and fixing mistakes. To do this, you can use Git branches to create a separate line of development. These branches represent pointers to snapshots.

To deepen your knowledge, it is also important to understand how merging allows you to group changes from different branches and integrate new code into the main project.

#### Cloning

Cloning allows you to create a copy of an existing repository. This is the process of cloning all the repository data from GitHub to your local computer. This is an important skill if you want to retrieve a copy of your own repository or someone else's.

Let's compare standard cloning and cloning with submodules:

| Features | Standard cloning | Cloning with submodules |
| Order | `git clone`   | `git clone --recurse-submodules`  |
| Creates a directory | Yes, default repository name | Yes, it also initializes the submodules. |
| Retrieves the full history | Yes | Yes |
| Protocol options | HTTPS, SSH, Git | HTTPS, SSH, Git |

#### Customizing Git

Every company and every user has specific needs, which is why they use Git to adapt accordingly. To do this, they use Git customization to integrate it into workflows. However, to do this, it is necessary to learn Git configuration and its various commands, which are organized according to the following three levels:

- **Local:** Repository-specific settings allowing per-project customization
- **Global:** User-specific settings that apply to all repositories
- **System:** Settings applicable to all users on the system

I have also included in the table below some commonly used Git configuration commands to help you master customization:

| Commands | Function |
| `git config --global user.name`  | Sets the global username for all commits. |
| `git config --global core.editor emacs`  | Sets the default editor for Git commands. |
| `git config --global color.ui auto`  | Enables color output in the terminal. |
| `git config --global alias.co checkout` | Creates an alias for the checkout command. |
| `git config --local commit.template .gitmessage`  | Sets a commit message template for a specific repository. |

### 4. Learn Git and GitHub through practice

Tutorials alone will not be enough to help you understand all of Git's features. It is best to start projects from scratch. Here's how to do it:

- Participate in interactive sessions to practice using Git commands in real time.
- Please review existing issues and pull requests on repositories.
- Participate in open source projects on GitHub

This is how you can gain more practical knowledge about Git, which does not come from reproducing exercises found on Google.

### 5. Build a project portfolio

A well-maintained GitHub portfolio can help you stand out. It's not just about uploading your source code, but also about how you manage projects and make regular commits. You can do this by creating a separate repository for each new project.

### 6. Keep challenging yourself

Mastering Git is an ongoing process. Just as you constantly adapt to new programming languages, it is also important to stay informed about new Git updates. This will help you learn new commands and integrate them into your projects.

## Example Git and GitHub learning plan

If you are new to Git, please follow a learning plan to progress in manageable steps. I have created an example learning plan that covers the entire journey to mastering Git and GitHub:

### Weeks 1 to 3: Introduction to version control and Git fundamentals

- Version control systems: Familiarize yourself with the concept of version control systems (VCS). Understand how Git differs from other version control systems, such as Subversion and BitKeeper. Discover the command-line interface and why it is essential for working with Git.
- Please install and configure Git: Please install Git on your system or sign up on GitHub and configure a username, an email address, and other preferences.
- Basic Git operations: Master the fundamental commands needed to create a new repository.
- View and undo changes: Discover which commands allow you to view the history of changes and undo mistakes. This will also help you understand how going back in a project will benefit you later on.

### Weeks 4 to 6: Advanced Git features and collaboration

- Branching and merging: Discover how Git branching allows you to work on different features and fixes. Start with the commands for creating, managing, and merging branches. Then move on to practicing conflict resolution.
- Remote repositories: To improve collaboration, please familiarize yourself with using remote repositories. Understand the concepts of forking, creating pull requests, and code review.
- Tagging and releasing: Gain knowledge about Git tags and releases.
- Rebasing: Learn how to use rebasing to clean up your commit history.

### From week 7 onward: Mastering Git and beyond

- Submodules: Discover how submodules manage your project's dependencies on other Git repositories. To master this, it is necessary to understand the appropriate commands.
- Advanced Git configurations: Please discover how to customize your Git environment by configuring aliases.

## The best methods for learning Git and GitHub

When you are mastering complex version control systems such as Git, it is best not to limit yourself to a single method. It is best to seek help from online tutorials, books, courses, and other learning resources.

### Online courses


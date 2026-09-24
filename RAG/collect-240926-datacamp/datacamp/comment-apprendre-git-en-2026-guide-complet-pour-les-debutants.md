---
id: collect-240926-datacamp/datacamp/comment-apprendre-git-en-2026-guide-complet-pour-les-debutants
title: "comment-apprendre-git-en-2026-guide-complet-pour-les-debutants"
domain: datacamp
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["copilot", "license", "open source", "research"]
source: docs/RAG/clean_en/datacamp/comment-apprendre-git-en-2026-guide-complet-pour-les-debutants.md
source_anchor: ""
source_lines: [1, 320]
sha256: 7027754416ea5dfeaacb86f6010af343dc8446c845c500547e4716d88cfd3676
---

# comment-apprendre-git-en-2026-guide-complet-pour-les-debutants

<!-- source: https://www.datacamp.com/fr/blog/how-to-learn-git -->

Course

As a data professional, you use a version control tool to track changes made to code and collaborate with your team. Git is one such tool, used by more than 100 million developers worldwide.

A good command of Git is more important than ever, as companies now expect this skill for any position in software engineering and data.

In this article, I covered all the essential aspects to know about Git to follow an effective curriculum, as well as some resources and a detailed learning plan.

## What is Git?

Git is an open source tool for managing different versions of code. It's similar to a folder on your computer where you store your code. Each time you make a change, Git records the changes as a snapshot, allowing you to apply or undo the changes.

It also promotes teamwork, allowing you to make your changes separately and merge them. In case of conflict, for example if two people modify the same part of the code, Git allows you to choose which changes to keep and which to discard.

*Number of developers using GitHub worldwide (in millions). Image source.*

### What makes Git so popular?

With more than 70% market share, Git has become an essential tool for developers around the world. Here's what makes it so popular:

- Fast and allows you to work offline.
- Provides a safe environment for less experienced developers to experiment without compromising the main code.
- Freely accessible without any financial constraints.

### Main features of Git

Among its most useful features are the following:

- Distributed version control: Each user can have a complete copy of the repository. This means you can work offline while having access to all the data you need. If the main server fails, any user's repository can restore it.
- Open source: Anyone can download and modify it, as it is distributed under a free license. Git's local configuration makes it responsive and easy to set up without an Internet connection.
- Minimal data loss: Git is designed to prevent any data loss. You can add data to the repository and don't risk losing committed snapshots.
- Snapshots over deltas: Some version control systems record changes as deltas, which track changes from one version to another. However, Git allows you to save snapshots of the entire project with each commit. This way, you can access any version of a file at any time.
- Automation and CI/CD: Git integrates seamlessly with CI/CD, and you can automate many tasks, such as testing, scheduling, project management, tagging, and integration. This will allow you to streamline your workflows and maintain consistency.
- Branching and merging: This feature makes it easy to manage different lines of development. First, you create separate branches to test ideas without disrupting the main code. Then, you merge these branches into the main project.

**To get started easily, please see this guide to learn how to set up Git.** 

### Different Git platforms

Effectively managing your team's source code depends heavily on the Git hosting provider you choose.

That's why it's essential to choose a platform that fits your budget and integrates with your existing tools. Below, I present the three most popular Git platforms:

1. GitHub is the reference platform for open source projects. It's beginner-friendly and hosts millions of public repositories. In addition, it offers basic project management tools such as issue tracking and project boards.
2. GitLab stands out for its exceptional CI/CD capabilities. It's ideal for working in a dynamic environment where automated workflows and security features are needed. That's why you can use it to optimize your development process.
3. Bitbucket is suited for small teams and private repositories. It's free for teams of up to five users and offers basic CI/CD integration through Bitbucket Pipelines. This tool is also known for its integration with other Atlassian tools, such as Jira and Confluence by . However, support for Bitbucket Server ended in 2024, which could raise security concerns if you continue to use it.

## Learn the basics of Git today

## Why is it so useful to learn Git?

Git has become an essential skill in today's job market, indispensable for anyone seriously looking to enter the tech field.

To help you better understand where it can be used, I covered its applications in various sectors and explained how learning it can help you land well-paid jobs:

### Git has various applications.

Git has become an indispensable tool in many sectors, beyond traditional software development. Let's look at these different applications:

- Research and data science: With Git, you can manage scripts, Jupyter notebooks, and research papers.
- Web development: You can use it to manage the code, assets, and configurations of your website. It is an essential element of development for everyone, from independent developers to large teams.
- DevOps practices: DevOps teams can also use it to automate and manage their infrastructure with fewer errors in the deployment process.
- Mobile app development: Tools such as GitHub Copilot make the mobile app development process easier. It provides instant code suggestions, rapid prototyping, and reduces the risk of errors while you work.
- Machine learning: You can use it in the field of machine learning to control versions of code, notebooks, and models.

Git is also essential for careers in the data field, such as data engineering and machine learning. **If you are considering moving toward these careers and want to understand the role of Git in a broader context, we invite you to consult our Professional Data Engineer and Machine Learning Fundamentals courses.** 

### There is a high demand for Git skills.

If you simply master the basic Git commands, that would be enough to get started in the tech field. However, as your role evolves, it is necessary to refine your existing Git skills in order to progress further.

More than 6,000 job listings on Indeed, ranging from Tableau developers to C++ developers, highlight the demand for Git expertise. Here are the salaries you can expect in positions requiring Git skills:

- Application developer: $58,975 - $141,044 per year
- Control engineer: $102,000 to $150,000 per year
- Front-end developer: $42,500 to $155,500 per year
- Data Scientist: $125,000 to $203,000 per year

## How to learn Git and GitHub from scratch in 2026

Git and GitHub have completely transformed the way you work on your code and collaborate on projects. They make your task considerably easier. However, if you don't know where to start, here's how to proceed:

### 1. Understand why you are learning Git

Before you start learning Git, please make sure that it is not just about learning to use a tool, but rather a comprehensive approach to project management. It would be best if you also took your needs and goals into consideration. To do this, please ask yourself the following questions before you begin:

- What do I already know about this tool?
- Do you want to acquire basic knowledge, or does your role require a deeper understanding of the tool?
- Do you want to contribute to open source projects, collaborate with teams on a complex codebase, or optimize your personal workflow?

Once you have answered these questions, you will be able to better structure your learning path.

### 2. Start with the basics of Git and GitHub

Once your goals are identified, please master the fundamental principles and understand how they work. I have highlighted a few fundamental steps to get started:

#### Create a Git repository

*Please click on "New repository" in the upper left corner — image source.*

To create a new GitHub repository, please click on "New repository" in the upper right corner of the page. The git `init` command can also be used to create a new repository. Please note that you need to create a GitHub account beforehand.

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

DataCamp offers courses suitable for beginners that break down the complex concepts of Git and GitHub into simple lessons. These courses provide you with basic knowledge of Git and advanced skills on GitHub. Once you have gone through them, you will be able to confidently embark on any project related to data science.

Here are a few of my recommendations for you:

- **To better understand Git: Introduction to the Git course**
- **To develop your foundational concepts: Course on GitHub concepts**

### Online tutorials

Traditional learning methods, such as online tutorials, are still widely used and effective for acquiring complex skills like Git and GitHub. At DataCamp, we also offer detailed tutorials that provide step-by-step guides on Git and GitHub.

I have therefore gathered the most relevant tutorials that will help you learn how to install Git, clone branches, and acquire other advanced skills:

- **To install Git: Git installation tutorial**
- **To gain basic knowledge of GitHub and Git: GitHub and Git tutorial for beginners**
- **To learn how to clone a branch: Git branch cloning tutorial**
- **To understand how to clone a specific branch: How to clone a specific branch**
- **To use the Git reset and revert commands: Git Reset and Revert tutorial for beginners**
- **To learn the commands for merging branches and resolving problems:  Tutorial: How to resolve merge conflicts in Git**
- **To understand how to perform Git push and pull requests: Git Push and Pull tutorial** 

### Git cheat sheets

Who doesn't appreciate cheat sheets? Indeed, they have been my reference resources for memorizing key terms and the most complex commands. That is why I also recommend DataCamp's comprehensive cheat sheet for Git enthusiasts:

- **Download DataCamp's Git cheat sheet**

### Books

If you think the era of books is long gone, you are probably mistaken, because there are many remarkable works on Git, and many people prefer to learn from books.

If you are also a reading enthusiast, you can consult these works on Git to deepen your knowledge:

- **For the fundamentals: Git Pocket Guide: A Practical Introduction** 
- **For Git features: Version Control with Git**
- **For project management and collaboration: Getting Started with GitHub**

## Tips for learning Git and GitHub

These tips will help you determine how much time you should devote to practicing Git and how often you should practice it.

### Please select your area of interest

Before starting your learning journey, it is important to determine what you want to focus your efforts on. Since Git is an additional skill, it is important to devote time to perfecting your core skills. For example, as a data scientist, you can adopt a mixed approach that consists of dividing your time between learning coding and version control.

### Practice regularly

Learning from online tutorials and physical resources is an excellent approach. However, it is recommended to develop an interest in hands-on learning before committing to concrete projects. It offers practical experience, which most recruiters look for in candidates.

### Work on concrete projects

The best way to master Git is to work on concrete projects and solve problems encountered in different fields, such as data science, machine learning, or software development. You can also make your projects public so that others can contribute to your work and the GitHub online community.

### Join a community

Online communities are always a great way to learn anything. Therefore, if you are learning Git, it is time to join a project. To this end, we recommend checking out platforms such as Reddit, as they host active groups where you can ask relevant questions and propose solutions.

*Git community on Reddit. Image by the author*

This is an excellent way to interact with experts and benefit from their knowledge and experience.

### Please do not rush

While it is tempting to prioritize speed to get a job quickly, this approach can lead to gaps in your knowledge. This is why it is important to take the time to practice and explore different scenarios in order to gain a solid understanding of Git and its concepts.

## Final conclusions

Git has become indispensable for succeeding in this competitive job market, especially if you work in the technology field. In reality, recruiters prioritize Git over other version control systems because it fosters more effective collaboration within teams and improves workflow.

You can use tutorials and online courses to start acquiring the basic concepts. However, it is equally important to work on concrete projects in order to gain practical experience and build a solid portfolio to advance in your career.

In addition to Git, if you also want to master a programming language and are still unsure which one to choose, Python ranked as the third most used programming language in 2023. You can follow this guide to learn Python from scratch.

## Become a data engineer

I am a content strategist who enjoys simplifying complex topics. I have helped companies like Splunk, Hackernoon, and Tiiny Host create engaging and informative content for their audience.

## Frequently asked questions

### How do I get started with Git and GitHub as a novice?

**If you are starting your career, the complexity of Git and GitHub may seem intimidating. However, by taking small, regular steps and following a structured learning plan, it is possible to master Git and GitHub in a few weeks.**

### Do I need to install Git to use GitHub?

**Many beginners confuse Git and GitHub. Git is free software, while GitHub is a cloud-based hosting service offering some paid features. It is possible to use GitHub without having to install Git beforehand.**

### Is it necessary to know a programming language to use Git?

**No, Git does not work with any programming language. Since it is a command-line tool, you can store your source code in any language.**

### Is Git useful for computer programmers?

**Yes, today, programmers prefer Git for various reasons. However, it is in team collaboration that it excels. This saves them from having to rack their brains to understand the latest part of the code when they collaborate on a project.**

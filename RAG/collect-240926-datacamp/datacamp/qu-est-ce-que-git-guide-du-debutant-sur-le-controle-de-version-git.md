---
id: collect-240926-datacamp/datacamp/qu-est-ce-que-git-guide-du-debutant-sur-le-controle-de-version-git
title: "qu-est-ce-que-git-guide-du-debutant-sur-le-controle-de-version-git"
domain: datacamp
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["exploit", "open source"]
source: docs/RAG/clean_en/datacamp/qu-est-ce-que-git-guide-du-debutant-sur-le-controle-de-version-git.md
source_anchor: ""
source_lines: [1, 178]
sha256: 84112d442721fbb2020083dc246dfab53518b0a8a75977f7b06b4ad5ce2c2777
---

# qu-est-ce-que-git-guide-du-debutant-sur-le-controle-de-version-git

<!-- source: https://www.datacamp.com/fr/blog/all-about-git -->

If you’ve ever read anything about code, programming, or software development, you’ve heard of Git.

Practical (and free), this tool is the most widely used version control system in the world. It’s so widespread that more than 90% of professional developers use it, not to mention experts in other fields.

In many ways, Git has become synonymous with version control. But what is version control, and why is it so important?

Join us as we dive into the world of Git. We’ll take a close look at what Git is, who uses it, and its history.

## What Is Git?

Git is a distributed version control system (dVCS). As the name suggests, version control is about managing and tracking different versions of a given project.

### What Is a Version Control System (VCS)?

A VCS tracks and records changes made to any file (or group of files), allowing you to retrieve specific iterations later or on demand. VCSs are sometimes called source code management (SCM) or revision control systems (RCS).

Version control allows many members of a team to work together on a project, even if they are not in the same room, or even in the same country.

For example, let’s say you’re a songwriter. You’re working at home on a new song, but you’re not quite satisfied with it. So you decide to collaborate with two other songwriters to rework the parts that need improvement.

You and the two other songwriters start tweaking the lyrics and the score, each working on your own. When the other musicians send you their versions, you like some of the changes and like others less.

Now imagine that you could see every change in every version, test them to hear the result, and then synchronize the edits you want to keep across the versions.

That’s exactly what Git makes possible. Everyone can work locally (on their own computer), save the changes that work, and then synchronize those changes to a Git repository so others can see the new version.

Git is often thought of as a software development tool — which it is — but it can be used to version any type of file: lines of code, a website mockup, or even a song.

### The Benefits of Version Control

Beyond collaboration, version control offers other advantages:

- **Attributable changes**. Every change can be tied to a team member.
- **Fine-grained traceability and easy rollback**. Because every change is tracked, even the smallest ones, it’s easy to return to an earlier version if needed. As you can imagine, this is essential in software development.
- **Better organization and communication.** Commit messages — which explain to the team why you made a change — encourage good communication. They also help you find your way if you forget what you changed in the past!
- **Concurrency**. In software projects, developers often modify source code. Usually, several people work on different topics. One improves security, another develops a new feature. Git allows them to work in parallel while limiting conflicts between their changes.
- **Branches and merges.** Team members can create separate branches to work on, then merge their changes into the main branch. Branches are temporary and can be deleted after merging.

### Is Git the Only Version Control System?

No, Git is not the only VCS, but it is the most popular and is considered the reference tool. Other well-known systems include Fossil, Mercurial, and Subversion.

These systems have nuances — for example, in how they handle key functions such as branching and merging — but the general idea remains the same. The big difference lies in whether they are centralized or distributed.

#### Centralized and Distributed Version Control Systems

Centralized and distributed systems, like Git, serve the same function.

The main difference: in centralized systems, a central server receives the latest versions to which team members push their work. You can think of it as a single project shared by everyone.

With distributed VCSs, each member has a local copy (clone) of the project’s complete history on their device. So they don’t need to be online to modify or work on their code. Instead of a central server, they retrieve that clone from an online repository.

With Git, each member’s clone is a repository that can contain all changes since the beginning of the project.

## The History of Git

Git was developed in 2005 by Finnish software engineer Linus Torvalds, who also created the Linux operating system kernel.

Git was created to meet an immediate need. Before its invention, Linux developers around the world used the proprietary software BitKeeper, itself a dVCS.

Because this software belonged to a company, it created tensions within the Linux community, which was mostly committed to open source.

In exchange for free use, BitMover, the company behind BitKeeper, imposed restrictions on the Linux community. According to the Linux Journal, one of them prohibited working on competing version control projects.

As one might expect, a Linux developer began reverse-engineering BitKeeper to create an open-source product. True to its warning, BitMover stopped providing its services to the Linux kernel, plunging the distributed development system into uncertainty.

To resolve the impasse, Torvalds paused his work on Linux for the first time since 1991 and created Git, a stable version of which was released just a few months after development began.

Interestingly, even before the adoption of BitKeeper, developers sent their "patches" (modifications) to Torvalds individually, which he integrated as they came in. And in 2016, 11 years after Git's release, BitKeeper became open source.

### Why is Git called Git?

In his very first code commit on Git in 2005, Linus Torvalds added a readme file that sheds light on the origin of the name. Here is an excerpt:

Unless you prefer the "Global Information Tracker" version, the name Git is a slightly mocking nod to its capabilities — or their supposed absence.

### A brief history of VCS

Version control systems existed well before Git or even BitKeeper. Here is a short timeline:

- 1972 - SCCS, the first VCS, is created by Bell Labs. It bears little resemblance to today's systems.
- 1982 - The Revision Control System (RCS) is developed by a computer scientist at Purdue University.
- 1986 - The Concurrent Versions System (CVS) is developed. It is the first VCS to offer a centralized repository accessible to multiple users.
- 1995 - Development of Perforce, a VCS still popular today.
- 2000 - Arrival of Subversion (or SVN), more sophisticated. And of BitKeeper, one of the first dVCS, which popularized distributed systems.
- 2005 - Invention of Git, which quickly became the preferred tool of developers around the world.

## Git and GitHub, version control and repositories

Git and GitHub are complementary technologies. Git is a version control system, while GitHub is a cloud platform that hosts Git repositories and helps teams manage them.

Designed in 2008 to facilitate collaborative work with Git, this software-as-a-service (SaaS) model excelled at that goal, attracting millions of users worldwide.

Beyond Git's standard functions, GitHub offers its own features such as bug tracking, task management tools, and continuous integration (CI). GitHub operates on a freemium model: many features are free, but full access requires a paid subscription. GitHub has belonged to Microsoft since 2018.

GitHub is not the only repository hosting service, but with millions of users and hundreds of millions of projects, it is by far the most popular. Many large companies are present there, including DataCamp.

Among competing services are GitLab, fully free and open source, designed for Git, and Bitbucket, which supports Git and Mercurial.

As we mentioned, Git and version control are not just about code or software development. The same is true for GitHub, even though it is not optimized for non-code projects.

## Git, much more than a tool for developers

Git can be used for any collaborative project where version management matters: writing a voluminous user guide or even creating liturgical music (a real project to discover on GitHub).

Although it is mostly associated with the heart of software development, professionals in related fields use Git daily. This is the case for data scientists and analysts: they need to manage the code that supports their work, and Git meets exactly that need.

At DataCamp, we teach the essential tools and technologies for working with data, including Git. Find our selection of immersive and engaging Git courses here.

## Why is Git so popular?

Git is widely praised for many reasons, starting with the fact that it is free and open source.

- **Speed**. Git is fast, especially considering that developers create branches and merge entire repositories. Since everyone has a local copy, there is no need to wait for a multitude of small changes to be pushed to a server.
- **Very fine traceability.** Git offers extremely detailed versioning: even the smallest modifications are committed, and developers can leave a timestamped comment explaining each change.
- **Offline work.** With local copies of the complete repository, there is no need to be online before you are ready to push your changes.
- **Ubiquity**. Today, Git is so common that its spread further fuels its popularity. More than 90% of developers use Git, and a company has little reason to adopt another tool if everyone already knows Git.
- **Collaboration**. Git facilitates teamwork: it simplifies merging different versions of the same project while minimizing potential conflicts. With GitHub, developers have an agile collaborative ecosystem that supports their work.

## How does Git work?

To understand the full power and efficiency of Git, one must look at a few technical aspects. Here are the basic principles:

1. **Repository (repo).** A Git repository is a directory where all the files of a given project are stored. It contains all the revisions and the history of the project. When you initialize Git in a folder (`git init`), it becomes a repository.
2. **Commits**. Each change — or set of changes — that you validate in Git is called a commit. Each commit has a unique identifier (SHA-1 hash) that allows Git to track the modifications and their order.
3. **Staging area.** Before validating your changes with a commit, you "stage" them. The staging area is a preparation space where you gather your modifications before validation. To add files to it, use the command `git add`.
4. **Branches**. Git allows you to create multiple lines of development thanks to branches. The default branch is called `master`. To develop a feature or fix a bug, you can create a new branch (`git branch <branch-name>`) in order to isolate your changes without impacting the main line.
5. **Merge.** Once your changes are finished on a branch, you can merge them into the `master` branch (or any other) with the command `git merge`.
6. **Remote repositories.** Even if you work locally, Git also allows you to connect to remote repositories via `git remote`. This is particularly useful for collaborating. As mentioned, the most common remote repository is GitHub.
7. **Push and pull.** Once connected to a remote repository, you can `push` your changes so that others can see them and collaborate, and you can `pull` remote updates to synchronize your local copy.
8. **Fetch**. Similar to `pull`, the command `git fetch` retrieves updates from a remote repository without automatically merging them into your current branch. You can thus examine the changes before integrating them.
9. **Clone**. To get a copy of an existing Git repository, use `git clone`. This creates on your machine a new directory with all the files and the history of the repository.
10. **Conflict resolution.** When several people work on the same portion of code, conflicts can arise. Git includes mechanisms to flag them and allow manual resolution before merging.
11. **Log**. To consult the history of commits, use `git log`. This command displays the list of commits, their unique identifiers, and the associated messages.

Understanding these technical elements constitutes a solid base for working with Git. As you become familiar with these concepts and commands, you will appreciate the flexibility, power, and efficiency that Git brings to version control.

## Want to get started with Git?

Git is the most widely used distributed version control system in the world, and it has profoundly changed the way developers and related professions manage their projects.

Companies like Google, Netflix, and many others use Git as a standard in their technology stack. Git is so ubiquitous that, for any project related to software or code, one can assume it is part of the process.

It is also an indispensable skill for data professions, such as analysts and data scientists. We need to version the code that allows us to exploit data and create software tools to carry out our analyses.

Git is the de facto standard for VCS. If you want to work in IT or an adjacent field, it is an essential skill. Admittedly, Git is not known for its simplicity, but it is fairly easy to master the basics and gradually deepen your knowledge.

DataCamp can support you. Our Introduction to Git course is designed to teach you the essentials of Git in a fun and engaging way. Once you are comfortable, you can consider taking a GitHub certification to showcase your skills.

To discover why more than nine million learners worldwide love DataCamp, sign up today for your first Git course!

## FAQ

### What is the main purpose of Git?

Git is a distributed version control system designed to track changes in source code during software development. It allows multiple developers to work simultaneously on the same project, ensuring that their modifications do not conflict.

### How does Git differ from other version control systems?

Git is a distributed version control system: each developer has a complete copy of the project history on their machine. This contrasts with centralized systems, where a single central repository serves as the reference and from which code is pulled.

### Is Git only for software developers?

Although Git is mostly associated with software development, its versioning capabilities are useful for many types of projects: documentation, design, writing, etc.

### What is the connection between Git and GitHub?

Git is a version control system, while GitHub is a cloud platform that hosts Git repositories. GitHub offers additional features such as bug tracking, task management, and collaboration tools.

### Do I need to be online to use Git?

No. One of the advantages of Git is being able to work offline on your local repository. An Internet connection is only necessary when you want to push your changes to a remote repository or retrieve updates.

### Is Git secure?

Git includes several mechanisms to guarantee the integrity and authenticity of code. For example, commit signing relies on cryptography to verify the source and integrity of validations.

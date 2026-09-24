---
id: collect-240926-datacamp/datacamp/qu-est-ce-que-github-le-guide-ultime
title: "qu-est-ce-que-github-le-guide-ultime"
domain: datacamp
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "agents", "aws", "distribution", "license", "open source"]
source: docs/RAG/clean_en/datacamp/qu-est-ce-que-github-le-guide-ultime.md
source_anchor: ""
source_lines: [1, 249]
sha256: 57a72983846862d1aff5b4b88c2b2e57f572b6653a7c0eeecc8d5095791afb21
---

# qu-est-ce-que-github-le-guide-ultime

<!-- source: https://www.datacamp.com/fr/blog/what-is-github -->

Course

*GitHub logo. Source: GitHub Logos and Usage* 

Imagine you're working on a data science project and have made good progress. Suddenly, a bug appears. You'd like to go back to the last working version, but you no longer remember all the changes you made. Or perhaps you're collaborating with others, and merging everyone's contributions becomes a headache. If these situations sound familiar, you're not alone.

These common problems are solved with GitHub, the go-to platform for version control and collaboration. In this article, we'll look at how GitHub can transform the way you manage your data projects. We'll also explore collaboration techniques and strategies to boost productivity.

Let's start with the fundamentals of version control.

## What is version control?

Version control is a system that records the history of changes made to files over time. It allows multiple people to work together on a project while keeping an accurate record of changes. Without version control, tracking code changes quickly becomes chaotic and error-prone, especially in a team when multiple contributors work in parallel on different parts of the code.

## What is GitHub for?

As you might expect, GitHub excels at version control. But the platform goes far beyond that, with many other uses, including:

- 
**Building a project portfolio:** GitHub lets you create a public profile to showcase your skills and data projects to recruiters or colleagues.
- 
**Collaborating:** GitHub makes it easy to work as a team on projects, share code snippets, and review each other's contributions.
- 
**Contributing to open source:** With GitHub, you can explore and contribute to open source data science projects, accelerating your learning and innovation.

## How does GitHub work?

To get the most out of GitHub, it's essential to understand its key components and how they work together.

- **Repositories:** These are folders that store your project files and their version history. Think of them as a digital filing cabinet for your data projects. Each repository has a unique URL and contains files, branches, and commits.
- **Forks:** A fork is a personal copy of another user's repository. You can make changes to it independently, then propose to merge them back into the original repository.
- **Pull requests:** PRs are a formal way to propose your changes to the project owner for review and merging. They make code review and collaboration easier.
- **Issues:** They are used to track tasks, bugs, or improvements.
- **Branches:** A branch is a parallel version of a repository. Create branches to develop a feature or fix a problem, then merge them into the main branch when they're ready. To learn more, check out this tutorial on Git Clone Branch.
- **Merging:** Merging combines your changes with the original project to keep everything organized and up to date, for example by merging a feature branch into the main branch.

## Git vs. GitHub

You may be wondering what the connection is between Git and GitHub. These two terms are sometimes confused, but there's a key difference.

**Git is a distributed version control system (DVCS) that helps developers manage their code. It tracks changes and allows you to create different versions, or branches, of the code, which makes collaborative work easier. Git also offers features like the staging area and commit history, providing fine-grained traceability of changes.**

GitHub, for its part, adds features like access control, bug tracking, task management, and wikis, which simplifies collaboration on projects. With GitHub, you can manage your code, track changes, review contributions, and discuss issues — all in one place. The platform also integrates with many tools and services to streamline development workflows. 

| Category | Git | GitHub | 
| Definition | Distributed version control system | Web platform built on top of Git | 
| Purpose | Helps manage code, track changes, and create branches | Hosts Git repositories and provides additional collaboration tools | 
| Features | Staging area, commit history, branches, and merging | Access control, bug tracking, task management, wikis, and integrations | 
| Benefit | Enables collaborative work and detailed tracking of code changes | Improves collaboration, project management, and code review processes | 

## How to use GitHub

So far, we've defined GitHub and version control, and compared Git to GitHub. Now let's get practical.

We'll first look at how to create a GitHub account, customize your experience, and choose a plan. Then we'll create a repository, configuring it, adding a description, and managing its visibility. Next we'll cover creating branches to work on different versions of the project, before moving on to commits, where we'll learn how to modify files and document changes.

### Create an account

Here are the steps to create a GitHub account and get started:

1. Go to GitHub and click the **Sign****up** button.
2. Follow the instructions to create your account. Enter your email address, choose a username and a password.
3. Customize your experience by choosing a suitable plan and adjusting your preferences during setup. The free plan is more than enough for beginners and junior data profiles.

    *Creating a GitHub account. Image by the author*

### Create a repository

After creating your account, the next step is to create a repository. Here's how to proceed for your first repository:

1. 
Click the **+** icon in the top right and select **New****repository**.
2. 
Add a name and a description, then choose whether the repository should be public or private. Public repositories are visible to everyone. Private repositories are accessible only to you and the collaborators you invite.
3. 
You can optionally add a README file, a `.gitgnore` file, and a license. You can also add them later.
4. 
Click **Create****repository**.

*Creating a repository. Image by the author*

Once these steps are completed, a quick setup window for your new repository appears. You can start by creating a new file or uploading an existing file to the repository.

*Setting up a new repository. Image by the author*

*Uploading files. Image by the author*

### Create branches

Once the repository is ready, create branches. Here's how to proceed:

1. 
In your repository, click **Branch:main**, near the top of the page.
2. 
Then click the **New branch** button in the top right.
3. 
Enter a branch name and click **Create new branch**.
4. 
You can switch from one branch to another via the branch dropdown menu and select the one you want to work on.

*Creating branches. Image by the author*

### Make commits

After creating branches, move on to commits. Steps to follow:

1. 
Go to the file you want to modify in your repository.
2. 
Modify it by clicking the pencil icon, then make your changes in the text editor.
3. 
Click **Commit****Changes**. Scroll down to the **Commit****changes** section. Enter a commit message accurately describing your changes — this is essential.

*Commits. Image by the author*

### Create a pull request

Once your commits are made, create a pull request. Procedure:

1. 
Go to the **Pull requests** tab of your repository.
2. 
Click **New****pull****request**. GitHub will automatically compare the changes between branches.
3. 
Review the differences to confirm that everything is correct, by comparing your branch with the main branch.
4. 
Create the pull request by clicking **Create****pull****request**. Add a title and a description.
5. 
Add reviewers if necessary, then submit. This is optional, but very useful if you want expert insight and feedback before merging.

Creating a pull request. Image by the author

### Merge branches

After your pull request has been reviewed and approved, all that's left is to merge. Steps:

1. Once the pull request is approved, open it in the **Pull****requests** tab.
2. Click the **Merge****pull****request** button.
3. Confirm the merge with **Confirm****merge**.

Beyond the GitHub interface, you can perform most actions with Git on the command line. For example, create a local Git repository and retrieve changes from a remote repository with the `git pull` command. Becoming familiar with this environment becomes crucial throughout your career. DataCamp's Introduction to Git course is an excellent resource for progressing. You can also start with our GitHub and Git Tutorial for Beginners. 

## Alternatives to GitHub

GitHub is the most popular platform for version control and collaboration. However, several alternatives offer specific approaches and advantages. It is important for data professionals to know about these solutions that may better meet certain needs.

Here is an overview of some alternatives to GitHub.

### GitLab

This is a DevOps platform that offers Git repository management, issue tracking, and CI/CD pipelines. GitLab is appreciated for its all-in-one approach integrating a wide range of tools and services.

*GitLab logo*

Some key features:

- **Integrated CI/CD:** Built-in pipelines for continuous integration and deployment, automating tests and production releases.
- **Issue tracking:** A system for managing tasks, bugs, and feature requests, with boards, milestones, and labels.
- **Auto DevOps:** Automatically configures pipelines based on the project structure, which speeds up setup.
- **Container registry:** A built-in registry for managing Docker images.
- **Security features:** Built-in security testing tools for continuous monitoring and detection of vulnerabilities in your code.

### Bitbucket

Bitbucket is a product of the Atlassian suite that manages Git repositories, including Mercurial repositories. It integrates naturally with other Atlassian products such as Jira and Trello.

*Bitbucket logo*

Key features:

- **Branch permissions:** Fine-grained control of access and modification rights on branches, for greater security.
- **CI/CD integration:** Integration with Bitbucket Pipelines for continuous integration and delivery, with automated testing and deployments.
- **Jira integration:** Direct connection to Jira for comprehensive project management and tracking of development progress.
- **Pull requests:** Inline comments and PR reviews for efficient code reviews.
- **Deployment:** Support for deployments via Bitbucket Pipelines for smooth releases.

### SourceForge

SourceForge is one of the earliest hosting and project management platforms for open source projects. It offers many tools for developers, including code repositories, bug tracking, and project management features.

Key features:

- **Project hosting:** Free hosting for open source projects, with unlimited bandwidth and storage.
- **Download statistics:** Detailed statistics to track project popularity and progress.
- **Discussion forums:** Built-in forums to engage the community and provide support.
- **File distribution system:** An advanced system for managing downloads and releases.
- **SVN and Git support:** Support for Subversion (SVN) and Git.

### AWS CodeCommit

AWS CodeCommit is a fully managed source code management service from Amazon Web Services. It lets you store and manage your code securely in the cloud, with seamless integration with other AWS services.

Key features:

- **Scalability:** Scales to your repository needs without infrastructure management.
- **Security:** Integration with AWS Identity and Access Management (IAM) for granular access control.
- **Integration:** Works with CodePipeline, CodeBuild, and CodeDeploy for a complete CI/CD chain.
- **High availability:** Built on AWS infrastructure for high availability and durability of your repositories.
- **Encryption:** Data encrypted in transit and at rest to strengthen security.

### Cursor Origin

Cursor Origin is a Git forge offered by Anysphere, the team behind the Cursor AI code editor. Unlike older platforms, it is designed from the ground up for AI-assisted development: it hosts your repositories like any Git forge, but treats commits and pull requests generated by agents as first-class elements. You can also mirror your existing GitHub repositories to Origin, which makes it possible to adopt it without overhauling your current workflow.

Key features:

- **"AI-agent-first" design:** Cursor coding agents can push commits, open pull requests, and act directly on repositories, ideal for agent-driven development.
- **GitHub mirroring:** Mirror your existing GitHub repositories to Origin for gradual adoption rather than a full migration.
- **Cursor integration:** Tight integration with the Cursor editor and its cloud agents to bring hosting, editing, and agent execution together within a single ecosystem.
- **Standard Git workflows:** Repositories, branches, and pull requests work as expected, with a CLI-driven setup.
- **Pull request review:** Built-in review to validate AI-generated PRs before merging them into the main branch.

For a step-by-step walkthrough via the CLI and mirroring a GitHub repository, see our Cursor Origin tutorial.

## Conclusion

In this article, we reviewed the basics of GitHub, its importance, and how it works. We first clarified the concept of version control and its role in mastering changes to code and data. Finally, we offered a practical guide to GitHub: signing up, creating repositories, commits, managing pull requests, and merging branches.

The key to success in data is continuous learning. To go further with GitHub and its applications in data science, listen to our DataFramed podcast episode on The Future of Programming with Kyle Daigle, which shares GitHub's COO vision on the evolution of development. You can also explore GitHub Concepts, a complete course on the GitHub ecosystem, and check out our guide to GitHub certifications to understand how they can boost your career.

Experienced data professional and writer passionate about empowering aspiring data experts.

## FAQ

### Is GitHub difficult to learn?

**While Git, the underlying technology, takes time to learn, GitHub offers a user-friendly interface for managing Git features.**

### What is the difference between Git and GitHub?

**Git is a version control system that tracks file changes, while GitHub is a platform that hosts Git repositories online.**

### Is GitHub safe for storing sensitive data?

**Public repositories on GitHub are visible to everyone. Avoid storing sensitive data such as passwords or API keys there. Use private repositories instead.**

### Do I have to pay to use GitHub?

**GitHub offers a free plan with unlimited public repositories and private repositories (limited storage). Paid plans offer additional features such as more storage and advanced collaboration options.**

### What are the alternatives to GitHub?

**GitLab, Bitbucket, SourceForge, AWS CodeCommit**

### How can I demonstrate my GitHub knowledge in an interview?

To showcase your GitHub skills in an interview, consider earning a GitHub certification that validates your achievements, strengthens your resume, boosts your confidence, and demonstrates your commitment to learning.

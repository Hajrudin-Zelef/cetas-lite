---
id: collect-240926-misc/misc/introduction
title: "Introduction"
domain: opencode
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "open source"]
source: docs/RAG/clean_en/misc/introduction.md
source_anchor: ""
source_lines: [1, 121]
sha256: 5b78878ce208e0fefd69f710895c6a957f920964c082a98ec3454f4f09733885
---

# Introduction

<!-- source: https://opencode.ai/docs/fr -->

# Introduction

Start with OpenCode.

**OpenCode** is an open source AI coding agent. It is available as a terminal-based interface, a desktop application, or an IDE extension.

Let's get started.

To use OpenCode in your terminal, you will need:

1. 
A modern terminal emulator such as:
2. 
API keys for the LLM providers you want to use.

The easiest way to install OpenCode is to use the installation script.

You can also install it with the following commands:

- 
**Node.js**
- 
**Via Homebrew on macOS and Linux**We recommend using the OpenCode tap for the most recent versions. The official formula `brew install opencode` is maintained by the Homebrew team and is updated less frequently.
- 
**Via Paru on Arch Linux**

- 
**Via Chocolatey**
- 
**Via Scoop**
- 
**Via NPM**
- 
**Via Mise**
- 
**Via Docker**

Support for installing OpenCode on Windows using Bun is currently in development.

You can also get the binary from the Releases.

With OpenCode, you can use any LLM provider by configuring its API keys.

If you're new to LLM providers, we recommend using OpenCode Zen. It's a curated list of models that have been tested and verified by the OpenCode team.

1. 
Run the `/connect` command in the TUI, select opencode, and go to opencode.ai/auth.
2. 
Sign in, add your billing information, and copy your API key.
3. 
Paste your API key.

You can also select one of the other providers. Learn more.

Now that you've configured a provider, you can navigate to a project you want to work on.

And run OpenCode.

Next, initialize OpenCode for the project by running the following command.

This will allow OpenCode to analyze your project and create an `AGENTS.md` file at the root of the project.

This helps OpenCode understand the project structure and the coding patterns used.

You are now ready to use OpenCode to work on your project. Feel free to ask it anything!

If you're new to using an AI coding agent, here are some examples that might help.

You can ask OpenCode to explain the codebase to you.

This is helpful if there's a part of the codebase you haven't worked on.

You can ask OpenCode to add new features to your project. However, we recommend first asking it to create a plan.

1. **Create a plan**

OpenCode has a *Plan Mode* that disables its ability to make changes and instead suggests *how* it will implement the feature.

Access it using the **Tab** key. You'll see an indicator for this in the bottom right corner.

Now let's describe what we want it to do.

You want to give OpenCode enough details to understand what you want. It helps to talk to it as if you were talking to a junior developer on your team.

1. **Iterate on the plan**

Once it gives you a plan, you can give it feedback or add more details.

OpenCode can scan any images you give it and add them to the prompt. You can do this by dragging and dropping an image into the terminal.

1. **Build the feature**

Once you feel comfortable with the plan, switch back to *Build Mode* by pressing the **Tab** key again.

And ask it to make the changes.

For simpler changes, you can ask OpenCode to build directly without having to review the plan beforehand.

You need to provide enough details for OpenCode to make the right changes.

Let's say you ask OpenCode to make some changes.

But you realize it's not what you wanted. You **can undo** the changes using the `/undo` command.

OpenCode will now revert the changes you made and show your original message again.

From there, you can modify the prompt and ask OpenCode to try again.

Or you **can redo** the changes using the `/redo` command.

The conversations you have with OpenCode can be shared with your team.

This will create a link to the current conversation and copy it to your clipboard.

Here is an example conversation with OpenCode.

And that's it! You're now a pro at using OpenCode.

To make it your own, we recommend picking a theme, customizing the keybinds, configuring code formatters, creating custom commands, or playing around with the OpenCode Config.

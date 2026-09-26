---
id: collect-240926-mindstudio/mindstudio/the-five-types-of-software-you-can-build-with-ai-and-how-to-pick-one-2
title: "the-five-types-of-software-you-can-build-with-ai-and-how-to-pick-one"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "OpenAI", "Z.ai"]
dates: []
keywords: ["agent", "claude", "glm"]
source: docs/RAG/clean_en/mindstudio/the-five-types-of-software-you-can-build-with-ai-and-how-to-pick-one.md
source_anchor: ""
source_lines: [74, 98]
sha256: 823a2a8047e6184bb00ffb559ee10ff8080a47de5093f72c10ce4aaebed2e6a3
---

# the-five-types-of-software-you-can-build-with-ai-and-how-to-pick-one

For more control, some builders separate every piece themselves: a coding agent (Codex or Claude Code) working on files stored in GitHub, a dedicated database service like Supabase, and a hosting service like Vercel for publishing. This setup has more accounts and more moving parts, but it also means any single piece can be swapped without disturbing the rest. The model can change while the database stays put. The hosting service can change while the code doesn’t move.

It’s worth separating two things people often conflate: the coding agent (the tool that reads your files, writes changes, and shows previews) and the underlying model powering it (which might come from OpenAI, Anthropic, or an open-source option like GLM). A coding agent and its model don’t have to come from the same company, and picking one doesn’t lock in the other.

## Frequently Asked Questions

### What’s the easiest way for a nontechnical person to start building software with AI?

Starting with a web app built through a hosted AI builder is generally the lowest-friction path. It avoids app store approval, server setup, and installing any programming language, while still producing something usable across a phone and a computer.

### Do I need to know how to code to build any of these five shapes?

No. Hosted builders are designed so a plain-language description produces a working preview. Coding knowledge becomes more useful, though still not strictly required, once a project moves toward the more advanced setup involving a coding agent, GitHub, and separate database and hosting services.

### When should I use Supabase instead of a builder’s built-in database?

Use a separate database service when the data matters enough that you want the option to move it later, connect it to another tool, or keep it independent of the interface built on top of it. If the goal is simply to get a first working version running, a builder’s built-in database is a reasonable starting point.

### Is a native phone app always better than a web app?

No. A native app adds real overhead, including app store submission and platform-specific work. It’s only necessary when the software needs a phone capability a browser can’t reliably provide, such as background location, Bluetooth, or persistent push notifications.

### What hardware do people typically use for physical AI projects?

Small, general-purpose single-board computers are common when a project needs meaningful processing near a sensor. Simple microcontroller boards work well for narrow jobs like reading one sensor. Home automation platforms make sense when the goal is connecting devices that already exist in a house.

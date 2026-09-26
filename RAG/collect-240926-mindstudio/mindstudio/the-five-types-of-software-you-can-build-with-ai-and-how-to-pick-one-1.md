---
id: collect-240926-mindstudio/mindstudio/the-five-types-of-software-you-can-build-with-ai-and-how-to-pick-one-1
title: "the-five-types-of-software-you-can-build-with-ai-and-how-to-pick-one"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Apple"]
dates: []
keywords: ["agent", "agents", "claude", "research"]
source: docs/RAG/clean_en/mindstudio/the-five-types-of-software-you-can-build-with-ai-and-how-to-pick-one.md
source_anchor: ""
source_lines: [1, 73]
sha256: d1a6d764650bd35ea3487dc309b9165f8da795db2a73d44d709def18380008a1
---

# the-five-types-of-software-you-can-build-with-ai-and-how-to-pick-one

<!-- source: https://www.mindstudio.ai/blog/five-types-of-personal-software -->

## What are the five shapes of software you can build with AI?

Nearly every personal software idea falls into one of five shapes: a local tool that runs on one computer, a web app that opens from a link, a native phone app, a background service with no screen at all, or a hardware project that connects software to the physical world. Figuring out which shape your idea needs, before you touch any AI builder or coding tool, determines what you sign up for and how much complexity you take on.

## TL;DR

- **Personal software** no longer requires a developer. AI builders and coding agents let anyone turn a specific, small-scale need into working software.
- The five shapes are **local tools, web apps, native apps, background services, and hardware projects** , and each one implies a different set of tools.
- **Web apps are the default starting point** for most people because they work across devices without app store approval or extra infrastructure.
- **Native apps are only necessary** when the software depends on deep phone features like push notifications, Bluetooth, background location, or NFC.
- Tools like **Lovable, Replit, and Bolt** package an AI coding agent with a ready-made environment, so nontechnical builders never touch a command line.
- More advanced builders can separate the pieces themselves, using a **coding agent, GitHub, a database service like Supabase, and a hosting service like Vercel** , which trades convenience for portability.
- The **AI model and the coding tool are not the same thing** . A coding agent like Codex or Claude Code can run different underlying models, so the choice of harness and the choice of model are two separate decisions.

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

## Why does the “shape” of software matter before you pick a tool?

Most people start a software project by picking a tool first: they open an AI builder, type a description, and hope something useful comes out. That works for simple cases, but it skips a step that saves time later. Every idea has a shape determined by where its data comes from, who needs to see it, what device it needs to reach, and what happens if it fails.

A tool that tracks ferry arrival times by listening to a ship’s radio broadcast needs a physical receiver near an antenna, a small always-on computer, and a simple phone display. A shared household maintenance tracker needs a database two people can edit, photo storage, and a browser-based interface. Neither needs an app store listing. Neither needs millions of users’ worth of infrastructure. Once you describe the change you want in your life, the shape of the software is usually already implied, before you’ve chosen a single piece of software to build it with.

## What counts as a local tool?

A local tool runs on one computer and stores its data in files or a small local database. There’s no login system, no server, and often no website. This is the simplest shape available and fits things like a document organizer, a private research assistant, or a utility that renames and sorts photos. If only one person ever needs to use it, and it never needs to be reached from a phone or shared with anyone else, a local tool avoids a lot of unnecessary complexity.

## When is a web app the right choice?

A web app opens from a link, runs in a browser, and can usually be pinned to a phone’s home screen like an app icon. It’s the shape recommended as the default for most personal software, because one version works across a laptop, an iPhone, and an Android device, and it can be shared with a spouse or a few coworkers without going through app store review.

This is the shape hosted AI builders are best at. Tools like **Lovable, Replit, and Bolt** let a builder describe what they want in plain language, then generate a working preview and a way to publish it, all inside an environment that’s already configured. There’s no programming language to install and no server to start manually.

Lovable in particular is suited to a clean web app when the goal is the shortest path from an idea to a working interface. Along the way, it presents a real decision: whether to let the tool manage its own database (Lovable Cloud) or connect to a separate database service (Supabase, built on the Postgres standard). The easier path keeps everything inside one service. The more deliberate path takes more setup but keeps data portable if the project ever needs to move, or if the underlying interface gets replaced without touching the data.

## Why would you build a native app instead of a web app?

## One coffee. One working app.

You bring the idea. Remy manages the project.

A native phone app is what most people picture when they hear the word “app,” largely because platforms like Apple’s App Store have trained users to expect that experience. But native apps require more work: app store submission, platform-specific development, and ongoing maintenance across updates.

Native development only becomes necessary when the software depends on a phone capability a web app can’t reliably access, such as persistent push notifications, Bluetooth connections, background location tracking, or NFC. If none of those apply, starting with a web app avoids a significant amount of unnecessary overhead.

## What is a background service, and when do you need one?

A background service is software with no screen at all. It wakes up on a schedule or in response to an event, does a job, and sends the result somewhere else. Examples include checking a public record every morning, processing a file as soon as it arrives, or sending an alert when a sensor crosses a threshold.

These are often the simplest projects in practice: take data from one place, move it to another, on a timer or a trigger. They’re easy to overlook because there’s no interface to point to, but they’re frequently the backbone of a larger personal software project, quietly doing the data collection or monitoring that a web app later displays.

## How do hardware projects differ from the rest?

A hardware project is software that has to live close to the physical world, reading a signal, controlling a device, or capturing something a sensor detects. This covers things like a custom radio receiver, a home display, a physical button, a camera, or a sensor feeding data into a larger system.

A few small computers show up repeatedly here. A general-purpose single-board computer is useful when a project needs real processing power near the sensor. A simple microcontroller board is enough when the job is narrow, like reading one sensor or controlling one device. And home automation platforms are useful when a house already has several smart devices and the goal is just connecting them into one system.

## How do you choose between Lovable, Replit, Codex, and Claude Code?

These tools solve different problems, even though they’re often compared as if they’re interchangeable.

Lovable, Replit, and Bolt are hosted builders: they combine an AI coding agent with a ready-made environment and a publish button. Lovable leans toward fast, clean web apps. Replit is a stronger fit when a project needs more than a browser interface, such as a scripting language, an always-running server, or a scheduled job, because it keeps the build environment, database, and publishing step inside one service. Bolt offers a similar full-stack path: interface, database, authentication, and storage in one place. The tradeoff with all three is the same: convenience in exchange for depending more heavily on a single service.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.


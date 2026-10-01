---
id: collect-261001-ia-llm/ia-llm/deepseek-harness-architecture-1
title: "DeepSeek Harness Architecture"
domain: ia-llm
role: reference
task: reference
actors: ["DeepSeek"]
dates: []
keywords: ["deepseek", "agent", "agents", "distribution", "memory", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/deepseek-harness-architecture.md
source_anchor: ""
source_lines: [1, 82]
sha256: 94dc22edd3a2fd6df9ab8cf9dd9ad8928fd593a9b71c21483e510dca9908d9f2
---

# DeepSeek Harness Architecture 

Read this before changing anything under `packages/`. It assumes you know Cordis; if you do not, start with the primer or the tutorial.

We recommend using an agent to explore the codebase and understand its architecture.

## Cordis 

Cordis is the framework under dsh: plugins contribute services, typed events, and reversible effects to a shared context. Every part of the product is a plugin, including the model adapter, the tool registry, the session log, and the agent loop itself, so each is replaceable from configuration.

There is no privileged core to patch: you extend dsh by mounting a plugin beside the others, and registrations are effects that unwind when their plugin unloads.

## Profiles and bundles 

A running `dsh` is a plugin tree composed at boot from ordered layers.

A **profile** is a named composition stored in the Harness home. It lists the bundles it stacks, holds any out-of-tree plugins it installs, and keeps the user's own `cordis.patch.yml`. `web`, `headless`, `sdk`, `sdk-minimal`, and `acp` ship as templates.

A **bundle** is a distribution format for Cordis config rows and the code they mount, so whatever it inserts stays patchable by the layers above it.

Each declares itself in its own `package.json` under a `dsh` field: `dsh.profile` lists a profile's bundles, and `dsh.bundle` points at a bundle's patch file.

`dsh-base` is the shared first layer of the `web`, `headless`, `sdk`, and `acp` profiles: model adapters, tools, persistence, sandbox and approval policy, settings, credentials, telemetry. `dsh-web-app` adds the browser application, `dsh-headless` adds a one-shot runner with no server, `dsh-sdk-app` adds the SDK JSON-RPC server, and `dsh-acp-app` adds the automation-only ACP server. `dsh-sdk-minimal` is the deliberate exception: one bundle owns its complete explicit SDK tree and does not apply `dsh-base`.

Layers apply to an empty entry list in this order: each bundle in the profile's listed order, then the profile's `cordis.patch.yml`, then the home-level one, then any `--patch` overlay. A patch targets a row by id and replaces its whole config, or inserts new rows.

YAML controls HMR: base enables config-only `dsh-hmr`; headless, SDK and ACP disable it; `sdk-minimal` omits it. Profile patches override these defaults. HMR coordinates watching and reloads; the launcher provides profile data and readiness.

Base includes Plugin Manager for Web and agents.

To see the tree your machine boots:

`dsh --profile web --dump-config`
Any row it prints can be replaced by a patch of your own.

Composition mechanics are in app-boot; config fields are in the generated config catalog.

## Application launch 

Supported Node applications launch through named `dsh` profiles. The shipped profiles are `web`, `headless`, `sdk`, `sdk-minimal`, and `acp`, selected with `dsh --profile <name>` or `dsh <name>`. `plugin` names the management command; a profile with that name requires `--profile plugin`. The TypeScript SDK resolves its same-version `dsh` dependency and selects `sdk`; custom plugin composition remains a profile plus ordered patch files, not another executable or inline application tree. `sdk-minimal` is a repository-owned standalone bundle behind the same launcher, not a caller-supplied Cordis tree.

Vendored CLIs, build-only and test-only executables, direct in-process plugin mounting, and the private browser WebWorker preview are not Harness application launchers. `verify-application-entrypoints` keeps every package bin, executable source, root demo, and the root `start:web` and `dev:web` scripts in an explicit class and rejects a Node application path that bypasses `dsh`.

The Python SDK follows the same application architecture. Its runtime wheel packages the normal `dsh` CLI as `deepseek-harness-sdk-runtime-<platform>-<arch>`, and the client launches `dsh --profile sdk` with an explicit Harness home by default. The minimal example selects the shipped `sdk-minimal` profile. Python exposes profile selection and ordered patch files rather than a complete Cordis tree; persistent external plugins are installed through `dsh plugin`. The removed private direct-config carrier has no compatibility bin or fallback parser.

## Desktop application 

The Electron desktop application carries its exact dsh production runtime in signed resources and owns the reserved `$DSH_HOME/profiles/desktop`. Shared profile helpers initialize its files, reconcile installed bundles, and resolve installation and bundle dependencies without replacing pnpm-owned packages. CLI and Desktop share product data, while executable packages, activation choices, and lockfiles remain separate. The public CLI cannot manage Desktop’s profile.

Electron starts the private Desktop Host in Electron Node mode. The Host invokes the shared CLI profile runner and complete Web application. The window immediately loads packaged Web assets and waits for boot injections before activating client plugins in the same document. Web owns RPC and streams; the desktop carrier connects the local page to the authenticated Host. Node IPC carries boot injections, readiness, fatal errors, and shutdown. Desktop defaults to port `19387`; profile configuration can override it. Shell-owned UI runs plugin transactions through bundled pnpm with normal user and profile configuration.

## Core packages 

Here are some core packages that contribute to the Cordis tree.

| Package | Owns | `ctx` key | 
|---|---|---|
| `core/session` | The append-only `SessionEvent` log and in-memory store | `ctx.sessions` | 
| `core/system-prompt` | Prompt-section and tool-schema assembly | `ctx.systemPrompt` | 
| `core/tools` | The scoped tool registry and guarded execution pipeline | `ctx.tools` | 
| `core/agent` | The `Agent` interface, live registry, and`agent/*` events | `ctx.agents` | 
| `core/agent-loop` | The default driver implementing that interface | `ctx.agentLoop` | 
| `core/scope` | The per-agent scoped-registration primitive | library, no key | 
| `llm/llm` | Message and stream vocabulary plus the adapter seam | `ctx.llm` | 
| `webhook/webhook` | Authenticated-delivery dispatch and Workspace Session creation | `ctx.webhookRuntime` | 

## Events 

Events are the extension points, and picking the right domain is the first decision in most changes.

- **Session events** are durable facts appended to the log and broadcast through`session/event` . Use one when the fact must survive a reload.
- **Agent events** (`agent/*` ) carry a live`Agent` : inbox, step, status, request, validation, continuation. Use one to observe or intercept work in flight.
- **Capability events** attach policy and adapters to a seam (`fs/*` ,`tools/*` ,`telemetry/*` ) without importing the loop.

AgentLoop awaits serial `agent/created` initialization before starting queued work. Initialization failure rolls back creation; agent-loop defines teardown ordering.

The event map lists every event's producers and consumers.

## Turn flow 

A **step** is one model request plus the tools it calls. A **turn** is zero or more steps: it opens before its first input is claimed and closes once nothing is owed.


---
id: vague2-nerdykings/nerdykings/code-world-model-rival-genie-3
title: "Code World Model : Le Vrai Rival De Genie 3 ?"
domain: nerdykings
role: reference
task: article
actors: ["China", "Google", "MiniMax"]
dates: ["2026-08-26", "2026-09-23"]
keywords: ["agent", "agentic", "memory", "reasoning", "video generation"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/code-world-model-rival-genie-3.md
source_anchor: ""
source_lines: [1, 64]
sha256: fbf51ed9d05e254c34162e187ecd519f90cdca3a93248c5056f89bd0a387c06a
---

# Code World Model : Le Vrai Rival De Genie 3 ?

## Metadata

- **Source** : https://www.nerdykings.com/blog/code-world-model-rival-genie-3.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Since Google DeepMind's Genie 3, world models have a clear public image: an interactive world generated directly as video, frame by frame, according to user actions. Impressive, but with a known flaw — over time these worlds drift, forget objects, and change their shapes. A paper published August 26, 2026, **Code World Model (CWM)**, proposes a radically different architecture: instead of entrusting everything to a single video AI, give the world's logic to a coding agent.

The project comes from a paper titled "Code World Model: Coding Agent as World Brain," by three researchers from West Lake University in China and Nanyang Technological University in Singapore. Their goal: turn a coding agent into the true brain of the world. A large language model no longer just chats — it writes executable code that manages environment rules, keeps its state, and computes the consequences of each event. The video model handles only visual rendering.

Why current video world models drift: systems like Genie 3 observe previous frames plus user actions, then generate the most likely continuation. Trained on huge video datasets, they create the illusion of a real game. But unlike a classic video game, no precise program records every object, position, and rule — much of this information stays buried in the neural network's internal representations, hard to inspect, modify, and especially maintain over time. That's why objects change shape, characters forget what they held, and vehicles leaving the camera reappear differently. Over seconds it's unnoticeable, but for a world meant to evolve over minutes or hours, each action must produce a precise, time-persistent consequence — exactly where this architecture hits its limit.

CWM removes that responsibility from the video generator. A large language model acts as a coding agent: when a user requests an action or a new event occurs, it analyzes the situation and generates executable code to update the world (positions, speeds, objects, collisions, vehicle states, past events). Example: a character destroys a car. A classic video model must generate explosion images and try to visually remember the car no longer exists. With CWM, the program directly records that the vehicle was destroyed — even if the camera looks away for a long time. When the camera returns, the system knows the car must still be destroyed. The coding agent isn't invoked every frame — mainly for complex decisions, a new behavior, or a rule change. The rest (moving objects, computing trajectories, detecting collisions) runs deterministically on its own.

The **proxy** is the bridge between code and rendering. The system has a logical world of coordinates, variables, and rules — nothing like a video. Sending all the code directly to the video generator would be too abstract. So researchers created an intermediate representation called proxy: an extremely simplified scene where characters become basic capsules, vehicles simple boxes, and movements precise trajectories. It isn't meant to be pretty — only to indicate where elements are and how they should evolve. A compiler turns the logical world state into proxy video, sent to **MiniMax H3** with a text description of the desired visual style. The proxy imposes positions, movements, and trajectories; MiniMax H3 adds characters, textures, lighting, decor, and visual details. Code decides what exists and happens, the proxy translates it visually, and the video model builds the final appearance.

Versus Genie 3: Google DeepMind's Genie 3 remains the public reference — a real-time interactive world generated frame by frame by a video AI, with no coding agent or explicit logical state. Visually stunning, but all world memory lives in the network weights, with exactly the drift described. CWM doesn't aim to rival raw visual quality (Genie 3 leads there). Its proposition is to separate two functions usually mixed in one neural model: world logic (rules, memory, objects, evolution) and visual representation (textures, colors, styles). This separation makes world state easier to inspect and correct — if a rule is problematic, developers can theoretically consult the program and fix it precisely without retraining the whole video model. Crucially, the same logic can receive several appearances: a character controlled by the same proxy can become a realistic human, an animated hero, or a completely different creature, without changing a line of code — precisely the blind spot of a 100% video approach like Genie 3.

Limits: despite impressive demos, we remain far from a fully AI-generated GTA. The worlds shown are relatively simple. A real game would require simultaneously managing physics, characters, inventories, collisions, objectives, and thousands of persistent events — and in the current prototype the agent still relies on structures and logic prepared in advance by researchers. It can modify, combine, and extend them, but doesn't yet build all the rules of a complex universe from scratch. Rendering keeps some limits of classic video models: code can impose an object's precise position but can't guarantee every visual detail stays coherent. Researchers haven't yet demonstrated this architecture could run a vast, complex, photorealistic world in real time — the terrain where Genie 3 has progressed most.

The author's view: the value isn't replacing Genie 3 tomorrow (researchers present it as a proof of concept). The architectural bet matters: instead of betting everything on one giant neural network for logic and visuals, CWM reintroduces determinism — code that can be read and corrected — at the heart of a generative system, echoing modular coding-agent architectures that separate reasoning, decision, and execution. Genie 3 still leads on visual quality and fluidity, and CWM is no finished product. But if separating logic and rendering holds at scale, it could become the real answer to the world-model problem: a world that truly remembers what happens, instead of one that dreams anew at every frame.

## Key points

- Code World Model (CWM) is from "Code World Model: Coding Agent as World Brain," published August 26, 2026.
- Authors: three researchers from West Lake University (China) and Nanyang Technological University (Singapore).
- Core idea: an LLM coding agent holds the world's logic/state/rules; the video model only renders visuals.
- Solves video world-model drift (forgotten/changed objects) by persisting state as explicit executable code.
- A "proxy" (simplified capsule/box scene) bridges logical state and video rendering via a compiler.
- Rendering uses **MiniMax H3**, given proxy video plus a text style description.
- Separation of logic and visuals makes state inspectable/correctable and allows swapping appearances without code changes.
- Limits: simple worlds, pre-built structures, no demonstrated real-time photorealistic large-scale world; Genie 3 still leads visually.

## Technical data / figures

| Item | Value |
|---|---|
| Paper title | Code World Model: Coding Agent as World Brain |
| Publication date | August 26, 2026 |
| Institutions | West Lake University (CN), Nanyang Technological University (SG) |
| Video renderer | MiniMax H3 |
| Intermediate representation | Proxy (capsules, boxes, trajectories) |
| Comparison target | Google DeepMind Genie 3 |
| Agent role | Writes executable code for world logic/state |
| Deterministic handling | Object movement, trajectories, collisions |
| Status | Proof of concept / prototype |

- Key concepts: **world model**, **logical state**, **proxy representation**, **coding agent as world brain**
- Compared architectures: CWM (logic + render split) vs Genie 3 (pure video generation)

## Why this source matters for the RAG

This article documents an alternative world-model architecture that separates deterministic world logic (code) from generative rendering, addressing the drift problem of video-based models like Genie 3. It is valuable for a RAG knowledge base on world models, generative environments, and hybrid neuro-symbolic/agentic architectures.

## Source URL

https://www.nerdykings.com/blog/code-world-model-rival-genie-3.html

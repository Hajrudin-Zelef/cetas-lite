---
id: collect-240926-nerdykings/nerdykings/code-world-model-le-vrai-rival-de-genie-3-1
title: "Code World Model: The True Rival of Genie 3?"
domain: nerdykings
role: reference
task: reference
actors: ["China", "Google", "MiniMax"]
dates: ["2026-08-26"]
keywords: ["agent", "benchmark", "memory"]
source: docs/RAG/clean_en/nerdykings/code-world-model-le-vrai-rival-de-genie-3.md
source_anchor: ""
source_lines: [1, 42]
sha256: 2fb9114d299805b8e57a39f41aa94266807e3e0e64b226e194096d32ac3bfdcf
---

# Code World Model: The True Rival of Genie 3?

<!-- source: https://www.nerdykings.com/blog/code-world-model-rival-genie-3.html -->

# Code World Model: The True Rival of Genie 3?

Since **Genie 3** from Google DeepMind, we have a fairly clear picture of what a **world model** looks like: an interactive world generated directly as video, frame by frame, based on your actions. Impressive, but with a known flaw — over time, this kind of world drifts, forgets objects, changes their shape. A paper published on August 26, 2026, **Code World Model (CWM)**, proposes a radically different architecture to solve this problem: what if, instead of entrusting everything to a single video AI, we gave the logic of the world to a coding agent? Let me explain the principle, how it compares to Genie 3, and where the limits are.

## The paper that arrives with a clear ambition

The project is called **Code World Model**, or **CWM**. It comes from a paper titled *"Code World Model: Coding Agent as World Brain"*, published on August 26, 2026 by three researchers from **West Lake University** in China and **Nanyang Technological University** in Singapore. Their goal can be summed up in one sentence: transform a coding agent into a true **brain of the world**. Concretely, in their system, a large language model is no longer just there to chat — it writes executable code that manages the rules of the environment, preserves its state, and calculates the consequences of each event. The video model, for its part, only handles the visual rendering.

## Why current video world models drift

To understand the value of this approach, we need to revisit how current video world models work — the kind of system behind **Genie 3**. Simplified, the model observes previous frames as well as the user's actions, then generates the most probable continuation. If you move forward, it produces the next frames of the movement; if you turn the camera, it generates the new viewing angle. Trained on enormous quantities of video, these models give the impression that a real game is running behind the screen.

Except that unlike a classic video game, there is no precise program that records each object, each position, and each rule of the world. A large part of this information remains buried in the internal representations of the neural network — difficult to inspect, to modify, and above all to maintain over time. That's why these worlds can drift: an object changes shape, a character forgets what it was holding, a vehicle that leaves the camera's field of view reappears differently. Over a few seconds, it goes unnoticed. But for an interactive world meant to evolve over several minutes, or even several hours, each action must produce a precise consequence that is preserved over time — and that's exactly where this kind of architecture reaches its limit.

## CWM: a coding agent as the brain of the world

The researchers behind Code World Model therefore decided to remove this responsibility from the video generator. In their system, a large language model acts as a **coding agent**. When a user requests an action or a new event occurs, this agent analyzes the situation and generates executable code to update the world: position of characters, speed, objects present in the scene, collisions, state of vehicles, past events.

Take the example given by the researchers: a character destroys a car. With a classic video model, you have to produce explosion frames, then try to remember *visually* that this car no longer exists. With CWM, the program directly records that the vehicle has been destroyed — even if the camera looks elsewhere for a long time. This information remains inscribed in the logical state of the world, and when the camera returns, the system knows that the car must still be destroyed. The coding agent does not intervene on every frame, moreover: it is mainly called upon for complex decisions, a new behavior to add, or a rule to modify. The rest — moving objects, calculating trajectories, detecting collisions — runs on its own, deterministically.

## The proxy: the bridge between code and the MiniMax H3 rendering

The issue is that at this stage, the system has a logical world made of coordinates, variables, and rules — and that doesn't look remotely like a video. Sending all the code directly to the video generator would be far too abstract. The researchers therefore created an intermediate representation called a **proxy**: an extremely simplified version of the scene, where characters become basic capsules, vehicles simple boxes, and movements precise trajectories. This proxy does not aim to be pretty — it serves only to indicate where the elements are and how they should evolve.

A compiler then transforms the logical state of the world into proxy video, sent to **MiniMax H3** with a textual description of the desired visual style. The proxy imposes positions, movements, and trajectories; MiniMax H3 adds the characters, textures, lighting, decor, and all visual details. The code therefore decides what exists and what happens, the proxy translates that visually, and the video model builds the final appearance.

## Google's Genie 3: the real rival to beat

This is where the comparison with **Genie 3** becomes interesting. Google DeepMind's Genie 3 remains, to this day, the benchmark for mainstream world models: an interactive world generated in real time, frame by frame, directly by a video AI — without a coding agent, without explicit logical state. The result is visually stunning, but all the memory of the world lives in the network's weights, with exactly the type of drift described above over time.

CWM isn't trying to compete on raw visual quality — Genie 3 retains a clear lead there. Its proposition is to separate two functions usually mixed within the same neural model: on one side, the logic of the world, its rules, its memory, its objects and their evolution; on the other, its visual representation, textures, colors, styles. This separation makes the state of the world much easier to inspect and correct — if a rule poses a problem, developers can theoretically consult the program and correct it precisely, without retraining the entire video model. And above all, the same logic can receive multiple different appearances: a character controlled by the same proxy can become a realistic human, an animated hero, or a completely different creature, without changing a single line of code. This is precisely the blind spot of a 100% video approach like that of Genie 3.

## The limits: we're still far from an AI-generated GTA

Despite impressive demonstrations, we remain very far from a new GTA entirely generated and simulated by AI. The worlds presented remain relatively simple. Building a real game would require simultaneously managing physics, characters, inventories, collisions, objectives, and thousands of persistent events — and in the current prototype, the agent still relies on structures and logic prepared in advance by the researchers. It can modify them, combine them, extend them, but it doesn't yet build all the rules of a complex universe from scratch on its own.

The rendering also retains certain limitations of classic video models: the code can impose that an object be at a precise position, without guaranteeing that every visual detail will remain perfectly consistent. And the researchers have not yet demonstrated that this architecture could run a vast, complex, and photorealistic world in real time — the very terrain where Genie 3 has made the most progress.

## My take


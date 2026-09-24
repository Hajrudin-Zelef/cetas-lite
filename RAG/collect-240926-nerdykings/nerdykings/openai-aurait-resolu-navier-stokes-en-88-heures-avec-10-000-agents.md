---
id: collect-240926-nerdykings/nerdykings/openai-aurait-resolu-navier-stokes-en-88-heures-avec-10-000-agents
title: "OpenAI Would Have Solved Navier-Stokes in 88 Hours With 10,000 Agents"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: ["2026-09-01"]
keywords: ["agents", "astra", "claude", "energy", "gpt-6", "lean", "reasoning", "research", "training"]
source: docs/RAG/clean_en/nerdykings/openai-aurait-resolu-navier-stokes-en-88-heures-avec-10-000-agents.md
source_anchor: ""
source_lines: [1, 49]
sha256: 8a6466e2af5b5530ec4edb2129af6ecb7252ddacb969726604b858d3909bc390
---

# OpenAI Would Have Solved Navier-Stokes in 88 Hours With 10,000 Agents

<!-- source: https://www.nerdykings.com/blog/openai-navier-stokes-88-heures.html -->

# OpenAI Would Have Solved Navier-Stokes in 88 Hours With 10,000 Agents

OpenAI claims to have obtained, in just **88 hours**, a proof for a problem that has resisted mathematicians for nearly 90 years. And the model behind this result isn't even GPT-6 Astra: it would be a new internal model, still in training, presented as significantly more powerful, and deployed through nearly **10,000 agents** working in parallel. Their goal: to solve a formulation of the **Navier-Stokes** problem, one of the seven millennium problems, carrying a one-million-dollar reward. If the result is confirmed, it would be one of the most important scientific discoveries ever produced by a model. Except that the story is far more complicated than what OpenAI's announcement suggests — the proof doesn't necessarily solve the problem in the way one might imagine, it relies on very recent human work, and a serious scientific priority controversy has already erupted. Let's look at what OpenAI actually demonstrated.

## The Navier-Stokes equations, a 90-year-old problem

The **Navier-Stokes** equations describe the motion of fluids: water, air, smoke, blood, or even the turbulence around an airplane. They are already used in countless simulations. The real problem is that no one knows whether they remain mathematically stable in all situations. Take a fluid whose motion is perfectly regular at the start: does its evolution always remain regular, or could a tiny region end up accelerating indefinitely until it reaches infinite speed in a finite time? This mathematical breaking point is called a **singularity**, or *blow-up*. In reality, a fluid obviously cannot reach infinite speed — but if a singularity can appear in the equations, it reveals a fundamental limitation of the mathematical model.

Since 1934, thanks to the mathematician Jean Leray, we have known that so-called "weak" solutions exist. No one, however, had managed to determine whether these solutions could always remain perfectly regular. In 2000, the Clay Institute therefore placed this question among the **seven millennium problems**, with a one-million-dollar prize for a valid solution.

## The scenario that OpenAI's system would have constructed

OpenAI's system claims to have found a precise scenario in which the Navier-Stokes equations actually end up producing a singularity. At the start, the fluid is perfectly motionless. An external force gradually creates a vortex, which spins faster and faster, tightens around its center and stretches upward and downward. As the vortex becomes narrower, its speed increases until it theoretically becomes infinite in a finite time — the famous singularity.

But for the demonstration to be valid, it is not enough to apply an infinite force to the fluid: that would amount to causing an explosion with infinite energy, then claiming to have discovered something. The whole difficulty consisted in using a perfectly **regular and bounded** force, while nevertheless succeeding in making a singularity appear — and that is exactly what OpenAI's system would have managed to construct. This type of force is explicitly allowed in certain official versions of the Navier-Stokes problem. If all the mathematical conditions are indeed respected, the demonstration could therefore genuinely solve one of the formulations envisioned by the Clay Institute.

## 10,000 agents, 88 hours, 130 billion tokens

The way OpenAI would have obtained this proof is almost as impressive as the result itself. On September 1, 2026, the company learns that researchers would have made significant advances on two millennium problems. It then decides to test a new internal model, still in training, presented as much more powerful than GPT-6 Astra, and deploys thousands of agents capable of working in parallel: some look for solutions, others verify the calculations, others gather the most promising discoveries.

Initially, about a hundred agents work for 50 hours on the **Euler equations**, which also describe the motion of fluids but without taking viscosity into account. After encouraging results, OpenAI concentrates its resources on Navier-Stokes: approximately **10,000 agents** are mobilized, and after **88 hours** of work, the system manages to construct a complete mathematical demonstration. In total, the agents produced **2.7 million messages** and nearly **130 billion tokens** for this single problem. GPT-6 Astra then intervenes for nearly **17 additional hours** to translate this demonstration into **Lean**, a mathematical verification software, in order to check each step of the reasoning and verify that the conclusions indeed follow from the starting hypotheses — which considerably reduces the risk that an error or a hallucination is hidden in the proof.

## Why it is not (yet) "solved"

Except that Lean only verifies what has been encoded. Mathematicians must therefore still ensure that the hypotheses given to the software correspond exactly to the official problem posed by the Clay Institute. For the moment, Navier-Stokes is therefore still considered **unsolved**, and the one-million-dollar reward has not yet been awarded. This discovery also does not come out of nowhere: the AI did not invent the entire approach from scratch.

## An intellectual debt — and a scientific priority controversy

A large part of the strategy relies on the work of **Diego Córdoba and Luis Martínez Zoroa**, two mathematicians who had developed a succession of layers and vortices producing a cascade toward a singularity. Their work came close to the desired result, but the force obtained did not yet satisfy all the properties required by the Millennium Problem. For his part, **Tristan Buckmaster**, a professor at New York University, had attempted to cross this final step with Claude and Codex, teaming up with a researcher employed by Anthropic but working here in a personal capacity. They had obtained important results on the Euler equations and other nearby systems — and rumors about their progress would later reach OpenAI, which would then have launched its own massive operation.

When OpenAI announced its result a few days later, a controversy immediately erupted. Buckmaster asked whether the private conversations he had had with Codex could have influenced OpenAI's model, and the company's initial response left room for uncertainty. On September 10, OpenAI updated its statement after an internal investigation: the company now claims that the Codex prompts in question could not have influenced the system, either directly or through training. So one cannot really assert that OpenAI used their data — on the other hand, the intellectual debt to the work of Córdoba and Martínez Zoroa is very clear. And this affair raises a fundamental question: how can scientific priority be protected when a company can mobilize 10,000 agents as soon as it hears about a promising lead?

## Why math advances faster than everything else with AI

This discovery also helps explain why artificial intelligence progresses so quickly in mathematics. When a model writes an article or produces an image, the quality remains partly subjective — a human often needs to examine the result. But in formal mathematics, a proof can be tested automatically: the model proposes a step, Lean verifies it, and the system immediately gets a clear signal, valid or invalid.

This loop makes it possible to generate, test, correct, and start over at a speed impossible for a human. Add to that thousands of agents simultaneously exploring different directions, and scientific research begins to look like a gigantic automated process. The result still depends on decades of human work and on choosing the right lead, but the scale and speed have suddenly changed.

## My opinion

If this proof is confirmed by mathematicians, it would be simply enormous: an AI producing a genuine advance on a problem more than 90 years old. And the craziest part is that the model behind this discovery hasn't even been released to the public yet, and it would already be far more powerful than GPT-6 Astra — even though that one has just come out. Now imagine thousands of agents like those working in parallel on the greatest scientific problems: if it's true, OpenAI would have turned 90 years of research into 88 hours of work.

But I remain cautious on two points. First, technically: Lean only verifies what it is given to verify, and until independent mathematicians have confirmed that the hypotheses match exactly the official Clay Institute problem, it is not "solved." Second, ethically: this story of scientific priority with Buckmaster and Codex makes me uncomfortable. That the intellectual debt to Córdoba and Martínez Zoroa be acknowledged, very well — but the fact that a company can throw 10,000 agents at a lead as soon as it hears about it somewhere completely changes the rules of the game for researchers who have neither that computing power nor that speed. If this way of doing research becomes widespread, the real question will no longer be only "can AI solve 90-year-old problems?" but "who owns the discovery when the lead comes from a human and the execution from a machine?"

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.

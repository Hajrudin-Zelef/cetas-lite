---
id: collect-261001-ia-llm/ia-llm/vibe-physics-the-ai-grad-student-1
title: "vibe-physics-the-ai-grad-student"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Sakana"]
dates: ["2024-08", "2025-02", "2025-08"]
keywords: ["agents", "claude", "distribution", "gemini", "reasoning", "research", "sakana", "training"]
source: docs/RAG/collect-261001-ia-llm/vibe-physics-the-ai-grad-student.md
source_anchor: ""
source_lines: [1, 38]
sha256: 6e9f009a1ad61d7f847e2a2f22aa56a89bc6efd31b34fbd0e17fc5bfa42b0729
---

# vibe-physics-the-ai-grad-student

## Subscribe to Anthropic Science

Features on AI-assisted discoveries, practical workflows, and field notes across the sciences.

**Can AI do theoretical physics? In this guest post, professor of physics Matthew Schwartz decided to find out by supervising Claude through a real research calculation, start to finish, without ever touching a file himself. His account of what happened is below.** 

I’m Matthew Schwartz, a professor of physics at Harvard and a principal investigator in the NSF Institute for Artificial Intelligence and Fundamental Interactions (IAIFI). My area of expertise is quantum field theory, which asks what matter is, how particles interact, and why the Universe has the rules it does. One might say I wrote the book on the subject. I’ve been working with modern machine learning tools for over a decade. My first modern ML paper, from 2016, was an early application of deep learning to particle physics. In a *Nature Reviews Physics* piece in 2022, I compared the timescale of AI and human evolution, arguing that transferring understanding between biological and artificial intelligence would become a fundamental challenge. Since then, I’ve been trying to push AI towards more symbolic work (manipulating mathematical expressions rather than numerical data) and the core questions in theoretical physics.

There has been a lot of recent hype about AI scientists doing end-to-end research autonomously. In August 2024, Sakana AI released their AI Scientist, a system designed to automate the entire research lifecycle—from generating hypotheses to writing papers. In February 2025, Google released an AI co-scientist built on Gemini, promising to help researchers generate and evaluate hypotheses at scale. And in August 2025, the Allen Institute for AI (Ai2) launched the open-source Asta ecosystem, featuring tools like CodeScientist and AutoDiscovery to find patterns in complex datasets. Since then, a new entrant has appeared every few months—FutureHouse’s Kosmos, the Autoscience Institute’s Carl, the Simons Foundation’s Denario project, among others—each promising some version of end-to-end autonomous research. Even as these approaches are visionary, their successes to date seem a bit forced: run hundreds or thousands of trials and define the best one as interesting. While I believe we are not far from end-to-end science, I’m not convinced we can skip the intermediate steps. Maybe LLMs need to go to graduate school before advancing straight to the Ph.D.

In mathematics, automated end-to-end AI agents have produced impressive results, at least for a certain class of problems. An early breakthrough was DeepMind’s FunSearch, launched in 2023, and later AlphaEvolve, which used LLMs to make new discoveries in combinatorics. A related project, AlphaProof, earned a silver medal at the 2024 International Mathematical Olympiad, solving problems that stumped all but five human contestants, and in 2025, an advanced version of Gemini achieved the gold-medal standard. And, just as in science, more achievements have continued to follow.

What about theoretical physics? End-to-end AI scientists have found their footing in data-rich domains, but theoretical physics is not one of them. Unlike mathematics, theoretical physics problems can be more nebulous—less about formal proof search and more about physical intuition, choosing the right approximations, and navigating a landscape of subtleties that often trip up even experienced researchers. Even so, there are problems in physics where AI might be better suited. Not yet the paradigm-shifting questions at the frontier, but those where the conceptual framework is established and the goal well-defined. To find out if AI can solve these types of theory problems, I supervised Claude through a real research calculation at the level of a second-year grad student.

In grad school, at least at my institution, first-year theory students (G1s) typically just take classes. Research often begins in the second year. G2 students start with well-defined projects that have a guarantee of success—often follow-ups from previous studies where the methods are established and the endpoints clear. This gives them a chance to learn the techniques, make mistakes in a controlled setting, and build confidence. It's also easy for me as an advisor: I can check their work, spot where they've gone off track, and quickly reorient them.

Advanced students (G3+) work on more open-ended, creative problems. These require choosing your own direction, deciding which approximations matter, and sometimes realizing the original question was wrong (such is the nature of research).

For this experiment, I deliberately chose a G2-style problem. My reasoning was that LLMs can already do all the coursework, so they are past the G1 stage. But if AI can't do the G2 projects—the ones with training wheels, where I know the answer and can check every step—then it certainly can't do the G3+ projects where creativity and good judgment are essential.

The problem I chose was resumming the Sudakov shoulder in the C-parameter. For context, when you smash electrons and positrons at a collider, debris sprays out; the C-parameter is a single number that describes the shape of that spray, and its distribution has been measured with extreme precision. The theory that's supposed to predict that distribution is quantum chromodynamics, the study of the strong nuclear force, which holds nuclei together and powers the sun. The C-parameter is well-defined on paper but brutally hard to calculate, so you approximate. Every approximation is a stress-test—failures tell you something about the foundations of quantum field theory itself: what are the right building blocks and effective degrees of freedom (particles? jets? clouds of gluons?), and what gaps might lead to new insights? At one particular spot on the distribution, a kink called the Sudakov shoulder, the standard approximations break down, and the math starts producing nonsense. The goal of the project was to fix the prediction at this point.

I picked this problem because it connects directly to the foundations of our understanding of quantum theory. But more importantly, it’s a highly technical calculation that I was confident I could do myself. The physics is understood in principle; what's missing is a careful, complete treatment.

The dream was that I could ask:

*Write a paper on resummation to NLL level of the Sudakov shoulder in the C-parameter in e+e- collisions. Include a derivation of the factorization formula, comparison with previous results, numerical checks against Monte Carlo calculations using EVENT2, and a final plot of the resummed distribution with uncertainty bands.*

and out would pop the paper. We are not there yet, of course. I tried giving this prompt to all the frontier models, and—predictably—they all failed pitifully. But I wanted to see if I could *coach* the model to succeed: to show, rather than tell it.

To go about this scientifically, I encapsulated all the work. The rules were strict:

My question was: is there a set of prompts, like instructions to a talented G2, that can guide an AI to produce a high-quality physics paper (one that is genuinely interesting and pushes the field forward)?

I knew from experience that LLMs struggle with context and organization over long projects. So I started by asking Claude to come up with a plan of attack: what tasks needed to be done in what order. I also asked GPT 5.2 and Gemini 3.0. Then, I had all three LLMs merge the best ideas from each, using web interfaces and copying one to another. Next, I gave those merges to Claude, asking it to break the outline into detailed subsections. The result is here. There were 102 separate tasks across seven stages.

From there, I turned to Claude Code, using the extension in VS Code.


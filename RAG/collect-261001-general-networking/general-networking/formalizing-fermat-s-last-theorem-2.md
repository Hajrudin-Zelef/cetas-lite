---
id: collect-261001-general-networking/general-networking/formalizing-fermat-s-last-theorem-2
title: "formalizing-fermat-s-last-theorem"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Kevin Buzzard", "United States", "Z.ai"]
dates: []
keywords: ["fermat", "agents", "claude", "consumer", "cost", "cyber", "formalization", "glm", "lean", "research", "safeguards"]
source: docs/RAG/collect-261001-general-networking/formalizing-fermat-s-last-theorem.md
source_anchor: ""
source_lines: [54, 78]
sha256: 1ebcc4f4854696af5b6b2089232baf9b51171624b6b4cfd7d51f1723563f275b
---

# formalizing-fermat-s-last-theorem

Formalization is also a major factor in how humans can gain confidence in AI-generated mathematical results. As AI and AI-assisted mathematicians produce more (purported) proofs than ever before, AI-assisted formalization takes part of the load off human reviewers. We expect it will become common to produce a formalized proof alongside any write-up intended for a human reader. Although we do not think a formalized proof should replace a human-understandable exposition, it may be the only feasible way for the mathematical community to keep up with AI-generated contributions.

Writing Lean also seems to help Claude prove novel results. Many of our recent Claude-authored results have been formalized in parallel with their proofs, and Claude appears to use these partial proofs to independently check its hypotheses much like it writes numerical simulations to check that it’s on the right track.

Formalizing FLT was a token-intensive project, but it is also the largest Lean proof ever constructed. Anthropic researchers did a small experiment using three personal Claude Max plans to formalize applications of the Hardy-Littlewood Circle Method. Collaborating entirely through Prove2Me, the agents jointly completed a formalization of Vinogradov’s Three Primes Theorem in just three days. We think with the right scaffold, collaborative formalization of major results with consumer AI subscriptions is achievable.

To this end, Anthropic as well as other labs have recently expanded their support for external researchers—including mathematicians working on pure math and formalization—with free and discounted subscriptions and research credits. We also offer dedicated grants for larger scientific projects, which could include formalizing other major theorems or improving Lean or Mathlib.

With AI rapidly changing what it looks like to do math research, mathematicians—at Anthropic and elsewhere—are grappling with what that means for their work. Formalization, however, is a place where we feel unambiguously good about the role of AI. As formalization becomes a more commonplace tool, we are hopeful that it will help maintain trust in the common body of mathematical knowledge.

Our formalization effort is a small piece of the long history of Fermat’s theorem and the development of formal mathematics. The first full proof from Andrew Wiles together with Richard Taylor was a culmination of more than 300 years of mathematics, integrating ideas from Gerhard Frey, Jean-Pierre Serre, Ken Ribet, Barry Mazur, Robert Langlands, Jerrold Tunnell, Yutaka Taniyama, Goro Shimura, and André Weil, among others. Claude’s proof follows the exposition by Henri Darmon, Fred Diamond, and Richard Taylor.

Our proof adapts pieces from the Imperial College London FLT project led by Kevin Buzzard and the flt-regular project. Lean and Mathlib are both their own labors of love and have received contributions from hundreds of mathematicians, many working with the Lean FRO. We thank Kevin Buzzard for reviewing the proof and for his comments.

The full proof is available on GitHub along with a written walk-through of the proof.

We built an index of how well today’s robots can perform US job tasks. Robots can already do three-quarters of physical tasks, mostly in limited settings, but are cost-competitive for just 0.3% of them.

Read more
We’re launching a new study using Anthropic Interviewer to learn from your experiences with AI, and we invite you to participate.

Read more
Like Claude Mythos Preview, GLM-5.3 has strong capabilities for autonomously building end-to-end cyber exploits. But GLM-5.3 is unlike other frontier models in that it has been released without meaningful safeguards to limit misuse.

Read more

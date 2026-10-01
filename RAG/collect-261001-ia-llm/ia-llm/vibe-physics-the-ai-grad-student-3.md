---
id: collect-261001-ia-llm/ia-llm/vibe-physics-the-ai-grad-student-3
title: "vibe-physics-the-ai-grad-student"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: ["2025-08", "2025-12", "2027-03"]
keywords: ["agentic", "agents", "claude", "memory", "opus 4", "research"]
source: docs/RAG/collect-261001-ia-llm/vibe-physics-the-ai-grad-student.md
source_anchor: ""
source_lines: [79, 128]
sha256: f206aefcd0507d0eba813f2f58dd994dddfe64cc7fec1a325870073335c128c1
---

# vibe-physics-the-ai-grad-student

At this stage, all that remained was massaging the text and the figures. To be fair, the style of scientific writing varies tremendously between disciplines. And although I gave some examples, it couldn’t match my style. I went back and forth between micromanaging sentences—“rewrite this,” “be more positive about previous work”—and letting it get by with its choppy, repetitive style. (In truth, I have misgivings about whether human-readable prose is the right medium for science communication going forward. But that’s a different post.) For the figures, Claude didn’t care a whit about font size, label placement, etc., so there was a lot of “move this label up a little” and so on. But these things are relatively painless with Claude—you just say move this, move that, and it requires no concentration, unlike adjusting label placement by hand within Python code, which requires recalling and looking up finicky syntax

The final money plot was:

This looks similar to the other one, but—after many cross-checks—I can confirm it’s actually correct.

Beyond these larger structural problems, there was a constant stream of smaller errors that required intervention. Some highlights:

When I asked Claude to verify that its formulas expanded correctly to fixed order, it kept producing “verification” documents that invented coefficients that were not in the paper. When I called Claude out, it said:

*No, this document is not helpful either. It has serious problems: It invents terms not in our paper... ‘using the standard SCET consistency condition, the coefficients in Appendix B are constructed to satisfy...’ This is not a verification*.

It was generating plausible-sounding justifications for answers it hadn’t actually derived.

The one-loop soft function calculation (one of the core ingredients in the paper) was supposed to be straightforward: set up the eikonal integrals, regulate in dimensional regularization, and extract the finite part. But Claude’s initial setup in the draft seemed off. When I asked it to start from scratch:

*I see the issue—line 317 and 371 just assert that soft radiation increases C linearly as δC ~ ω/Q without derivation. Let me verify my calculation very carefully from first principles. The soft contribution is definitely δC = (3/2)*k_x^2*—quadratic in the out-of-plane direction, not linear!*

It had been asserting a result without checking. Eventually, GPT solved the integral and Claude incorporated it. They needed each other, and I needed both of them.

When I gave Claude Code the implementation guide for NNLL resummation, it couldn’t implement it. It would see a formula in our paper and simplify it based on patterns from other examples, not accounting for the specifics of our particular case. After hours of debugging:

*You’re absolutely right—I cheated! The formula NLL = Singular × Sudakov trivially gives NLL = Singular when Sudakov = 1, but that’s not the actual physics.*

When I started reading the draft in detail, it was a mess. In particular, there were lots of “zombie sections” it forgot about, repetitions, and guesses that it pretended to derive. I had to go section by section, having Claude reorganize things, like:

The formula you reference in deriving the factorization formula in Eq. 13 is for 3 partons. You need to start with the all-orders formula Eq. 9 and expand when there are 3 partons plus soft and collinear radiation.

Claude had no trouble doing this once I pointed it out. But it didn’t do it without my prompting.

The final paper is a valuable contribution to quantum field theory. Notably, it has a new factorization theorem. There aren’t that many of these, and it’s these kinds of theorems that lead to a deeper understanding of quantum field theory. And it makes novel predictions about the physical world that can be tested with data. Again, this is relatively rare these days. I’m proud of the paper. People are reading it, using it for physics, and engaging in a follow-up project looking at comparison to data from experiments.

Given Claude’s contribution to this paper, I wanted to have Claude as co-author. Unfortunately, current arXiv policy forbids this. The justification is that LLMs can’t take responsibility. This is a good point. So I added to my acknowledgments:

M.D.S. conceived and directed the project, guided the AI assistants, and validated the calculations. Claude Opus 4.5, an AI research assistant developed by Anthropic, performed all calculations including the SCET factorization theorem derivation, one-loop soft and jet function calculations, EVENT2 Monte Carlo simulations, numerical analysis, figure generation, and manuscript preparation. The work was conducted using Claude Code, Anthropic’s agentic coding tool. M.D.S. is fully responsible for the scientific content and integrity of this paper.

Such recognition of integrity and responsibility is important. After all, it would not be good for science if people put out AI slop and then blamed the LLM for its errors. On the other hand, grad students are often on papers with implicit responsibility for the content even when they cannot fully understand it, which is why everyone knows it’s truly the PI’s fault when something is wrong.

One final recommendation I’d give is to move away from the web-based LLMs. These have been around for a while, and are good. But for me the real phase transition was running Claude Code with access to files, terminal commands, agents, skills, memory etc. It makes a big difference.

This paper started out as an experiment: how close are we to end-to-end science with AI? My conclusion is that current LLMs are at the G2 level. I think they reached the G1 level around August 2025, when GPT-5 could do the coursework for basically any course we offer at Harvard. By December 2025, Claude Opus 4.5 was at the G2 level.

What this means is that although LLMs cannot yet do original theoretical physics research autonomously, they can vastly accelerate the research done by experts. For this project (which I completed with Claude in two weeks), I’d estimate that it would have taken me and a G2 student 1-2 years, and me without AI around 3-5 months. Ultimately, it accelerated my own research tenfold. That’s game-changing!

There are two natural follow-up questions that arise from this project. How do we get from here to an AI Ph.D.? And what are human grad students supposed to do now?

I don’t have great answers to these questions. By blunt extrapolation, LLMs will be at the Ph.D or postdoc level in around a year (March 2027). I’m not sure how we’ll get there—maybe we need domain experts to train them, maybe they will train themselves, maybe it will be some combination of the two. I am more confident that the bottleneck is not creativity. LLMs are profoundly creative. They simply lack a sense of which paths might be fruitful before walking them. I think we can distill what is missing in current LLMs to a single word: *Taste*.

In physics, taste is the intangible sense about which research directions might lead somewhere. I've been doing research in theoretical physics for a long time and have learned to tell pretty quickly whether an idea is promising or not. I suspect anyone who has honed a craft for a long time—whether in science, carpentry, or design—would recognize this: experience produces a kind of judgment that AI has not yet mastered. We do not give enough credit to taste. When solving problems is hard, the solution gets the glory, but when knowledge and technical strength are ubiquitous, it’s the taste to come up with good ideas that distinguishes great work.


---
id: collect-261001-ia-llm/ia-llm/claude-computes-a-nine-loop-amplitude-in-n-4-super-yang-mills-2
title: "claude-computes-a-nine-loop-amplitude-in-n-4-super-yang-mills"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "China", "OpenAI"]
dates: []
keywords: ["claude", "compute", "accelerator", "cost", "fable 5", "gpt-6", "research", "training"]
source: docs/RAG/collect-261001-ia-llm/claude-computes-a-nine-loop-amplitude-in-n-4-super-yang-mills.md
source_anchor: ""
source_lines: [45, 92]
sha256: 88e189da2642841c1b2cc65f3b0cea06151bd13eb6409057d499a868cff1ec6f
---

# claude-computes-a-nine-loop-amplitude-in-n-4-super-yang-mills

But he hadn’t computed it, and neither had anyone else in the field. The way he found the eight-loop answer was already a bit indirect, via a surprising link to a different but related formula called a form-factor, a kind of partial amplitude involving different particles that turns out to be a bit easier to calculate. He was expecting to find the next loop even more indirectly, potentially by a different kind of AI method. If people thought it was possible to just run the usual bootstrap method for one more loop, someone would have done it.

Apparently, there are folks at Anthropic who read my blog.

At the end of August, Liam Fitzpatrick and Siddharth Mishra-Sharma, two physicists at Anthropic, reached out to me to say they had tackled one of the challenges in my post. After verifying the result with Lance, they talked me through how they got it.

True to the spirit of the challenge, they didn’t use millions of dollars in computer power. They used Fable 5.1, working within Claude Science, a platform scientists can pay to use. Claude Science is what folks in the biz call a “harness,” a program that uses the Claude LLM with structured rules and prompts in order to get more robust and scientifically useful behavior.

Apparently, after asking Claude which problem it was most likely to be able to tackle, they gave it a simple prompt:

“The problem is to compute the Six-particle (hexagon) amplitude in planar N=4 SYM at nine loops.”

From there, they just kept telling it to keep going, with comments like:

“I'm going to sleep and won't be available for another several hours. Keep working on this until I tell you to stop. Give me updates every 4-6 hours.”

Claude ended up doing the calculation two different ways: the original bootstrap, and the indirect form-factor approach. Either approach would have cost an end-user around one or two thousand dollars, mostly due to the expense of running Claude for so long. The bootstrap calculation, done with the Python programming language with package SymPy, took around $100 of the budget, corresponding to running 96 CPUs for a week.

Running 96 CPUs for a week might have felt like a lot when I was doing this kind of work ten years ago, but it’s pretty affordable now if you have a good reason.

As it turned out, the result wasn’t all that far away for humans either. A few days after I heard from Anthropic, we heard from Song He, an amplitudeologist at the Chinese Academy of Sciences in Beijing. Song’s group had already gotten the majority of the result. They’d used some AI assistance, based on GPT-6, but not the kind of one-shot almost human-less approach Anthropic used.

Everyone has been friendly here, which is a bit of a relief. The humans, Lance and Song and their collaborators, will get to publish the results, taking time to explain them and analyze them for the benefit of future researchers. Claude’s role is done, for now.

I set my challenge because I wanted a better sense of what current AI can do, and where it could go from here. So what have I learned?

I’d thought this could be a chance to see AI overcome a computational barrier in a surprising way. Instead, it did something it turned out humans were also able to do. Claude used known methods, with a bit more compute than people had tried to use before. It may have gotten a boost from using Python, and not Maple (Lance’s favorite program for math) or Mathematica (mine), and it may have used much better software engineering practices than we would have, but not super-intelligently so.

My biggest takeaway is that there is more low-hanging fruit out there than you’d expect. Even when a goal is simple and well-defined, sometimes it’s going to look much less achievable to experts than it actually is. There are people with a computer science background who’ve been telling me for years that amplitudeologists could make a lot more progress just by hiring a few programmers. They should feel vindicated.

It’s also noteworthy that Claude Science accomplished this in one shot, without any scientific oversight more sophisticated than “keep going.” These are finicky, messy calculations. If I’d used a week of time on 96 CPUs to do this kind of calculation, then I’d almost certainly end up using two weeks: it’s practically guaranteed I’d screw up something on the first try. I don’t know how many mistakes Claude made internally on the way, but the harness got it to the end without an outside collaborator’s input. I’m not sure that surprises me, at this point. But if you didn’t know it could do that because you’re still thinking of AI as so error-prone that it’s unusable, then this should be your takeaway: It can do this kind of thing reliably now.

Things definitely seem to be moving fast. In March, AI was accomplishing physics projects like a student: smaller-scale tasks with a lot of hand-holding and mistakes. In contrast, this is a real frontier calculation, the kind of thing normally tackled by the top experts in amplitudes. While it’s possible that this is just a much more AI-friendly problem, I don’t think it’s just that: I think the technology has genuinely gotten better.

How far can I generalize this? That I’m not sure of.

These toy model theories tend to be the focus of small sub-communities. The real-world amplitudes calculations are a wider field, with many groups trying to beat each other to the frontier. It’s possible there’s less low-hanging fruit there. But I wouldn’t count on it. I know people who work on those calculations have been increasingly using AI for coding. If people aren’t already checking whether AI science harnesses can one-shot frontier calculations there, they ought to (and they ought to have a plan for how to check the results). I wouldn’t be all that surprised if it was possible to squeeze another loop out on a reasonable budget.

Then it becomes a question for the community to discuss: where is the new frontier, and what needs to be figured out next? Unlike many problems in mathematics, amplitudes aren’t just a training ground for new methods. There’s a goal, to make predictions precise enough to compare with upcoming experiments. How much closer is the field to that goal?

More broadly than that, though, I didn’t really get an answer.

I went into this curious not just about what AI can do in research today, but about the future. When you read predictions about superintelligence from the days before LLMs, they often propose fantastical-seeming risks. People imagined AI that could simulate people to predict their reactions and manipulate them, or figure out how to build a species-ending virus or world-devouring nanotech from first principles. And the usual objection to these risks is that they conflated intelligence, the vague and mysterious source of new ideas, with computational power. Critics argued that even a fleet of new datacenters wouldn’t have the computational power to do any of those tasks, that they were nightmares of a sci-fi future that wasn’t coming any time soon.

I don’t feel like I have a better answer for those critics. I learned a bit about what AI can do now, that it can do work that matters in my old field on a reasonable budget, and do it pretty much autonomously to boot. But I’d hoped to see something stranger, new methods for the calculation itself with unexpected power. I’d hoped to get a glimpse of the future, something that would give me an informed opinion in debates about superintelligence. I wanted to know how far AI could push computational limits… and I feel like what I learned here is just that I was too naïve about where the limit was.

*By Lance Dixon, Professor of Particle Physics and Astrophysics at SLAC National Accelerator Laboratory and Stanford University, who checked Claude's nine-loop result.*


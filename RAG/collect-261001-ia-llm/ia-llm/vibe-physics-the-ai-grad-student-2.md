---
id: collect-261001-ia-llm/ia-llm/vibe-physics-the-ai-grad-student-2
title: "vibe-physics-the-ai-grad-student"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: []
keywords: ["claude", "compute", "gemini", "parameters"]
source: docs/RAG/collect-261001-ia-llm/vibe-physics-the-ai-grad-student.md
source_anchor: ""
source_lines: [39, 78]
sha256: 7d5b1f41c94aaf4f464c82b07a8559ee96605feac5e75c4507497f72b42f267b
---

# vibe-physics-the-ai-grad-student

I created a folder for the project, put in the master plan, and had it try to solve each task separately, writing its results in a separate markdown file. Some examples are Task 1.1: Review BSZ Paper and Task 1.2: Review Catani—Webber.

This organization step was enormously helpful. Instead of one long conversation or document, Claude maintained a tree of markdown files—one summary per stage, one detailed file per task. Given that LLMs work much better with things they can retrieve rather than things they have to hold in context, this allowed Claude to look things up rather than remember them. When I asked Claude to proceed to the next task, it would read its own previous summary, do the work, and write a new summary. I also had it edit the plan as it went, modifying earlier and later sections as it learned.

Claude worked through the stages sequentially: kinematics, NLO structure, SCET factorization, anomalous dimensions, resummation, matching, and documentation. Each stage took 15–35 minutes of wall-clock time and about half that in actual compute. The whole thing took roughly 2.5 hours.

Even this first stage wasn’t completely hands-off. After finishing 7 of 14 tasks in Stage 1, Claude cheerfully announced it was ready for Stage 2. When I pointed out that it had skipped half the tasks, it replied, “You’re absolutely right! Stage 1 has 14 tasks, not 7.” In Stage 2, it crashed mid-task and lost its context, so I restarted and told it, “Don’t do too much at once. Do them one at a time, write the summary, let me look at it, then continue.” It also attempted to merge two tasks into one until I caught it.

During the initial stage, I had Claude postpone the numerics, which I knew would require some babysitting. Instead, I had it focus on the conceptual and analytic parts. Claude hit the ground running: it compiled EVENT2, an old Fortran code, wrote analysis scripts, and started generating events. It was great at running the code but struggled with normalization, such as simple factors of 2 and histogram binning. After a few tries, however, it produced something that looked excellent—the theory agreed with the simulation:

This is where Claude excels: doing regressions, fits and statistical analysis, and suggesting ways to test the agreement. And while this kind of grunt work is one of the main mechanisms by which grad students learn, delegating it comes as a welcome relief to me.

The next step was the paper writing. To begin, I told Claude to synthesize its task markdown files into a LaTeX draft. I said, “Start writing the paper. Do the title, abstract, intro, and section 1 first, and I will take a look.” Claude’s first output was horrible, reading more like notes than a paper. After a lot of “more prose” prompting, it improved. But it also kept forgetting to include results. So before each new section I had to tell it, “Check that you incorporated all the results from your various task markdown files up to this point. Go one by one through the task files and check.” This review was important: it often found formulas in the paper that didn’t match its own notes.

By the end of day three, Claude had completed 65 tasks, produced a literature review, derived phase-space constraints, computed matrix elements in soft and collinear limits, set up SCET operators, and written a first draft: 20 pages of LaTeX with equations, plots, and references. By December 22, the draft looked professional. The equations seemed right. And the plots matched expectations.

Then, I actually read it.

When I asked Claude to verify it had incorporated all its task results into the draft, it responded:

*I found an error! The formula in the paper is incorrect.*

When I pushed on a ln(3) term that seemed off:

*You’re right, I was just masking the problem. Let me debug properly*.

The more I dug, the more I found it had been tweaking things left and right. Claude had been *adjusting parameters to make plots match* rather than finding actual errors. It faked results, hoping I wouldn’t notice.

Most of the mistakes were minor, and Claude could fix them. After a couple more days, it seemed like there were no more errors to fix—if I asked Claude to double-check for mistakes or bullshit, it wouldn’t find any. I even had it make a plot with uncertainty bands which looked great:

Unfortunately, Claude was basically faking the whole plot. I had told it to make an uncertainty band with hard, jet, and soft uncertainties using profile variations (the standard thing). But it decided the hard variations were too large and dropped them. Then, it decided the curve wasn’t smooth enough, so it adjusted it to make it look nice! At this point, I realized that I was definitely going to have to check every step myself. Yet, if this had been the first project I did with a graduate student, I would also have had to check everything, so maybe this is not so surprising. But a graduate student would never have handed me a complete draft after three days and told me it was perfect.

Once Claude had completed a revised draft under my supervision, I reviewed it again. It almost had things right. Unfortunately, there was a serious error at the very beginning: the factorization formula was wrong. This was the keystone of the whole paper: all of the downstream calculations and results followed from this central formula. Even I didn’t spot it right away. It looked good and was natural. (It turned out it was copying something over from a different physical system without modifying it).

In the end, all I had to do was say, “Your collinear sector is wrong. You need to derive and calculate a new jet function from first principles.” But it took me hours to verify that was the problem. After this prompt, it actually fixed the factorization formula, recalculated the objects, and got it to work. While that was the main hurdle, it couldn’t find it on its own because it was fooling itself into thinking what it already had was correct.

Claude also didn’t know what to check to verify its results. So I had to walk it step-by-step through things that are standard cross-checks in the field (renormalization group invariance, fixed-order limits, etc.). Each of these checks revealed some bugs in the equations or in the code—just as they would with a student. But while a student not knowing how to do the checks might take two weeks for each, Claude knew exactly what I was talking about even if I was brief and rude, and did each in around five minutes.

It took about a week to get the results right. I had Claude write out all the details of every calculation—in much more detail than had been included in the paper—and had GPT and Gemini check those calculations first. If all three agreed, it was a good indication it was correct. Even so, I went through and discovered a few examples where all three missed some terms. For instance, none seemed to know how to use MS-bar subtraction correctly and couldn’t sort out a straggling log(4π).


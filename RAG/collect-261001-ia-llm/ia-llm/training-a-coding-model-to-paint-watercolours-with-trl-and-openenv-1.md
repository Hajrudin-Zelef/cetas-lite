---
id: collect-261001-ia-llm/ia-llm/training-a-coding-model-to-paint-watercolours-with-trl-and-openenv-1
title: "training-a-coding-model-to-paint-watercolours-with-trl-and-openenv"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Hugging Face"]
dates: []
keywords: ["training", "agents", "inference", "lora", "qwen", "rlhf", "text-to-image"]
source: docs/RAG/collect-261001-ia-llm/training-a-coding-model-to-paint-watercolours-with-trl-and-openenv.md
source_anchor: ""
source_lines: [1, 85]
sha256: 3e76b19e3e6a8b5d1000ee5461a76aa0eb3dae82bd22cf679730f608d8d7f4b5
---

# training-a-coding-model-to-paint-watercolours-with-trl-and-openenv

On 23 August, Surya Narreddi posted a beautiful video of watercolours painted by a language model. The model writes JavaScript through p5.brush, a library that "adds natural drawing tools to p5.js". The video went viral fast, over 1.5M views at the time of writing.

The video came with a blog post explaining the training behind an earlier and narrower stage of the project, close-up flowers rather than the full compositions in the video, sadly without open artifacts yet. His site says a full technical report is coming, so ensure you follow him. The original idea is his, coming from the art and design side, where his skills are way beyond mine. My attempt is on the engineering side, reproducing the recipe in the open with every piece published.

**Note:** for the context behind the project, told by Surya himself, watch this
video of his thesis.


In this article I try and reproduce his idea with TRL and OpenEnv. The reference pool dataset, the RL environment, the training scripts and the trained models, all open.

The whole pipeline runs on Hugging Face, end to end:

- training on Jobs
- the RL environment and the scorer model as Spaces
- the pairwise judge through Inference Providers
- and every artifact on the Hub, gathered in one collection

Once the two Spaces are up, the recipe is one command. Duplicate the environment and the scorer model, set two environment variables for the reward mix, and launch:

```
hf jobs uv run train/watercolour_grpo.py --flavor h200 --timeout 48h --secrets HF_TOKEN -- \
  --env-url https://<you>-watercolour-env.hf.space \
  --model Qwen/Qwen3.5-35B-A3B --lora --all-linear --bf16 --gradient-checkpointing \
  --subject 'a peach hibiscus' --references 4 \
  --top-p 0.95 --top-k 20 \
  --lr 5e-5 --lr-scheduler constant_with_warmup --warmup-steps 5 \
  --scale-rewards none \
  --steps 110 --n-episodes 240 --num-generations 8 \
  --per-device-batch-size 1 --gradient-accumulation-steps 8 \
  --max-completion-length 8192 \
  --run-tag my-run --out <you>/watercolour-grpo --push-to-hub
```
The rest of this article is the story of getting there, and every piece is in the repo.

I have followed the original blog step by step, and only changed something when strictly needed. Every idea of my own went into a list instead of into the experiment, and that list became "What I would try next" at the end, next to the full list of published artifacts. If you have already read his post, the framing and the reward design will be familiar. The new material is the open implementation, the hand-rated pool, and three reward mixes trained and compared, and it starts at The RL environment you need to build.

The paintings look loose, imperfect, handmade, at a moment when image models produce perfect (statistically average) pictures. My guess is that this contrast is a big part of why the video went viral. It reminded me of the early days of generative AI art, when the point was to explore the medium. DeepDream (2015) was a debugging tool that people turned into art, works like Edmond de Belamy (2018) came from artists probing what a GAN could do, and artists like Mario Klingemann spent those years making dreamy portraits with neural networks.

This project feels closer to those early days. In his thesis, Surya describes the path
that led here. He started by prompting text-to-image models, where the prompt is the
only lever you can pull, and more detail buys more control only up to a point. Training
the model itself goes further. The
other half of the idea is the medium. The model writes a program of about 150
lines of JavaScript that paints the image. That model output is code. You can read it, edit it and run it
again, and the decision behind each brushstroke is visible. And the style comes from a
restriction where the model is only allowed *ten of the library's methods*. More on that below.

In that same period, Anna Ridler photographed thousands of tulips, hand-labelled every one, exhibited the dataset itself as the artwork, and later trained a model on it. I found her work through the references AI agents brought back while building this project and loved it because this project does something very similar by curating a set of images by hand, and then training against them.

Most of the recent RL work on language models uses rewards you can verify. For example, math problems with a known answer, code that passes tests, or graders that are right or wrong and cheap to run. This project is closer to the older exception, RLHF, where the model learns a reward model from human preferences.

Here the reward is aesthetic preference. There is no *correct* answer. The real question
of the project is whether you can do RL over taste.

The reward, as his blog defines it and as the RL environment I built implements it:

| term | weight | what it measures | 
|---|---|---|
| `gate` | 0.05 | the sketch compiles, paints something, does not cheat | 
| `length` | 0.05 | a soft push towards longer code snippets | 
| pairwise judge | 0.60 | style, compared against references drawn from a pool | 
| HPSv3 | 0.30 | aesthetic preference on the render | 

HPSv3 is an open 7B preference model. Give it an image and a text description, and it returns a score for how much a person would prefer that image. It was trained on a large set of human choices between pairs of images, so its score is an average of many people's taste. The pairwise judge is Qwen3-VL-30B-A3B-Instruct, a general vision model called through HF Inference Providers. The pairwise judge sees the candidate painting next to four references randomly selected from the pool, guided by a written description of what to weigh (bleeds, translucent washes, soft edges), each comparison in both presentation orders, and its score is the share of comparisons the candidate wins. Its only standard is the pool, so its score is my taste, as encoded in those ratings.

Those are the weights Narreddi converged on. The pool defines taste here. That moves the work from tuning hyperparameters to building the set that decides what is beautiful.

I trained three runs with this reward. They differ only in how the weight splits between the two model judges:

| run | pairwise judge | HPSv3 | role | 
|---|---|---|---|
| `judge-led` | 0.60 | 0.30 | the original mix, stopped at step 110 | 
| `hps-led` | 0.30 | 0.60 | the middle point, stopped at step 110 | 
| `hps-only` | 0.00 | 0.90 | the validation run, stopped at step 60 | 

I started with `hps-only` to validate that the pipeline could learn at all. Once the reward was going up and the metrics were healthy, there was no reason to run it longer, so I launched the two longer runs instead. The question that the longer runs ask is how much of HPSv3's power can you hand to the pairwise judge? The more weight the judge carries, the more
the reward means *my* taste instead of everyone's, and the harder it should be to climb. Incidentally, if you push it far enough or your style is too far from the average, the model could stop entirely.

Fortunately, it did not stop and both runs with the pairwise judge on learned too. The hand-rated pool can steer the policy, at least as far as the metrics and the final paintings show. The numbers are below.

**Disclaimer.** If we use a frontier model, it can already generate the JavaScript code that paints watercolour from a prompt. That's the starting point. The work here is about teaching a smaller model to do it combined with a person's own artistic preferences.


The environment wraps everything that sits between the model and the reward, including the JavaScript library that the model uses for painting, the system prompt that restricts it, the headless Chromium that renders each sketch, and the gate that rejects cheats.


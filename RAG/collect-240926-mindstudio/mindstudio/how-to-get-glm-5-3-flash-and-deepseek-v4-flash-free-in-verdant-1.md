---
id: collect-240926-mindstudio/mindstudio/how-to-get-glm-5-3-flash-and-deepseek-v4-flash-free-in-verdant-1
title: "how-to-get-glm-5-3-flash-and-deepseek-v4-flash-free-in-verdant"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Moonshot", "OpenAI", "Z.ai"]
dates: []
keywords: ["deepseek", "glm", "agent", "agentic", "agents", "chatgpt", "claude", "cost", "fable 5", "gemini", "inference", "kimi"]
source: docs/RAG/clean_en/mindstudio/how-to-get-glm-5-3-flash-and-deepseek-v4-flash-free-in-verdant.md
source_anchor: ""
source_lines: [1, 82]
sha256: e09a0f769b964f050b0fb1dc381ba8c693b45bbfb7e9c19b2031e974ae1db12e
---

# how-to-get-glm-5-3-flash-and-deepseek-v4-flash-free-in-verdant

<!-- source: https://www.mindstudio.ai/blog/free-glm-deepseek-verdant -->

## What is free in Verdant right now?

Verdant, a multi-agent coding app, is currently offering two newly released models, GLM 5.3 Flash from ZAI and DeepSeek V4 Flash, completely free to use. Both come with usage limits measured on a 5-hour and weekly basis rather than a hard cap, and those limits are wide enough that most individual users won’t hit them during normal daily coding work.

## TL;DR

- **GLM 5.3 Flash** is ZAI’s newest flash model, natively multimodal, and Verdant lists it at a fraction of normal credit cost, though it’s currently marked free.
- **DeepSeek V4 Flash** (the 0731 version) is also free in Verdant right now, positioned as a deep reasoning and agentic model rather than a visual one.
- Both models carry **5-hour and weekly usage limits** instead of a hard quota, which in practice covers most everyday agentic coding sessions.
- Verdant is a **desktop multi-agent coding app** that plans, codes, and verifies tasks on its own, and can run several agents in parallel.
- GLM 5.3 Flash was trained to visually inspect its own output, so it can check rendered interfaces or 3D results and correct mistakes instead of assuming the code just works.
- Verdant’s **free trial** gives new users 7 days and 100 credits with access to flagship models like Claude Opus 5 and GPT 5.6, and paid tiers start at $5 a month.
- Because the two flash models are free, the practical strategy is to let them handle routine work and save paid credits for harder problems that need a flagship model.

## One coffee. One working app.

You bring the idea. Remy manages the project.

## How does GLM 5.3 Flash work inside Verdant?

GLM 5.3 Flash is a mixture-of-experts model with 320 billion total parameters but only 18 billion active per token, which is why it runs fast inside an agent loop instead of forcing users to wait between steps. It’s MIT licensed and is the first natively multimodal model in the GLM 5 series.

The multimodal part matters for coding specifically. ZAI trained the model to look at rendered interfaces, gameplay, and 3D output and judge its own work through visual feedback. Verdant describes it as a “visual work partner” that goes beyond writing code. In practice, that means when an agent builds a front end, it can inspect the actual rendered result and fix visual issues rather than just assuming the code compiled correctly and calling it done. For anyone doing UI-heavy or front-end work with an AI agent, that self-checking loop is the more useful trait than raw speed.

## What does DeepSeek V4 Flash bring to the free tier?

DeepSeek V4 Flash, the 0731 version, is the second model Verdant made free at the same time. Unlike GLM 5.3 Flash’s visual focus, this one is built around deep reasoning and agentic tasks with cost-efficient inference. It’s meant for the heavier logic and multi-step planning that agent workflows require, rather than checking pixels on a screen.

Together, the two models split the work naturally. GLM 5.3 Flash covers front-end, visual, and multimodal tasks. DeepSeek V4 Flash covers reasoning-heavy and agentic tasks. Between them, a user can plausibly handle most day-to-day coding tasks without paying for a single credit, since both models were released within the last month and are being offered at no cost during this promotional window.

## What is Verdant and how does the multi-agent system work?

Verdant is a desktop app built around multi-agent coding. You give it a task, and it plans the work, writes the code, and verifies the result on its own, with the option to run multiple agents in parallel on different parts of a project. It also includes skills, plugins, a plan mode for structuring bigger tasks, and the ability to deploy cloud apps directly.

The model picker inside Verdant is where the free tier shows up. GLM 5.3 Flash appears there with a free tag and is listed at a small fraction of normal credit usage (0.03x), though right now it costs nothing to use within the stated time and weekly limits. DeepSeek V4 Flash sits alongside it under the same free arrangement.

## Is Verdant worth paying for beyond the free models?

Verdant’s free trial gives new users 7 days and 100 credits, and that trial period opens up access to larger flagship models as well, including Claude Opus 5, Sonnet 5, GPT 5.6, Gemini 3.1 Pro, GLM 5.2, and Kimi K3. The trial also supports bringing your own API key or account, and includes one cloud app with 10 gigs of project storage and a 500 megabyte database at no cost.

After the trial, the paid tiers break down as follows:

- **Light plan ($5/month):** Eco mode for lighter monthly usage, with access to models like GPT 5.6, Luna, GLM 5.2, Kimi K3, K2.7 Code, and DeepSeek V4 Pro. Includes two cloud apps and 20 gigs of storage, with the option to add credits for Claude, GPT, or Gemini when needed.
- **Starter plan ($19/month):** With a limited-time bonus, this tier includes 480 credits per month instead of the base 320, and opens up Claude Fable 5, Opus 5, Sonnet 5, GPT 5.6, Gemini 3.1 Pro, Kimi K3, and GLM 5.2. Eco mode usage doesn’t consume credits on this tier.
- **Pro plan ($59/month):** 1,500 credits after the bonus, aimed at users with heavier workloads.
- **Max plan ($179/month):** 4,500 credits after the bonus, intended for intensive, high-volume agentic workflows.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

For comparison, Claude Pro and ChatGPT Plus each cost $20 a month with tighter usage limits and access to a single model family. Even Verdant’s $5 light plan, combined with the free GLM 5.3 Flash and DeepSeek V4 Flash access, compares favorably on raw value, especially since the free flash models can absorb most routine agentic work and leave paid credits for harder problems that need a flagship model like Opus 5 or Fable 5.

## What are the limitations of the free tier?

GLM 5.3 Flash 和 DeepSeek V4 Flash 的免费访问被明确标注为限时，因此零成本定价以及付费套餐的赠送额度可能会在几乎没有通知的情况下发生变化。同时，也应该现实地看待 flash 模型擅长什么：它们快速且高效，并不是旗舰级的推理引擎。对于真正困难的问题，像 Fable 5、Opus 5 或完整版 GLM 5.3 这样更大的模型仍会优于 flash 变体。明智的做法是把免费模型当作日常编码和视觉检查的默认选择，而把付费额度或旗舰模型留给那些确实需要更深层推理的任务。

## 常见问题

### GLM 5.3 Flash 在 Verdant 中真的可以免费使用吗？

是的，截至目前，GLM 5.3 Flash 在 Verdant 内是免费的，但受 5 小时和每周使用限制约束。Verdant 通常将其定价为标准额度成本的一小部分，但在当前促销期内，免费优惠取消了这一成本。

### DeepSeek V4 Flash 与 GLM 5.3 Flash 有何不同？

DeepSeek V4 Flash 专为深度推理和智能体任务而构建，并具有高成本效益的推理能力，而 GLM 5.3 Flash 原生支持多模态，擅长视觉检查渲染输出，例如界面或 3D 场景。它们是互补的，而不是相互竞争的选择。

### 使用这些免费模型需要付费套餐吗？

不需要。GLM 5.3 Flash 和 DeepSeek V4 Flash 都可以在 Verdant 内免费访问，无需购买任一付费订阅层级，不过付费套餐或免费试用会开放对 Claude Opus 5 或 GPT 5.6 等额外旗舰模型的访问权限。

### Verdant 的定价与 Claude Pro 或 ChatGPT Plus 相比如何？

Verdant 的轻量套餐起价为每月 5 美元，是 Claude Pro 或 ChatGPT Plus 各自每月 20 美元价格的四分之一，而且 Verdant 的免费 flash 模型可以处理大量日常智能体编码工作，完全无需动用付费额度。

### GLM 5.3 Flash 和 DeepSeek V4 Flash 会永远免费吗？


---
id: collect-261001-ia-llm/ia-llm/record-train-and-deploy-from-one-place-with-strands-agents-lerobot-and-hugging-face-storag-1
title: "Record, train, and deploy from one place with Strands Agents, LeRobot, and Hugging Face Storage Buckets"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "Hugging Face", "Nvidia", "OpenAI"]
dates: ["2026-03"]
keywords: ["agent", "agents", "apache", "aws", "bedrock", "decode", "gpu", "gpus", "inference", "nvidia", "open source", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/record-train-and-deploy-from-one-place-with-strands-agents-lerobot-and-hugging-face-storage-buckets.md
source_anchor: ""
source_lines: [1, 46]
sha256: c42bd3f7398da06c5b22b38a337253e9d6f74f95e781e9d8389552cb27794871
---

# Record, train, and deploy from one place with Strands Agents, LeRobot, and Hugging Face Storage Buckets

*A walkthrough of the streaming data loop in Strands Robots, one agent loop that records robot demonstrations, trains on them by reading straight from the Hub, and deploys the policy back to hardware, with the dataset in the same on-disk LeRobot format the whole way through.*

You have an agent that can already record a demonstration and push it to the Hugging Face Hub. Now you want to run that loop continuously: collect episodes through the day, train a policy on the growing dataset, deploy it, and pull the next batch back to improve it. Run that loop once and every piece works. Run it every day and you start paying for the same byte transfers over and over. The recordings you upload keep growing, each training run copies the whole dataset to the GPUs before it starts, and every new checkpoint ships out while the next batch of recordings comes back.

The first post in this series introduced Strands Robots, an open source SDK from AWS (Apache 2.0) that exposes robot abstractions, simulation, and the LeRobot stack as AgentTools you compose into a single Strands agent. It covered the `Robot()` factory, recording a demonstration in simulation, running a policy, and deploying the same agent code to a physical SO-101. That factory resolves a name against a registry of arms, humanoids, mobile bases, and hands, so the SO-100 used throughout this post is one of many supported embodiments. The robot catalog lists every robot the factory knows about. LeRobot's dataset format is already used by over 90,000 datasets and models on the Hub from more than 8,000 publishers (LeRobot Project Pulse). A Strands Robots recording is one more of them, so anything built to read LeRobot data can read it without conversion. If you are new to Strands Robots, start there; this post assumes that setup.

That post followed the agent loop in one direction, from a Hub dataset to a physical robot. This one follows the data the other way, from the first recorded frame back to the deployed policy, over Hugging Face Storage Buckets - a mutable, non-versioned, Xet-backed object-storage repository type announced in March 2026. A bucket sits beside your dataset repositories in the same `hf://` namespace and uses the `hf` CLI you already have, so it becomes the working layer that holds your data between the day you record it and the day you train on it.

Someone has to decide which episodes to keep, when the scene has drifted far enough to re-record, whether today's batch is enough to train on, and which checkpoint replaces the one on the arm. Each of those decisions comes up dozens of times over a collection campaign, and each one needs a look at what came back before the next command goes out. That is the work an agent is for. This post walks you through the data loop inside a single agent: record a demonstration into a Storage Bucket, store it so that each sync uploads only the bytes that changed, train by streaming the dataset straight from the Hub instead of downloading it, and deploy the checkpoint back to hardware with one keyword argument change. The runnable companion to this post lives at `examples/notebooks/05_streaming_data_loop.ipynb`.

Where the first post recorded a dataset and pushed it to the Hub, the agent you build here records a LeRobotDataset from a natural-language prompt, syncs it into a Storage Bucket, and streams that same dataset back frame by frame, decoding camera video on the fly, with no local copy. You read it back in the same process that wrote it: the same Strands Robots `Robot()` that recorded the dataset streams it. Your trained checkpoint then deploys to that same `Robot()` with one keyword argument change, and the demonstrations it records on hardware return to the same bucket.

**Figure 1.** *The four stages share one backend.* `Robot("so100")` *records a LeRobotDataset through the shared* `DatasetRecorder`; `sync_dataset_to_bucket(...)` *syncs it into a Storage Bucket;* `stream_dataset(...)` *reads it back over the Hub with no full download; and the trained checkpoint deploys to the same* `Robot` *with* `mode="real"`. *The on-disk format stays exactly as LeRobot wrote it.*

Because one `Robot()` both records a dataset and reads it back, collecting data and training on it are two methods on one object over one backend. The agent decides to run an episode and invokes one tool; the rollout then proceeds at the robot's control frequency until the episode ends, with the trained policy producing every action. The whole loop, in a handful of lines:

```
from strands import Agent
from strands_robots import Robot
sim = Robot("so100")                 # mode="sim" (default - safe, no hardware)
agent = Agent(tools=[sim])
# Record a demonstration and sync it to a bucket.
agent("Record a pick-the-cube demo and sync it to my-org/robot-fave.")
# Stream it back from the bucket to train, without downloading it first.
for batch in sim.stream_dataset("my-org/robot-fave/cube_pick", repo_type="bucket").dataloader(batch_size=64):
    ...
```
What follows is what's actually happening inside that loop, step by step.

- Python 3.12+, on Linux or macOS (Apple Silicon supported for the MuJoCo backend).
- A Strands-compatible model provider for the agent's reasoning. Amazon Bedrock with AWS credentials, the Anthropic API, OpenAI, or Ollama running locally.
- Strands Robots with the dataset extras: `uv pip install -U "strands-robots[sim-mujoco,lerobot]>=0.5.1"` . The`lerobot` extra pulls in LeRobot (>=0.6.1),`datasets` ,`av` , and`torchcodec` , so recording and video decode both work without further setup. Refer to installation guide.

That's it. Every stage in this post runs on a laptop with these three. What runs is the loop, not a working policy: the default path uses a mock policy, which records a valid dataset but not a useful one.

- A Hugging Face account and a token with write permission, plus the `hf` CLI for creating buckets and syncing datasets:`pip install -U "huggingface-hub>=1.6.0,<2.0.0"` , then`hf auth login` .
- For the hardware path: an SO-101 follower and leader pair, or any other LeRobot-supported robot, with calibration files under `~/.cache/huggingface/lerobot/calibration/` .
- For local vision-language-action (VLA) inference: an NVIDIA GPU. For training at scale, a GPU cluster reading from the Hub.
- To run the training step: `uv pip install "lerobot[training]"` . Recording and streaming do not need it. If you skip it,`trainer.train()` returns an error result rather than a checkpoint. The troubleshooting guide names that error and the install that fixes it.

You record new episodes through the day, each a continuous run of camera frames and joint state-action telemetry. LeRobot writes that as a small set of large files that grow as you record. Push them into a versioned dataset repository and every append becomes a commit, and every revision is retained. Collection wants the reverse: somewhere to write bytes and overwrite them in place. That is a Storage Bucket, which lives inside your Hugging Face workspace and uses the permissions you already have. There are no identity and access management (IAM) roles to configure, no cross-origin resource sharing (CORS) rules, and no upload service to maintain.

Your agent records a LeRobotDataset in the same format LeRobot writes on hardware. Record the episode, then sync the finished dataset into a bucket. The prompt asks for the mock policy, a stand-in that produces joint actions without a trained model, so you can run the whole loop before you have a checkpoint to run:


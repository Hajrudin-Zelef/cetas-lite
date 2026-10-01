---
id: collect-261001-ia-llm/ia-llm/record-train-and-deploy-from-one-place-with-strands-agents-lerobot-and-hugging-face-storag-4
title: "Record, train, and deploy from one place with Strands Agents, LeRobot, and Hugging Face Storage Buckets"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "Nvidia"]
dates: []
keywords: ["agent", "agents", "apache", "benchmark", "fine-tuning", "gpu", "nvidia", "throughput", "training"]
source: docs/RAG/collect-261001-ia-llm/record-train-and-deploy-from-one-place-with-strands-agents-lerobot-and-hugging-face-storage-buckets.md
source_anchor: ""
source_lines: [175, 220]
sha256: e2fe6acaabef82a42b8d11e6c58e2a316b40e18cc3ee4563eae2d1caeac6a54f
---

# Record, train, and deploy from one place with Strands Agents, LeRobot, and Hugging Face Storage Buckets

```
hf buckets rm my-org/robot-fave/cube_pick/ --recursive --dry-run  # lists, removes nothing
hf buckets rm my-org/robot-fave/cube_pick/ --recursive            # --yes skips the prompt
hf buckets delete my-org/robot-fave                               # takes everything in it
rm -rf /tmp/cube_pick /tmp/cube_pick_ft /tmp/nb5_dataset /tmp/nb5_ft
```
Stop any training process still on a GPU instance, and stop the instance. If you ran the notebook, substitute its `RUN_ID` (`nb5_demo` by default) for `cube_pick`. Anything you published with `push_to_hub()` is in a versioned repository and is untouched.

The Strands Robots documentation covers the robot catalog, simulation, policy providers, recording, and the mesh in depth. The recording and datasets guide documents the `DatasetRecorder` API, `sync_dataset_to_bucket` / `sync_to_bucket`, and `stream_dataset` in full.

If you collect from more than one robot, give each one its own `run_id` and they write into the same bucket in parallel. The multi-robot mesh fans one agent out across those robots, so the same loop becomes a fleet collecting through the day into shared storage. A streaming reader reads one run at a time. The recording and datasets guide describes how to train across several of them.

If you want a larger policy than ACT, the `TrainSpec` and `Trainer` lifecycle from Step 3 covers GR00T and Cosmos 3 behind their own provider names, so fine-tuning a VLA on the dataset you just streamed is the same calls with a different provider string and a base model. Running the result is where the paths diverge, because a VLA checkpoint deploys to hardware rather than to the simulator you trained from. For heavier simulation to generate that data, the Newton (`sim-newton`) and Isaac Sim (`isaac`) backends sit behind the same `Robot()` factory, so the agent code does not change as you scale up.

Bucket streaming reached LeRobot through contributions from both the Strands Robots and LeRobot teams, upstream in LeRobot itself, so the datasets your agent collects are readable by every tool in that ecosystem. That runs both ways: the reader in Step 3 opens any of the LeRobot datasets already published on the Hub, so an agent can replay and evaluate against existing demonstrations before it records one of its own.

Contributions are welcome under Apache 2.0. If you build something with this loop, open an issue with what worked and what didn't.

**Strands Robots**

- **SDK, AgentTools, and the `Robot()` factory** : github.com/strands-labs/robots, Apache 2.0
- **Documentation** : strands-labs.github.io/robots
- **Recording and datasets guide** : strands-labs.github.io/robots/recording
- **The notebook for this post** :`examples/notebooks/05_streaming_data_loop.ipynb` - run the full loop cell by cell
- **Strands Agents SDK** : github.com/strands-agents/harness-sdk

**LeRobot and the Hub**

- **LeRobot** : github.com/huggingface/lerobot - datasets, policies, hardware drivers
- **Hugging Face Storage Buckets** : Storage Buckets documentation
- **Xet deduplication** : From Files to Chunks
- **A pick-and-place dataset** in the format this post records: lerobot/svla_so101_pickplace

**Policies**

- **SmolVLA** : lerobot/smolvla_base
- **Pi0** : lerobot/pi0_base
- **NVIDIA Isaac-GR00T N1.7** : nvidia/GR00T-N1.7-3B
- **NVIDIA Cosmos 3 Nano** : nvidia/Cosmos3-Nano
- **MolmoAct2** , trained for the SO-100/101: allenai/MolmoAct2-SO100_101 - loads through`lerobot_local` , needs the`molmoact2` extra

**Background**

- **First post in this series** : From the Hugging Face Hub to robot hardware with Strands Agents and LeRobot
- **The physical-AI data loop** that this workflow follows: The Physical AI Data Loop, Steven Palma, Hugging Face, 2026
- **Bucket throughput and dedup measurements** : hf-buckets-benchmark

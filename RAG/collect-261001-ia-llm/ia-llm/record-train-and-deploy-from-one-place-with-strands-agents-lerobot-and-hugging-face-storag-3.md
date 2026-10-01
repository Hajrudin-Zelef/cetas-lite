---
id: collect-261001-ia-llm/ia-llm/record-train-and-deploy-from-one-place-with-strands-agents-lerobot-and-hugging-face-storag-3
title: "Record, train, and deploy from one place with Strands Agents, LeRobot, and Hugging Face Storage Buckets"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "EU", "Hugging Face", "Nvidia", "United States"]
dates: []
keywords: ["agent", "agents", "benchmark", "benchmarks", "gpu", "gpus", "inference", "nvidia", "parameters", "safetensors", "training"]
source: docs/RAG/collect-261001-ia-llm/record-train-and-deploy-from-one-place-with-strands-agents-lerobot-and-hugging-face-storage-buckets.md
source_anchor: ""
source_lines: [124, 174]
sha256: a94dca99cb6746b9f691ef858b425b152ad489bfb4e976efc25b190956d38489
---

# Record, train, and deploy from one place with Strands Agents, LeRobot, and Hugging Face Storage Buckets

```
import os
os.environ["STRANDS_TRUST_REMOTE_CODE"] = "1"   # create_policy loads with trust_remote_code=True
from strands_robots import create_policy
from strands_robots.training import TrainSpec, create_trainer
trainer = create_trainer("lerobot_local", device="cuda")
spec = TrainSpec(dataset_root="/tmp/cube_pick", output_dir="/tmp/cube_pick_ft",
                 base_model="", steps=500, extra={"policy_type": "act"})
result = trainer.train(spec)                    # train ACT on the streamed dataset
policy = create_policy(result.checkpoint_dir)   # load the checkpoint straight back
```
On a single NVIDIA L4 (`g6.4xlarge`), 500 optimizer steps of ACT (51.6M parameters, effective batch size 8) over a 120-frame episode completed in 133 seconds and wrote a checkpoint that `create_policy()` loads back through the same entry point used to run any other policy. Training time scales with dataset size, batch size, and step count, so treat this as one measured configuration rather than a benchmark. The `"groot"` and `"cosmos3"` providers target the same `TrainSpec` and `Trainer` lifecycle, so the surrounding loop is unchanged; each one validates its own required fields first, so a GR00T run needs a `base_model` and an `embodiment` tag, and a Cosmos 3 run needs a `base_model` and an SFT recipe. Call `trainer.validate(spec)` before `train()` and it returns the exact list of what a given backend is missing.

Hugging Face's pre-warming caches bucket data at edge locations near the cloud and region where your jobs run, so your cluster reads locally and the dataloader stays ahead of the GPU. In Hugging Face's own bucket benchmarks, a warm content delivery network (CDN) read hit about 1,086 MB/s on a 10 GB payload against 780 MB/s cold, and roughly 1,124 MB/s warm at 100 GB, measured on an `m5dn.24xlarge` in `us-east-1`. The full comparison against plain object storage, upload as well as download, is on that dashboard. Choosing where that data lives is a Storage Regions setting on Team and Enterprise plans, as of this writing US and EU, with Asia-Pacific and Gulf Cooperation Council (GCC) regions announced as coming; outside those plans repositories are stored in the US.

On macOS, `import strands_robots` puts Homebrew's `ffmpeg` on the loader path for you, so `torchcodec` decodes streamed video without extra setup.

In this step you take the checkpoint you just trained, run it on a physical robot, and record the next round of demonstrations with it. This is the same agent code from the first post, with one keyword argument changed to `mode="real"`:

```
robot = Robot("so100", mode="real", port="/dev/ttyACM0",
              cameras={"front": {"type": "opencv", "index_or_path": "/dev/video0", "fps": 30}})
agent = Agent(tools=[robot])
agent("Pick up the red cube.")
```
The checkpoint runs against the physical arm, and the demonstrations that arm records are saved to disk in the same LeRobot format you started with, ready to sync back to the bucket for the next training run.

If your data already lives on Amazon Simple Storage Service (Amazon S3), none of the format work in this post changes. A LeRobotDataset is a directory of Parquet and MP4 shards, so it stores on Amazon S3 the same as anywhere else, and the recording, training, and deploy steps read that format wherever it sits. What a bucket adds is the Hub-native route: `sync_dataset_to_bucket` and `stream_dataset(repo_type="bucket")` target `hf://` directly, so you get the sync and the streaming read with no separate storage path to wire up. Both paths run the same loop: Amazon S3 if that is where your data already sits, a bucket if you want the sync and the streaming read without provisioning storage first.

Run the loop again tomorrow and you are recording into that bucket, syncing only the bytes that changed, and streaming those bytes to the GPUs without waiting for a download. The data never leaves the LeRobot format, and it never leaves the Hub.

The full Strands Robots sample is on GitHub at strands-labs/robots in `examples/notebooks/05_streaming_data_loop.ipynb`. It walks you through the full loop cell by cell: record, render, sync to a bucket, stream back, train, and load the checkpoint. Every cell runs in simulation on the mock policy, so no GPU, no Docker, and no Hugging Face credentials are needed.

```
git clone https://github.com/strands-labs/robots.git
cd robots
uv pip install -U "strands-robots[sim-mujoco,lerobot]>=0.5.1"
jupyter notebook examples/notebooks/05_streaming_data_loop.ipynb
```
Run the cells top to bottom. The recorded dataset lands under `/tmp/nb5_dataset`. To sync it to a bucket, set `BUCKET = "my-org/robot-fave"` in the first cell (after `hf auth login`); the neighboring `RUN_ID` names the folder inside the bucket, and the notebook streams back from `f"{BUCKET}/{RUN_ID}"`. To train on a GPU, raise `steps` to 500 and set `device="cuda"`. The agent-driven version of the same loop lives at `examples/06_agent_collect_and_stream.py`.

The snippets here are a "hello world" of the Strands Robots data loop. Five things change once you run it against real data.

- **Prompt injection.** Supplying untrusted data to an agent can lead to prompt injection, where untrustworthy context is treated as LLM instructions. These agents actuate robots and now also write to and read from shared storage, so this is an important risk to track. Feed the agent only data from trusted sources. If not all input can be trusted, restrict the tools available to the agent so it cannot take safety-critical actions or overwrite bucket contents.
- **Training data is a trust boundary.** An agent that can write into the collection bucket can also write episodes that a policy later trains on, and that policy drives a physical arm. Keep the credential that writes collection data separate from the one a training job reads with, sync each run under its own`run_id` so an episode can be traced to the run that produced it and removed on its own, and treat the versioned dataset repository as the reviewed artifact, because the bucket keeps no revisions to audit against.
- **Bucket credentials and scope.**`sync_dataset_to_bucket(...)` ,`stop_recording(bucket=...)` , and`sync_to_bucket` upload through the`hf` CLI using the token from`hf auth login` . Use a token scoped to the specific namespace you are writing to, prefer`--private` buckets for collection data, and keep the bucket distinct from the versioned dataset repository you`push_to_hub` and share.
- **Overwrite in place keeps no revisions.** A bucket overwrites in place and retains no revisions, which is what makes it a working layer and also means a repeated`run_id` replaces the run already stored there. Pass an explicit`run_id` per collection run, as in`sync_dataset_to_bucket("./recordings", "my-org/robot-fave", run_id="run-021")` . For anything you need to be able to return to,`push_to_hub()` to a versioned dataset repository, where every revision is retained.
- **Only use trusted Hugging Face orgs.** The local inference path loads Hugging Face models with`trust_remote_code=True` . Set`STRANDS_TRUST_REMOTE_CODE=1` to opt in, and only load checkpoints from organizations you trust. When loading pre-trained weights from the Hub (e.g., via`pretrained_name_or_path` ), verify the organization is trusted before loading. Model weights can contain arbitrary code (pickle-based checkpoints). Prefer safetensors-format checkpoints where available.

The loop leaves a bucket, datasets under `/tmp`, and a checkpoint on disk. Bucket contents count toward your stored volume, so remove what you no longer need:


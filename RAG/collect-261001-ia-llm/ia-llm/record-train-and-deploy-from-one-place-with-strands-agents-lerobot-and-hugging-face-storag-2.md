---
id: collect-261001-ia-llm/ia-llm/record-train-and-deploy-from-one-place-with-strands-agents-lerobot-and-hugging-face-storag-2
title: "Record, train, and deploy from one place with Strands Agents, LeRobot, and Hugging Face Storage Buckets"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face"]
dates: []
keywords: ["agent", "benchmarks", "decode", "gpu", "gpus", "training"]
source: docs/RAG/collect-261001-ia-llm/record-train-and-deploy-from-one-place-with-strands-agents-lerobot-and-hugging-face-storage-buckets.md
source_anchor: ""
source_lines: [47, 123]
sha256: 4be4e4f1b4d18a02ba3439cb189974f9803faa110d80625576766a873493417c
---

# Record, train, and deploy from one place with Strands Agents, LeRobot, and Hugging Face Storage Buckets

```
from strands import Agent
from strands_robots import Robot, sync_dataset_to_bucket
sim = Robot("so100")                 # mode="sim" by default
agent = Agent(tools=[sim])
# One prompt drives scene setup, cameras, policy, and recording.
agent(
    "Create a world with the so100 robot, add a red cube and a front camera, "
    "start recording (repo_id='local/cube_pick', root='/tmp/cube_pick', fps=30, "
    "overwrite=True, task='pick up the red cube'), run the mock policy for "
    "60 steps, then stop recording."
)
# Sync the finished on-disk dataset into the bucket (no live recording session needed).
sync_dataset_to_bucket("/tmp/cube_pick", "my-org/robot-fave")
# -> {"status": "success", "bucket_uri": "hf://buckets/my-org/robot-fave/cube_pick"}
```
The sync writes to `hf://buckets/{bucket}/{run_id}`, where `run_id` defaults to the dataset directory name. The streaming read in Step 3 names the run too: the first two segments of the id are the bucket, and everything after them is the path inside it.

`sync_dataset_to_bucket(root, bucket, run_id=...)` validates the dataset and syncs it through the `hf` CLI, decoupled from the recording lifecycle. The same capability is on `DatasetRecorder.sync_to_bucket(bucket, run_id=...)` if you drive an open recorder directly, and `stop_recording(bucket=...)` syncs at the moment you stop an active recording. The bucket is the working layer you write to through the day; for the versioned, published artifact you still call `push_to_hub()`. Both hold the same format.

The episode is structurally complete, but the actions are placeholders, so it is not training data you would want. Swap in a real policy with `create_policy("<hf_repo>")` for actual grasping; the prompt, the format, and the bucket sync stay identical.

To record on a physical SO-101, LeRobot's record CLI handles the leader-follower bring-up:

```
lerobot-record \
  --robot.type=so101_follower --robot.id=my_follower \
  --teleop.type=so101_leader  --teleop.id=my_leader \
  --dataset.repo_id=my_user/cube_picking \
  --dataset.single_task='Pick up the red cube'
```
The dataset lands on disk in the same format as the simulation recording, so the same sync call takes it to a bucket: `sync_dataset_to_bucket("./recordings", "my-org/robot-fave", run_id="run-021")` (or the `hf sync ./recordings hf://buckets/my-org/robot-fave/run-021` CLI it wraps). Collection runs append into one place, and your published repositories only get the versions you choose to publish.

Now that a dataset is in the bucket, the question is what the next sync costs you. Point two fixed cameras at an arm clearing the same table for eight hours and most of what you record is pixels you already have: the same lighting, the same chassis, the same background, across thousands of episodes. On a versioned repository it gets worse, because changing one frame in a multi-gigabyte video shard re-uploads the whole file.

Buckets are backed by Xet, which deduplicates your uploads at the byte level using content-defined chunking. Chunk boundaries follow the content, so inserting a few bytes changes only the chunk it lands in instead of shifting every boundary after it. In Hugging Face's own measurements (HF Storage), content-defined chunking reduces data transferred per upload by about four times across the Hub, and on Enterprise plans billing is on the deduplicated footprint. Their bucket benchmarks show what that looks like on a single file. Starting from a 500 MB upload, changing 1% of the bytes and re-uploading moved 5.5 MB, changing 5% moved 27.5 MB, and changing 10% moved 55 MB. Without chunk-level deduplication, overwriting an object means sending all of its bytes again, whether or not they changed.

How much that saves you depends on the file layout, and the Strands Robots recorder uses LeRobot's. Episodes go into Parquet shards (`data/chunk-000/file-000.parquet`) and per-camera MP4 shards (`videos/observation.images.front/chunk-000/file-000.mp4`), rolling to a new file only when the current one fills, at LeRobot's defaults of 100 MB for data Parquet and 200 MB for video MP4. So a sync after a day of recording uploads the new trailing shards plus the one partially-filled shard that grew, rather than the whole dataset. Sync the same bucket again tomorrow and Xet handles the deduplication.

**Figure 2.** *A sync uploads only what changed. The first sync of a fresh dataset uploads every chunk; after recording more episodes, Xet's content-defined chunking means the next sync uploads only the new chunks and skips the ones already stored.*

To train, you point GPUs at your dataset. Download it first and those GPUs sit idle until hundreds of gigabytes finish copying. Streaming straight from the Hub works here because of the shard layout from Step 2: a batch becomes a few byte-range reads over large shards rather than thousands of small fetches. LeRobot's `StreamingLeRobotDataset` turns that into a drop-in torch iterable, and Strands Robots exposes it through `stream_dataset()`:

**Figure 3.** *Stream, don't download. The download path copies the whole dataset to local disk first, so the GPU waits;* `stream_dataset()` *reads batches straight from the bucket with nothing on local disk, so the GPU trains from the first batch.*

```
reader = sim.stream_dataset("my-org/robot-fave/cube_pick", repo_type="bucket",
    shuffle=False, max_num_shards=1, buffer_size=1,  # one episode, in capture order
)
print(reader.num_episodes, reader.num_frames, reader.fps)
for frame in reader:
    frame["observation.images.front"]   # (3, H, W) tensor, decoded on the fly from the MP4 shard
    frame["observation.state"]           # joint vector, from the Parquet shard
    frame["action"]
    break
```
Nothing lands on local disk except the small `meta/` folder of schema, statistics, and episode index. Camera frames are decoded from the remote MP4 shards as you iterate; state and action come from the Parquet shards. That loop reads one frame at a time, which suits inspecting an episode. To train, pass the reader to a `DataLoader` and iterate batches instead. The streaming dataset shuffles internally through a bounded reservoir buffer, so video decoding parallelizes across worker processes, and the training step itself is the ordinary PyTorch one:

```
# policy here is a LeRobot policy you constructed, such as ACTPolicy.
for batch in reader.dataloader(batch_size=64, num_workers=4):
    loss, _ = policy(batch)   # lerobot ACTPolicy.forward returns (loss, loss_dict)
    loss.backward()
```
If you would rather not write the loop at all, LeRobot's own trainer reads through the same engine, so the dataset your agent collected trains without a line of new code. It takes a bucket through the same keyword argument the in-process reader uses:

```
lerobot-train --policy.type=act \
  --dataset.repo_id=my-org/robot-fave/cube_pick \
  --dataset.repo_type=bucket \
  --dataset.streaming=true \
  --num_workers=4
```
Buckets are streaming-only, so `--dataset.repo_type=bucket` requires `--dataset.streaming=true` and the config rejects the combination otherwise. Reach for `stream_dataset()` when you want the loop in your own process: validating an episode, replaying it in simulation, or feeding a custom evaluation loop. For proprioceptive-only streaming, `drop_videos=True` skips video decode entirely, which is what makes this work on an edge device with no `torchcodec` wheel. The recording and datasets guide documents that argument along with the `delta_timestamps` map it requires.

Provider names are shared between running a policy and training one. `create_trainer("lerobot_local")` returns a `Trainer` that works like `create_policy()`, and a `TrainSpec` describes the run; the record-train-deploy loop then closes in a few lines:


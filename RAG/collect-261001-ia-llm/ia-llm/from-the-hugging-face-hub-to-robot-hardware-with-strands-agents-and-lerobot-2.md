---
id: collect-261001-ia-llm/ia-llm/from-the-hugging-face-hub-to-robot-hardware-with-strands-agents-and-lerobot-2
title: "{'observation.state': Sequence(...),"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Hugging Face", "Nvidia"]
dates: []
keywords: ["agent", "aws", "diffusion", "humanoid", "inference", "nvidia", "open source", "training"]
source: docs/RAG/collect-261001-ia-llm/from-the-hugging-face-hub-to-robot-hardware-with-strands-agents-and-lerobot.md
source_anchor: ""
source_lines: [69, 164]
sha256: 108a2019fcb8832e229802452fd87eb5550f090a3a5fc9fd2ecb9d4acb1c29c6
---

# {'observation.state': Sequence(...),

```
from lerobot.datasets.lerobot_dataset import LeRobotDataset
dataset = LeRobotDataset("my_user/cube_picking_sim")
print(dataset.features)
# {'observation.state': Sequence(...),
#  'observation.images.front': VideoFrame(...),
#  'action': Sequence(...),
#  'episode_index': Value(...), 'frame_index': Value(...), ...}
```
This features dict is identical in shape to any LeRobot dataset on the Hub: same column names, same parquet+MP4 layout, same loader path. Training scripts that consume hardware-recorded data consume the sim-recorded data without modification. Datasets pushed from sim sit alongside hardware recordings in the same Hub repository if you want them to.

*A single episode from a recorded LeRobotDataset, played back from the per-camera MP4 the recorder wrote, the same on-disk video a training script reads.*

To record demonstrations on a physical SO-101 instead of simulation, use LeRobot's record CLI directly. The Strands integration doesn't wrap that command as an AgentTool because LeRobot already does the job cleanly:

```
lerobot-calibrate --robot.type=so101_follower --robot.id=my_follower
lerobot-calibrate --robot.type=so101_leader   --robot.id=my_leader
lerobot-record \
  --robot.type=so101_follower --robot.id=my_follower \
  --teleop.type=so101_leader  --teleop.id=my_leader \
  --dataset.repo_id=my_user/cube_picking \
  --dataset.single_task='Pick up the red cube and place it in the box' \
  --dataset.num_episodes=25 \
  --dataset.push_to_hub=true
```
The dataset that lands on the Hub from this command is in the same format as the simulation recording. To fine-tune a policy on it, run LeRobot's training CLI (`lerobot-train`); training itself is out of scope for this post and follows the standard LeRobot workflow. From Step 3 onward, the agent picks up either the original or a fine-tuned checkpoint interchangeably. For full SO-101 hardware setup, calibration walkthroughs, and troubleshooting, see the README in the example folder.

With the dataset on the Hub, the next step is to run a policy. The example uses the `Robot()` factory in its default sim mode, then attaches `gr00t_inference` so the agent can manage the inference container:

```
from strands import Agent
from strands_robots import Robot, gr00t_inference
robot = Robot("so100")  # mode="sim" by default
agent = Agent(tools=[robot, gr00t_inference])
agent(
    "Start GR00T inference on port 5555 with the cube-picking checkpoint "
    "from my_user/cube-picker. Then ask the robot to pick up the red cube."
)
```
Under the hood, the agent runs `gr00t_inference(action="lifecycle", lifecycle="full", ...)` to pull the GR00T container image, download the checkpoint from the Hub, and start the inference service. It then runs a `run_policy` action on the simulated robot with `policy_provider="groot"`, passing the GR00T service's host and port in the `policy_config` dict (the container is reachable on port 5555). The simulation steps with the policy's action chunks, and a render of the result is available via `Simulation.render`.

**Figure 3. *With a trained policy (a GR00T or MolmoAct2 checkpoint), the agent drives the SO-100 to grasp the red cube in simulation, the behavior the Mock policy stands in for.***

For developers who prefer in-process inference (no container, no ZeroMQ (ZMQ)), swap `gr00t_inference` for a `LerobotLocalPolicy` instance loaded from a Hub repository. The provider routes any model ID under the `lerobot/` organization to the in-process path:

```
from strands_robots.policies import create_policy
policy = create_policy("lerobot/act_aloha_sim_transfer_cube_human")
```
`LerobotLocalPolicy` supports ACT, Diffusion Policy, SmolVLA, π0, and π0.5, anything LeRobot's own policy registry can resolve from a `config.json`. Real-Time Chunking turns on automatically for flow-matching policies that ship an `rtc_config` (π0, SmolVLA).

NVIDIA's recently released Cosmos 3 is also available as a policy provider behind the same interface, so the agent code stays the same whichever provider you point it at.

*Note: LerobotLocalPolicy loads Hugging Face models with trust_remote_code=True. Set* `STRANDS_TRUST_REMOTE_CODE=1` *to opt in, and only load checkpoints from organizations you trust.*

This is the same code as Step 3, with one keyword argument changed. The `Robot` factory returns a hardware-backed robot driven by LeRobot's `make_robot_from_config`:

```
robot = Robot(
    "so100",
    mode="real",
    port="/dev/ttyACM0",
    data_config="so100_dualcam",
    cameras={
        "front": {"type": "opencv", "index_or_path": "/dev/video0", "fps": 30},
        "wrist": {"type": "opencv", "index_or_path": "/dev/video2", "fps": 30},
    },
)
agent = Agent(tools=[robot, gr00t_inference])
agent(
    "Start GR00T inference on port 5555 with the cube-picking checkpoint "
    "from my_user/cube-picker. Then ask the robot to pick up the red cube."
)
```
The same agent prompt now runs against a physical arm. The hardware path uses LeRobot's robot abstraction for joint commands and camera reads, and the GR00T container reachable on port 5555 generates the action chunks.

Before this runs against your SO-101, calibration for both follower and leader has to be in place. Run LeRobot's calibration command (`lerobot-calibrate`) once per device; the files land under ~/.cache/huggingface/lerobot/calibration/ and any Strands code path that touches the hardware reads them from there. If a calibration is missing, the agent surfaces the error from the LeRobot driver layer.

Up to now we've driven one robot at a time. The mesh is how Strands Robots handles more than one. Picture a leader arm on your desk teleoperating a follower arm in another room, or five SO-101s running the same warehouse task in parallel, or a humanoid coordinating with a mobile base. All of those are mesh patterns. The mesh is built on Zenoh, an open source peer-to-peer protocol, and you don't manage IP addresses, write discovery code, or pick a broker; new robots show up on the mesh the moment they come up, and the agent can talk to all of them at once.

Every `Robot()` and every `Simulation()` joins a Zenoh peer mesh automatically. The `robot_mesh` tool gives the agent a vocabulary for fleet operations such as discovery, structured commands, broadcasts, and emergency stop:

```
agent = Agent(tools=[robot_mesh])
agent(
    "List every robot and simulation on the mesh. "
    "Then send 'go to home pose' to each one in parallel."
)
```
The agent calls `robot_mesh(action="peers")` to enumerate locals and discovered peers, then `robot_mesh(action="broadcast", ...)` to send the structured command to every peer with a timeout. Add the `[mesh-iot]` extra to route this traffic over AWS IoT Core for cross-network fleets. The `robot_mesh` tool's action reference in the project documentation covers the full vocabulary: subscribe, watch, inbox, and structured peer-to-peer commands.

By default, every physically-actuating mesh action pauses for a human approval interrupt before it runs: the fleet-wide broadcast and emergency_stop, plus the single-peer tell, send, and stop. You can tune this set with the STRANDS_MESH_HITL_ACTIONS environment variable (set it to all, none, or a comma-separated subset). The first time you run this example, you'll see a robot_mesh-broadcast-approval prompt in your terminal; type y (or yes / approve) to authorize the broadcast. The approval is delivered out-of-band of the LLM's tool arguments, so a prompt-injection attempt that tries to slip an approval flag into the command body cannot bypass the gate.

The transport scales without touching agent code. The built-in Zenoh mesh is the automatic fallback: on the LAN, Zenoh multicast handles peer discovery with no broker, and adding the [mesh-iot] extra routes traffic through AWS IoT Core (MQTT5 with mTLS) for cloud fleets, with a BridgeTransport that fans LAN and cloud behind one API (select it with STRANDS_MESH_BACKEND=bridge).


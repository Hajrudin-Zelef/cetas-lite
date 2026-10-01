---
id: collect-261001-ia-llm/ia-llm/from-the-hugging-face-hub-to-robot-hardware-with-strands-agents-and-lerobot-3
title: "{'observation.state': Sequence(...),"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Hugging Face", "Nvidia"]
dates: []
keywords: ["agent", "agentic", "agents", "apache", "aws", "bedrock", "benchmark", "gpu", "inference", "nvidia", "research", "robotics"]
source: docs/RAG/collect-261001-ia-llm/from-the-hugging-face-hub-to-robot-hardware-with-strands-agents-and-lerobot.md
source_anchor: ""
source_lines: [165, 219]
sha256: 2d13302d5a4a8d0d53d15fb3d0163e45ab6f3dfd7ac859277f9575941f8d2333
---

# {'observation.state': Sequence(...),

For production fleets, Device Connect, a device-aware networking layer developed in collaboration with Arm, handles discovery, presence, structured RPC, event routing, and safety. The same robot_mesh tool dispatches through Device Connect when it is available and falls back to the built-in Zenoh mesh otherwise, so the agent code in this post is unchanged either way. See the Device Connect documentation for setup and current availability.

The full sample is on GitHub at strands-labs/robots in the examples/lerobot/ folder. It packages all five steps into a single CLI script (hub_to_hardware.py) and a notebook (hub_to_hardware.ipynb). The CLI defaults run end-to-end in simulation with the Mock policy. No GPU, no Docker, no Hugging Face credentials needed.

```
uv pip install "strands-robots[sim-mujoco,lerobot,mesh]"
git clone https://github.com/strands-labs/robots.git
cd robots
export STRANDS_MESH_LOCAL_DEV=1
python examples/lerobot/hub_to_hardware.py
```
The recorded dataset lands at `~/.cache/huggingface/lerobot/local/strands-cube-pick/`. To push to the Hugging Face Hub instead of keeping it local, pass `--hf-user <your-user>` after exporting HF_TOKEN with write scope. For real grasping behavior in Step 3, pass `--policy groot --checkpoint <hf_repo>` (requires Docker + NVIDIA GPU) or `--policy lerobot_local --checkpoint <hf_repo>` (requires a GPU and `STRANDS_TRUST_REMOTE_CODE=1`).

The notebook (examples/lerobot/hub_to_hardware.ipynb) walks through the same workflow cell by cell, with narration between each step. Open it in JupyterLab and run top-to-bottom in simulation mode.

The code snippets shown in this setup represent a “hello world” example of setting up Strands Robots with HuggingFace. For more serious, production-ready use cases there are some important considerations users should be aware of:

Supplying untrusted data into agents can lead to prompt injection, where untrustworthy context is treated as LLM instructions. Given the actuation of these robots in physical space, this is an important risk to track. To mitigate this behavior, developers should be careful to feed the robots only data that comes from a trusted source. If not all input data can be trusted, developers should restrict the tools available to the agent to prevent the robots from making safety-critical actions.

The `STRANDS_MESH_LOCAL_DEV=1` setting shared in the code snippets in this blog initializes the robot mesh without authentication or access controls. This means that any device on the same network can provide commands to the robot fleet. This is acceptable for trusted development environments, but is not suitable for untrusted networks or production environments. For these use cases, `STRANDS_MESH_AUTH_MODE=mtls` is required.

The robot_mesh tool's physically-actuating actions affect peers on the network: broadcast and emergency_stop reach every peer, while tell, send, and stop reach a single targeted peer. To prevent an agent from issuing these commands autonomously (or under prompt injection), all five are gated behind a human-in-the-loop interrupt by default. When the agent invokes a gated action, the Strands runtime pauses the agent loop and asks the operator to approve out-of-band of the LLM's tool arguments. You can adjust the gated set with the STRANDS_MESH_HITL_ACTIONS environment variable (all, none, or a comma-separated subset). Per-action rate limits, command validation, and an audit trail run alongside the interrupt. Outside an agent loop (a bare script or unit test), the gated actions fail closed.

The preceding workflow starts a GR00T container, opens serial ports on hardware, and writes a local dataset cache. To return your environment to a clean state:

- **Stop the GR00T inference container:**`agent.tool.gr00t_inference(action="stop", port=5555)` , or use`lifecycle="teardown"` to remove the container as well.
- **Release the serial ports:** if you ran the hardware path, disconnect the SO-101 follower and leader.
- **Optionally remove the local dataset cache:** the recorded dataset lives under`~/.cache/huggingface/lerobot/<repo_id>` . Datasets you pushed to the Hub are unaffected.

The integration's central design choice is that Strands Robots doesn't reimplement what LeRobot already provides. Hardware abstraction, calibration, and the dataset format stay upstream. Strands adds the AgentTool surface that makes them composable from natural language.

Two consequences follow. For users, every dataset on the Hub is an asset an agent can extend, fine-tune from, and deploy against with no conversion step. For developers, simulation data and hardware data share a single file format, so training scripts written for one consume the other unchanged. The line between sim and real becomes a deployment detail, not an architectural divide.

**Figure 4. *The Strands Robots catalog spans arms, humanoids, quadrupeds, and hands, all in the same MuJoCo simulation and behind the same `Robot()` factory. The SO-100 in this post is one of many supported embodiments.***

The full Strands Robots documentation covers the robot catalog, simulation, policy providers, the mesh, and Device Connect in depth. For larger workloads, the strands-labs/robots-sim repository hosts heavier simulation backends including Isaac Sim and Newton, plus a LIBERO benchmark example. Both backends plug into the same Robot abstraction shown in this post, so the agent code stays the same as you scale up.

Contributions are welcome under Apache 2.0. If you build something with this workflow, open an issue with what worked and what didn't. The SDK improves fastest when developer feedback lands directly on the surface that needs it.

- **Strands Robots** (SDK, AgentTools, Robot factory): github.com/strands-labs/robots, Apache 2.0
- **Strands Robots docs** (full documentation): strands-labs.github.io/robots
- **Strands Robots Sim** (examples, simulation backends): github.com/strands-labs/robots-sim
- **The example:** examples/lerobot/hub_to_hardware.py and hub_to_hardware.ipynb
- **How to Build Physical AI Agents: Natural Language for Real-World Robotics** : Live Stream and Blog
- **Diving Deep on Physical AI | S1E4 | Automate with NVIDIA NeMo Agent Toolkit and Bedrock AgentCore** : Live Stream
- **LeRobot** : github.com/huggingface/lerobot - datasets, policies, hardware drivers
- **Strands Agents SDK** : github.com/strands-agents/harness-sdk
- **SmolVLA** : SmolVLA
- **Pi0** : Pi0
- **NVIDIA Isaac-GR00T N1.7** : GR00T N1.7
- **NVIDIA Cosmos3 Nano** : Cosmos 3 Nano

**Cagatay Cali** is a Research Engineer at AWS focused on Agentic AI and robotics. He designs interfaces that connect AI agents to physical robots, enabling developers to control robotic systems through natural language and making agents and robotics development accessible to builders at any skill level.

**Sundar Raghavan** is a Sr Solutions Architect at AWS on the Agentic AI Foundations team. He leads the developer experience for Amazon Bedrock AgentCore, owning the SDK and CLI, and drives the framework and ecosystem integrations strategy. He focuses on how developers build, deploy, and scale production AI agents on AWS. He is currently extending that focus into physical AI, collaborating on Strands Robots to bring the same agent developer experience to robotics.

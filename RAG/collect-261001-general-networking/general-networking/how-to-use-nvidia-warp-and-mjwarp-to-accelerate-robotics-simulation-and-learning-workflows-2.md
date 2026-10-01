---
id: collect-261001-general-networking/general-networking/how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows-2
title: "how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "compute", "gpu", "memory", "throughput"]
source: docs/RAG/collect-261001-general-networking/how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows.md
source_anchor: ""
source_lines: [107, 214]
sha256: 8c7c6f327df56bc2373e3f4574de8b6a3148fe9957e9528b5d4d88e841bb471f
---

# how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows

```
with wp.ScopedCapture() as capture:
     mjw.step(m, d)
wp.capture_launch(capture.graph)
```
2. **Size nconmax / naconmax / njmax tightly**: memory and work scale with them. Tune with mjwarp-testspeed: --measure_alloc and watch overflows in mjwarp-viewer.

**Additional tuning considerations. After sizing contact and constraint buffers, test solver iteration limits without changing task behavior. Meshes and CCD settings can increase memory use; nccdmax / naccdmax can reduce CCD buffer allocation when the measured contact counts allow it. MJWarp’s compact solver uses MuJoCo’s Newton constraint solver and sleeping, not the separate Newton physics-engine framework. Compact-solver and multi-GPU configuration are beyond this walkthrough; consult the MJWarp performance-tuning documentation.**

To train policies on MJWarp physics:

- *Isaac Lab* via*Newton*
- *mjlab* (manager API directly on MJWarp + PyTorch)
- *MuJoCo Playground* via MJX (impl='warp')

**Install / try:** pip install mujoco-warp · mjwarp-viewer path/to/scene.xml · *Colab tutorial*

**The scene.** Nothing here is MJWarp-specific yet: an SO-101 arm, a table, and two cubes to stack, written as ordinary MJCF.

*Figure 2. SO-101 pick-and-place scene, rendered from the MuJoCo CPU simulation. The task is to grasp the red 44 mm cube and stack it on the blue cube; the same robot and scene are used for MJWarp validation.*

```
<mujoco model="so101_pick_place">
  <include file="so101.xml"/>
  <worldbody>
    <light pos="0.3 0 1.5" dir="0 0 -1" directional="true"/>
    <geom name="floor" type="plane" size="0 0 0.05"/>
    <geom name="table" type="box" pos="0.35 -0.04 0.012"
          size="0.16 0.26 0.012" rgba="0.32 0.32 0.32 1"
          friction="1 0.005 0.0005" condim="3"/>
    <body name="red_cube" pos="0.33 -0.13 0.046">
      <freejoint name="red_cube_joint"/>
      <geom type="box" size="0.022 0.022 0.022" mass="0.08"
            rgba="0.85 0.05 0.04 1" friction="1.2 0.005 0.0005" condim="3"/>
    </body>
    <body name="blue_cube" pos="0.33 0.06 0.046">
      <freejoint name="blue_cube_joint"/>
      <geom type="box" size="0.022 0.022 0.022" mass="0.08"
            rgba="0.05 0.20 0.90 1" friction="1.2 0.005 0.0005" condim="3"/>
    </body>
  </worldbody>
</mujoco>
```
For an MJCF box, the size values are half-extents: size=”0.022 …” defines a cube with 44 mm edges. The task uses this size for its success thresholds. The arm base is at the origin, its reach is along +X, and the cubes are arranged along Y.

In the companion repository this file is generated rather than hand-written: resolve_pick_place_scene() copies the Menagerie arm into .generated/, fills the table and cube coordinates from a robot profile, and writes scene_pick_place.xml. The walkthrough uses the SO-101 profile; the optional reBot variant is described below.

**Loading it.** Compilation and stepping are ordinary MuJoCo:

```
import mujoco
mjm = mujoco.MjModel.from_xml_path("scene_pick_place.xml")
mjd = mujoco.MjData(mjm)
fps = 50 # controller rate
sim_substeps = 10 # physics steps per control frame
frame_dt = 1.0 / fps
mjm.opt.timestep = frame_dt / sim_substeps
controller = PickPlaceController(spec=spec) # waypoints + damped-least-squares IK
for _ in range(600): # 600 control frames
    ctrl = controller.step(mjm, mjd, frame_dt)
    for _ in range(sim_substeps):
        mjd.ctrl[: mjm.nu] = ctrl
        mujoco.mj_step(mjm, mjd)
```
Keep that shape in mind: **compute controls once per frame, step physics sim_substeps times.** Gate 2 changes only the inner loop, which is what makes the migration easy  to review.

**Match the simulation and control rates. At 50 control frames per second and 10 physics substeps per frame, use a physics timestep of 0.002 seconds. Set it before the CPU rollout and before uploading the model with mjw.put_model so both backends advance the same simulated time:**

```
mjm.opt.timestep = frame_dt / sim_substeps # 50 Hz × 10 substeps -> 0.002 s
```
Without that line, every later measurement inherits the mismatch: parity comparisons, throughput numbers quoted as “simulated seconds,” and any learned policy whose action rate no longer matches deployment.

**Check whether the cubes are stacked successfully.** With 44 mm cubes, success becomes two measurable conditions: a horizontal center error of xy_err ≤ 0.015 m (measured between the cube centers) and a vertical separation of 0.035 m ≤ dz ≤ 0.055 m between cube centers (one cube edge, with slack for settling). Evaluate both conditions after the cubes have settled; a successful process exit alone does not establish task success.

Run the CPU task from the companion checkout. Publication blocker: confirm the accessible repository URL and pinned dependency and asset versions before publishing these instructions; the repository placeholder below is not an executable URL.

```
git clone https://github.com/NVIDIA/accelerated-computing-hub.git blogs
cd blogs/tutorials/sim2real-blogs/notebooks/mujoco
uv venv --python 3.12 && source .venv/bin/activate
uv pip install -r requirements.txt
cd /tutorials/sim2real-blogs/notebooks/mujoco
python solutions/so101_pick_place_solution.py --headless-steps 600 --debug
```
The run ends by printing the two numbers above (stack check: xy_err=… dz=…), which is the assertion the rest of the article compares against. so101_pick_place.py next to it is the same program with the physics steps left as exercises.

The arm comes straight from *MuJoCo Menagerie* pinned to a known-good commit, since Menagerie assets change, so treat the scene as a template. Optional reBot variant. The companion code also exposes --robot rebot with a separate profile for the scene layout, gripper, and capacity limits (nconmax=256, njmax=500). This walkthrough uses SO-101. Validate the reBot asset and task separately before reporting its results.

Run **one** world on the GPU first, with the host still in the loop, so you can watch the same task in the same viewer and compare the same two numbers. Upload the model, allocate batched state, seed it from the initialized host state, and run one forward pass before stepping:

```
wp.init()
import mujoco_warp as mjw
device = wp.get_device()
m = mjw.put_model(mjm)
d = mjw.make_data(mjm, nworld=1, nconmax=spec.nconmax, njmax=spec.njmax)
wp.copy(d.qpos, wp.array(mjd.qpos[None, :], dtype=wp.float32, device=device))
wp.copy(d.qvel, wp.array(mjd.qvel[None, :], dtype=wp.float32, device=device))
wp.copy(d.ctrl, wp.array(mjd.ctrl[None, :], dtype=wp.float32, device=device))
mjw.forward(m, d)
```
Every device array carries a leading world dimension, which is why the host state is indexed as mjd.qpos[None, :], shape (1, nq) instead of (nq,). Scaling to thousands of worlds later changes only that leading dimension, not the calls. mjw.put_model() also doubles as a compatibility check: it raises if the model uses unsupported features rather than silently dropping them.

Seeding the three fields explicitly is the transparent option, and it makes clear exactly what crosses to the device; mjw.put_data(mjm, mjd, nworld=…) carries the whole initialized struct over in one call instead.

The frame loop is then the Gate 1 loop with its inner step redirected to the GPU and mirrored back:


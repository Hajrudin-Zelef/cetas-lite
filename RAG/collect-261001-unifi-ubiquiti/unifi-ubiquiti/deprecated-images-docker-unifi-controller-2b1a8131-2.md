---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/deprecated-images-docker-unifi-controller-2b1a8131-2
title: "deprecated-images-docker-unifi-controller-2b1a8131"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2023-09-06"]
keywords: ["memory", "parameters"]
source: docs/RAG/collect-261001-unifi-ubiquiti/deprecated-images-docker-unifi-controller-2b1a8131.md
source_anchor: ""
source_lines: [144, 187]
sha256: e731e0ba6b9ee098291b8d0aa14947a9c09c329ce8280d279aa3a1fd62f32bcc
---

# deprecated-images-docker-unifi-controller-2b1a8131

-  Recreate a new container with the same docker run parameters as instructed above (if mapped correctly to a host folder, your /config folder and settings will be preserved)
-  You can also remove the old dangling images:
Via Watchtower auto-updater (only use if you don't remember the original parameters)¶
-  Pull the latest image at its tag and replace it with the same env variables in one run:
-  You can also remove the old dangling images: docker image prune
Warning
We do not endorse the use of Watchtower as a solution to automated updates of existing Docker containers. In fact we generally discourage automated updates. However, this is a useful tool for one-time manual updates of containers where you have forgotten the original parameters. In the long term, we highly recommend using Docker Compose.
Image Update Notifications - Diun (Docker Image Update Notifier)¶
Tip
We recommend Diun for update notifications. Other tools that automatically update containers unattended are not recommended or supported.
Building locally¶
If you want to make local modifications to these images for development purposes or just to customize the logic:
git clone https://github.com/linuxserver/docker-unifi-controller.git
cd docker-unifi-controller
docker build \
  --no-cache \
  --pull \
  -t lscr.io/linuxserver/unifi-controller:latest .
The ARM variants can be built on x86_64 hardware using multiarch/qemu-user-static
Once registered you can define the dockerfile to use with -f Dockerfile.aarch64.
Versions¶
- 01.01.24: - Deprecate.
- 05.09.23: - Add deprecation warning as per https://info.linuxserver.io/issues/2023-09-06-unifi-controller.
- 04.09.23: - Bump JRE to 17 to support v7.5.
- 02.05.23: - Cleanup apt-get install during build to reduce image size.
- 18.03.23: - Add mongoless branch.
- 10.03.23: - Test writing to /run/unifi and symlink to /config/run if it fails.
- 20.02.23: - Migrate to s6v3, install deb package on build, fix permissions.
- 23.01.23: - Exclude run from/config volume.
- 30.11.22: - Bump JRE to 11.
- 01.06.22: - Deprecate armhf.
- 23.12.21: - Move min/max memory config from run to system.properties.
- 22.12.21: - Move deb package install to first init to avoid overlayfs performance issues.
- 13.12.21: - Rebase 64 bit containers to Focal.
- 11.12.21: - Add java opts to mitigate CVE-2021-44228.
- 11.06.21: - Allow for changing Java initial mem via new optional environment variable.
- 12.01.21: - Deprecate the LTS tag as Unifi no longer releases LTS stable builds. Existing users can switch to thelatest tag. Direct upgrade from 5.6.42 (LTS) to 6.0.42 (latest) tested successfully.
- 17.07.20: - Rebase 64 bit containers to Bionic and Mongo 3.6.
- 16.06.20: - Add logrotate.
- 02.06.20: - Updated port list & descriptions. Moved some ports to optional.
- 14.11.19: - Changed url for deb package to match new Ubiquity domain.
- 29.07.19: - Allow for changing Java mem limit via new optional environment variable.
- 23.03.19: - Switching to new Base images, shift to arm32v7 tag.
- 10.02.19: - Initial release of new unifi-controller image with new tags and pipeline logic

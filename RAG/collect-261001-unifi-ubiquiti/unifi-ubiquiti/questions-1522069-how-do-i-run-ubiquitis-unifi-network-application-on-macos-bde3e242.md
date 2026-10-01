---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1522069-how-do-i-run-ubiquitis-unifi-network-application-on-macos-bde3e242
title: "questions-1522069-how-do-i-run-ubiquitis-unifi-network-application-on-macos-bde3e242"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: ["2020-05-15", "2022-04-10", "2023-01-13", "2023-08-27"]
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1522069-how-do-i-run-ubiquitis-unifi-network-application-on-macos-bde3e242.md
source_anchor: ""
source_lines: [1, 52]
sha256: 7eeb70e3472e077a42c0919c70b9210590f564d4075e5b9a2866ed7f59016bdb
---

# questions-1522069-how-do-i-run-ubiquitis-unifi-network-application-on-macos-bde3e242

I recently purchased two UniFi nanoHD access points but not a Cloud Key. How can I run the UniFi Network Application (previously called UniFi Network Controller) on my macOS computer?
2 Answers 2
I'm aware of two different ways to do this.
Using Docker
This might be a good option if you're already familiar with Docker, but it may not be worth the learning curve otherwise. It appears there are two well-maintained Docker images for the UniFi Network Application.
- The LinuxServer.io team's Docker image. I haven't tried it but it looks like a decent choice.
- An image from a guy named Jacob Alberty. I have tried this one and it works well.
Extra effort must be made to make adoption work when running in Docker. The default network type for Docker (bridge) means containers are not accessible on the network. For Layer 2 adoption, the documentation for Jacob Alberty's image mentions using host networking or the macvlan network. I've used host networking with this Docker image in Linux and it works well, however at the time of this writing (2022-04-10) host networking is not available in macOS Docker Desktop. I haven't tried using macvlan but it might work well (with either image, even though the LinuxServer.io documentation doesn't mention it). For Layer 3 adoption, it looks like you must enter the Docker host's IP in the UniFi Network Application's settings (the documentation for both Docker images has details of where to find the setting).
The Jacob Alberty image appears to include some support for specifying an https certificate. I haven't tried it, but if using a certificate trusted by your browser is important for you then this might be a reason to choose the Jacob Alberty image.
Installing Natively
This works fine and isn't tremendously difficult to get running, but it is unfortunately more work than it should be.
- Download your preferred version Look for "UniFi Network Application [version] for macOS" on the Unifi download page. If you have no preference then go with the newest version.
- Install Network Application Find UniFi.pkg in Finder. Control+click and select Open with... → Installer. Without this macOS complains that "'UniFi.pkg' can’t be opened because Apple cannot check it for malicious software." Then select "Open" if prompted and proceed through the installer.
- Install Java 11 Note from 2023-08-27: You can try skipping this step—it might not be needed anymore. I was able to install version 7.4.162 and I think I don't have Java installed, so it might be bundled now. Or I could be wrong. I don't see anything about it in the release notes. Note: Versions prior to 7.3.x required Java 8 so if you've chosen to install an older version of the Network Application then see an earlier revision of this answer for advice on installing Java 8 and also the command for starting older versions of the Network Application. For versions 7.3.x and newer you need Java 11. (As of 2023-01-13 the release notes state "Java 12 and later are not yet supported." Presumably Java 13 through 19 are also not supported.) Installing with Homebrew is easy: 
  - Follow the instructions to install Homebrew if you don't already have it.
  - In a terminal, run brew install openjdk@11 . For me this installed 24 dependencies, too.
  - As suggested by the output of the above command, create a symlink to allow the system Java wrappers to find this JDK: sudo ln -sfn /usr/local/opt/openjdk@11/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk-11.jdk
- Adjust security settings Note: These instructions changed significantly due to the Settings overhaul in macOS 13 Ventura. You can reference an earlier revision of this answer for guidance on macOS 12 and older, but also note that those instructions might be inaccurate for newer versions of the Network Application and Java 11. 
  - Navigate to System Settings → Network. Select Firewall then click the Options... button. EnsureBlock all incoming connections is off. No other changes are required. Click OK to save.A few notes: 
    - If Automatically allow downloaded signed software to receive incoming connections is enabled then that is sufficient to give Network Application access.
    - If Automatically allow downloaded signed software to receive incoming connections is disabled then you'll be prompted to allow incoming connections when Network Application is started.
    - Proactively adding Applications/UniFi to the list of apps with allowed incoming connections doesn't seem to work. I don't know why. But it's not important—macOS will prompt you about it later, if needed.
    - Instead of the above firewall changes you could completely turn off your firewall, but I don't recommend it. Firewalls provide a useful layer of protection.
  - If 
  - Attempt to start the Network Application. You'll use this process anytime you need to start it in the future: Either run the UniFi app from the Applications folder or use ⌘+Space and search for "UniFi." You should get this error dialog:
  - "Damaged" is not accurate—macOS is just preventing it from running correctly. Click the Cancel button then navigate to System Settings → Privacy & Security. Scroll down past all the applications and look for the "Security" heading. You should see the following dialog:
  - Use the "Open Anyway" button to approve it. You'll get a second dialog that says the application is damaged:
 
  - I have no idea what macOS is complaining about here. Ignore this and click the "Open" button. (Side note: This area of the macOS user experience really needs work.) If macOS asks, "Do you want the application 'UniFi.app' to accept incoming network connections?" you must click "Allow" for Network Application to work correctly.
Two follow-up notes:
- You can undo any changes to your firewall when you're done with the Network Application (assuming you're not planning on leaving it running all the time). Though the firewall changes are mostly harmless if all you did was allow UniFi to receive incoming connections.
- You can also run the Network Application "headlessly." I haven't tried it, but the instructions at https://akrabat.com/run-the-unifi-controller-headless-on-mac/ look promising. Note that you'll need to use a different exec string than the ones suggested on that page.
- 
        1Wow this is so needlessly complex, thanks for the writeup. In my case I already had java 12 installed with brew so I had to execute the JAR using the correct virtual machine: /Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home/bin/javaPeter Johnson– Peter Johnson2020-05-15 15:15:59 +00:00Commented May 15, 2020 at 15:15
- 
            
            
- 
        1This got me 90% of the way there, but there's an important omission. After step 3, substep 1 (installing brew), you have to tap the openjdk repo first:brew tap adoptopenjdk/openjdk .
                
                spamguy– spamguy2020-10-24 22:09:54 +00:00Commented Oct 24, 2020 at 22:09
- 
        1The above cask install command no longer works with newer versions of brew. Instead you needbrew install --cask adoptopenjdk8 catchdave– catchdave2021-04-27 21:28:54 +00:00Commented Apr 27, 2021 at 21:28
- 
        1brew now suggests usingtemurin8 instead ofadoptopenjdk8 Dave– Dave2021-08-19 07:23:51 +00:00Commented Aug 19, 2021 at 7:23
- 
        1@MarkDoliner it seems this is no longer required and Unifi finally runs on newer versions of Java, happy days!Hackeron– Hackeron2023-01-07 14:52:11 +00:00Commented Jan 7, 2023 at 14:52
I got it working in docker in a few simple steps! I've made a docker config on guthub with full instructions, but here are the basic steps:
- Decide where UniFi should save persistant data on your machine.
- Run a docker compose script.
- Browse to https://localhost:8443 and complete the initial setup.
- Go to settings and enter the IP address of your machine.

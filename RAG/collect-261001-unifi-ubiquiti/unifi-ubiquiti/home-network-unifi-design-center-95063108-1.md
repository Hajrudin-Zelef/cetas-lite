---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/home-network-unifi-design-center-95063108-1
title: "home-network-unifi-design-center-95063108"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/home-network-unifi-design-center-95063108.md
source_anchor: ""
source_lines: [1, 50]
sha256: 350e9f74ec98446205d45bcdee1f7ed4569e5f50382e7ed0b1f4c24d5181517b
---

# home-network-unifi-design-center-95063108

When you are building or remodeling a house, or need to plan a network installation for an office building, then you should take a look at UniFi Design Center. It allows you to visualize the wireless network coverage and also helps you with planning your network cable installation.
The most difficult part of planning a network installation is finding the best spots for the access points. Where do you place them, and how do you get the best coverage? But also, how much ethernet cable do you need for the installation?
In this article
In this article, we will look at the UniFi Design Center, how to use it, and some tips to create the best network layout.
UniFi Design Center
UniFi Design Center is an online tool that you can use freely to plan your network installation. It allows you to upload a floorplan of your building, set the scale and height, and layout all your network equipment.
To plan your wireless network installation, you can first trace all the walls and windows in your floor plan. This way the tool knows where the walls are, allowing it to calculate the wireless network performance for each room.
You can also plan your network camera installation with the Design Center. It will show the field of view of each camera, so you can make sure that every spot is covered.
There are also a few limitations to the design center that are good to know. Even though you can add multiple floorplans in one project, the wireless network coverage doesn’t take the different floors into account.
When it comes to the cable installation we face a similar problem. Even though we have set the height of the floors, we can’t configure the height for each ethernet drop. So you will need to account for some extra length for each ethernet wall socket.
Preparing the Floorplan
The first step when you start a new project is to prepare the floorplan of your building. To add a new floorplan you have two options, upload a plan, which can be a PNG, JPEG or PDF file, or draw a floorplan from scratch.
I recommend using a drawing of the floor plan as an underlayer if you have one because this makes it a lot easier. You can use Floorplanner.com or Sketchup to create a floor plan for your house for free.
When you create a new project, you will first need to upload the image of the floor plan. You can later add multiple floors if needed (1).
After you have uploaded the floor plan, you will need to set the floor plan scale (2). For the scale, you will need to draw a line on the floor plan and enter the width of that part. It’s also a good idea to set the ceiling height because this will affect the cable length.
You can set the ceiling height in the Multi-Floor menu (3). To change the scale after you have set it, simply click on the scale icon in the lower-left corner (4)
Drawing the Walls
To plan your wireless network, you will first need to draw all the walls in the UniFi Design Center. The walls are used to calculate how much dB is deducted from your wireless network signal. We have three types of walls that we can use:
- Glass – 3dB reduction
- Drywall – 5dB reduction
- Concrete – 15dB reduction
When drawing the outer walls, the window openings can be a bit of a challenge at first. I find it easiest to first click the left mouse button to place the wall and then hit the right mouse button to stop the drawing. Then continue on the other side of the window opening with the wall. You can fill the openings later with a glass “wall”.
For door openings, I recommend drawing in a piece of drywall. Also, if you have any full-height cabinets or closets, I recommend drawing them as a wall also. This way you will get the most realistic results in the wireless coverage measurement.
Placing the Access Points
With the walls drawn, we can start with placing the access points. You want to design the wireless network based on the 5Ghz band. This will give you the highest throughput and best performance. For residential networks, the UniFi U6+ is more than enough. This access point is small but powerful. If you are designing the network for a small or medium business, you might want to look at the U6 Pro.
- Make sure that you have enabled the WiFi coverage view in the top center
- Set it to 5 GHz (in the lower-right corner)
- Click on Place Devices
- Expand WiFi and select the access point you want to place
- Place the access point on the map
You want to aim for at least -62 dBm in the location where the wireless network will be used the most. If you look at the example above, you see we have good coverage in the bedrooms, office, and living room on 5 GHz. These are the areas where the wireless network will be used the most.
The coverage is a bit lower in the kitchen/dining room, but it is still enough. And keep in mind we have a strong signal on 2.4 GHz as well.
Now in this particular network design, there is an extra network cable drop inside the closest behind the dining table, so it’s possible to add an additional access point later if needed.
Planning Network Cables
The next step is to plan all the network cables. We want to plan where all the ethernet drops will come, so we know how much network cable we are going to need.
If you go to Place Devices you will find wall sockets and cable connections under the Accessories. First, place all the sockets in your floor plan. It’s also a good idea to place the network equipment that you are planning to use in the drawing.
Now, you can draw all the network cables manually, tracing the complete path on how you are going to install them. But I find it easier to first lay out the Cable Route in the floor plan and then use the Auto-draw cable option.
- Click on Draw Cabling
- Choose Cable Route
- Draw the main routes where the network cables will go
The cable routes are displayed as a grey line in the drawing. As I mentioned at the beginning of the article, the cable router doesn’t account for any height difference. So, if your cables are placed in the ceiling, then you will need to add an extra 5 meters (15 feet) for each wall outlet.
When done, click on Cable Route (1) again,n and this time choose Auto-draw Cables (4). It will show which devices it was able to connect to and the length of each cable (we can look that up later as well). If you now hover over a device, you will see how the cable is routed to the nearest switch.
The design tool takes available ports into account when routing each ethernet connection. One problem with the auto-routing option is that it doesn’t take the PoE requirements into account.
Placing Cameras
The last step in planning the network installation is to place the cameras. The view angle of the camera is displayed in blue. You can rotate the camera by dragging the blue arc around the camera icon.
When you have placed all the cameras, you will need to connect does as well to your network. Just use the Auto-draw cable option again to connect all the devices. In the example above, you will see that the G4 Pro doorbell has a wireless connection to the nearest access point.
But in this case, the signal is too weak for the doorbell. So an option would be to place the doorbell on the opposite side of the door or use the PoE version of the doorbell.
Reviewing the Network Layout
When you have completed the network layout, we start with extracting information from UniFi Design Center. On the left-hand side of the tool, you see a couple of options. The first thing you want to take a look at is the Port View.
Here you can see how each device is connected and how many ports you have left. Important to check here is that all devices that require a PoE port is connected to a PoE-capable switch. This is something the tool unfortunately doesn’t take into account.

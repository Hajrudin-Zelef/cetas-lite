---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-62-2
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-62.md
source_anchor: ""
source_lines: [146, 302]
sha256: eac10a0c198d52a3e541bac721c1be6da0ab5b3e15f402cb517badfef1716910
---

# Summary

- **undo**
- **redo**
- **Safe Mode**
- Currently loaded session

More about Safe mode and undoing performed actions read in this article.

On the right side is located:

- an indicator that shows whether the WinBox session uses encryption
- WinBox traffic indicator displayed as a green bar,
- Custom info fields that can be added by the user by right-clicking on the toolbar and picking available info fields from the list

# Work Area and Child Windows

WinBox has an MDI interface meaning that all menu configuration (child) widows are attached to the main (parent) WinBox window and is showed in the work area.

Child windows can not be dragged out of the working area. Notice in the screenshot above that the **Interface** window is dragged out of the visible working area and a horizontal scroll bar appeared at the bottom. If any window is outside visible work area boundaries the vertical or/and horizontal scrollbars will appear.

## Child window menu bar

Each child window has its own toolbar. Most of the windows have the same set of toolbar buttons:

- **Add** - add a new item to the list
- **Remove** - remove the selected item from the list
- **Enable** - enable selected item (the same as**enable** command from console)
- **Disable** - disable selected item (the same as**disable** command from console)
- **Comment** - add or edit a comment
- **Sort** - allows to sort out items depending on various parameters.`Read more >>` 

Almost all windows have a quick search input field on the right side of the toolbar. Any text entered in this field is searched through all the items and highlighted as illustrated in the screenshot below

Notice that on the right side next to the quick find input filed there is a drop-down box. For the currently opened (IP Route) window, this drop-down box allows to quickly sort out items by routing tables. For example, if the **main** is selected, then only routes from the main routing table will be listed. 

A similar drop-down box is also in all firewall windows to quickly sort out rules by chains.

## Sorting out displayed items

Almost every window has a **Sort** button. When clicking on this button several options appear as illustrated in the screenshot below

The example shows how to quickly filter out routes that are in the 10.0.0.0/8 range

1. Press **Sort** button
2. Choose **Dst.Address** from the first drop-down box.
3. Choose **in** form the second drop-down box. "in" means that filter will check if DST address value is in range of the specified network.
4. Enter the network against which values will be compared (in our example enter "10.0.0.0/8")
5. These buttons are to add or remove another filter to the stack.
6. Press the **Filter** button to apply our filter.

As you can see from the screenshot WinBox sorted out only routes that are within the 10.0.0.0/8 range.

Comparison operators (Number **3** in the screenshot) may be different for each window. For example "IP Route" window has only two **is** and **in**. Other windows may have operators such as "is not", "contains", "contains not".

WinBox allows building a stack of filters. For example, if there is a need to filter by destination address and gateway, then

- set the first filter as described in the example above,
- press **[+]** button to add another filter bar in the stack.
- set up a second filter to filter by the gateway
- press the **Filter** button to apply filters.

You can also remove unnecessary filters from the stack by pressing the **[-]** button.

## Customizing list of displayed columns

By default, WinBox shows the most commonly used parameters. However sometimes it is needed to see other parameters, for example, "BGP AS Path" or other BGP attributes to monitor if routes are selected properly.

WinBox allows to customize displayed columns for each individual window. For example to add BGP AS path column:

- Click on the little arrow button (**1** ) on the right side of the column titles or right mouse click on the route list.
- From popped up menu move to **Show Columns** (**2** ) and from the sub-menu pick the desired column, in our case click on**BGP AS Path** (**3** )

Changes made to window layout are saved and next time when WinBox is opened the same column order and size are applied.

### Detail mode

It is also possible to enable **Detail mode**. In this mode all parameters are displayed in columns, the first column is the parameter name, the second column is the parameter's value.

To enable detail mode right mouse click on the item list and from the popup menu pick **Detail mode**

### Category view

It is possible to list items by categories. In this mode, all items will be grouped alphabetically or by another category. For example, items may be categorized alphabetically if sorted by name, items can also be categorized by type like in the screenshot below.

To enable Category view, right mouse click on the item list and from the popup menu pick **Show Categories**

## Drag & Drop

It is possible to upload and download files to/from the router using WinBox drag & drop functionality. You can also download the file by pressing the right mouse button on it and selecting "Download".

## Traffic monitoring

WinBox can be used as a tool to monitor the traffic of every interface, queue, or firewall rule in real-time. The screenshot below shows Ethernet traffic monitoring graphs.

## Item copy

This shows how easy it is to copy an item in WinBox. In this example, we will use the COPY button to make a Dynamic PPPoE server interface into a Static interface.

This image shows us the initial state, as you see DR indicates "D" which means Dynamic:

Double-Click on the interface and click on COPY:

A new interface window will appear, a new name will be created automatically (in this case pppoe-in1)

After this Down/Up event this interface will be Static:

# Transferring Settings

- Managed router transfer - In the File menu, use Save As and Open functions to save the managed router list to file and open it up again on a new workstation.

- Router sessions transfer - In the Tools menu, use Export and Import functions to save existing sessions to file and import them again on a new workstation.

# WinBox v3 Keyboard Shortcuts

| Shortcut | Description | 
|---|---|
| **Ctrl** +**F** | Find | 
| **Ctrl** +**G** | Find Next | 
| **F3**  | Find / Find next | 
| **Ctrl** +**M** | Add or edit a comment | 
| **Ctrl** +**E** or**Num+** | Enables a selected setting | 
| **Ctrl**  +**D**  or**Num-** | Disables a selected setting | 
| **Ctrl** +**+** | Zoom in WinBox | 
| **Ctrl**  +**-** | Zoom out WinBox | 
| **Tab** | Choose next control | 
| **Tab**  +**Shift** | Choose previous control | 
| **Space** | Select focused control | 
| **F4** or**Esc** | Close window | 
| **F6**  | Focus previous window | 
| **F6**  +**Shift** | Focus next window | 
| **Insert**  | Add new entry into list | 
| **Delete**  | Delete entry from list | 

# WinBox v4 Keyboard Shortcuts

| Shortcut | Description | 
|---|---|
| **Ctrl** +**E** or**Num+** | Enables a selected setting | 
| **Ctrl**  +**D**  or**Num-** | Disables a selected setting | 
| **Ctrl/CMD** +**+** | Zoom in WinBox | 
| **Ctrl/CMD** +**-** | Zoom out WinBox | 
| **Ctrl/CMD**  +**0** | Reset zoom | 
| **Ctrl** + **Tab** | Choose next tab | 
| **Ctrl** +**Shift** +**Tab** | Choose previous tab | 
| **Tab** | Choose next control | 
| **Shift** +**Tab** | Choose previous control | 
| **Space** | Select focused control | 
| **Ctrl + F4**  or**Alt** /**Cmd + W**  | Close window | 
| **Alt** +**A**  /**Cmd** +**Shift** +**S** | Focus previous window | 
| **Alt** /**Cmd** +**S** | Focus next window | 
| **Delete**  | Delete entry from list | 
| **Alt** /**Cmd**  +**Opt** +**F**  | Global menu search | 
| **Alt** /**Cmd** +**T**  | Open new Terminal window | 

# Troubleshooting

#### WinBox cannot connect to the router's IP address, devices do not show up in the Neighbors list


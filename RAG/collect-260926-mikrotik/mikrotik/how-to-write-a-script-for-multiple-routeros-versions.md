---
id: collect-260926-mikrotik/mikrotik/how-to-write-a-script-for-multiple-routeros-versions
title: "how-to-write-a-script-for-multiple-routeros-versions"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/how-to-write-a-script-for-multiple-routeros-versions.md
source_anchor: ""
source_lines: [1, 61]
sha256: 23410303d9dde7bec838c17eb6f512a1f449dd724afd6dc4713102015fc9146f
---

# how-to-write-a-script-for-multiple-routeros-versions

I’m scripting for v4.x and v5.x routerOS but some commands are differents in versions. As an example, to set a remote logging server

```
/ system logging action
```

in v4

```
set remote name="remote" target=remote remote=xxx.xxx.xxxx.xxx:514
```

while in v5

```
set remote name="remote" target=remote remote=xxx.xxx.xxxx.xxx remote-port=514
```

Obviously the remote-port attribute isn’t recognised in v4.

Is there a way to wrrite code for both versions? I tried something like this

```
:if ([/system package get system version] > 5) do={
    set remote name="remote" target=remote remote=10.20.0.1 remote-port=514
} else={
    set remote name="remote" target=remote remote=10.20.0.1:514
}
```

which works in v5 while in v4 give an error in the remote-port attribute.

Suggests?

             
            
           
          
            
            
              You can use :parse to run commands. For example:

```
:local cmdv5 [:parse "set remote name=\"remote\" target=remote remote=10.20.0.1 remote-port=514"];
:local cmdv4 [:parse "set remote name=\"remote\" target=remote remote=10.20.0.1:514"];
:if ([/system package get system version] > 5) do={
    $cmdv5;
} else={
    $cmdv4;
}
```

             
            
           
          
            
            
              mrz,

I think there’s a bigger issue here as well.  It would be nice to keep backward compatibility when adding new functionality/properties. (This would help me tremendously).

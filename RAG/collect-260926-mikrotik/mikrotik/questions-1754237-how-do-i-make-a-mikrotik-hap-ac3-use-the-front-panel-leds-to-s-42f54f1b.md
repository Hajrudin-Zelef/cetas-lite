---
id: collect-260926-mikrotik/mikrotik/questions-1754237-how-do-i-make-a-mikrotik-hap-ac3-use-the-front-panel-leds-to-s-42f54f1b
title: "TYPE       LEDS"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1754237-how-do-i-make-a-mikrotik-hap-ac3-use-the-front-panel-leds-to-s-42f54f1b.md
source_anchor: ""
source_lines: [1, 42]
sha256: 8f8bb645dd55619b8cc5e20f05e3243f80ea09e59f0d473dcf5391fdaf68a11f
---

# TYPE       LEDS

In the CLI, set the LEDs up so that led1-5 are used only by individual LED groups:
/system leds print
/system leds set 0 leds=""
/system leds add type=off leds=led1
/system leds add type=off leds=led2
/system leds add type=off leds=led3
/system leds add type=off leds=led4
/system leds add type=off leds=led5
Now when you check the LEDs it should look something like this:
[admin@house] > /system leds print
Flags: * - DEFAULT
Columns: TYPE, LEDS
#   TYPE       LEDS
0 * off
1 * poe-fault  poe-led
2 * off        led1
3 * off        led2
4   off        led3
5   off        led4
6   off        led5
(that it adds more "default" entries seems to be a bug in routerOS 7.7)
Now add a script and edit it:
/system script add name=usage-leds
/system script edit usage-leds source
Put this in the editor:
/interface monitor-traffic ether1 once do={
    :local r (rx-bits-per-second / 1024);
    :if ($r > 0) do={ /system leds set [find leds=led1] type=on } \
        else={ /system leds set [find leds=led1] type=off };
    :if ($r >= 138240) do={ /system leds set [find leds=led2] type=on } \
        else={ /system leds set [find leds=led2] type=off };
    :if ($r >= 276480) do={ /system leds set [find leds=led3] type=on } \
        else={ /system leds set [find leds=led3] type=off };
    :if ($r >= 414720) do={ /system leds set [find leds=led4] type=on } \
        else={ /system leds set [find leds=led4] type=off };
    :if ($r >= 552960) do={ /system leds set [find leds=led5] type=on } \
        else={ /system leds set [find leds=led5] type=off };
}
The constants are in kbps, and can be set to your desired threshold.  These ones are even from "any traffic" for the first, in fractions up to 540mbps (540*1024) for full deflection.
Now schedule it to run:
system scheduler add disabled=no interval=1s name="load LEDs" on-event=usage-leds
Run a speedtest or something and watch the LEDs indicate usage.

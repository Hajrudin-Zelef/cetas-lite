---
id: collect-260926-mikrotik/mikrotik/help-on-restoring-routeros-on-rb951g-2hnd-2
title: "help-on-restoring-routeros-on-rb951g-2hnd"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/help-on-restoring-routeros-on-rb951g-2hnd.md
source_anchor: ""
source_lines: [132, 223]
sha256: e4ccf466c77c7b0f56ae982e14a67e7bb153b3f78f85ac9e31a495486ee925dc
---

# help-on-restoring-routeros-on-rb951g-2hnd

Thank you very much

             
            
           
          
            
            
              USB to serial converter can only be connected to USB port of your device, and it only works with RouterOS itself but not for the bootloader, so it’s not an option.

Out of the 9 pins on the DB-9 connector, you only need three - SG, TxD and RxD. But you have to convert the voltage levels from TTL (0 and 5 V) to standard RS-232 levels (+12 and -12 V, respectively). So you need a discrete circuit or an integrated one, like max-232, and the easiest way may be to find a complete module, which you connect using passive wires alone (i. e. no electronic components to add), at your favourite Arduino shop. You connect the ground, TTL sides of Tx and Rx, and power supply (+3 V) to the RB side, and connect the same ground and the RS-232 sides of Tx and Rx to the 9-pin connector. There is an issue - you have to connect the output of one device to the input of the other one, and the Rx and Tx names may be confusing, so better to measure voltage against ground before connecting. What gives anything between 3 and 25 volts in either polarity against ground is an output, what gives less is most likely input. This is an example of such module (there are many different ones out there).

             
            
           
          
            
            
              I used the USB to serial cable to netinstall RB 433 when I bricked it some years ago

USB → PC

Serial > RB 433

 
            
           
          
            
            
              
Yeah, but the RB433 probably had true RS-232 levels, or the USB to serial cable had TTL levels at the serial end as it was one of those suggested by @ysha - the photo shows it must have been the first case.

             
            
           
          
            
            
              So Is it safe to use the one USB-Serial Converter I used on the RB433 (I still have it) ?

I can use a solder to fix on the cables to the board but I don’t know on which pads to connect on the RouterBoard

             
            
           
          
            
            
              
One more time:

- If you want to connect the USB end of that cable to the 951’s USB port, it won’t help as the bootloader doesn’t work with serial over USB.
- If you want to connect the USB end of that cable to the PC, and the 9-pin end to the 951, you need the RS-232/TTL level converter as suggested by me to use your cable. Or you may replace the combination of these two items by a single USB-to-TTL-serial cable as recommended by @ysha.

Regarding the pins to connect - at the photo, there are four pads in the UART area - GND, Rx, Tx, and 3V. For @ysha’s suggestion, it is enough to use GND, Rx,Tx as the converter electronic is powered from the USB end. For my suggestion, you need also the 3V to power the electronic of the level converter, as no power is sent across the DB-9 connector.

             
            
           
          
            
            
              Thank you @sindy & @ysha for your input. I’ll give it a try and will let you know how it goes

             
            
           
          
            
            
              @AMatt, there is still a connection option, but to do this, you will need to open the plastic case of your adapter to access the converter board your adapter and connect to contacts with TTL levels. For example as shown in the diagram below, connect to the pins of the FT232R GND, RXD, TXD chip (between FT232R and RS232 level converter). But you must have the appropriate qualifications to understand the schematic diagram of the adapter you have (or find a similar one on the Internet). The connection points depend on the USB-TTL chip used in your adapter.

)

             
            
           
          
            
            
              Thank you @ysha, I’ll try and get the adapter tomorrow. If I couldn’t find it in the computer and IT Gadget store, I’ll give it a try and try to connect as you suggested.

However, most of USB to UART TTL adapters are using +5v and the RB 951 UART says +3v, is it safe to proceed with the +5v adapter ?

             
            
           
          
            
            
              You shouldn’t do that, then it is better to order something from proposed earlier above. It is possible to convert levels +5 and +3.3 V but additional schemes are required for example https://electronics.stackexchange.com/questions/186168/how-to-convert-uart-voltage-from-5v-to-3-3v. I would prefer to use a ready-made USB-TTL adapter or find another adapter that has a chip powered by a voltage of +3.3V.

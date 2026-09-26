---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-39-2
title: "arborescence-des-pages-39"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-39.md
source_anchor: ""
source_lines: [186, 350]
sha256: c23e4f3fc749e6865bf655db96e45e9f8b358ae3ba290f059a213bab35f4d733
---

# arborescence-des-pages-39

```
[admin@1p_DUT_wAP ac] /> :global test {{"11";"12";"13"};{"21";"22";"23"}}   
[admin@1p_DUT_wAP ac] > :set ($test->1->1) "22_changed"
[admin@1p_DUT_wAP ac] > :put [($test->1->1)]           
22_changed
[admin@1p_DUT_wAP ac] > :environment print  
test={{"11"; "12"; "13"}; {"21"; "22_changed"; "23"}}
```
### Read value of global variable defined in other script

Lets say we have one script that declares variable and sets the value:

```
/system script add name=script1 source={
  :global myVar "hello!"
}
```
And we want to write the value of that variable in log from another script. Simply adding `/log info $myVar` will fail to return correct value, because second script does not know anything about variables defined in another scripts. To make it work properly variable need to be defined, so correct second script code is:

```
/system script add name=script2 source={
  :global myVar;
  :log info "value is: $myvar"
}
```
### Accessing global variable from function

Logically you would think that globally defined variables should be accessible in functions too, but that is not really the case. Lets see an example:

```
:global myVar "test"
:global myFunc do={
  :put "global var=$myVar"
}
[$myFunc]
```
Output is:

global var=

So obviously global variable is not accessible directly. To make it work we need do declare global variable inside the function:

```
:global myVar "test"
:global myFunc do={
  :global myVar;
  :put "global var=$myVar"
}
[$myFunc]
```
Output:

global var=test

### Running function from another function

Same as above applies also to functions. If you want to run function from another function then it need to be declared.

```
:global test do={
  :return ($1 + 1)
}
:global testtest do={
  :local x 5
  :local y [$test $x]
  :put "typeof = $[:typeof $y]"
  :put "testets_res=$y"
}
```
Code above will not work as expected, output will be:

typeof = nil
testets_res=

To fix this we need to declare global "test" in "testtest" function

```
:global testtest do={
  :global test
  :local x 5
  :local y [$test $x]
  :put "typeof = $[:typeof $y]"
  :put "testets_res=$y"
}
```
### Always use unique variable names

One of the most common scripting mistakes that most users are doing is not using unique varible names, for example, variable defined in function has the same name as globally defined variable, which leads to unexpected result:

```
:global my2 "123"
:global myFunc do={ :global my2; :put $my2; :set my2 "lala"; :put $my2 }
$myFunc my2=1234
:put "global value $my2"
```
Output will be:

1234
lala
global value 123


Another common case is when user defined variable have the same name as RouterOS built in variable, for example, we want to print route with dst address defined in variable:

[admin@1p_DUT_wAP ac] /ip route> :global "dst-address" "0.0.0.0/0"
[admin@1p_DUT_wAP ac] /ip route> print where dst-address=$"dst-address" 
Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, 
B - blackhole, U - unreachable, P - prohibit 
 #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
 0 ADS  0.0.0.0/0                          10.155.136.1              1
 1 ADC  10.155.136.0/24    10.155.136.41   ether1                    0

Obviously result is not as expected, simple solution, use unique variable name:

[admin@1p_DUT_wAP ac] /ip route> :global myDst "0.0.0.0/0"
[admin@1p_DUT_wAP ac] /ip route> print where dst-address=$myDst             
Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, 
B - blackhole, U - unreachable, P - prohibit 
 #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
 0 ADS  0.0.0.0/0                          10.155.136.1              1

### Get values from looped interactive commands like "monitor"

Frequently asked question s how to get values in script returned by, for example, monitor command? First problem with such commands is that they are running infinitely until user action is applied, obviously you cannot do that from script. Instead you can run with additional parameter once, it will allow to execute command only once and stop. Another problem is getting variable value sin script, there is no as-value, there is no get, but they have do. What it does is allows to access variables returned by the command as in example below:

```
[admin@1p_DUT_wAP ac] /interface> monitor-traffic ether1 once do={:global myBps $"rx-bits-per-second" }
...
[admin@1p_DUT_wAP ac] /interface> :environment print 
myBps=71464
```
### Get file content received by fetch tool

Fetch tool allows for ease of use downloading file content into memory and allowing to access this data by script. To make it work use as-value parameter and output=user:

[admin@rack1_b34_CCR1036] > :put ([/tool fetch ftp://admin:@10.155.136.41/test.txt
 output=user as-value ]->"data")
my file content

### Check script permissions

Lets say we have a script that creates and writes content to the file:

```
/system script add name=script1 policy=ftp,read,write source={
       /file print file=test;
       /file set test.txt content="my content"
}
```
Now lets add scheduler that will try to execute this script:

/system scheduler
add interval=10s name=test on-event=script2 policy=read,write

So now we wait 10 seconds, file not created, we wait another 10 seconds and still no file. What is going on? If you look closely script requires policy "ftp", to create a file, but scheduler has only "read" and "write" policies, so script will not be executed. Fix is to set scheduler to run with correct policies "read,write,ftp".

This applies also if you are trying to run script from netwatch, ppp on event and so on, which are limited to specific policies "read,write,test,reboot", so you will not be able to run advanced scripts that creates backups, creates files and so on.

Limitation could be fixed by using dont-require-permissions, but be very careful, read below.

### Be careful when using dont-require-permissions

It is possible to set script with dont-require-permissions parameter. Basically it allows anyone without adequate permissions to execute the script. For example, if script has policies "read,write,test,sensitive", but user or application that executes the script has less, for example, "read,write", then by setting `dont-require-permissions=yes` will allow to run script anyway.

This could potentially allow to change sensitive information using script even if user doe snot have enough permissions.

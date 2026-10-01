---
id: collect-261001-general-networking/general-networking/questions-40056-which-ip-address-is-being-referred-to-in-this-situation-6f9d048d
title: "questions-40056-which-ip-address-is-being-referred-to-in-this-situation-6f9d048d"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: ["2018-02-19"]
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-40056-which-ip-address-is-being-referred-to-in-this-situation-6f9d048d.md
source_anchor: ""
source_lines: [1, 14]
sha256: e42ad96e1857b6f12aae0a8bbd8b03d52f8c129519f1f229e8e80848c95d0591
---

# questions-40056-which-ip-address-is-being-referred-to-in-this-situation-6f9d048d

I manage a small event venue space, and a client for an upcoming event recently asked me for our IP address. He's planning on setting up a few stations (laptops) to do sales and collect money for fundraising. These laptops will be connected to our wireless internet. Apparently he wants to know our IP address "so that multiple sales from your location that day aren’t flagged as spam." Can anyone please tell me which IP address he is referring to and how I can find that IP address for him? Our ISP is Webpass, and we use Cisco Meraki MX80 with Meraki MR34 access points. Thanks
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you can provide and accept your own answer.Ron Maupin– Ron Maupin ♦2018-02-19 02:59:27 +00:00Commented Feb 19, 2018 at 2:59
3 Answers 3
I think your client is asking for your public IP address (static/Dynamic) provided by your ISP.Please check the authenticity of the request before providing the information(IP address) to your client.
If your sales process includes a call to a third party payment gateway/Banking API then it makes sense.
Sometimes when multiple clients with the same Public IP address send back to back queries to their API/services then the Automatic security routines/algorithms will start classifying them as fishy/trigger captachas or even block the IP address altogether. You client might have foreseen this problem and taking actions to prevent the problem by white listing the IP address upfront.
Just Google "what is my public IP address" You will get to know the IP address.
If you can also access the ISP provided Modem's GUI from your PC then also check it from the Web GUI as well.The IP address must be listed on WAN IP address section.
What he needs to know is the publiv IP address(es) that will be used when clients on your wifi network make requests to services on the Internet so that he can whitelist them in his abuse-control system.
Depending on your network setup that may be an individual IP address or it may be a range of addresses.
If you are sure you only have a single internet IP address you can go to a "what is my IP" type site. If you have a pool of addresses you should have a record somewhere of what that pool is.
Put simply, if you are on the same wireless network the vendor will be using and you search "what is my ip" in a Google search window, you will get a reply with your public IP (also called the outbound NAT IP. Hopefully there is only one and not a range.)
There are other sites which will provide you a similar response such as https://www.whatismyip.com/ but the Google search is a nice easter egg.

---
id: collect-261001-huawei/huawei/questions-849428-f84beb1e
title: "questions-849428-f84beb1e"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "research"]
source: docs/RAG/collect-261001-huawei/questions-849428-f84beb1e.md
source_anchor: ""
source_lines: [1, 16]
sha256: bce54649b25566af977f427a29c56ae3c000631ffeea8a3c1fae58dd79bdd89c
---

# questions-849428-f84beb1e

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have ubuntu 14 server that runs apache and has huawei router.
Recently I purchased and installed a SSL certificate for my domain (say www.abc.com) according to guide here.
Then I tested my ssl cert using their tool. and it gives an error as common name not matched. Tool say common name as cpe.huawei.com
So I also checked using command openssl s_client -connect www.abc.com:443
And It shows Secure Renegotiation not support. And server certificate retrieve here when I decoded using a tool show common name as cpe.huawei.com.
So I know some another certificate run or cached on the server. But I search every general directory and couldn't find another cert. Also, my apache error log file has no error log regarding this issue.
I need help in overriding new cert over all other certs and remove any cached ssl cert on my server.
"CPE" is a telco abbreviation for "Customer Premise Equipment", i.e. anything that they control but that they put on the customer site. In this case, it's going to refer to your router (it also commonly used for an ONT).
Your huawei router is intercepting requests directed to your public IP address on port 443. My guess is that if you were to browse to https://[your public ip] from outside your network, you will get a certificate error. Click through the certificate error and you'll either see the login page for your router, or some huawei API endpoint.
Your correct course of action here is to contact your ISP or huawei's support channel and get their assistance in resolving this, as they are hijacking port 443 on your connection.

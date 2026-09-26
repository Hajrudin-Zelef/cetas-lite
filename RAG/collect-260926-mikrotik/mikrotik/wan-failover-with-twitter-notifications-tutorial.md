---
id: collect-260926-mikrotik/mikrotik/wan-failover-with-twitter-notifications-tutorial
title: "wan-failover-with-twitter-notifications-tutorial"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["consumer"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/wan-failover-with-twitter-notifications-tutorial.md
source_anchor: ""
source_lines: [1, 120]
sha256: 05679c2ced33c1c65800bdaa0d5456068e41956e8727159805606cc1856c930d
---

# wan-failover-with-twitter-notifications-tutorial

**Introduction**

In some scenarios, you want to be able to failover to another WAN in case your main link goes down.

I prepared this tutorial to explain how to achieve that with RouterOS and the ability to send a notification using Twitter on Failover events.

**How it works**

Basically, the script will ping Google DNS 8.8.8.8 every 5 minutes using the main route (wan1), if it fails to ping, the script will change the distance of the main route by setting the value to 3 or any value higher than the second route (wan2). If the link is back, it will change the distance back to the original value.

**Setup**

- Configure two WAN interfaces in the router. In my case I have them on ether1 and ether6 ports as the below diagrams.



- Add two routes for both WANs with distance of 1 for the main WAN link and 2 for the backup link

- Then, created a new route entry that you will be using to test the internet connection on wan1 as it is your default wan connection and set the Distance to 100 and Route Mark to wan1_test


**Pre-requisites**

**Create a separate Twitter account for notifications then create a Twitter App**

- Fill the app creation form and create your app
 
 
- From Keys and Access Tokens, write down your Consumer Key, and Consumer Secret

- From Keys and Access Tokens → Your Access Token, Click Crate my access token

- From Keys and Access Tokens → Your Access Token, write down your generated Access Token, and Access Token Secret


**Create a web application to consume twitter API**

For that, I used Visual Studio and ASP.NET MVC web application. Feel free to develop the application with any programming language you are comfortable with.

- Define string variables for Consumer Key, Consumer Secret, Access Token, and Access Token Secret

```
string consumerKey = "xXxXxXxXxXxXxXxXxXxXxXxX"; // The application's consumer key
        string consumerSecret = "xXxXxXxXxXxXxXxXxXxXxXxX"; // The application's consumer secret
        string accessToken = "xXxXxXxXxXxXxXxXxXxXxXxX"; // The access token granted after OAuth authorization
        string accessTokenSecret = "xXxXxXxXxXxXxXxXxXxXxXxX"; // The access token secret granted after OAuth authorization
```




- Install NuGet Package Spring.Social.Twitter version 2.0.0-M1
 
 
- Create a new function in your controller, this function will be called from RouterOS script every time failover happens and that function will tweet the failover details.

```
public ActionResult UpdateConnectionStatus(string currentConnection)
        {
            ITwitter twitter = new TwitterTemplate(consumerKey, consumerSecret, accessToken, accessTokenSecret);
            if (currentConnection != null)
            {
                twitter.TimelineOperations.UpdateStatusAsync(String.Format("Internet service is now provided by {0} network", currentConnection));
                return new HttpStatusCodeResult(HttpStatusCode.OK);
            }
            else
            {
                return new HttpStatusCodeResult(HttpStatusCode.NotFound);
            }
        }
```




- Compile and publish the application on IIS

**Scripting RouterOS**

- Create a new script to check the internet connection on your default wan using the test route that we created in the above step as following

```
:global currentConnection;
:local GW1 ether1-wan1
:local RouteNameGW1 wan1
:local RouteTableTest wan1_test
:local PingTarget 8.8.8.8
:local PingResult
:set PingResult [ping $PingTarget count=5 routing-table=$RouteTableTest interface=$GW1]
:if ( $PingResult > 2 ) do={
/ip route set [/ip route find where comment=$RouteNameGW1] distance=1;
:if ( $currentConnection = "Mobily 4G" ) do={
:set currentConnection "STC Fiber";
/tool fetch address=<ip address where you host your internal app> src-path="/currentConnection=$currentConnection" mode=http keep-result=no url="http://<ip address where you host your internal app>/Home/UpdateConnectionStatus\?currentConnection=STC%20Fiber";
}
:if ( [ :typeof $currentConnection ] = "nothing" ) do={
:set currentConnection "STC Fiber";
}
:set PingResult 0;
} else={
/ip route set [/ip route find where comment=$RouteNameGW1] distance=3;
:if ( $currentConnection = "STC Fiber" ) do={
:set currentConnection "Mobily 4G";
/tool fetch address=<ip address where you host your internal app> src-path="/currentConnection=$currentConnection" mode=http keep-result=no url="http://<ip address where you host your internal app>/Home/UpdateConnectionStatus\?currentConnection=Mobily%204G";
}
:if ( [ :typeof $currentConnection ] = "nothing" ) do={
:set currentConnection "Mobily 4G";}
:set PingResult 0; }
```




- Schedule the script to run every 5 minutes
 
 
- Follow the twitter account, that you use to tweet these notifications, from your personal account and set SMS notifications on that account
 
 
- Now whenever my connection fails over the backup link, I get the following notification on my phone.

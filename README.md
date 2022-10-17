# Description
Get the latest IP of the [vpngate](https://vpngate.net/) website, save or update to MySQL database.
* Obtain only L2TP/IPsec-enabled IPs
# Getting Started
1. You need a server in the true free network world.
   * [A Powerful Server From Oracle Cloud — Always Free](https://levelup.gitconnected.com/a-powerful-server-from-oracle-cloud-always-free-cbc73d9fbfee?gi=29b7a0d35b2f)
2. You need a MySQL server you can access easily.
   * You can use [ngrok](https://ngrok.com/) as a passthrough to access your local MySQL.
3. Run the application by crontab.
   * `*/5 * * * * /home/gate -DSN=<username>:<pw>@tcp(<HOST>:<port>)/<dbname>`
# How to Connect
* [Windows](https://www.vpngate.net/en/howto_l2tp.aspx#windows)
* [Mac OS X](https://www.vpngate.net/en/howto_l2tp.aspx#mac)
* [iPhone / iPad (iOS)](https://www.vpngate.net/en/howto_l2tp.aspx#ios)
* [Android](https://www.vpngate.net/en/howto_l2tp.aspx#android)
# Experience
1. Install the TablePlus app on your iPhone.
2. Access your MySQL via TablePlus.
3. Open the "gate" table order by create_time desc.
4. Select an IP your like, and set the IP according to the method described above.
5. Access Google, then welcome to the real web.
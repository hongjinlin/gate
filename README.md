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
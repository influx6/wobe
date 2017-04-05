WobeEcho App
============

Wobe demonstrates a simple service deployed to both [Heroku](https://heroku.com) and [Now](https://zeit.co/now) using docker.

Install
-------

-	Have `heroku` setup on your system as described in https://devcenter.heroku.com/articles/getting-started-with-go

-	Have `now` setup on your system as described in https://zeit.co/now

Heroku
------

Reverse Only Service: https://wobe.herokuapp.com/

To be able to deploy to the [Heroku](https://heroku.com) platform simply do:

-	Navigate to `deploy/wobe-heroku`

-	Run `make` on the terminal

Now
---

Reverse Only Service: https://wobe-now-jzokyjlqdt.now.sh/

To be able to deploy to the [Now](https://zeit.co/now) platform simply do:

-	Navigate to `deploy/wobe-now`

-	Run `make create` and `make` on the terminal

Testing
-------

-	Test App by sending the following through the terminal

```bash
curl -v $HOST:$PORT/reverse -d '{"input": "bomba"}'
```

-	Test `echo` endpoint by switching into the `echo` branch and deploy once again

```bash
curl -v $HOST:$PORT/reverse -d '{"input": "bomba"}'
curl -v $HOST:$PORT/echo -d '{"input": "bomba"}'
```

Note that `$HOST` is a placeholder for the URL address of the deployed app and that `$PORT` is a placeholder for the port deployed to on the platform used.

-	The Now deployment by default uses port `7080`

-	You will need to locate the `$PORT` data related to the heroku deployment from your Heroku App dashboard.

Scaling
-------

### Now Cloud Platform

Based on giving plans on the Now platform, you can do the following to scale:

-	Nothing, Now automatically starts up multiple instances of your container based on your docker image and scales with them based on the load on requests coming in.

-	The higher grade plans expand the capability which that is provided and the level of support possible.

Only issue to watch out for is:

-	If the amount of traffic that requires a lot of scale exceeds available bandwidth limit then you generally will have issues on scaling, so always check the plans.

### Heroku Cloud Platform

Based on giving plans on the Heroku platform, you can do the following to scale:

-	Depending on dyno types scaling has limits on Heroku so, ensure to check https://devcenter.heroku.com/articles/scaling#scaling-limits.

-	Move subscription of app from a free-tier dyno to a professional-tier dyno which provides quick access to the manual scaling ability where you can add more dynos to handle concurrent requests managed, this allows you to scale up or down manually through the App Dashboard.

-	If using the performance-tier dyno, then auto-scaling can be configured for the apps associated with this tier. See more at https://devcenter.heroku.com/articles/scaling

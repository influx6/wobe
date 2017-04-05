Wobe App
========

Wobe demonstrates a simple service deployed on heroku using the experiemental docker service builds.

Install
-------

-	Have heroku setup on your system as describe in https://devcenter.heroku.com/articles/getting-started-with-go

Heroku
------

To be able to deploy to the [Heroku](https://heroku.com) platform simply do:

-	Navigate to `deploy/wobe-heroku`

-	Run `make` on the terminal

Now
---

Server: https://wobe-now-jzokyjlqdt.now.sh/

To be able to deploy to the [Now](https://zeit.co/now) platform simply do:

-	Navigate to `deploy/wobe-now`

-	Run `make` on the terminal

Testing
-------

-	Test App by sending the following through the terminal

```bash
curl -v localhost:$PORT/reverse -d '{"input": "bomba"}'
curl -v localhost:$PORT/reverse -d '{"input": "bomba"}'
```

Note that `$PORT` is a placeholder for the port deployed to on the platform used.

-	The Now deployment by default uses port `7080`

-	You will need to locate the `$PORT` data related to the heroku deployment from your Heroku App dashboard.

Scaling
-------

### Now Cloud Platform

Based on giving plans on the Now platform, you can do the following to scale:

### Heroku Cloud Platform

Based on giving plans on the Heroku platform, you can do the following to scale:

CREATEAPP := heroku create wobe
DEPLOYAPP := heroku container:push web --app wobe
CLEANAPP := heroku destroy wobe
STARTAPP := heroku ps:scale web=1 wobe

clean:
	$(CLEANAPP)

deploy:
	$(DEPLOYAPP)

start:
	$(STARTAPP)

create:
	$(CREATEAPP)
	$(start)

$(deploy)

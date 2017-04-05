CREATEAPP := heroku create wobe
DEPLOYAPP := heroku container:push web --app wobe
CLEANAPP := heroku destroy wobe

deploy:
	$(DEPLOYAPP)

create:
	$(CREATEAPP)

clean:
	$(CLEANAPP)

$(deploy)

CREATEAPP := heroku create wobe
DEPLOYAPP := heroku container:push web
CLEANAPP := heroku destroy wobe

deploy:
	$(CREATEAPP)
	$(DEPLOYAPP)

clean:
	$(CLEANAPP)

$(deploy)

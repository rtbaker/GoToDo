default: develop

fixcs:
	vendor/bin/phpcbf

test:
	vendor/bin/phpcs
	vendor/bin/phpstan
	bin/phpunit

coverage:
	

start-develop:
	docker compose -f development/docker-compose-dev.yml up -d

stop-develop:
	docker compose -f development/docker-compose-dev.yml stop
	
develop:
	symfony server:start --allow-all-ip --port=9090


default: develop

fixcs:
	vendor/bin/phpcbf

test:
	vendor/bin/phpcs
	vendor/bin/phpstan --memory-limit=512M
	bin/phpunit

coverage:
	

start-develop:
	docker compose -f development/docker-compose-dev.yml up -d

stop-develop:
	docker compose -f development/docker-compose-dev.yml stop
	
develop:
	XDEBUG_MODE=debug symfony server:start --allow-all-ip --port=9090


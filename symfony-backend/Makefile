default: develop

fixcs:
	vendor/bin/phpcbf

test:
	vendor/bin/phpcs
	vendor/bin/phpstan --memory-limit=512M
	bin/phpunit

setup-test-db:
	bin/console --env=test doctrine:database:drop --if-exists --force
	bin/console --env=test doctrine:database:create --if-not-exists
	bin/console --env=test doc:mig:mig -n
	bin/console --env=test doctrine:fixtures:load -n
	
coverage:
	

start-develop:
	docker compose -f development/docker-compose-dev.yml up -d

stop-develop:
	docker compose -f development/docker-compose-dev.yml stop
	
develop:
	XDEBUG_MODE=debug symfony server:start --allow-all-ip --port=9090


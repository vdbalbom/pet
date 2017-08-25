run:
	docker-compose stop && \
	docker-compose rm && \
	docker-compose build pet-run && \
	docker-compose run pet-run

test:
	docker-compose stop && \
	docker-compose rm && \
	docker-compose build pet-test && \
	docker-compose run pet-test

FROM rabbitmq:3.9.13-management-alpine
RUN rabbitmq-plugins enable --offline rabbitmq_mqtt rabbitmq_federation_management rabbitmq_stomp

EXPOSE 8080
FROM maven:3.8-eclipse-temurin-19-alpine
COPY target/order-service-0.0.1-SNAPSHOT.jar order-service-0.0.1.jar
EXPOSE 8080
ENTRYPOINT ["java","-jar","/order-service-0.0.1.jar"]

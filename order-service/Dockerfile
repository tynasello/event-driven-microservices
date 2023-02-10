# Multi Stage build

# Base Image and name stage as "builder"
FROM maven:3.8-eclipse-temurin-19-alpine AS builder

# Create App Directory inside our container
WORKDIR /usr/app/src/

# Copy files
COPY src ./
COPY pom.xml ../

RUN mvn -f /usr/app/pom.xml clean package

#### 2nd Stage ####

FROM eclipse-temurin:19-jdk-alpine

WORKDIR /usr/lib/

# Copy the Jar from the first Stage (builder) to the 2nd stage working directory
COPY --from=builder /usr/app/target/order-service-0.0.1-SNAPSHOT.jar ./order-service.jar

# Expose the port to the inner container communication network
EXPOSE 8080

# Run the Java Application
ENTRYPOINT [ "java","-jar","/usr/lib/order-service.jar"]

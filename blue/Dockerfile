FROM openjdk:8-jdk

ADD target/blue-0.1.0.jar /app.jar

ENV JAVA_OPTS=""

ENTRYPOINT [ "bash", "-c", "java -jar /app.jar" ]

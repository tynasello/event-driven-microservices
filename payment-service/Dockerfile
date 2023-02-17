FROM python:3.10

EXPOSE 8081 

WORKDIR /app
COPY . .

RUN pip install -r requirements.txt

ENTRYPOINT [ "python", "src/main.py" ]

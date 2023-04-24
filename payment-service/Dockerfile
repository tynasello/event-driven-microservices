FROM python:3.10

EXPOSE 8083

WORKDIR /app
COPY . .

RUN pip install -r requirements.txt

ENTRYPOINT [ "python", "src/main.py" ]

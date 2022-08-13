FROM python:3.7.13-alpine
WORKDIR /ml_engine
EXPOSE 5002
COPY . .
RUN apk add gcc musl-dev python3-dev libffi-dev openssl-dev
RUN pip install -r requirements.txt
CMD ["python", "app.py"]
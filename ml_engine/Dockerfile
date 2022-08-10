FROM python:2.7
WORKDIR /ml_engine
EXPOSE 5002
COPY . .
RUN pip install -r requirements.txt
CMD python app.py
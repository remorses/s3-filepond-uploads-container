FROM python:3.7.4-slim-stretch

WORKDIR /src

ENV PYTHONUNBUFFERED=1

COPY requirements.txt /src/

RUN pip install -r requirements.txt

COPY . /src/

CMD python -m src
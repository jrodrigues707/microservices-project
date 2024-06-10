from kafka import KafkaProducer
import json
import time
import random
from datetime import datetime

def generate_message():
    return {
        "transaction_id": random.randint(1000, 9999),
        "user_id": random.randint(1, 100),
        "product_id": random.randint(1, 50),
        "quantity": random.randint(1, 10),
        "price": round(random.uniform(10.0, 500.0), 2),
        "timestamp": datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    }

producer = KafkaProducer(
    bootstrap_servers='kafka:9092',
    value_serializer=lambda v: json.dumps(v).encode('utf-8'))

while True:
    message = generate_message()
    producer.send('my-topic', message)
    time.sleep(1)
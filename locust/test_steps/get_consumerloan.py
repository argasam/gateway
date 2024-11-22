import time
from locust import HttpUser, TaskSet, task, between

class GetConsumerLoanStep(TaskSet):
    wait_time = between(1, 5)

    @task
    def get_consumerloan(self):
        self.client.get("/consumerloan")

class GetConsumerLoanTest(HttpUser):
    tasks = [GetConsumerLoanStep]
    wait_time = lambda self: 1  # Add wait_time as needed
    host = "http://localhost:8080/gateway"  # Default to None


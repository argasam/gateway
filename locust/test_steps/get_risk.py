import time
from locust import HttpUser, TaskSet, task, between


class GetRiskStep(TaskSet):
    wait_time = between(1, 5)

    @task
    def get_risk(self):
        self.client.get("/risk")

class GetRiskTest(HttpUser):
    tasks = [GetRiskStep]
    wait_time = lambda self: 1  # Add wait_time as needed
    host = "http://localhost:8080/gateway"  # Default to None
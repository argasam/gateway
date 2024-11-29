import time
from locust import HttpUser, TaskSet, task, between


class GetRiskStep(TaskSet):
    @task
    def get_risk(self):
        self.client.get("/risk")

class GetRiskTest(HttpUser):
    tasks = [GetRiskStep]
    wait_time = between(1,5)
    host = "http://localhost:8080/gateway"
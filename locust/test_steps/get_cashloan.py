import time
from locust import HttpUser, TaskSet, task, between


class GetCashLoanStep(TaskSet):
    wait_time = between(1, 5)

    @task
    def get_cashloan(self):
        self.client.get("/cashloan")

class GetCashLoanTest(HttpUser):
    tasks = [GetCashLoanStep]
    wait_time = lambda self: 1  # Add wait_time as needed
    host = "http://localhost:8080/gateway"  # Default to None
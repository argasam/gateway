import time
from locust import HttpUser, TaskSet, task, between

class GetConsumerLoanStep(TaskSet):
    @task
    def get_consumerloan(self):
        self.client.get("/consumerloan")

class GetConsumerLoanTest(HttpUser):
    tasks = [GetConsumerLoanStep]
    wait_time = between(1,5)  
    host = "http://localhost:8080/gateway"  


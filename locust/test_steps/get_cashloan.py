from locust import HttpUser, TaskSet, task, between


class GetCashLoanStep(TaskSet):
    @task
    def get_cashloan(self):
        self.client.get("/cashloan")

class GetCashLoanTest(HttpUser):
    tasks = [GetCashLoanStep]
    wait_time = between(1,5) 
    host = "http://localhost:8080/gateway"
from locust import HttpUser, TaskSet, task, between

class GetSearchConsumerLoanStep(TaskSet):
    @task
    def get_searchconsumerloan(self):
        with self.client.get("/consumerloan", catch_response=True) as response:
            if response.status_code == 200:
                cookies = response.cookies.get_dict()
            else:
                response.failure("Failed to fetch /consumerloan")
        # print("Cookies received:", cookies)

        with self.client.get("/consumerloan/search", cookies=cookies, catch_response=True) as search:
            # print(search.status_code)
            if search.status_code != 200:
                search.failure("Failed to search consumer loan")
                return

class GetSearchConsumerLoanTest(HttpUser):
    tasks = [GetSearchConsumerLoanStep]
    wait_time = between(1, 5)
    host = "http://localhost:8080/gateway" 
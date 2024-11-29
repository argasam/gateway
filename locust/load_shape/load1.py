from locust import LoadTestShape

class StagesShapeNo1(LoadTestShape):

    stages = [
        {"duration": 15, "users": 30, "spawn_rate": 2},
        {"duration": 30, "users": 30, "spawn_rate": 1},
    ]

    def tick(self):
        run_time = self.get_run_time()

        total_time = 0
        for stage in self.stages:
            total_time += stage["duration"]
            if run_time < total_time:
                tick_data = (stage["users"], stage["spawn_rate"])
                return tick_data

        return None
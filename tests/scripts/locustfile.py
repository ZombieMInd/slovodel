from locust import HttpUser, task, between

class LoggerUser(HttpUser):
    wait_time = between(0.5, 1)
    message = {
        "user_uuid": "bd4cb967-a824-4ada-ad75-f74820793819",
        "timestamp": 2987428975,
        "events": [{
            "event_type": "string",
            "event_txt": "string"
        },{
            "event_type": "string",
            "event_txt": "string"
        },{
            "event_type": "string",
            "event_txt": "string"
        }]
    }

    @task
    def messagesJSON(self):
        self.client.post("/log", json=self.message)


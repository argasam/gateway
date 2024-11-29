# !Before doing anything
### Create Venv
```bash
python3 -m venv venv2
source venv2/bin/activate
pip install -r requirements.txt
```

# command to run headless (Non-UI)
```bash
locust -f "replace w/ direction to file","replace with custom load" --headless --html performance_reports/results_$(date +%Y%m%d%H%M%S).html
```

### example
```bash
locust -f locust/test_steps/get_consumerloan_search.py,locust/load_shape/load1.py --headless --html locust/performance_reports/results_$(date +%Y%m%d%H%M%S).html
```

# command to run with UI
```bash
locust -f locust --class-picker
```

generate-backend:
	openapi-generator generate -i openapi.yaml -g nodejs-express-server -o backend


run-solver:
	minizinc --solver gurobi  backend/minizinc/scheduler.mzn backend/minizinc/data.dzn
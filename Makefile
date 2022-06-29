generate-backend:
	openapi-generator generate -i openapi.yaml -g go-server -o backend

generate-frontend:
	openapi-generator generate -i openapi.yaml --output frontend/src/openapi

run-solver:
	minizinc --solver gurobi  backend/minizinc/scheduler.mzn backend/minizinc/data.dzn

run-jaev-solver:
	minizinc --solver gurobi  backend/minizinc/scheduler_jaev.mzn backend/minizinc/data_jaev.dzn

run-backend:
	cd backend; go run main.go

run-frontend:
	cd frontend; npm run dev

run:
	make run-backend; make run-frontend

run-greedy-search:

	java -cp "/Users/dennismaes/Library/Mobile Documents/com~apple~CloudDocs/School/2021-2022/Thesis/applicatie/thesis/target/scheduler-1.0.0.jar:/Users/dennismaes/.m2/repository/org/xerial/sqlite-jdbc/3.34.0/sqlite-jdbc-3.34.0.jar:/Users/dennismaes/.m2/repository/org/javatuples/javatuples/1.2/javatuples-1.2.jar" Main
generate-backend:
	openapi-generator generate -i openapi.yaml -g go-server -o backend

run-solver:
	minizinc --solver gurobi  backend/minizinc/scheduler.mzn backend/minizinc/data.dzn

run-backend:
	cd backend; go run main.go

run-frontend:
	cd frontend; npm run dev

run:
	make run-backend; make run-frontend
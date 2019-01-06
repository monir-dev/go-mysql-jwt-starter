GCC=go
GCMD=run
GPATH=main.go

run:
	$(GCC) $(GCMD) $(GPATH)

install:
	make install_env
	make install_db
	make install_routes

install_env:
	dep ensure -add github.com/joho/godotenv
install_db:
	dep ensure -add github.com/go-xorm/xorm
install_routes:
	dep ensure -add github.com/gorilla/mux



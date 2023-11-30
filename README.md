## GinkGo Sandbox

Examples of using ginkgo libraries for unit tests. Mocks are generated using mockery. 


#### Used
- [SQLMock](https://github.com/DATA-DOG/go-sqlmock)
- [Mockery](https://vektra.github.io/mockery/v2.38/)
- [Ginkgo](https://github.com/onsi/ginkgo)
- [apitest](https://github.com/steinfletcher/apitest)
- [gomega](https://github.com/onsi/gomega)

### Running
```sh
go run ./...
# or 
ginkgo ./...
```

### Generating mocks
Run `mockery`
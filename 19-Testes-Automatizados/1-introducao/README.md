### Para rodar os testes em um pacote, modo verboso/detalhista, porcentagem de cobertura:
- go test 
- go test ""Nome do arquivo"
- go test -v
- go test -v --cover
  
### Para rodar todos os testes, modo verboso/detalhista, porcentagem de cobertura em varias pastas e pacotes:
- go test ./...
- go test ./... -v
- go test ./... --cover
- go test ./... --cover -v

### Para salvar os testes em um arquivo:
- go test --coverprofile cobertura.txt 

### Para ler o arquivo de testes:
- go tool cover --func=cobertura.txt 

### Para ler o arquivo de testes em html de froma mais visual:
- go tool cover --html=cobertura.txt 
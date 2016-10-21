# Trovilo - 0.0.1
Admin Finder Wrote in Golang

##Uso
### Opções(Flag)

Opção               |  Descrição | Aplicação       
------------------- |  --------- | ---------
`-url=`          | Seta a url | http://site.com/
`-Tor=`          | Seta o tor() | true or false
`-Wordlist`      | Seta o Caminho para Wordlist | wordlist/word.txt  



* Exemplos usando `-action=`:
AQUI VEM TEUS EXEMPLOS QUE BUGUEI

* Exemplos usando `-limit=`:
AQUI VEM TEUS EXEMPLOS QUE BUGUEI

###Verificador 

Uso : `./gonion -validate=<nome do arquivo com os links>`  

###Buscador (by RANGE)

Uso : `./gonion -action=generate <quantidade de links a ser gerada>`

**Obs: É recomendado deixar o comando rodando em background, para isso usar** `&` :

Uso : `./gonion -action=generate <quantidade de links a ser gerada> &`

###Buscador (by Crawling)
Para esse algoritmo você terá de passar um link ou um arquivo de links como parametro.

* Exemplos:
```
./gonion -action=bar :txt_de_links.txt  
./gonion -action=bar link
```

## Funcionamento
### Generate
_O gerador do Gonion cria links aleatórios na faixa de 22 caracteres e os requisita utilizando o protocolo SOCKS para resolver os enderaços .onion_
### Bar
_O algoritmo do Bar usa scraping para buscar no(s) site(s) passado(s) como parametr outros links e usa o mesmo como parametro para uma nova busca. Para usar o Bar é necessário setar o limite de Goroutines, uma quantidade pequena pode resultar num mal funcionamento do programa é recomendado que use algo a partir de 500 Goroutines._
### Verificador  
_O verificador do Gonion lê os arquivos de link e os testa utilizando o protocolo SOCKS para resolver os enderaços .onion_

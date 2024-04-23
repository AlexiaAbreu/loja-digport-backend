# **Considerações Aula 6**


#### adicionar lista de todos os produtos ao sclice "produtosEscolhidosPeloNome" que correspondem ao nome específico 
 
``` 
 produtosEscolhidosPeloNome = append(produtosEscolhidosPeloNome, produto) 
```

---

#### Se nenhum produto com o nome específico foi encontrado, retorne todos os produtos 

```
 if len(produtosEscolhidosPeloNome) == 0 {
      return produtos
 }
```
---

## **Uso da Biblioteca externa - gorilla/mux**      


>  O roteador ("r") é uma **estrutura de dados** fornecida pelo pacote mux que  permite **definir rotas** e **associar manipuladores** a essas rotas.     

> Ao criar um novo roteador usando mux.NewRouter(), você está criando uma instância do roteador que pode ser usada para definir e gerenciar as rotas do seu aplicativo. Essa instância do roteador, frequentemente armazenada em uma variável como "r" por convenção, é usada para adicionar rotas ao servidor HTTP e especificar os manipuladores para essas rotas.

>No exemplo fornecido, "roteador" é usado para registrar os manipuladores para as rotas "/produtos" e "/produtos/{nomeDoProduto}", **indicando quais funções devem ser chamadas quando essas rotas são acessadas** pelo cliente. Depois que todas as rotas são registradas no roteador, o roteador é passado para http.ListenAndServe() para iniciar o servidor HTTP que irá lidar com essas rotas.

---
### **Antes das minhas mudanças:**

####  Primeiro parâmetro é o endpoint que irá retornar os dados para quem chamar, segundo parâmetro é o handler em si, que pega a lista de produtos e transforma em JSON
````
	http.HandleFunc("/produtos", produtosHandler) 
	http.ListenAndServe(":8080", r)

````
### Aqui é feita a escrita do JSON
````
json.NewEncoder(writer).Encode(catalogoDeProdutos) 
````







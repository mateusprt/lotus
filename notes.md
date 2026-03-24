# Construtos da linguagem Lótus

---

# Declarações (Statements)

| Construto | Descrição | Exemplo |
|---|---|---|
| `var statement` | Declaração de variável com atribuição de valor. | `var x = 10;` |
| `return statement` | Retorna um valor dentro de uma função. | `return x + y;` |
| `block statement` | Bloco de código executado sequencialmente. | `{ var x = 5; x; }` |
| `assignment statement` | Permite reatribuir um novo valor a uma variável já declarada. | `x = 20;` |
| `break statement` | Interrompe imediatamente a execução de um loop. | `break;` |
| `continue statement` | Interrompe a iteração atual do loop e continua para a próxima. | `continue;` |

---

# Identificadores

| Construto | Descrição | Exemplo |
|---|---|---|
| `identifier` | Nome que referencia variáveis, funções ou parâmetros. | `x`, `add`, `result` |

---

# Estruturas de Dados

| Construto | Descrição | Exemplo |
|---|---|---|
| `array literal` | Estrutura ordenada de valores. | `[1, 2, 3]` |
| `hash literal` | Estrutura chave-valor (map/dicionário). | `{"name": "Monkey"}` |
| `index expression` | Acesso a elemento de array ou hash. | `arr[0]`, `user["name"]` |
| `struct literal` | Define uma estrutura composta por campos nomeados. | `var user = struct { name: "Ana", age: 30 };` |
| `member expression` | Permite acessar um campo de uma struct. | `user.name` |
| `field assignment expression` | Permite modificar diretamente o valor de um campo de uma struct. | `user.age = 31;` |
| `const statement` | Declara um identificador cujo binding não pode ser reatribuído após a inicialização. | `const PI = 3.14;` |

---

# Expressões

| Construto | Descrição | Exemplo |
|---|---|---|
| `integer literal` | Número inteiro literal. | `10` |
| `boolean literal` | Valor booleano literal. | `true`, `false` |
| `string literal` | Texto literal. | `"hello"` |
| `prefix expression` | Operador antes do operando. | `!true`, `-5` |
| `infix expression` | Operador entre dois operandos. | `5 + 5`, `x * y`, `a == b` |
| `member expression` | Expressão que acessa membros (campos) de uma struct. | `user.name` |
| `constant identifier` | Identificador associado a um valor imutável definido por `const`. | `PI` |
---

# Controle de Fluxo

| Construto | Descrição | Exemplo |
|---|---|---|
| `if expression` | Estrutura condicional que executa um bloco se a condição for verdadeira. | `if (x < y) { x } else { y }` |
| `for loop` | Estrutura de repetição com inicialização, condição e incremento. | `for (var i = 0; i < 10; i = i + 1) { print(i); }` |

---

# Funções

| Construto | Descrição | Exemplo |
|---|---|---|
| `function literal` | Define uma função anônima utilizando a palavra-chave `function`. | `function(x, y) { x + y }` |
| `call expression` | Executa uma função passando argumentos. | `add(2, 3)` |
| `struct mutation via function` | Funções podem modificar campos de structs recebidas como argumento através de atribuição de campo. | `function(u) { u.age = u.age + 1 }` |

---

# Funções Built-in

| Construto | Descrição | Exemplo |
|---|---|---|
| `len` | Retorna o tamanho de uma string ou array. | `len("hello")` |
| `first` | Retorna o primeiro elemento de um array. | `first([1,2,3])` |
| `last` | Retorna o último elemento de um array. | `last([1,2,3])` |
| `push` | Adiciona um elemento ao final de um array. | `push([1,2], 3)` |
| `print` | Imprime valores no console. | `print("hello world")` |

# Comentários
| `//` | Comentários. | `// Isso é um comentário` |
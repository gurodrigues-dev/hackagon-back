

# Arquitetura do Hackagon

Este documento descreve a arquitetura básica de um webserver (hackagon), incluindo componentes principais e suas funções.

## Visão Geral

Um webserver é um software projetado para servir conteúdo na World Wide Web. Ele aceita solicitações HTTP de clientes, como navegadores da web, e fornece respostas com páginas da web, arquivos, dados ou outros recursos, conforme necessário. No nosso caso, recebemos solicitações e devolvemos dados de um item ou série em específico.

O Hackagon, foi criado com o intuito de ser uma inovação em projetos acadêmicos da `Faculdade Impacta de Tecnologia`. Ele foi criado tal qual, um Hackaton, porém com sua principal linguagem, herdando o nome: Hacka`go`n

## Componentes Principais

### 1. Frontend

O "front-end" refere-se à parte de um aplicativo ou site que os usuários interagem diretamente. Ele engloba a interface do usuário, o design, a interação e a experiência geral do usuário. O front-end é construído utilizando tecnologias como HTML (Hypertext Markup Language), CSS (Cascading Style Sheets) e JavaScript.

Decidimos escolher o framework React, que mais nos representa de acordo com nossas habilidades, nas experiências descritas acima

React é uma biblioteca JavaScript de código aberto, mantida pelo Facebook, que é utilizada para construir interfaces de usuário (UI) interativas e escaláveis. Ela permite aos desenvolvedores criar aplicativos web com interfaces de usuário dinâmicas e reativas, tendo uma série de vantagens:

- Componentização: React é baseado em componentes. Um componente em React é uma peça reutilizável e independente de código que encapsula uma parte específica da interface do usuário. Por exemplo, um componente pode representar um botão, um formulário, uma barra de navegação, etc.

- Virtual DOM (Document Object Model): Uma das características mais poderosas do React é o uso do Virtual DOM. O Virtual DOM é uma representação virtual da estrutura de uma página da web que o React mantém na memória. Sempre que ocorre uma alteração no estado de um componente, o React compara o Virtual DOM atual com o estado anterior, identifica as diferenças e atualiza apenas as partes que foram modificadas, minimizando assim a quantidade de manipulação direta do DOM real.

- JSX (JavaScript XML): JSX é uma extensão da sintaxe JavaScript que permite escrever marcação HTML dentro do código JavaScript. Ele permite que os desenvolvedores descrevam a estrutura da interface do usuário de uma forma mais declarativa e intuitiva, integrando HTML com JavaScript.

- State Management: O estado (state) em React é uma peça fundamental. O estado representa os dados mutáveis que podem afetar o comportamento da interface do usuário. React permite gerenciar o estado de forma eficiente através do conceito de "stateful components" (componentes com estado) e "stateless components" (componentes sem estado). O estado é atualizado de forma reativa, o que significa que quando o estado de um componente muda, o React automaticamente re-renderiza o componente para refletir essas alterações.

- Lifecycle Methods: React fornece métodos de ciclo de vida que permitem aos desenvolvedores executar código em pontos específicos durante o ciclo de vida de um componente. Por exemplo, métodos como componentDidMount(), componentDidUpdate(), e componentWillUnmount() permitem aos desenvolvedores controlar o comportamento do componente em diferentes estágios do ciclo de vida.

- Gestão de Eventos: React permite a manipulação de eventos de forma semelhante ao JavaScript puro, mas com uma sintaxe um pouco diferente. Os eventos são tratados de forma eficiente e podem ser ligados diretamente aos elementos da interface do usuário.

> React é uma abordagem poderosa para a construção de interfaces de usuário interativas e dinâmicas em aplicações web. Ao utilizar conceitos como componentização, Virtual DOM, JSX, gerenciamento de estado e métodos de ciclo de vida, os desenvolvedores podem criar aplicações web escaláveis, eficientes e de fácil manutenção.

### 2. Nginx (Web Server & Proxy Reverso + Load Balancer)

Um servidor web é um software que permite armazenar, processar e servir conteúdo web para clientes, geralmente navegadores da web, mediante solicitação. Ele desempenha um papel crucial na hospedagem de sites, aplicativos web e serviços online, permitindo que os usuários acessem e interajam com recursos disponíveis na internet.

Decidimos escolher o Nginx para dar continuidade ao projeto, pela quantidade de poderio e customização sobre o server.

O Nginx é um servidor web de código aberto extremamente poderoso e versátil, conhecido por sua alta performance, escalabilidade e capacidade de lidar com um grande volume de tráfego. Aqui está uma descrição detalhada sobre o Nginx, seu poder como servidor web, o uso de proxy reverso e sua capacidade de criação e utilização de balanceadores de carga:

- Conhecido por sua eficiência e desempenho. Ele foi projetado para lidar com um grande número de conexões simultâneas e processar solicitações de forma rápida e eficaz, o que o torna ideal para lidar com o tráfego intenso da web.

- O Nginx é otimizado para uso eficiente de recursos, consumindo menos memória e processamento em comparação com outros servidores web tradicionais, como o Apache porém altamente escalável e pode ser facilmente configurado para lidar com um grande volume de tráfego e aplicativos web de alta carga.

- Nginx é frequentemente usado como um servidor proxy reverso, atuando como intermediário entre os clientes e os servidores de aplicativos. Isso permite que o Nginx direcione o tráfego para diferentes servidores com base em determinados critérios.

- O proxy reverso do Nginx pode distribuir o tráfego entre vários servidores de aplicativos para garantir uma distribuição uniforme da carga e melhorar a disponibilidade e a confiabilidade do serviço.

- Além disso, o Nginx pode ser configurado para fazer cache de conteúdo estático, reduzindo a carga nos servidores de aplicativos e melhorando os tempos de resposta para os usuários finais.

> O Nginx é uma ferramenta extremamente poderosa para servir conteúdo web, gerenciar tráfego e melhorar o desempenho e a confiabilidade de aplicativos web. Sua capacidade de atuar como servidor web, proxy reverso e balanceador de carga o torna uma escolha popular entre desenvolvedores e administradores de sistemas para hospedar e escalar aplicativos web de alto desempenho.

### 3. Backend

O backend de uma aplicação web é a parte do sistema que lida com o processamento dos dados, a lógica de negócios e a interação com o banco de dados. Enquanto o frontend lida com a interface com o usuário e a apresentação dos dados, o backend trabalha nos bastidores para garantir que a aplicação funcione corretamente e forneça os dados certos para o frontend.

A escolha da linguagem de programação para o backend depende de vários fatores, como requisitos de desempenho, facilidade de desenvolvimento, escalabilidade, entre outros. Go, também conhecida como Golang, é uma linguagem de programação criada pela Google, lançada em 2009, que tem ganhado popularidade especialmente para o desenvolvimento de sistemas de backend. Aqui estão algumas razões pelas quais usar Go pode ser vantajoso:

- Go é conhecida por sua performance excepcional. Ela é compilada para código nativo, o que resulta em execução mais rápida e eficiente em comparação com linguagens interpretadas.

- Go tem suporte nativo a concorrência por meio de goroutines e canais, o que facilita a criação de sistemas concorrentes e distribuídos de forma eficiente. Isso é particularmente útil para aplicações que precisam lidar com um grande volume de requisições simultâneas.

- Go tem uma sintaxe simples e concisa que facilita a leitura e a manutenção do código. Ela foi projetada para minimizar a quantidade de código boilerplate e para promover boas práticas de desenvolvimento.

- Go possui uma biblioteca padrão rica que inclui pacotes para realizar tarefas comuns, como manipulação de strings, I/O, criptografia, e muito mais. Isso reduz a dependência de bibliotecas de terceiros e simplifica o desenvolvimento de aplicações.

- A arquitetura de microserviços tem se tornado cada vez mais popular, e Go é uma escolha sólida para o desenvolvimento de microserviços devido à sua eficiência em termos de desempenho e concorrência, bem como sua facilidade de implementação e manutenção.

- A comunidade em torno de Go é robusta e ativa, o que significa que há uma abundância de recursos, documentação e ferramentas disponíveis para os desenvolvedores.

> Em resumo, Go é uma escolha poderosa para o desenvolvimento de sistemas de backend, especialmente para aplicações que exigem alta performance, concorrência eficiente e facilidade de manutenção. Sua sintaxe simples, biblioteca padrão abrangente e suporte nativo a concorrência fazem dela uma opção atraente para uma ampla gama de cenários de desenvolvimento de software.

### 4. PostgreSQL

Um banco de dados é um sistema organizado para armazenar e gerenciar conjuntos de dados de forma estruturada, eficiente e segura. Ele é projetado para permitir a inserção, atualização, recuperação e manipulação de informações de maneira confiável e escalável. Os bancos de dados desempenham um papel fundamental em muitos aplicativos e sistemas de software, fornecendo um meio centralizado para armazenar e acessar dados.

Então escolhemos Postgres, por alguns motivos...

- Rico em Recursos: O PostgreSQL é conhecido por sua riqueza de recursos e capacidades avançadas, incluindo suporte a transações ACID (Atomicidade, Consistência, Isolamento, Durabilidade), chaves estrangeiras, gatilhos (triggers), procedimentos armazenados e replicação.

- Flexibilidade: O PostgreSQL suporta uma ampla variedade de tipos de dados, incluindo tipos nativos, tipos geométricos, tipos de dados personalizados e muito mais. Isso o torna flexível o suficiente para lidar com uma ampla gama de requisitos de aplicativos.

- Escalabilidade: O PostgreSQL é altamente escalável e pode lidar com grandes volumes de dados e cargas de trabalho intensivas com eficiência. Ele suporta técnicas de particionamento de tabelas, replicação síncrona e assíncrona e clustering para escalabilidade horizontal.

- Desempenho: O PostgreSQL é otimizado para desempenho e oferece recursos avançados de otimização de consultas, como índices, estatísticas, planejamento de consultas e execução paralela de consultas.

> Em resumo, o PostgreSQL é uma escolha poderosa e versátil para sistemas de banco de dados relacional, oferecendo uma combinação única de recursos avançados, escalabilidade, desempenho e flexibilidade, tudo isso em um pacote open-source.

Essas são algumas de nossas tabelas, para facilitar a compreensão sobre nossa aplicação:

##### Tabela de Usuários

| id | nickname | email    | password |
|----|----------|----------|----------|
|  71  |    gurodrigues-dev      |     teste@gmail.com     |     123teste     |

##### Tabela de Questões

| id | title    | description | date |
|----|----------|-------------|------|
|  72  |     Soma de Salários     |       José e sua esposa precisavam validar quanto daria a soma de seus salarios      |   01-03-2004-00:00:00   |

##### Tabela de Testes

| id | questionid  | param1 | param2 | response |
|----|------------------|--------|--------|----------|
|  73  |       72           |   2     |   2     |    4      |

##### Tabela de Respostas de Usuários

| id | userid | questionid | status | created_at |
|----|--------|------------|--------|------------|
|  74  |    71    |    72        |    running    |     01-03-2015-00:00:00       |


### 5. JDoodle

O JDoodle é uma plataforma online que oferece serviços de compilação e execução de código em várias linguagens de programação diretamente no navegador. Sua API de compilação é uma interface que permite aos desenvolvedores integrar a funcionalidade de compilação e execução de código em seus próprios aplicativos e sistemas.

A API de compilação do JDoodle oferece uma maneira simples e eficiente de executar código em tempo real em várias linguagens de programação, sem a necessidade de configurar ambientes de desenvolvimento localmente. Isso é especialmente útil para aprendizado de programação, testes rápidos de algoritmos e prototipagem de código.

A API de compilação do JDoodle geralmente funciona da seguinte forma:

- Envio do Código: O usuário envia o código fonte que deseja compilar e executar através da API do JDoodle. O código pode ser escrito em uma variedade de linguagens de programação suportadas, como C, C++, Java, Python, Ruby, entre outras.

- Seleção da Linguagem: O usuário especifica a linguagem de programação em que o código está escrito.

- Envio de Parâmetros Adicionais (Opcional): O usuário pode fornecer parâmetros adicionais, como entradas de teste, opções de compilação ou configurações específicas para a execução do código.

- Compilação e Execução: A API do JDoodle recebe o código fonte e o submete aos servidores do JDoodle, onde é compilado (se necessário) e executado. O resultado da compilação ou execução é retornado ao usuário, geralmente em formato JSON, que pode incluir mensagens de erro, saída do programa, status de execução, entre outros.

- Tratamento de Erros: A API do JDoodle também fornece informações detalhadas sobre erros de compilação ou execução, o que permite aos usuários depurar e corrigir problemas em seus códigos.

> A API de compilação do JDoodle é útil em uma variedade de cenários, incluindo aprendizado de programação, testes de algoritmos, entrevistas de codificação, e prototipagem rápida de código. Ela simplifica o processo de compilação e execução de código, eliminando a necessidade de configurações complexas de ambiente de desenvolvimento e oferecendo uma solução fácil de usar e acessível diretamente no navegador.

### 6. WebScraper

Um web scraper é uma ferramenta ou programa que automatiza a extração de dados de páginas da web. Ele navega pelas páginas, analisa o HTML, extrai os dados relevantes e os armazena em um formato estruturado para posterior processamento e análise.

O Robot Framework oferece uma biblioteca chamada "SeleniumLibrary" que permite automatizar a interação com páginas da web, preenchimento de formulários, navegação e extração de dados. Com a combinação do SeleniumLibrary e outras bibliotecas, como a "RequestsLibrary", o Robot Framework pode ser configurado para realizar tarefas de web scraping e crawling.


- O Robot Framework é um framework de automação de testes e automação de processos de software que suporta testes de aceitação, automação de interfaces de usuário (UI) e automação de tarefas repetitivas. Embora o Robot Framework seja amplamente conhecido por suas capacidades de teste de software, ele também pode ser usado para automação de web scraping e crawling (rastreamento).

Como o Robot Framework Pode Ser Usado como Web Scraper?
O Robot Framework oferece uma biblioteca chamada "SeleniumLibrary" que permite automatizar a interação com páginas da web, preenchimento de formulários, navegação e extração de dados. Com a combinação do SeleniumLibrary e outras bibliotecas, como a "RequestsLibrary", o Robot Framework pode ser configurado para realizar tarefas de web scraping e crawling. Aqui está como você pode fazer isso:

> Instale as Bibliotecas Necessárias:

Para usar o SeleniumLibrary, você precisa ter o Selenium WebDriver instalado e configurado em seu sistema.
Instale a SeleniumLibrary e a RequestsLibrary no ambiente do Robot Framework.
Escreva Test Cases: Escreva test cases no formato do Robot Framework para descrever as ações que você deseja automatizar. Isso pode incluir abrir páginas da web, preencher formulários, clicar em links, e extrair dados.

Use Palavras-Chave do SeleniumLibrary: O SeleniumLibrary fornece uma variedade de palavras-chave para interagir com elementos da página da web, como clicar em botões, preencher campos de formulário, extrair texto, e assim por diante.

Implemente Lógica de Crawling e Scraping: Use as capacidades do Robot Framework para implementar a lógica de crawling e scraping. Isso pode incluir a navegação por várias páginas, a extração de dados de tabelas ou listas, e o armazenamento dos dados extraídos em um formato estruturado, como CSV ou JSON.

Execute os Testes: Execute os testes do Robot Framework para iniciar o processo de web scraping e crawling. O framework automatizará as ações definidas nos test cases, navegando pelas páginas da web, interagindo com os elementos e extrai os dados conforme necessário.

Analise os Dados: Após a conclusão da execução dos testes, você pode analisar os dados extraídos e usá-los para os fins desejados, como análise de dados, geração de relatórios, ou alimentação de sistemas de backend.

### 7. Discord-Bot

O Discord é uma plataforma de comunicação por voz e texto projetada para comunidades de jogos, embora tenha se expandido para atender a uma variedade de comunidades online. Ele oferece salas de bate-papo (chamadas de servidores) onde os usuários podem se conectar, conversar por texto, voz e vídeo, compartilhar conteúdo e colaborar em projetos.

Os bots da comunidade no Discord são programas automatizados que oferecem uma variedade de funcionalidades para os usuários dentro dos servidores do Discord. Eles são construídos usando a API do Discord e podem ser programados para realizar uma ampla gama de tarefas, desde moderação do servidor até jogos e entretenimento.

### Desenho da Arquitetura

![Texto alternativo](https://i.imgur.com/uUbNNVq.png)

## Como Usar

Para configurar e implantar um servidor web:

1. Escolha um servidor web adequado para suas necessidades.
2. Configure o servidor web de acordo com os requisitos do seu projeto.
3. Carregue os recursos do seu site no sistema de arquivos do servidor.
4. Teste o servidor web para garantir que esteja respondendo corretamente às solicitações dos clientes.

## Contribuição

Se você encontrar problemas ou quiser contribuir para este projeto, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).



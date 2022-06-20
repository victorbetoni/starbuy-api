
# Starbuy API

Todas as rotas de nossa API descritas

OBS: Rotas que possuem a anotação "Requer autenticação" devem possuir o header de autorização com o JWT do usuário
```
  'Authorization': 'Bearer ${token}'
```

## Produtos

### Overview:
```http
  POST /item
  GET  /items
  GET  /item/${id}
  GET  /item
```

### Adicionar item
#### Autenticação requerida

```http
  POST /item
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :---------------|
| `item.title`  | `string`   | **Obrigatório** |
| `item.seller`      | `string`   | **Obrigatório** |
| `item.price`      | `number`   | **Obrigatório** |
| `item.stock`      | `number`   | **Obrigatório** |
| `item.category`      | `number`   | **Obrigatório** |
| `assets`  | `string[]`   | **Obrigatório** |

Request:
```json
{
  "item": {
    "title": "Nome do produto",
    "seller": "Username do vendedor",
    "price": 19.99,
    "stock": 10,
    "category": 1,
    "description": "description"
  },
  "assets": [
    "url de alguma imagem"
  ]
}
```

Response:
```json
{
		"item": {
			"identifier": "ea848cc08a1348338ceb4b5d593bc26d",
			"title": "Nome do produto",
			"seller": "vasco2004",
			"price": 19.99,
			"stock": 10,
			"category": 1,
			"description": "description"
		},
		"assets": [
			"url de alguma imagem"
		]
}
```

### Pegar todos os produtos
```http
  GET /items
```

Response:
```json
[
    {
        "item": {
            "identifier": "ddbe01ad9f7446dfa0dcb35249080514",
            "title": "Fender American Pro Jazzmaster RW Candy Apple Red LH",
            "seller": {
                "username": "victorbetoni",
                "email": "victorbetoni@protonmail.com",
                "name": "Victor Betoni",
                "birthdate": "2005-09-09T00:00:00Z",
                "seller": false,
                "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                "city": "LIMEIRA-SP",
                "registration": "2019-01-01T00:00:00Z"
            },
            "price": 2300.99,
            "stock": 10,
            "category": 7,
            "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
        },
        "assets": [
            "https://muzikercdn.com/uploads/products/4382/438279/thumb_d_gallery_base_fa379a17.jpg"
        ]
    }
]
```

### Retornar um item específico
 
```http
  GET /item/${id}
```
Response:
```json
{
    "item": {
        "item": {
            "identifier": "ddbe01ad9f7446dfa0dcb35249080514",
            "title": "Fender American Pro Jazzmaster RW Candy Apple Red LH",
            "seller": {
                "username": "victorbetoni",
                "email": "victorbetoni@protonmail.com",
                "name": "Victor Betoni",
                "birthdate": "2005-09-09T00:00:00Z",
                "seller": false,
                "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                "city": "LIMEIRA-SP",
                "registration": "2019-01-01T00:00:00Z"
            },
            "price": 2300.99,
            "stock": 10,
            "category": 7,
            "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
        },
        "assets": [
            "https://muzikercdn.com/uploads/products/4382/438279/thumb_d_gallery_base_fa379a17.jpg"
        ]
    },
    "average": 5
}
```

### Buscar produtos a partir de uma string
```http
  GET /item/search/${query}
```
Response:
```json
[
    {
        "item": {
            "identifier": "ddbe01ad9f7446dfa0dcb35249080514",
            "title": "Fender American Pro Jazzmaster RW Candy Apple Red LH",
            "seller": {
                "username": "victorbetoni",
                "email": "victorbetoni@protonmail.com",
                "name": "Victor Betoni",
                "birthdate": "2005-09-09T00:00:00Z",
                "seller": false,
                "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                "city": "LIMEIRA-SP",
                "registration": "2019-01-01T00:00:00Z"
            },
            "price": 2300.99,
            "stock": 10,
            "category": 7,
            "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
        },
        "assets": [
            "https://muzikercdn.com/uploads/products/4382/438279/thumb_d_gallery_base_fa379a17.jpg"
        ]
    }
]
```

### Buscar produtos a partir de uma categoria
```http
  GET /item/category/${id}
```
Response:
```json
[
    {
        "item": {
            "identifier": "9fcd608e6bfa48049f04da7071466112",
            "title": "Chuveiro Acqua Duo 127V 5500W, Lorenzetti, 7510100, Branco, Pequeno",
            "seller": {
                "username": "juhtc",
                "email": "juhtc@gmail.com",
                "name": "Julia Teles Cruz",
                "birthdate": "2004-12-12T00:00:00Z",
                "seller": true,
                "profile_picture": "https://midias.correiobraziliense.com.br/_midias/jpg/2022/02/03/675x450/1_juliana_bonde-7409319.jpg?20220203195909?20220203195909",
                "city": "CAMPINAS-SP",
                "registration": "2022-06-17T00:00:00Z"
            },
            "price": 328.58,
            "stock": 842,
            "category": 3,
            "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque diam ex, feugiat in eros nec, volutpat consequat sapien. Praesent at risus interdum, ultricies nulla id, mattis velit. Aliquam sollicitudin pharetra est. Vivamus lacinia metus ultricies, volutpat dui sed, vehicula velit. Curabitur ut gravida nisl, sed porta magna. Mauris nec ligula nec magna suscipit fringilla. Donec ac condimentum purus. Proin ac ullamcorper eros. In bibendum erat nec tellus malesuada cursus. Sed faucibus nisl non urna velit."
        },
        "assets": [
            "https://images-na.ssl-images-amazon.com/images/I/51UqEdnYTAL.__AC_SX300_SY300_QL70_ML2_.jpg"
        ]
    },
    {
        "item": {
            "identifier": "9df6bf5eef3811ec8ea00242ac120002",
            "title": "Umidificador de Ar Digital, Branco, 2L, 18 Watts, Bivolt, Elgin",
            "seller": {
                "username": "juhtc",
                "email": "juhtc@gmail.com",
                "name": "Julia Teles Cruz",
                "birthdate": "2004-12-12T00:00:00Z",
                "seller": true,
                "profile_picture": "https://midias.correiobraziliense.com.br/_midias/jpg/2022/02/03/675x450/1_juliana_bonde-7409319.jpg?20220203195909?20220203195909",
                "city": "CAMPINAS-SP",
                "registration": "2022-06-17T00:00:00Z"
            },
            "price": 98.89,
            "stock": 12,
            "category": 3,
            "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque diam ex, feugiat in eros nec, volutpat consequat sapien. Praesent at risus interdum, ultricies nulla id, mattis velit. Aliquam sollicitudin pharetra est. Vivamus lacinia metus ultricies, volutpat dui sed, vehicula velit. Curabitur ut gravida nisl, sed porta magna. Mauris nec ligula nec magna suscipit fringilla. Donec ac condimentum purus. Proin ac ullamcorper eros. In bibendum erat nec tellus malesuada cursus. Sed faucibus nisl non urna velit."
        },
        "assets": [
            "https://m.media-amazon.com/images/I/616Wn5HT++L._AC_SY879_.jpg"
        ]
    }
]
```
## Usuários

### Overview:
```http
  POST /register
  POST /login
  GET  /user/${username}
```

### LOGIN [POST] - Faz o login de um usuário e retorna um token JWT
```http
  POST /login
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :---------------|
| `username`  | `string`   | **Obrigatório** |
| `password`  | `string`   | **Obrigatório** |

Request:
```json
  {
    "username": "user",
    "password": "password"
  }
```

Response:
```json
{
  "jwt": "token jwt",
  "message": "Sessão iniciada com sucesso",
  "status": true,
  "user": {
    "username": "juhtc",
    "email": "juhtc@gmail.com",
    "name": "Julia Teles Cruz",
    "birthdate": "2004-12-12T00:00:00Z",
    "seller": true,
    "profile_picture": "https://midias.correiobraziliense.com.br/_midias/jpg/2022/02/03/675x450/1_juliana_bonde-7409319.jpg?20220203195909?20220203195909",
    "city": "CAMPINAS-SP",
    "registration": "2022-06-17T00:00:00Z"
  }
}
```

### REGISTER [POST] - Registra um usuário
```http
  POST /register
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :---------------|
| `username`  | `string`   | **Obrigatório** |
| `name`      | `string`   | **Obrigatório** |
| `email`      | `string`   | **Obrigatório** |
| `city`      | `string`   | **Obrigatório** |
| `birthdate`      | `string`   | **Obrigatório** |
| `seller`      | `boolean`   | **Obrigatório** |
| `password`  | `string`   | **Obrigatório** |

Request:
```json
{
  "username": "paulao",
  "name": "Paulo Eduardo Crystal",
  "email": "paulao@gmail.com",
  "city": "IRACEMAPOLIS-SP",
  "birthdate": "2004-10-10",
  "seller": false,
  "profile_picture": "url qualquer",
  "password": "admin"
}
```

Response:
```json
{
  "jwt": "token jwt",
  "message": "Registrado com sucesso",
  "status": true,
  "user": {
    "username": "paulao",
    "email": "paulao@gmail.com",
    "name": "Paulo Eduardo Crystal",
    "birthdate": "2004-10-10",
    "seller": false,
    "profile_picture": "url qualquer",
    "city": "IRACEMAPOLIS-SP",
    "registration": "2022-06-19"
  }
}
```

### Retornar um usuário específico

```http
  GET /user/${username}
```
Response:
```json
{
  "user": {
    "username": "vasco2004",
    "email": "vasco@gmail.com",
    "name": "Fernando Antonio",
    "birthdate": "2004-10-10T00:00:00Z",
    "seller": true,
    "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
    "city": "LIMEIRA-SP",
    "registration": "2021-08-09T00:00:00Z"
  },
  "rating": 0
}
```

## Carrinho

### Overview:
```http
  POST    /cart
  DELETE  /cart/${produto}
  GET     /cart
```

### Adicionar um item ao carrinho
#### Requer atenticação
```http
  POST /cart
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :---------------|
| `item`      | `string`   | **Obrigatório** |
| `quantity`  | `number`   | **Obrigatório** |

Request:
```json
  {
    "item": "id do produto",
    "quantity": 3
  }
```

Response:
```json
  {
    "status": "true se tudo tiver dado certo e false se tiver dado erro",
    "message": "mensagem de sucesso ou de erro"
  }
```

### Pegar o carrinho de um usuario
#### Requer autenticação
```http
  GET /cart
```

Response:
```json
  [
    {
        "quantity": 1,
        "item": {
            "item": {
                "identifier": "58b9c68cb8e441c2915d2be54e28d757",
                "title": "Notebook ASUS ROG Zephyrus Duo 15 GX550LXS-HF157T Cinza",
                "seller": {
                    "username": "victorbetoni",
                    "email": "victorbetoni@protonmail.com",
                    "name": "Victor Betoni",
                    "birthdate": "2005-09-09T00:00:00Z",
                    "seller": false,
                    "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                    "city": "LIMEIRA-SP",
                    "registration": "2019-01-01T00:00:00Z"
                },
                "price": 40449.99,
                "stock": 12,
                "category": 1,
                "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
            },
            "assets": [
                "https://br.store.asus.com/media/catalog/product/cache/80ca4d3b08020b5abcd9abd5d305c273/a/s/asus_selos_zephyrusduo15_v1.png"
            ]
        }
    },
    {
        "quantity": 1,
        "item": {
            "item": {
                "identifier": "9df6bf5eef3811ec8ea00242ac120002",
                "title": "Umidificador de Ar Digital, Branco, 2L, 18 Watts, Bivolt, Elgin",
                "seller": {
                    "username": "juhtc",
                    "email": "juhtc@gmail.com",
                    "name": "Julia Teles Cruz",
                    "birthdate": "2004-12-12T00:00:00Z",
                    "seller": true,
                    "profile_picture": "https://midias.correiobraziliense.com.br/_midias/jpg/2022/02/03/675x450/1_juliana_bonde-7409319.jpg?20220203195909?20220203195909",
                    "city": "CAMPINAS-SP",
                    "registration": "2022-06-17T00:00:00Z"
                },
                "price": 98.89,
                "stock": 12,
                "category": 3,
                "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque diam ex, feugiat in eros nec, volutpat consequat sapien. Praesent at risus interdum, ultricies nulla id, mattis velit. Aliquam sollicitudin pharetra est. Vivamus lacinia metus ultricies, volutpat dui sed, vehicula velit. Curabitur ut gravida nisl, sed porta magna. Mauris nec ligula nec magna suscipit fringilla. Donec ac condimentum purus. Proin ac ullamcorper eros. In bibendum erat nec tellus malesuada cursus. Sed faucibus nisl non urna velit."
            },
            "assets": [
                "https://m.media-amazon.com/images/I/616Wn5HT++L._AC_SY879_.jpg"
            ]
        }
    }
]
```

### Adicionar um item ao carrinho
#### Requer atenticação
```http
  DELETE /cart/${item}
```

Response:
```json
  {
    "status": "true se tudo tiver dado certo e false se tiver dado erro",
    "message": "mensagem de sucesso ou de erro"
  }
```


## Compras

### Overview:
```http
  POST    /order
  GET     /order/${id}
  GET     /orders
```

### Adicionar uma compra
#### Requer atenticação
```http
  POST /cart
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :---------------|
| `item`      | `string`   | **Obrigatório** |
| `quantity`  | `number`   | **Obrigatório** |

Request:
```json
  {
    "item": "id do produto",
    "quantity": 3
  }
```

Response:
```json
  {
    "status": "true se tudo tiver dado certo e false se tiver dado erro",
    "message": "mensagem de sucesso ou de erro"
  }
```

### Pegar todas as compras de um usuário
#### Requer atenticação
```http
  GET /order
```

Response:
```json
  [
    {
        "identifier": "f4a39854a125420bb8f953646c05fc3e",
        "customer": {
            "username": "vasco2004",
            "email": "vasco@gmail.com",
            "name": "Fernando Antonio",
            "birthdate": "2004-10-10T00:00:00Z",
            "seller": true,
            "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
            "city": "LIMEIRA-SP",
            "registration": "2021-08-09T00:00:00Z"
        },
        "seller": {
            "username": "victorbetoni",
            "email": "victorbetoni@protonmail.com",
            "name": "Victor Betoni",
            "birthdate": "2005-09-09T00:00:00Z",
            "seller": false,
            "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
            "city": "LIMEIRA-SP",
            "registration": "2019-01-01T00:00:00Z"
        },
        "item": {
            "item": {
                "identifier": "ddbe01ad9f7446dfa0dcb35249080514",
                "title": "Fender American Pro Jazzmaster RW Candy Apple Red LH",
                "seller": {
                    "username": "victorbetoni",
                    "email": "victorbetoni@protonmail.com",
                    "name": "Victor Betoni",
                    "birthdate": "2005-09-09T00:00:00Z",
                    "seller": false,
                    "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                    "city": "LIMEIRA-SP",
                    "registration": "2019-01-01T00:00:00Z"
                },
                "price": 2300.99,
                "stock": 10,
                "category": 7,
                "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
            },
            "assets": [
                "https://muzikercdn.com/uploads/products/4382/438279/thumb_d_gallery_base_fa379a17.jpg"
            ]
        },
        "price": 9203.96,
        "quantity": 4
    },
    {
        "identifier": "9d3ec8a33f4449c18356c2cf2f4a6d5f",
        "customer": {
            "username": "vasco2004",
            "email": "vasco@gmail.com",
            "name": "Fernando Antonio",
            "birthdate": "2004-10-10T00:00:00Z",
            "seller": true,
            "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
            "city": "LIMEIRA-SP",
            "registration": "2021-08-09T00:00:00Z"
        },
        "seller": {
            "username": "victorbetoni",
            "email": "victorbetoni@protonmail.com",
            "name": "Victor Betoni",
            "birthdate": "2005-09-09T00:00:00Z",
            "seller": false,
            "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
            "city": "LIMEIRA-SP",
            "registration": "2019-01-01T00:00:00Z"
        },
        "item": {
            "item": {
                "identifier": "5e946e32a740436db62ea34c435cd421",
                "title": "MOCHILA NOTEBOOK ATÉ 15.6 COM ENTRADA USB - GOLDEN WOLF GB00370 PRETO",
                "seller": {
                    "username": "victorbetoni",
                    "email": "victorbetoni@protonmail.com",
                    "name": "Victor Betoni",
                    "birthdate": "2005-09-09T00:00:00Z",
                    "seller": false,
                    "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                    "city": "LIMEIRA-SP",
                    "registration": "2019-01-01T00:00:00Z"
                },
                "price": 110.5,
                "stock": 30,
                "category": 2,
                "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
            },
            "assets": [
                "https://images.tcdn.com.br/img/img_prod/593198/mochila_notebook_ate_15_6_com_entrada_usb_golden_wolf_gb00370_preto_2765_1_50022ae490a7eb45ac4c6aae2f106799.jpg"
            ]
        },
        "price": 110.5,
        "quantity": 1
    }
]
```

### Pegar uma compra especifica
#### Requer atenticação
```http
  GET /order/${id}
```

Response:
```json
{
    "identifier": "f4a39854a125420bb8f953646c05fc3e",
    "customer": {
        "username": "vasco2004",
        "email": "vasco@gmail.com",
        "name": "Fernando Antonio",
        "birthdate": "2004-10-10T00:00:00Z",
        "seller": true,
        "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
        "city": "LIMEIRA-SP",
        "registration": "2021-08-09T00:00:00Z"
    },
    "seller": {
        "username": "victorbetoni",
        "email": "victorbetoni@protonmail.com",
        "name": "Victor Betoni",
        "birthdate": "2005-09-09T00:00:00Z",
        "seller": false,
        "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
        "city": "LIMEIRA-SP",
        "registration": "2019-01-01T00:00:00Z"
    },
    "item": {
        "item": {
            "identifier": "ddbe01ad9f7446dfa0dcb35249080514",
            "title": "Fender American Pro Jazzmaster RW Candy Apple Red LH",
            "seller": {
                "username": "victorbetoni",
                "email": "victorbetoni@protonmail.com",
                "name": "Victor Betoni",
                "birthdate": "2005-09-09T00:00:00Z",
                "seller": false,
                "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                "city": "LIMEIRA-SP",
                "registration": "2019-01-01T00:00:00Z"
            },
            "price": 2300.99,
            "stock": 10,
            "category": 7,
            "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
        },
        "assets": [
            "https://muzikercdn.com/uploads/products/4382/438279/thumb_d_gallery_base_fa379a17.jpg"
        ]
    },
    "price": 9203.96,
    "quantity": 4
}
```
## Avaliações

### Overview:
```http
  POST   /review
  GET    /review?user=${usuario}&product=${produto}
  GET    /user/reviews
  GET    /user/reviews/received/${user}
  GET    /item/reviews/${item}
```

### Adicionar uma nova avaliação
#### Requer autenticação

```http
  POST  /review
```

Request: 
```json
{
  "item": "id do produto",
  "message": "mensagem da avaliação",
  "rate": 10
}
```
Response:
```json
  {
    "status": "true se tudo tiver dado certo e false se tiver dado erro",
    "message": "mensagem de sucesso ou de erro"
  }
```

### Pegar uma avaliação específica
#### Requer autenticação

```http
  POST  /review?user=${usuario}&product=${produto}
```

Response:
```json
  {
    "user": {
        "username": "vasco2004",
        "email": "vasco@gmail.com",
        "name": "Fernando Antonio",
        "birthdate": "2004-10-10T00:00:00Z",
        "seller": true,
        "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
        "city": "LIMEIRA-SP",
        "registration": "2021-08-09T00:00:00Z"
    },
    "item": {
        "item": {
            "identifier": "ddbe01ad9f7446dfa0dcb35249080514",
            "title": "Fender American Pro Jazzmaster RW Candy Apple Red LH",
            "seller": {
                "username": "victorbetoni",
                "email": "victorbetoni@protonmail.com",
                "name": "Victor Betoni",
                "birthdate": "2005-09-09T00:00:00Z",
                "seller": false,
                "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                "city": "LIMEIRA-SP",
                "registration": "2019-01-01T00:00:00Z"
            },
            "price": 2300.99,
            "stock": 10,
            "category": 7,
            "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
        },
        "assets": [
            "https://muzikercdn.com/uploads/products/4382/438279/thumb_d_gallery_base_fa379a17.jpg"
        ]
    },
    "message": "Produto muito bom! Recomendo! Só a entrega atrasou um pouco...",
    "rate": 9
}
```

### Pegar todas as avaliações feitas por um usuário
#### Requer autenticação

```http
  GET  /user/reviews
```

Response
```json
{
    "reviews": [
        {
            "user": {
                "username": "vasco2004",
                "email": "vasco@gmail.com",
                "name": "Fernando Antonio",
                "birthdate": "2004-10-10T00:00:00Z",
                "seller": true,
                "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
                "city": "LIMEIRA-SP",
                "registration": "2021-08-09T00:00:00Z"
            },
            "item": {
                "item": {
                    "identifier": "2f5c8464ef3a11ec8ea00242ac120002",
                    "title": "O mínimo que você precisa saber para não ser um idiota Capa comum",
                    "seller": {
                        "username": "juhtc",
                        "email": "juhtc@gmail.com",
                        "name": "Julia Teles Cruz",
                        "birthdate": "2004-12-12T00:00:00Z",
                        "seller": true,
                        "profile_picture": "https://midias.correiobraziliense.com.br/_midias/jpg/2022/02/03/675x450/1_juliana_bonde-7409319.jpg?20220203195909?20220203195909",
                        "city": "CAMPINAS-SP",
                        "registration": "2022-06-17T00:00:00Z"
                    },
                    "price": 220,
                    "stock": 99,
                    "category": 4,
                    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque diam ex, feugiat in eros nec, volutpat consequat sapien. Praesent at risus interdum, ultricies nulla id, mattis velit. Aliquam sollicitudin pharetra est. Vivamus lacinia metus ultricies, volutpat dui sed, vehicula velit. Curabitur ut gravida nisl, sed porta magna. Mauris nec ligula nec magna suscipit fringilla. Donec ac condimentum purus. Proin ac ullamcorper eros. In bibendum erat nec tellus malesuada cursus. Sed faucibus nisl non urna velit."
                },
                "assets": [
                    "https://salvadornorteonline.com.br/salvadornorteonline/2021/12/O-Minimo-Que-Voce-Precisa-Saber-Para-Nao-Ser-Um-Idiota.png-2021-11-30-193053.png"
                ]
            },
            "message": "Ótimo livro do mestre olavo",
            "rate": 7
        },
        {
            "user": {
                "username": "vasco2004",
                "email": "vasco@gmail.com",
                "name": "Fernando Antonio",
                "birthdate": "2004-10-10T00:00:00Z",
                "seller": true,
                "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
                "city": "LIMEIRA-SP",
                "registration": "2021-08-09T00:00:00Z"
            },
            "item": {
                "item": {
                    "identifier": "5e946e32a740436db62ea34c435cd421",
                    "title": "MOCHILA NOTEBOOK ATÉ 15.6 COM ENTRADA USB - GOLDEN WOLF GB00370 PRETO",
                    "seller": {
                        "username": "victorbetoni",
                        "email": "victorbetoni@protonmail.com",
                        "name": "Victor Betoni",
                        "birthdate": "2005-09-09T00:00:00Z",
                        "seller": false,
                        "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                        "city": "LIMEIRA-SP",
                        "registration": "2019-01-01T00:00:00Z"
                    },
                    "price": 110.5,
                    "stock": 30,
                    "category": 2,
                    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
                },
                "assets": [
                    "https://images.tcdn.com.br/img/img_prod/593198/mochila_notebook_ate_15_6_com_entrada_usb_golden_wolf_gb00370_preto_2765_1_50022ae490a7eb45ac4c6aae2f106799.jpg"
                ]
            },
            "message": "Otima mochila, atendeu minhas expectativas",
            "rate": 10
        }

}
```

### Pegar todas as avaliações recebidas por um usuário

```http
  GET    /user/reviews/received/${user}
```

Response:
```json
{
    "reviews": [
        {
            "user": {
                "username": "juhtc",
                "email": "juhtc@gmail.com",
                "name": "Julia Teles Cruz",
                "birthdate": "2004-12-12T00:00:00Z",
                "seller": true,
                "profile_picture": "https://midias.correiobraziliense.com.br/_midias/jpg/2022/02/03/675x450/1_juliana_bonde-7409319.jpg?20220203195909?20220203195909",
                "city": "CAMPINAS-SP",
                "registration": "2022-06-17T00:00:00Z"
            },
            "item": {
                "item": {
                    "identifier": "ddbe01ad9f7446dfa0dcb35249080514",
                    "title": "Fender American Pro Jazzmaster RW Candy Apple Red LH",
                    "seller": {
                        "username": "victorbetoni",
                        "email": "victorbetoni@protonmail.com",
                        "name": "Victor Betoni",
                        "birthdate": "2005-09-09T00:00:00Z",
                        "seller": false,
                        "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                        "city": "LIMEIRA-SP",
                        "registration": "2019-01-01T00:00:00Z"
                    },
                    "price": 2300.99,
                    "stock": 10,
                    "category": 7,
                    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
                },
                "assets": [
                    "https://muzikercdn.com/uploads/products/4382/438279/thumb_d_gallery_base_fa379a17.jpg"
                ]
            },
            "message": "Pessimo produto, fui tocar um Sepultura e toquei Baroes da Pisadinha sem querer",
            "rate": 1
        },
        {
            "user": {
                "username": "vasco2004",
                "email": "vasco@gmail.com",
                "name": "Fernando Antonio",
                "birthdate": "2004-10-10T00:00:00Z",
                "seller": true,
                "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
                "city": "LIMEIRA-SP",
                "registration": "2021-08-09T00:00:00Z"
            },
            "item": {
                "item": {
                    "identifier": "ddbe01ad9f7446dfa0dcb35249080514",
                    "title": "Fender American Pro Jazzmaster RW Candy Apple Red LH",
                    "seller": {
                        "username": "victorbetoni",
                        "email": "victorbetoni@protonmail.com",
                        "name": "Victor Betoni",
                        "birthdate": "2005-09-09T00:00:00Z",
                        "seller": false,
                        "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                        "city": "LIMEIRA-SP",
                        "registration": "2019-01-01T00:00:00Z"
                    },
                    "price": 2300.99,
                    "stock": 10,
                    "category": 7,
                    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
                },
                "assets": [
                    "https://muzikercdn.com/uploads/products/4382/438279/thumb_d_gallery_base_fa379a17.jpg"
                ]
            },
            "message": "Produto muito bom! Recomendo! Só a entrega atrasou um pouco...",
            "rate": 9
        },
        {
            "user": {
                "username": "vasco2004",
                "email": "vasco@gmail.com",
                "name": "Fernando Antonio",
                "birthdate": "2004-10-10T00:00:00Z",
                "seller": true,
                "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
                "city": "LIMEIRA-SP",
                "registration": "2021-08-09T00:00:00Z"
            },
            "item": {
                "item": {
                    "identifier": "5e946e32a740436db62ea34c435cd421",
                    "title": "MOCHILA NOTEBOOK ATÉ 15.6 COM ENTRADA USB - GOLDEN WOLF GB00370 PRETO",
                    "seller": {
                        "username": "victorbetoni",
                        "email": "victorbetoni@protonmail.com",
                        "name": "Victor Betoni",
                        "birthdate": "2005-09-09T00:00:00Z",
                        "seller": false,
                        "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                        "city": "LIMEIRA-SP",
                        "registration": "2019-01-01T00:00:00Z"
                    },
                    "price": 110.5,
                    "stock": 30,
                    "category": 2,
                    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
                },
                "assets": [
                    "https://images.tcdn.com.br/img/img_prod/593198/mochila_notebook_ate_15_6_com_entrada_usb_golden_wolf_gb00370_preto_2765_1_50022ae490a7eb45ac4c6aae2f106799.jpg"
                ]
            },
            "message": "Otima mochila, atendeu minhas expectativas",
            "rate": 10
        }
    ],
    "average": 6.666666666666667
}
```

### Pegar todas as avaliações de um produto

```http
  GET    /item/reviews/${item}
```

Request:
```json
{
    "reviews": [
        {
            "user": {
                "username": "juhtc",
                "email": "juhtc@gmail.com",
                "name": "Julia Teles Cruz",
                "birthdate": "2004-12-12T00:00:00Z",
                "seller": true,
                "profile_picture": "https://midias.correiobraziliense.com.br/_midias/jpg/2022/02/03/675x450/1_juliana_bonde-7409319.jpg?20220203195909?20220203195909",
                "city": "CAMPINAS-SP",
                "registration": "2022-06-17T00:00:00Z"
            },
            "item": {
                "item": {
                    "identifier": "ddbe01ad9f7446dfa0dcb35249080514",
                    "title": "Fender American Pro Jazzmaster RW Candy Apple Red LH",
                    "seller": {
                        "username": "victorbetoni",
                        "email": "victorbetoni@protonmail.com",
                        "name": "Victor Betoni",
                        "birthdate": "2005-09-09T00:00:00Z",
                        "seller": false,
                        "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                        "city": "LIMEIRA-SP",
                        "registration": "2019-01-01T00:00:00Z"
                    },
                    "price": 2300.99,
                    "stock": 10,
                    "category": 7,
                    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
                },
                "assets": [
                    "https://muzikercdn.com/uploads/products/4382/438279/thumb_d_gallery_base_fa379a17.jpg"
                ]
            },
            "message": "Pessimo produto, fui tocar um Sepultura e toquei Baroes da Pisadinha sem querer",
            "rate": 1
        },
        {
            "user": {
                "username": "vasco2004",
                "email": "vasco@gmail.com",
                "name": "Fernando Antonio",
                "birthdate": "2004-10-10T00:00:00Z",
                "seller": true,
                "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
                "city": "LIMEIRA-SP",
                "registration": "2021-08-09T00:00:00Z"
            },
            "item": {
                "item": {
                    "identifier": "ddbe01ad9f7446dfa0dcb35249080514",
                    "title": "Fender American Pro Jazzmaster RW Candy Apple Red LH",
                    "seller": {
                        "username": "victorbetoni",
                        "email": "victorbetoni@protonmail.com",
                        "name": "Victor Betoni",
                        "birthdate": "2005-09-09T00:00:00Z",
                        "seller": false,
                        "profile_picture": "https://avatars.githubusercontent.com/u/40972803?v=4",
                        "city": "LIMEIRA-SP",
                        "registration": "2019-01-01T00:00:00Z"
                    },
                    "price": 2300.99,
                    "stock": 10,
                    "category": 7,
                    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id imperdiet augue, a scelerisque orci. Etiam risus velit, tempor vitae lorem eget, lobortis semper felis. Integer sed pretium enim. Integer eu metus et nisi mattis finibus. Proin leo est, venenatis viverra aliquam in, pretium nec elit. Donec viverra fringilla sapien, vitae fermentum nisl iaculis et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Etiam aliquam quis arcu ut lacinia. Pellentesque ac lectus."
                },
                "assets": [
                    "https://muzikercdn.com/uploads/products/4382/438279/thumb_d_gallery_base_fa379a17.jpg"
                ]
            },
            "message": "Produto muito bom! Recomendo! Só a entrega atrasou um pouco...",
            "rate": 9
        }
    ],
    "average": 5
}
```
## Endereços

### Overview:
```http
  POST   /user/address
  GET    /user/address/${id}
  GET    /user/address
```

### Adicionar um novo Endereços
#### Requer autenticação

```http
  POST /user/address
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :---------------|
| `cep`      | `string`   | **Obrigatório** |
| `number`  | `string`   | **Obrigatório** |
| `complement`  | `string`   | |

Request
```json
{
  "cep": "12352-124",
  "number": "192",
  "complement": "Portão amarelo"
}
```

Response:
```json
  {
    "status": "true se tudo tiver dado certo e false se tiver dado erro",
    "message": "mensagem de sucesso ou de erro"
  }
```

### Pegar um endereço especifico
#### Requer autenticação

```http
  GET /user/address/${id}
```
Response
```json
{
    "identifier": "fc900b70dad141208dd975af708eddf5",
    "user": {
        "username": "vasco2004",
        "email": "vasco@gmail.com",
        "name": "Fernando Antonio",
        "birthdate": "2004-10-10T00:00:00Z",
        "seller": true,
        "profile_picture": "https://a.espncdn.com/i/teamlogos/soccer/500/3454.png",
        "city": "LIMEIRA-SP",
        "registration": "2021-08-09T00:00:00Z"
    },
    "cep": "12345789",
    "number": 12,
    "complement": "Muro amarelo"
}
```

### Pegar todos os endereços de um usuário
#### Requer autenticação

```http
  GET /user/address
```
Response
```json
[
    {
        "identifier": "fc900b70dad141208dd975af708eddf5",
        "user": "vasco2004",
        "cep": "12345789",
        "number": 12,
        "complement": "Muro amarelo"
    },
    {
        "identifier": "1d058321e9bd42f0998cda1c34a4de04",
        "user": "vasco2004",
        "cep": "98765432",
        "number": 145,
        "complement": "Atrás do hospital"
    }
]
```

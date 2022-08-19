# [Contessa](https://contessa-cocktail.netlify.app)
Contessaでは、その日その時の気分に応じて、条件を入力し様々なカクテルの検索を行うことができます。  
*URL*: [https://contessa-cocktail.netlify.app](https://contessa-cocktail.netlify.app)

#### 使用技術
・フロントエンド：Nuxt.js(2.15.8), CSS, Vuetify  
・バックエンド(API)：Go(1.18.1), Echo, Gorm, Air  
・インフラ：Docker(20.10.13), docker-compose(1.29.2), Netlify, MySQL(8.0.28), AWS(ALB, EC2, Route53, ACM, VPC, IAM)  

#### 機能一覧
・JWT認証(新規登録, ログイン, ログアウト)  
・REST API(カクテル検索/カクテル情報表示)  
・カクテルのレビュー追加/表示  
・カクテルのレーティング追加/表示  
・カクテルのお気に入り登録/解除/表示  

#### インフラ構成図
![スクリーンショット 2022-08-19 19 09 59](https://user-images.githubusercontent.com/95117254/185603149-8c07131d-e936-42bb-b215-48093a0538df.png)  

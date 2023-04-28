演習１ 商品一覧の出力
SELECT * FROM products;

演習２ 商品名と価格を出力
SELECT
    name, 
    price 
FROM
    products
;

演習３ 男性ユーザを１０件出力
SELECT 
    id,
    last_name,
    gender
FROM
    users
WHERE 
    gender = 1
LIMIT
    10
;

57 ２０１７年１月の合計売上金額を出力
SELECT
    sum(amount)
FROM
    orders 
WHERE
    order_time between '2017-01-01 00:00:00'
    and '2017-02-01 00:00:00'
;

58 全商品の平均価格を出力
SELECT avg(price)  FROM products;

59 商品価格の最小値を出力
SELECT min(price) FROM products;

60 商品価格の最大値を出力
SELECT max(price)  FROM products ;

62 女性ユーザ数を出力
SELECT count(*) FROM users WHERE gender = 2;

63 ２０１７年１月アクセスしたユニークユーザ数を出力
SELECT count(distinct user_id)  FROM access_logs WHERE request_month = '2017-01-01';

64 都道府県別のユーザ数を出力
SELECT
    prefecture_id,
    count(*)
FROM
    users
GROUP BY
    prefecture_id
;

65 2017年の月別ユニークユーザ数を出力
SELECT
	request_month,
    count(distinct user_id)
 FROM
	access_logs 
 WHERE 
	request_month between '2017-01-01' 
    and '2018-01-01' 
GROUP BY
	request_month
;

66 2017年の月間ユニークユーザ数が630人以上の月を出力
SELECT
	request_month,
    count(distinct user_id)
 FROM
	access_logs 
 WHERE 
	request_month between '2017-01-01' 
    and '2018-01-01' 
GROUP BY
	request_month
HAVING
	count(distinct user_id) >= 630;
;

演習４
①アクセスログ一覧を出力
SELECT * FROM access_logs;

②２０１７年１月〜２０１７年６月のアクセスログを出力
SELECT
    * 
FROM
    access_logs
WHERE
    request_month between '2017-01-01'
    and '2017-07-01'
;

③月ごとのリクエスト数を出力
SELECT
    request_month,
    count(*)
FROM
    access_logs
WHERE
    request_month between '2017-01-01'
    and '2017-07-01'
GROUP BY
    request_month
;

④アクセス数が１０００以上の月だけ出力
SELECT
    request_month,
    count(*)
FROM
    access_logs
WHERE
    request_month between '2017-01-01'
    and '2017-07-01'
GROUP BY
    request_month
HAVING
	count(*) >= 1000
;

70 商品を価格が低い順に出力
SELECT * FROM products ORDER BY price;

71 商品一覧を価格が高い順に出力する。価格が同じ時は登録順で並び替える。
SELECT
    *
FROM
    products
ORDER BY
	price desc,
    id
;

72 演習：ユーザを生年月日が古い順に出力。同じ場合は都道府県IDの昇順で出力。
SELECT
    *
FROM
    users
ORDER BY
	birthday,
    prefecture_id
;

78 商品価格の税込価格を出力。少数第１位で四捨五入。
SELECT
    id,
    name,
    round(price * 1.1, 0)
FROM
    products
;

80 演習：宛名「名字＋さん」とメールアドレスを女性だけ出力
SELECT
    concat(last_name, 'さん'),
    email
FROM
    users
WHERE
	gender = 2
;

81 ユーザID、名字、名前、都道府県名を出力
SELECT
	u.id,
    u.last_name,
    u.first_name,
    p.name
FROM
    users as u
INNER JOIN
	prefectures as p
    on u.prefecture_id = p.id
;

82 ユーザID、名字、名前、都道府県名を女性のみ出力
SELECT
	u.id,
    u.last_name,
    u.first_name,
    p.name
FROM
    users as u
INNER JOIN
	prefectures as p
    on u.prefecture_id = p.id
WHERE
	u.gender = 2;
;

83 演習：2017年1月に東京都に住むユーザの情報一覧を出力（注文ID、注文日時、注文金額合計、ユーザID、名字、名前）
SELECT
	o.id order_id,
    o.order_time order_time,
    o.amount amount,
	u.id user_id,
    u.last_name last_name,
    u.first_name first_name
FROM
    orders AS o
INNER JOIN
	users AS u
    on o.user_id = u.id
WHERE
	order_time BETWEEN '2017-01-01-00:00:00'
    AND '2017-02-01-00:00:00:'
	AND u.prefecture_id = 13
ORDER BY order_id
;

94 すべての商品の販売個数一覧を出力
SELECT
	p.id,
    p.name,
	SUM(od.product_qty)
FROM
    products AS p
LEFT JOIN
	order_details AS od
    ON p.id = od.product_id
GROUP BY
	p.id
;

95 注文詳細情報、商品情報、名字、名前を含む注文一覧を出力
SELECT
	o.id order_id,
    o.user_id,
    u.last_name,
    u.first_name,
    o.amount,
    o.order_time,
    p.name product_name,
    od.product_qty qty,
    p.price product_price
FROM
    orders o
INNER JOIN
	order_details od
    ON o.id = od.order_id
INNER JOIN
	products p
    ON od.product_id = p.id
INNER JOIN
	users u
    ON o.user_id = u.id
;

96 IDが３の商品カテゴリ名を出力
SELECT
	p.id product_id,
    p.name product_name,
    c.name category_name
FROM
	 products p
INNER JOIN
	products_categories pc
    ON p.id = pc.product_id
INNER JOIN
	categories c
    ON pc.category_id = c.id
WHERE
	p.id = 3
;

97 ユーザとアドミンユーザを足し合わせた一覧を出力（email,性,名,性別）
SELECT
	email,
    last_name,
    first_name,
    gender
FROM
	 users
UNION
SELECT
	email,
    last_name,
    first_name,
    gender
FROM
	 admin_users
;

-- 100 都道府県別のユーザ数を出力
create view prefecture_user_counts(name, count)
as
SELECT
	p.name,
    count(*) count
FROM
	users u
INNER JOIN
	prefectures p
    ON u.prefecture_id = p.id
GROUP BY
	u.prefecture_id
;

SELECT
	name,
    count
FROM
	prefecture_user_counts;

-- 104 2017年12月に商品を購入していないユーザを出力
SELECT
	id,
    last_name,
    email
FROM
	users
WHERE id NOT IN(
    SELECT
		user_id
	FROM
		orders
	WHERE
		order_time BETWEEN '2017-12-01 00:00:00'
		AND '2018-01-01 00:00:00')
;

-- 105 演習:2017年12月に商品を購入したユーザを出力
SELECT
	id,
    last_name,
    email
FROM
	users
WHERE id IN (
    SELECT
		user_id
	FROM
		orders
	WHERE
		order_time BETWEEN '2017-12-01 00:00:00'
		AND '2018-01-01 00:00:00')
;

-- 107 応用:全商品の平均単価より、単価が高い商品を単価が高い順に出力。同額の場合はid順。
SELECT
	*
FROM
	products
WHERE 
	price > 
    (
		SELECT
			avg(price)
		FROM
			products
	)
ORDER BY
	price desc,
    id
;

-- 109 ユーザのアクティビティ度合いを出力
SELECT
	u.id user_id,
    count(*) num,
    CASE
		WHEN count(*) >= 5 then'A'
		WHEN count(*) >= 2 then'B'
		ELSE 'C'
	END user_rank
FROM
	users u
INNER JOIN
	orders o
    on u.id = o.user_id
GROUP BY u.id
ORDER BY user_rank
;

-- 112 演習:累計販売個数でランク分けしてランクが高い順に出力
SELECT
	p.id product_id,
    p.name product_name,
    sum(product_qty),
    CASE
		WHEN sum(product_qty) >= 20 then'A'
		WHEN sum(product_qty) >= 10 then'B'
		ELSE 'C'
	END product_rank
FROM
	products p
INNER JOIN
	order_details od
    on p.id = od.product_id
GROUP BY p.id
ORDER BY product_rank
;

-- 115 平均客単価を出力。少数第一位で出力
SELECT
	ROUND(AVG(amount), 0)
FROM
	orders
;

-- 116 月別平均客単価を出力。少数第一位で出力
SELECT
	date_format(order_time, '%Y%m') order_year_month,
	ROUND(AVG(amount), 0) average_customer_spend
FROM
	orders
GROUP BY
	date_format(order_time, '%Y%m')
ORDER BY order_year_month
;

-- 117 都道府県別平均客単価を出力
SELECT
	u.prefecture_id,
	pr.name prefecture_name,
	ROUND(AVG(amount), 0) average_customer
FROM
	users u
INNER JOIN
	prefectures pr
	ON u.prefecture_id = pr.id
INNER JOIN
	orders o
    ON u.id = o.user_id
GROUP BY u.prefecture_id
ORDER BY u.prefecture_id
;

-- 118 都道府県別月別平均客単価を出力 
SELECT
	u.prefecture_id prefecture_id,
	pr.name prefecture_name,
    date_format(order_time, '%Y%m') order_year_month,
	ROUND(AVG(amount), 0) average_customer
FROM
	users u
INNER JOIN
	prefectures pr
	ON u.prefecture_id = pr.id
INNER JOIN
	orders o
    ON u.id = o.user_id
GROUP BY prefecture_id, order_year_month
ORDER BY prefecture_id, order_year_month
;

-- 120 新商品を一件追加してほしい
INSERT INTO products (name, price) values ('新商品A', 1000);

-- 122 新商品を三件追加してほしい
INSERT INTO 
    products (name, price) 
values 
    ('新商品C', 3000),
    ('新商品D', 4000),
    ('新商品E', 5000)
;

-- 123 全商品の値段を１０％OFFにする
UPDATE products SET price = price * 0.9;

-- 124 商品idが３の商品名を「SQL入門」に変更
UPDATE products SET name  = 'SQL入門' WHERE id = 3;

-- 125 累計販売数が１０を超えている商品のみ価格を５％増やす
UPDATE 
    products
SET
    price = price * 1.05
WHERE 
    id IN
    (
    SELECT
        product_id
    FROM
        order_details
    GROUP BY
        product_id
    HAVING
        sum(product_qty) >= 10
    )
;
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
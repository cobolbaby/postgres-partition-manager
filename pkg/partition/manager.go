package partition

/*
	PG Native分区表操作须知:

	// 查看所有分区表
	SELECT distinct inhparent::regclass FROM pg_inherits

	// 查看某张表的所有子分区以及子分区定义
	// -- 依赖 ::regclass
	SELECT
		i.inhparent::regclass AS parent_name,
		t.inhrelid::regclass AS part_name,
		pg_get_expr(t.relpartbound, t.oid) AS part_expr
	FROM pg_inherits AS i
	JOIN pg_class AS t ON t.oid = i.inhrelid
	WHERE i.inhparent = 'ictf3.ictlogtestpart_ao'::regclass;
	// -- 依赖 join
	SELECT
		nmsp_parent.nspname AS parent_schema,
		parent.relname      AS parent,
		child.relname       AS child
	FROM pg_inherits
		JOIN pg_class parent            ON pg_inherits.inhparent = parent.oid
		JOIN pg_class child             ON pg_inherits.inhrelid   = child.oid
		JOIN pg_namespace nmsp_parent   ON nmsp_parent.oid  = parent.relnamespace
	WHERE parent.relname='parent_table_name';

	// 判断分区键是否合理
	SELECT DATE_TRUNC('year', view_date)::DATE, COUNT(*) FROM website_views GROUP BY 1 ORDER BY 1;

	// 既有工具，更推荐
	// https://www.citusdata.com/blog/2018/01/24/citus-and-pg-partman-creating-a-scalable-time-series-database-on-postgresql/

	// SQL工具脚本，参考如下连接改造
	// https://www.cybertec-postgresql.com/en/partition-management-do-you-really-need-a-tool-for-that/

	WITH q_last_part AS (
		select
			*,
			((regexp_match(part_expr, $$ TO \('(.*)'\)$$))[1])::timestamp without time zone as last_part_end
		from
			(
				select
					format('%I.%I', n.nspname, p.relname) as parent_name,
					format('%I.%I', n.nspname, c.relname) as part_name,
					pg_catalog.pg_get_expr(c.relpartbound, c.oid) as part_expr
				from
					pg_class p
					join pg_inherits i ON i.inhparent = p.oid
					join pg_class c on c.oid = i.inhrelid
					join pg_namespace n on n.oid = c.relnamespace
				where
					p.relname = 'ictlogtestpart_ao'
					and n.nspname = 'ictf3'
					and p.relkind = 'p'
			) x
		order by
			last_part_end desc
		limit
			1
	)
	SELECT
		format(
			$$CREATE TABLE IF NOT EXISTS %s_%s%s%s PARTITION OF %s FOR VALUES FROM ('%s') TO ('%s')$$,
			parent_name,
			extract(year from last_part_end),
			lpad((extract(month from last_part_end))::text, 2, '0'),
			lpad((extract(day from last_part_end))::text, 2, '0'),
			parent_name,
			last_part_end,
			last_part_end + '1 day' :: interval
		) AS sql_to_exec
	FROM
		q_last_part;
	-- 变量是 表名 ictf3.ictlogtestpart_ao 和 分区时间范围 1 day

	// 过期的 Part
	WITH q_expired_part AS (
		select
			*,
			((regexp_match(part_expr, $$ TO \('(.*)'\)$$))[1])::timestamp without time zone as part_end
		from
			(
				select
					format('%I.%I', n.nspname, p.relname) as parent_name,
					format('%I.%I', n.nspname, c.relname) as part_name,
					pg_catalog.pg_get_expr(c.relpartbound, c.oid) as part_expr
				from
					pg_class p
					join pg_inherits i ON i.inhparent = p.oid
					join pg_class c on c.oid = i.inhrelid
					join pg_namespace n on n.oid = c.relnamespace
				where
					p.relname = 'ictlogtestpart_ao'
					and n.nspname = 'ictf3'
					and p.relkind = 'p'
			) x
	)
	SELECT
		format('DROP TABLE IF EXISTS %s', part_name) as sql_to_exec
	FROM
		q_expired_part
	WHERE
		part_end < CURRENT_DATE - '7 days'::interval;
		and part_name !~* 'his$';
	-- 变量是 表名 ictf3.ictlogtestpart_ao 和 过期定义 7 day，但得刨去默认分区

*/

import (
	"fmt"
)

func Add(args []string) {
	fmt.Println("add partition")
}

func add() error {
	return nil
}

func Drop(args []string) {
	fmt.Println("drop partition")
}

func drop() error {
	return nil
}

func Migrate(args []string) {
	fmt.Println("migrate partition")
}

func migrate() error {
	return nil
}

// TODO:
func Autopilot() {
	fmt.Println("autopilot partition")

	// Check Partition Status

	// Add necessary partitions

	// Migrate partitions to Minio or Greenplum

	// Drop unnecessary partitions

}

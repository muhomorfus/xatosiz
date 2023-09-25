package migrations

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddBaseTables, downAddBaseTables)
}

func upAddBaseTables(tx *sql.Tx) error {
	_, err := tx.Exec(fmt.Sprintf(`
		create table xatosiz.group_queue (
			uuid UUID
		) engine = Kafka()
		SETTINGS
			kafka_broker_list = '%s',
			kafka_topic_list = 'groups',
			kafka_group_name = 'clickhouse-groups',
			kafka_format = 'JSONEachRow',
		    kafka_flush_interval_ms = 1000;
	`, Brokers))
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		create table xatosiz.group (
			uuid UUID, 
			primary key(uuid)
		) engine = MergeTree() order by (uuid);
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		create materialized view xatosiz.group_mv to xatosiz.group as
		select *
		from xatosiz.group_queue;
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		create table xatosiz.trace_queue (
		uuid UUID,
		group_uuid UUID,
		parent_uuid Nullable(UUID),
		title String,
		time_start DateTime64(9, 'UTC'),
		time_end Nullable(DateTime64(9, 'UTC')),
		component String
		) engine = Kafka()
		SETTINGS
			kafka_broker_list = '%s',
			kafka_topic_list = 'traces',
			kafka_group_name = 'clickhouse-traces',
			kafka_format = 'JSONEachRow',
		    kafka_flush_interval_ms = 1000;
	`, Brokers))
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		create table xatosiz.trace_raw (
			uuid UUID,
			group_uuid UUID,
			parent_uuid UUID,
			title String,
			time_start DateTime64(9, 'UTC'),
			time_end Nullable(DateTime64(9, 'UTC')),
			component String,
			updated_at DateTime64(9, 'UTC'),
			primary key(uuid)
		) engine = MergeTree() order by (uuid);
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		create materialized view xatosiz.trace_mv to xatosiz.trace_raw as
		select
		    uuid,
		    group_uuid,
		    parent_uuid,
		    title,
		    time_start,
		    time_end,
		    component,
		    _timestamp_ms as updated_at
		from xatosiz.trace_queue;
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		create view xatosiz.trace as
		select
		    tr.uuid as uuid,
		    tr.group_uuid as group_uuid,
		    tr.parent_uuid as parent_uuid,
		    tr.title as title,
		    tr.time_start as time_start,
		    tr.time_end as time_end,
		    tr.component as component
		from
		xatosiz.trace_raw tr
		inner join
		(select uuid as uuid, max(updated_at) as updated_at from xatosiz.trace_raw group by uuid) tu
		on tr.uuid = tu.uuid and tr.updated_at = tu.updated_at;
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		create table xatosiz.event_queue (
			uuid UUID,
			group_uuid UUID,
			trace_uuid UUID,
			message String,
			time DateTime64(9, 'UTC'),
			priority String,
			payload Nullable(String),
			fixed Boolean
		) engine = Kafka()
		SETTINGS
			kafka_broker_list = '%s',
			kafka_topic_list = 'events',
			kafka_group_name = 'clickhouse-events',
			kafka_format = 'JSONEachRow',
		    kafka_flush_interval_ms = 1000;
	`, Brokers))
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		create table xatosiz.event_raw (
			uuid UUID,
			group_uuid UUID,
			trace_uuid UUID,
			message String,
			time DateTime64(9, 'UTC'),
			priority String,
			payload Nullable(String),
			fixed Boolean,
			updated_at DateTime64(9, 'UTC'),
			primary key(uuid)
		) engine = MergeTree() order by (uuid);
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		create materialized view xatosiz.event_mv to xatosiz.event_raw as
		select 
		    uuid,
		    group_uuid,
		    trace_uuid,
		    message,
		    time,
		    priority,
		    payload,
		    fixed,
		    _timestamp_ms as updated_at
		from xatosiz.event_queue;
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		create view xatosiz.event as
		select
		    er.uuid as uuid,
		    er.group_uuid as group_uuid,
		    er.trace_uuid as trace_uuid,
		    er.message as message,
		    er.time as time,
		    er.priority as priority,
		    er.payload as payload,
		    er.fixed as fixed
		from
		xatosiz.event_raw er
		inner join
		(select uuid as uuid, max(updated_at) as updated_at from xatosiz.event_raw group by uuid) eu
		on er.uuid = eu.uuid and er.updated_at = eu.updated_at;
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		create table xatosiz.alert_hit_queue (
			uuid UUID,
			config_uuid UUID,
			time DateTime64(9, 'UTC')
		) engine = Kafka()
		SETTINGS
			kafka_broker_list = '%s',
			kafka_topic_list = 'alert_hits',
			kafka_group_name = 'clickhouse-alert-hits',
			kafka_format = 'JSONEachRow',
		    kafka_flush_interval_ms = 1000;
	`, Brokers))
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		create table xatosiz.alert_hit (
			uuid UUID,
			config_uuid UUID,
			time DateTime64(9, 'UTC'),
			primary key(uuid)
		) engine = MergeTree() order by (uuid);
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		create materialized view xatosiz.alert_hit_mv to xatosiz.alert_hit as
		select *
		from xatosiz.alert_hit_queue;
	`)
	if err != nil {
		return err
	}

	return nil
}

func downAddBaseTables(tx *sql.Tx) error {
	_, err := tx.Exec("drop view if exists xatosiz.alert_hit_mv;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop table if exists xatosiz.alert_hit;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop table if exists xatosiz.alert_hit_queue;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop view if exists xatosiz.event;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop view if exists xatosiz.event_mv;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop table if exists xatosiz.event_raw;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop table if exists xatosiz.event_queue;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop view if exists xatosiz.trace;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop view if exists xatosiz.trace_mv;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop table if exists xatosiz.trace_raw;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop table if exists xatosiz.trace_queue;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop view if exists xatosiz.group_mv;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop table if exists xatosiz.group;")
	if err != nil {
		return err
	}

	_, err = tx.Exec("drop table if exists xatosiz.group_queue;")
	if err != nil {
		return err
	}

	return nil
}

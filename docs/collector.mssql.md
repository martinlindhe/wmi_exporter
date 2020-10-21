# mssql collector

The mssql collector exposes metrics about the MSSQL server

|||
-|-
Metric name prefix  | `mssql`
Classes             | [`Win32_PerfRawData_MSSQLSERVER_SQLServerAccessMethods`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-access-methods-object)<br/>[`Win32_PerfRawData_MSSQLSERVER_SQLServerAvailabilityReplica`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-availability-replica)<br/>[`Win32_PerfRawData_MSSQLSERVER_SQLServerBufferManager`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-buffer-manager-object)<br/>[`Win32_PerfRawData_MSSQLSERVER_SQLServerDatabaseReplica`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-database-replica)<br/>[`Win32_PerfRawData_MSSQLSERVER_SQLServerDatabases`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-databases-object?view=sql-server-2017)<br/>[`Win32_PerfRawData_MSSQLSERVER_SQLServerGeneralStatistics`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-general-statistics-object)<br/>[`Win32_PerfRawData_MSSQLSERVER_SQLServerLocks`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-locks-object)<br/>[`Win32_PerfRawData_MSSQLSERVER_SQLServerMemoryManager`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-memory-manager-object)<br/>[`Win32_PerfRawData_MSSQLSERVER_SQLServerSQLStatistics`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-sql-statistics-object)<br/>[`Win32_PerfRawData_MSSQLSERVER_SQLServerSQLErrors`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-sql-errors-object)<br/>[`Win32_PerfRawData_MSSQLSERVER_SQLServerTransactions`](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-transactions-object)
Enabled by default? | No

## Flags

### `--collectors.mssql.classes-enabled`

Comma-separated list of MSSQL WMI classes to use. Supported values are `accessmethods`, `availreplica`, `bufman`, `databases`, `dbreplica`, `genstats`, `locks`, `memmgr`, `sqlstats`, `sqlerrors` and `transactions`.

### `--collectors.mssql.class-print`

If true, print available mssql WMI classes and exit.  Only displays if the mssql collector is enabled.

## Metrics

Name | Description | Type | Labels
-----|-------------|------|-------
`windows_mssql_collector_duration_seconds` | The time taken for each sub-collector to return | counter | `collector`, `instance`
`windows_mssql_collector_success` | 1 if sub-collector succeeded, 0 otherwise | counter | `collector`, `instance`
`windows_mssql_accessmethods_au_batch_cleanups` | The total number of batches that were completed successfully by the background task that cleans up deferred dropped allocation units | counter | `instance`
`windows_mssql_accessmethods_au_cleanups` | The total number of allocation units that were successfully dropped the background task that cleans up deferred dropped allocation units. Each allocation unit drop requires multiple batches | counter | `instance`
`windows_mssql_accessmethods_by_reference_lob_creates` | The total count of large object (lob) values that were passed by reference. By-reference lobs are used in certain bulk operations to avoid the cost of passing them by value | counter | `instance`
`windows_mssql_accessmethods_by_reference_lob_uses` | The total count of by-reference lob values that were used. By-reference lobs are used in certain bulk operations to avoid the cost of passing them by-value | counter | `instance`
`windows_mssql_accessmethods_lob_read_aheads` | The total count of lob pages on which readahead was issued | counter | `instance`
`windows_mssql_accessmethods_column_value_pulls` | The total count of column values that were pulled in-row from off-row | counter | `instance`
`windows_mssql_accessmethods_column_value_pushes` | The total count of column values that were pushed from in-row to off-row | counter | `instance`
`windows_mssql_accessmethods_deferred_dropped_aus` | The total number of allocation units waiting to be dropped by the background task that cleans up deferred dropped allocation units | counter | `instance`
`windows_mssql_accessmethods_deferred_dropped_rowsets` | The number of rowsets created as a result of aborted online index build operations that are waiting to be dropped by the background task that cleans up deferred dropped rowsets | counter | `instance`
`windows_mssql_accessmethods_dropped_rowset_cleanups` | The number of rowsets per second created as a result of aborted online index build operations that were successfully dropped by the background task that cleans up deferred dropped rowsets | counter | `instance`
`windows_mssql_accessmethods_dropped_rowset_skips` | The number of rowsets per second created as a result of aborted online index build operations that were skipped by the background task that cleans up deferred dropped rowsets created | counter | `instance`
`windows_mssql_accessmethods_extent_deallocations` | Number of extents deallocated per second in all databases in this instance of SQL Server | counter | `instance`
`windows_mssql_accessmethods_extent_allocations` | Number of extents allocated per second in all databases in this instance of SQL Server | counter | `instance`
`windows_mssql_accessmethods_au_batch_cleanup_failures` | The number of batches per second that failed and required retry, by the background task that cleans up deferred dropped allocation units. Failure could be due to lack of memory or disk space, hardware failure and other reasons | counter | `instance`
`windows_mssql_accessmethods_leaf_page_cookie_failures` | The number of times that a leaf page cookie could not be used during an index search since changes happened on the leaf page. The cookie is used to speed up index search | counter | `instance`
`windows_mssql_accessmethods_tree_page_cookie_failures` | The number of times that a tree page cookie could not be used during an index search since changes happened on the parent pages of those tree pages. The cookie is used to speed up index search | counter | `instance`
`windows_mssql_accessmethods_forwarded_records` | Number of records per second fetched through forwarded record pointers | counter | `instance`
`windows_mssql_accessmethods_free_space_page_fetches` | Number of pages fetched per second by free space scans. These scans search for free space within pages already allocated to an allocation unit, to satisfy requests to insert or modify record fragments | counter | `instance`
`windows_mssql_accessmethods_free_space_scans` | Number of scans per second that were initiated to search for free space within pages already allocated to an allocation unit to insert or modify record fragment. Each scan may find multiple pages | counter | `instance`
`windows_mssql_accessmethods_full_scans` | Number of unrestricted full scans per second. These can be either base-table or full-index scans | counter | `instance`
`windows_mssql_accessmethods_index_searches` | Number of index searches per second. These are used to start a range scan, reposition a range scan, revalidate a scan point, fetch a single index record, and search down the index to locate where to insert a new row | counter | `instance`
`windows_mssql_accessmethods_insysxact_waits` | Number of times a reader needs to wait for a page because the InSysXact bit is set | counter | `instance`
`windows_mssql_accessmethods_lob_handle_creates` | Count of temporary lobs created | counter | `instance`
`windows_mssql_accessmethods_lob_handle_destroys` | Count of temporary lobs destroyed | counter | `instance`
`windows_mssql_accessmethods_lob_ss_provider_creates` | Count of LOB Storage Service Providers (LobSSP) created. One worktable created per LobSSP | counter | `instance`
`windows_mssql_accessmethods_lob_ss_provider_destroys` | Count of LobSSP destroyed | counter | `instance`
`windows_mssql_accessmethods_lob_ss_provider_truncations` | Count of LobSSP truncated | counter | `instance`
`windows_mssql_accessmethods_mixed_page_allocations` | Number of pages allocated per second from mixed extents. These could be used for storing the IAM pages and the first eight pages that are allocated to an allocation unit | counter | `instance`
`windows_mssql_accessmethods_page_compression_attempts` | Number of pages evaluated for page-level compression. Includes pages that were not compressed because significant savings could be achieved. Includes all objects in the instance of SQL Server | counter | `instance`
`windows_mssql_accessmethods_page_deallocations` | Number of pages deallocated per second in all databases in this instance of SQL Server. These include pages from mixed extents and uniform extents | counter | `instance`
`windows_mssql_accessmethods_page_allocations` | Number of pages allocated per second in all databases in this instance of SQL Server. These include pages allocations from both mixed extents and uniform extents | counter | `instance`
`windows_mssql_accessmethods_page_compressions` | Number of data pages that are compressed by using PAGE compression. Includes all objects in the instance of SQL Server | counter | `instance`
`windows_mssql_accessmethods_page_splits` | Number of page splits per second that occur as the result of overflowing index pages | counter | `instance`
`windows_mssql_accessmethods_probe_scans` | Number of probe scans per second that are used to find at most one single qualified row in an index or base table directly | counter | `instance`
`windows_mssql_accessmethods_range_scans` | Number of qualified range scans through indexes per second | counter | `instance`
`windows_mssql_accessmethods_scan_point_revalidations` | Number of times per second that the scan point had to be revalidated to continue the scan | counter | `instance`
`windows_mssql_accessmethods_ghost_record_skips` | Number of ghosted records per second skipped during scans | counter | `instance`
`windows_mssql_accessmethods_table_lock_escalations` | Number of times locks on a table were escalated to the TABLE or HoBT granularity | counter | `instance`
`windows_mssql_accessmethods_leaf_page_cookie_uses` | Number of times a leaf page cookie is used successfully during an index search since no change happened on the leaf page. The cookie is used to speed up index search | counter | `instance`
`windows_mssql_accessmethods_tree_page_cookie_uses` | Number of times a tree page cookie is used successfully during an index search since no change happened on the parent page of the tree page. The cookie is used to speed up index search | counter | `instance`
`windows_mssql_accessmethods_workfile_creates` | Number of work files created per second. For example, work files could be used to store temporary results for hash joins and hash aggregates | counter | `instance`
`windows_mssql_accessmethods_worktables_creates` | Number of work tables created per second. For example, work tables could be used to store temporary results for query spool, lob variables, XML variables, and cursors | counter | `instance`
`windows_mssql_accessmethods_worktables_from_cache_ratio` | Percentage of work tables created where the initial two pages of the work table were not allocated but were immediately available from the work table cache | counter | `instance`
`windows_mssql_availreplica_received_from_replica_bytes` | Number of bytes received from the availability replica per second. Pings and status updates will generate network traffic even on databases with no user updates | counter | `instance`, `replica`
`windows_mssql_availreplica_sent_to_replica_bytes` | Number of bytes sent to the remote availability replica per second. On the primary replica this is the number of bytes sent to the secondary replica. On the secondary replica this is the number of bytes sent to the primary replica | counter | `instance`, `replica`
`windows_mssql_availreplica_sent_to_transport_bytes` | Actual number of bytes sent per second over the network to the remote availability replica. On the primary replica this is the number of bytes sent to the secondary replica. On the secondary replica this is the number of bytes sent to the primary replica | counter | `instance`, `replica`
`windows_mssql_availreplica_initiated_flow_controls` | Time in milliseconds that log stream messages waited for send flow control, in the last second | counter | `instance`, `replica`
`windows_mssql_availreplica_flow_control_wait_seconds` | Number of times flow-control initiated in the last second. Flow Control Time (ms/sec) divided by Flow Control/sec is the average time per wait | counter | `instance`, `replica`
`windows_mssql_availreplica_receives_from_replica` | Number of Always On messages received from thereplica per second | counter | `instance`, `replica`
`windows_mssql_availreplica_resent_messages` | Number of Always On messages resent in the last second | counter | `instance`, `replica`
`windows_mssql_availreplica_sends_to_replica` | Number of Always On messages sent to this availability replica per second | counter | `instance`, `replica`
`windows_mssql_availreplica_sends_to_transport` | Actual number of Always On messages sent per second over the network to the remote availability replica | counter | `instance`, `replica`
`windows_mssql_bufman_background_writer_pages` | Number of pages flushed to enforce the recovery interval settings | counter | `instance`
`windows_mssql_bufman_buffer_cache_hit_ratio` | Indicates the percentage of pages found in the buffer cache without having to read from disk. The ratio is the total number of cache hits divided by the total number of cache lookups over the last few thousand page accesses | counter | `instance`
`windows_mssql_bufman_checkpoint_pages` | Indicates the number of pages flushed to disk per second by a checkpoint or other operation that require all dirty pages to be flushed | counter | `instance`
`windows_mssql_bufman_database_pages` | Indicates the number of pages in the buffer pool with database content | counter | `instance`
`windows_mssql_bufman_extension_allocated_pages` | Total number of non-free cache pages in the buffer pool extension file | counter | `instance`
`windows_mssql_bufman_extension_free_pages` | Total number of free cache pages in the buffer pool extension file | counter | `instance`
`windows_mssql_bufman_extension_in_use_as_percentage` | _Not yet documented_ | counter | `instance`
`windows_mssql_bufman_extension_outstanding_io` | Percentage of the buffer pool extension paging file occupied by buffer manager pages | counter | `instance`
`windows_mssql_bufman_extension_page_evictions` | Number of pages evicted from the buffer pool extension file per second | counter | `instance`
`windows_mssql_bufman_extension_page_reads` | Number of pages read from the buffer pool extension file per second | counter | `instance`
`windows_mssql_bufman_extension_page_unreferenced_seconds` | Average seconds a page will stay in the buffer pool extension without references to it | counter | `instance`
`windows_mssql_bufman_extension_page_writes` | Number of pages written to the buffer pool extension file per second | counter | `instance`
`windows_mssql_bufman_free_list_stalls` | Indicates the number of requests per second that had to wait for a free page | counter | `instance`
`windows_mssql_bufman_integral_controller_slope` | The slope that integral controller for the buffer pool last used, times -10 billion | counter | `instance`
`windows_mssql_bufman_lazywrites` | Indicates the number of buffers written per second by the buffer manager's lazy writer | counter | `instance`
`windows_mssql_bufman_page_life_expectancy_seconds` | Indicates the number of seconds a page will stay in the buffer pool without references | counter | `instance`
`windows_mssql_bufman_page_lookups` | Indicates the number of requests per second to find a page in the buffer pool | counter | `instance`
`windows_mssql_bufman_page_reads` | Indicates the number of physical database page reads that are issued per second | counter | `instance`
`windows_mssql_bufman_page_writes` | Indicates the number of physical database page writes that are issued per second | counter | `instance`
`windows_mssql_bufman_read_ahead_pages` | Indicates the number of pages read per second in anticipation of use | counter | `instance`
`windows_mssql_bufman_read_ahead_issuing_seconds` | Time (microseconds) spent issuing readahead | counter | `instance`
`windows_mssql_bufman_target_pages` | Ideal number of pages in the buffer pool | counter | `instance`
`windows_mssql_dbreplica_database_flow_control_wait_seconds` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_database_initiated_flow_controls` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_received_file_bytes` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_group_commits` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_group_commit_stall_seconds` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_apply_pending_queue` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_apply_ready_queue` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_compressed_bytes` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_decompressed_bytes` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_received_bytes` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_compression_cachehits` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_compression_cachemisses` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_compressions` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_decompressions` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_remaining_for_undo` | The amount of log, in bytes, remaining to complete the undo phase | counter | `instance`, `replica`
`windows_mssql_dbreplica_log_send_queue` | Amount of log records in the log files of the primary database, in kilobytes, that haven't been sent to the secondary replica | counter | `instance`, `replica`
`windows_mssql_dbreplica_mirrored_write_transactions` | Number of transactions that were written to the primary database and then waited to commit until the log was sent to the secondary database, in the last second | counter | `instance`, `replica`
`windows_mssql_dbreplica_recovery_queue_records` | Amount of log records in the log files of the secondary replica that have not been redone | counter | `instance`, `replica`
`windows_mssql_dbreplica_redo_blocks` | Number of times the redo thread was blocked on locks held by readers of the database | counter | `instance`, `replica`
`windows_mssql_dbreplica_redo_remaining_bytes` | The amount of log, in kilobytes, remaining to be redone to finish the reverting phase | counter | `instance`, `replica`
`windows_mssql_dbreplica_redone_bytes` | Amount of log records redone on the secondary database in the last second | counter | `instance`, `replica`
`windows_mssql_dbreplica_redones` | _Not yet documented_ | counter | `instance`, `replica`
`windows_mssql_dbreplica_total_log_requiring_undo` | Total kilobytes of log that must be undone | counter | `instance`, `replica`
`windows_mssql_dbreplica_transaction_delay_seconds` | Delay in waiting for unterminated commit acknowledgment for all the current transactions | counter | `instance`, `replica`
`windows_mssql_databases_active_transactions` | Number of active transactions for the database | counter | `instance`, `database`
`windows_mssql_databases_backup_restore_operations` | Read/write throughput for backup and restore operations of a database per second | counter | `instance`, `database`
`windows_mssql_databases_bulk_copy_rows` | Number of rows bulk copied per second | counter | `instance`, `database`
`windows_mssql_databases_bulk_copy_bytes` | Amount of data bulk copied (in kilobytes) per second | counter | `instance`, `database`
`windows_mssql_databases_commit_table_entries` | he size (row count) of the in-memory portion of the commit table for the database | counter | `instance`, `database`
`windows_mssql_databases_data_files_size_bytes` | Cumulative size (in kilobytes) of all the data files in the database including any automatic growth. Monitoring this counter is useful, for example, for determining the correct size of tempdb | counter | `instance`, `database`
`windows_mssql_databases_dbcc_logical_scan_bytes` | Number of logical read scan bytes per second for database console commands (DBCC) | counter | `instance`, `database`
`windows_mssql_databases_group_commit_stall_seconds` | Group stall time (microseconds) per second | counter | `instance`, `database`
`windows_mssql_databases_log_flushed_bytes` | Total number of log bytes flushed | counter | `instance`, `database`
`windows_mssql_databases_log_cache_hit_ratio` | Percentage of log cache reads satisfied from the log cache | counter | `instance`, `database`
`windows_mssql_databases_log_cache_reads` | Reads performed per second through the log manager cache | counter | `instance`, `database`
`windows_mssql_databases_log_files_size_bytes` | Cumulative size (in kilobytes) of all the transaction log files in the database | counter | `instance`, `database`
`windows_mssql_databases_log_files_used_size_bytes` | The cumulative used size of all the log files in the database | counter | `instance`, `database`
`windows_mssql_databases_log_flushes` | Total wait time (in milliseconds) to flush the log. On an Always On secondary database, this value indicates the wait time for log records to be hardened to disk | counter | `instance`, `database`
`windows_mssql_databases_log_flush_waits` | Number of commits per second waiting for the log flush | counter | `instance`, `database`
`windows_mssql_databases_log_flush_wait_seconds` | Number of commits per second waiting for the log flush | counter | `instance`, `database`
`windows_mssql_databases_log_flush_write_seconds` | Time in milliseconds for performing writes of log flushes that were completed in the last second | counter | `instance`, `database`
`windows_mssql_databases_log_growths` | Total number of times the transaction log for the database has been expanded | counter | `instance`, `database`
`windows_mssql_databases_log_pool_cache_misses` | Number of requests for which the log block was not available in the log pool | counter | `instance`, `database`
`windows_mssql_databases_log_pool_disk_reads` | Number of disk reads that the log pool issued to fetch log blocks | counter | `instance`, `database`
`windows_mssql_databases_log_pool_hash_deletes` | Rate of raw hash entry deletes from the Log Pool | counter | `instance`, `database`
`windows_mssql_databases_log_pool_hash_inserts` | Rate of raw hash entry inserts into the Log Pool | counter | `instance`, `database`
`windows_mssql_databases_log_pool_invalid_hash_entries` | Rate of hash lookups failing due to being invalid | counter | `instance`, `database`
`windows_mssql_databases_log_pool_log_scan_pushes` | Rate of Log block pushes by log scans, which may come from disk or memory | counter | `instance`, `database`
`windows_mssql_databases_log_pool_log_writer_pushes` | Rate of Log block pushes by log writer thread | counter | `instance`, `database`
`windows_mssql_databases_log_pool_empty_free_pool_pushes` | Rate of Log block push fails due to empty free pool | counter | `instance`, `database`
`windows_mssql_databases_log_pool_low_memory_pushes` | Rate of Log block push fails due to being low on memory | counter | `instance`, `database`
`windows_mssql_databases_log_pool_no_free_buffer_pushes` | Rate of Log block push fails due to free buffer unavailable | counter | `instance`, `database`
`windows_mssql_databases_log_pool_req_behind_trunc` | Log pool cache misses due to block requested being behind truncation LSN | counter | `instance`, `database`
`windows_mssql_databases_log_pool_requests_old_vlf` | Log Pool requests that were not in the last VLF of the log | counter | `instance`, `database`
`windows_mssql_databases_log_pool_requests` | The number of log-block requests processed by the log pool | counter | `instance`, `database`
`windows_mssql_databases_log_pool_total_active_log_bytes` | Current total active log stored in the shared cache buffer manager in bytes | counter | `instance`, `database`
`windows_mssql_databases_log_pool_total_shared_pool_bytes` | Current total memory usage of the shared cache buffer manager in bytes | counter | `instance`, `database`
`windows_mssql_databases_log_shrinks` | Total number of log shrinks for this database | counter | `instance`, `database`
`windows_mssql_databases_log_truncations` | The number of times the transaction log has been truncated (in Simple Recovery Model) | counter | `instance`, `database`
`windows_mssql_databases_log_used_percent` | Percentage of space in the log that is in use | counter | `instance`, `database`
`windows_mssql_databases_pending_repl_transactions` | Number of transactions in the transaction log of the publication database marked for replication, but not yet delivered to the distribution database | counter | `instance`, `database`
`windows_mssql_databases_repl_transactions` | Number of transactions per second read out of the transaction log of the publication database and delivered to the distribution database | counter | `instance`, `database`
`windows_mssql_databases_shrink_data_movement_bytes` | Amount of data being moved per second by autoshrink operations, or DBCC SHRINKDATABASE or DBCC SHRINKFILE statements | counter | `instance`, `database`
`windows_mssql_databases_tracked_transactions` | Number of committed transactions recorded in the commit table for the database | counter | `instance`, `database`
`windows_mssql_databases_transactions` | Number of transactions started for the database per second | counter | `instance`, `database`
`windows_mssql_databases_write_transactions` | Number of transactions that wrote to the database and committed, in the last second | counter | `instance`, `database`
`windows_mssql_databases_xtp_controller_dlc_fetch_latency_seconds` | Average latency in microseconds between log blocks entering the Direct Log Consumer and being retrieved by the XTP controller, per second | counter | `instance`, `database`
`windows_mssql_databases_xtp_controller_dlc_peak_latency_seconds` | The largest recorded latency, in microseconds, of a fetch from the Direct Log Consumer by the XTP controller | counter | `instance`, `database`
`windows_mssql_databases_xtp_controller_log_processed_bytes` | The amount of log bytes processed by the XTP controller thread, per second | counter | `instance`, `database`
`windows_mssql_databases_xtp_memory_used_bytes` | The amount of memory used by XTP in the database | counter | `instance`, `database`
`windows_mssql_genstats_active_temp_tables` | Number of temporary tables/table variables in use | counter | `instance`
`windows_mssql_genstats_connection_resets` | Total number of logins started from the connection pool | counter | `instance`
`windows_mssql_genstats_event_notifications_delayed_drop` | Number of event notifications waiting to be dropped by a system thread | counter | `instance`
`windows_mssql_genstats_http_authenticated_requests` | Number of authenticated HTTP requests started per second | counter | `instance`
`windows_mssql_genstats_logical_connections` | Number of logical connections to the system | counter | `instance`
`windows_mssql_genstats_logins` | Total number of logins started per second. This does not include pooled connections | counter | `instance`
`windows_mssql_genstats_logouts` | Total number of logout operations started per second | counter | `instance`
`windows_mssql_genstats_mars_deadlocks` | Number of MARS deadlocks detected | counter | `instance`
`windows_mssql_genstats_non_atomic_yields` | Number of non-atomic yields per second | counter | `instance`
`windows_mssql_genstats_blocked_processes` | Number of currently blocked processes | counter | `instance`
`windows_mssql_genstats_soap_empty_requests` | Number of empty SOAP requests started per second | counter | `instance`
`windows_mssql_genstats_soap_method_invocations` | Number of SOAP method invocations started per second | counter | `instance`
`windows_mssql_genstats_soap_session_initiate_requests` | Number of SOAP Session initiate requests started per second | counter | `instance`
`windows_mssql_genstats_soap_session_terminate_requests` | Number of SOAP Session terminate requests started per second | counter | `instance`
`windows_mssql_genstats_soapsql_requests` | Number of SOAP SQL requests started per second | counter | `instance`
`windows_mssql_genstats_soapwsdl_requests` | Number of SOAP Web Service Description Language requests started per second | counter | `instance`
`windows_mssql_genstats_sql_trace_io_provider_lock_waits` | Number of waits for the File IO Provider lock per second | counter | `instance`
`windows_mssql_genstats_tempdb_recovery_unit_ids_generated` | Number of duplicate tempdb recovery unit id generated | counter | `instance`
`windows_mssql_genstats_tempdb_rowset_ids_generated` | Number of duplicate tempdb rowset id generated | counter | `instance`
`windows_mssql_genstats_temp_tables_creations` | Number of temporary tables/table variables created per second | counter | `instance`
`windows_mssql_genstats_temp_tables_awaiting_destruction` | Number of temporary tables/table variables waiting to be destroyed by the cleanup system thread | counter | `instance`
`windows_mssql_genstats_trace_event_notification_queue_size` | Number of trace event notification instances waiting in the internal queue to be sent through Service Broker | counter | `instance`
`windows_mssql_genstats_transactions` | Number of transaction enlistments (local, DTC, bound all combined) | counter | `instance`
`windows_mssql_genstats_user_connections` | Counts the number of users currently connected to SQL Server | counter | `instance`
`windows_mssql_locks_average_wait_seconds` | Average amount of wait time (in milliseconds) for each lock request that resulted in a wait | counter | `instance`, `resource`
`windows_mssql_locks_lock_requests` | Number of new locks and lock conversions per second requested from the lock manager | counter | `instance`, `resource`
`windows_mssql_locks_lock_timeouts` | Number of lock requests per second that timed out, including requests for NOWAIT locks | counter | `instance`, `resource`
`windows_mssql_locks_lock_timeouts_excluding_NOWAIT` | Number of lock requests per second that timed out, but excluding requests for NOWAIT locks | counter | `instance`, `resource`
`windows_mssql_locks_lock_waits` | Total wait time (in milliseconds) for locks in the last second | counter | `instance`, `resource`
`windows_mssql_locks_lock_wait_seconds` | Number of lock requests per second that required the caller to wait | counter | `instance`, `resource`
`windows_mssql_locks_deadlocks` | Number of lock requests per second that resulted in a deadlock | counter | `instance`, `resource`
`windows_mssql_memmgr_connection_memory_bytes` | Specifies the total amount of dynamic memory the server is using for maintaining connections | counter | `instance`
`windows_mssql_memmgr_database_cache_memory_bytes` | Specifies the amount of memory the server is currently using for the database pages cache | counter | `instance`
`windows_mssql_memmgr_external_benefit_of_memory` | An internal estimation of the performance benefit from adding memory to a specific cache | counter | `instance`
`windows_mssql_memmgr_free_memory_bytes` | Specifies the amount of committed memory currently not used by the server | counter | `instance`
`windows_mssql_memmgr_granted_workspace_memory_bytes` | Specifies the total amount of memory currently granted to executing processes, such as hash, sort, bulk copy, and index creation operations | counter | `instance`
`windows_mssql_memmgr_lock_blocks` | Specifies the current number of lock blocks in use on the server (refreshed periodically). A lock block represents an individual locked resource, such as a table, page, or row | counter | `instance`
`windows_mssql_memmgr_allocated_lock_blocks` | Specifies the current number of allocated lock blocks. At server startup, the number of allocated lock blocks plus the number of allocated lock owner blocks depends on the SQL Server Locks configuration option. If more lock blocks are needed, the value increases | counter | `instance`
`windows_mssql_memmgr_lock_memory_bytes` | Specifies the total amount of dynamic memory the server is using for locks | counter | `instance`
`windows_mssql_memmgr_lock_owner_blocks` | Specifies the current number of allocated lock owner blocks. At server startup, the number of allocated lock owner blocks and the number of allocated lock blocks depend on the SQL Server Locks configuration option. If more lock owner blocks are needed, the value increases dynamically | counter | `instance`
`windows_mssql_memmgr_allocated_lock_owner_blocks` | _Not yet documented_ | counter | `instance`
`windows_mssql_memmgr_log_pool_memory_bytes` | Total amount of dynamic memory the server is using for Log Pool | counter | `instance`
`windows_mssql_memmgr_maximum_workspace_memory_bytes` | Indicates the maximum amount of memory available for executing processes, such as hash, sort, bulk copy, and index creation operations | counter | `instance`
`windows_mssql_memmgr_outstanding_memory_grants` | Specifies the total number of processes that have successfully acquired a workspace memory grant | counter | `instance`
`windows_mssql_memmgr_pending_memory_grants` | Specifies the total number of processes waiting for a workspace memory grant | counter | `instance`
`windows_mssql_memmgr_optimizer_memory_bytes` | Specifies the total amount of dynamic memory the server is using for query optimization | counter | `instance`
`windows_mssql_memmgr_reserved_server_memory_bytes` | ndicates the amount of memory the server has reserved for future usage. This counter shows the current unused amount of memory initially granted that is shown in Granted Workspace Memory | counter | `instance`
`windows_mssql_memmgr_sql_cache_memory_bytes` | Specifies the total amount of dynamic memory the server is using for the dynamic SQL cache | counter | `instance`
`windows_mssql_memmgr_stolen_server_memory_bytes` | Specifies the amount of memory the server is using for purposes other than database pages | counter | `instance`
`windows_mssql_memmgr_target_server_memory_bytes` | Indicates the ideal amount of memory the server can consume | counter | `instance`
`windows_mssql_memmgr_total_server_memory_bytes` | Specifies the amount of memory the server has committed using the memory manager | counter | `instance`
`windows_mssql_sqlstats_auto_parameterization_attempts` | Number of failed auto-parameterization attempts per second. This should be small. Note that auto-parameterizations are also known as simple parameterizations in later versions of SQL Server | counter | `instance`
`windows_mssql_sqlstats_batch_requests` | _Not yet documented_ | counter | `instance`
`windows_mssql_sqlstats_failed_auto_parameterization_attempts` | _Not yet documented_ | counter | `instance`
`windows_mssql_sqlstats_forced_parameterizations` | Number of successful forced parameterizations per second | counter | `instance`
`windows_mssql_sqlstats_guided_plan_executions` | Number of plan executions per second in which the query plan has been generated by using a plan guide | counter | `instance`
`windows_mssql_sqlstats_misguided_plan_executions` | Number of plan executions per second in which a plan guide could not be honored during plan generation | counter | `instance`
`windows_mssql_sqlstats_safe_auto_parameterization_attempts` | Number of safe auto-parameterization attempts per second | counter | `instance`
`windows_mssql_sqlstats_sql_attentions` | Number of attentions per second | counter | `instance`
`windows_mssql_sqlstats_sql_compilations` | Number of SQL compilations per second | counter | `instance`
`windows_mssql_sqlstats_sql_recompilations` | Number of statement recompiles per second | counter | `instance`
`windows_mssql_sqlstats_unsafe_auto_parameterization_attempts` | Number of unsafe auto-parameterization attempts per second. | counter | `instance`
`windows_mssql_sql_errors_total` | Information for all errors | counter | `instance`, `resource`
`windows_mssql_transactions_tempdb_free_space_bytes` | The amount of space (in kilobytes) available in tempdb | gauge | `instance`
`windows_mssql_transactions_longest_transaction_running_seconds` | The length of time (in seconds) since the start of the transaction that has been active longer than any other current transaction | gauge | `instance`
`windows_mssql_transactions_nonsnapshot_version_active_total` | The number of currently active transactions that are not using snapshot isolation level and have made data modifications that have generated row versions in the tempdb version store | counter | `instance`
`windows_mssql_transactions_snapshot_active_total` | The number of currently active transactions using the snapshot isolation level | counter | `instance`
`windows_mssql_transactions_active` | The number of currently active transactions of all types | gauge | `instance`
`windows_mssql_transactions_update_conflicts_total` | The percentage of those transactions using the snapshot isolation level that have encountered update conflicts within the last second | counter | `instance`
`windows_mssql_transactions_update_snapshot_active_total` | The number of currently active transactions using the snapshot isolation level and have modified data | counter | `instance`
`windows_mssql_transactions_version_cleanup_rate_bytes` | The rate (in kilobytes per second) at which row versions are removed from the snapshot isolation version store in tempdb | gauge | `instance`
`windows_mssql_transactions_version_generation_rate_bytes` | The rate (in kilobytes per second) at which new row versions are added to the snapshot isolation version store in tempdb | gauge | `instance`
`windows_mssql_transactions_version_store_size_bytes` | he amount of space (in kilobytes) in tempdb being used to store snapshot isolation level row versions | gauge | `instance`
`windows_mssql_transactions_version_store_units` | The number of active allocation units in the snapshot isolation version store in tempdb | counter | `instance`
`windows_mssql_transactions_version_store_creation_units` | The number of allocation units that have been created in the snapshot isolation store since the instance of the Database Engine was started | counter | `instance`
`windows_mssql_transactions_version_store_truncation_units` | The number of allocation units that have been removed from the snapshot isolation store since the instance of the Database Engine was started | counter | `instance`

### Example metric
_This collector does not yet have explained examples, we would appreciate your help adding them!_

## Useful queries

### Buffer Cache Hit Ratio

When you read the counter in perfmon you will get the the percentage pages found in the buffer cache. This percentage is calculated internally based on the total number of cache hits divided by the total number of cache lookups over the last few thousand page accesses.
This collector retrieves the two internal values separately. In order to calculate the Buffer Cache Hit Ratio in PromQL.

```
windows_mssql_bufman_buffer_cache_hits{instance="host:9182", exported_instance="MSSQLSERVER"} /
windows_mssql_bufman_buffer_cache_lookups{instance="host:9182", exported_instance="MSSQLSERVER"}
```

This principal can be used for following metrics too:
- AccessMethodsWorktablesFromCacheHitRatio
  - accessmethods_worktables_from_cache_hits
  - accessmethods_worktables_from_cache_lookups
- LogCacheHitRatio
  - databases_log_cache_hits
  - databases_log_cache_lookups
- AverageLockWaitTime
  - locks_wait_time_seconds
  - locks_count

## Alerting examples
_This collector does not yet have alerting examples, we would appreciate your help adding them!_

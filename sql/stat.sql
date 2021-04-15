-- SELECT * FROM polyswap.src_transactions;

use polyswap;

-- select DATE_FORMAT(FROM_UNIXTIME(create_time,"%Y%m%d"),"%Y%u")  weeks,count(caseid) count from tc_case group by weeks;

select DATE_FORMAT(FROM_UNIXTIME(`time`,"%Y%m%d"),"%Y%m%d") dates, count(*) from (
select * from src_transactions group by user
) group by dates;


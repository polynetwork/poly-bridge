use polyswap;

set @txhash='ba454d232fec8d1d89bd9ec3708bc65dff9dd2f69e8ce7410371e11fc1af0f15';
select * from wrapper_transactions where hash=@txhash;
select * from src_transactions where hash=@txhash;
select * from src_transfers where tx_hash=@txhash;
select *,@polyhash:=`hash` from poly_transactions where src_hash=@txhash;
select *,@dsthash:=`hash` from dst_transactions where poly_hash=@polyhash; 
select * from dst_transfers where tx_hash=@dsthash;

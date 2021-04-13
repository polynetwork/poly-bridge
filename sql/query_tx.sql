use polyswap;

-- query wrapper, src_transactions, src_transfers, poly_transactions, dst_transactions, dst_transfers
-- set src chain transaction hash as @txhash, e.g:
set @txhash='c83b1f0fd245d8719ffc865d06981de4ada0294c391c803fccf638b824f39d21';
select * from wrapper_transactions where hash=@txhash;
select * from src_transactions where hash=@txhash;
select * from src_transfers where tx_hash=@txhash;
select *,@polyhash:=hash from poly_transactions where src_hash=@txhash;
select *,@dsthash:=hash from dst_transactions where poly_hash=@polyhash;
select * from dst_transfers where tx_hash=@dsthash;
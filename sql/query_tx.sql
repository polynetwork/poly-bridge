use polyswap;

-- query wrapper, src_transactions, src_transfers, poly_transactions, dst_transactions, dst_transfers
-- set src chain transaction hash as @txhash, e.g:

set @txhash='af394e1a44d84aa573572b459e4fd43c377e167feeffde59c30d2de9aa4d6d8f';
select * from wrapper_transactions where hash=@txhash;
select * from src_transactions where hash=@txhash;
select * from src_transfers where tx_hash=@txhash;
select *,@polyhash:=`hash` from poly_transactions where src_hash=@txhash;
select *,@dsthash:=`hash` from dst_transactions where poly_hash=@polyhash; 
select * from dst_transfers where tx_hash=@dsthash;

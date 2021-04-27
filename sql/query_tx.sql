use polyswap;

-- query wrapper, src_transactions, src_transfers, poly_transactions, dst_transactions, dst_transfers
-- set src chain transaction hash as @txhash, e.g:

-- set @amount=138;
-- select * from src_transfers where amount=@amount and standard=0;
-- select * from dst_transfers where amount=@amount and standard=0;

set @txhash='1c7de65638b671b8c26443434497a3eb05891a4d379f97f760382adeacf0a057';
select * from wrapper_transactions where hash=@txhash;
select * from src_transactions where hash=@txhash;
select * from src_transfers where tx_hash=@txhash;
select *,@polyhash:=`hash` from poly_transactions where src_hash=@txhash;
select *,@dsthash:=`hash` from dst_transactions where poly_hash=@polyhash; 
select * from dst_transfers where tx_hash=@dsthash;

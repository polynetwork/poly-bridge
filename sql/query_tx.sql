use polyswap;

-- query wrapper, src_transactions, src_transfers, poly_transactions, dst_transactions, dst_transfers
-- set src chain transaction hash as @txhash, e.g:

-- set @amount=138;
-- select * from src_transfers where amount=@amount and standard=0;
-- select * from dst_transfers where amount=@amount and standard=0;

set @txhash='1d80e0552771d661fd53f29fc4ef63d836cdd246d8bf77765a36f253b59c4173';
select * from wrapper_transactions where hash=@txhash;
select * from src_transactions where hash=@txhash;
select * from src_transfers where tx_hash=@txhash;
select *,@polyhash:=`hash` from poly_transactions where src_hash=@txhash;
select *,@dsthash:=`hash` from dst_transactions where poly_hash=@polyhash; 
select * from dst_transfers where tx_hash=@dsthash;

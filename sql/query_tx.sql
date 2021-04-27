use polyswap;

-- query wrapper, src_transactions, src_transfers, poly_transactions, dst_transactions, dst_transfers
-- set src chain transaction hash as @txhash, e.g:

-- set @amount=138;
-- select * from src_transfers where amount=@amount and standard=0;
-- select * from dst_transfers where amount=@amount and standard=0;

-- select * from poly_transactions where hash='a9bedadc5a99e0e6112bff8108a7997c487b9797b212c32aa48809e8b7783db1';
-- select * from dst_transactions where hash='a9bedadc5a99e0e6112bff8108a7997c487b9797b212c32aa48809e8b7783db1';

set @txhash='c767b2004fcc0f9d8f21af81f84a84b425a997c3a2d6a8882e5ba997e656c7f8';
select * from wrapper_transactions where hash=@txhash;
select * from src_transactions where hash=@txhash;
select * from src_transfers where tx_hash=@txhash;
select *,@polyhash:=`hash` from poly_transactions where src_hash=@txhash;
select *,@dsthash:=`hash` from dst_transactions where poly_hash=@polyhash; 
select * from dst_transfers where tx_hash=@dsthash;

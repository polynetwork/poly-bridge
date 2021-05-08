use polyswap;

-- query wrapper, src_transactions, src_transfers, poly_transactions, dst_transactions, dst_transfers
-- set src chain transaction hash as @txhash, e.g:

-- set @amount=138;
-- select * from src_transfers where amount=@amount and standard=0;
-- select * from dst_transfers where amount=@amount and standard=0;

-- select * from poly_transactions where hash='a9bedadc5a99e0e6112bff8108a7997c487b9797b212c32aa48809e8b7783db1';
-- select * from dst_transactions where hash='a9bedadc5a99e0e6112bff8108a7997c487b9797b212c32aa48809e8b7783db1';

-- set @txhash='c767b2004fcc0f9d8f21af81f84a84b425a997c3a2d6a8882e5ba997e656c7f8';
-- select * from wrapper_transactions where hash=@txhash;
-- select * from src_transactions where hash=@txhash;
-- select * from src_transfers where tx_hash=@txhash;
-- select *,@polyhash:=`hash` from poly_transactions where src_hash=@txhash;
-- select *,@dsthash:=`hash` from dst_transactions where poly_hash=@polyhash; 
-- select * from dst_transfers where tx_hash=@dsthash;

explain SELECT src_transactions.hash as src_hash, poly_transactions.hash as poly_hash, 
dst_transactions.hash as dst_hash, src_transactions.chain_id as chain_id, src_transfers.asset as token_hash 
FROM `src_transactions` 
left join src_transfers on src_transactions.hash = src_transfers.tx_hash 
left join poly_transactions on src_transactions.hash = poly_transactions.src_hash 
left join dst_transactions on poly_transactions.hash = dst_transactions.poly_hash 
WHERE 
-- (src_transactions.hash = '1adf9495d2b677dde48ee9749e60d0c038b34e865cee7fb003e43cdabfab742c') 
(poly_transactions.hash = '1adf9495d2b677dde48ee9749e60d0c038b34e865cee7fb003e43cdabfab742c') 
or (dst_transactions.hash = '1adf9495d2b677dde48ee9749e60d0c038b34e865cee7fb003e43cdabfab742c') 
ORDER BY src_transactions.time desc;

-- explain select * from src_transactions where hash in ('1adf9495d2b677dde48ee9749e60d0c038b34e865cee7fb003e43cdabfab742c', 'adf9495d2b677dde48ee9749e60d0c038b34e865cee7fb003e43cdabfab742c');